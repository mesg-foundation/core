// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/types/service.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Type of trigger available. Event to react to an event. Result to react to a result of an execution.
type Service_Workflow_Trigger_Type int32

const (
	Service_Workflow_Trigger_Unknown Service_Workflow_Trigger_Type = 0
	Service_Workflow_Trigger_Event   Service_Workflow_Trigger_Type = 1
	Service_Workflow_Trigger_Result  Service_Workflow_Trigger_Type = 2
)

var Service_Workflow_Trigger_Type_name = map[int32]string{
	0: "Unknown",
	1: "Event",
	2: "Result",
}

var Service_Workflow_Trigger_Type_value = map[string]int32{
	"Unknown": 0,
	"Event":   1,
	"Result":  2,
}

func (x Service_Workflow_Trigger_Type) String() string {
	return proto.EnumName(Service_Workflow_Trigger_Type_name, int32(x))
}

func (Service_Workflow_Trigger_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 5, 0, 0}
}

// Type of condition available to compare the values.
type Service_Workflow_Trigger_Filter_Predicate int32

const (
	Service_Workflow_Trigger_Filter_Unknown Service_Workflow_Trigger_Filter_Predicate = 0
	Service_Workflow_Trigger_Filter_EQ      Service_Workflow_Trigger_Filter_Predicate = 1
)

var Service_Workflow_Trigger_Filter_Predicate_name = map[int32]string{
	0: "Unknown",
	1: "EQ",
}

var Service_Workflow_Trigger_Filter_Predicate_value = map[string]int32{
	"Unknown": 0,
	"EQ":      1,
}

func (x Service_Workflow_Trigger_Filter_Predicate) String() string {
	return proto.EnumName(Service_Workflow_Trigger_Filter_Predicate_name, int32(x))
}

func (Service_Workflow_Trigger_Filter_Predicate) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 5, 0, 0, 0}
}

// Service represents the service's type.
type Service struct {
	Hash                 string                 `protobuf:"bytes,10,opt,name=hash,proto3" json:"hash,omitempty"`
	Sid                  string                 `protobuf:"bytes,12,opt,name=sid,proto3" json:"sid,omitempty"`
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Configuration        *Service_Configuration `protobuf:"bytes,8,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Tasks                []*Service_Task        `protobuf:"bytes,5,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Events               []*Service_Event       `protobuf:"bytes,6,rep,name=events,proto3" json:"events,omitempty"`
	Dependencies         []*Service_Dependency  `protobuf:"bytes,7,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Repository           string                 `protobuf:"bytes,9,opt,name=repository,proto3" json:"repository,omitempty"`
	Source               string                 `protobuf:"bytes,13,opt,name=source,proto3" json:"source,omitempty"`
	Workflows            []*Service_Workflow    `protobuf:"bytes,14,rep,name=workflows,proto3" json:"workflows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Service) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service) GetConfiguration() *Service_Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *Service) GetTasks() []*Service_Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *Service) GetEvents() []*Service_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Service) GetDependencies() []*Service_Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *Service) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *Service) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Service) GetWorkflows() []*Service_Workflow {
	if m != nil {
		return m.Workflows
	}
	return nil
}

// Events are emitted by the service whenever the service wants.
type Service_Event struct {
	Key                  string               `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Data                 []*Service_Parameter `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Event) Reset()         { *m = Service_Event{} }
func (m *Service_Event) String() string { return proto.CompactTextString(m) }
func (*Service_Event) ProtoMessage()    {}
func (*Service_Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 0}
}

func (m *Service_Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Event.Unmarshal(m, b)
}
func (m *Service_Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Event.Marshal(b, m, deterministic)
}
func (m *Service_Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Event.Merge(m, src)
}
func (m *Service_Event) XXX_Size() int {
	return xxx_messageInfo_Service_Event.Size(m)
}
func (m *Service_Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Event proto.InternalMessageInfo

func (m *Service_Event) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Service_Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service_Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service_Event) GetData() []*Service_Parameter {
	if m != nil {
		return m.Data
	}
	return nil
}

