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
)

type Monitor struct {
	MonitorManager  gateway.MonitorManager
	SchedulerClient *client.SchedulerClient
}

// 获取一个worker的系统指标
func getWorkerInfo(ctx context.Context, host string, port int64) (*model.SystemState, error) {
	url := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connect to worker rpc server failed, %w", err)
	}
	defer conn.Close()
	client := worker.NewWorkerClient(conn)
	rsp, err := client.GetContainerStat(ctx, &worker.GetContainerStatReq{})
	if err != nil {
		return nil, fmt.Errorf("grpc GetContainerStat fail, %w", err)
	}
	cpuState := &model.CPUState{
		CPUPercent: rsp.CpuPercent,
	}
	memoryState := &model.MemoryState{
		MemoryUsed:    rsp.MemoryUsage,
		MemoryMaxUsed: rsp.MemoryMaxUsage,
	}
	gpuState := &model.GPUState{
		//TODO:GPU
	}
	return &model.SystemState{
		CPUState:    cpuState,
		MemoryState: memoryState,
		GPUState:    gpuState,
	}, nil
}

func NewMonitor(monitorManager gateway.MonitorManager, schedulerClient *client.SchedulerClient) *Monitor {
	return &Monitor{
		MonitorManager:  monitorManager,
		SchedulerClient: schedulerClient,
	}
}

// 获取所有worker的系统指标
func (obj *Monitor) GetAllWorkerInfo(ctx context.Context) ([]*model.TaskSysInfo, error) {
	rsp, err := obj.SchedulerClient.Client.ListWorkers(ctx, &scheduler.ListWorkerReq{})
	if err != nil {
		return nil, fmt.Errorf("grpc ListTask failed: %w", err)
	}
	retWorkers := make([]*model.TaskSysInfo, 0)
	for i, v := range rsp.Workers {
		//如果worker不在运行则跳出本次循环
		logger.Logger.Infof("work: %+v", v)

		if v.GetRunningTask().GetTaskId() == "" {
			continue
		}
		retWorkers = append(retWorkers, &model.TaskSysInfo{
			RunningWorkerPort: v.Port,
			RunningWorkerHost: v.Host,
		})
		//grpc调用worker获取系统指标
		systemstate, err := getWorkerInfo(ctx, v.Host, v.Port)
		if err != nil {
			return nil, fmt.Errorf("grpc GetContainerStat failed: %w", err)
		}
		logger.Logger.Infof("systemstate: %v", systemstate)
		systemstate.AlgorithmName = v.RunningTask.AlgorithmName
		systemstate.AlgorithmVersion = v.RunningTask.AlgorithmVersion
		systemstate.TaskId = v.RunningTask.TaskId
		retWorkers[i].TaskSystemState = systemstate
	}
	return retWorkers, nil
}

// 持久化数据
func (obj *Monitor) PersistenceInfo(ctx context.Context, workers []*model.TaskSysInfo) error {
	if len(workers) == 0 {
		logger.Logger.Infof("no Info need to persist")
		return nil
	}
	err := obj.MonitorManager.StorageInfo(ctx, workers)
	if err != nil {
		return fmt.Errorf("StorageInfo failed, %w", err)
	}
	return nil
}

func (obj *Monitor) UpdateInfo(ctx context.Context) error {
	// 调用 GetInfo 方法获取系统指标
	workers, err := obj.GetAllWorkerInfo(ctx)
	if err != nil {
		return fmt.Errorf("GetInfo failed: %w", err)
	}
	// 持久化数据
	err = obj.PersistenceInfo(ctx, workers)
	if err != nil {
		return fmt.Errorf("PersistenceInfo failed: %w", err)
	}
	return nil
}

func (obj *Monitor) StartMonitoring(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	//定时器
	for {
		select {
		case <-ticker.C:
			err := obj.UpdateInfo(ctx)
			if err != nil {
				logger.Logger.Errorf("UpdateInfo failed: %s", err)
			}
		case <-ctx.Done():
			return
		}

	}
}
