package model

type Summary struct {
	//cpu平均使用率
	AvgCPUPercent float32
	//memeory平均使用率、平均使用字节、平均空闲字节
	AvgMemoryUsage float32
	AvgMemoryUsed  int64
	AvgMemoryFree  int64
	//gpu平均使用率、gpumMmeory平均使用率、gpumMmeory平均使用字节、gpumMmeory平均空闲字节
	AvgGpuUsage       float32
	AvgGpuMemoryUsage float32
	AvgGpuMemoryUsed  int64
	AvgGpuMemoryFree  int64
}
