// Code generated by protoc-gen-go.
// source: hello.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	hello.proto
	profile.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
	EmptyMessage
	IDMessage
	ProfileMessage
*/
package helloworld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type HelloRequest struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HelloResponse struct {
	Reply string `protobuf:"bytes,1,opt,name=reply" json:"reply,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "helloworld.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for HelloService service

type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/helloworld.HelloService/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloService service

type HelloServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(HelloServiceServer).SayHello(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x73, 0xca, 0xf3, 0x8b, 0x72, 0x52,
	0x94, 0xb4, 0xb8, 0x78, 0x3c, 0x40, 0xbc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x29,
	0x2e, 0x8e, 0xf4, 0xa2, 0xd4, 0xd4, 0x92, 0xcc, 0xbc, 0x74, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x38, 0x5f, 0x49, 0x95, 0x8b, 0x17, 0xaa, 0xb6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48,
	0x84, 0x8b, 0xb5, 0x28, 0xb5, 0x20, 0xa7, 0x12, 0xaa, 0x12, 0xc2, 0x31, 0xf2, 0x87, 0x1a, 0x19,
	0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xcf, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0x16, 0x12,
	0x92, 0xd0, 0x43, 0xd8, 0xad, 0x87, 0x6c, 0xb1, 0x94, 0x24, 0x16, 0x19, 0x88, 0x35, 0x49, 0x6c,
	0x60, 0x67, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x40, 0xf2, 0xf1, 0x9a, 0xc5, 0x00, 0x00,
	0x00,
}
