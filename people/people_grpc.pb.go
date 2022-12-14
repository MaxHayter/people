// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package people

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PeopleServiceClient is the client API for PeopleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeopleServiceClient interface {
	Registrate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Login(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type peopleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPeopleServiceClient(cc grpc.ClientConnInterface) PeopleServiceClient {
	return &peopleServiceClient{cc}
}

func (c *peopleServiceClient) Registrate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/people.PeopleService/Registrate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peopleServiceClient) Login(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/people.PeopleService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeopleServiceServer is the server API for PeopleService service.
// All implementations must embed UnimplementedPeopleServiceServer
// for forward compatibility
type PeopleServiceServer interface {
	Registrate(context.Context, *Request) (*emptypb.Empty, error)
	Login(context.Context, *Request) (*Result, error)
	mustEmbedUnimplementedPeopleServiceServer()
}

// UnimplementedPeopleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPeopleServiceServer struct {
}

func (UnimplementedPeopleServiceServer) Registrate(context.Context, *Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registrate not implemented")
}
func (UnimplementedPeopleServiceServer) Login(context.Context, *Request) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedPeopleServiceServer) mustEmbedUnimplementedPeopleServiceServer() {}

// UnsafePeopleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeopleServiceServer will
// result in compilation errors.
type UnsafePeopleServiceServer interface {
	mustEmbedUnimplementedPeopleServiceServer()
}

func RegisterPeopleServiceServer(s grpc.ServiceRegistrar, srv PeopleServiceServer) {
	s.RegisterService(&PeopleService_ServiceDesc, srv)
}

func _PeopleService_Registrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeopleServiceServer).Registrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/people.PeopleService/Registrate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeopleServiceServer).Registrate(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeopleService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeopleServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/people.PeopleService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeopleServiceServer).Login(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// PeopleService_ServiceDesc is the grpc.ServiceDesc for PeopleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeopleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "people.PeopleService",
	HandlerType: (*PeopleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registrate",
			Handler:    _PeopleService_Registrate_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _PeopleService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "people.proto",
}
