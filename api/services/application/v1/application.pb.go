// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/application/v1/application.proto

/*
Package application is a generated protocol buffer package.

It is generated from these files:
	github.com/ehazlett/stellar/api/services/application/v1/application.proto

It has these top-level messages:
	CreateRequest
	DeleteRequest
	ListRequest
	App
	ListResponse
	GetRequest
	GetResponse
	RestartRequest
*/
package application

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import stellar_services_node_v1 "github.com/ehazlett/stellar/api/services/node/v1"

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

type CreateRequest struct {
	Name     string                              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels   []string                            `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
	Services []*stellar_services_node_v1.Service `protobuf:"bytes,3,rep,name=services" json:"services,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{0} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreateRequest) GetServices() []*stellar_services_node_v1.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type DeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{1} }

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{2} }

type App struct {
	Name     string                              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Services []*stellar_services_node_v1.Service `protobuf:"bytes,2,rep,name=services" json:"services,omitempty"`
}

func (m *App) Reset()                    { *m = App{} }
func (m *App) String() string            { return proto.CompactTextString(m) }
func (*App) ProtoMessage()               {}
func (*App) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{3} }

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *App) GetServices() []*stellar_services_node_v1.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type ListResponse struct {
	Applications []*App `protobuf:"bytes,1,rep,name=applications" json:"applications,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{4} }

func (m *ListResponse) GetApplications() []*App {
	if m != nil {
		return m.Applications
	}
	return nil
}

type GetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{5} }

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetResponse struct {
	Application *App `protobuf:"bytes,1,opt,name=application" json:"application,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{6} }

func (m *GetResponse) GetApplication() *App {
	if m != nil {
		return m.Application
	}
	return nil
}

type RestartRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *RestartRequest) Reset()                    { *m = RestartRequest{} }
func (m *RestartRequest) String() string            { return proto.CompactTextString(m) }
func (*RestartRequest) ProtoMessage()               {}
func (*RestartRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{7} }

