// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/types/execution.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Status represents the status of a single execution.
// Note that a valid execution must have only one status
// flag at time.
type Status int32

const (
	// Unknown status represents any status unknown to execution.
	Status_Unknown Status = 0
	// Created is an initial status after execution creation.
	Status_Created Status = 1
	// InProgress informs that processing of execution has been started.
	Status_InProgress Status = 2
	// Completed is a success status after execution was processed.
	Status_Completed Status = 3
	// Failed is an error status after execution was processed.
	Status_Failed Status = 4
)

var Status_name = map[int32]string{
	0: "Unknown",
	1: "Created",
	2: "InProgress",
	3: "Completed",
	4: "Failed",
}

var Status_value = map[string]int32{
	"Unknown":    0,
	"Created":    1,
	"InProgress": 2,
	"Completed":  3,
	"Failed":     4,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_42e00237fcad5e7f, []int{0}
}

// Execution represents a single execution run in engine.
type Execution struct {
	// Hash is a unique hash to identify execution.
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// parentHash is the unique hash of parent execution. if execution is triggered by another one, dependency execution considered as the parent.
	ParentHash []byte `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	// eventHash is unique event hash.
	EventHash []byte `protobuf:"bytes,3,opt,name=eventHash,proto3" json:"eventHash,omitempty"`
	// Status is the current status of execution.
	Status Status `protobuf:"varint,4,opt,name=status,proto3,enum=types.Status" json:"status,omitempty"`
	// instanceHash is hash of the instance that can proceed an execution
	InstanceHash []byte `protobuf:"bytes,5,opt,name=instanceHash,proto3" json:"instanceHash,omitempty"`
	// taskKey is the key of the task of this execution.
	TaskKey string `protobuf:"bytes,6,opt,name=taskKey,proto3" json:"taskKey,omitempty"`
	// inputs data of the execution.
	Inputs *types.Struct `protobuf:"bytes,7,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// outputs are the returned data of successful execution.
	Outputs *types.Struct `protobuf:"bytes,8,opt,name=outputs,proto3" json:"outputs,omitempty"`
	// error message of a failed execution.
	Error string `protobuf:"bytes,9,opt,name=error,proto3" json:"error,omitempty"`
	// tags are optionally associated with execution by the user.
	Tags []string `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	// processHash is the unique hash of the process associated to this execution.
	ProcessHash []byte `protobuf:"bytes,11,opt,name=processHash,proto3" json:"processHash,omitempty"`
	// step of the process.
	StepID               string   `protobuf:"bytes,12,opt,name=stepID,proto3" json:"stepID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Execution) Reset()         { *m = Execution{} }
func (m *Execution) String() string { return proto.CompactTextString(m) }
func (*Execution) ProtoMessage()    {}
func (*Execution) Descriptor() ([]byte, []int) {
	return fileDescriptor_42e00237fcad5e7f, []int{0}
}
func (m *Execution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Execution.Unmarshal(m, b)
}
func (m *Execution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Execution.Marshal(b, m, deterministic)
}
func (m *Execution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Execution.Merge(m, src)
}
func (m *Execution) XXX_Size() int {
	return xxx_messageInfo_Execution.Size(m)
}
func (m *Execution) XXX_DiscardUnknown() {
	xxx_messageInfo_Execution.DiscardUnknown(m)
}

var xxx_messageInfo_Execution proto.InternalMessageInfo

func (m *Execution) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Execution) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *Execution) GetEventHash() []byte {
	if m != nil {
		return m.EventHash
	}
	return nil
}

func (m *Execution) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_Unknown
}

func (m *Execution) GetInstanceHash() []byte {
	if m != nil {
		return m.InstanceHash
	}
	return nil
}

func (m *Execution) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *Execution) GetInputs() *types.Struct {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Execution) GetOutputs() *types.Struct {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Execution) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Execution) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Execution) GetProcessHash() []byte {
	if m != nil {
		return m.ProcessHash
	}
	return nil
}

func (m *Execution) GetStepID() string {
	if m != nil {
		return m.StepID
	}
	return ""
}

func init() {
	proto.RegisterEnum("types.Status", Status_name, Status_value)
	proto.RegisterType((*Execution)(nil), "types.Execution")
}

func init() { proto.RegisterFile("protobuf/types/execution.proto", fileDescriptor_42e00237fcad5e7f) }

