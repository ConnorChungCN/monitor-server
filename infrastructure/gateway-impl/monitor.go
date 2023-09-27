package gatewayimpl

import (
	"context"
	"fmt"
	"time"

	"hanglok-tech.com/monitor-server/domain/model"
	"hanglok-tech.com/monitor-server/infrastructure/client"
	"hanglok-tech.com/monitor-server/infrastructure/logger"
)

// 暂定义查询结果为平均值
// type ResultByTaskId struct {
// 	TaskId           string
// 	AvgCPUPercent    float32
// 	AvgMemoryUsed    uint64
// 	AvgMemoryMaxUsed uint64
// 	// CudaVersion   string
// 	// AttachedGPUs  string
// 	// GPUsInfo      []*model.GPUInfo
// }

type MonitorGateway struct {
	InfluxDBClient  *client.InfluxDBClient
	SchedulerClient *client.SchedulerClient
}

func NewMonitorGateway(influxDBClient *client.InfluxDBClient, schedulerClient *client.SchedulerClient) (*MonitorGateway, error) {
	return &MonitorGateway{
		InfluxDBClient:  influxDBClient,
		SchedulerClient: schedulerClient,
	}, nil
}

func (obj *MonitorGateway) StorageInfo(ctx context.Context, workers []*model.TaskSysInfo) error {
	cpuDataPoints := make([]*model.StorgeDataPoint, len(workers))
	for i, v := range workers {
		logger.Logger.Infof("worker:%+v", v)
		cpuDataPoints[i] = &model.StorgeDataPoint{
			Tags: map[string]string{
				"AlgorithmName":    v.TaskSystemState.AlgorithmName,
				"AlgorithmVersion": v.TaskSystemState.AlgorithmVersion,
				"TaskId":           v.TaskSystemState.TaskId,
			},
			Fields: map[string]interface{}{
				"CPUPercent": v.TaskSystemState.CPUState.CPUPercent,
			},
			Timestamp: time.Now(),
		}
	}
	memoryDataPoints := make([]*model.StorgeDataPoint, len(workers))
	for i, v := range workers {
		memoryDataPoints[i] = &model.StorgeDataPoint{
			Tags: map[string]string{
				"AlgorithmName":    v.TaskSystemState.AlgorithmName,
				"AlgorithmVersion": v.TaskSystemState.AlgorithmVersion,
				"TaskId":           v.TaskSystemState.TaskId,
			},
			Fields: map[string]interface{}{
				"MemoryUsed":    v.TaskSystemState.MemoryState.MemoryUsed,
				"MemoryMaxUsed": v.TaskSystemState.MemoryState.MemoryMaxUsed,
			},
			Timestamp: time.Now(),
		}
	}
	gpuDataPoints := make([]*model.StorgeDataPoint, len(workers))
	for i, v := range workers {
		gpuDataPoints[i] = &model.StorgeDataPoint{
			Tags: map[string]string{
				"AlgorithmName":    v.TaskSystemState.AlgorithmName,
				"AlgorithmVersion": v.TaskSystemState.AlgorithmVersion,
				"TaskId":           v.TaskSystemState.TaskId,
			},
			Fields: map[string]interface{}{
				//TODO:gpu指标
				"CudaVersion":  v.TaskSystemState.GPUState.CudaVersion,
				"AttachedGPUs": v.TaskSystemState.GPUState.AttachedGPUs,
			},
			Timestamp: time.Now(),
		}
	}
	obj.InfluxDBClient.WriteData("containerCPUState", cpuDataPoints)
	obj.InfluxDBClient.WriteData("containerMemoryState", memoryDataPoints)
	obj.InfluxDBClient.WriteData("containerGPUState", gpuDataPoints)
	return nil
}

func (obj *MonitorGateway) FindByTaskId(ctx context.Context, taskId string) (*model.ResultByTaskId, error) {
	// 初始化一个 ResultByTaskId 结构体
	result := &model.ResultByTaskId{
		TaskId: taskId,
	}
	// 查询 containerCPUState 中的数据
	// 使用 WHERE "TaskId"='%s'会限定只有一个Series
	cpuQueryString := fmt.Sprintf(`SELECT "CPUPercent" FROM containerCPUState WHERE "TaskId"='%s'`, taskId)
	cpuRsp, err := obj.InfluxDBClient.QueryData(cpuQueryString)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	if len(cpuRsp.Results) > 0 && len(cpuRsp.Results[0].Series) > 0 {
		var cpuPercent []float32
		var totalCpuPercent float32 = 0
		for _, values := range cpuRsp.Results[0].Series[0].Values {
			// values 是一个 []interface{}，其中包含了每条记录的字段值
			// 将 values 中的字段值提取出来并进行相应的处理
			cpuPercent = append(cpuPercent, float32(values[1].(float64)))
			totalCpuPercent += float32(values[1].(float64))

		}
		result.AvgCPUPercent = totalCpuPercent / float32(len(cpuPercent))
	}
	// 查询 containerMemoryState 中的数据
	memoryQueryString := fmt.Sprintf(`SELECT "MemoryUsed", "MemoryMaxUsed" FROM containerMemoryState WHERE "TaskId"='%s'`, taskId)
	memoryRsp, err := obj.InfluxDBClient.QueryData(memoryQueryString)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	//memoryRsp.Results[0].Series[0].Columns字段名数组
	//memoryRsp.Results[0].Series[0].Values是[时间, memoryUsed的值, memoryMaxUsed的值]
	if len(memoryRsp.Results) > 0 && len(memoryRsp.Results[0].Series) > 0 {
		var memoryUsed []uint64
		var memoryMaxUsed []uint64
		var totalMemoryUsed uint64 = 0
		var totalmemoryMaxUsed uint64 = 0
		for _, values := range memoryRsp.Results[0].Series[0].Values {
			// values 是一个 []interface{}，其中包含了每条记录的字段值
			// 将 values 中的字段值提取出来并进行相应的处理
			memoryUsed = append(memoryUsed, uint64(values[1].(float64)))
			memoryMaxUsed = append(memoryMaxUsed, uint64(values[2].(float64)))
			totalMemoryUsed += uint64(values[1].(float64))
			totalmemoryMaxUsed += uint64(values[2].(float64))

		}
		result.AvgMemoryUsed = totalMemoryUsed / uint64(len(memoryUsed))
		result.AvgMemoryMaxUsed = totalmemoryMaxUsed / uint64(len(memoryMaxUsed))
	}
	// 查询 containerGPUState 中的数据
	// gpuQueryString := fmt.Sprintf(`SELECT "CudaVersion", "AttachedGPUs" FROM containerGPUState WHERE "TaskId"='%s'`, taskId)
	// gpuRsp, err := obj.InfluxDBClient.QueryData(gpuQueryString)
	// if err != nil {
	//     return nil, fmt.Errorf("query failed: %w", err)
	// }
	// if len(gpuRsp.Results) > 0 && len(gpuRsp.Results[0].Series) > 0 {
	//     result.CudaVersion = gpuRsp.Results[0].Series[0].Values[0][1].(string)
	//     result.AttachedGPUs = gpuRsp.Results[0].Series[0].Values[0][2].(string)
	// }

	return result, nil
}
