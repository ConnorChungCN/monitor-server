// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"hanglok-tech.com/monitor-server/adapter/grpc/service"
	"hanglok-tech.com/monitor-server/app/executor"
	monitorScheduler "hanglok-tech.com/monitor-server/app/scheduler"
	"hanglok-tech.com/monitor-server/domain/gateway"
	"hanglok-tech.com/monitor-server/infrastructure/client"
	"hanglok-tech.com/monitor-server/infrastructure/config"
	gatewayimpl "hanglok-tech.com/monitor-server/infrastructure/gateway-impl"
)

func initMonitorService(config *config.ProjectConfig) (*service.MonitorServer, error) {
	wire.Build(
		service.NewMonitorServer,
		executor.NewExecutor,
		wire.NewSet(gatewayimpl.NewMonitorGateway, wire.Bind(new(gateway.MonitorManager), new(*gatewayimpl.MonitorGateway))),
		client.NewSchedulerClient,
		client.NewInfluxDBClient,
	)
	return &service.MonitorServer{}, nil
}

func initMonitor(config *config.ProjectConfig) (*monitorScheduler.Monitor, error) {
	wire.Build(
		monitorScheduler.NewMonitor,
		wire.NewSet(gatewayimpl.NewMonitorGateway, wire.Bind(new(gateway.MonitorManager), new(*gatewayimpl.MonitorGateway))),
		client.NewSchedulerClient,
		client.NewInfluxDBClient,
	)
	return &monitorScheduler.Monitor{}, nil
}
