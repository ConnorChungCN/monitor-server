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
}

type QueryCpuInfo struct {
	Time       string
	CpuPercent float64
}
