package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hanglok-tech.com/monitor-server/adapter/grpc/proto/monitor"
	"hanglok-tech.com/monitor-server/app/executor"
)

type MonitorServer struct {
	monitor.UnimplementedMonitorServer
	Monitor *executor.Executor
}

func NewMonitorServer(monitor *executor.Executor) *MonitorServer {
	return &MonitorServer{
		Monitor: monitor,
	}
}

func (obj *MonitorServer) FindTaskInfoById(ctx context.Context, req *monitor.FindTaskInfoByIdReq) (*monitor.FindTaskInfoByIdRsp, error) {
	findResult, err := obj.Monitor.FindTaskInfoById(ctx, req.TaskId)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Find Task Info By Id failed, %w", err)
	}
	return &monitor.FindTaskInfoByIdRsp{
		AvgCPUPercent:    findResult.AvgCPUPercent,
		AvgMemoryUsed:    findResult.AvgMemoryUsed,
		AvgMemoryMaxUsed: findResult.AvgMemoryMaxUsed,
	}, nil
}
