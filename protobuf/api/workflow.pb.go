// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/api/workflow.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	types "github.com/mesg-foundation/engine/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
type CreateWorkflowRequest struct {
	Key                  string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Nodes                []*types.Workflow_Node `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Edges                []*types.Workflow_Edge `protobuf:"bytes,5,rep,name=edges,proto3" json:"edges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateWorkflowRequest) Reset()         { *m = CreateWorkflowRequest{} }
func (m *CreateWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWorkflowRequest) ProtoMessage()    {}
func (*CreateWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2568c098005b7b27, []int{0}
}

func (m *CreateWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWorkflowRequest.Unmarshal(m, b)
}
func (m *CreateWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *CreateWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWorkflowRequest.Merge(m, src)
}
func (m *CreateWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWorkflowRequest.Size(m)
}
func (m *CreateWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWorkflowRequest proto.InternalMessageInfo

func (m *CreateWorkflowRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CreateWorkflowRequest) GetNodes() []*types.Workflow_Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *CreateWorkflowRequest) GetEdges() []*types.Workflow_Edge {
	if m != nil {
		return m.Edges
	}
	return nil
}

// The response's data for the `Create` API.
type CreateWorkflowResponse struct {
	// The workflow's hash created.
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWorkflowResponse) Reset()         { *m = CreateWorkflowResponse{} }
func (m *CreateWorkflowResponse) String() string { return proto.CompactTextString(m) }
func (*CreateWorkflowResponse) ProtoMessage()    {}
func (*CreateWorkflowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2568c098005b7b27, []int{1}
}

func (m *CreateWorkflowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWorkflowResponse.Unmarshal(m, b)
}
func (m *CreateWorkflowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWorkflowResponse.Marshal(b, m, deterministic)
}
func (m *CreateWorkflowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWorkflowResponse.Merge(m, src)
}
func (m *CreateWorkflowResponse) XXX_Size() int {
	return xxx_messageInfo_CreateWorkflowResponse.Size(m)
}
func (m *CreateWorkflowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWorkflowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWorkflowResponse proto.InternalMessageInfo

func (m *CreateWorkflowResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// The request's data for the `Delete` API.
type DeleteWorkflowRequest struct {
	// The workflow's hash to delete.
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteWorkflowRequest) Reset()         { *m = DeleteWorkflowRequest{} }
func (m *DeleteWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteWorkflowRequest) ProtoMessage()    {}
func (*DeleteWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2568c098005b7b27, []int{2}
}

func (m *DeleteWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteWorkflowRequest.Unmarshal(m, b)
}
func (m *DeleteWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *DeleteWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteWorkflowRequest.Merge(m, src)
}
func (m *DeleteWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteWorkflowRequest.Size(m)
}
func (m *DeleteWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteWorkflowRequest proto.InternalMessageInfo

func (m *DeleteWorkflowRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// The response's data for the `Delete` API, doesn't contain anything.
type DeleteWorkflowResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteWorkflowResponse) Reset()         { *m = DeleteWorkflowResponse{} }
func (m *DeleteWorkflowResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteWorkflowResponse) ProtoMessage()    {}
func (*DeleteWorkflowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2568c098005b7b27, []int{3}
}

func (m *DeleteWorkflowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteWorkflowResponse.Unmarshal(m, b)
}
func (m *DeleteWorkflowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteWorkflowResponse.Marshal(b, m, deterministic)
}
func (m *DeleteWorkflowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteWorkflowResponse.Merge(m, src)
}
func (m *DeleteWorkflowResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteWorkflowResponse.Size(m)
}
func (m *DeleteWorkflowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteWorkflowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteWorkflowResponse proto.InternalMessageInfo

// The request's data for the `Get` API.
type GetWorkflowRequest struct {
	// The workflow's hash to fetch.
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWorkflowRequest) Reset()         { *m = GetWorkflowRequest{} }
func (m *GetWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*GetWorkflowRequest) ProtoMessage()    {}
func (*GetWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2568c098005b7b27, []int{4}
}

func (m *GetWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWorkflowRequest.Unmarshal(m, b)
}
func (m *GetWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *GetWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWorkflowRequest.Merge(m, src)
}
func (m *GetWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_GetWorkflowRequest.Size(m)
}
func (m *GetWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWorkflowRequest proto.InternalMessageInfo

func (m *GetWorkflowRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// The request's data for the `List` API.
type ListWorkflowRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWorkflowRequest) Reset()         { *m = ListWorkflowRequest{} }
func (m *ListWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowRequest) ProtoMessage()    {}
func (*ListWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2568c098005b7b27, []int{5}
}

func (m *ListWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowRequest.Unmarshal(m, b)
}
func (m *ListWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *ListWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowRequest.Merge(m, src)
}
func (m *ListWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowRequest.Size(m)
}
func (m *ListWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowRequest proto.InternalMessageInfo

// The response's data for the `List` API.
type ListWorkflowResponse struct {
	// List of workflows that match the request's filters.
	Workflows            []*types.Workflow `protobuf:"bytes,1,rep,name=workflows,proto3" json:"workflows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListWorkflowResponse) Reset()         { *m = ListWorkflowResponse{} }
func (m *ListWorkflowResponse) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowResponse) ProtoMessage()    {}
func (*ListWorkflowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2568c098005b7b27, []int{6}
}

func (m *ListWorkflowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowResponse.Unmarshal(m, b)
}
func (m *ListWorkflowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowResponse.Marshal(b, m, deterministic)
}
func (m *ListWorkflowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowResponse.Merge(m, src)
}
func (m *ListWorkflowResponse) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowResponse.Size(m)
}
func (m *ListWorkflowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowResponse proto.InternalMessageInfo

func (m *ListWorkflowResponse) GetWorkflows() []*types.Workflow {
	if m != nil {
		return m.Workflows
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateWorkflowRequest)(nil), "api.CreateWorkflowRequest")
	proto.RegisterType((*CreateWorkflowResponse)(nil), "api.CreateWorkflowResponse")
	proto.RegisterType((*DeleteWorkflowRequest)(nil), "api.DeleteWorkflowRequest")
	proto.RegisterType((*DeleteWorkflowResponse)(nil), "api.DeleteWorkflowResponse")
	proto.RegisterType((*GetWorkflowRequest)(nil), "api.GetWorkflowRequest")
	proto.RegisterType((*ListWorkflowRequest)(nil), "api.ListWorkflowRequest")
	proto.RegisterType((*ListWorkflowResponse)(nil), "api.ListWorkflowResponse")
}

func init() { proto.RegisterFile("protobuf/api/workflow.proto", fileDescriptor_2568c098005b7b27) }

var fileDescriptor_2568c098005b7b27 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x29, 0x05, 0x22, 0xe3, 0x41, 0x33, 0x02, 0xae, 0x25, 0x26, 0xa4, 0x27, 0xe2, 0x9f,
	0x25, 0xe2, 0xd9, 0x13, 0x12, 0x2e, 0xc6, 0x43, 0x2f, 0x9e, 0x4b, 0x3a, 0x40, 0x03, 0x61, 0xd7,
	0xee, 0x12, 0xc2, 0xd1, 0x8b, 0x9f, 0xdb, 0x74, 0xdb, 0x6a, 0x5c, 0xd6, 0xc4, 0xdb, 0x66, 0xe6,
	0xb7, 0xaf, 0xef, 0xbd, 0x2e, 0xf4, 0x65, 0x26, 0xb4, 0x98, 0xef, 0x16, 0xa3, 0x58, 0xa6, 0xa3,
	0xbd, 0xc8, 0xd6, 0x8b, 0x8d, 0xd8, 0x73, 0x33, 0x45, 0x3f, 0x96, 0x69, 0x70, 0xfd, 0x4d, 0xe8,
	0x83, 0x24, 0x65, 0x31, 0xe1, 0x87, 0x07, 0xdd, 0x49, 0x46, 0xb1, 0xa6, 0xb7, 0x72, 0x11, 0xd1,
	0xfb, 0x8e, 0x94, 0xc6, 0x73, 0xf0, 0xd7, 0x74, 0x60, 0xf5, 0x81, 0x37, 0x6c, 0x47, 0xf9, 0x11,
	0x6f, 0xa0, 0xb9, 0x15, 0x09, 0x29, 0xd6, 0x18, 0xf8, 0xc3, 0xd3, 0x71, 0x87, 0x1b, 0x45, 0x5e,
	0x5d, 0xe4, 0xaf, 0x22, 0xa1, 0xa8, 0x40, 0x72, 0x96, 0x92, 0x25, 0x29, 0xd6, 0x74, 0xb3, 0xd3,
	0x64, 0x49, 0x51, 0x81, 0x84, 0x77, 0xd0, 0xb3, 0x2d, 0x28, 0x29, 0xb6, 0x8a, 0x10, 0xa1, 0xb1,
	0x8a, 0xd5, 0x8a, 0x79, 0xc6, 0x84, 0x39, 0x87, 0xb7, 0xd0, 0x7d, 0xa6, 0x0d, 0x1d, 0x1b, 0x76,
	0xc1, 0x0c, 0x7a, 0x36, 0x5c, 0x48, 0x87, 0x43, 0xc0, 0x19, 0xe9, 0xff, 0x68, 0x74, 0xe1, 0xe2,
	0x25, 0x55, 0x36, 0x1a, 0x4e, 0xa1, 0xf3, 0x7b, 0x5c, 0x7a, 0xbe, 0x87, 0x76, 0xd5, 0xb1, 0x62,
	0x9e, 0x49, 0x7f, 0x66, 0xa5, 0x8f, 0x7e, 0x88, 0xf1, 0x67, 0x1d, 0x4e, 0xaa, 0x39, 0x4e, 0xa0,
	0x55, 0x34, 0x81, 0x01, 0x8f, 0x65, 0xca, 0x9d, 0x7f, 0x26, 0xe8, 0x3b, 0x77, 0x65, 0xae, 0x5a,
	0x2e, 0x52, 0x64, 0x2e, 0x45, 0x9c, 0x6d, 0x95, 0x22, 0x7f, 0x94, 0x53, 0xc3, 0x07, 0xf0, 0x67,
	0xa4, 0xf1, 0xd2, 0x50, 0xc7, 0x45, 0x05, 0x76, 0xa4, 0xb0, 0x86, 0x4f, 0xd0, 0xc8, 0x0b, 0x41,
	0x66, 0xee, 0x38, 0x2a, 0x0b, 0xae, 0x1c, 0x9b, 0xea, 0x8b, 0xf3, 0x96, 0x79, 0x90, 0x8f, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x12, 0x57, 0x17, 0x2b, 0xd3, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkflowClient is the client API for Workflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkflowClient interface {
	// Create a Workflow from a Workflow Definition.
	// It will return an unique identifier which is used to interact with the Workflow.
	Create(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*CreateWorkflowResponse, error)
	// Delete a workflow.
	// An error is returned if one or more Instances of the workflow are running.
	Delete(ctx context.Context, in *DeleteWorkflowRequest, opts ...grpc.CallOption) (*DeleteWorkflowResponse, error)
	// Get returns a workflow matching the criteria of the request.
	Get(ctx context.Context, in *GetWorkflowRequest, opts ...grpc.CallOption) (*types.Workflow, error)
	// List returns workflows specified in a request.
	List(ctx context.Context, in *ListWorkflowRequest, opts ...grpc.CallOption) (*ListWorkflowResponse, error)
}

type workflowClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowClient(cc *grpc.ClientConn) WorkflowClient {
	return &workflowClient{cc}
}

func (c *workflowClient) Create(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*CreateWorkflowResponse, error) {
	out := new(CreateWorkflowResponse)
	err := c.cc.Invoke(ctx, "/api.Workflow/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) Delete(ctx context.Context, in *DeleteWorkflowRequest, opts ...grpc.CallOption) (*DeleteWorkflowResponse, error) {
	out := new(DeleteWorkflowResponse)
	err := c.cc.Invoke(ctx, "/api.Workflow/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) Get(ctx context.Context, in *GetWorkflowRequest, opts ...grpc.CallOption) (*types.Workflow, error) {
	out := new(types.Workflow)
	err := c.cc.Invoke(ctx, "/api.Workflow/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) List(ctx context.Context, in *ListWorkflowRequest, opts ...grpc.CallOption) (*ListWorkflowResponse, error) {
	out := new(ListWorkflowResponse)
	err := c.cc.Invoke(ctx, "/api.Workflow/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServer is the server API for Workflow service.
type WorkflowServer interface {
	// Create a Workflow from a Workflow Definition.
	// It will return an unique identifier which is used to interact with the Workflow.
	Create(context.Context, *CreateWorkflowRequest) (*CreateWorkflowResponse, error)
	// Delete a workflow.
	// An error is returned if one or more Instances of the workflow are running.
	Delete(context.Context, *DeleteWorkflowRequest) (*DeleteWorkflowResponse, error)
	// Get returns a workflow matching the criteria of the request.
	Get(context.Context, *GetWorkflowRequest) (*types.Workflow, error)
	// List returns workflows specified in a request.
	List(context.Context, *ListWorkflowRequest) (*ListWorkflowResponse, error)
}

// UnimplementedWorkflowServer can be embedded to have forward compatible implementations.
type UnimplementedWorkflowServer struct {
}

func (*UnimplementedWorkflowServer) Create(ctx context.Context, req *CreateWorkflowRequest) (*CreateWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedWorkflowServer) Delete(ctx context.Context, req *DeleteWorkflowRequest) (*DeleteWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedWorkflowServer) Get(ctx context.Context, req *GetWorkflowRequest) (*types.Workflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedWorkflowServer) List(ctx context.Context, req *ListWorkflowRequest) (*ListWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterWorkflowServer(s *grpc.Server, srv WorkflowServer) {
	s.RegisterService(&_Workflow_serviceDesc, srv)
}

func _Workflow_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Workflow/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).Create(ctx, req.(*CreateWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Workflow/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).Delete(ctx, req.(*DeleteWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Workflow/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).Get(ctx, req.(*GetWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Workflow/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).List(ctx, req.(*ListWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workflow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Workflow",
	HandlerType: (*WorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Workflow_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Workflow_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Workflow_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Workflow_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/workflow.proto",
}
