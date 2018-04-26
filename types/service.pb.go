// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProtoService struct {
	Name         string                      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description  string                      `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Visibility   string                      `protobuf:"bytes,3,opt,name=visibility" json:"visibility,omitempty"`
	Publish      string                      `protobuf:"bytes,4,opt,name=publish" json:"publish,omitempty"`
	Tasks        map[string]*ProtoTask       `protobuf:"bytes,5,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Events       map[string]*ProtoEvent      `protobuf:"bytes,6,rep,name=events" json:"events,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Dependencies map[string]*ProtoDependency `protobuf:"bytes,7,rep,name=dependencies" json:"dependencies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ProtoService) Reset()                    { *m = ProtoService{} }
func (m *ProtoService) String() string            { return proto.CompactTextString(m) }
func (*ProtoService) ProtoMessage()               {}
func (*ProtoService) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ProtoService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtoService) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProtoService) GetVisibility() string {
	if m != nil {
		return m.Visibility
	}
	return ""
}

func (m *ProtoService) GetPublish() string {
	if m != nil {
		return m.Publish
	}
	return ""
}

func (m *ProtoService) GetTasks() map[string]*ProtoTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *ProtoService) GetEvents() map[string]*ProtoEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ProtoService) GetDependencies() map[string]*ProtoDependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

type ProtoTask struct {
	Name        string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                     `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Verifiable  bool                       `protobuf:"varint,3,opt,name=verifiable" json:"verifiable,omitempty"`
	Payable     bool                       `protobuf:"varint,4,opt,name=payable" json:"payable,omitempty"`
	Fees        *ProtoFee                  `protobuf:"bytes,5,opt,name=fees" json:"fees,omitempty"`
	Inputs      map[string]*ProtoParameter `protobuf:"bytes,6,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Outputs     map[string]*ProtoOutput    `protobuf:"bytes,7,rep,name=outputs" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ProtoTask) Reset()                    { *m = ProtoTask{} }
func (m *ProtoTask) String() string            { return proto.CompactTextString(m) }
func (*ProtoTask) ProtoMessage()               {}
func (*ProtoTask) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ProtoTask) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtoTask) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProtoTask) GetVerifiable() bool {
	if m != nil {
		return m.Verifiable
	}
	return false
}

func (m *ProtoTask) GetPayable() bool {
	if m != nil {
		return m.Payable
	}
	return false
}

func (m *ProtoTask) GetFees() *ProtoFee {
	if m != nil {
		return m.Fees
	}
	return nil
}

func (m *ProtoTask) GetInputs() map[string]*ProtoParameter {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *ProtoTask) GetOutputs() map[string]*ProtoOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type ProtoFee struct {
	Developer string `protobuf:"bytes,1,opt,name=developer" json:"developer,omitempty"`
	Validator string `protobuf:"bytes,2,opt,name=validator" json:"validator,omitempty"`
	Executor  string `protobuf:"bytes,3,opt,name=executor" json:"executor,omitempty"`
	Emittors  string `protobuf:"bytes,4,opt,name=emittors" json:"emittors,omitempty"`
}

func (m *ProtoFee) Reset()                    { *m = ProtoFee{} }
func (m *ProtoFee) String() string            { return proto.CompactTextString(m) }
func (*ProtoFee) ProtoMessage()               {}
func (*ProtoFee) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *ProtoFee) GetDeveloper() string {
	if m != nil {
		return m.Developer
	}
	return ""
}

func (m *ProtoFee) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *ProtoFee) GetExecutor() string {
	if m != nil {
		return m.Executor
	}
	return ""
}

func (m *ProtoFee) GetEmittors() string {
	if m != nil {
		return m.Emittors
	}
	return ""
}

type ProtoEvent struct {
	Name        string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                     `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Data        map[string]*ProtoParameter `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ProtoEvent) Reset()                    { *m = ProtoEvent{} }
func (m *ProtoEvent) String() string            { return proto.CompactTextString(m) }
func (*ProtoEvent) ProtoMessage()               {}
func (*ProtoEvent) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *ProtoEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtoEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProtoEvent) GetData() map[string]*ProtoParameter {
	if m != nil {
		return m.Data
	}
	return nil
}

type ProtoOutput struct {
	Name        string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                     `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Data        map[string]*ProtoParameter `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ProtoOutput) Reset()                    { *m = ProtoOutput{} }
func (m *ProtoOutput) String() string            { return proto.CompactTextString(m) }
func (*ProtoOutput) ProtoMessage()               {}
func (*ProtoOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *ProtoOutput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtoOutput) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProtoOutput) GetData() map[string]*ProtoParameter {
	if m != nil {
		return m.Data
	}
	return nil
}

type ProtoParameter struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Optional    bool   `protobuf:"varint,4,opt,name=optional" json:"optional,omitempty"`
}

