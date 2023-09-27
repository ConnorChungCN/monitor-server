package model

type TaskSysInfo struct {
	TaskSystemState   *SystemState
	RunningWorkerPort int64
	RunningWorkerHost string
}

type ResultByTaskId struct {
	TaskId           string
	AvgCPUPercent    float32
	AvgMemoryUsed    uint64
	AvgMemoryMaxUsed uint64
	// CudaVersion   string
	// AttachedGPUs  string
	// GPUsInfo      []*model.GPUInfo
}
