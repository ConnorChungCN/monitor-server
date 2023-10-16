package model

type QueryAvgTaskInfo struct {
	TaskId           string
	AlgorithmName    string
	AlgorithmVersion string
	AvgCPUPercent    float32
	AvgMemoryUsage   float32
	AvgMemoryUsed    int64
	AvgMemoryFree    int64

	AvgGpuUsage       float32
	AvgGpuMemoryUsage float32
	AvgGpuMemoryUsed  int64
	AvgGpuMemoryFree  int64
}

type QueryAllTaskInfo struct {
	TaskId           string
	AlgorithmName    string
	AlgorithmVersion string
	CpuResult        []*QueryCpuInfo
	MemResult        []*QueryMemInfo
	GpuResult        []*QueryGpuInfo
}

type QueryCpuInfo struct {
	Time       string
	CpuPercent float64
}

type QueryMemInfo struct {
	Time  string
	Usage float64
	Used  int64
	Free  int64
}

type QueryGpuInfo struct {
	Time        string
	Id          string
	ProductName string
	GpuUsage    float64
	MemoryUsage float64
	MemoryUsed  int64
	MemoryFree  int64
}

type GpuInstance struct {
	Time        string
	Id          string
	ProductName string
	GpuUsage    float64
	MemoryUsage float64
	MemoryUsed  int64
	MemoryFree  int64
}
