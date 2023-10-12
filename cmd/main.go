package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"hanglok-tech.com/monitor-server/adapter/grpc/proto/monitor"
	"hanglok-tech.com/monitor-server/infrastructure/config"
	"hanglok-tech.com/monitor-server/utils"
)

const (
	interval = 200 * time.Millisecond
)

func main() {
	// run handler
	flag.Parse()
	utils.RunHandlers()
	// set log
	logger := grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	grpclog.SetLoggerV2(logger)

	// config
	config := config.GetConfig()
	fmt.Printf("get config: %+v", config)
	// grpc
	url := fmt.Sprintf("%s:%d", config.Host, config.Port)
	svc, err := initMonitorService(config)
	if err != nil {
		fmt.Printf("initMonitorService error: %s", err)
		panic(err)
	}
	lis, err := net.Listen("tcp", url)
	if err != nil {
		fmt.Printf("listen port error: %s", err)
		panic(err)
	}
	// new 服务
	s := grpc.NewServer()
	monitor.RegisterMonitorServer(s, svc)

	// rpc
	go func() {
		fmt.Printf("Grpc started")
		if err := s.Serve(lis); err != nil {
			fmt.Printf("Serve error: %s", err)
			panic(err)
		}
	}()

	//monitorScheduler
	monitor, err := initMonitor(config)
	if err != nil {
		fmt.Printf("start consumer error: %s", err)
		os.Exit(1)
	}
	//开始监控
	ctx, cancel := context.WithCancel(context.TODO())
	go monitor.StartMonitoring(ctx, interval)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	s.Stop()
	cancel()
	log.Println("All Server shutdown")
}
