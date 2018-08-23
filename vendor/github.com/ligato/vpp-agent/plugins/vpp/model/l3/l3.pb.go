// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: l3.proto

/*
Package l3 is a generated protocol buffer package.

It is generated from these files:
	l3.proto

It has these top-level messages:
	StaticRoutes
	ArpTable
	ProxyArpRanges
	ProxyArpInterfaces
	STNTable
	IPScanNeighbor
*/
package l3

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

type StaticRoutes_Route_RouteType int32

const (
	StaticRoutes_Route_INTRA_VRF StaticRoutes_Route_RouteType = 0
	StaticRoutes_Route_INTER_VRF StaticRoutes_Route_RouteType = 1
	StaticRoutes_Route_DROP      StaticRoutes_Route_RouteType = 2
)

var StaticRoutes_Route_RouteType_name = map[int32]string{
	0: "INTRA_VRF",
	1: "INTER_VRF",
	2: "DROP",
}
var StaticRoutes_Route_RouteType_value = map[string]int32{
	"INTRA_VRF": 0,
	"INTER_VRF": 1,
	"DROP":      2,
}

func (x StaticRoutes_Route_RouteType) String() string {
	return proto.EnumName(StaticRoutes_Route_RouteType_name, int32(x))
}
func (StaticRoutes_Route_RouteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorL3, []int{0, 0, 0}
}

type IPScanNeighbor_Mode int32

const (
	IPScanNeighbor_DISABLED IPScanNeighbor_Mode = 0
	IPScanNeighbor_IPv4     IPScanNeighbor_Mode = 1
	IPScanNeighbor_IPv6     IPScanNeighbor_Mode = 2
	IPScanNeighbor_BOTH     IPScanNeighbor_Mode = 3
)

var IPScanNeighbor_Mode_name = map[int32]string{
	0: "DISABLED",
	1: "IPv4",
	2: "IPv6",
	3: "BOTH",
}
var IPScanNeighbor_Mode_value = map[string]int32{
	"DISABLED": 0,
	"IPv4":     1,
	"IPv6":     2,
	"BOTH":     3,
}

func (x IPScanNeighbor_Mode) String() string {
	return proto.EnumName(IPScanNeighbor_Mode_name, int32(x))
}
func (IPScanNeighbor_Mode) EnumDescriptor() ([]byte, []int) { return fileDescriptorL3, []int{5, 0} }

// Static routes
type StaticRoutes struct {
	Routes []*StaticRoutes_Route `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
}

func (m *StaticRoutes) Reset()                    { *m = StaticRoutes{} }
func (m *StaticRoutes) String() string            { return proto.CompactTextString(m) }
func (*StaticRoutes) ProtoMessage()               {}
func (*StaticRoutes) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{0} }

func (m *StaticRoutes) GetRoutes() []*StaticRoutes_Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type StaticRoutes_Route struct {
	Type              StaticRoutes_Route_RouteType `protobuf:"varint,10,opt,name=type,proto3,enum=l3.StaticRoutes_Route_RouteType" json:"type,omitempty"`
	VrfId             uint32                       `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	Description       string                       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DstIpAddr         string                       `protobuf:"bytes,3,opt,name=dst_ip_addr,json=dstIpAddr,proto3" json:"dst_ip_addr,omitempty"`
	NextHopAddr       string                       `protobuf:"bytes,4,opt,name=next_hop_addr,json=nextHopAddr,proto3" json:"next_hop_addr,omitempty"`
	OutgoingInterface string                       `protobuf:"bytes,5,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	Weight            uint32                       `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	Preference        uint32                       `protobuf:"varint,7,opt,name=preference,proto3" json:"preference,omitempty"`
	// (a poor man's primary and backup)
	ViaVrfId uint32 `protobuf:"varint,8,opt,name=via_vrf_id,json=viaVrfId,proto3" json:"via_vrf_id,omitempty"`
}