// Task is a function that requires inputs and returns output.
type Service_Task struct {
	Key                  string               `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Inputs               []*Service_Parameter `protobuf:"bytes,6,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              []*Service_Parameter `protobuf:"bytes,7,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Task) Reset()         { *m = Service_Task{} }
func (m *Service_Task) String() string { return proto.CompactTextString(m) }
func (*Service_Task) ProtoMessage()    {}
func (*Service_Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 1}
}

func (m *Service_Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Task.Unmarshal(m, b)
}
func (m *Service_Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Task.Marshal(b, m, deterministic)
}
func (m *Service_Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Task.Merge(m, src)
}
func (m *Service_Task) XXX_Size() int {
	return xxx_messageInfo_Service_Task.Size(m)
}
func (m *Service_Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Task proto.InternalMessageInfo

func (m *Service_Task) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Service_Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service_Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service_Task) GetInputs() []*Service_Parameter {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Service_Task) GetOutputs() []*Service_Parameter {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// Parameter describes the task's inputs, the task's outputs, and the event's data.
type Service_Parameter struct {
	Key                  string               `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string               `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Optional             bool                 `protobuf:"varint,4,opt,name=optional,proto3" json:"optional,omitempty"`
	Repeated             bool                 `protobuf:"varint,9,opt,name=repeated,proto3" json:"repeated,omitempty"`
	Object               []*Service_Parameter `protobuf:"bytes,10,rep,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Parameter) Reset()         { *m = Service_Parameter{} }
func (m *Service_Parameter) String() string { return proto.CompactTextString(m) }
func (*Service_Parameter) ProtoMessage()    {}
func (*Service_Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 2}
}

func (m *Service_Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Parameter.Unmarshal(m, b)
}
func (m *Service_Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Parameter.Marshal(b, m, deterministic)
}
func (m *Service_Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Parameter.Merge(m, src)
}
func (m *Service_Parameter) XXX_Size() int {
	return xxx_messageInfo_Service_Parameter.Size(m)
}
func (m *Service_Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Parameter proto.InternalMessageInfo

func (m *Service_Parameter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Service_Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service_Parameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service_Parameter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Service_Parameter) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

func (m *Service_Parameter) GetRepeated() bool {
	if m != nil {
		return m.Repeated
	}
	return false
}

func (m *Service_Parameter) GetObject() []*Service_Parameter {
	if m != nil {
		return m.Object
	}
	return nil
}

// A configuration is the configuration of the main container of the service's instance.
type Service_Configuration struct {
	Volumes              []string `protobuf:"bytes,1,rep,name=volumes,proto3" json:"volumes,omitempty"`
	VolumesFrom          []string `protobuf:"bytes,2,rep,name=volumesFrom,proto3" json:"volumesFrom,omitempty"`
	Ports                []string `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	Args                 []string `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	Command              string   `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty"`
	Env                  []string `protobuf:"bytes,6,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service_Configuration) Reset()         { *m = Service_Configuration{} }
func (m *Service_Configuration) String() string { return proto.CompactTextString(m) }
func (*Service_Configuration) ProtoMessage()    {}
func (*Service_Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 3}
}

func (m *Service_Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Configuration.Unmarshal(m, b)
}
func (m *Service_Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Configuration.Marshal(b, m, deterministic)
}
func (m *Service_Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Configuration.Merge(m, src)
}
func (m *Service_Configuration) XXX_Size() int {
	return xxx_messageInfo_Service_Configuration.Size(m)
}
func (m *Service_Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Configuration proto.InternalMessageInfo

func (m *Service_Configuration) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Service_Configuration) GetVolumesFrom() []string {
	if m != nil {
		return m.VolumesFrom
	}
	return nil
}

func (m *Service_Configuration) GetPorts() []string {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Service_Configuration) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Service_Configuration) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Service_Configuration) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

// A dependency is a configuration of an other container that runs separately from the service.
type Service_Dependency struct {
	Key                  string   `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty"`
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Volumes              []string `protobuf:"bytes,2,rep,name=volumes,proto3" json:"volumes,omitempty"`
	VolumesFrom          []string `protobuf:"bytes,3,rep,name=volumesFrom,proto3" json:"volumesFrom,omitempty"`
	Ports                []string `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	Args                 []string `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty"`
	Command              string   `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty"`
	Env                  []string `protobuf:"bytes,9,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service_Dependency) Reset()         { *m = Service_Dependency{} }
func (m *Service_Dependency) String() string { return proto.CompactTextString(m) }
func (*Service_Dependency) ProtoMessage()    {}
func (*Service_Dependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 4}
}

func (m *Service_Dependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Dependency.Unmarshal(m, b)
}
func (m *Service_Dependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Dependency.Marshal(b, m, deterministic)
}
func (m *Service_Dependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Dependency.Merge(m, src)
}
func (m *Service_Dependency) XXX_Size() int {
	return xxx_messageInfo_Service_Dependency.Size(m)
}
func (m *Service_Dependency) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Dependency.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Dependency proto.InternalMessageInfo

func (m *Service_Dependency) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Service_Dependency) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Service_Dependency) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Service_Dependency) GetVolumesFrom() []string {
	if m != nil {
		return m.VolumesFrom
	}
	return nil
}

func (m *Service_Dependency) GetPorts() []string {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Service_Dependency) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Service_Dependency) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Service_Dependency) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

// A workflow is a configuration to trigger a specific task when certains conditions of a trigger are valid.
type Service_Workflow struct {
	Key                  string                    `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Trigger              *Service_Workflow_Trigger `protobuf:"bytes,1,opt,name=trigger,proto3" json:"trigger,omitempty"`
	Tasks                []*Service_Workflow_Task  `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Service_Workflow) Reset()         { *m = Service_Workflow{} }
func (m *Service_Workflow) String() string { return proto.CompactTextString(m) }
func (*Service_Workflow) ProtoMessage()    {}
func (*Service_Workflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 5}
}

func (m *Service_Workflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Workflow.Unmarshal(m, b)
}
func (m *Service_Workflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Workflow.Marshal(b, m, deterministic)
}
func (m *Service_Workflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Workflow.Merge(m, src)
}
func (m *Service_Workflow) XXX_Size() int {
	return xxx_messageInfo_Service_Workflow.Size(m)
}
func (m *Service_Workflow) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Workflow.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Workflow proto.InternalMessageInfo

func (m *Service_Workflow) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Service_Workflow) GetTrigger() *Service_Workflow_Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (m *Service_Workflow) GetTasks() []*Service_Workflow_Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

// Trigger is the configuration that will trigger a workflow.
type Service_Workflow_Trigger struct {
	Type                 Service_Workflow_Trigger_Type      `protobuf:"varint,1,opt,name=type,proto3,enum=types.Service_Workflow_Trigger_Type" json:"type,omitempty"`
	InstanceHash         string                             `protobuf:"bytes,2,opt,name=instanceHash,proto3" json:"instanceHash,omitempty"`
	Key                  string                             `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Filters              []*Service_Workflow_Trigger_Filter `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *Service_Workflow_Trigger) Reset()         { *m = Service_Workflow_Trigger{} }
func (m *Service_Workflow_Trigger) String() string { return proto.CompactTextString(m) }
func (*Service_Workflow_Trigger) ProtoMessage()    {}
func (*Service_Workflow_Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 5, 0}
}

func (m *Service_Workflow_Trigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Workflow_Trigger.Unmarshal(m, b)
}
func (m *Service_Workflow_Trigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Workflow_Trigger.Marshal(b, m, deterministic)
}
func (m *Service_Workflow_Trigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Workflow_Trigger.Merge(m, src)
}
func (m *Service_Workflow_Trigger) XXX_Size() int {
	return xxx_messageInfo_Service_Workflow_Trigger.Size(m)
}
func (m *Service_Workflow_Trigger) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Workflow_Trigger.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Workflow_Trigger proto.InternalMessageInfo

func (m *Service_Workflow_Trigger) GetType() Service_Workflow_Trigger_Type {
	if m != nil {
		return m.Type
	}
	return Service_Workflow_Trigger_Unknown
}

func (m *Service_Workflow_Trigger) GetInstanceHash() string {
	if m != nil {
		return m.InstanceHash
	}
	return ""
}

func (m *Service_Workflow_Trigger) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Service_Workflow_Trigger) GetFilters() []*Service_Workflow_Trigger_Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

// Filter is applied on the data of the event/result.
type Service_Workflow_Trigger_Filter struct {
	Key                  string                                    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Predicate            Service_Workflow_Trigger_Filter_Predicate `protobuf:"varint,2,opt,name=predicate,proto3,enum=types.Service_Workflow_Trigger_Filter_Predicate" json:"predicate,omitempty"`
	Value                string                                    `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *Service_Workflow_Trigger_Filter) Reset()         { *m = Service_Workflow_Trigger_Filter{} }
func (m *Service_Workflow_Trigger_Filter) String() string { return proto.CompactTextString(m) }
func (*Service_Workflow_Trigger_Filter) ProtoMessage()    {}
func (*Service_Workflow_Trigger_Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 5, 0, 0}
}

