package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"hanglok-tech.com/monitor-server/domain/model"
	"hanglok-tech.com/monitor-server/infrastructure/client"
	"hanglok-tech.com/monitor-server/infrastructure/config"
	"hanglok-tech.com/monitor-server/utils"
)

func main() {
	// run handler
	flag.Parse()
	utils.RunHandlers()
	config := config.GetConfig()
	cli, err := client.NewInfluxDBClient(config)
	if err != nil {
		log.Fatal(err)
	}
	var bp []*model.StorgeDataPoint
	var j int = 95
	v := 15
	for i := 0; i < 5; i++ {
		bp = append(bp, &model.StorgeDataPoint{
			Tags: map[string]string{
				"AlgorithmName":    "bone",
				"AlgorithmVersion": "0.0.1",
				"TaskId":           "dgasdgasd",
			},
			Fields: map[string]interface{}{
				"CPUPersent": j,
			},
			Timestamp: time.Now(),
		})
		j += 1
		v += 1
		fmt.Printf("i: %d\n", i)
		time.Sleep(time.Second)
	}
	cli.WriteData("test", bp)
	cpuQueryString := fmt.Sprintf(`SELECT "CPUPersent" FROM test WHERE "TaskId" = 'dgasdgasd'`)
	cpuRsp, err := cli.QueryData(cpuQueryString)
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range cpuRsp.Results[0].Series[0].Values {
		for j, value := range row {
			log.Printf("j:%d value:%v\n", j, value)
		}
	}
}
