package executor

import (
	"context"
	"fmt"

	"hanglok-tech.com/monitor-server/domain/gateway"
	"hanglok-tech.com/monitor-server/domain/model"
	"hanglok-tech.com/monitor-server/infrastructure/logger"
)

// 实现service的功能，grpc调用tsdb获取系统指标
type Executor struct {
	MonitorManager gateway.MonitorManager
}

func NewExecutor(monitorManager gateway.MonitorManager) *Executor {
	return &Executor{
		MonitorManager: monitorManager,
	}
}

func (obj *Executor) QuerySystemMetrics(ctx context.Context, taskId string) (*SystemMetrics, error) {
	cpuResult, err := obj.MonitorManager.QueryCpuStatsByTask(ctx, taskId)
	if err != nil {
		logger.Logger.Errorf("query cpu stats fail: %s", err)
		return nil, fmt.Errorf("query cpu stats fail: %s", err)
	}
	memResult, err := obj.MonitorManager.QueryMemStatsByTask(ctx, taskId)
	if err != nil {
		logger.Logger.Errorf("query mem stats fail: %s", err)
		return nil, fmt.Errorf("query mem stats fail: %s", err)
	}
	gpuResult, err := obj.MonitorManager.QueryGpuStatsByTask(ctx, taskId)
	if err != nil {
		logger.Logger.Errorf("query gpu stats fail: %s", err)
		return nil, fmt.Errorf("query gpu stats fail: %s", err)
	}
	return &SystemMetrics{
		CpuResult: cpuResult,
		MemResult: memResult,
		GpuResult: gpuResult,
	}, nil
}

func (obj *Executor) QuerySummary(ctx context.Context, taskId string) (*model.Summary, error) {
	findResult, err := obj.MonitorManager.QuerySummaryByTask(ctx, taskId)
	if err != nil {
		logger.Logger.Errorf("FindByTaskId fail: %s", err)
		return nil, fmt.Errorf("FindByTaskId fail: %s", err)
	}
	return findResult, nil
}
