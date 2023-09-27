package model

type SystemState struct {
	TaskId           string
	AlgorithmName    string
	AlgorithmVersion string
	CPUState         *CPUState
	MemoryState      *MemoryState
	GPUState         *GPUState
}

type CPUState struct {
	CPUPercent float32
}

type MemoryState struct {
	MemoryUsed    uint64
	MemoryMaxUsed uint64
}

type GPUState struct {
	CudaVersion  string
	AttachedGPUs string
	GPUsInfo     []*GPUInfo
}

type GPUInfo struct {
	ProductName string
	MemoryTotal float64
	MemoryUsed  float64
	MemoryFree  float64
	GPUUsage    float64
	MemoryUsage float64
}
