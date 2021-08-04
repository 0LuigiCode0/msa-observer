// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package msa_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Call(ctx context.Context, in *RequestCall, opts ...grpc.CallOption) (*ResponseCall, error)
	AddService(ctx context.Context, in *RequestAddService, opts ...grpc.CallOption) (*Response, error)
	DeleteService(ctx context.Context, in *RequestDelService, opts ...grpc.CallOption) (*Response, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Call(ctx context.Context, in *RequestCall, opts ...grpc.CallOption) (*ResponseCall, error) {
	out := new(ResponseCall)
	err := c.cc.Invoke(ctx, "/msa_service.Service/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AddService(ctx context.Context, in *RequestAddService, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/msa_service.Service/AddService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteService(ctx context.Context, in *RequestDelService, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/msa_service.Service/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	Call(context.Context, *RequestCall) (*ResponseCall, error)
	AddService(context.Context, *RequestAddService) (*Response, error)
	DeleteService(context.Context, *RequestDelService) (*Response, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Call(context.Context, *RequestCall) (*ResponseCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedServiceServer) AddService(context.Context, *RequestAddService) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddService not implemented")
}
func (UnimplementedServiceServer) DeleteService(context.Context, *RequestDelService) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCall)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msa_service.Service/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Call(ctx, req.(*RequestCall))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AddService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msa_service.Service/AddService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddService(ctx, req.(*RequestAddService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDelService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msa_service.Service/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteService(ctx, req.(*RequestDelService))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msa_service.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Service_Call_Handler,
		},
		{
			MethodName: "AddService",
			Handler:    _Service_AddService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _Service_DeleteService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msa/proto/msa_servoce.proto",
}
