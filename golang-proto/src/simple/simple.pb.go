// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple/simple.proto

package personpb

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

type Person struct {
	Age                  int32    `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	ProfilePicture       []byte   `protobuf:"bytes,4,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	IsVerfied            bool     `protobuf:"varint,5,opt,name=isVerfied,proto3" json:"isVerfied,omitempty"`
	Height               float32  `protobuf:"fixed32,6,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3c868e94d57426, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Person) GetProfilePicture() []byte {
	if m != nil {
		return m.ProfilePicture
	}
	return nil
}

func (m *Person) GetIsVerfied() bool {
	if m != nil {
		return m.IsVerfied
	}
	return false
}

func (m *Person) GetHeight() float32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*Person)(nil), "example.person.Person")
}

func init() { proto.RegisterFile("simple/simple.proto", fileDescriptor_9b3c868e94d57426) }

var fileDescriptor_9b3c868e94d57426 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xb1, 0x4a, 0xc6, 0x30,
	0x14, 0x85, 0xc9, 0x5f, 0x1b, 0x9a, 0x8b, 0x54, 0x89, 0x20, 0x01, 0x15, 0x82, 0x8b, 0x99, 0x74,
	0xf0, 0x0d, 0x7c, 0x00, 0x29, 0x19, 0x1c, 0x5c, 0x4a, 0xaa, 0xb7, 0x6d, 0xa0, 0x6d, 0x42, 0x12,
	0xc1, 0x17, 0xf3, 0xfd, 0xa4, 0x49, 0xc5, 0xe9, 0x9c, 0xf3, 0x7d, 0xd3, 0x81, 0xab, 0x68, 0x57,
	0xbf, 0xe0, 0x53, 0x89, 0x47, 0x1f, 0x5c, 0x72, 0xbc, 0xc5, 0x6f, 0x53, 0x26, 0x86, 0xe8, 0xb6,
	0xfb, 0x1f, 0x02, 0xb4, 0xcb, 0x95, 0x5f, 0x42, 0x65, 0x26, 0x14, 0x44, 0x12, 0x55, 0xeb, 0xbd,
	0xf2, 0x3b, 0x80, 0xd1, 0x86, 0x98, 0xfa, 0xcd, 0xac, 0x28, 0x4e, 0x92, 0x28, 0xa6, 0x59, 0x26,
	0xaf, 0x66, 0x45, 0x7e, 0x03, 0x6c, 0x31, 0x7f, 0xb6, 0xca, 0xb6, 0xd9, 0x41, 0x96, 0x0f, 0x70,
	0xe1, 0x83, 0x1b, 0xed, 0x82, 0xbd, 0xb7, 0x1f, 0xe9, 0x2b, 0xa0, 0x38, 0x93, 0x44, 0x9d, 0xeb,
	0xf6, 0xc0, 0x5d, 0xa1, 0xfc, 0x16, 0x98, 0x8d, 0x6f, 0x18, 0x46, 0x8b, 0x9f, 0xa2, 0x96, 0x44,
	0x35, 0xfa, 0x1f, 0xf0, 0x6b, 0xa0, 0x33, 0xda, 0x69, 0x4e, 0x82, 0x4a, 0xa2, 0x4e, 0xfa, 0x58,
	0x2f, 0xf0, 0xde, 0x94, 0x07, 0x7e, 0x18, 0x68, 0xbe, 0xf6, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x74, 0x36, 0xf5, 0x36, 0xf1, 0x00, 0x00, 0x00,
}