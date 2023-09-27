package config

import (
	"flag"
	"log"
	"os"

	"github.com/spf13/viper"
	"hanglok-tech.com/monitor-server/utils"
)

// 配置对象
var Config ProjectConfig

func GetConfig() *ProjectConfig {
	return &Config
}

type InfluxDBConfig struct {
	// Host     string `mapstructure:"host"`
	// Port     int    `mapstructure:"port"`
	URL string `mapstructure:"url"`
	// Username string `mapstructure:"username"`
	// Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type SchedulerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ProjectConfig struct {
	Host      string          `mapstructure:"host"`
	Port      int             `mapstructure:"port"`
	Scheduler SchedulerConfig `mapstructure:"scheduler"`
	InfluxDB  InfluxDBConfig  `mapstructure:"influxdb"`
}

// 配置选项
type configOption struct {
	path *string
}

func (o *configOption) Name() string {
	return "config_option"
}

func (o *configOption) Handle() {
	p := *o.path
	if _, err := os.Stat(p); err != nil && !os.IsExist(err) {
		log.Panicf("cannot open config[%s], %s", p, err)
	}
	viper.SetConfigFile(p)
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("readin config failed, %s", err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		log.Panicf("unmarshal config failed, %s", err)
	}
}

func init() {
	utils.AddHandler(&configOption{
		path: flag.String("config", "./config/config.yml", "app config file path"),
	})
}
