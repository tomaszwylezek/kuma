// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/plugins/ca/builtin/config/builtin_ca_config.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// BuiltinCertificateAuthorityConfig defines configuration for Builtin CA
// plugin
type BuiltinCertificateAuthorityConfig struct {
	// Configuration of CA Certificate
	CaCert               *BuiltinCertificateAuthorityConfig_CaCert `protobuf:"bytes,1,opt,name=caCert,proto3" json:"caCert,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *BuiltinCertificateAuthorityConfig) Reset()         { *m = BuiltinCertificateAuthorityConfig{} }
func (m *BuiltinCertificateAuthorityConfig) String() string { return proto.CompactTextString(m) }
func (*BuiltinCertificateAuthorityConfig) ProtoMessage()    {}
func (*BuiltinCertificateAuthorityConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_20d0073fd2f18c7e, []int{0}
}

func (m *BuiltinCertificateAuthorityConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuiltinCertificateAuthorityConfig.Unmarshal(m, b)
}
func (m *BuiltinCertificateAuthorityConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuiltinCertificateAuthorityConfig.Marshal(b, m, deterministic)
}
func (m *BuiltinCertificateAuthorityConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuiltinCertificateAuthorityConfig.Merge(m, src)
}
func (m *BuiltinCertificateAuthorityConfig) XXX_Size() int {
	return xxx_messageInfo_BuiltinCertificateAuthorityConfig.Size(m)
}
func (m *BuiltinCertificateAuthorityConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BuiltinCertificateAuthorityConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BuiltinCertificateAuthorityConfig proto.InternalMessageInfo

func (m *BuiltinCertificateAuthorityConfig) GetCaCert() *BuiltinCertificateAuthorityConfig_CaCert {
	if m != nil {
		return m.CaCert
	}
	return nil
}

// CaCert defines configuration for Certificate of CA.
type BuiltinCertificateAuthorityConfig_CaCert struct {
	// RSAbits of the certificate
	RSAbits *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=RSAbits,proto3" json:"RSAbits,omitempty"`
	// Expiration time of the certificate
	Expiration           string   `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuiltinCertificateAuthorityConfig_CaCert) Reset() {
	*m = BuiltinCertificateAuthorityConfig_CaCert{}
}
func (m *BuiltinCertificateAuthorityConfig_CaCert) String() string { return proto.CompactTextString(m) }
func (*BuiltinCertificateAuthorityConfig_CaCert) ProtoMessage()    {}
func (*BuiltinCertificateAuthorityConfig_CaCert) Descriptor() ([]byte, []int) {
	return fileDescriptor_20d0073fd2f18c7e, []int{0, 0}
}

func (m *BuiltinCertificateAuthorityConfig_CaCert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuiltinCertificateAuthorityConfig_CaCert.Unmarshal(m, b)
}
func (m *BuiltinCertificateAuthorityConfig_CaCert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuiltinCertificateAuthorityConfig_CaCert.Marshal(b, m, deterministic)
}
func (m *BuiltinCertificateAuthorityConfig_CaCert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuiltinCertificateAuthorityConfig_CaCert.Merge(m, src)
}
func (m *BuiltinCertificateAuthorityConfig_CaCert) XXX_Size() int {
	return xxx_messageInfo_BuiltinCertificateAuthorityConfig_CaCert.Size(m)
}
func (m *BuiltinCertificateAuthorityConfig_CaCert) XXX_DiscardUnknown() {
	xxx_messageInfo_BuiltinCertificateAuthorityConfig_CaCert.DiscardUnknown(m)
}

var xxx_messageInfo_BuiltinCertificateAuthorityConfig_CaCert proto.InternalMessageInfo

func (m *BuiltinCertificateAuthorityConfig_CaCert) GetRSAbits() *wrappers.UInt32Value {
	if m != nil {
		return m.RSAbits
	}
	return nil
}

func (m *BuiltinCertificateAuthorityConfig_CaCert) GetExpiration() string {
	if m != nil {
		return m.Expiration
	}
	return ""
}

func init() {
	proto.RegisterType((*BuiltinCertificateAuthorityConfig)(nil), "kuma.plugins.ca.BuiltinCertificateAuthorityConfig")
	proto.RegisterType((*BuiltinCertificateAuthorityConfig_CaCert)(nil), "kuma.plugins.ca.BuiltinCertificateAuthorityConfig.CaCert")
}

func init() {
	proto.RegisterFile("pkg/plugins/ca/builtin/config/builtin_ca_config.proto", fileDescriptor_20d0073fd2f18c7e)
}

var fileDescriptor_20d0073fd2f18c7e = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x59, 0x8b, 0x55, 0x63, 0x21, 0x6c, 0x75, 0x1c, 0x72, 0x9c, 0x56, 0x57, 0x4d, 0xe0,
	0x0e, 0x05, 0xcb, 0xbb, 0xad, 0x2c, 0x5d, 0xd1, 0xc2, 0xe6, 0x9c, 0x0d, 0xd9, 0x38, 0xdc, 0x9a,
	0x84, 0xec, 0x04, 0xf5, 0x9f, 0xfa, 0x73, 0xc4, 0x64, 0x17, 0xc4, 0xc6, 0x72, 0xde, 0x9b, 0xf7,
	0xcd, 0x63, 0xc4, 0xb5, 0x3f, 0x18, 0xe9, 0xfb, 0x68, 0xc8, 0x0e, 0x52, 0xa1, 0x6c, 0x23, 0xf5,
	0x4c, 0x56, 0x2a, 0x67, 0x3b, 0x32, 0xd3, 0xb8, 0x57, 0xb8, 0xcf, 0x0a, 0xf8, 0xe0, 0xd8, 0x55,
	0xe7, 0x87, 0xf8, 0x86, 0x30, 0xe6, 0x40, 0xe1, 0x7c, 0x61, 0x9c, 0x33, 0xbd, 0x96, 0xc9, 0x6e,
	0x63, 0x27, 0xdf, 0x03, 0x7a, 0xaf, 0xc3, 0x90, 0x03, 0x57, 0x5f, 0x85, 0xb8, 0xdc, 0x65, 0x58,
	0xad, 0x03, 0x53, 0x47, 0x0a, 0x59, 0x6f, 0x23, 0xbf, 0xba, 0x40, 0xfc, 0x59, 0x27, 0x78, 0x75,
	0x2f, 0x4a, 0x85, 0x3f, 0xfe, 0xac, 0x58, 0x16, 0xab, 0xb3, 0xf5, 0x2d, 0xfc, 0xb9, 0x03, 0xff,
	0x32, 0xa0, 0x4e, 0x80, 0x66, 0x04, 0xcd, 0x5f, 0x44, 0x99, 0x95, 0xea, 0x46, 0x1c, 0x37, 0x0f,
	0xdb, 0x96, 0x78, 0x18, 0xe9, 0x17, 0x90, 0x4b, 0xc3, 0x54, 0x1a, 0x1e, 0xef, 0x2c, 0x6f, 0xd6,
	0x4f, 0xd8, 0x47, 0xdd, 0x4c, 0xcb, 0xd5, 0x42, 0x08, 0xfd, 0xe1, 0x29, 0x20, 0x93, 0xb3, 0xb3,
	0xa3, 0x65, 0xb1, 0x3a, 0x6d, 0x7e, 0x29, 0xbb, 0x93, 0xe7, 0x32, 0xff, 0xa6, 0x2d, 0x13, 0x68,
	0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x8f, 0xfc, 0xd5, 0x55, 0x01, 0x00, 0x00,
}