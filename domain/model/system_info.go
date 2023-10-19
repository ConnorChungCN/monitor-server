package model

import "time"

type SystemState struct {
	CpuStats    *CpuStats
	MemoryStats *MemoryStats
	GpuStats    []*GpuInstanceStats
}

type CpuStats struct {
	TaskId           string
	AlgorithmName    string
	AlgorithmVersion string
	Time             time.Time
	CpuUsage         float64
}

type MemoryStats struct {
	TaskId           string
	AlgorithmName    string
	AlgorithmVersion string
	Time             time.Time
	Usage            float64
	Used             int64
	Free             int64
}

type GpuInstanceStats struct {
	TaskId           string
	AlgorithmName    string
	AlgorithmVersion string
	Time             time.Time
	CudaVersion      string
	Id               string
	ProductName      string
	GpuUsage         float64
	MemoryUsage      float64
	MemoryUsed       int64
	MemoryFree       int64
}
