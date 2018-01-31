// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/sum.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	pb/sum.proto
	pb/echo.proto

It has these top-level messages:
	SumRequest
	SumResponse
	WaitRequest
	WaitResponse
	StrMessage
	StringMessage
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SumRequest is a request for Summator service.
type SumRequest struct {
	// A is the number we're adding to. Can't be zero for the sake of example.
	A int64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	// B is the number we're adding.
	B int64 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *SumRequest) Reset()                    { *m = SumRequest{} }
func (m *SumRequest) String() string            { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()               {}
func (*SumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SumRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumResponse struct {
	Sum   int64  `protobuf:"varint,1,opt,name=sum" json:"sum,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *SumResponse) Reset()                    { *m = SumResponse{} }
func (m *SumResponse) String() string            { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()               {}
func (*SumResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SumResponse) GetSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *SumResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// WaitRequest is only for test
type WaitRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *WaitRequest) Reset()                    { *m = WaitRequest{} }
func (m *WaitRequest) String() string            { return proto.CompactTextString(m) }
func (*WaitRequest) ProtoMessage()               {}
func (*WaitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WaitRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type WaitResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *WaitResponse) Reset()                    { *m = WaitResponse{} }
func (m *WaitResponse) String() string            { return proto.CompactTextString(m) }
func (*WaitResponse) ProtoMessage()               {}
func (*WaitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WaitResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type StrMessage struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StrMessage) Reset()                    { *m = StrMessage{} }
func (m *StrMessage) String() string            { return proto.CompactTextString(m) }
func (*StrMessage) ProtoMessage()               {}
func (*StrMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StrMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "example.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "example.SumResponse")
	proto.RegisterType((*WaitRequest)(nil), "example.WaitRequest")
	proto.RegisterType((*WaitResponse)(nil), "example.WaitResponse")
	proto.RegisterType((*StrMessage)(nil), "example.StrMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Summator service

type SummatorClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*WaitResponse, error)
	Hello(ctx context.Context, in *StrMessage, opts ...grpc.CallOption) (*StrMessage, error)
}

type summatorClient struct {
	cc *grpc.ClientConn
}

func NewSummatorClient(cc *grpc.ClientConn) SummatorClient {
	return &summatorClient{cc}
}

func (c *summatorClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := grpc.Invoke(ctx, "/example.Summator/Sum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *summatorClient) Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*WaitResponse, error) {
	out := new(WaitResponse)
	err := grpc.Invoke(ctx, "/example.Summator/Wait", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *summatorClient) Hello(ctx context.Context, in *StrMessage, opts ...grpc.CallOption) (*StrMessage, error) {
	out := new(StrMessage)
	err := grpc.Invoke(ctx, "/example.Summator/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Summator service

type SummatorServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Wait(context.Context, *WaitRequest) (*WaitResponse, error)
	Hello(context.Context, *StrMessage) (*StrMessage, error)
}

func RegisterSummatorServer(s *grpc.Server, srv SummatorServer) {
	s.RegisterService(&_Summator_serviceDesc, srv)
}

func _Summator_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummatorServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Summator/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummatorServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Summator_Wait_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummatorServer).Wait(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Summator/Wait",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummatorServer).Wait(ctx, req.(*WaitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Summator_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummatorServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Summator/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummatorServer).Hello(ctx, req.(*StrMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Summator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.Summator",
	HandlerType: (*SummatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Summator_Sum_Handler,
		},
		{
			MethodName: "Wait",
			Handler:    _Summator_Wait_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _Summator_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/sum.proto",
}

func init() { proto.RegisterFile("pb/sum.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0xe5, 0xfe, 0x00, 0x9d, 0x56, 0xa2, 0x98, 0x16, 0x55, 0xa1, 0x0b, 0xb0, 0x58, 0x54,
	0x5d, 0x34, 0x02, 0xc4, 0xa6, 0x27, 0x60, 0x01, 0x02, 0x25, 0x0b, 0xd6, 0x13, 0xc9, 0x2a, 0x91,
	0xe2, 0x38, 0xf8, 0xa7, 0xb0, 0xe6, 0x0a, 0x1c, 0x8d, 0x2b, 0x70, 0x0b, 0x36, 0xc8, 0x76, 0xda,
	0x06, 0xc4, 0xce, 0x6f, 0xf4, 0xe6, 0x9b, 0x79, 0x63, 0x18, 0x54, 0x59, 0xac, 0xad, 0x58, 0x54,
	0x4a, 0x1a, 0x49, 0xf7, 0xf9, 0x1b, 0x8a, 0xaa, 0xe0, 0xd1, 0x74, 0x25, 0xe5, 0xaa, 0xe0, 0x31,
	0x56, 0x79, 0x8c, 0x65, 0x29, 0x0d, 0x9a, 0x5c, 0x96, 0x3a, 0xd8, 0xd8, 0x0c, 0x20, 0xb5, 0x22,
	0xe1, 0x2f, 0x96, 0x6b, 0x43, 0x07, 0x40, 0x70, 0x42, 0xce, 0xc8, 0xac, 0x9d, 0x10, 0x74, 0x2a,
	0x9b, 0xb4, 0x82, 0xca, 0xd8, 0x0d, 0xf4, 0xbd, 0x53, 0x57, 0xb2, 0xd4, 0x9c, 0x0e, 0xa1, 0xad,
	0xad, 0xa8, 0xcd, 0xee, 0x49, 0x47, 0xd0, 0xe5, 0x4a, 0x49, 0xe5, 0x5b, 0x7a, 0x49, 0x10, 0xec,
	0x1c, 0xfa, 0x4f, 0x98, 0x9b, 0xcd, 0x04, 0x0a, 0x9d, 0x12, 0x05, 0xf7, 0x7d, 0xbd, 0xc4, 0xbf,
	0xd9, 0x05, 0x0c, 0x82, 0xa5, 0x46, 0x6f, 0x41, 0xa4, 0x09, 0x62, 0x00, 0xa9, 0x51, 0xf7, 0x5c,
	0x6b, 0x5c, 0x79, 0xcf, 0x1a, 0x0b, 0xbb, 0x01, 0x05, 0x71, 0xf5, 0x4d, 0xe0, 0x20, 0xb5, 0x42,
	0xa0, 0x91, 0x8a, 0xde, 0x41, 0x3b, 0xb5, 0x82, 0x1e, 0x2f, 0xea, 0x4b, 0x2c, 0x76, 0x41, 0xa3,
	0xd1, 0xef, 0x62, 0x18, 0xcc, 0xa2, 0xf7, 0xcf, 0xaf, 0x8f, 0xd6, 0x88, 0x1d, 0xc6, 0xeb, 0xcb,
	0xb8, 0x36, 0xb8, 0x93, 0x2e, 0xc9, 0x9c, 0x3e, 0x42, 0xc7, 0x2d, 0x49, 0x77, 0x9d, 0x8d, 0x58,
	0xd1, 0xf8, 0x4f, 0xb5, 0x06, 0x9e, 0x7a, 0xe0, 0x98, 0x0d, 0x9b, 0xc0, 0x57, 0xcc, 0x8d, 0x23,
	0x3e, 0x40, 0xf7, 0x96, 0x17, 0x85, 0x6c, 0x6e, 0xb8, 0x0d, 0x18, 0xfd, 0x57, 0x64, 0x53, 0xcf,
	0x3b, 0x61, 0x47, 0x4d, 0xde, 0xb3, 0x83, 0x2c, 0xc9, 0x3c, 0xdb, 0xf3, 0x5f, 0x7a, 0xfd, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x6a, 0x61, 0x56, 0xfe, 0x09, 0x02, 0x00, 0x00,
}
