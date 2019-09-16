// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api_options.proto

package native_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type APISourceType int32

const (
	APISourceType_SOURCE_BOTH   APISourceType = 0
	APISourceType_SOURCE_SERVER APISourceType = 1
	APISourceType_SOURCE_CLIENT APISourceType = 2
)

var APISourceType_name = map[int32]string{
	0: "SOURCE_BOTH",
	1: "SOURCE_SERVER",
	2: "SOURCE_CLIENT",
}

var APISourceType_value = map[string]int32{
	"SOURCE_BOTH":   0,
	"SOURCE_SERVER": 1,
	"SOURCE_CLIENT": 2,
}

func (x APISourceType) Enum() *APISourceType {
	p := new(APISourceType)
	*p = x
	return p
}

func (x APISourceType) String() string {
	return proto.EnumName(APISourceType_name, int32(x))
}

func (x *APISourceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(APISourceType_value, data, "APISourceType")
	if err != nil {
		return err
	}
	*x = APISourceType(value)
	return nil
}

func (APISourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_62b81fd83060afa0, []int{0}
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b81fd83060afa0, []int{0}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

var E_NeedsSetupConnection = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1038,
	Name:          "native_api.needs_setup_connection",
	Tag:           "varint,1038,opt,name=needs_setup_connection,def=1",
	Filename:      "api_options.proto",
}

var E_NeedsAuthentication = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1039,
	Name:          "native_api.needs_authentication",
	Tag:           "varint,1039,opt,name=needs_authentication,def=1",
	Filename:      "api_options.proto",
}

var E_Id = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         1036,
	Name:          "native_api.id",
	Tag:           "varint,1036,opt,name=id,def=0",
	Filename:      "api_options.proto",
}

var E_Source = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*APISourceType)(nil),
	Field:         1037,
	Name:          "native_api.source",
	Tag:           "varint,1037,opt,name=source,enum=native_api.APISourceType,def=0",
	Filename:      "api_options.proto",
}

var E_Ifdef = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1038,
	Name:          "native_api.ifdef",
	Tag:           "bytes,1038,opt,name=ifdef",
	Filename:      "api_options.proto",
}

var E_Log = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1039,
	Name:          "native_api.log",
	Tag:           "varint,1039,opt,name=log,def=1",
	Filename:      "api_options.proto",
}

var E_NoDelay = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1040,
	Name:          "native_api.no_delay",
	Tag:           "varint,1040,opt,name=no_delay,def=0",
	Filename:      "api_options.proto",
}

func init() {
	proto.RegisterEnum("native_api.APISourceType", APISourceType_name, APISourceType_value)
	proto.RegisterType((*Void)(nil), "native_api.void")
	proto.RegisterExtension(E_NeedsSetupConnection)
	proto.RegisterExtension(E_NeedsAuthentication)
	proto.RegisterExtension(E_Id)
	proto.RegisterExtension(E_Source)
	proto.RegisterExtension(E_Ifdef)
	proto.RegisterExtension(E_Log)
	proto.RegisterExtension(E_NoDelay)
}

func init() { proto.RegisterFile("api_options.proto", fileDescriptor_62b81fd83060afa0) }

