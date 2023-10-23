package gateway

import (
	"context"

	"hanglok-tech.com/monitor-server/domain/model"
)

type MonitorManager interface {
	//持久化数据
	StorageMetrics(ctx context.Context, workers []*model.SystemState) error
	//查询所有系统指标
	QueryCpuStatsByTask(ctx context.Context, taskId string) ([]*model.CpuStats, error)
	QueryMemStatsByTask(ctx context.Context, taskId string) ([]*model.MemoryStats, error)
	QueryGpuStatsByTask(ctx context.Context, taskId string) ([]*model.GpuInstanceStats, error)
	//查询平均系统指标
	QuerySummaryByTask(ctx context.Context, taskId string) (*model.Summary, error)
}
