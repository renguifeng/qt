// Code generated by protoc-gen-go.
// source: profile.proto
// DO NOT EDIT!

package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProfileMessage_ProfileType int32

const (
	ProfileMessage_USER  ProfileMessage_ProfileType = 0
	ProfileMessage_ADMIN ProfileMessage_ProfileType = 1
)

var ProfileMessage_ProfileType_name = map[int32]string{
	0: "USER",
	1: "ADMIN",
}
var ProfileMessage_ProfileType_value = map[string]int32{
	"USER":  0,
	"ADMIN": 1,
}

func (x ProfileMessage_ProfileType) String() string {
	return proto.EnumName(ProfileMessage_ProfileType_name, int32(x))
}
func (ProfileMessage_ProfileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{2, 0}
}

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type IDMessage struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *IDMessage) Reset()                    { *m = IDMessage{} }
func (m *IDMessage) String() string            { return proto.CompactTextString(m) }
func (*IDMessage) ProtoMessage()               {}
func (*IDMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *IDMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ProfileMessage struct {
	Id          string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ProfileType ProfileMessage_ProfileType `protobuf:"varint,2,opt,name=profile_type,json=profileType,enum=helloworld.ProfileMessage_ProfileType" json:"profile_type,omitempty"`
	Name        string                     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *ProfileMessage) Reset()                    { *m = ProfileMessage{} }
func (m *ProfileMessage) String() string            { return proto.CompactTextString(m) }
func (*ProfileMessage) ProtoMessage()               {}
func (*ProfileMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ProfileMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProfileMessage) GetProfileType() ProfileMessage_ProfileType {
	if m != nil {
		return m.ProfileType
	}
	return ProfileMessage_USER
}

func (m *ProfileMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyMessage)(nil), "helloworld.EmptyMessage")
	proto.RegisterType((*IDMessage)(nil), "helloworld.IDMessage")
	proto.RegisterType((*ProfileMessage)(nil), "helloworld.ProfileMessage")
	proto.RegisterEnum("helloworld.ProfileMessage_ProfileType", ProfileMessage_ProfileType_name, ProfileMessage_ProfileType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ProfileService service

type ProfileServiceClient interface {
	// List returns all available profiles associated with the account.
	List(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (ProfileService_ListClient, error)
	// Get
	Get(ctx context.Context, in *IDMessage, opts ...grpc.CallOption) (*ProfileMessage, error)
	// Update saves profile.
	Update(ctx context.Context, in *ProfileMessage, opts ...grpc.CallOption) (*IDMessage, error)
}

type profileServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfileServiceClient(cc *grpc.ClientConn) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) List(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (ProfileService_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ProfileService_serviceDesc.Streams[0], c.cc, "/helloworld.ProfileService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &profileServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProfileService_ListClient interface {
	Recv() (*ProfileMessage, error)
	grpc.ClientStream
}

type profileServiceListClient struct {
	grpc.ClientStream
}

func (x *profileServiceListClient) Recv() (*ProfileMessage, error) {
	m := new(ProfileMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profileServiceClient) Get(ctx context.Context, in *IDMessage, opts ...grpc.CallOption) (*ProfileMessage, error) {
	out := new(ProfileMessage)
	err := grpc.Invoke(ctx, "/helloworld.ProfileService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Update(ctx context.Context, in *ProfileMessage, opts ...grpc.CallOption) (*IDMessage, error) {
	out := new(IDMessage)
	err := grpc.Invoke(ctx, "/helloworld.ProfileService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProfileService service

type ProfileServiceServer interface {
	// List returns all available profiles associated with the account.
	List(*EmptyMessage, ProfileService_ListServer) error
	// Get
	Get(context.Context, *IDMessage) (*ProfileMessage, error)
	// Update saves profile.
	Update(context.Context, *ProfileMessage) (*IDMessage, error)
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProfileServiceServer).List(m, &profileServiceListServer{stream})
}

type ProfileService_ListServer interface {
	Send(*ProfileMessage) error
	grpc.ServerStream
}

type profileServiceListServer struct {
	grpc.ServerStream
}

func (x *profileServiceListServer) Send(m *ProfileMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ProfileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.ProfileService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Get(ctx, req.(*IDMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.ProfileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Update(ctx, req.(*ProfileMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ProfileService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProfileService_Update_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ProfileService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "profile.proto",
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0xdf, 0x4a, 0x3a, 0x41,
	0x14, 0xc7, 0x7f, 0xb3, 0xfe, 0xf9, 0xe5, 0x51, 0x37, 0x39, 0x20, 0x2c, 0x5b, 0x17, 0x32, 0x17,
	0x21, 0x5d, 0xb8, 0x62, 0x77, 0xdd, 0x05, 0x4a, 0x08, 0x1a, 0xb1, 0xe6, 0x75, 0x4c, 0xed, 0xc9,
	0x06, 0xd6, 0x9d, 0x61, 0x77, 0x30, 0x24, 0xba, 0xa8, 0x57, 0xe8, 0x11, 0x7a, 0xa4, 0x5e, 0xa1,
	0x07, 0x09, 0x46, 0xd3, 0x0d, 0xd2, 0xbb, 0xf9, 0xc2, 0xf7, 0x7c, 0xce, 0x67, 0x66, 0xa0, 0xae,
	0x53, 0xf5, 0x20, 0x63, 0xea, 0xe8, 0x54, 0x19, 0x85, 0xf0, 0x48, 0x71, 0xac, 0x9e, 0x54, 0x1a,
	0x47, 0xfe, 0xf1, 0x4c, 0xa9, 0x59, 0x4c, 0x81, 0xd0, 0x32, 0x10, 0x49, 0xa2, 0x8c, 0x30, 0x52,
	0x25, 0xd9, 0xaa, 0xc9, 0x5d, 0xa8, 0x0d, 0xe6, 0xda, 0x2c, 0xc7, 0x94, 0x65, 0x62, 0x46, 0xfc,
	0x08, 0x2a, 0xc3, 0xfe, 0x3a, 0xa0, 0x0b, 0x8e, 0x8c, 0x3c, 0xd6, 0x62, 0xed, 0x4a, 0xe8, 0xc8,
	0x88, 0x7f, 0x30, 0x70, 0xaf, 0x57, 0x8b, 0x76, 0x54, 0x70, 0x08, 0xb5, 0xb5, 0xca, 0xad, 0x59,
	0x6a, 0xf2, 0x9c, 0x16, 0x6b, 0xbb, 0xbd, 0x93, 0xce, 0x56, 0xa8, 0xf3, 0x9b, 0xf0, 0x13, 0x6f,
	0x96, 0x9a, 0xc2, 0xaa, 0xde, 0x06, 0x44, 0x28, 0x26, 0x62, 0x4e, 0x5e, 0xc1, 0xc2, 0xed, 0x99,
	0x73, 0xa8, 0xe6, 0xfa, 0x78, 0x00, 0xc5, 0xe9, 0x64, 0x10, 0x36, 0xfe, 0x61, 0x05, 0x4a, 0x17,
	0xfd, 0xf1, 0xf0, 0xaa, 0xc1, 0x7a, 0xaf, 0xce, 0xc6, 0x72, 0x42, 0xe9, 0x42, 0xde, 0x13, 0x8e,
	0xa1, 0x38, 0x92, 0x99, 0x41, 0x2f, 0xef, 0x91, 0xbf, 0xb7, 0xef, 0xef, 0x36, 0xe4, 0xf5, 0xb7,
	0xcf, 0xaf, 0x77, 0xe7, 0x3f, 0x96, 0x82, 0x45, 0x37, 0xd0, 0x5d, 0x86, 0x23, 0x28, 0x5c, 0x92,
	0xc1, 0x66, 0x7e, 0x66, 0xf3, 0x6a, 0x7b, 0x51, 0x68, 0x51, 0x35, 0x04, 0x8b, 0x0a, 0x9e, 0x65,
	0xf4, 0x82, 0x21, 0x94, 0xa7, 0x3a, 0x12, 0x86, 0x70, 0xcf, 0xa4, 0xff, 0xf7, 0x32, 0xde, 0xb4,
	0xc0, 0x43, 0x3f, 0x07, 0x3c, 0x67, 0xa7, 0x77, 0x65, 0xfb, 0xbb, 0x67, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x61, 0xac, 0x88, 0x52, 0x18, 0x02, 0x00, 0x00,
}