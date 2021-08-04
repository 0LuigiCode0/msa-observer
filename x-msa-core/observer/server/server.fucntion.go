package server

import (
	"context"
	"fmt"
	"net"
	"x-msa-core/grpc/msa_observer"
	"x-msa-core/monitor/dependents"

	"google.golang.org/grpc"
)

func NewObserverServer(key, addr string) ObserverServer {
	server := &observerServer{
		server:   grpc.NewServer(),
		key:      key,
		addr:     addr,
		monitors: dependents.NewMDM(),
	}
	msa_observer.RegisterObserverServer(server.server, server)

	return server
}

func (o *observerServer) Start() error {
	gListen, err := net.Listen("tcp", o.addr)
	if err != nil {
		return fmt.Errorf("glistener error: %v", err)
	}
	if err := o.server.Serve(gListen); err != nil {
		if err == grpc.ErrServerStopped {
			return nil
		}
		return fmt.Errorf("gserve error: %v", err)
	}

	return nil
}

func (o *observerServer) Monitors() dependents.MonitorDependentsManager { return o.monitors }

func (o *observerServer) Close() { o.server.Stop() }

func (o *observerServer) SetPushFirst(f func(ctx context.Context, req *msa_observer.RequestPushFirst) (*msa_observer.ResponsePushFirst, error)) {
	o.pushFirst = f
}
func (o *observerServer) SetPushStats(f func(ctx context.Context, req *msa_observer.RequestPushStats) (*msa_observer.Response, error)) {
	o.pushStats = f
}
func (o *observerServer) SetPushStatus(f func(ctx context.Context, req *msa_observer.RequestPushStatus) (*msa_observer.Response, error)) {
	o.pushStatus = f
}
func (o *observerServer) SetWho(f func(ctx context.Context, req *msa_observer.RequestWho) (*msa_observer.ResponseWho, error)) {
	o.who = f
}
func (o *observerServer) SetRestartService(f func(ctx context.Context, req *msa_observer.RequestRestartService) (*msa_observer.Response, error)) {
	o.restartService = f
}

func (o *observerServer) PushFirst(ctx context.Context, req *msa_observer.RequestPushFirst) (*msa_observer.ResponsePushFirst, error) {
	return o.pushFirst(ctx, req)
}
func (o *observerServer) PushStats(ctx context.Context, req *msa_observer.RequestPushStats) (*msa_observer.Response, error) {
	return o.pushStats(ctx, req)
}
func (o *observerServer) PushStatus(ctx context.Context, req *msa_observer.RequestPushStatus) (*msa_observer.Response, error) {
	return o.pushStatus(ctx, req)
}
func (o *observerServer) Who(ctx context.Context, req *msa_observer.RequestWho) (*msa_observer.ResponseWho, error) {
	return o.who(ctx, req)
}
func (o *observerServer) RestartService(ctx context.Context, req *msa_observer.RequestRestartService) (*msa_observer.Response, error) {
	return o.restartService(ctx, req)
}
