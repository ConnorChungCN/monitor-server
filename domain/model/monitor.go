package model

type InquireResult struct {
	TaskId           string
	AlgorithmName    string
	AlgorithmVersion string
	CpuInquireResult []*CpuInquire
}

type CpuInquire struct {
	Time       string
	CpuPercent float64
}
