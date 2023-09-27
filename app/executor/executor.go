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

func (obj *Executor) FindTaskInfoById(ctx context.Context, taskId string) (*model.Summary, error) {
	findResult, err := obj.MonitorManager.FindSummaryByTaskId(ctx, taskId)
	if err != nil {
		logger.Logger.Errorf("FindByTaskId fail: %s", err)
		return nil, fmt.Errorf("FindByTaskId fail: %s", err)
	}
	return findResult, nil
}
