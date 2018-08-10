// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/hashicorp/nomad/plugins/base/proto/base.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hclspec "github.com/hashicorp/nomad/plugins/shared/hclspec"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PluginType enumerates the type of plugins Nomad supports
type PluginType int32

const (
	PluginType_UNKNOWN PluginType = 0
	PluginType_BASE    PluginType = 1
	PluginType_DRIVER  PluginType = 2
	PluginType_DEVICE  PluginType = 3
)

var PluginType_name = map[int32]string{
	0: "UNKNOWN",
	1: "BASE",
	2: "DRIVER",
	3: "DEVICE",
}
var PluginType_value = map[string]int32{
	"UNKNOWN": 0,
	"BASE":    1,
	"DRIVER":  2,
	"DEVICE":  3,
}

func (x PluginType) String() string {
	return proto.EnumName(PluginType_name, int32(x))
}
func (PluginType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_base_d24a6a23488adb59, []int{0}
}

// PluginInfoRequest is used to request the plugins basic information.
type PluginInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginInfoRequest) Reset()         { *m = PluginInfoRequest{} }
func (m *PluginInfoRequest) String() string { return proto.CompactTextString(m) }
func (*PluginInfoRequest) ProtoMessage()    {}
func (*PluginInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_d24a6a23488adb59, []int{0}
}
func (m *PluginInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInfoRequest.Unmarshal(m, b)
}
func (m *PluginInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInfoRequest.Marshal(b, m, deterministic)
}
func (dst *PluginInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInfoRequest.Merge(dst, src)
}
func (m *PluginInfoRequest) XXX_Size() int {
	return xxx_messageInfo_PluginInfoRequest.Size(m)
}
func (m *PluginInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInfoRequest proto.InternalMessageInfo

// PluginInfoResponse returns basic information about the plugin such
// that Nomad can decide whether to load the plugin or not.
type PluginInfoResponse struct {
	// type indicates what type of plugin this is.
	Type PluginType `protobuf:"varint,1,opt,name=type,proto3,enum=hashicorp.nomad.plugins.base.proto.PluginType" json:"type,omitempty"`
	// plugin_api_version indicates the version of the Nomad Plugin API
	// this plugin is built against.
	PluginApiVersion string `protobuf:"bytes,2,opt,name=plugin_api_version,json=pluginApiVersion,proto3" json:"plugin_api_version,omitempty"`
	// plugin_version is the semver version of this individual plugin.
	// This is divorce from Nomad’s development and versioning.
	PluginVersion string `protobuf:"bytes,3,opt,name=plugin_version,json=pluginVersion,proto3" json:"plugin_version,omitempty"`
	// name is the name of the plugin
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginInfoResponse) Reset()         { *m = PluginInfoResponse{} }
func (m *PluginInfoResponse) String() string { return proto.CompactTextString(m) }
func (*PluginInfoResponse) ProtoMessage()    {}
func (*PluginInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_d24a6a23488adb59, []int{1}
}
func (m *PluginInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInfoResponse.Unmarshal(m, b)
}
func (m *PluginInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInfoResponse.Marshal(b, m, deterministic)
}
func (dst *PluginInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInfoResponse.Merge(dst, src)
}
func (m *PluginInfoResponse) XXX_Size() int {
	return xxx_messageInfo_PluginInfoResponse.Size(m)
}
func (m *PluginInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInfoResponse proto.InternalMessageInfo

func (m *PluginInfoResponse) GetType() PluginType {
	if m != nil {
		return m.Type
	}
	return PluginType_UNKNOWN
}

func (m *PluginInfoResponse) GetPluginApiVersion() string {
	if m != nil {
		return m.PluginApiVersion
	}
	return ""
}

func (m *PluginInfoResponse) GetPluginVersion() string {
	if m != nil {
		return m.PluginVersion
	}
	return ""
}

func (m *PluginInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// ConfigSchemaRequest is used to request the configurations schema.
type ConfigSchemaRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigSchemaRequest) Reset()         { *m = ConfigSchemaRequest{} }
func (m *ConfigSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigSchemaRequest) ProtoMessage()    {}
func (*ConfigSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_d24a6a23488adb59, []int{2}
}
func (m *ConfigSchemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSchemaRequest.Unmarshal(m, b)
}
func (m *ConfigSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSchemaRequest.Marshal(b, m, deterministic)
}
func (dst *ConfigSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSchemaRequest.Merge(dst, src)
}
func (m *ConfigSchemaRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigSchemaRequest.Size(m)
}
func (m *ConfigSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSchemaRequest proto.InternalMessageInfo

// ConfigSchemaResponse returns the plugins configuration schema.
type ConfigSchemaResponse struct {
	// spec is the plugins configuration schema
	Spec                 *hclspec.Spec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ConfigSchemaResponse) Reset()         { *m = ConfigSchemaResponse{} }
func (m *ConfigSchemaResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigSchemaResponse) ProtoMessage()    {}
func (*ConfigSchemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_d24a6a23488adb59, []int{3}
}
func (m *ConfigSchemaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSchemaResponse.Unmarshal(m, b)
}
func (m *ConfigSchemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSchemaResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigSchemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSchemaResponse.Merge(dst, src)
}
func (m *ConfigSchemaResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigSchemaResponse.Size(m)
}
func (m *ConfigSchemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSchemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSchemaResponse proto.InternalMessageInfo

func (m *ConfigSchemaResponse) GetSpec() *hclspec.Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// SetConfigRequest is used to set the configuration
type SetConfigRequest struct {
	// msgpack_config is the configuration encoded as MessagePack.
	MsgpackConfig        []byte   `protobuf:"bytes,1,opt,name=msgpack_config,json=msgpackConfig,proto3" json:"msgpack_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetConfigRequest) Reset()         { *m = SetConfigRequest{} }
func (m *SetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*SetConfigRequest) ProtoMessage()    {}
func (*SetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_d24a6a23488adb59, []int{4}
}
func (m *SetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigRequest.Unmarshal(m, b)
}
func (m *SetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigRequest.Marshal(b, m, deterministic)
}
func (dst *SetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigRequest.Merge(dst, src)
}
func (m *SetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_SetConfigRequest.Size(m)
}
func (m *SetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigRequest proto.InternalMessageInfo

func (m *SetConfigRequest) GetMsgpackConfig() []byte {
	if m != nil {
		return m.MsgpackConfig
	}
	return nil
}

// SetConfigResponse is used to respond to setting the configuration
type SetConfigResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetConfigResponse) Reset()         { *m = SetConfigResponse{} }
func (m *SetConfigResponse) String() string { return proto.CompactTextString(m) }
func (*SetConfigResponse) ProtoMessage()    {}
func (*SetConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_d24a6a23488adb59, []int{5}
}
func (m *SetConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigResponse.Unmarshal(m, b)
}
func (m *SetConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigResponse.Marshal(b, m, deterministic)
}
func (dst *SetConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigResponse.Merge(dst, src)
}
func (m *SetConfigResponse) XXX_Size() int {
	return xxx_messageInfo_SetConfigResponse.Size(m)
}
func (m *SetConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PluginInfoRequest)(nil), "hashicorp.nomad.plugins.base.proto.PluginInfoRequest")
	proto.RegisterType((*PluginInfoResponse)(nil), "hashicorp.nomad.plugins.base.proto.PluginInfoResponse")
	proto.RegisterType((*ConfigSchemaRequest)(nil), "hashicorp.nomad.plugins.base.proto.ConfigSchemaRequest")
	proto.RegisterType((*ConfigSchemaResponse)(nil), "hashicorp.nomad.plugins.base.proto.ConfigSchemaResponse")
	proto.RegisterType((*SetConfigRequest)(nil), "hashicorp.nomad.plugins.base.proto.SetConfigRequest")
	proto.RegisterType((*SetConfigResponse)(nil), "hashicorp.nomad.plugins.base.proto.SetConfigResponse")
	proto.RegisterEnum("hashicorp.nomad.plugins.base.proto.PluginType", PluginType_name, PluginType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BasePluginClient is the client API for BasePlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BasePluginClient interface {
	// PluginInfo describes the type and version of a plugin.
	PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error)
	// ConfigSchema returns the schema for parsing the plugins configuration.
	ConfigSchema(ctx context.Context, in *ConfigSchemaRequest, opts ...grpc.CallOption) (*ConfigSchemaResponse, error)
	// SetConfig is used to set the configuration.
	SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error)
}

type basePluginClient struct {
	cc *grpc.ClientConn
}

func NewBasePluginClient(cc *grpc.ClientConn) BasePluginClient {
	return &basePluginClient{cc}
}

func (c *basePluginClient) PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error) {
	out := new(PluginInfoResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.base.proto.BasePlugin/PluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basePluginClient) ConfigSchema(ctx context.Context, in *ConfigSchemaRequest, opts ...grpc.CallOption) (*ConfigSchemaResponse, error) {
	out := new(ConfigSchemaResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.base.proto.BasePlugin/ConfigSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basePluginClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error) {
	out := new(SetConfigResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.base.proto.BasePlugin/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasePluginServer is the server API for BasePlugin service.
type BasePluginServer interface {
	// PluginInfo describes the type and version of a plugin.
	PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoResponse, error)
	// ConfigSchema returns the schema for parsing the plugins configuration.
	ConfigSchema(context.Context, *ConfigSchemaRequest) (*ConfigSchemaResponse, error)
	// SetConfig is used to set the configuration.
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)
}

func RegisterBasePluginServer(s *grpc.Server, srv BasePluginServer) {
	s.RegisterService(&_BasePlugin_serviceDesc, srv)
}

func _BasePlugin_PluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasePluginServer).PluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.base.proto.BasePlugin/PluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasePluginServer).PluginInfo(ctx, req.(*PluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasePlugin_ConfigSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasePluginServer).ConfigSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.base.proto.BasePlugin/ConfigSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasePluginServer).ConfigSchema(ctx, req.(*ConfigSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasePlugin_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasePluginServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.base.proto.BasePlugin/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasePluginServer).SetConfig(ctx, req.(*SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BasePlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad.plugins.base.proto.BasePlugin",
	HandlerType: (*BasePluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginInfo",
			Handler:    _BasePlugin_PluginInfo_Handler,
		},
		{
			MethodName: "ConfigSchema",
			Handler:    _BasePlugin_ConfigSchema_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _BasePlugin_SetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/hashicorp/nomad/plugins/base/proto/base.proto",
}

func init() {
	proto.RegisterFile("github.com/hashicorp/nomad/plugins/base/proto/base.proto", fileDescriptor_base_d24a6a23488adb59)
}

var fileDescriptor_base_d24a6a23488adb59 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x2f, 0xd7, 0x78, 0xe7, 0xcd, 0xfd, 0x20, 0xee, 0x29, 0x94, 0x3c, 0x1d, 0x01, 0xe1,
	0x90, 0x63, 0x03, 0xd5, 0xd3, 0x13, 0x5f, 0xae, 0xa9, 0x79, 0x28, 0x42, 0x95, 0x44, 0xab, 0xf8,
	0x12, 0xb6, 0xdb, 0x6d, 0x12, 0x6c, 0xb2, 0x6b, 0x36, 0x15, 0x2a, 0xf8, 0xe4, 0xb3, 0x7f, 0x94,
	0xff, 0x99, 0x64, 0x37, 0x69, 0xa3, 0x28, 0xa6, 0x4f, 0x19, 0x66, 0x3e, 0xf3, 0xdd, 0x99, 0xef,
	0x04, 0x6e, 0xe2, 0xb4, 0x4c, 0x56, 0x33, 0x4c, 0x79, 0xe6, 0x26, 0x44, 0x26, 0x29, 0xe5, 0x85,
	0x70, 0x73, 0x9e, 0x91, 0xb9, 0x2b, 0x96, 0xab, 0x38, 0xcd, 0xa5, 0x3b, 0x23, 0x92, 0xb9, 0xa2,
	0xe0, 0x25, 0x57, 0x21, 0x56, 0x21, 0x72, 0x36, 0x38, 0x56, 0x38, 0xae, 0x71, 0xbc, 0x65, 0xec,
	0xdb, 0x0e, 0xea, 0x32, 0x21, 0x05, 0x9b, 0xbb, 0x09, 0x5d, 0x4a, 0xc1, 0x68, 0xf5, 0x8d, 0xaa,
	0x40, 0x2b, 0x38, 0xe7, 0x70, 0xef, 0x8d, 0x02, 0xc7, 0xf9, 0x82, 0x07, 0xec, 0xf3, 0x8a, 0xc9,
	0xd2, 0xf9, 0x69, 0x00, 0x6a, 0x67, 0xa5, 0xe0, 0xb9, 0x64, 0xc8, 0x03, 0xb3, 0x5c, 0x0b, 0xd6,
	0x37, 0x2e, 0x8c, 0xcb, 0xb3, 0x01, 0xc6, 0xff, 0x1f, 0x10, 0x6b, 0x95, 0xb7, 0x6b, 0xc1, 0x02,
	0xd5, 0x8b, 0xae, 0x00, 0x69, 0x2c, 0x22, 0x22, 0x8d, 0xbe, 0xb0, 0x42, 0xa6, 0x3c, 0xef, 0xef,
	0x5f, 0x18, 0x97, 0x47, 0x81, 0xa5, 0x2b, 0x43, 0x91, 0x4e, 0x75, 0x1e, 0x3d, 0x84, 0xb3, 0x9a,
	0x6e, 0xc8, 0x9e, 0x22, 0x4f, 0x75, 0xb6, 0xc1, 0x10, 0x98, 0x39, 0xc9, 0x58, 0xdf, 0x54, 0x45,
	0x15, 0x3b, 0x0f, 0xe0, 0x7c, 0xc4, 0xf3, 0x45, 0x1a, 0x87, 0x34, 0x61, 0x19, 0x69, 0x56, 0xfb,
	0x00, 0xf7, 0x7f, 0x4f, 0xd7, 0xbb, 0xdd, 0x82, 0x59, 0xb9, 0xa2, 0x76, 0x3b, 0x1e, 0x5c, 0xfd,
	0x73, 0x37, 0xed, 0x26, 0xae, 0xdd, 0xc4, 0xa1, 0x60, 0x34, 0x50, 0x9d, 0xce, 0x73, 0xb0, 0x42,
	0x56, 0x6a, 0xf1, 0xfa, 0xb5, 0x6a, 0xfe, 0x4c, 0xc6, 0x82, 0xd0, 0x4f, 0x11, 0x55, 0x05, 0xa5,
	0x7f, 0x12, 0x9c, 0xd6, 0x59, 0x4d, 0x57, 0x47, 0x68, 0xb5, 0xea, 0x89, 0x1e, 0xbd, 0x00, 0xd8,
	0xba, 0x87, 0x8e, 0xe1, 0xf0, 0xdd, 0xe4, 0xd5, 0xe4, 0xf5, 0xfb, 0x89, 0xb5, 0x87, 0xee, 0x82,
	0xe9, 0x0d, 0x43, 0xdf, 0x32, 0x10, 0xc0, 0xc1, 0xcb, 0x60, 0x3c, 0xf5, 0x03, 0x6b, 0x5f, 0xc5,
	0xfe, 0x74, 0x3c, 0xf2, 0xad, 0xde, 0xe0, 0x47, 0x0f, 0xc0, 0x23, 0x92, 0x69, 0x05, 0xf4, 0xad,
	0xd1, 0xaa, 0xee, 0x89, 0xae, 0xbb, 0x5f, 0xae, 0xf5, 0x57, 0xd8, 0x4f, 0x77, 0x6d, 0xd3, 0x8b,
	0x38, 0x7b, 0xe8, 0xbb, 0x01, 0x27, 0x6d, 0xd7, 0xd1, 0xb3, 0x2e, 0x52, 0x7f, 0x39, 0x9f, 0x7d,
	0xb3, 0x7b, 0xe3, 0x66, 0x8a, 0xaf, 0x70, 0xb4, 0x71, 0x19, 0x3d, 0xe9, 0x22, 0xf4, 0xe7, 0x3d,
	0xed, 0xeb, 0x1d, 0xbb, 0x9a, 0xb7, 0xbd, 0xc3, 0x8f, 0x77, 0x54, 0x71, 0x76, 0xa0, 0x3e, 0x8f,
	0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x73, 0x54, 0x60, 0xaa, 0x18, 0x04, 0x00, 0x00,
}
