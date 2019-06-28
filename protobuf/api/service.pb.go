// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/api/service.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	types "github.com/mesg-foundation/engine/protobuf/types"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request's data for the `Create` API.
type CreateServiceRequest struct {
	// The service definition to use to create the Service.
	Definition           *types.Service `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateServiceRequest) Reset()         { *m = CreateServiceRequest{} }
func (m *CreateServiceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceRequest) ProtoMessage()    {}
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{0}
}

func (m *CreateServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceRequest.Unmarshal(m, b)
}
func (m *CreateServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceRequest.Marshal(b, m, deterministic)
}
func (m *CreateServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceRequest.Merge(m, src)
}
func (m *CreateServiceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServiceRequest.Size(m)
}
func (m *CreateServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceRequest proto.InternalMessageInfo

func (m *CreateServiceRequest) GetDefinition() *types.Service {
	if m != nil {
		return m.Definition
	}
	return nil
}

// The response's data for the `Create` API.
type CreateServiceResponse struct {
	// The service's hash created.
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceResponse) Reset()         { *m = CreateServiceResponse{} }
func (m *CreateServiceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServiceResponse) ProtoMessage()    {}
func (*CreateServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{1}
}

func (m *CreateServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceResponse.Unmarshal(m, b)
}
func (m *CreateServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceResponse.Marshal(b, m, deterministic)
}
func (m *CreateServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceResponse.Merge(m, src)
}
func (m *CreateServiceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServiceResponse.Size(m)
}
func (m *CreateServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceResponse proto.InternalMessageInfo

func (m *CreateServiceResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// The request's data for the `Delete` API.
type DeleteServiceRequest struct {
	// The service's hash to delete.
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteServiceRequest) Reset()         { *m = DeleteServiceRequest{} }
func (m *DeleteServiceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteServiceRequest) ProtoMessage()    {}
func (*DeleteServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{2}
}

func (m *DeleteServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteServiceRequest.Unmarshal(m, b)
}
func (m *DeleteServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteServiceRequest.Marshal(b, m, deterministic)
}
func (m *DeleteServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteServiceRequest.Merge(m, src)
}
func (m *DeleteServiceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteServiceRequest.Size(m)
}
func (m *DeleteServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteServiceRequest proto.InternalMessageInfo

func (m *DeleteServiceRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// The response's data for the `Delete` API, doesn't contain anything.
type DeleteServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteServiceResponse) Reset()         { *m = DeleteServiceResponse{} }
func (m *DeleteServiceResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteServiceResponse) ProtoMessage()    {}
func (*DeleteServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{3}
}

func (m *DeleteServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteServiceResponse.Unmarshal(m, b)
}
func (m *DeleteServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteServiceResponse.Marshal(b, m, deterministic)
}
func (m *DeleteServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteServiceResponse.Merge(m, src)
}
func (m *DeleteServiceResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteServiceResponse.Size(m)
}
func (m *DeleteServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteServiceResponse proto.InternalMessageInfo

// The request's data for the `Get` API.
type GetServiceRequest struct {
	// The service's hash to fetch.
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceRequest) Reset()         { *m = GetServiceRequest{} }
func (m *GetServiceRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceRequest) ProtoMessage()    {}
func (*GetServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{4}
}

func (m *GetServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceRequest.Unmarshal(m, b)
}
func (m *GetServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceRequest.Marshal(b, m, deterministic)
}
func (m *GetServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceRequest.Merge(m, src)
}
func (m *GetServiceRequest) XXX_Size() int {
	return xxx_messageInfo_GetServiceRequest.Size(m)
}
func (m *GetServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceRequest proto.InternalMessageInfo

func (m *GetServiceRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// The request's data for the `List` API.
type ListServiceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListServiceRequest) Reset()         { *m = ListServiceRequest{} }
func (m *ListServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ListServiceRequest) ProtoMessage()    {}
func (*ListServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{5}
}

func (m *ListServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServiceRequest.Unmarshal(m, b)
}
func (m *ListServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServiceRequest.Marshal(b, m, deterministic)
}
func (m *ListServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServiceRequest.Merge(m, src)
}
func (m *ListServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ListServiceRequest.Size(m)
}
func (m *ListServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListServiceRequest proto.InternalMessageInfo

// The response's data for the `List` API.
type ListServiceResponse struct {
	// List of services that match the request's filters.
	Services             []*types.Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListServiceResponse) Reset()         { *m = ListServiceResponse{} }
func (m *ListServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ListServiceResponse) ProtoMessage()    {}
func (*ListServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{6}
}

func (m *ListServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServiceResponse.Unmarshal(m, b)
}
func (m *ListServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServiceResponse.Marshal(b, m, deterministic)
}
func (m *ListServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServiceResponse.Merge(m, src)
}
func (m *ListServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ListServiceResponse.Size(m)
}
func (m *ListServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListServiceResponse proto.InternalMessageInfo

func (m *ListServiceResponse) GetServices() []*types.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateServiceRequest)(nil), "api.CreateServiceRequest")
	proto.RegisterType((*CreateServiceResponse)(nil), "api.CreateServiceResponse")
	proto.RegisterType((*DeleteServiceRequest)(nil), "api.DeleteServiceRequest")
	proto.RegisterType((*DeleteServiceResponse)(nil), "api.DeleteServiceResponse")
	proto.RegisterType((*GetServiceRequest)(nil), "api.GetServiceRequest")
	proto.RegisterType((*ListServiceRequest)(nil), "api.ListServiceRequest")
	proto.RegisterType((*ListServiceResponse)(nil), "api.ListServiceResponse")
}

func init() { proto.RegisterFile("protobuf/api/service.proto", fileDescriptor_0615fe53b372bcb1) }

var fileDescriptor_0615fe53b372bcb1 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x5b, 0x3a, 0xa6, 0x3e, 0x41, 0xf0, 0xd9, 0xb9, 0x1a, 0x3c, 0x8c, 0x5c, 0x94, 0x09,
	0x29, 0xcc, 0xa3, 0xa7, 0xa1, 0xb8, 0x8b, 0xa7, 0xf9, 0x09, 0x32, 0x7d, 0x63, 0x01, 0x69, 0x63,
	0x93, 0x09, 0x7e, 0x78, 0x41, 0x96, 0x66, 0xc3, 0x36, 0x39, 0x78, 0x2b, 0xef, 0xfd, 0xfa, 0x7b,
	0xfd, 0xff, 0x29, 0x30, 0xdd, 0xd4, 0xb6, 0x5e, 0x6d, 0xd7, 0xa5, 0xd4, 0xaa, 0x34, 0xd4, 0x7c,
	0xa9, 0x37, 0x12, 0x6e, 0x88, 0x99, 0xd4, 0x8a, 0x5d, 0x1f, 0x00, 0xfb, 0xad, 0xc9, 0x74, 0x11,
	0xfe, 0x0c, 0xf9, 0x63, 0x43, 0xd2, 0xd2, 0x6b, 0x3b, 0x5e, 0xd2, 0xe7, 0x96, 0x8c, 0x45, 0x01,
	0xf0, 0x4e, 0x6b, 0x55, 0x29, 0xab, 0xea, 0xaa, 0x48, 0x27, 0xe9, 0xed, 0xe9, 0xec, 0x4c, 0x38,
	0x83, 0xd8, 0xa3, 0x7f, 0x08, 0x7e, 0x07, 0xa3, 0x9e, 0xc7, 0xe8, 0xba, 0x32, 0x84, 0x08, 0x83,
	0x8d, 0x34, 0x1b, 0xa7, 0x38, 0x59, 0xba, 0x67, 0x3e, 0x85, 0xfc, 0x89, 0x3e, 0x28, 0x38, 0x1a,
	0x63, 0xc7, 0x30, 0xea, 0xb1, 0xad, 0x98, 0xdf, 0xc0, 0xf9, 0x82, 0xec, 0x3f, 0x0c, 0x39, 0xe0,
	0x8b, 0x32, 0x3d, 0x92, 0xcf, 0xe1, 0xa2, 0x33, 0xf5, 0x9f, 0x3b, 0x85, 0x63, 0x5f, 0x90, 0x29,
	0xd2, 0x49, 0x16, 0x49, 0x7d, 0xd8, 0xcf, 0x7e, 0x52, 0x38, 0xf2, 0x53, 0x9c, 0xc3, 0xb0, 0xcd,
	0x8f, 0x57, 0x42, 0x6a, 0x25, 0x62, 0xa5, 0x32, 0x16, 0x5b, 0xf9, 0x38, 0xc9, 0x4e, 0xd1, 0x26,
	0xf5, 0x8a, 0x58, 0x45, 0x5e, 0x11, 0x6f, 0x24, 0xc1, 0x12, 0xb2, 0x05, 0x59, 0xbc, 0x74, 0x50,
	0xd0, 0x0e, 0xeb, 0x45, 0xe1, 0x09, 0x3e, 0xc0, 0x60, 0xd7, 0x02, 0x8e, 0xdd, 0x1b, 0x61, 0x4d,
	0xac, 0x08, 0x17, 0xfb, 0x6b, 0xab, 0xa1, 0xfb, 0x85, 0xee, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x96, 0x7f, 0x64, 0x89, 0x83, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// Create a Service from a Service Definition.
	// It will return an unique identifier which is used to interact with the Service.
	Create(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
	// Delete a Service.
	// An error is returned if one or more Instances of the Service are running.
	Delete(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error)
	// Get returns a Service matching the criteria of the request.
	Get(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*types.Service, error)
	// List returns services specified in a request.
	List(ctx context.Context, in *ListServiceRequest, opts ...grpc.CallOption) (*ListServiceResponse, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Create(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, "/api.Service/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Delete(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error) {
	out := new(DeleteServiceResponse)
	err := c.cc.Invoke(ctx, "/api.Service/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Get(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*types.Service, error) {
	out := new(types.Service)
	err := c.cc.Invoke(ctx, "/api.Service/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) List(ctx context.Context, in *ListServiceRequest, opts ...grpc.CallOption) (*ListServiceResponse, error) {
	out := new(ListServiceResponse)
	err := c.cc.Invoke(ctx, "/api.Service/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// Create a Service from a Service Definition.
	// It will return an unique identifier which is used to interact with the Service.
	Create(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	// Delete a Service.
	// An error is returned if one or more Instances of the Service are running.
	Delete(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error)
	// Get returns a Service matching the criteria of the request.
	Get(context.Context, *GetServiceRequest) (*types.Service, error)
	// List returns services specified in a request.
	List(context.Context, *ListServiceRequest) (*ListServiceResponse, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Service/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Create(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Service/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Delete(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Service/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Get(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Service/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).List(ctx, req.(*ListServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Service_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Service_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Service_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Service_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/service.proto",
}
