package model

type SystemState struct {
	TaskId           string
	AlgorithmName    string
	AlgorithmVersion string
	CpuStats         *CpuStats
	MemoryStats      *MemoryStats
	GpuStats         *GpuStats
}

type CpuStats struct {
	CPUPercent float32
}

type MemoryStats struct {
	Usage float64
	Used  int64
	Free  int64
}

type GpuStats struct {
	CudaVersion string
	GPUsInfo    []*GpuInstanceStats
}

type GpuInstanceStats struct {
	Id          string
	ProductName string
	GpuUsage    float64
	MemoryUsage float64
	MemoryUsed  int64
	MemoryFree  int64
}
