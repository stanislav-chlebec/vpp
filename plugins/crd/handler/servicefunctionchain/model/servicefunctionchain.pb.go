// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: servicefunctionchain.proto

package model

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type ServiceFunctionChain_ServiceFunction_Type int32

const (
	ServiceFunctionChain_ServiceFunction_Pod               ServiceFunctionChain_ServiceFunction_Type = 0
	ServiceFunctionChain_ServiceFunction_ExternalInterface ServiceFunctionChain_ServiceFunction_Type = 1
)

var ServiceFunctionChain_ServiceFunction_Type_name = map[int32]string{
	0: "Pod",
	1: "ExternalInterface",
}

var ServiceFunctionChain_ServiceFunction_Type_value = map[string]int32{
	"Pod":               0,
	"ExternalInterface": 1,
}

func (x ServiceFunctionChain_ServiceFunction_Type) String() string {
	return proto.EnumName(ServiceFunctionChain_ServiceFunction_Type_name, int32(x))
}

func (ServiceFunctionChain_ServiceFunction_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a1073e00293c62d6, []int{0, 0, 0}
}

// ServiceFunctionChain is used to store definition of a service function chain as a k8s CRD resource.
type ServiceFunctionChain struct {
	// Name of the chain.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Name of the custom pod network where the chain resides
	// (if applicable, can be left blank for the default pod network).
	Network string `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	// List of service functions (chain elements) in the chain.
	Chain                []*ServiceFunctionChain_ServiceFunction `protobuf:"bytes,3,rep,name=chain,proto3" json:"chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ServiceFunctionChain) Reset()         { *m = ServiceFunctionChain{} }
func (m *ServiceFunctionChain) String() string { return proto.CompactTextString(m) }
func (*ServiceFunctionChain) ProtoMessage()    {}
func (*ServiceFunctionChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1073e00293c62d6, []int{0}
}
func (m *ServiceFunctionChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceFunctionChain.Unmarshal(m, b)
}
func (m *ServiceFunctionChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceFunctionChain.Marshal(b, m, deterministic)
}
func (m *ServiceFunctionChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceFunctionChain.Merge(m, src)
}
func (m *ServiceFunctionChain) XXX_Size() int {
	return xxx_messageInfo_ServiceFunctionChain.Size(m)
}
func (m *ServiceFunctionChain) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceFunctionChain.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceFunctionChain proto.InternalMessageInfo

func (m *ServiceFunctionChain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceFunctionChain) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *ServiceFunctionChain) GetChain() []*ServiceFunctionChain_ServiceFunction {
	if m != nil {
		return m.Chain
	}
	return nil
}