func (m *StaticRoutes_Route) Reset()                    { *m = StaticRoutes_Route{} }
func (m *StaticRoutes_Route) String() string            { return proto.CompactTextString(m) }
func (*StaticRoutes_Route) ProtoMessage()               {}
func (*StaticRoutes_Route) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{0, 0} }

func (m *StaticRoutes_Route) GetType() StaticRoutes_Route_RouteType {
	if m != nil {
		return m.Type
	}
	return StaticRoutes_Route_INTRA_VRF
}

func (m *StaticRoutes_Route) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *StaticRoutes_Route) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StaticRoutes_Route) GetDstIpAddr() string {
	if m != nil {
		return m.DstIpAddr
	}
	return ""
}

func (m *StaticRoutes_Route) GetNextHopAddr() string {
	if m != nil {
		return m.NextHopAddr
	}
	return ""
}

func (m *StaticRoutes_Route) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *StaticRoutes_Route) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *StaticRoutes_Route) GetPreference() uint32 {
	if m != nil {
		return m.Preference
	}
	return 0
}

func (m *StaticRoutes_Route) GetViaVrfId() uint32 {
	if m != nil {
		return m.ViaVrfId
	}
	return 0
}

// IP ARP entries
type ArpTable struct {
	ArpEntries []*ArpTable_ArpEntry `protobuf:"bytes,1,rep,name=arp_entries,json=arpEntries" json:"arp_entries,omitempty"`
}

func (m *ArpTable) Reset()                    { *m = ArpTable{} }
func (m *ArpTable) String() string            { return proto.CompactTextString(m) }
func (*ArpTable) ProtoMessage()               {}
func (*ArpTable) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{1} }

func (m *ArpTable) GetArpEntries() []*ArpTable_ArpEntry {
	if m != nil {
		return m.ArpEntries
	}
	return nil
}

type ArpTable_ArpEntry struct {
	Interface   string `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	IpAddress   string `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	PhysAddress string `protobuf:"bytes,3,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	Static      bool   `protobuf:"varint,4,opt,name=static,proto3" json:"static,omitempty"`
}

func (m *ArpTable_ArpEntry) Reset()                    { *m = ArpTable_ArpEntry{} }
func (m *ArpTable_ArpEntry) String() string            { return proto.CompactTextString(m) }
func (*ArpTable_ArpEntry) ProtoMessage()               {}
func (*ArpTable_ArpEntry) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{1, 0} }

func (m *ArpTable_ArpEntry) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *ArpTable_ArpEntry) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *ArpTable_ArpEntry) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *ArpTable_ArpEntry) GetStatic() bool {
	if m != nil {
		return m.Static
	}
	return false
}

// Proxy ARP ranges
type ProxyArpRanges struct {
	RangeLists []*ProxyArpRanges_RangeList `protobuf:"bytes,1,rep,name=range_lists,json=rangeLists" json:"range_lists,omitempty"`
}

func (m *ProxyArpRanges) Reset()                    { *m = ProxyArpRanges{} }
func (m *ProxyArpRanges) String() string            { return proto.CompactTextString(m) }
func (*ProxyArpRanges) ProtoMessage()               {}
func (*ProxyArpRanges) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{2} }

func (m *ProxyArpRanges) GetRangeLists() []*ProxyArpRanges_RangeList {
	if m != nil {
		return m.RangeLists
	}
	return nil
}

type ProxyArpRanges_RangeList struct {
	Label  string                            `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Ranges []*ProxyArpRanges_RangeList_Range `protobuf:"bytes,2,rep,name=ranges" json:"ranges,omitempty"`
}

func (m *ProxyArpRanges_RangeList) Reset()                    { *m = ProxyArpRanges_RangeList{} }
func (m *ProxyArpRanges_RangeList) String() string            { return proto.CompactTextString(m) }
func (*ProxyArpRanges_RangeList) ProtoMessage()               {}
func (*ProxyArpRanges_RangeList) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{2, 0} }

