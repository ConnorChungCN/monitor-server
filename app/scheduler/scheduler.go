package monitorScheduler

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hanglok-tech.com/monitor-server/adapter/grpc/proto/scheduler"
	"hanglok-tech.com/monitor-server/adapter/grpc/proto/worker"
	"hanglok-tech.com/monitor-server/domain/gateway"
	"hanglok-tech.com/monitor-server/domain/model"
	"hanglok-tech.com/monitor-server/infrastructure/client"
	"hanglok-tech.com/monitor-server/infrastructure/logger"
	"hanglok-tech.com/monitor-server/infrastructure/myerrors"
)

// 客户端连接失活时间5分钟
const WorkerClientTimeout = 60 * 5

type WorkerClient struct {
	conn         *grpc.ClientConn
	useTimestamp int64
	worker.WorkerClient
}

func NewWorkerClient(url string) (*WorkerClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connect to worker rpc server failed, %w", err)
	}
	client := worker.NewWorkerClient(conn)
	return &WorkerClient{
		conn:         conn,
		useTimestamp: time.Now().Unix(),
		WorkerClient: client,
	}, nil
}

func (cli *WorkerClient) Close() error {
	return cli.conn.Close()
}

func (cli *WorkerClient) IsTooOld(sec int64) bool {
	return time.Now().Unix()-cli.useTimestamp > sec
}

func (cli *WorkerClient) UpdateTimestamp() {
	cli.useTimestamp = time.Now().Unix()
}

type Monitor struct {
	MonitorManager  gateway.MonitorManager
	SchedulerClient *client.SchedulerClient
	workerClients   map[string]*WorkerClient
}

func NewMonitor(monitorManager gateway.MonitorManager, schedulerClient *client.SchedulerClient) *Monitor {
	return &Monitor{
		MonitorManager:  monitorManager,
		SchedulerClient: schedulerClient,
		workerClients:   make(map[string]*WorkerClient),
	}
}

func (obj *Monitor) getWorkerClient(ctx context.Context, url string) (*WorkerClient, error) {
	// 清理太久没用的client
	for k, c := range obj.workerClients {
		if c.IsTooOld(WorkerClientTimeout) {
			c.Close()
			delete(obj.workerClients, k)
		}
	}
	// 尝试获取
	cli, ok := obj.workerClients[url]
	if ok {
		cli.UpdateTimestamp()
		return cli, nil
	}
	cli, err := NewWorkerClient(url)
	if err != nil {
		return nil, fmt.Errorf("new worker client failed, %w", err)
	}
	obj.workerClients[url] = cli
	return cli, nil
}

// 获取一个worker的系统指标（cpu、memory、gpu）
func (obj *Monitor) getWorkerInfo(ctx context.Context, host string, port int64, taskId, algoName, algoVersion string) (*model.SystemState, error) {
	url := fmt.Sprintf("%s:%d", host, port)

	client, err := obj.getWorkerClient(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get worker client failed, %w", err)
	}

	rsp, err := client.GetTaskMetric(ctx, &worker.GetTaskMetricReq{})
	if err != nil {
		return nil, fmt.Errorf("grpc GetContainerStat fail, %w", err)
	}
	logger.Logger.Infof("rsp: %+v", rsp)

	cpuStats := rsp.GetCpuStats()
	if cpuStats == nil {
		return nil, myerrors.ErrTaskFinish
	}
	logger.Logger.Infof("cpuState: %+v\n", cpuStats)

	memoryStats := rsp.GetMemoryStats()
	if memoryStats == nil {
		return nil, myerrors.ErrTaskFinish
	}
	logger.Logger.Infof("memoryState: %+v\n", memoryStats)

	gpuStats := rsp.GetGpuStats()
	if gpuStats == nil {
		return nil, myerrors.ErrTaskFinish
	}
	logger.Logger.Infof("gpuStats: %+v\n", gpuStats)
	cudaVersion := gpuStats.CudaVersion
	var gpuInfo []*model.GpuInstanceStats
	for _, v := range gpuStats.GpuInstanceStats {
		gpuInfo = append(gpuInfo, &model.GpuInstanceStats{
			TaskId:           taskId,
			AlgorithmName:    algoName,
			AlgorithmVersion: algoVersion,
			CudaVersion:      cudaVersion,
			Id:               v.Id,
			ProductName:      v.ProductName,
			GpuUsage:         float64(v.GpuUsage),
			MemoryUsage:      float64(v.MemoryUsage),
			MemoryUsed:       int64(v.MemoryUsed),
			MemoryFree:       int64(v.MemoryFree),
		})
	}

	return &model.SystemState{
		CpuStats: &model.CpuStats{
			TaskId:           taskId,
			AlgorithmName:    algoName,
			AlgorithmVersion: algoVersion,
			CpuUsage:         float64(cpuStats.Usage),
		},
		MemoryStats: &model.MemoryStats{
			TaskId:           taskId,
			AlgorithmName:    algoName,
			AlgorithmVersion: algoVersion,
			Usage:            float64(memoryStats.Usage),
			Used:             int64(memoryStats.Used),
			Free:             int64(memoryStats.Free),
		},
		GpuStats: gpuInfo,
	}, nil
}

// 获取所有worker的系统指标
func (obj *Monitor) GetBusyWorkerInfo(ctx context.Context) ([]*model.SystemState, error) {
	rsp, err := obj.SchedulerClient.Client.ListWorkers(ctx, &scheduler.ListWorkerReq{})
	if err != nil {
		return nil, fmt.Errorf("grpc ListTask failed: %w", err)
	}
	retWorkers := make([]*model.SystemState, 0)
	for _, v := range rsp.Workers {
		//如果worker不在运行则跳出本次循环
		if v.GetRunningTask().GetTaskId() == "" {
			continue
		}
		port := v.Port
		host := v.Host
		//grpc调用worker获取系统指标
		systemstate, err := obj.getWorkerInfo(ctx, host, port,
			v.RunningTask.TaskId, v.RunningTask.Algorithm.Name, v.RunningTask.Algorithm.Version)
		if err != nil {
			return nil, fmt.Errorf("grpc GetContainerStat failed: %w", err)
		}
		retWorkers = append(retWorkers, systemstate)
	}
	if len(retWorkers) == 0 {
		logger.Logger.Infof("no Info need to persist")
		return nil, nil
	}
	return retWorkers, nil
}

func (obj *Monitor) StartMonitoring(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	//定时器
	for {
		select {
		case <-ticker.C:
			go func(ctx context.Context) {
				workers, err := obj.GetBusyWorkerInfo(ctx)
				if err != nil {
					logger.Logger.Errorf("no worker running :%s", err)
					return
				}
				// 持久化数据
				err = obj.MonitorManager.StorageMetrics(ctx, workers)
				if err != nil {
					logger.Logger.Errorf("Storage Info failed:%s", err)
				}
			}(ctx)
		case <-ctx.Done():
			return
		}

	}
}
