package model

type QueryAvgTaskInfo struct {
	TaskId         string
	AvgCPUPercent  float32
	AvgMemoryUsage float32
	AvgMemoryUsed  int64
	AvgMemoryFree  int64
}

type QueryAllTaskInfo struct {
	TaskId           string
	AlgorithmName    string
	AlgorithmVersion string
	CpuResult        []*QueryCpuInfo
	MemResult        []*QueryMemInfo
	GpuResylt        []*QueryGpuInfo
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
	CudaVersion string
	GpuInstance []*GpuInstance
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
