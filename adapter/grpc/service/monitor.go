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

func (obj *MonitorServer) QueryAllTaskInfo(ctx context.Context, req *monitor.QueryAllTaskInfoReq) (*monitor.QueryAllTaskInfoRsp, error) {
	findResult, err := obj.Monitor.QuerySummary(ctx, req.TaskId)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Find Task Info By Id failed, %w", err)
	}
	//TODO:增加Memory、GPU
	var cpuInquire []*monitor.CpuQuery
	var memInquire []*monitor.MemQuery
	for _, i := range findResult.CpuResult {
		cpuInquire = append(cpuInquire, &monitor.CpuQuery{
			Time:       i.Time,
			CpuPercent: float32(i.CpuPercent),
		})
	}
	for _, i := range findResult.MemResult {
		memInquire = append(memInquire, &monitor.MemQuery{
			Time:     i.Time,
			MemUsage: float32(i.Usage),
			MemUsed:  i.Used,
			MemFree:  i.Free,
		})
	}
	return &monitor.QueryAllTaskInfoRsp{
		CpuQuery: cpuInquire,
		MemQuery: memInquire,
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
