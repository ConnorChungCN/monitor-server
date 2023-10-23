package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hanglok-tech.com/monitor-server/adapter/grpc/proto/monitor"
	"hanglok-tech.com/monitor-server/app/executor"
)

const (
	timeFormat = "2006-01-02 15:04:05"
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

// TODO：对时间信息进行转换——>年月日
func (obj *MonitorServer) QueryAllTaskInfo(ctx context.Context, req *monitor.QueryAllTaskInfoReq) (*monitor.QueryAllTaskInfoRsp, error) {
	findResult, err := obj.Monitor.QuerySystemMetrics(ctx, req.TaskId)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Find Task Info By Id failed, %w", err)
	}
	var cpuInquire []*monitor.CpuQuery
	var memInquire []*monitor.MemQuery
	var gpuInquire []*monitor.GpuQuery
	for _, i := range findResult.CpuResult {
		cpuInquire = append(cpuInquire, &monitor.CpuQuery{
			Time:       i.Time.Format(timeFormat),
			CpuPercent: float32(i.CpuUsage),
		})
	}
	for _, i := range findResult.MemResult {
		memInquire = append(memInquire, &monitor.MemQuery{
			Time:     i.Time.Format(timeFormat),
			MemUsage: float32(i.Usage),
			MemUsed:  i.Used,
			MemFree:  i.Free,
		})
	}
	//GPU proto rsp
	for _, i := range findResult.GpuResult {
		gpuInquire = append(gpuInquire, &monitor.GpuQuery{
			Time:        i.Time.Format(timeFormat),
			Id:          i.Id,
			ProdctName:  i.ProductName,
			GpuUsage:    float32(i.GpuUsage),
			GpuMemUsage: float32(i.MemoryUsage),
			GpuMemUsed:  i.MemoryUsed,
			GpuMemFree:  i.MemoryFree,
		})
	}
	return &monitor.QueryAllTaskInfoRsp{
		CpuQuery: cpuInquire,
		MemQuery: memInquire,
		GpuQuery: gpuInquire,
	}, nil
}

func (obj *MonitorServer) QueryAvgTaskInfo(ctx context.Context, req *monitor.QueryAvgTaskInfoReq) (*monitor.QueryAvgTaskInfoRsp, error) {
	findResult, err := obj.Monitor.QuerySummary(ctx, req.TaskId)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Find Task Info By Id failed, %w", err)
	}
	return &monitor.QueryAvgTaskInfoRsp{
		AvgCPUPercent: findResult.AvgCPUPercent,

		AvgMemoryUsage: findResult.AvgMemoryUsage,
		AvgMemoryUsed:  findResult.AvgMemoryUsed,
		AvgMemoryFree:  findResult.AvgMemoryFree,

		AvgGpuUsage:       findResult.AvgGpuUsage,
		AvgGpuMemoryUsage: findResult.AvgGpuMemoryUsage,
		AvgGpuMemoryUsed:  findResult.AvgGpuMemoryUsed,
		AvgGpuMemoryFree:  findResult.AvgGpuMemoryFree,
	}, nil
}
