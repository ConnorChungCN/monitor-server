package gatewayimpl

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"hanglok-tech.com/monitor-server/domain/model"
	"hanglok-tech.com/monitor-server/infrastructure/client"
)

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

func (obj *MonitorGateway) StorageInfo(ctx context.Context, workers []*model.SystemState) error {
	if workers == nil {
		return nil
	}
	cpuDataPoints := make([]*client.DataPoint, len(workers))
	memoryDataPoints := make([]*client.DataPoint, len(workers))
	gpuDataPoints := make([]*client.DataPoint, len(workers))
	for i, v := range workers {
		cpuDataPoints[i] = &client.DataPoint{
			Tags: map[string]string{
				"AlgorithmName":    v.AlgorithmName,
				"AlgorithmVersion": v.AlgorithmVersion,
				"TaskId":           v.TaskId,
			},
			Fields: map[string]interface{}{
				"CPUPercent": v.CpuStats.CPUPercent,
			},
			Timestamp: time.Now(),
		}
		memoryDataPoints[i] = &client.DataPoint{
			Tags: map[string]string{
				"AlgorithmName":    v.AlgorithmName,
				"AlgorithmVersion": v.AlgorithmVersion,
				"TaskId":           v.TaskId,
			},
			Fields: map[string]interface{}{
				"MemoryUsage": v.MemoryStats.Usage,
				"MemoryUsed":  v.MemoryStats.Used,
				"MemoryFree":  v.MemoryStats.Free,
			},
			Timestamp: time.Now(),
		}
		for j := 0; j < len(v.GpuStats.GPUsInfo); j++ {
			gpuDataPoints[i] = &client.DataPoint{
				Tags: map[string]string{
					"AlgorithmName":    v.AlgorithmName,
					"AlgorithmVersion": v.AlgorithmVersion,
					"TaskId":           v.TaskId,
				},
				Fields: func() map[string]interface{} {
					fields := make(map[string]interface{})
					fields["CudaVersion"] = v.GpuStats.CudaVersion
					fields[fmt.Sprintf("Id-%d", j)] = v.GpuStats.GPUsInfo[j].Id
					fields[fmt.Sprintf("ProductName-%d", j)] = v.GpuStats.GPUsInfo[j].ProductName
					fields[fmt.Sprintf("GpuUsage-%d", j)] = v.GpuStats.GPUsInfo[j].GpuUsage
					fields[fmt.Sprintf("MemoryUsage-%d", j)] = v.GpuStats.GPUsInfo[j].MemoryUsage
					fields[fmt.Sprintf("MemoryUsed-%d", j)] = v.GpuStats.GPUsInfo[j].MemoryUsed
					fields[fmt.Sprintf("MemoryFree-%d", j)] = v.GpuStats.GPUsInfo[j].MemoryFree
					return fields
				}(),
				Timestamp: time.Now(),
			}
		}
	}
	obj.InfluxDBClient.WriteData("containerCPUState", cpuDataPoints)
	obj.InfluxDBClient.WriteData("containerMemState", memoryDataPoints)
	obj.InfluxDBClient.WriteData("containerGPUState", gpuDataPoints)
	return nil
}

