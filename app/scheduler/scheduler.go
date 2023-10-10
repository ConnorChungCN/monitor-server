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

type Monitor struct {
	MonitorManager  gateway.MonitorManager
	SchedulerClient *client.SchedulerClient
}

func NewMonitor(monitorManager gateway.MonitorManager, schedulerClient *client.SchedulerClient) *Monitor {
	return &Monitor{
		MonitorManager:  monitorManager,
		SchedulerClient: schedulerClient,
	}
}

// 获取一个worker的系统指标（cpu、memory）
func (obj *Monitor) getWorkerInfo(ctx context.Context, host string, port int64) (*model.SystemState, error) {
	url := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connect to worker rpc server failed, %w", err)
	}
	defer conn.Close()
	client := worker.NewWorkerClient(conn)
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
	var gpuInfo []*model.GpuInstanceStats
	for _, v := range gpuStats.GpuInstanceStats {
		gpuInfo = append(gpuInfo, &model.GpuInstanceStats{
			Id:          v.Id,
			ProductName: v.ProductName,
			GpuUsage:    float64(v.GpuUsage),
			MemoryUsage: float64(v.MemoryUsage),
			MemoryUsed:  int64(v.MemoryUsed),
			MemoryFree:  int64(v.MemoryFree),
		})
	}

	return &model.SystemState{
		CpuStats: &model.CpuStats{
			CPUPercent: cpuStats.GetUsage(),
		},
		MemoryStats: &model.MemoryStats{
			Usage: float64(rsp.MemoryStats.Usage),
			Used:  int64(rsp.MemoryStats.Used),
			Free:  int64(rsp.MemoryStats.Free),
		},
		GpuStats: &model.GpuStats{
			CudaVersion: rsp.GpuStats.CudaVersion,
			GPUsInfo:    gpuInfo,
		},
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
		systemstate, err := obj.getWorkerInfo(ctx, host, port)
		if err != nil {
			return nil, fmt.Errorf("grpc GetContainerStat failed: %w", err)
		}
		systemstate.AlgorithmName = v.RunningTask.Algorithm.Name
		systemstate.AlgorithmVersion = v.RunningTask.Algorithm.Version
		systemstate.TaskId = v.RunningTask.TaskId
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
			workers, err := obj.GetBusyWorkerInfo(ctx)
			if err != nil {
				logger.Logger.Errorf("no worker running :%s", err)
				break
			}
			// 持久化数据
			err = obj.MonitorManager.StorageInfo(ctx, workers)
			if err != nil {
				logger.Logger.Errorf("Storage Info failed:%s", err)
			}
		case <-ctx.Done():
			return
		}

	}
}