type ServiceFunctionChain_ServiceFunction struct {
	// Name of the service function (optional).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the service function.
	Type ServiceFunctionChain_ServiceFunction_Type `protobuf:"varint,2,opt,name=type,proto3,enum=model.ServiceFunctionChain_ServiceFunction_Type" json:"type,omitempty"`
	// Pod selector (k8s labels) identifying the pod(s)
	// (applicable for pod service function type).
	PodSelector map[string]string `protobuf:"bytes,3,rep,name=pod_selector,json=podSelector,proto3" json:"pod_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Interface trough which the traffic enters the service function. Applicable for:
	// - pods with multiple interfaces
	// - external interface if it is the last element (service function) in the chain
	InputInterface string `protobuf:"bytes,4,opt,name=input_interface,json=inputInterface,proto3" json:"input_interface,omitempty"`
	// Interface trough which the traffic leaves the service function. Applicable for:
	// - pods with multiple interfaces
	// - external interface if it is the first element (service function) in the chain
	OutputInterface      string   `protobuf:"bytes,5,opt,name=output_interface,json=outputInterface,proto3" json:"output_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceFunctionChain_ServiceFunction) Reset()         { *m = ServiceFunctionChain_ServiceFunction{} }
func (m *ServiceFunctionChain_ServiceFunction) String() string { return proto.CompactTextString(m) }
func (*ServiceFunctionChain_ServiceFunction) ProtoMessage()    {}
func (*ServiceFunctionChain_ServiceFunction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1073e00293c62d6, []int{0, 0}
}
func (m *ServiceFunctionChain_ServiceFunction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceFunctionChain_ServiceFunction.Unmarshal(m, b)
}
func (m *ServiceFunctionChain_ServiceFunction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceFunctionChain_ServiceFunction.Marshal(b, m, deterministic)
}
func (m *ServiceFunctionChain_ServiceFunction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceFunctionChain_ServiceFunction.Merge(m, src)
}
func (m *ServiceFunctionChain_ServiceFunction) XXX_Size() int {
	return xxx_messageInfo_ServiceFunctionChain_ServiceFunction.Size(m)
}
func (m *ServiceFunctionChain_ServiceFunction) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceFunctionChain_ServiceFunction.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceFunctionChain_ServiceFunction proto.InternalMessageInfo

func (m *ServiceFunctionChain_ServiceFunction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceFunctionChain_ServiceFunction) GetType() ServiceFunctionChain_ServiceFunction_Type {
	if m != nil {
		return m.Type
	}
	return ServiceFunctionChain_ServiceFunction_Pod
}

func (m *ServiceFunctionChain_ServiceFunction) GetPodSelector() map[string]string {
	if m != nil {
		return m.PodSelector
	}
	return nil
}

func (m *ServiceFunctionChain_ServiceFunction) GetInputInterface() string {
	if m != nil {
		return m.InputInterface
	}
	return ""
}

func (m *ServiceFunctionChain_ServiceFunction) GetOutputInterface() string {
	if m != nil {
		return m.OutputInterface
	}
	return ""
}

func init() {
	proto.RegisterEnum("model.ServiceFunctionChain_ServiceFunction_Type", ServiceFunctionChain_ServiceFunction_Type_name, ServiceFunctionChain_ServiceFunction_Type_value)
	proto.RegisterType((*ServiceFunctionChain)(nil), "model.ServiceFunctionChain")
	proto.RegisterType((*ServiceFunctionChain_ServiceFunction)(nil), "model.ServiceFunctionChain.ServiceFunction")
	proto.RegisterMapType((map[string]string)(nil), "model.ServiceFunctionChain.ServiceFunction.PodSelectorEntry")
}

func init() { proto.RegisterFile("servicefunctionchain.proto", fileDescriptor_a1073e00293c62d6) }

var fileDescriptor_a1073e00293c62d6 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xbf, 0x34, 0xc9, 0x57, 0x9c, 0x4a, 0x13, 0x87, 0x0a, 0x21, 0xa7, 0xd2, 0x83, 0x56,
	0x84, 0x20, 0xf5, 0x22, 0x22, 0x82, 0x68, 0x05, 0x6f, 0x25, 0xf5, 0x1e, 0x62, 0x32, 0xc5, 0xd0,
	0x74, 0x77, 0xd9, 0x6e, 0xaa, 0xf9, 0x07, 0x9e, 0xfd, 0xc5, 0x92, 0x4d, 0x5a, 0x31, 0xf4, 0xd2,
	0xdb, 0xcc, 0xbb, 0xcf, 0xbc, 0x33, 0x3b, 0xbb, 0xe0, 0xaf, 0x49, 0x6e, 0xb2, 0x84, 0x16, 0x05,
	0x4b, 0x54, 0xc6, 0x59, 0xf2, 0x1e, 0x67, 0x2c, 0x10, 0x92, 0x2b, 0x8e, 0xf6, 0x8a, 0xa7, 0x94,
	0x8f, 0xbe, 0x2d, 0x18, 0xcc, 0x6b, 0xea, 0xb9, 0xa1, 0x1e, 0x2b, 0x0a, 0x11, 0x2c, 0x16, 0xaf,
	0xc8, 0x33, 0x86, 0xc6, 0xf8, 0x28, 0xd4, 0x31, 0x7a, 0xd0, 0x65, 0xa4, 0x3e, 0xb8, 0x5c, 0x7a,
	0x1d, 0x2d, 0x6f, 0x53, 0x7c, 0x00, 0x5b, 0x9b, 0x7b, 0xe6, 0xd0, 0x1c, 0xf7, 0x26, 0x97, 0x81,
	0x76, 0x0f, 0xf6, 0x39, 0xb7, 0xc5, 0xb0, 0xae, 0xf4, 0xbf, 0x4c, 0x70, 0x5a, 0x47, 0x7b, 0x87,
	0x78, 0x02, 0x4b, 0x95, 0x82, 0xf4, 0x04, 0xfd, 0xc9, 0xd5, 0x01, 0x9d, 0x82, 0xd7, 0x52, 0x50,
	0xa8, 0xab, 0x31, 0x82, 0x63, 0xc1, 0xd3, 0x68, 0x4d, 0x39, 0x25, 0x8a, 0xcb, 0x66, 0xee, 0xbb,
	0x43, 0xdc, 0x66, 0x3c, 0x9d, 0x37, 0xe5, 0x53, 0xa6, 0x64, 0x19, 0xf6, 0xc4, 0xaf, 0x82, 0xe7,
	0xe0, 0x64, 0x4c, 0x14, 0x2a, 0xca, 0x98, 0x22, 0xb9, 0x88, 0x13, 0xf2, 0x2c, 0x7d, 0x8b, 0xbe,
	0x96, 0x5f, 0xb6, 0x2a, 0x5e, 0x80, 0xcb, 0x0b, 0xf5, 0x97, 0xb4, 0x35, 0xe9, 0xd4, 0xfa, 0x0e,
	0xf5, 0xef, 0xc1, 0x6d, 0x37, 0x45, 0x17, 0xcc, 0x25, 0x95, 0xcd, 0x86, 0xaa, 0x10, 0x07, 0x60,
	0x6f, 0xe2, 0xbc, 0xa0, 0xe6, 0x8d, 0xea, 0xe4, 0xb6, 0x73, 0x63, 0x8c, 0xce, 0xc0, 0xaa, 0x56,
	0x80, 0x5d, 0x30, 0x67, 0x3c, 0x75, 0xff, 0xe1, 0x29, 0x9c, 0x4c, 0x3f, 0x15, 0x49, 0x16, 0xe7,
	0xbb, 0x2e, 0xae, 0xf1, 0xf6, 0x5f, 0x7f, 0x91, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0,
	0xba, 0x6b, 0x38, 0x40, 0x02, 0x00, 0x00,
}
