// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/config.proto

package config // import "github.com/linkerd/linkerd2/controller/gen/config"

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Global struct {
	LinkerdNamespace string `protobuf:"bytes,1,opt,name=linkerd_namespace,json=linkerdNamespace,proto3" json:"linkerd_namespace,omitempty"`
	CniEnabled       bool   `protobuf:"varint,2,opt,name=cni_enabled,json=cniEnabled,proto3" json:"cni_enabled,omitempty"`
	Registry         string `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty"`
	Version          string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// null indicates TLS is disabled.
	// Otherwise, a non-null struct indicates the equivalence
	// of --tls=optional.
	IdentityContext      *IdentityContext `protobuf:"bytes,5,opt,name=identity_context,json=identityContext,proto3" json:"identity_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Global) Reset()         { *m = Global{} }
func (m *Global) String() string { return proto.CompactTextString(m) }
func (*Global) ProtoMessage()    {}
func (*Global) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_9cd844db14ba3623, []int{0}
}
func (m *Global) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Global.Unmarshal(m, b)
}
func (m *Global) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Global.Marshal(b, m, deterministic)
}
func (dst *Global) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Global.Merge(dst, src)
}
func (m *Global) XXX_Size() int {
	return xxx_messageInfo_Global.Size(m)
}
func (m *Global) XXX_DiscardUnknown() {
	xxx_messageInfo_Global.DiscardUnknown(m)
}

var xxx_messageInfo_Global proto.InternalMessageInfo

func (m *Global) GetLinkerdNamespace() string {
	if m != nil {
		return m.LinkerdNamespace
	}
	return ""
}

func (m *Global) GetCniEnabled() bool {
	if m != nil {
		return m.CniEnabled
	}
	return false
}

func (m *Global) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *Global) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Global) GetIdentityContext() *IdentityContext {
	if m != nil {
		return m.IdentityContext
	}
	return nil
}

type Proxy struct {
	ProxyImage              *Image                `protobuf:"bytes,1,opt,name=proxy_image,json=proxyImage,proto3" json:"proxy_image,omitempty"`
	ProxyInitImage          *Image                `protobuf:"bytes,2,opt,name=proxy_init_image,json=proxyInitImage,proto3" json:"proxy_init_image,omitempty"`
	DestinationApiPort      *Port                 `protobuf:"bytes,3,opt,name=destination_api_port,json=destinationApiPort,proto3" json:"destination_api_port,omitempty"`
	ControlPort             *Port                 `protobuf:"bytes,4,opt,name=control_port,json=controlPort,proto3" json:"control_port,omitempty"`
	IgnoreInboundPorts      []*Port               `protobuf:"bytes,5,rep,name=ignore_inbound_ports,json=ignoreInboundPorts,proto3" json:"ignore_inbound_ports,omitempty"`
	IgnoreOutboundPorts     []*Port               `protobuf:"bytes,6,rep,name=ignore_outbound_ports,json=ignoreOutboundPorts,proto3" json:"ignore_outbound_ports,omitempty"`
	InboundPort             *Port                 `protobuf:"bytes,7,opt,name=inbound_port,json=inboundPort,proto3" json:"inbound_port,omitempty"`
	MetricsPort             *Port                 `protobuf:"bytes,8,opt,name=metrics_port,json=metricsPort,proto3" json:"metrics_port,omitempty"`
	OutboundPort            *Port                 `protobuf:"bytes,9,opt,name=outbound_port,json=outboundPort,proto3" json:"outbound_port,omitempty"`
	Resource                *ResourceRequirements `protobuf:"bytes,10,opt,name=resource,proto3" json:"resource,omitempty"`
	ProxyUid                int64                 `protobuf:"varint,11,opt,name=proxy_uid,json=proxyUid,proto3" json:"proxy_uid,omitempty"`
	LogLevel                *LogLevel             `protobuf:"bytes,12,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	DisableExternalProfiles bool                  `protobuf:"varint,13,opt,name=disable_external_profiles,json=disableExternalProfiles,proto3" json:"disable_external_profiles,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *Proxy) Reset()         { *m = Proxy{} }
