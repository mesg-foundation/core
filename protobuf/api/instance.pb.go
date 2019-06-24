// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/api/instance.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import definition "github.com/mesg-foundation/core/protobuf/definition"

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

// The request's data for the `Get` API.
//
// **Example**
// TODO: add JSON example
type GetInstanceRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInstanceRequest) Reset()         { *m = GetInstanceRequest{} }
func (m *GetInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstanceRequest) ProtoMessage()    {}
func (*GetInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_278b77604059a607, []int{0}
}
func (m *GetInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInstanceRequest.Unmarshal(m, b)
}
func (m *GetInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInstanceRequest.Marshal(b, m, deterministic)
}
func (dst *GetInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInstanceRequest.Merge(dst, src)
}
func (m *GetInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_GetInstanceRequest.Size(m)
}
func (m *GetInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInstanceRequest proto.InternalMessageInfo

func (m *GetInstanceRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetInstanceResponse struct {
	Instance             *definition.Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetInstanceResponse) Reset()         { *m = GetInstanceResponse{} }
func (m *GetInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*GetInstanceResponse) ProtoMessage()    {}
func (*GetInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_278b77604059a607, []int{1}
}
func (m *GetInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInstanceResponse.Unmarshal(m, b)
}
func (m *GetInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInstanceResponse.Marshal(b, m, deterministic)
}
func (dst *GetInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInstanceResponse.Merge(dst, src)
}
func (m *GetInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_GetInstanceResponse.Size(m)
}
func (m *GetInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInstanceResponse proto.InternalMessageInfo

func (m *GetInstanceResponse) GetInstance() *definition.Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

// The request's data for the `List` API.
//
// **Example**
// TODO: add JSON example
type ListInstancesRequest struct {
	ServiceHash          string   `protobuf:"bytes,1,opt,name=serviceHash,proto3" json:"serviceHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListInstancesRequest) Reset()         { *m = ListInstancesRequest{} }
func (m *ListInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*ListInstancesRequest) ProtoMessage()    {}
func (*ListInstancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_278b77604059a607, []int{2}
}
func (m *ListInstancesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstancesRequest.Unmarshal(m, b)
}
func (m *ListInstancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstancesRequest.Marshal(b, m, deterministic)
}
func (dst *ListInstancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstancesRequest.Merge(dst, src)
}
func (m *ListInstancesRequest) XXX_Size() int {
	return xxx_messageInfo_ListInstancesRequest.Size(m)
}
func (m *ListInstancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstancesRequest proto.InternalMessageInfo

func (m *ListInstancesRequest) GetServiceHash() string {
	if m != nil {
		return m.ServiceHash
	}
	return ""
}

// The response's data for the `List` API.
//
// **Example**
// TODO: add JSON example
type ListInstancesResponse struct {
	Instances            []*definition.Instance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListInstancesResponse) Reset()         { *m = ListInstancesResponse{} }
func (m *ListInstancesResponse) String() string { return proto.CompactTextString(m) }
func (*ListInstancesResponse) ProtoMessage()    {}
func (*ListInstancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_278b77604059a607, []int{3}
}
func (m *ListInstancesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstancesResponse.Unmarshal(m, b)
}
func (m *ListInstancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstancesResponse.Marshal(b, m, deterministic)
}
func (dst *ListInstancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstancesResponse.Merge(dst, src)
}
func (m *ListInstancesResponse) XXX_Size() int {
	return xxx_messageInfo_ListInstancesResponse.Size(m)
}
func (m *ListInstancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstancesResponse proto.InternalMessageInfo

func (m *ListInstancesResponse) GetInstances() []*definition.Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

// The request's data for the `Create` API.
//
// **Example**
// TODO: add JSON example
type CreateInstanceRequest struct {
	ServiceHash          string   `protobuf:"bytes,1,opt,name=serviceHash,proto3" json:"serviceHash,omitempty"`
	Env                  []string `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateInstanceRequest) Reset()         { *m = CreateInstanceRequest{} }
func (m *CreateInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateInstanceRequest) ProtoMessage()    {}
func (*CreateInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_278b77604059a607, []int{4}
}
func (m *CreateInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInstanceRequest.Unmarshal(m, b)
}
func (m *CreateInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInstanceRequest.Marshal(b, m, deterministic)
}
func (dst *CreateInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInstanceRequest.Merge(dst, src)
}
func (m *CreateInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateInstanceRequest.Size(m)
}
func (m *CreateInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInstanceRequest proto.InternalMessageInfo

func (m *CreateInstanceRequest) GetServiceHash() string {
	if m != nil {
		return m.ServiceHash
	}
	return ""
}

func (m *CreateInstanceRequest) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

// The response's data for the `Create` API.
//
// **Example**
// TODO: add JSON example
type CreateInstanceResponse struct {
	Instance             *definition.Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateInstanceResponse) Reset()         { *m = CreateInstanceResponse{} }
func (m *CreateInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateInstanceResponse) ProtoMessage()    {}
func (*CreateInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_278b77604059a607, []int{5}
}
func (m *CreateInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInstanceResponse.Unmarshal(m, b)
}
func (m *CreateInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInstanceResponse.Marshal(b, m, deterministic)
}
func (dst *CreateInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInstanceResponse.Merge(dst, src)
}
func (m *CreateInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateInstanceResponse.Size(m)
}
func (m *CreateInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInstanceResponse proto.InternalMessageInfo

func (m *CreateInstanceResponse) GetInstance() *definition.Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

// The request's data for the `Delete` API.
//
// **Example**
// TODO: add JSON example
type DeleteInstanceRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	DeleteData           bool     `protobuf:"varint,2,opt,name=deleteData,proto3" json:"deleteData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteInstanceRequest) Reset()         { *m = DeleteInstanceRequest{} }
func (m *DeleteInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteInstanceRequest) ProtoMessage()    {}
func (*DeleteInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_278b77604059a607, []int{6}
}
func (m *DeleteInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteInstanceRequest.Unmarshal(m, b)
}
func (m *DeleteInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteInstanceRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteInstanceRequest.Merge(dst, src)
}
func (m *DeleteInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteInstanceRequest.Size(m)
}
func (m *DeleteInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteInstanceRequest proto.InternalMessageInfo

func (m *DeleteInstanceRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *DeleteInstanceRequest) GetDeleteData() bool {
	if m != nil {
		return m.DeleteData
	}
	return false
}

// The response's data for the `Delete` API.
//
// **Example**
// TODO: add JSON example
type DeleteInstanceResponse struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteInstanceResponse) Reset()         { *m = DeleteInstanceResponse{} }
func (m *DeleteInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteInstanceResponse) ProtoMessage()    {}
func (*DeleteInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_278b77604059a607, []int{7}
}
func (m *DeleteInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteInstanceResponse.Unmarshal(m, b)
}
func (m *DeleteInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteInstanceResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteInstanceResponse.Merge(dst, src)
}
func (m *DeleteInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteInstanceResponse.Size(m)
}
func (m *DeleteInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteInstanceResponse proto.InternalMessageInfo

func (m *DeleteInstanceResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetInstanceRequest)(nil), "api.GetInstanceRequest")
	proto.RegisterType((*GetInstanceResponse)(nil), "api.GetInstanceResponse")
	proto.RegisterType((*ListInstancesRequest)(nil), "api.ListInstancesRequest")
	proto.RegisterType((*ListInstancesResponse)(nil), "api.ListInstancesResponse")
	proto.RegisterType((*CreateInstanceRequest)(nil), "api.CreateInstanceRequest")
	proto.RegisterType((*CreateInstanceResponse)(nil), "api.CreateInstanceResponse")
	proto.RegisterType((*DeleteInstanceRequest)(nil), "api.DeleteInstanceRequest")
	proto.RegisterType((*DeleteInstanceResponse)(nil), "api.DeleteInstanceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InstanceClient is the client API for Instance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InstanceClient interface {
	// Get returns an Instance matching the criteria of the request.
	Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error)
	// List returns all Instances matching the criteria of the request.
	List(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	// Create an Instance from a Service's hash and custom environmental variables.
	// It will return an unique identifier which is used to interact with the Instance.
	Create(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error)
	// Delete an Instance.
	Delete(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceResponse, error)
}

type instanceClient struct {
	cc *grpc.ClientConn
}

func NewInstanceClient(cc *grpc.ClientConn) InstanceClient {
	return &instanceClient{cc}
}

func (c *instanceClient) Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error) {
	out := new(GetInstanceResponse)
	err := c.cc.Invoke(ctx, "/api.Instance/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) List(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := c.cc.Invoke(ctx, "/api.Instance/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) Create(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error) {
	out := new(CreateInstanceResponse)
	err := c.cc.Invoke(ctx, "/api.Instance/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) Delete(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceResponse, error) {
	out := new(DeleteInstanceResponse)
	err := c.cc.Invoke(ctx, "/api.Instance/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceServer is the server API for Instance service.
type InstanceServer interface {
	// Get returns an Instance matching the criteria of the request.
	Get(context.Context, *GetInstanceRequest) (*GetInstanceResponse, error)
	// List returns all Instances matching the criteria of the request.
	List(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	// Create an Instance from a Service's hash and custom environmental variables.
	// It will return an unique identifier which is used to interact with the Instance.
	Create(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error)
	// Delete an Instance.
	Delete(context.Context, *DeleteInstanceRequest) (*DeleteInstanceResponse, error)
}

func RegisterInstanceServer(s *grpc.Server, srv InstanceServer) {
	s.RegisterService(&_Instance_serviceDesc, srv)
}

func _Instance_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Instance/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Get(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Instance/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).List(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Instance/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Create(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Instance/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Delete(ctx, req.(*DeleteInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Instance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Instance",
	HandlerType: (*InstanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Instance_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Instance_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Instance_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Instance_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/instance.proto",
}

func init() {
	proto.RegisterFile("protobuf/api/instance.proto", fileDescriptor_instance_278b77604059a607)
}

var fileDescriptor_instance_278b77604059a607 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x4e, 0xf2, 0x40,
	0x10, 0xc0, 0x29, 0x25, 0x04, 0x86, 0xcb, 0x97, 0xf9, 0x00, 0x6b, 0x49, 0x4c, 0xb3, 0xa7, 0x1e,
	0x4c, 0x31, 0x78, 0x31, 0x5e, 0x3c, 0x40, 0x82, 0x8a, 0xa7, 0xbe, 0xc1, 0x02, 0x43, 0xd8, 0xc4,
	0x6c, 0x2b, 0xbb, 0xf0, 0x1e, 0xbe, 0xb1, 0x61, 0xfb, 0x87, 0xda, 0xae, 0x31, 0xf1, 0xb6, 0x99,
	0xf9, 0xf5, 0x37, 0xb3, 0x33, 0x5d, 0x98, 0xa4, 0x87, 0x44, 0x27, 0xeb, 0xe3, 0x6e, 0xca, 0x53,
	0x31, 0x15, 0x52, 0x69, 0x2e, 0x37, 0x14, 0x99, 0x28, 0xba, 0x3c, 0x15, 0x3e, 0x2b, 0x89, 0x2d,
	0xed, 0x84, 0x14, 0x5a, 0x24, 0xb2, 0x06, 0xb2, 0x10, 0x70, 0x49, 0xfa, 0x25, 0x0f, 0xc6, 0xf4,
	0x71, 0x24, 0xa5, 0x11, 0xa1, 0xb3, 0xe7, 0x6a, 0xef, 0x39, 0x81, 0x13, 0xf6, 0x63, 0x73, 0x66,
	0x4b, 0xf8, 0xff, 0x8d, 0x54, 0x69, 0x22, 0x15, 0xe1, 0x1d, 0xf4, 0x0a, 0xa5, 0xc1, 0x07, 0xb3,
	0x61, 0x74, 0x29, 0x17, 0x95, 0x7c, 0x49, 0xb1, 0x07, 0x18, 0xbe, 0x09, 0x55, 0x9a, 0x54, 0x51,
	0x34, 0x80, 0x81, 0xa2, 0xc3, 0x49, 0x6c, 0xe8, 0xf9, 0x52, 0xbb, 0x1a, 0x62, 0x2b, 0x18, 0xd5,
	0xbe, 0xcc, 0x9b, 0x98, 0x41, 0xbf, 0xd0, 0x2b, 0xcf, 0x09, 0xdc, 0x1f, 0xbb, 0xb8, 0x60, 0x67,
	0xd9, 0xfc, 0x40, 0x5c, 0x53, 0xfd, 0xf2, 0xbf, 0xf6, 0x81, 0xff, 0xc0, 0x25, 0x79, 0xf2, 0xda,
	0x81, 0x1b, 0xf6, 0xe3, 0xf3, 0x91, 0xbd, 0xc2, 0xb8, 0x2e, 0xfb, 0xf3, 0x7c, 0x56, 0x30, 0x5a,
	0xd0, 0x3b, 0x35, 0x1b, 0xb3, 0x6c, 0x05, 0x6f, 0x00, 0xb6, 0x06, 0x5e, 0x70, 0xcd, 0xbd, 0x76,
	0xe0, 0x84, 0xbd, 0xb8, 0x12, 0x61, 0xb7, 0x30, 0xae, 0xcb, 0xf2, 0xc6, 0x2c, 0xb6, 0xd9, 0x67,
	0x1b, 0x7a, 0x05, 0x88, 0x8f, 0xe0, 0x2e, 0x49, 0xe3, 0x55, 0xc4, 0x53, 0x11, 0x35, 0x7f, 0x12,
	0xdf, 0x6b, 0x26, 0x32, 0x35, 0x6b, 0xe1, 0x13, 0x74, 0xce, 0x9b, 0xc2, 0x6b, 0xc3, 0xd8, 0xd6,
	0xed, 0xfb, 0xb6, 0x54, 0x29, 0x98, 0x43, 0x37, 0x1b, 0x28, 0x66, 0x9c, 0x75, 0x55, 0xfe, 0xc4,
	0x9a, 0xab, 0x4a, 0xb2, 0xcb, 0xe7, 0x12, 0xeb, 0x58, 0x73, 0x89, 0x7d, 0x4a, 0xac, 0xb5, 0xee,
	0x9a, 0x87, 0x72, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x18, 0x2e, 0xfb, 0x70, 0x03, 0x00,
	0x00,
}
