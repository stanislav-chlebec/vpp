// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bd.proto

package l2

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

type BridgeDomain struct {
	Name                 string                              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Flood                bool                                `protobuf:"varint,2,opt,name=flood,proto3" json:"flood,omitempty"`
	UnknownUnicastFlood  bool                                `protobuf:"varint,3,opt,name=unknown_unicast_flood,json=unknownUnicastFlood,proto3" json:"unknown_unicast_flood,omitempty"`
	Forward              bool                                `protobuf:"varint,4,opt,name=forward,proto3" json:"forward,omitempty"`
	Learn                bool                                `protobuf:"varint,5,opt,name=learn,proto3" json:"learn,omitempty"`
	ArpTermination       bool                                `protobuf:"varint,6,opt,name=arp_termination,json=arpTermination,proto3" json:"arp_termination,omitempty"`
	MacAge               uint32                              `protobuf:"varint,7,opt,name=mac_age,json=macAge,proto3" json:"mac_age,omitempty"`
	Interfaces           []*BridgeDomain_Interface           `protobuf:"bytes,100,rep,name=interfaces" json:"interfaces,omitempty"`
	ArpTerminationTable  []*BridgeDomain_ArpTerminationEntry `protobuf:"bytes,102,rep,name=arp_termination_table,json=arpTerminationTable" json:"arp_termination_table,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *BridgeDomain) Reset()         { *m = BridgeDomain{} }
func (m *BridgeDomain) String() string { return proto.CompactTextString(m) }
func (*BridgeDomain) ProtoMessage()    {}
func (*BridgeDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd_406d6f6aaa9ef47c, []int{0}
}
func (m *BridgeDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BridgeDomain.Unmarshal(m, b)
}
func (m *BridgeDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BridgeDomain.Marshal(b, m, deterministic)
}
func (dst *BridgeDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeDomain.Merge(dst, src)
}
func (m *BridgeDomain) XXX_Size() int {
	return xxx_messageInfo_BridgeDomain.Size(m)
}
func (m *BridgeDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeDomain.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeDomain proto.InternalMessageInfo

func (m *BridgeDomain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BridgeDomain) GetFlood() bool {
	if m != nil {
		return m.Flood
	}
	return false
}

func (m *BridgeDomain) GetUnknownUnicastFlood() bool {
	if m != nil {
		return m.UnknownUnicastFlood
	}
	return false
}

func (m *BridgeDomain) GetForward() bool {
	if m != nil {
		return m.Forward
	}
	return false
}

func (m *BridgeDomain) GetLearn() bool {
	if m != nil {
		return m.Learn
	}
	return false
}

func (m *BridgeDomain) GetArpTermination() bool {
	if m != nil {
		return m.ArpTermination
	}
	return false
}

func (m *BridgeDomain) GetMacAge() uint32 {
	if m != nil {
		return m.MacAge
	}
	return 0
}

func (m *BridgeDomain) GetInterfaces() []*BridgeDomain_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *BridgeDomain) GetArpTerminationTable() []*BridgeDomain_ArpTerminationEntry {
	if m != nil {
		return m.ArpTerminationTable
	}
	return nil
}

type BridgeDomain_Interface struct {
	Name                    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BridgedVirtualInterface bool     `protobuf:"varint,2,opt,name=bridged_virtual_interface,json=bridgedVirtualInterface,proto3" json:"bridged_virtual_interface,omitempty"`
	SplitHorizonGroup       uint32   `protobuf:"varint,3,opt,name=split_horizon_group,json=splitHorizonGroup,proto3" json:"split_horizon_group,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *BridgeDomain_Interface) Reset()         { *m = BridgeDomain_Interface{} }
func (m *BridgeDomain_Interface) String() string { return proto.CompactTextString(m) }
func (*BridgeDomain_Interface) ProtoMessage()    {}
func (*BridgeDomain_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd_406d6f6aaa9ef47c, []int{0, 0}
}
func (m *BridgeDomain_Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BridgeDomain_Interface.Unmarshal(m, b)
}
func (m *BridgeDomain_Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BridgeDomain_Interface.Marshal(b, m, deterministic)
}
func (dst *BridgeDomain_Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeDomain_Interface.Merge(dst, src)
}
func (m *BridgeDomain_Interface) XXX_Size() int {
	return xxx_messageInfo_BridgeDomain_Interface.Size(m)
}
func (m *BridgeDomain_Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeDomain_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeDomain_Interface proto.InternalMessageInfo

func (m *BridgeDomain_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BridgeDomain_Interface) GetBridgedVirtualInterface() bool {
	if m != nil {
		return m.BridgedVirtualInterface
	}
	return false
}

func (m *BridgeDomain_Interface) GetSplitHorizonGroup() uint32 {
	if m != nil {
		return m.SplitHorizonGroup
	}
	return 0
}

type BridgeDomain_ArpTerminationEntry struct {
	IpAddress            string   `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	PhysAddress          string   `protobuf:"bytes,2,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BridgeDomain_ArpTerminationEntry) Reset()         { *m = BridgeDomain_ArpTerminationEntry{} }
func (m *BridgeDomain_ArpTerminationEntry) String() string { return proto.CompactTextString(m) }
func (*BridgeDomain_ArpTerminationEntry) ProtoMessage()    {}
func (*BridgeDomain_ArpTerminationEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd_406d6f6aaa9ef47c, []int{0, 1}
}
func (m *BridgeDomain_ArpTerminationEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BridgeDomain_ArpTerminationEntry.Unmarshal(m, b)
}
func (m *BridgeDomain_ArpTerminationEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BridgeDomain_ArpTerminationEntry.Marshal(b, m, deterministic)
}
func (dst *BridgeDomain_ArpTerminationEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeDomain_ArpTerminationEntry.Merge(dst, src)
}
func (m *BridgeDomain_ArpTerminationEntry) XXX_Size() int {
	return xxx_messageInfo_BridgeDomain_ArpTerminationEntry.Size(m)
}
func (m *BridgeDomain_ArpTerminationEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeDomain_ArpTerminationEntry.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeDomain_ArpTerminationEntry proto.InternalMessageInfo

func (m *BridgeDomain_ArpTerminationEntry) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *BridgeDomain_ArpTerminationEntry) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*BridgeDomain)(nil), "l2.BridgeDomain")
	proto.RegisterType((*BridgeDomain_Interface)(nil), "l2.BridgeDomain.Interface")
	proto.RegisterType((*BridgeDomain_ArpTerminationEntry)(nil), "l2.BridgeDomain.ArpTerminationEntry")
}