var fileDescriptor_42e00237fcad5e7f = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x8a, 0xd3, 0x50,
	0x10, 0x36, 0x9b, 0x36, 0x31, 0x93, 0xee, 0x12, 0x06, 0xd1, 0x83, 0x2c, 0x4b, 0x58, 0x10, 0x82,
	0x60, 0xa2, 0xf5, 0x0d, 0x5c, 0x15, 0x17, 0x11, 0x24, 0x8b, 0x37, 0xde, 0x9d, 0xa6, 0xb3, 0x69,
	0x68, 0x7b, 0x4e, 0x38, 0x3f, 0x6a, 0x5f, 0xc1, 0xa7, 0x96, 0x4c, 0x1a, 0xad, 0x37, 0x7b, 0x37,
	0xdf, 0xcf, 0x4c, 0xf2, 0xcd, 0x1c, 0xb8, 0xea, 0x8d, 0x76, 0x7a, 0xe5, 0xef, 0x2b, 0x77, 0xe8,
	0xc9, 0x56, 0xf4, 0x8b, 0x1a, 0xef, 0x3a, 0xad, 0x4a, 0x16, 0x70, 0xce, 0xf4, 0xf3, 0xcb, 0x56,
	0xeb, 0x76, 0x47, 0xd5, 0x5f, 0xb7, 0x75, 0xc6, 0x37, 0x6e, 0x34, 0x5d, 0xff, 0x0e, 0x21, 0xf9,
	0x30, 0x35, 0x22, 0xc2, 0x6c, 0x23, 0xed, 0x46, 0x04, 0x79, 0x50, 0x2c, 0x6a, 0xae, 0xf1, 0x0a,
	0xa0, 0x97, 0x86, 0x94, 0xfb, 0x34, 0x28, 0x67, 0xac, 0x9c, 0x30, 0x78, 0x09, 0x09, 0xfd, 0x98,
	0xe4, 0x90, 0xe5, 0x7f, 0x04, 0xbe, 0x80, 0xc8, 0x3a, 0xe9, 0xbc, 0x15, 0xb3, 0x3c, 0x28, 0x2e,
	0x96, 0xe7, 0x25, 0xff, 0x55, 0x79, 0xc7, 0x64, 0x7d, 0x14, 0xf1, 0x1a, 0x16, 0x9d, 0xb2, 0x4e,
	0xaa, 0x86, 0x78, 0xce, 0x9c, 0xe7, 0xfc, 0xc7, 0xa1, 0x80, 0xd8, 0x49, 0xbb, 0xfd, 0x4c, 0x07,
	0x11, 0xe5, 0x41, 0x91, 0xd4, 0x13, 0xc4, 0x0a, 0xa2, 0x4e, 0xf5, 0xde, 0x59, 0x11, 0xe7, 0x41,
	0x91, 0x2e, 0x9f, 0x95, 0x63, 0xe6, 0x72, 0xca, 0x5c, 0xde, 0x71, 0xe6, 0xfa, 0x68, 0xc3, 0x37,
	0x10, 0x6b, 0xef, 0xb8, 0xe3, 0xf1, 0xc3, 0x1d, 0x93, 0x0f, 0x9f, 0xc0, 0x9c, 0x8c, 0xd1, 0x46,
	0x24, 0xfc, 0xed, 0x11, 0x0c, 0x0b, 0x73, 0xb2, 0xb5, 0x02, 0xf2, 0xb0, 0x48, 0x6a, 0xae, 0x31,
	0x87, 0xb4, 0x37, 0xba, 0x21, 0x6b, 0x39, 0x4a, 0xca, 0x51, 0x4e, 0x29, 0x7c, 0x3a, 0x2c, 0x85,
	0xfa, 0xdb, 0xf7, 0x62, 0xc1, 0xc3, 0x8e, 0xe8, 0xe5, 0x17, 0x88, 0xc6, 0xbd, 0x60, 0x0a, 0xf1,
	0x37, 0xb5, 0x55, 0xfa, 0xa7, 0xca, 0x1e, 0x0d, 0xe0, 0xc6, 0x90, 0x74, 0xb4, 0xce, 0x02, 0xbc,
	0x00, 0xb8, 0x55, 0x5f, 0x8d, 0x6e, 0x0d, 0x59, 0x9b, 0x9d, 0xe1, 0x39, 0x24, 0x37, 0x7a, 0xdf,
	0xef, 0x68, 0x90, 0x43, 0x04, 0x88, 0x3e, 0xca, 0x6e, 0x47, 0xeb, 0x6c, 0xf6, 0x6e, 0xf9, 0xfd,
	0x75, 0xdb, 0xb9, 0x8d, 0x5f, 0x95, 0x8d, 0xde, 0x57, 0x7b, 0xb2, 0xed, 0xab, 0x7b, 0xed, 0xd5,
	0x5a, 0x0e, 0xb7, 0xae, 0x48, 0xb5, 0x9d, 0x3a, 0x79, 0x16, 0x7c, 0x97, 0x55, 0xc4, 0xf8, 0xed,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0x40, 0x29, 0xf3, 0x5d, 0x02, 0x00, 0x00,
}