func (m *ProtoParameter) Reset()                    { *m = ProtoParameter{} }
func (m *ProtoParameter) String() string            { return proto.CompactTextString(m) }
func (*ProtoParameter) ProtoMessage()               {}
func (*ProtoParameter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *ProtoParameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtoParameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProtoParameter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ProtoParameter) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

type ProtoDependency struct {
	Key     string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Image   string   `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	Volumes []string `protobuf:"bytes,3,rep,name=volumes" json:"volumes,omitempty"`
	Ports   []string `protobuf:"bytes,4,rep,name=ports" json:"ports,omitempty"`
	Command string   `protobuf:"bytes,5,opt,name=command" json:"command,omitempty"`
}

func (m *ProtoDependency) Reset()                    { *m = ProtoDependency{} }
func (m *ProtoDependency) String() string            { return proto.CompactTextString(m) }
func (*ProtoDependency) ProtoMessage()               {}
func (*ProtoDependency) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *ProtoDependency) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ProtoDependency) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ProtoDependency) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *ProtoDependency) GetPorts() []string {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ProtoDependency) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func init() {
	proto.RegisterType((*ProtoService)(nil), "types.ProtoService")
	proto.RegisterType((*ProtoTask)(nil), "types.ProtoTask")
	proto.RegisterType((*ProtoFee)(nil), "types.ProtoFee")
	proto.RegisterType((*ProtoEvent)(nil), "types.ProtoEvent")
	proto.RegisterType((*ProtoOutput)(nil), "types.ProtoOutput")
	proto.RegisterType((*ProtoParameter)(nil), "types.ProtoParameter")
	proto.RegisterType((*ProtoDependency)(nil), "types.ProtoDependency")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0xd4, 0x4c,
	0x10, 0x94, 0xb3, 0xde, 0x4d, 0xdc, 0xce, 0xf7, 0x25, 0x19, 0x01, 0xb2, 0x96, 0x10, 0xa2, 0x45,
	0x40, 0x24, 0x90, 0x41, 0x21, 0x52, 0x10, 0xe7, 0x04, 0x29, 0x08, 0x85, 0xc8, 0x20, 0x71, 0x9e,
	0x5d, 0x77, 0x60, 0x14, 0xff, 0xc9, 0x1e, 0x5b, 0xf8, 0xc2, 0x85, 0xc7, 0xe1, 0xc4, 0x33, 0x20,
	0xde, 0x0b, 0x4d, 0x8f, 0xed, 0x1d, 0x67, 0x7d, 0x0a, 0x70, 0x73, 0x77, 0x75, 0xd5, 0xf4, 0x74,
	0x97, 0x6d, 0xf8, 0xaf, 0xc0, 0xbc, 0x12, 0x0b, 0xf4, 0xb3, 0x3c, 0x95, 0x29, 0x1b, 0xcb, 0x3a,
	0xc3, 0x62, 0xf6, 0xdd, 0x86, 0xcd, 0x0b, 0x95, 0x78, 0xaf, 0x51, 0xc6, 0xc0, 0x4e, 0x78, 0x8c,
	0x9e, 0xb5, 0x6f, 0x1d, 0x38, 0x01, 0x3d, 0xb3, 0x7d, 0x70, 0x43, 0x2c, 0x16, 0xb9, 0xc8, 0xa4,
	0x48, 0x13, 0x6f, 0x8d, 0x20, 0x33, 0xc5, 0xf6, 0x00, 0x2a, 0x51, 0x88, 0xb9, 0x88, 0x84, 0xac,
	0xbd, 0x11, 0x15, 0x18, 0x19, 0xe6, 0xc1, 0x7a, 0x56, 0xce, 0x23, 0x51, 0x7c, 0xf6, 0x6c, 0x02,
	0xdb, 0x90, 0x1d, 0xc1, 0x58, 0xf2, 0xe2, 0xaa, 0xf0, 0xc6, 0xfb, 0xa3, 0x03, 0xf7, 0x70, 0xcf,
	0xa7, 0xbe, 0x7c, 0xb3, 0x27, 0xff, 0x83, 0x2a, 0x38, 0x4d, 0x64, 0x5e, 0x07, 0xba, 0x98, 0x1d,
	0xc3, 0x04, 0x2b, 0x4c, 0x64, 0xe1, 0x4d, 0x88, 0x76, 0x7f, 0x88, 0x76, 0x4a, 0x15, 0x9a, 0xd7,
	0x94, 0xb3, 0x33, 0xd8, 0x0c, 0x31, 0xc3, 0x24, 0xc4, 0x64, 0x21, 0xb0, 0xf0, 0xd6, 0x89, 0xfe,
	0x70, 0x88, 0x7e, 0x62, 0xd4, 0x69, 0x91, 0x1e, 0x75, 0xfa, 0x06, 0x60, 0xd9, 0x18, 0xdb, 0x86,
	0xd1, 0x15, 0xd6, 0xcd, 0xd8, 0xd4, 0x23, 0x7b, 0x04, 0xe3, 0x8a, 0x47, 0x25, 0xd2, 0xbc, 0xdc,
	0xc3, 0x6d, 0xf3, 0x0c, 0x45, 0x0c, 0x34, 0xfc, 0x6a, 0xed, 0xa5, 0x35, 0x7d, 0x0b, 0xae, 0xd1,
	0xed, 0x80, 0xd8, 0xe3, 0xbe, 0xd8, 0x8e, 0x29, 0x46, 0x4c, 0x53, 0xed, 0x23, 0xec, 0xac, 0x34,
	0x3f, 0xa0, 0xf9, 0xb4, 0xaf, 0x79, 0xc7, 0xd4, 0xec, 0xf8, 0xb5, 0x21, 0x3c, 0xfb, 0x31, 0x02,
	0xa7, 0xeb, 0xff, 0x0f, 0xac, 0x82, 0xb9, 0xb8, 0x14, 0x7c, 0x1e, 0x21, 0x59, 0x65, 0x23, 0x30,
	0x32, 0x64, 0x15, 0x5e, 0x13, 0x68, 0x13, 0xd8, 0x86, 0xec, 0x01, 0xd8, 0x97, 0x88, 0xca, 0x29,
	0xaa, 0xdd, 0x2d, 0xb3, 0xdd, 0xd7, 0x88, 0x01, 0x81, 0xec, 0x08, 0x26, 0x22, 0xc9, 0xca, 0xce,
	0x19, 0xbb, 0xd7, 0xc7, 0xee, 0x9f, 0x11, 0xdc, 0xd8, 0x42, 0xd7, 0xb2, 0x63, 0x58, 0x4f, 0x4b,
	0x49, 0x34, 0xed, 0x88, 0x7b, 0x2b, 0xb4, 0x77, 0x1a, 0xd7, 0xbc, 0xb6, 0x7a, 0x7a, 0x01, 0xae,
	0xa1, 0x37, 0x30, 0xe4, 0x27, 0xfd, 0x21, 0xdf, 0x36, 0x75, 0x2f, 0x78, 0xce, 0x63, 0x94, 0x98,
	0x9b, 0xcb, 0x3b, 0x87, 0x4d, 0xf3, 0xa8, 0x01, 0xc9, 0x83, 0xbe, 0x24, 0x33, 0x25, 0x35, 0xd5,
	0xdc, 0xd9, 0x57, 0xd8, 0x68, 0x47, 0xc4, 0x76, 0xc1, 0x09, 0xb1, 0xc2, 0x28, 0xcd, 0x30, 0x6f,
	0x14, 0x97, 0x09, 0x85, 0x56, 0x3c, 0x12, 0x21, 0x97, 0x69, 0xde, 0x6c, 0x6e, 0x99, 0x60, 0x53,
	0xd8, 0xc0, 0x2f, 0xb8, 0x28, 0x15, 0xa8, 0x5f, 0xf0, 0x2e, 0x26, 0x2c, 0x16, 0x52, 0xa6, 0x79,
	0xd1, 0xbc, 0xdf, 0x5d, 0x3c, 0xfb, 0x69, 0x01, 0x2c, 0x6d, 0x7a, 0x43, 0xd3, 0x3c, 0x03, 0x3b,
	0xe4, 0x92, 0x7b, 0x23, 0x5a, 0xce, 0xdd, 0x15, 0xf7, 0xfb, 0x27, 0x5c, 0x72, 0xbd, 0x1a, 0x2a,
	0x9c, 0x9e, 0x83, 0xd3, 0xa5, 0xfe, 0xc2, 0x56, 0x66, 0xbf, 0x2c, 0x70, 0x8d, 0x01, 0xdf, 0xf0,
	0x1a, 0xcf, 0x7b, 0xd7, 0xd8, 0x5d, 0x5d, 0xdc, 0x3f, 0xbf, 0x47, 0x05, 0xff, 0xf7, 0xc1, 0x1b,
	0xde, 0x84, 0x81, 0xad, 0x8e, 0x6a, 0x9c, 0x40, 0xcf, 0xca, 0x05, 0x29, 0xa1, 0x3c, 0x6a, 0x5e,
	0xdd, 0x2e, 0x9e, 0x7d, 0xb3, 0x60, 0xeb, 0xda, 0x87, 0x65, 0xe0, 0x3a, 0xb7, 0x60, 0x2c, 0x62,
	0xfe, 0x09, 0x9b, 0x13, 0x75, 0xa0, 0xbe, 0x08, 0x55, 0x1a, 0x95, 0x31, 0x16, 0x34, 0x38, 0x27,
	0x68, 0x43, 0x55, 0x9f, 0xa5, 0xb9, 0x54, 0xa6, 0x53, 0x79, 0x1d, 0xa8, 0xfa, 0x45, 0x1a, 0xc7,
	0x3c, 0x09, 0xe9, 0x53, 0xe1, 0x04, 0x6d, 0x38, 0x9f, 0xd0, 0xbf, 0xef, 0xc5, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xfd, 0x0a, 0xe4, 0x5d, 0x0c, 0x07, 0x00, 0x00,
}