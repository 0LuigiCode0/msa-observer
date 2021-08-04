package client

import (
	"context"
	"x-msa-core/grpc/msa_monitor"

	"google.golang.org/grpc"
)

type MonitorClient interface {
	Close()

	AddService(req *msa_monitor.RequestAddService) (*msa_monitor.Response, error)
	DeleteService(req *msa_monitor.RequestDelService) (*msa_monitor.Response, error)
	RebuildService(req *msa_monitor.RequestRebuildService) (*msa_monitor.Response, error)
	SetService(req *msa_monitor.RequestSetService) (*msa_monitor.Response, error)
}

type monitorClient struct {
	client msa_monitor.MonitorClient
	conn   *grpc.ClientConn
	ctx    context.Context
	close  context.CancelFunc
}