func (m *ProxyArpRanges_RangeList) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ProxyArpRanges_RangeList) GetRanges() []*ProxyArpRanges_RangeList_Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type ProxyArpRanges_RangeList_Range struct {
	FirstIp string `protobuf:"bytes,1,opt,name=first_ip,json=firstIp,proto3" json:"first_ip,omitempty"`
	LastIp  string `protobuf:"bytes,2,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
}

func (m *ProxyArpRanges_RangeList_Range) Reset()         { *m = ProxyArpRanges_RangeList_Range{} }
func (m *ProxyArpRanges_RangeList_Range) String() string { return proto.CompactTextString(m) }
func (*ProxyArpRanges_RangeList_Range) ProtoMessage()    {}
func (*ProxyArpRanges_RangeList_Range) Descriptor() ([]byte, []int) {
	return fileDescriptorL3, []int{2, 0, 0}
}

func (m *ProxyArpRanges_RangeList_Range) GetFirstIp() string {
	if m != nil {
		return m.FirstIp
	}
	return ""
}

func (m *ProxyArpRanges_RangeList_Range) GetLastIp() string {
	if m != nil {
		return m.LastIp
	}
	return ""
}

// Proxy ARP interfaces
type ProxyArpInterfaces struct {
	InterfaceLists []*ProxyArpInterfaces_InterfaceList `protobuf:"bytes,1,rep,name=interface_lists,json=interfaceLists" json:"interface_lists,omitempty"`
}

func (m *ProxyArpInterfaces) Reset()                    { *m = ProxyArpInterfaces{} }
func (m *ProxyArpInterfaces) String() string            { return proto.CompactTextString(m) }
func (*ProxyArpInterfaces) ProtoMessage()               {}
func (*ProxyArpInterfaces) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{3} }

func (m *ProxyArpInterfaces) GetInterfaceLists() []*ProxyArpInterfaces_InterfaceList {
	if m != nil {
		return m.InterfaceLists
	}
	return nil
}

type ProxyArpInterfaces_InterfaceList struct {
	Label      string                                        `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Interfaces []*ProxyArpInterfaces_InterfaceList_Interface `protobuf:"bytes,2,rep,name=interfaces" json:"interfaces,omitempty"`
}

func (m *ProxyArpInterfaces_InterfaceList) Reset()         { *m = ProxyArpInterfaces_InterfaceList{} }
func (m *ProxyArpInterfaces_InterfaceList) String() string { return proto.CompactTextString(m) }
func (*ProxyArpInterfaces_InterfaceList) ProtoMessage()    {}
func (*ProxyArpInterfaces_InterfaceList) Descriptor() ([]byte, []int) {
	return fileDescriptorL3, []int{3, 0}
}

func (m *ProxyArpInterfaces_InterfaceList) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ProxyArpInterfaces_InterfaceList) GetInterfaces() []*ProxyArpInterfaces_InterfaceList_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type ProxyArpInterfaces_InterfaceList_Interface struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ProxyArpInterfaces_InterfaceList_Interface) Reset() {
	*m = ProxyArpInterfaces_InterfaceList_Interface{}
}
func (m *ProxyArpInterfaces_InterfaceList_Interface) String() string {
	return proto.CompactTextString(m)
}
func (*ProxyArpInterfaces_InterfaceList_Interface) ProtoMessage() {}
func (*ProxyArpInterfaces_InterfaceList_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptorL3, []int{3, 0, 0}
}

func (m *ProxyArpInterfaces_InterfaceList_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// STN (Steal The NIC) feature table
type STNTable struct {
	StnEntries []*STNTable_STNTableEntry `protobuf:"bytes,1,rep,name=stn_entries,json=stnEntries" json:"stn_entries,omitempty"`
}

func (m *STNTable) Reset()                    { *m = STNTable{} }
func (m *STNTable) String() string            { return proto.CompactTextString(m) }
func (*STNTable) ProtoMessage()               {}
func (*STNTable) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{4} }

func (m *STNTable) GetStnEntries() []*STNTable_STNTableEntry {
	if m != nil {
		return m.StnEntries
	}
	return nil
}

