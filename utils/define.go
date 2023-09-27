package utils

import "time"

const (
	// 请求状态
	ErrCodeOther        = 9
	ErrCodeIncorrectPsd = 10
	ErrCodeNoUser       = 11
	// 错误信息
	ErrMsgIncorrectPsd = "Incorrect password"
	ErrMsgNoUser       = "User not exists"
)

const (
	WorkerSurvival = time.Second * 60
	AlgoTimeout    = time.Minute * 10
)

// worker状态
const (
	WorkerWaiting = iota
	WorkerRunning
)

const (
	TaskStatusNotRun = iota
	TaskStatusRunning
	TaskStatusStop
)

// 任务状态码
const (
	TaskStateSuccess = iota
	TaskStateTimeout
	TaskStateFailed
	TaskStateExecption
	TaskStateStop
)

// pipeline status
const (
	PipelineStatusNotStart = iota
	PipelineStatusRunning
	PipelineStatusBlocked
	PipelineStatusSuccess
	PipelineStatusFail
)

const (
	PipelineQueueName = "pipeline_task"
)
