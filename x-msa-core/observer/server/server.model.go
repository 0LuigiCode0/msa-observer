package server

import (
	"context"
	"x-msa-core/grpc/msa_observer"
	"x-msa-core/monitor/dependents"

	"google.golang.org/grpc"
)

type ObserverServer interface {
	Start() error

	Monitors() dependents.MonitorDependentsManager

	Close()

	SetPushFirst(f func(ctx context.Context, req *msa_observer.RequestPushFirst) (*msa_observer.ResponsePushFirst, error))
	SetPushStats(f func(ctx context.Context, req *msa_observer.RequestPushStats) (*msa_observer.Response, error))
	SetPushStatus(f func(ctx context.Context, req *msa_observer.RequestPushStatus) (*msa_observer.Response, error))
	SetWho(f func(ctx context.Context, req *msa_observer.RequestWho) (*msa_observer.ResponseWho, error))
	SetRestartService(f func(ctx context.Context, req *msa_observer.RequestRestartService) (*msa_observer.Response, error))
}

type observerServer struct {
	msa_observer.ObserverServer

	server *grpc.Server
	key    string
	addr   string

	pushFirst      func(context.Context, *msa_observer.RequestPushFirst) (*msa_observer.ResponsePushFirst, error)
	pushStats      func(context.Context, *msa_observer.RequestPushStats) (*msa_observer.Response, error)
	pushStatus     func(context.Context, *msa_observer.RequestPushStatus) (*msa_observer.Response, error)
	who            func(context.Context, *msa_observer.RequestWho) (*msa_observer.ResponseWho, error)
	restartService func(context.Context, *msa_observer.RequestRestartService) (*msa_observer.Response, error)

	monitors dependents.MonitorDependentsManager
}
