package client

import (
	"context"
	"x-msa-core/grpc/msa_observer"

	"google.golang.org/grpc"
)

type ObserverClient interface {
	Close()

	PushFirst(req *msa_observer.RequestPushFirst) (*msa_observer.ResponsePushFirst, error)
	PushStats(req *msa_observer.RequestPushStats) (*msa_observer.Response, error)
	PushStatus(req *msa_observer.RequestPushStatus) (*msa_observer.Response, error)
	Who(req *msa_observer.RequestWho) (*msa_observer.ResponseWho, error)
	RestartService(req *msa_observer.RequestRestartService) (*msa_observer.Response, error)
}

type observerClient struct {
	client msa_observer.ObserverClient
	conn   *grpc.ClientConn
	ctx    context.Context
	close  context.CancelFunc
}