type STNTable_STNTableEntry struct {
	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Interface string `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
}

func (m *STNTable_STNTableEntry) Reset()                    { *m = STNTable_STNTableEntry{} }
func (m *STNTable_STNTableEntry) String() string            { return proto.CompactTextString(m) }
func (*STNTable_STNTableEntry) ProtoMessage()               {}
func (*STNTable_STNTableEntry) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{4, 0} }

func (m *STNTable_STNTableEntry) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *STNTable_STNTableEntry) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

// Enables/disables IP neighbor scanning
type IPScanNeighbor struct {
	Mode           IPScanNeighbor_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=l3.IPScanNeighbor_Mode" json:"mode,omitempty"`
	ScanInterval   uint32              `protobuf:"varint,2,opt,name=scan_interval,json=scanInterval,proto3" json:"scan_interval,omitempty"`
	MaxProcTime    uint32              `protobuf:"varint,3,opt,name=max_proc_time,json=maxProcTime,proto3" json:"max_proc_time,omitempty"`
	MaxUpdate      uint32              `protobuf:"varint,4,opt,name=max_update,json=maxUpdate,proto3" json:"max_update,omitempty"`
	ScanIntDelay   uint32              `protobuf:"varint,5,opt,name=scan_int_delay,json=scanIntDelay,proto3" json:"scan_int_delay,omitempty"`
	StaleThreshold uint32              `protobuf:"varint,6,opt,name=stale_threshold,json=staleThreshold,proto3" json:"stale_threshold,omitempty"`
}

func (m *IPScanNeighbor) Reset()                    { *m = IPScanNeighbor{} }
func (m *IPScanNeighbor) String() string            { return proto.CompactTextString(m) }
func (*IPScanNeighbor) ProtoMessage()               {}
func (*IPScanNeighbor) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{5} }

func (m *IPScanNeighbor) GetMode() IPScanNeighbor_Mode {
	if m != nil {
		return m.Mode
	}
	return IPScanNeighbor_DISABLED
}

func (m *IPScanNeighbor) GetScanInterval() uint32 {
	if m != nil {
		return m.ScanInterval
	}
	return 0
}

func (m *IPScanNeighbor) GetMaxProcTime() uint32 {
	if m != nil {
		return m.MaxProcTime
	}
	return 0
}

func (m *IPScanNeighbor) GetMaxUpdate() uint32 {
	if m != nil {
		return m.MaxUpdate
	}
	return 0
}

func (m *IPScanNeighbor) GetScanIntDelay() uint32 {
	if m != nil {
		return m.ScanIntDelay
	}
	return 0
}

func (m *IPScanNeighbor) GetStaleThreshold() uint32 {
	if m != nil {
		return m.StaleThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*StaticRoutes)(nil), "l3.StaticRoutes")
	proto.RegisterType((*StaticRoutes_Route)(nil), "l3.StaticRoutes.Route")
	proto.RegisterType((*ArpTable)(nil), "l3.ArpTable")
	proto.RegisterType((*ArpTable_ArpEntry)(nil), "l3.ArpTable.ArpEntry")
	proto.RegisterType((*ProxyArpRanges)(nil), "l3.ProxyArpRanges")
	proto.RegisterType((*ProxyArpRanges_RangeList)(nil), "l3.ProxyArpRanges.RangeList")
	proto.RegisterType((*ProxyArpRanges_RangeList_Range)(nil), "l3.ProxyArpRanges.RangeList.Range")
	proto.RegisterType((*ProxyArpInterfaces)(nil), "l3.ProxyArpInterfaces")
	proto.RegisterType((*ProxyArpInterfaces_InterfaceList)(nil), "l3.ProxyArpInterfaces.InterfaceList")
	proto.RegisterType((*ProxyArpInterfaces_InterfaceList_Interface)(nil), "l3.ProxyArpInterfaces.InterfaceList.Interface")
	proto.RegisterType((*STNTable)(nil), "l3.STNTable")
	proto.RegisterType((*STNTable_STNTableEntry)(nil), "l3.STNTable.STNTableEntry")
	proto.RegisterType((*IPScanNeighbor)(nil), "l3.IPScanNeighbor")
	proto.RegisterEnum("l3.StaticRoutes_Route_RouteType", StaticRoutes_Route_RouteType_name, StaticRoutes_Route_RouteType_value)
	proto.RegisterEnum("l3.IPScanNeighbor_Mode", IPScanNeighbor_Mode_name, IPScanNeighbor_Mode_value)
}

