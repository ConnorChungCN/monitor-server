package gateway

import (
	"context"

	"hanglok-tech.com/monitor-server/domain/model"
)

type MonitorManager interface {
	//持久化数据
	StorageInfo(ctx context.Context, workers []*model.SystemState) error
	//查询所有系统指标
	QuerySummary(ctx context.Context, taskId string) (*model.QueryAllTaskInfo, error)
	//查询平均系统指标
	QueryAvg(ctx context.Context, taskId string) (*model.QueryAvgTaskInfo, error)
}