func (m *Proxy) String() string { return proto.CompactTextString(m) }
func (*Proxy) ProtoMessage()    {}
func (*Proxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_9cd844db14ba3623, []int{1}
}
func (m *Proxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proxy.Unmarshal(m, b)
}
func (m *Proxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proxy.Marshal(b, m, deterministic)
}
func (dst *Proxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proxy.Merge(dst, src)
}
func (m *Proxy) XXX_Size() int {
	return xxx_messageInfo_Proxy.Size(m)
}
func (m *Proxy) XXX_DiscardUnknown() {
	xxx_messageInfo_Proxy.DiscardUnknown(m)
}

var xxx_messageInfo_Proxy proto.InternalMessageInfo

func (m *Proxy) GetProxyImage() *Image {
	if m != nil {
		return m.ProxyImage
	}
	return nil
}

func (m *Proxy) GetProxyInitImage() *Image {
	if m != nil {
		return m.ProxyInitImage
	}
	return nil
}

func (m *Proxy) GetDestinationApiPort() *Port {
	if m != nil {
		return m.DestinationApiPort
	}
	return nil
}

func (m *Proxy) GetControlPort() *Port {
	if m != nil {
		return m.ControlPort
	}
	return nil
}

func (m *Proxy) GetIgnoreInboundPorts() []*Port {
	if m != nil {
		return m.IgnoreInboundPorts
	}
	return nil
}

func (m *Proxy) GetIgnoreOutboundPorts() []*Port {
	if m != nil {
		return m.IgnoreOutboundPorts
	}
	return nil
}

func (m *Proxy) GetInboundPort() *Port {
	if m != nil {
		return m.InboundPort
	}
	return nil
}

func (m *Proxy) GetMetricsPort() *Port {
	if m != nil {
		return m.MetricsPort
	}
	return nil
}

func (m *Proxy) GetOutboundPort() *Port {
	if m != nil {
		return m.OutboundPort
	}
	return nil
}

func (m *Proxy) GetResource() *ResourceRequirements {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Proxy) GetProxyUid() int64 {
	if m != nil {
		return m.ProxyUid
	}
	return 0
}

func (m *Proxy) GetLogLevel() *LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return nil
}

func (m *Proxy) GetDisableExternalProfiles() bool {
	if m != nil {
		return m.DisableExternalProfiles
	}
	return false
}

type Image struct {
	ImageName            string   `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	PullPolicy           string   `protobuf:"bytes,2,opt,name=pull_policy,json=pullPolicy,proto3" json:"pull_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_9cd844db14ba3623, []int{2}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *Image) GetPullPolicy() string {
	if m != nil {
		return m.PullPolicy
	}
	return ""
}

