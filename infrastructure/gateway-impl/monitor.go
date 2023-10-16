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
	cpuDataPoints := make([]*client.DataPoint, 0)
	memoryDataPoints := make([]*client.DataPoint, 0)
	gpuDataPoints := make([]*client.DataPoint, 0)
	for _, v := range workers {
		timeStamp := time.Now()
		cpuDataPoints = append(cpuDataPoints, &client.DataPoint{
			Tags: map[string]string{
				"AlgorithmName":    v.AlgorithmName,
				"AlgorithmVersion": v.AlgorithmVersion,
				"TaskId":           v.TaskId,
			},
			Fields: map[string]interface{}{
				"CPUPercent": v.CpuStats.CPUPercent,
			},
			Timestamp: timeStamp,
		})
		memoryDataPoints = append(memoryDataPoints, &client.DataPoint{
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
			Timestamp: timeStamp,
		})
		for _, gpuInfo := range v.GpuStats.GPUsInfo {
			gpuDataPoints = append(gpuDataPoints, &client.DataPoint{
				Tags: map[string]string{
					"AlgorithmName":    v.AlgorithmName,
					"AlgorithmVersion": v.AlgorithmVersion,
					"TaskId":           v.TaskId,
					"CudaVersion":      v.GpuStats.CudaVersion,
				},
				Fields: map[string]interface{}{
					"Id":          gpuInfo.Id,
					"ProductName": gpuInfo.ProductName,
					"GpuUsage":    gpuInfo.GpuUsage,
					"MemoryUsage": gpuInfo.MemoryUsage,
					"MemoryUsed":  gpuInfo.MemoryUsed,
					"MemoryFree":  gpuInfo.MemoryFree,
				},
				Timestamp: timeStamp,
			})
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
			timeInt, _ := values[0].(json.Number).Int64()
			timestamp := time.Unix(timeInt, 0)
			cpuPercentFloat, _ := values[1].(json.Number).Float64()
			cpuInquire = append(cpuInquire, &model.QueryCpuInfo{
				Time:       timestamp.String(),
				CpuPercent: cpuPercentFloat,
			})
		}
		result.CpuResult = cpuInquire
	}
	// 查询 containerMemoryState 中的数据
	memoryQueryString := fmt.Sprintf(`SELECT "MemoryUsage", "MemoryUsed", "MemoryFree" FROM containerMemState WHERE "TaskId"='%s'`, taskId)
	memoryRsp, err := obj.InfluxDBClient.QueryData(memoryQueryString)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	//memoryRsp.Results[0].Series[0].Columns字段名数组
	//memoryRsp.Results[0].Series[0].Values是[时间, memoryUsed的值, memoryMaxUsed的值]
	if len(memoryRsp.Results) > 0 && len(memoryRsp.Results[0].Series) > 0 {
		var memInquire []*model.QueryMemInfo
		for _, values := range memoryRsp.Results[0].Series[0].Values {
			timeInt, _ := values[0].(json.Number).Int64()
			timestamp := time.Unix(timeInt, 0)
			usageFloat, _ := values[1].(json.Number).Float64()
			usedInt, _ := values[2].(json.Number).Int64()
			freeInt, _ := values[3].(json.Number).Int64()
			memInquire = append(memInquire, &model.QueryMemInfo{
				Time:  timestamp.String(),
				Usage: usageFloat,
				Used:  usedInt,
				Free:  freeInt,
			})
		}
		result.MemResult = memInquire
	}
	// 查询 containerGPUState 中的数据
	gpuQueryString := fmt.Sprintf(`SELECT "Id", "ProductName", "GpuUsage", "MemoryUsage", "MemoryUsed", "MemoryFree" FROM containerGPUState WHERE "TaskId"='%s'`, taskId)
	gpuRsp, err := obj.InfluxDBClient.QueryData(gpuQueryString)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	if len(gpuRsp.Results) > 0 && len(gpuRsp.Results[0].Series) > 0 {
		var GpuInquire []*model.QueryGpuInfo
		for _, values := range gpuRsp.Results[0].Series[0].Values {
			timeInt, _ := values[0].(json.Number).Int64()
			timestamp := time.Unix(timeInt, 0)
			idString := values[1].(string)
			ProductNameString := values[2].(string)
			GpuUsageFloat, _ := values[3].(json.Number).Float64()
			MemoryUsageFloat, _ := values[4].(json.Number).Float64()
			MemoryUsedInt, _ := values[5].(json.Number).Int64()
			MemoryFreeInt, _ := values[6].(json.Number).Int64()
			GpuInquire = append(GpuInquire, &model.QueryGpuInfo{
				Time:        timestamp.String(),
				Id:          idString,
				ProductName: ProductNameString,
				GpuUsage:    GpuUsageFloat,
				MemoryUsage: MemoryUsageFloat,
				MemoryUsed:  MemoryUsedInt,
				MemoryFree:  MemoryFreeInt,
			})
		}
		result.GpuResult = GpuInquire
	}
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
	// 查询 containerMemState 中的数据
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
			memoryUsedFloat, _ := values[1].(json.Number).Int64()
			memoryFreeFloat, _ := values[2].(json.Number).Int64()
			memoryUsageFloat, _ := values[3].(json.Number).Float64()
			memoryUsed = append(memoryUsed, int64(memoryUsedFloat))
			memoryFree = append(memoryFree, int64(memoryFreeFloat))
			memoryUsage = append(memoryUsage, float32(memoryUsageFloat))
			totalMemoryUsed += memoryUsedFloat
			totalmemoryFree += memoryFreeFloat
			totalmemoryUsage += float32(memoryUsageFloat)
		}
		result.AvgMemoryUsed = totalMemoryUsed / int64(len(memoryUsed))
		result.AvgMemoryFree = totalmemoryFree / int64(len(memoryFree))
		result.AvgMemoryUsage = totalmemoryUsage / float32(len(memoryUsage))
	}
	// 查询 containerGPUState 中的数据
	gpuQueryString := fmt.Sprintf(`SELECT "GpuUsage", "MemoryUsage", "MemoryUsed", "MemoryFree" FROM containerGPUState WHERE "TaskId"='%s'`, taskId)
	gpuRsp, err := obj.InfluxDBClient.QueryData(gpuQueryString)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	if len(gpuRsp.Results) > 0 && len(gpuRsp.Results[0].Series) > 0 {
		var gpuUsage []float32
		var gpuMemoryUsed []int64
		var gpuMemoryFree []int64
		var gpuemoryUsage []float32
		var totalGpuUsage float32 = 0
		var totalGpuMemoryUsed int64 = 0
		var totalGpumemoryFree int64 = 0
		var totalGpumemoryUsage float32 = 0
		for _, values := range gpuRsp.Results[0].Series[0].Values {
			GpuUsageFloat, _ := values[1].(json.Number).Float64()
			GpuMemoryUsageFloat, _ := values[2].(json.Number).Float64()
			GpuMemoryUsedInt, _ := values[3].(json.Number).Int64()
			GpuMemoryFreeInt, _ := values[4].(json.Number).Int64()

			gpuUsage = append(gpuUsage, float32(GpuUsageFloat))
			gpuemoryUsage = append(gpuemoryUsage, float32(GpuMemoryUsageFloat))
			gpuMemoryUsed = append(gpuMemoryUsed, int64(GpuMemoryUsedInt))
			gpuMemoryFree = append(gpuMemoryFree, int64(GpuMemoryFreeInt))

			totalGpuUsage += float32(GpuUsageFloat)
			totalGpumemoryUsage += float32(GpuMemoryUsageFloat)
			totalGpuMemoryUsed += GpuMemoryUsedInt
			totalGpumemoryFree += GpuMemoryFreeInt
		}
		result.AvgGpuUsage = totalGpuUsage / float32(len(gpuUsage))
		result.AvgGpuMemoryUsage = totalGpumemoryUsage / float32(len(gpuemoryUsage))
		result.AvgGpuMemoryUsed = totalGpuMemoryUsed / int64(len(gpuMemoryUsed))
		result.AvgGpuMemoryFree = totalGpumemoryFree / int64(len(gpuMemoryFree))
	}
	return result, nil
}
