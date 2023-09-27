package client

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"hanglok-tech.com/monitor-server/domain/model"
	"hanglok-tech.com/monitor-server/infrastructure/config"
	"hanglok-tech.com/monitor-server/infrastructure/logger"
)

type InfluxDBClient struct {
	Client   client.Client
	Database string
}

// type DataPoint struct {
// 	Tags      map[string]string
// 	Fields    map[string]interface{}
// 	Timestamp time.Time
// }

// 判断数据库是否存在
func dbExists(c client.Client, dbName string) (bool, error) {
	resp, err := c.Query(client.NewQuery("SHOW DATABASES", "", ""))
	if err != nil {
		return false, err
	}

	for _, result := range resp.Results {
		for _, row := range result.Series {
			for _, value := range row.Values {
				if value[0] == dbName {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

// 创建数据库
func createDatabase(c client.Client, dbName string) error {
	q := client.NewQuery("CREATE DATABASE "+dbName, "", "")
	_, err := c.Query(q)
	return err
}

func NewInfluxDBClient(conf *config.ProjectConfig) (*InfluxDBClient, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: conf.InfluxDB.URL,
		// Username: conf.InfluxDB.Username,
		// Password: conf.InfluxDB.Password,
	})
	if err != nil {
		return nil, err
	}
	// 初始化数据库
	exists, err := dbExists(c, conf.InfluxDB.Database)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = createDatabase(c, conf.InfluxDB.Database)
		if err != nil {
			return nil, err
		}
	}
	logger.Logger.Info("InfluxDB client setup")
	return &InfluxDBClient{
		Client:   c,
		Database: conf.InfluxDB.Database,
	}, nil
}

// 查询数据
func (obj *InfluxDBClient) QueryData(queryString string) (*client.Response, error) {
	q := client.NewQuery(queryString, obj.Database, "s")
	resp, err := obj.Client.Query(q)
	return resp, err
}

// 写入数据
func (obj *InfluxDBClient) WriteData(measurement string, dataPoints []*model.StorgeDataPoint) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  obj.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}
	for _, dp := range dataPoints {
		pt, err := client.NewPoint(measurement, dp.Tags, dp.Fields, dp.Timestamp)
		if err != nil {
			return err
		}
		bp.AddPoint(pt)
	}
	return obj.Client.Write(bp)
}

// 删除数据
func (obj *InfluxDBClient) DeleteData(measurement string, tags map[string]string, startTime, endTime time.Time) error {
	q := client.NewQueryWithParameters(
		"DELETE FROM "+measurement+" WHERE time >= $start AND time <= $end",
		obj.Database,
		"s",
		map[string]interface{}{
			"start": startTime.Unix(),
			"end":   endTime.Unix(),
		},
	)
	_, err := obj.Client.Query(q)
	return err
}