var fileDescriptor_62b81fd83060afa0 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x5f, 0x4b, 0xeb, 0x30,
	0x18, 0xc6, 0xcf, 0x76, 0xb6, 0x9d, 0x9d, 0xc8, 0x74, 0xab, 0x22, 0xea, 0x85, 0x0e, 0xaf, 0xc4,
	0x8b, 0x4e, 0x14, 0x44, 0x72, 0xb7, 0xcd, 0x8a, 0x03, 0x75, 0x92, 0x4e, 0x11, 0xbc, 0xa8, 0xb1,
	0x79, 0xdb, 0x05, 0x4a, 0x12, 0x9a, 0x74, 0xb0, 0xef, 0xe0, 0xbf, 0x8f, 0x2c, 0x6d, 0x37, 0xdd,
	0x40, 0xa8, 0xb7, 0x4f, 0xdf, 0xdf, 0xaf, 0x0f, 0x4f, 0x50, 0x8b, 0x2a, 0xee, 0x49, 0x65, 0xb8,
	0x14, 0xda, 0x56, 0xb1, 0x34, 0xd2, 0x42, 0x82, 0x1a, 0x3e, 0x01, 0x8f, 0x2a, 0xbe, 0xd3, 0x0e,
	0xa5, 0x0c, 0x23, 0xe8, 0x64, 0x5f, 0x9e, 0x93, 0xa0, 0xc3, 0x40, 0xfb, 0x31, 0x57, 0x46, 0xc6,
	0xf9, 0xf5, 0x7e, 0x0d, 0x55, 0x26, 0x92, 0xb3, 0xc3, 0x0b, 0xd4, 0xe8, 0xde, 0x0e, 0x5c, 0x99,
	0xc4, 0x3e, 0x8c, 0xa6, 0x0a, 0xac, 0x35, 0xb4, 0xe2, 0x0e, 0xef, 0x48, 0xdf, 0xf1, 0x7a, 0xc3,
	0xd1, 0x65, 0xf3, 0x8f, 0xd5, 0x42, 0x8d, 0x59, 0xe0, 0x3a, 0xe4, 0xde, 0x21, 0xcd, 0xd2, 0x42,
	0xd4, 0xbf, 0x1a, 0x38, 0x37, 0xa3, 0x66, 0x19, 0x3f, 0xa2, 0x4d, 0x01, 0xc0, 0xb4, 0xa7, 0xc1,
	0x24, 0xca, 0xf3, 0xa5, 0x10, 0xe0, 0xa7, 0xf5, 0xac, 0x5d, 0x3b, 0x2f, 0x63, 0xcf, 0xcb, 0xd8,
	0xd7, 0x60, 0xc6, 0x92, 0x0d, 0xf3, 0xf6, 0x5b, 0x6f, 0xf5, 0x76, 0xe9, 0xa0, 0x8e, 0x2b, 0x26,
	0x4e, 0x80, 0x6c, 0x64, 0x12, 0x37, 0x75, 0xf4, 0xbf, 0x14, 0xf8, 0x01, 0xe5, 0xb9, 0x47, 0x13,
	0x33, 0x06, 0x61, 0xb8, 0x4f, 0x7f, 0xa5, 0x7e, 0x5f, 0x54, 0xaf, 0x67, 0x8a, 0xee, 0x92, 0x01,
	0x9f, 0xa0, 0x32, 0x67, 0xd6, 0xde, 0x0f, 0x1e, 0xad, 0x69, 0x08, 0x73, 0xd1, 0x4b, 0x2a, 0x6a,
	0xe0, 0xd2, 0x11, 0x29, 0x73, 0x86, 0x9f, 0x50, 0x4d, 0x67, 0x83, 0x15, 0x83, 0xaf, 0x29, 0xb8,
	0x7a, 0xbc, 0x6d, 0x7f, 0x3f, 0x8e, 0xbd, 0xb4, 0x37, 0x5e, 0x5c, 0x9b, 0xcc, 0xbc, 0xf8, 0x14,
	0x55, 0x79, 0xc0, 0x20, 0x28, 0xfe, 0x41, 0xb6, 0xde, 0x7f, 0x92, 0x9f, 0xe3, 0x33, 0xf4, 0x37,
	0x92, 0x61, 0x31, 0xb5, 0x34, 0x4c, 0x8a, 0xe0, 0x1e, 0xaa, 0x0b, 0xe9, 0x31, 0x88, 0xe8, 0xb4,
	0x18, 0xff, 0xc8, 0xf1, 0x6a, 0x40, 0x23, 0x0d, 0xe4, 0x9f, 0x90, 0xe7, 0x29, 0xf7, 0x19, 0x00,
	0x00, 0xff, 0xff, 0x63, 0x3d, 0x1c, 0x98, 0x95, 0x02, 0x00, 0x00,
}