func (obj *MonitorGateway) QuerySummary(ctx context.Context, taskId string) (*model.QueryAllTaskInfo, error) {
	// 初始化一个 ResultByTaskId 结构体
	result := &model.QueryAllTaskInfo{
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
		var cpuInquire []*model.QueryCpuInfo
		for _, values := range cpuRsp.Results[0].Series[0].Values {
			// values 是一个 []interface{}，其中包含了每条记录的字段值
			// 将 values 中的字段值提取出来并进行相应的处理。values[0]是时间戳，value[1]是CPUPercent的值
			cpuFloat, _ := values[1].(json.Number).Float64()
			timeString := values[0].(json.Number).String()
			timestamp, _ := time.Parse(time.RFC3339, timeString)
			value := cpuFloat
			cpuInquire = append(cpuInquire, &model.QueryCpuInfo{
				Time:       timestamp.String(),
				CpuPercent: value,
			})
		}
		result.CpuResult = cpuInquire
	}
	// // 查询 containerMemoryState 中的数据
	// memoryQueryString := fmt.Sprintf(`SELECT "MemoryFree", "MemoryUsage", "MemoryUsed" FROM containerMemState WHERE "TaskId"='%s'`, taskId)
	// memoryRsp, err := obj.InfluxDBClient.QueryData(memoryQueryString)
	// if err != nil {
	// 	return nil, fmt.Errorf("query failed: %w", err)
	// }
	// //memoryRsp.Results[0].Series[0].Columns字段名数组
	// //memoryRsp.Results[0].Series[0].Values是[时间, memoryUsed的值, memoryMaxUsed的值]
	// if len(memoryRsp.Results) > 0 && len(memoryRsp.Results[0].Series) > 0 {
	// 	var memoryUsed []int64
	// 	var memoryMaxUsed []int64
	// 	var totalMemoryUsed int64 = 0
	// 	var totalmemoryMaxUsed int64 = 0
	// 	for _, values := range memoryRsp.Results[0].Series[0].Values {
	// 		// values 是一个 []interface{}，其中包含了每条记录的字段值
	// 		// 将 values 中的字段值提取出来并进行相应的处理
	// 		memoryUsed = append(memoryUsed, int64(values[1].(float64)))
	// 		memoryMaxUsed = append(memoryMaxUsed, int64(values[2].(float64)))
	// 		totalMemoryUsed += int64(values[1].(float64))
	// 		totalmemoryMaxUsed += int64(values[2].(float64))
	// 	}
	// 	result.AvgMemoryUsed = totalMemoryUsed / int64(len(memoryUsed))
	// 	result.AvgMemoryMaxUsed = totalmemoryMaxUsed / int64(len(memoryMaxUsed))
	// }
	// // 查询 containerGPUState 中的数据
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

func (obj *MonitorGateway) QueryAvg(ctx context.Context, taskId string) (*model.QueryAvgTaskInfo, error) {
	// 初始化一个 ResultByTaskId 结构体
	result := &model.QueryAvgTaskInfo{
		TaskId: taskId,
	}
	cpuQueryString := fmt.Sprintf(`SELECT "CPUPercent" FROM containerCPUState WHERE "TaskId"='%s'`, taskId)
	cpuRsp, err := obj.InfluxDBClient.QueryData(cpuQueryString)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	if len(cpuRsp.Results) > 0 && len(cpuRsp.Results[0].Series) > 0 {
		var cpuPercent []float32
		var totalCpuPercent float32 = 0
		for _, values := range cpuRsp.Results[0].Series[0].Values {
			cpuFloat, _ := values[1].(json.Number).Float64()
			cpuPercent = append(cpuPercent, float32(cpuFloat))
			totalCpuPercent += float32(cpuFloat)
		}
		result.AvgCPUPercent = totalCpuPercent / float32(len(cpuPercent))
	}
	// 查询 containerMemoryState 中的数据
	memoryQueryString := fmt.Sprintf(`SELECT "MemoryUsed", "MemoryFree", "MemoryUsage" FROM containerMemState WHERE "TaskId"='%s'`, taskId)
	memoryRsp, err := obj.InfluxDBClient.QueryData(memoryQueryString)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	if len(memoryRsp.Results) > 0 && len(memoryRsp.Results[0].Series) > 0 {
		var memoryUsed []int64
		var memoryFree []int64
		var memoryUsage []float32
		var totalMemoryUsed int64 = 0
		var totalmemoryFree int64 = 0
		var totalmemoryUsage float32 = 0
		for _, values := range memoryRsp.Results[0].Series[0].Values {
			memoryUsedFloat, _ := values[1].(json.Number).Float64()
			memoryFreeFloat, _ := values[2].(json.Number).Float64()
			memoryUsageFloat, _ := values[3].(json.Number).Float64()
			memoryUsed = append(memoryUsed, int64(memoryUsedFloat))
			memoryFree = append(memoryFree, int64(memoryFreeFloat))
			memoryUsage = append(memoryUsage, float32(memoryUsageFloat))
			totalMemoryUsed += int64(values[1].(float64))
			totalmemoryFree += int64(values[2].(float64))
			totalmemoryUsage += float32(values[3].(float64))
		}
		result.AvgMemoryUsed = totalMemoryUsed / int64(len(memoryUsed))
		result.AvgMemoryFree = totalmemoryFree / int64(len(memoryFree))
		result.AvgMemoryUsage = totalmemoryUsage / float32(len(memoryUsage))
	}
	return result, nil
}