func (m *Service_Workflow_Trigger_Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Workflow_Trigger_Filter.Unmarshal(m, b)
}
func (m *Service_Workflow_Trigger_Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Workflow_Trigger_Filter.Marshal(b, m, deterministic)
}
func (m *Service_Workflow_Trigger_Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Workflow_Trigger_Filter.Merge(m, src)
}
func (m *Service_Workflow_Trigger_Filter) XXX_Size() int {
	return xxx_messageInfo_Service_Workflow_Trigger_Filter.Size(m)
}
func (m *Service_Workflow_Trigger_Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Workflow_Trigger_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Workflow_Trigger_Filter proto.InternalMessageInfo

func (m *Service_Workflow_Trigger_Filter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Service_Workflow_Trigger_Filter) GetPredicate() Service_Workflow_Trigger_Filter_Predicate {
	if m != nil {
		return m.Predicate
	}
	return Service_Workflow_Trigger_Filter_Unknown
}

func (m *Service_Workflow_Trigger_Filter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Definition of the task to execute when the workflow is triggered.
type Service_Workflow_Task struct {
	InstanceHash         string   `protobuf:"bytes,1,opt,name=instanceHash,proto3" json:"instanceHash,omitempty"`
	TaskKey              string   `protobuf:"bytes,2,opt,name=taskKey,proto3" json:"taskKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service_Workflow_Task) Reset()         { *m = Service_Workflow_Task{} }
func (m *Service_Workflow_Task) String() string { return proto.CompactTextString(m) }
func (*Service_Workflow_Task) ProtoMessage()    {}
func (*Service_Workflow_Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 5, 1}
}

func (m *Service_Workflow_Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Workflow_Task.Unmarshal(m, b)
}
func (m *Service_Workflow_Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Workflow_Task.Marshal(b, m, deterministic)
}
func (m *Service_Workflow_Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Workflow_Task.Merge(m, src)
}
func (m *Service_Workflow_Task) XXX_Size() int {
	return xxx_messageInfo_Service_Workflow_Task.Size(m)
}
func (m *Service_Workflow_Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Workflow_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Workflow_Task proto.InternalMessageInfo

func (m *Service_Workflow_Task) GetInstanceHash() string {
	if m != nil {
		return m.InstanceHash
	}
	return ""
}

func (m *Service_Workflow_Task) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func init() {
	proto.RegisterEnum("types.Service_Workflow_Trigger_Type", Service_Workflow_Trigger_Type_name, Service_Workflow_Trigger_Type_value)
	proto.RegisterEnum("types.Service_Workflow_Trigger_Filter_Predicate", Service_Workflow_Trigger_Filter_Predicate_name, Service_Workflow_Trigger_Filter_Predicate_value)
	proto.RegisterType((*Service)(nil), "types.Service")
	proto.RegisterType((*Service_Event)(nil), "types.Service.Event")
	proto.RegisterType((*Service_Task)(nil), "types.Service.Task")
	proto.RegisterType((*Service_Parameter)(nil), "types.Service.Parameter")
	proto.RegisterType((*Service_Configuration)(nil), "types.Service.Configuration")
	proto.RegisterType((*Service_Dependency)(nil), "types.Service.Dependency")
	proto.RegisterType((*Service_Workflow)(nil), "types.Service.Workflow")
	proto.RegisterType((*Service_Workflow_Trigger)(nil), "types.Service.Workflow.Trigger")
	proto.RegisterType((*Service_Workflow_Trigger_Filter)(nil), "types.Service.Workflow.Trigger.Filter")
	proto.RegisterType((*Service_Workflow_Task)(nil), "types.Service.Workflow.Task")
}

func init() { proto.RegisterFile("protobuf/types/service.proto", fileDescriptor_578ec16019661f8f) }

var fileDescriptor_578ec16019661f8f = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x8a, 0xec, 0x44,
	0x10, 0x36, 0x33, 0x99, 0x64, 0x52, 0xb3, 0xbb, 0x0e, 0xed, 0xa2, 0x6d, 0x58, 0x74, 0x58, 0x44,
	0x56, 0x39, 0xce, 0x2c, 0x11, 0x41, 0x2f, 0x04, 0x39, 0x9e, 0x73, 0x10, 0x04, 0x39, 0xb6, 0x2b,
	0x82, 0x77, 0xbd, 0x49, 0x4d, 0x36, 0xce, 0xa4, 0x3b, 0x74, 0x77, 0x66, 0x99, 0x0b, 0x5f, 0x40,
	0x7c, 0x09, 0x7d, 0x06, 0x9f, 0xc1, 0x7b, 0xdf, 0xc0, 0x47, 0x91, 0xee, 0x24, 0xf3, 0xc7, 0xfe,
	0xf8, 0x77, 0x57, 0x3f, 0x5f, 0xa5, 0xea, 0xab, 0x9f, 0x0e, 0x9c, 0x55, 0x4a, 0x1a, 0x79, 0x5d,
	0xcf, 0x67, 0x66, 0x5d, 0xa1, 0x9e, 0x69, 0x54, 0xab, 0x22, 0xc5, 0xa9, 0x33, 0x93, 0x81, 0x33,
	0x9e, 0xff, 0xf4, 0x2a, 0x84, 0xdf, 0x34, 0x0e, 0x42, 0xc0, 0xbf, 0xe1, 0xfa, 0x86, 0xc2, 0xc4,
	0xbb, 0x88, 0x98, 0x93, 0xc9, 0x18, 0xfa, 0xba, 0xc8, 0xe8, 0x91, 0x33, 0x59, 0xd1, 0xa2, 0x04,
	0x2f, 0x91, 0x7a, 0x0d, 0xca, 0xca, 0x64, 0x02, 0xa3, 0x0c, 0x75, 0xaa, 0x8a, 0xca, 0x14, 0x52,
	0xd0, 0x9e, 0x73, 0xed, 0x9a, 0xc8, 0x53, 0x38, 0x4e, 0xa5, 0x98, 0x17, 0x79, 0xad, 0xb8, 0xc3,
	0x0c, 0x27, 0xde, 0xc5, 0x28, 0x39, 0x9b, 0xba, 0x32, 0xa6, 0x6d, 0x09, 0xd3, 0xcf, 0x77, 0x31,
	0x6c, 0x3f, 0x84, 0xbc, 0x07, 0x03, 0xc3, 0xf5, 0x42, 0xd3, 0xc1, 0xa4, 0x7f, 0x31, 0x4a, 0x5e,
	0x3b, 0x88, 0xbd, 0xe2, 0x7a, 0xc1, 0x1a, 0x04, 0x79, 0x02, 0x01, 0xae, 0x50, 0x18, 0x4d, 0x03,
	0x87, 0x3d, 0x3d, 0xc0, 0x3e, 0xb7, 0x4e, 0xd6, 0x62, 0xc8, 0xa7, 0x70, 0x94, 0x61, 0x85, 0x22,
	0x43, 0x91, 0x16, 0xa8, 0x69, 0xe8, 0x62, 0xde, 0x3c, 0x88, 0x79, 0xd6, 0x41, 0xd6, 0x6c, 0x0f,
	0x4e, 0xde, 0x02, 0x50, 0x58, 0x49, 0x5d, 0x18, 0xa9, 0xd6, 0x34, 0x72, 0xe4, 0x77, 0x2c, 0xe4,
	0x75, 0x08, 0xb4, 0xac, 0x55, 0x8a, 0xf4, 0xd8, 0xf9, 0x5a, 0x8d, 0x7c, 0x04, 0xd1, 0xad, 0x54,
	0x8b, 0xf9, 0x52, 0xde, 0x6a, 0x7a, 0xe2, 0x72, 0xbe, 0x71, 0x90, 0xf3, 0xbb, 0xd6, 0xcf, 0xb6,
	0xc8, 0xf8, 0x47, 0x18, 0xb8, 0xf2, 0xed, 0x6c, 0x16, 0xb8, 0xa6, 0x7e, 0x33, 0x9b, 0x05, 0xae,
	0xff, 0xe5, 0x6c, 0x9e, 0x80, 0x9f, 0x71, 0xc3, 0x69, 0xdf, 0x95, 0x40, 0x0f, 0x4a, 0x78, 0xc9,
	0x15, 0x2f, 0xd1, 0xa0, 0x62, 0x0e, 0x15, 0xff, 0xe6, 0x81, 0x6f, 0x5b, 0xdd, 0xa5, 0x1f, 0xfe,
	0xd7, 0xf4, 0x97, 0x10, 0x14, 0xa2, 0xaa, 0x37, 0xb3, 0xba, 0xbf, 0x80, 0x16, 0x47, 0x12, 0x08,
	0x65, 0x6d, 0x5c, 0x48, 0xf8, 0x48, 0x48, 0x07, 0x8c, 0xff, 0xf0, 0x20, 0xda, 0x98, 0xff, 0xb7,
	0xda, 0x09, 0xf8, 0x36, 0x33, 0xed, 0x37, 0x51, 0x56, 0x26, 0x31, 0x0c, 0xa5, 0xf3, 0xf2, 0xa5,
	0x9b, 0xcd, 0x90, 0x6d, 0x74, 0xeb, 0x53, 0x58, 0x21, 0x37, 0x98, 0xb9, 0x45, 0x19, 0xb2, 0x8d,
	0x6e, 0xfb, 0x20, 0xaf, 0x7f, 0xc0, 0xd4, 0x50, 0x78, 0xac, 0x0f, 0x0d, 0x2e, 0xfe, 0xc5, 0x83,
	0xe3, 0xbd, 0x8b, 0x21, 0x14, 0xc2, 0x95, 0x5c, 0xd6, 0x25, 0x6a, 0xea, 0x4d, 0xfa, 0x17, 0x11,
	0xeb, 0x54, 0xcb, 0xa5, 0x15, 0x5f, 0x28, 0x59, 0xd2, 0x9e, 0xf3, 0xee, 0x9a, 0xc8, 0x29, 0x0c,
	0x2a, 0xa9, 0x8c, 0x76, 0x7b, 0x10, 0xb1, 0x46, 0xb1, 0x0c, 0xb9, 0xca, 0x35, 0xf5, 0x9d, 0xd1,
	0xc9, 0x36, 0x4b, 0x2a, 0xcb, 0x92, 0x8b, 0x8c, 0x0e, 0x1c, 0xf1, 0x4e, 0xb5, 0x7d, 0x45, 0xb1,
	0x72, 0x83, 0x8c, 0x98, 0x15, 0xe3, 0xdf, 0x3d, 0x80, 0xed, 0xe5, 0xdc, 0xd1, 0xf8, 0x53, 0x18,
	0x14, 0x25, 0xcf, 0xbb, 0xce, 0x37, 0xca, 0x2e, 0x91, 0xde, 0x83, 0x44, 0xfa, 0x0f, 0x10, 0xf1,
	0xef, 0x22, 0x12, 0xfc, 0x13, 0x22, 0xd1, 0x96, 0xc8, 0x9f, 0x3e, 0x0c, 0xbb, 0x73, 0xec, 0x68,
	0xf4, 0xb7, 0x34, 0x3e, 0x81, 0xd0, 0xa8, 0x22, 0xcf, 0x51, 0x39, 0x22, 0xa3, 0xe4, 0xed, 0x7b,
	0x4e, 0x79, 0x7a, 0xd5, 0xc0, 0x58, 0x87, 0x27, 0x49, 0xf7, 0xae, 0xf5, 0xdc, 0xdc, 0xcf, 0xee,
	0x0d, 0xdc, 0x3e, 0x70, 0xf1, 0xcf, 0x7d, 0x08, 0xdb, 0x0f, 0x91, 0x8f, 0xdb, 0x25, 0xb4, 0x79,
	0x4f, 0x92, 0x77, 0x1e, 0xc9, 0x3b, 0xbd, 0x5a, 0x57, 0xd8, 0xae, 0xea, 0x39, 0x1c, 0x15, 0x42,
	0x1b, 0x2e, 0x52, 0xfc, 0xc2, 0xbe, 0xfc, 0xcd, 0x86, 0xef, 0xd9, 0xee, 0xa0, 0xfa, 0x19, 0x84,
	0xf3, 0x62, 0x69, 0x50, 0x35, 0x1d, 0x1e, 0x25, 0xef, 0x3e, 0x96, 0xf2, 0x85, 0x83, 0xb3, 0x2e,
	0x2c, 0xfe, 0xd5, 0x83, 0xa0, 0xb1, 0x75, 0x9f, 0xf7, 0xb6, 0x9f, 0xff, 0x0a, 0xa2, 0x4a, 0x61,
	0x56, 0xa4, 0xdc, 0xa0, 0xab, 0xe8, 0x24, 0xb9, 0xfc, 0x7b, 0x09, 0xa6, 0x2f, 0xbb, 0x38, 0xb6,
	0xfd, 0x84, 0x5d, 0x87, 0x15, 0x5f, 0xd6, 0xdd, 0x91, 0x36, 0xca, 0xf9, 0x04, 0xa2, 0x0d, 0x9a,
	0x8c, 0x20, 0xfc, 0x56, 0x2c, 0x84, 0xbc, 0x15, 0xe3, 0x57, 0x48, 0x00, 0xbd, 0xe7, 0x5f, 0x8f,
	0xbd, 0xf3, 0xf7, 0xc1, 0xb7, 0xad, 0xda, 0x77, 0x46, 0xed, 0xe3, 0x3b, 0xf6, 0x08, 0x40, 0xc0,
	0x50, 0xd7, 0x4b, 0x33, 0xee, 0xc5, 0xcf, 0xda, 0x37, 0xf1, 0xb0, 0xa1, 0xde, 0x1d, 0x0d, 0xa5,
	0x10, 0xda, 0x19, 0x7e, 0x89, 0xeb, 0xb6, 0xdf, 0x9d, 0xfa, 0x34, 0xf9, 0xfe, 0x32, 0x2f, 0xcc,
	0x4d, 0x7d, 0x3d, 0x4d, 0x65, 0x39, 0x2b, 0x51, 0xe7, 0x1f, 0xcc, 0x65, 0x2d, 0x32, 0x77, 0xdb,
	0x33, 0x14, 0x79, 0x21, 0x70, 0xb6, 0xff, 0x57, 0xbf, 0x0e, 0x9c, 0xfe, 0xe1, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x26, 0xcb, 0x5a, 0x3f, 0xee, 0x07, 0x00, 0x00,
}
