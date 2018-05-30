// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/element/api/services/health/v1/health.proto

/*
Package health is a generated protocol buffer package.

It is generated from these files:
	github.com/ehazlett/element/api/services/health/v1/health.proto

It has these top-level messages:
	HealthResponse
*/
package health

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type HealthResponse struct {
	OSName      string                     `protobuf:"bytes,1,opt,name=os_name,json=osName,proto3" json:"os_name,omitempty"`
	OSVersion   string                     `protobuf:"bytes,2,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	Uptime      *google_protobuf1.Duration `protobuf:"bytes,3,opt,name=uptime" json:"uptime,omitempty"`
	Cpus        int64                      `protobuf:"varint,4,opt,name=cpus,proto3" json:"cpus,omitempty"`
	MemoryTotal int64                      `protobuf:"varint,5,opt,name=memory_total,json=memoryTotal,proto3" json:"memory_total,omitempty"`
	MemoryFree  int64                      `protobuf:"varint,6,opt,name=memory_free,json=memoryFree,proto3" json:"memory_free,omitempty"`
	MemoryUsed  int64                      `protobuf:"varint,7,opt,name=memory_used,json=memoryUsed,proto3" json:"memory_used,omitempty"`
}

func (m *HealthResponse) Reset()                    { *m = HealthResponse{} }
func (m *HealthResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()               {}
func (*HealthResponse) Descriptor() ([]byte, []int) { return fileDescriptorHealth, []int{0} }

func (m *HealthResponse) GetOSName() string {
	if m != nil {
		return m.OSName
	}
	return ""
}

func (m *HealthResponse) GetOSVersion() string {
	if m != nil {
		return m.OSVersion
	}
	return ""
}

func (m *HealthResponse) GetUptime() *google_protobuf1.Duration {
	if m != nil {
		return m.Uptime
	}
	return nil
}

func (m *HealthResponse) GetCpus() int64 {
	if m != nil {
		return m.Cpus
	}
	return 0
}

func (m *HealthResponse) GetMemoryTotal() int64 {
	if m != nil {
		return m.MemoryTotal
	}
	return 0
}

func (m *HealthResponse) GetMemoryFree() int64 {
	if m != nil {
		return m.MemoryFree
	}
	return 0
}

func (m *HealthResponse) GetMemoryUsed() int64 {
	if m != nil {
		return m.MemoryUsed
	}
	return 0
}

func init() {
	proto.RegisterType((*HealthResponse)(nil), "element.services.health.v1.HealthResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Health service

type HealthClient interface {
	Health(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*HealthResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Health(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := grpc.Invoke(ctx, "/element.services.health.v1.Health/Health", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Health(context.Context, *google_protobuf.Empty) (*HealthResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/element.services.health.v1.Health/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Health(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "element.services.health.v1.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Health_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/element/api/services/health/v1/health.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/element/api/services/health/v1/health.proto", fileDescriptorHealth)
}

var fileDescriptorHealth = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6b, 0xdb, 0x30,
	0x14, 0xc6, 0x71, 0x92, 0x39, 0x44, 0xd9, 0x76, 0x10, 0x63, 0x78, 0x1e, 0x2c, 0xd9, 0x76, 0x09,
	0x63, 0x48, 0x24, 0x3b, 0x8d, 0x1c, 0x06, 0x61, 0x2d, 0x3d, 0x94, 0x06, 0x9c, 0x36, 0x87, 0x5e,
	0x8c, 0x93, 0xbc, 0xd8, 0x06, 0xcb, 0xcf, 0x58, 0xb2, 0x21, 0xfd, 0x27, 0xfa, 0x1f, 0xe6, 0x90,
	0xbf, 0xa4, 0x58, 0x52, 0x4a, 0xda, 0xd2, 0x43, 0x4f, 0xfe, 0xfc, 0x7d, 0x3f, 0xf1, 0x3e, 0x9e,
	0x44, 0xfe, 0xc5, 0xa9, 0x4a, 0xaa, 0x15, 0x5b, 0xa3, 0xe0, 0x90, 0x44, 0x77, 0x19, 0x28, 0xc5,
	0x21, 0x03, 0x01, 0xb9, 0xe2, 0x51, 0x91, 0x72, 0x09, 0x65, 0x9d, 0xae, 0x41, 0xf2, 0x04, 0xa2,
	0x4c, 0x25, 0xbc, 0x1e, 0x5b, 0xc5, 0x8a, 0x12, 0x15, 0x52, 0xdf, 0xc2, 0xec, 0x08, 0x32, 0x1b,
	0xd7, 0x63, 0xff, 0x6b, 0x8c, 0x18, 0x67, 0xc0, 0x35, 0xb9, 0xaa, 0xb6, 0x1c, 0x44, 0xa1, 0x76,
	0xe6, 0xa0, 0xff, 0xed, 0x79, 0xb8, 0xa9, 0xca, 0x48, 0xa5, 0x98, 0xdb, 0xfc, 0x53, 0x8c, 0x31,
	0x6a, 0xc9, 0x1b, 0x65, 0xdc, 0x1f, 0xf7, 0x2d, 0xf2, 0xf1, 0x42, 0x0f, 0x08, 0x40, 0x16, 0x98,
	0x4b, 0xa0, 0x3f, 0x49, 0x17, 0x65, 0x98, 0x47, 0x02, 0x3c, 0x67, 0xe8, 0x8c, 0x7a, 0x33, 0x72,
	0xd8, 0x0f, 0xdc, 0xf9, 0xe2, 0x2a, 0x12, 0x10, 0xb8, 0x28, 0x9b, 0x2f, 0xfd, 0x4d, 0x08, 0xca,
	0xb0, 0x86, 0x52, 0xa6, 0x98, 0x7b, 0x2d, 0xcd, 0x7d, 0x38, 0xec, 0x07, 0xbd, 0xf9, 0x62, 0x69,
	0xcc, 0xa0, 0x87, 0xd2, 0x4a, 0x3a, 0x26, 0x6e, 0x55, 0xa8, 0x54, 0x80, 0xd7, 0x1e, 0x3a, 0xa3,
	0xfe, 0xe4, 0x0b, 0x33, 0x65, 0xd9, 0xb1, 0x2c, 0xfb, 0x6f, 0xcb, 0x06, 0x16, 0xa4, 0x94, 0x74,
	0xd6, 0x45, 0x25, 0xbd, 0xce, 0xd0, 0x19, 0xb5, 0x03, 0xad, 0xe9, 0x77, 0xf2, 0x5e, 0x80, 0xc0,
	0x72, 0x17, 0x2a, 0x54, 0x51, 0xe6, 0xbd, 0xd3, 0x59, 0xdf, 0x78, 0xd7, 0x8d, 0x45, 0x07, 0xc4,
	0xfe, 0x86, 0xdb, 0x12, 0xc0, 0x73, 0x35, 0x41, 0x8c, 0x75, 0x5e, 0x02, 0x9c, 0x00, 0x95, 0x84,
	0x8d, 0xd7, 0x3d, 0x05, 0x6e, 0x24, 0x6c, 0x26, 0x4b, 0xe2, 0x9a, 0x85, 0xd0, 0xcb, 0x47, 0xf5,
	0xf9, 0x45, 0xdf, 0xb3, 0x66, 0xf3, 0xfe, 0x2f, 0xf6, 0xfa, 0x6d, 0xb1, 0xa7, 0x6b, 0x9d, 0x4d,
	0x6f, 0xff, 0xbe, 0xfd, 0x6d, 0x4c, 0x8d, 0x5a, 0xb9, 0x7a, 0xf0, 0x9f, 0x87, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x7e, 0x54, 0x6e, 0x56, 0x5f, 0x02, 0x00, 0x00,
}
