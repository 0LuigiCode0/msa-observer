package server

import (
	"context"
	"fmt"
	"net"
	"x-msa-core/grpc/msa_service"
	dep_observer "x-msa-core/observer/dependents"
	dep_service "x-msa-core/service/dependents"

	"google.golang.org/grpc"
)

func NewServiceServer(key, addr string) ServiceServer {
	server := &serviceServer{
		server:    grpc.NewServer(),
		key:       key,
		addr:      addr,
		observers: dep_observer.NewODM(),
		services:  dep_service.NewSDM(),
	}
	msa_service.RegisterServiceServer(server.server, server)

	return server
}

func (s *serviceServer) Start() error {
	gListen, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("glistener error: %v", err)
	}
	if err := s.server.Serve(gListen); err != nil {
		if err == grpc.ErrServerStopped {
			return nil
		}
		return fmt.Errorf("gserve error: %v", err)
	}

	return nil
}

func (s *serviceServer) Services() dep_service.ServiceDependentsManager    { return s.services }
func (s *serviceServer) Observers() dep_observer.ObserverDependentsManager { return s.observers }

func (s *serviceServer) Close() { s.server.Stop() }

func (s *serviceServer) SetCall(f func(ctx context.Context, req *msa_service.RequestCall) (*msa_service.ResponseCall, error)) {
	s.call = f
}

func (s *serviceServer) Call(ctx context.Context, req *msa_service.RequestCall) (*msa_service.ResponseCall, error) {
	return s.call(ctx, req)
}

func (s *serviceServer) AddService(ctx context.Context, req *msa_service.RequestAddService) (*msa_service.Response, error) {
	if err := s.services.Add(req.Key, fmt.Sprintf("%v:%v", req.Host, req.Port), req.GroupType); err != nil {
		return nil, fmt.Errorf("cannot add service, key %v,  addr %v: %v", req.Key, fmt.Sprintf("%v:%v", req.Host, req.Port), err)
	}
	return &msa_service.Response{
		Success: true,
	}, nil
}

func (s *serviceServer) DeleteService(ctx context.Context, req *msa_service.RequestDelService) (*msa_service.Response, error) {
	if err := s.services.Delete(req.Key); err != nil {
		return nil, fmt.Errorf("cannot remove service, key %v: %v", req.Key, err)
	}
	return &msa_service.Response{
		Success: true,
	}, nil
}