func init() { proto.RegisterFile("bd.proto", fileDescriptor_bd_406d6f6aaa9ef47c) }

var fileDescriptor_bd_406d6f6aaa9ef47c = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcb, 0xae, 0xd3, 0x30,
	0x18, 0x84, 0x95, 0x9e, 0x9e, 0xf6, 0xf4, 0x6f, 0x0b, 0xc2, 0xa1, 0xaa, 0xa9, 0x84, 0x14, 0x10,
	0x12, 0x59, 0x65, 0x11, 0x76, 0xdd, 0x15, 0x71, 0xdd, 0x46, 0xe5, 0xb2, 0xb3, 0x9c, 0xd8, 0x49,
	0x2d, 0x12, 0xdb, 0x72, 0x1c, 0xaa, 0xf2, 0x0a, 0x3c, 0x22, 0x2f, 0x83, 0xe2, 0xa4, 0x37, 0xe8,
	0x2e, 0xff, 0x3f, 0x5f, 0x66, 0x26, 0xb1, 0xe1, 0x21, 0x65, 0x91, 0x36, 0xca, 0x2a, 0x34, 0x28,
	0xe3, 0x97, 0x7f, 0x86, 0x30, 0x7b, 0x6b, 0x04, 0x2b, 0xf8, 0x3b, 0x55, 0x51, 0x21, 0x11, 0x82,
	0xa1, 0xa4, 0x15, 0xc7, 0x5e, 0xe0, 0x85, 0x93, 0xc4, 0x3d, 0xa3, 0xa7, 0x70, 0x9f, 0x97, 0x4a,
	0x31, 0x3c, 0x08, 0xbc, 0xf0, 0x21, 0xe9, 0x06, 0x14, 0xc3, 0xa2, 0x91, 0x3f, 0xa4, 0xda, 0x4b,
	0xd2, 0x48, 0x91, 0xd1, 0xda, 0x92, 0x8e, 0xba, 0x73, 0x94, 0xdf, 0x8b, 0x5f, 0x3a, 0xed, 0x83,
	0x7b, 0x07, 0xc3, 0x38, 0x57, 0x66, 0x4f, 0x0d, 0xc3, 0x43, 0x47, 0x1d, 0xc7, 0x36, 0xa3, 0xe4,
	0xd4, 0x48, 0x7c, 0xdf, 0x65, 0xb8, 0x01, 0xbd, 0x86, 0xc7, 0xd4, 0x68, 0x62, 0xb9, 0xa9, 0x84,
	0xa4, 0x56, 0x28, 0x89, 0x47, 0x4e, 0x7f, 0x44, 0x8d, 0xde, 0x9e, 0xb7, 0x68, 0x09, 0xe3, 0x8a,
	0x66, 0x84, 0x16, 0x1c, 0x8f, 0x03, 0x2f, 0x9c, 0x27, 0xa3, 0x8a, 0x66, 0x9b, 0x82, 0xa3, 0x35,
	0x80, 0x90, 0x96, 0x9b, 0x9c, 0x66, 0xbc, 0xc6, 0x2c, 0xb8, 0x0b, 0xa7, 0xf1, 0x2a, 0x2a, 0xe3,
	0xe8, 0xf2, 0xab, 0xa3, 0xcf, 0x47, 0x24, 0xb9, 0xa0, 0xd1, 0x77, 0x58, 0xfc, 0x93, 0x4e, 0x2c,
	0x4d, 0x4b, 0x8e, 0x73, 0x67, 0xf3, 0xea, 0x3f, 0x9b, 0xcd, 0x55, 0xa9, 0xf7, 0xd2, 0x9a, 0x43,
	0xe2, 0x5f, 0x37, 0xdd, 0xb6, 0x06, 0xab, 0xdf, 0x1e, 0x4c, 0x4e, 0x99, 0x37, 0xff, 0xf9, 0x1a,
	0x9e, 0xa5, 0xce, 0x9a, 0x91, 0x9f, 0xc2, 0xd8, 0x86, 0x96, 0xe4, 0xd4, 0xac, 0x3f, 0x87, 0x65,
	0x0f, 0x7c, 0xed, 0xf4, 0xb3, 0x5f, 0x04, 0x7e, 0xad, 0x4b, 0x61, 0xc9, 0x4e, 0x19, 0xf1, 0x4b,
	0x49, 0x52, 0x18, 0xd5, 0x68, 0x77, 0x2e, 0xf3, 0xe4, 0x89, 0x93, 0x3e, 0x75, 0xca, 0xc7, 0x56,
	0x58, 0x7d, 0x03, 0xff, 0x46, 0x73, 0xf4, 0x1c, 0x40, 0x68, 0x42, 0x19, 0x33, 0xbc, 0xae, 0xfb,
	0x72, 0x13, 0xa1, 0x37, 0xdd, 0x02, 0xbd, 0x80, 0x99, 0xde, 0x1d, 0xea, 0x13, 0x30, 0x70, 0xc0,
	0xb4, 0xdd, 0xf5, 0x48, 0x3a, 0x72, 0x17, 0xed, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0xe2, 0x08, 0x4c, 0x74, 0x02, 0x00, 0x00,
}
