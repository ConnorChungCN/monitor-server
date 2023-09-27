package model

type Summary struct {
	TaskId           string
	AvgCPUPercent    float32
	AvgMemoryUsed    int64
	AvgMemoryMaxUsed int64
}
