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
	//TODO:增加Memory、GPU
	var cpuInquire []*monitor.CpuInquire
	for _, i := range findResult.CpuInquireResult {
		cpuInquire = append(cpuInquire, &monitor.CpuInquire{
			Time:       i.Time,
			CpuPercent: float32(i.CpuPercent),
		})
	}
	return &monitor.FindTaskInfoByIdRsp{
		CpuInquire: cpuInquire,
	}, nil
}
