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

// 目前proto的定义不完善
func (obj *MonitorServer) QueryAllTaskInfo(ctx context.Context, req *monitor.QueryAllTaskInfoReq) (*monitor.QueryAllTaskInfoRsp, error) {
	findResult, err := obj.Monitor.QuerySummary(ctx, req.TaskId)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Find Task Info By Id failed, %w", err)
	}
	//TODO:增加Memory、GPU
	var cpuInquire []*monitor.CpuInquire
	for _, i := range findResult.CpuResult {
		cpuInquire = append(cpuInquire, &monitor.CpuInquire{
			Time:       i.Time,
			CpuPercent: float32(i.CpuPercent),
		})
	}
	return &monitor.QueryAllTaskInfoRsp{
		CpuInquire: cpuInquire,
	}, nil
}

func (obj *MonitorServer) QueryAvgTaskInfo(ctx context.Context, req *monitor.QueryAvgTaskInfoReq) (*monitor.QueryAvgTaskInfoRsp, error) {
	findResult, err := obj.Monitor.QueryAvg(ctx, req.TaskId)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Find Task Info By Id failed, %w", err)
	}
	//TODO:增加Memory、GPU
	return &monitor.QueryAvgTaskInfoRsp{
		AvgCPUPercent:  findResult.AvgCPUPercent,
		AvgMemoryUsage: findResult.AvgMemoryUsage,
		AvgMemoryUsed:  findResult.AvgMemoryUsed,
		AvgMemoryFree:  findResult.AvgMemoryFree,
	}, nil
}
