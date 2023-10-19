package executor

import "hanglok-tech.com/monitor-server/domain/model"

type SystemMetrics struct {
	CpuResult []*model.CpuStats
	MemResult []*model.MemoryStats
	GpuResult []*model.GpuInstanceStats
}