func init() { proto.RegisterFile("l3.proto", fileDescriptorL3) }

var fileDescriptorL3 = []byte{
	// 826 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0x5e, 0x67, 0x12, 0x8f, 0x5d, 0x1e, 0x67, 0x43, 0x8b, 0x9d, 0x35, 0xd1, 0xb0, 0x04, 0xb3,
	0x12, 0x23, 0x21, 0x7c, 0x98, 0xac, 0xf6, 0xc0, 0x8a, 0x43, 0x56, 0x33, 0x68, 0x2d, 0xcd, 0x66,
	0xa3, 0x4e, 0x98, 0xab, 0xd5, 0xb1, 0x3b, 0x49, 0x4b, 0x8e, 0x6d, 0x75, 0xf7, 0x84, 0xe4, 0xca,
	0x3b, 0xc0, 0x81, 0x27, 0xe0, 0x05, 0x78, 0x06, 0x5e, 0x86, 0x23, 0x47, 0x0e, 0xa8, 0xdb, 0x3f,
	0xf9, 0x11, 0x20, 0x2e, 0x71, 0xd5, 0x57, 0x5f, 0x57, 0x77, 0x55, 0x7f, 0x5d, 0x01, 0x2b, 0x1d,
	0x06, 0x05, 0xcf, 0x65, 0x8e, 0x5a, 0xe9, 0xd0, 0xff, 0xed, 0x0c, 0x2e, 0xa6, 0x92, 0x48, 0x16,
	0xe3, 0xfc, 0x51, 0x52, 0x81, 0x02, 0x30, 0xb9, 0xb6, 0x3c, 0x63, 0x70, 0x76, 0xed, 0xdc, 0x5c,
	0x06, 0xe9, 0x30, 0x38, 0x64, 0x04, 0xfa, 0x83, 0x2b, 0x56, 0xff, 0xcf, 0x16, 0x74, 0x34, 0x82,
	0x5e, 0x41, 0x5b, 0xee, 0x0a, 0xea, 0xc1, 0xc0, 0xb8, 0xee, 0xde, 0x0c, 0xfe, 0x79, 0x5d, 0xf9,
	0x3b, 0xdb, 0x15, 0x14, 0x6b, 0x36, 0x7a, 0x06, 0xe6, 0x86, 0x2f, 0x22, 0x96, 0x78, 0xc6, 0xc0,
	0xb8, 0x76, 0x71, 0x67, 0xc3, 0x17, 0x61, 0x82, 0x06, 0xe0, 0x24, 0x54, 0xc4, 0x9c, 0x15, 0x92,
	0xe5, 0x99, 0xd7, 0x1a, 0x18, 0xd7, 0x36, 0x3e, 0x84, 0xd0, 0x0b, 0x70, 0x12, 0x21, 0x23, 0x56,
	0x44, 0x24, 0x49, 0xb8, 0x77, 0xa6, 0x19, 0x76, 0x22, 0x64, 0x58, 0x8c, 0x92, 0x84, 0x23, 0x1f,
	0xdc, 0x8c, 0x6e, 0x65, 0xb4, 0xca, 0x2b, 0x46, 0xbb, 0xcc, 0xa1, 0xc0, 0x77, 0x79, 0xc9, 0xf9,
	0x1a, 0x50, 0xfe, 0x28, 0x97, 0x39, 0xcb, 0x96, 0x11, 0xcb, 0x24, 0xe5, 0x0b, 0x12, 0x53, 0xaf,
	0xa3, 0x89, 0x1f, 0xd5, 0x91, 0xb0, 0x0e, 0xa0, 0x4b, 0x30, 0x7f, 0xa0, 0x6c, 0xb9, 0x92, 0x9e,
	0xa9, 0xcf, 0x5a, 0x79, 0xe8, 0x05, 0x40, 0xc1, 0xe9, 0x82, 0x72, 0x9a, 0xc5, 0xd4, 0x3b, 0xd7,
	0xb1, 0x03, 0x04, 0x5d, 0x01, 0x6c, 0x18, 0x89, 0xaa, 0x3a, 0x2d, 0x1d, 0xb7, 0x36, 0x8c, 0x3c,
	0xa8, 0x52, 0xfd, 0x21, 0xd8, 0x4d, 0x53, 0x90, 0x0b, 0x76, 0x38, 0x9e, 0xe1, 0x51, 0xf4, 0x80,
	0xbf, 0xeb, 0x3d, 0xa9, 0xdc, 0x3b, 0xac, 0x5d, 0x03, 0x59, 0xd0, 0xbe, 0xc5, 0x1f, 0x26, 0xbd,
	0x96, 0xff, 0xbb, 0x01, 0xd6, 0x88, 0x17, 0x33, 0x32, 0x4f, 0x29, 0x7a, 0x0d, 0x0e, 0xe1, 0x45,
	0x44, 0x33, 0xc9, 0x59, 0x73, 0x71, 0xcf, 0xd4, 0x05, 0xd4, 0x14, 0x65, 0xdc, 0x65, 0x92, 0xef,
	0x30, 0x90, 0xd2, 0x62, 0x54, 0xf4, 0x7f, 0x2c, 0x93, 0xe8, 0x00, 0xba, 0x02, 0x7b, 0xdf, 0x02,
	0xa3, 0xec, 0x66, 0x03, 0xa0, 0x4f, 0x01, 0xaa, 0x4e, 0x53, 0x21, 0xaa, 0xeb, 0xb0, 0x99, 0xee,
	0x22, 0x15, 0x02, 0x7d, 0x0e, 0x17, 0xc5, 0x6a, 0x27, 0x1a, 0x42, 0x79, 0x1b, 0x8e, 0xc2, 0x6a,
	0xca, 0x25, 0x98, 0x42, 0xcb, 0x41, 0x5f, 0x84, 0x85, 0x2b, 0xcf, 0xff, 0xc3, 0x80, 0xee, 0x84,
	0xe7, 0xdb, 0xdd, 0x88, 0x17, 0x98, 0x64, 0x4b, 0x2a, 0xd0, 0xb7, 0xe0, 0x70, 0x65, 0x45, 0x29,
	0x13, 0xb2, 0xae, 0xe7, 0x4a, 0xd5, 0x73, 0x4c, 0x0c, 0xf4, 0xe7, 0x9e, 0x09, 0x89, 0x81, 0xd7,
	0xa6, 0xe8, 0xff, 0x62, 0x80, 0xdd, 0x44, 0xd0, 0xc7, 0xd0, 0x49, 0xc9, 0x9c, 0xa6, 0x55, 0x4d,
	0xa5, 0x83, 0xbe, 0x01, 0x53, 0xaf, 0x50, 0xb5, 0xa8, 0xec, 0xfe, 0x7f, 0x65, 0x2f, 0x2d, 0x5c,
	0xad, 0xe8, 0xbf, 0x81, 0x8e, 0x06, 0xd0, 0x27, 0x60, 0x2d, 0x18, 0xd7, 0x22, 0xac, 0xb2, 0x9f,
	0x6b, 0x3f, 0x2c, 0xd0, 0x73, 0x38, 0x4f, 0x49, 0x19, 0x29, 0x9b, 0x65, 0x2a, 0x37, 0x2c, 0xfc,
	0xbf, 0x0c, 0x40, 0xf5, 0x3e, 0x8d, 0xb2, 0x04, 0x7a, 0x0f, 0x4f, 0x9b, 0x66, 0x1f, 0x95, 0xfd,
	0xf2, 0xf0, 0x60, 0xfb, 0x05, 0x41, 0x63, 0xea, 0xf2, 0xbb, 0xec, 0xd0, 0x15, 0xfd, 0x9f, 0x0d,
	0x70, 0x8f, 0x18, 0xff, 0xd2, 0x86, 0x31, 0x40, 0xb3, 0xb2, 0x6e, 0x45, 0xf0, 0x7f, 0x76, 0xdc,
	0x7b, 0xf8, 0x20, 0x43, 0xff, 0x33, 0xb0, 0xf7, 0xcf, 0x05, 0x41, 0x3b, 0x23, 0xeb, 0x5a, 0x4c,
	0xda, 0xf6, 0x7f, 0x32, 0xc0, 0x9a, 0xce, 0xc6, 0xa5, 0x6e, 0xdf, 0x80, 0x23, 0x64, 0x76, 0xa2,
	0xdb, 0xbe, 0x1e, 0x1c, 0x15, 0xa5, 0x31, 0x2a, 0xf1, 0x0a, 0x99, 0xd5, 0xe2, 0xbd, 0x07, 0xf7,
	0x28, 0x78, 0x22, 0x51, 0xe3, 0x54, 0xa2, 0x47, 0xfa, 0x6e, 0x9d, 0xe8, 0xdb, 0xff, 0xb5, 0x05,
	0xdd, 0x70, 0x32, 0x8d, 0x49, 0x36, 0x56, 0x6f, 0x7a, 0x9e, 0x73, 0xf4, 0x15, 0xb4, 0xd7, 0x79,
	0x52, 0x1e, 0xbf, 0x7b, 0xf3, 0x5c, 0x1d, 0xeb, 0x98, 0x11, 0xbc, 0xcf, 0x13, 0x8a, 0x35, 0x09,
	0x7d, 0x01, 0xae, 0x88, 0x49, 0x56, 0x4e, 0x91, 0x0d, 0x49, 0xf5, 0x0e, 0x2e, 0xbe, 0x50, 0x60,
	0x58, 0x61, 0x6a, 0x24, 0xad, 0xc9, 0x36, 0x2a, 0x78, 0x1e, 0x47, 0x92, 0xad, 0xa9, 0x7e, 0x26,
	0x2e, 0x76, 0xd6, 0x64, 0x3b, 0xe1, 0x79, 0x3c, 0x63, 0x6b, 0xfd, 0xd0, 0x14, 0xe7, 0xb1, 0x48,
	0x88, 0xa4, 0xfa, 0xa9, 0xb8, 0xd8, 0x5e, 0x93, 0xed, 0xf7, 0x1a, 0x40, 0x2f, 0xa1, 0x5b, 0xef,
	0x13, 0x25, 0x34, 0x25, 0x3b, 0x3d, 0xad, 0xf6, 0x1b, 0xdd, 0x2a, 0x0c, 0x7d, 0x09, 0x4f, 0x85,
	0x24, 0x29, 0x8d, 0xe4, 0x8a, 0x53, 0xb1, 0xca, 0xd3, 0xa4, 0x9a, 0x58, 0x5d, 0x0d, 0xcf, 0x6a,
	0xd4, 0xbf, 0x81, 0xb6, 0x2a, 0x02, 0x5d, 0x80, 0x75, 0x1b, 0x4e, 0x47, 0x6f, 0xef, 0xef, 0x6e,
	0x7b, 0x4f, 0xd4, 0x98, 0x09, 0x27, 0x9b, 0x57, 0xe5, 0xc0, 0x09, 0x27, 0x9b, 0xd7, 0xbd, 0x96,
	0xb2, 0xde, 0x7e, 0x98, 0xbd, 0xeb, 0x9d, 0xcd, 0x4d, 0xfd, 0xef, 0x31, 0xfc, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x3f, 0x9d, 0x00, 0x97, 0x49, 0x06, 0x00, 0x00,
}
