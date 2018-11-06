// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nodeconfig.proto

package model

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// NodeConfig is used to store Contiv-specific node configuration entered via CRD.
type NodeConfig struct {
	// name of the node to which the configuration applies
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// main VPP interface used for the inter-node connectivity
	MainVppInterface *NodeConfig_InterfaceConfig `protobuf:"bytes,2,opt,name=main_vpp_interface,json=mainVppInterface" json:"main_vpp_interface,omitempty"`
	// other interfaces on VPP, not necessarily used for inter-node connectivity
	OtherVppInterfaces []*NodeConfig_InterfaceConfig `protobuf:"bytes,3,rep,name=other_vpp_interfaces,json=otherVppInterfaces" json:"other_vpp_interfaces,omitempty"`
	// interface to be stolen from the host stack and bound to VPP
	StealInterface string `protobuf:"bytes,4,opt,name=steal_interface,json=stealInterface,proto3" json:"steal_interface,omitempty"`
	// IP address of the default gateway
	Gateway string `protobuf:"bytes,5,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// whether to NAT external traffic or not
	NatExternalTraffic   bool     `protobuf:"varint,6,opt,name=nat_external_traffic,json=natExternalTraffic,proto3" json:"nat_external_traffic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeConfig) Reset()         { *m = NodeConfig{} }
func (m *NodeConfig) String() string { return proto.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()    {}
func (*NodeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeconfig_36cb1a12573cb08f, []int{0}
}
func (m *NodeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeConfig.Unmarshal(m, b)
}
func (m *NodeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeConfig.Marshal(b, m, deterministic)
}
func (dst *NodeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeConfig.Merge(dst, src)
}
func (m *NodeConfig) XXX_Size() int {
	return xxx_messageInfo_NodeConfig.Size(m)
}
func (m *NodeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NodeConfig proto.InternalMessageInfo

func (m *NodeConfig) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NodeConfig) GetMainVppInterface() *NodeConfig_InterfaceConfig {
	if m != nil {
		return m.MainVppInterface
	}
	return nil
}

func (m *NodeConfig) GetOtherVppInterfaces() []*NodeConfig_InterfaceConfig {
	if m != nil {
		return m.OtherVppInterfaces
	}
	return nil
}

func (m *NodeConfig) GetStealInterface() string {
	if m != nil {
		return m.StealInterface
	}
	return ""
}

func (m *NodeConfig) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *NodeConfig) GetNatExternalTraffic() bool {
	if m != nil {
		return m.NatExternalTraffic
	}
	return false
}

// InterfaceConfig stores configuration for a single interface.
type NodeConfig_InterfaceConfig struct {
	// interface name to which the configuration applies
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// ip address to statically assign to the interface
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	// if enabled, the interface will be assigned IP address dynamically via DHCP protocol
	UseDhcp              bool     `protobuf:"varint,3,opt,name=use_dhcp,json=useDhcp,proto3" json:"use_dhcp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeConfig_InterfaceConfig) Reset()         { *m = NodeConfig_InterfaceConfig{} }
func (m *NodeConfig_InterfaceConfig) String() string { return proto.CompactTextString(m) }
func (*NodeConfig_InterfaceConfig) ProtoMessage()    {}
func (*NodeConfig_InterfaceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeconfig_36cb1a12573cb08f, []int{0, 0}
}
func (m *NodeConfig_InterfaceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeConfig_InterfaceConfig.Unmarshal(m, b)
}
func (m *NodeConfig_InterfaceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeConfig_InterfaceConfig.Marshal(b, m, deterministic)
}
func (dst *NodeConfig_InterfaceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeConfig_InterfaceConfig.Merge(dst, src)
}
func (m *NodeConfig_InterfaceConfig) XXX_Size() int {
	return xxx_messageInfo_NodeConfig_InterfaceConfig.Size(m)
}
func (m *NodeConfig_InterfaceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeConfig_InterfaceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NodeConfig_InterfaceConfig proto.InternalMessageInfo

func (m *NodeConfig_InterfaceConfig) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *NodeConfig_InterfaceConfig) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *NodeConfig_InterfaceConfig) GetUseDhcp() bool {
	if m != nil {
		return m.UseDhcp
	}
	return false
}

func init() {
	proto.RegisterType((*NodeConfig)(nil), "model.NodeConfig")
	proto.RegisterType((*NodeConfig_InterfaceConfig)(nil), "model.NodeConfig.InterfaceConfig")
}

func init() { proto.RegisterFile("nodeconfig.proto", fileDescriptor_nodeconfig_36cb1a12573cb08f) }

var fileDescriptor_nodeconfig_36cb1a12573cb08f = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x54, 0x12, 0xda, 0x26, 0x8b, 0x48, 0x2b, 0xab, 0x07, 0x03, 0x97, 0x80, 0x84, 0xc8, 0x29,
	0x42, 0xf0, 0x09, 0xc0, 0x81, 0x4b, 0x91, 0x02, 0xe2, 0x6a, 0x19, 0x67, 0xd3, 0x58, 0x4a, 0x6c,
	0x2b, 0x71, 0x79, 0x7c, 0x12, 0x7f, 0x89, 0xe2, 0xd2, 0x34, 0x70, 0xe2, 0xb8, 0x33, 0xb3, 0xb3,
	0xe3, 0x31, 0x2c, 0x94, 0x2e, 0x50, 0x68, 0x55, 0xca, 0x75, 0x66, 0x5a, 0x6d, 0x35, 0x99, 0x34,
	0xba, 0xc0, 0xfa, 0xfc, 0x2b, 0x00, 0x58, 0xe9, 0x02, 0x6f, 0x1d, 0x47, 0x4e, 0x21, 0xea, 0x95,
	0x4c, 0xf1, 0x06, 0xa9, 0x97, 0x78, 0x69, 0x94, 0x87, 0x3d, 0xb0, 0xe2, 0x0d, 0x92, 0x47, 0x20,
	0x0d, 0x97, 0x8a, 0xbd, 0x19, 0xc3, 0xa4, 0xb2, 0xd8, 0x96, 0x5c, 0x20, 0xf5, 0x13, 0x2f, 0x3d,
	0xbc, 0x3e, 0xcb, 0x9c, 0x5f, 0xb6, 0xf7, 0xca, 0x1e, 0x76, 0x92, 0xed, 0x9c, 0x2f, 0xfa, 0xe5,
	0x17, 0x63, 0x06, 0x9c, 0x3c, 0xc1, 0x52, 0xdb, 0x0a, 0xdb, 0xdf, 0x8e, 0x1d, 0x0d, 0x92, 0xe0,
	0x7f, 0x96, 0xc4, 0xad, 0x8f, 0x3d, 0x3b, 0x72, 0x09, 0xf3, 0xce, 0x22, 0xaf, 0x47, 0x11, 0x0f,
	0xdc, 0x43, 0x62, 0x07, 0xef, 0xaf, 0x53, 0x98, 0xad, 0xb9, 0xc5, 0x77, 0xfe, 0x49, 0x27, 0x4e,
	0xb0, 0x1b, 0xc9, 0x15, 0x2c, 0x15, 0xb7, 0x0c, 0x3f, 0x2c, 0xb6, 0x8a, 0xd7, 0xcc, 0xb6, 0xbc,
	0x2c, 0xa5, 0xa0, 0xd3, 0xc4, 0x4b, 0xc3, 0x9c, 0x28, 0x6e, 0xef, 0x7f, 0xa8, 0xe7, 0x2d, 0x73,
	0x22, 0x60, 0xfe, 0x27, 0x1b, 0xb9, 0x80, 0x78, 0x48, 0x30, 0xee, 0xf3, 0x68, 0x40, 0x5d, 0xa9,
	0x31, 0xf8, 0xd2, 0xb8, 0x12, 0xa3, 0xdc, 0x97, 0x86, 0x1c, 0x43, 0xb8, 0xe9, 0x90, 0x15, 0x95,
	0x30, 0x34, 0x70, 0xf7, 0x66, 0x9b, 0x0e, 0xef, 0x2a, 0x61, 0x5e, 0xa7, 0xee, 0xe7, 0x6e, 0xbe,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x56, 0xb8, 0x45, 0xcd, 0x01, 0x00, 0x00,
}
