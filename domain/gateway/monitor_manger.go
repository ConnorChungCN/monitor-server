package gateway

import (
	"context"

	"hanglok-tech.com/monitor-server/domain/model"
)

type MonitorManager interface {
	//持久化数据
	StorageInfo(ctx context.Context, workers []*model.SystemState) error
	//查找持久化数据目前只通过taskId寻找
	FindSummaryByTaskId(ctx context.Context, taskId string) (*model.InquireResult, error)
}