func (m *RestartRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "stellar.services.application.v1.CreateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "stellar.services.application.v1.DeleteRequest")
	proto.RegisterType((*ListRequest)(nil), "stellar.services.application.v1.ListRequest")
	proto.RegisterType((*App)(nil), "stellar.services.application.v1.App")
	proto.RegisterType((*ListResponse)(nil), "stellar.services.application.v1.ListResponse")
	proto.RegisterType((*GetRequest)(nil), "stellar.services.application.v1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "stellar.services.application.v1.GetResponse")
	proto.RegisterType((*RestartRequest)(nil), "stellar.services.application.v1.RestartRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Application service

type ApplicationClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type applicationClient struct {
	cc *grpc.ClientConn
}

func NewApplicationClient(cc *grpc.ClientConn) ApplicationClient {
	return &applicationClient{cc}
}

func (c *applicationClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Restart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Application service

type ApplicationServer interface {
	Create(context.Context, *CreateRequest) (*google_protobuf.Empty, error)
	Delete(context.Context, *DeleteRequest) (*google_protobuf.Empty, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Restart(context.Context, *RestartRequest) (*google_protobuf.Empty, error)
}

func RegisterApplicationServer(s *grpc.Server, srv ApplicationServer) {
	s.RegisterService(&_Application_serviceDesc, srv)
}

func _Application_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Restart(ctx, req.(*RestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Application_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.application.v1.Application",
	HandlerType: (*ApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Application_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Application_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Application_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Application_Get_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Application_Restart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/application/v1/application.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/application/v1/application.proto", fileDescriptorApplication)
}

var fileDescriptorApplication = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xa6, 0x9b, 0xa5, 0xba, 0x2f, 0x5b, 0x0f, 0x73, 0x58, 0x42, 0x3c, 0x18, 0xe3, 0x1e, 0x0a,
	0xae, 0x33, 0x74, 0x3d, 0x2e, 0x1e, 0x6a, 0xad, 0x55, 0xf0, 0x20, 0x11, 0xa1, 0x78, 0x72, 0x12,
	0x9f, 0x6d, 0x60, 0x9a, 0x19, 0x33, 0xd3, 0x80, 0xfd, 0x1b, 0xfd, 0xa3, 0x24, 0x3f, 0x6a, 0x27,
	0x68, 0x9b, 0xb8, 0xa7, 0xe4, 0x85, 0xef, 0xc7, 0x9b, 0x7c, 0x1f, 0x03, 0xef, 0x57, 0xa9, 0x59,
	0x6f, 0x63, 0x9a, 0xc8, 0x0d, 0xc3, 0x35, 0xdf, 0x09, 0x34, 0x86, 0x69, 0x83, 0x42, 0xf0, 0x9c,
	0x71, 0x95, 0x32, 0x8d, 0x79, 0x91, 0x26, 0xa8, 0x19, 0x57, 0x4a, 0xa4, 0x09, 0x37, 0xa9, 0xcc,
	0x58, 0x31, 0xb1, 0x47, 0xaa, 0x72, 0x69, 0x24, 0x79, 0xd2, 0xd0, 0xe8, 0x9e, 0x42, 0x6d, 0x4c,
	0x31, 0xf1, 0x1f, 0xaf, 0xa4, 0x5c, 0x09, 0x64, 0x15, 0x3c, 0xde, 0x7e, 0x67, 0xb8, 0x51, 0xe6,
	0x67, 0xcd, 0xf6, 0xef, 0x7a, 0x2f, 0x92, 0xc9, 0x6f, 0x58, 0x6e, 0x50, 0x3e, 0x6b, 0x72, 0xb8,
	0x83, 0xd1, 0x2c, 0x47, 0x6e, 0x30, 0xc2, 0x1f, 0x5b, 0xd4, 0x86, 0x10, 0x38, 0xcf, 0xf8, 0x06,
	0xbd, 0x41, 0x30, 0x18, 0x5f, 0x44, 0xd5, 0x3b, 0xb9, 0x82, 0xa1, 0xe0, 0x31, 0x0a, 0xed, 0x9d,
	0x05, 0xce, 0xf8, 0x22, 0x6a, 0x26, 0xf2, 0x0a, 0x1e, 0xee, 0xb5, 0x3d, 0x27, 0x70, 0xc6, 0xee,
	0xed, 0x53, 0xfa, 0xd7, 0x51, 0x2a, 0xb3, 0x62, 0x42, 0x3f, 0xd5, 0x1f, 0xa2, 0x3f, 0x94, 0xf0,
	0x19, 0x8c, 0xde, 0xa0, 0xc0, 0x93, 0xde, 0xe1, 0x08, 0xdc, 0x0f, 0xa9, 0x36, 0x0d, 0x24, 0x5c,
	0x82, 0x33, 0x55, 0xea, 0x9f, 0x5b, 0xda, 0xdb, 0x9c, 0xfd, 0xff, 0x36, 0x4b, 0xb8, 0xac, 0x8d,
	0xb4, 0x92, 0x99, 0x46, 0xf2, 0x0e, 0x2e, 0xad, 0x14, 0xb4, 0x37, 0xa8, 0x24, 0xaf, 0x69, 0x47,
	0x56, 0x74, 0xaa, 0x54, 0xd4, 0x62, 0x86, 0x01, 0xc0, 0x02, 0xcd, 0xa9, 0x43, 0x7e, 0x06, 0xb7,
	0x42, 0x34, 0xd6, 0x6f, 0xc1, 0xb5, 0x04, 0x2a, 0x64, 0x5f, 0x67, 0x9b, 0x18, 0x5e, 0xc3, 0xa3,
	0x08, 0xb5, 0xe1, 0xf9, 0x29, 0xf3, 0xdb, 0x5f, 0x0e, 0xb8, 0xd3, 0x03, 0x8b, 0x7c, 0x84, 0x61,
	0x5d, 0x09, 0x42, 0x3b, 0x2d, 0x5b, 0xdd, 0xf1, 0xaf, 0x68, 0xdd, 0x53, 0xba, 0xef, 0x29, 0x9d,
	0x97, 0x3d, 0x2d, 0x15, 0xeb, 0xa0, 0x7b, 0x28, 0xb6, 0x1a, 0x71, 0x54, 0x31, 0x81, 0xf3, 0x32,
	0x2c, 0x72, 0xd3, 0xa9, 0x67, 0x95, 0xc7, 0x7f, 0xd1, 0x13, 0xdd, 0xc4, 0xf0, 0x15, 0x9c, 0x05,
	0x1a, 0xf2, 0xbc, 0x93, 0x75, 0x48, 0xd7, 0xbf, 0xe9, 0x07, 0x6e, 0x1c, 0x22, 0x78, 0xd0, 0x04,
	0x44, 0x58, 0x27, 0xb1, 0x1d, 0xe5, 0xb1, 0x5f, 0xf3, 0x7a, 0xfe, 0x65, 0x76, 0xcf, 0x9b, 0xe9,
	0xce, 0x1a, 0xe3, 0x61, 0x25, 0xfb, 0xf2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x43, 0xfe,
	0x87, 0xe7, 0x04, 0x00, 0x00,
}
