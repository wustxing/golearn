// Code generated by protoc-gen-go. DO NOT EDIT.
// source: map.proto

package test

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

type GenderType int32

const (
	GenderType_NULL GenderType = 0
	GenderType_GIRL GenderType = 1
	GenderType_BOY  GenderType = 2
)

var GenderType_name = map[int32]string{
	0: "NULL",
	1: "GIRL",
	2: "BOY",
}
var GenderType_value = map[string]int32{
	"NULL": 0,
	"GIRL": 1,
	"BOY":  2,
}

func (x GenderType) String() string {
	return proto.EnumName(GenderType_name, int32(x))
}
func (GenderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_map_044c0925a51412e8, []int{0}
}

type ComplexObj struct {
	Map                  map[string]*MapValue `protobuf:"bytes,1,rep,name=map" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Complexid            int32                `protobuf:"varint,2,opt,name=complexid" json:"complexid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ComplexObj) Reset()         { *m = ComplexObj{} }
func (m *ComplexObj) String() string { return proto.CompactTextString(m) }
func (*ComplexObj) ProtoMessage()    {}
func (*ComplexObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_map_044c0925a51412e8, []int{0}
}
func (m *ComplexObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComplexObj.Unmarshal(m, b)
}
func (m *ComplexObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComplexObj.Marshal(b, m, deterministic)
}
func (dst *ComplexObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplexObj.Merge(dst, src)
}
func (m *ComplexObj) XXX_Size() int {
	return xxx_messageInfo_ComplexObj.Size(m)
}
func (m *ComplexObj) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplexObj.DiscardUnknown(m)
}

var xxx_messageInfo_ComplexObj proto.InternalMessageInfo

func (m *ComplexObj) GetMap() map[string]*MapValue {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *ComplexObj) GetComplexid() int32 {
	if m != nil {
		return m.Complexid
	}
	return 0
}

type MapValue struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age                  int32      `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	GenderType           GenderType `protobuf:"varint,3,opt,name=genderType,enum=test.GenderType" json:"genderType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MapValue) Reset()         { *m = MapValue{} }
func (m *MapValue) String() string { return proto.CompactTextString(m) }
func (*MapValue) ProtoMessage()    {}
func (*MapValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_map_044c0925a51412e8, []int{1}
}
func (m *MapValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapValue.Unmarshal(m, b)
}
func (m *MapValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapValue.Marshal(b, m, deterministic)
}
func (dst *MapValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapValue.Merge(dst, src)
}
func (m *MapValue) XXX_Size() int {
	return xxx_messageInfo_MapValue.Size(m)
}
func (m *MapValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MapValue.DiscardUnknown(m)
}

var xxx_messageInfo_MapValue proto.InternalMessageInfo

func (m *MapValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MapValue) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *MapValue) GetGenderType() GenderType {
	if m != nil {
		return m.GenderType
	}
	return GenderType_NULL
}

func init() {
	proto.RegisterType((*ComplexObj)(nil), "test.ComplexObj")
	proto.RegisterMapType((map[string]*MapValue)(nil), "test.ComplexObj.MapEntry")
	proto.RegisterType((*MapValue)(nil), "test.MapValue")
	proto.RegisterEnum("test.GenderType", GenderType_name, GenderType_value)
}

func init() { proto.RegisterFile("map.proto", fileDescriptor_map_044c0925a51412e8) }

var fileDescriptor_map_044c0925a51412e8 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x9d, 0x6c, 0xaa, 0xcd, 0x14, 0xca, 0x32, 0xa7, 0x28, 0x1e, 0x42, 0xf1, 0x10, 0x15,
	0x82, 0xc4, 0x8b, 0x78, 0x54, 0xb4, 0x08, 0xa9, 0x85, 0x45, 0x05, 0x8f, 0x1b, 0x3b, 0x14, 0xb5,
	0x49, 0x96, 0x18, 0xc5, 0xfc, 0x1a, 0xff, 0xaa, 0xec, 0x26, 0x9a, 0xde, 0x1e, 0xef, 0x7d, 0xf3,
	0x1e, 0x0c, 0x06, 0x85, 0x36, 0x89, 0xa9, 0xab, 0xa6, 0x22, 0xbf, 0xe1, 0x8f, 0x66, 0xf6, 0x03,
	0x88, 0xd7, 0x55, 0x61, 0x36, 0xfc, 0xbd, 0xcc, 0xdf, 0xe8, 0x14, 0x45, 0xa1, 0x4d, 0x08, 0x91,
	0x88, 0x27, 0xe9, 0x7e, 0x62, 0x91, 0x64, 0x88, 0x93, 0x85, 0x36, 0x37, 0x65, 0x53, 0xb7, 0xca,
	0x52, 0x74, 0x88, 0xc1, 0x4b, 0x97, 0xbd, 0xae, 0x42, 0x2f, 0x82, 0x78, 0xa4, 0x06, 0xe3, 0xe0,
	0x16, 0xc7, 0x7f, 0x38, 0x49, 0x14, 0xef, 0xdc, 0x86, 0x10, 0x41, 0x1c, 0x28, 0x2b, 0xe9, 0x08,
	0x47, 0x5f, 0x7a, 0xf3, 0xc9, 0xee, 0x6e, 0x92, 0x4e, 0xbb, 0xa9, 0x85, 0x36, 0x4f, 0xd6, 0x55,
	0x5d, 0x78, 0xe9, 0x5d, 0xc0, 0x2c, 0x77, 0x3d, 0xce, 0x26, 0x42, 0xbf, 0xd4, 0x05, 0xf7, 0x45,
	0x4e, 0xdb, 0x6e, 0xbd, 0xe6, 0x7e, 0xdf, 0x4a, 0x3a, 0x43, 0x5c, 0x73, 0xb9, 0xe2, 0xfa, 0xa1,
	0x35, 0x1c, 0x8a, 0x08, 0xe2, 0x69, 0x2a, 0xbb, 0x81, 0xf9, 0xbf, 0xaf, 0xb6, 0x98, 0x93, 0x63,
	0xc4, 0x21, 0xa1, 0x31, 0xfa, 0xf7, 0x8f, 0x59, 0x26, 0x77, 0xac, 0x9a, 0xdf, 0xa9, 0x4c, 0x02,
	0xed, 0xa1, 0xb8, 0x5a, 0x3e, 0x4b, 0x2f, 0xdf, 0x75, 0xdf, 0x3b, 0xff, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x48, 0x67, 0x54, 0x5a, 0x4a, 0x01, 0x00, 0x00,
}