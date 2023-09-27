package client

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hanglok-tech.com/monitor-server/adapter/grpc/proto/scheduler"
	"hanglok-tech.com/monitor-server/infrastructure/config"
)

type SchedulerClient struct {
	Client scheduler.SchedulerClient
}

func NewSchedulerClient(config *config.ProjectConfig) (*SchedulerClient, error) {
	url := fmt.Sprintf("%s:%d", config.Scheduler.Host, config.Scheduler.Port)
	log.Printf("connect to rpc: %s", url)
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connect scheduler failed, %w", err)
	}
	cli := scheduler.NewSchedulerClient(conn)
	return &SchedulerClient{
		Client: cli,
	}, nil
}
