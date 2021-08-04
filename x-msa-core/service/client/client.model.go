package client

import (
	"context"
	"x-msa-core/grpc/msa_service"
	"x-msa-core/helper"

	"google.golang.org/grpc"
)

type ServiceClient interface {
	Close()

	Call(req *msa_service.RequestCall) (*msa_service.ResponseCall, error)
	AddService(req *msa_service.RequestAddService) (*msa_service.Response, error)
	DeleteService(req *msa_service.RequestDelService) (*msa_service.Response, error)
}

type serviceClient struct {
	client msa_service.ServiceClient
	conn   *grpc.ClientConn
	group  helper.GroupsType
	ctx    context.Context
	close  context.CancelFunc
}