type Port struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_9cd844db14ba3623, []int{3}
}
func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (dst *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(dst, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type ResourceRequirements struct {
	RequestCpu           string   `protobuf:"bytes,1,opt,name=request_cpu,json=requestCpu,proto3" json:"request_cpu,omitempty"`
	RequestMemory        string   `protobuf:"bytes,2,opt,name=request_memory,json=requestMemory,proto3" json:"request_memory,omitempty"`
	LimitCpu             string   `protobuf:"bytes,3,opt,name=limit_cpu,json=limitCpu,proto3" json:"limit_cpu,omitempty"`
	LimitMemory          string   `protobuf:"bytes,4,opt,name=limit_memory,json=limitMemory,proto3" json:"limit_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceRequirements) Reset()         { *m = ResourceRequirements{} }
func (m *ResourceRequirements) String() string { return proto.CompactTextString(m) }
func (*ResourceRequirements) ProtoMessage()    {}
func (*ResourceRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_9cd844db14ba3623, []int{4}
}
func (m *ResourceRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceRequirements.Unmarshal(m, b)
}
func (m *ResourceRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceRequirements.Marshal(b, m, deterministic)
}
func (dst *ResourceRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceRequirements.Merge(dst, src)
}
func (m *ResourceRequirements) XXX_Size() int {
	return xxx_messageInfo_ResourceRequirements.Size(m)
}
func (m *ResourceRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceRequirements proto.InternalMessageInfo

func (m *ResourceRequirements) GetRequestCpu() string {
	if m != nil {
		return m.RequestCpu
	}
	return ""
}

func (m *ResourceRequirements) GetRequestMemory() string {
	if m != nil {
		return m.RequestMemory
	}
	return ""
}

func (m *ResourceRequirements) GetLimitCpu() string {
	if m != nil {
		return m.LimitCpu
	}
	return ""
}

func (m *ResourceRequirements) GetLimitMemory() string {
	if m != nil {
		return m.LimitMemory
	}
	return ""
}

type IdentityContext struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentityContext) Reset()         { *m = IdentityContext{} }
func (m *IdentityContext) String() string { return proto.CompactTextString(m) }
func (*IdentityContext) ProtoMessage()    {}
func (*IdentityContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_9cd844db14ba3623, []int{5}
}
func (m *IdentityContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityContext.Unmarshal(m, b)
}
func (m *IdentityContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityContext.Marshal(b, m, deterministic)
}
func (dst *IdentityContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityContext.Merge(dst, src)
}
func (m *IdentityContext) XXX_Size() int {
	return xxx_messageInfo_IdentityContext.Size(m)
}
func (m *IdentityContext) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityContext.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityContext proto.InternalMessageInfo

type LogLevel struct {
	Level                string   `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevel) Reset()         { *m = LogLevel{} }
func (m *LogLevel) String() string { return proto.CompactTextString(m) }
func (*LogLevel) ProtoMessage()    {}
func (*LogLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_9cd844db14ba3623, []int{6}
}
func (m *LogLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevel.Unmarshal(m, b)
}
func (m *LogLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevel.Marshal(b, m, deterministic)
}
func (dst *LogLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevel.Merge(dst, src)
}
func (m *LogLevel) XXX_Size() int {
	return xxx_messageInfo_LogLevel.Size(m)
}
func (m *LogLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevel.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevel proto.InternalMessageInfo

func (m *LogLevel) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func init() {
	proto.RegisterType((*Global)(nil), "linkerd2.config.Global")
	proto.RegisterType((*Proxy)(nil), "linkerd2.config.Proxy")
	proto.RegisterType((*Image)(nil), "linkerd2.config.Image")
	proto.RegisterType((*Port)(nil), "linkerd2.config.Port")
	proto.RegisterType((*ResourceRequirements)(nil), "linkerd2.config.ResourceRequirements")
	proto.RegisterType((*IdentityContext)(nil), "linkerd2.config.IdentityContext")
	proto.RegisterType((*LogLevel)(nil), "linkerd2.config.LogLevel")
}

func init() { proto.RegisterFile("config/config.proto", fileDescriptor_config_9cd844db14ba3623) }

var fileDescriptor_config_9cd844db14ba3623 = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdb, 0x6a, 0xdb, 0x40,
	0x10, 0xc5, 0x89, 0xed, 0xd8, 0x23, 0x3b, 0x97, 0x8d, 0xd3, 0x2a, 0x29, 0xa5, 0xae, 0x21, 0x60,
	0x28, 0xd8, 0xd4, 0x81, 0x36, 0xe4, 0xa9, 0x69, 0x08, 0x26, 0x34, 0x6d, 0x8d, 0xa0, 0x2f, 0x7d,
	0x11, 0xb2, 0xb4, 0x51, 0x87, 0xae, 0x76, 0x95, 0xd5, 0x2a, 0xc4, 0x1f, 0xd3, 0xfe, 0x57, 0xff,
	0xa6, 0xec, 0xc5, 0xa9, 0x13, 0x17, 0x3d, 0x69, 0x75, 0xe6, 0x9c, 0xb3, 0xa3, 0xd1, 0xcc, 0xc0,
	0x7e, 0x2c, 0xf8, 0x0d, 0xa6, 0x63, 0xfb, 0x18, 0xe5, 0x52, 0x28, 0x41, 0x76, 0x18, 0xf2, 0x9f,
	0x54, 0x26, 0x93, 0x91, 0x85, 0x07, 0x7f, 0x6a, 0xd0, 0x9c, 0x32, 0x31, 0x8f, 0x18, 0x79, 0x03,
	0x7b, 0x2e, 0x1a, 0xf2, 0x28, 0xa3, 0x45, 0x1e, 0xc5, 0xd4, 0xaf, 0xf5, 0x6b, 0xc3, 0x76, 0xb0,
	0xeb, 0x02, 0x5f, 0x96, 0x38, 0x79, 0x05, 0x5e, 0xcc, 0x31, 0xa4, 0x3c, 0x9a, 0x33, 0x9a, 0xf8,
	0x1b, 0xfd, 0xda, 0xb0, 0x15, 0x40, 0xcc, 0xf1, 0xd2, 0x22, 0xe4, 0x08, 0x5a, 0x92, 0xa6, 0x58,
	0x28, 0xb9, 0xf0, 0x37, 0x8d, 0xc9, 0xc3, 0x3b, 0xf1, 0x61, 0xeb, 0x8e, 0xca, 0x02, 0x05, 0xf7,
	0xeb, 0x26, 0xb4, 0x7c, 0x25, 0x9f, 0x60, 0x17, 0x13, 0xca, 0x15, 0xaa, 0x45, 0x18, 0x0b, 0xae,
	0xe8, 0xbd, 0xf2, 0x1b, 0xfd, 0xda, 0xd0, 0x9b, 0xf4, 0x47, 0x4f, 0x52, 0x1f, 0x5d, 0x39, 0xe2,
	0x85, 0xe5, 0x05, 0x3b, 0xf8, 0x18, 0x18, 0xfc, 0x6a, 0x42, 0x63, 0x26, 0xc5, 0xfd, 0x82, 0xbc,
	0x07, 0x2f, 0xd7, 0x87, 0x10, 0xb3, 0x28, 0xb5, 0x1f, 0xe5, 0x4d, 0x9e, 0xad, 0x3b, 0xea, 0x68,
	0x00, 0x86, 0x6a, 0xce, 0xe4, 0x03, 0xec, 0x3a, 0x21, 0x47, 0xe5, 0xd4, 0x1b, 0x95, 0xea, 0x6d,
	0xab, 0xe6, 0xa8, 0xac, 0xc3, 0x14, 0x7a, 0x09, 0x2d, 0x14, 0xf2, 0x48, 0xa1, 0xe0, 0x61, 0x94,
	0x63, 0x98, 0x0b, 0xa9, 0x4c, 0x4d, 0xbc, 0xc9, 0xc1, 0x9a, 0xcb, 0x4c, 0x48, 0x15, 0x90, 0x15,
	0xc9, 0x79, 0x8e, 0x1a, 0x23, 0xa7, 0xd0, 0xd1, 0x15, 0x91, 0x82, 0x59, 0x83, 0x7a, 0x95, 0x81,
	0xe7, 0xa8, 0x46, 0x39, 0x85, 0x1e, 0xa6, 0x5c, 0x48, 0x1a, 0x22, 0x9f, 0x8b, 0x92, 0x27, 0xc6,
	0xa0, 0xf0, 0x1b, 0xfd, 0xcd, 0x8a, 0x14, 0xac, 0xe4, 0xca, 0x2a, 0x34, 0x54, 0x90, 0x2b, 0x38,
	0x70, 0x46, 0xa2, 0x54, 0xab, 0x4e, 0xcd, 0x2a, 0xa7, 0x7d, 0xab, 0xf9, 0xea, 0x24, 0xd6, 0xea,
	0x14, 0x3a, 0xab, 0xc9, 0xf8, 0x5b, 0x95, 0x5f, 0x83, 0xff, 0xb2, 0xd0, 0xca, 0x8c, 0x2a, 0x89,
	0x71, 0x61, 0x95, 0xad, 0x4a, 0xa5, 0xa3, 0x1a, 0xe5, 0x19, 0x74, 0x1f, 0xe5, 0xed, 0xb7, 0xab,
	0xa4, 0x1d, 0xb1, 0x92, 0x30, 0x39, 0xd7, 0xed, 0x5c, 0x88, 0x52, 0xc6, 0xd4, 0x07, 0x23, 0x3b,
	0x5e, 0x93, 0x05, 0x8e, 0x10, 0xd0, 0xdb, 0x12, 0x25, 0xcd, 0x28, 0x57, 0x45, 0xf0, 0x20, 0x23,
	0x2f, 0xa0, 0x6d, 0x7b, 0xa9, 0xc4, 0xc4, 0xf7, 0xfa, 0xb5, 0xe1, 0x66, 0xd0, 0x32, 0xc0, 0x37,
	0x4c, 0xc8, 0x3b, 0x68, 0x33, 0x91, 0x86, 0x8c, 0xde, 0x51, 0xe6, 0x77, 0xcc, 0x05, 0x87, 0x6b,
	0x17, 0x5c, 0x8b, 0xf4, 0x5a, 0x13, 0x82, 0x16, 0x73, 0x27, 0x72, 0x06, 0x87, 0x09, 0x16, 0x7a,
	0xe4, 0x42, 0x7a, 0xaf, 0xa8, 0xe4, 0x11, 0x0b, 0x73, 0x29, 0x6e, 0x90, 0xd1, 0xc2, 0xef, 0x9a,
	0xa9, 0x7c, 0xee, 0x08, 0x97, 0x2e, 0x3e, 0x73, 0xe1, 0xc1, 0x14, 0x1a, 0xb6, 0x47, 0x5f, 0x02,
	0x98, 0xd6, 0x36, 0x73, 0xef, 0x46, 0xbe, 0x6d, 0x10, 0x3d, 0xf0, 0x7a, 0xd6, 0xf3, 0x92, 0xe9,
	0xb6, 0x63, 0x18, 0x2f, 0x4c, 0xff, 0xb7, 0x03, 0xd0, 0xd0, 0xcc, 0x20, 0x83, 0x23, 0xa8, 0x9b,
	0x22, 0x11, 0xa8, 0x9b, 0xba, 0x6a, 0x87, 0x6e, 0x60, 0xce, 0x83, 0xdf, 0x35, 0xe8, 0xfd, 0xaf,
	0x30, 0xda, 0x55, 0xd2, 0xdb, 0x92, 0x16, 0x2a, 0x8c, 0xf3, 0xd2, 0xdd, 0x0a, 0x0e, 0xba, 0xc8,
	0x4b, 0x72, 0x0c, 0xdb, 0x4b, 0x42, 0x46, 0x33, 0x21, 0x97, 0x37, 0x77, 0x1d, 0xfa, 0xd9, 0x80,
	0xba, 0xac, 0x0c, 0x33, 0xb4, 0x2e, 0x6e, 0xd3, 0x18, 0x40, 0x7b, 0xbc, 0x86, 0x8e, 0x0d, 0x3a,
	0x07, 0xbb, 0x6e, 0x3c, 0x83, 0x59, 0xfd, 0x60, 0x0f, 0x76, 0x9e, 0x6c, 0x92, 0x41, 0x1f, 0x5a,
	0xcb, 0x52, 0x93, 0x1e, 0x34, 0xec, 0x4f, 0xb1, 0x09, 0xda, 0x97, 0x8f, 0x27, 0xdf, 0xdf, 0xa6,
	0xa8, 0x7e, 0x94, 0xf3, 0x51, 0x2c, 0xb2, 0xb1, 0xfb, 0x4f, 0xcb, 0xe7, 0x64, 0xec, 0xa6, 0x8f,
	0x51, 0x39, 0x4e, 0x29, 0x77, 0x2b, 0x78, 0xde, 0x34, 0x3b, 0xf8, 0xe4, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x67, 0x83, 0x49, 0x65, 0x9a, 0x05, 0x00, 0x00,
}