// Code generated by protoc-gen-go. DO NOT EDIT.
// source: status.proto

package oaas

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

type Code int32

const (
	Code_OK                   Code = 0
	Code_UNKNOWN              Code = 1
	Code_BALANCE_INSUFFICIENT Code = 2
)

var Code_name = map[int32]string{
	0: "OK",
	1: "UNKNOWN",
	2: "BALANCE_INSUFFICIENT",
}

var Code_value = map[string]int32{
	"OK":                   0,
	"UNKNOWN":              1,
	"BALANCE_INSUFFICIENT": 2,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dfe4fce6682daf5b, []int{0}
}

type Status struct {
	Code                 Code     `protobuf:"varint,1,opt,name=code,proto3,enum=oaas.Code" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfe4fce6682daf5b, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_OK
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("oaas.Code", Code_name, Code_value)
	proto.RegisterType((*Status)(nil), "oaas.Status")
}

func init() { proto.RegisterFile("status.proto", fileDescriptor_dfe4fce6682daf5b) }

var fileDescriptor_dfe4fce6682daf5b = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x49, 0x2c,
	0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4f, 0x4c, 0x2c, 0x56, 0x72,
	0xe2, 0x62, 0x0b, 0x06, 0x8b, 0x0a, 0xc9, 0x71, 0xb1, 0x24, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0xf0, 0x19, 0x71, 0xe9, 0x81, 0xa4, 0xf5, 0x9c, 0xf3, 0x53, 0x52, 0x83, 0xc0, 0xe2,
	0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x30, 0xae, 0x96, 0x29, 0x17, 0x0b, 0x48, 0x9d, 0x10, 0x1b, 0x17, 0x93, 0xbf, 0xb7,
	0x00, 0x83, 0x10, 0x37, 0x17, 0x7b, 0xa8, 0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0xa3, 0x90,
	0x04, 0x97, 0x88, 0x93, 0xa3, 0x8f, 0xa3, 0x9f, 0xb3, 0x6b, 0xbc, 0xa7, 0x5f, 0x70, 0xa8, 0x9b,
	0x9b, 0xa7, 0xb3, 0xa7, 0xab, 0x5f, 0x88, 0x00, 0x53, 0x12, 0x1b, 0xd8, 0x1d, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x83, 0x03, 0x5f, 0xf2, 0x97, 0x00, 0x00, 0x00,
}
