// Code generated by protoc-gen-go.
// source: page.proto
// DO NOT EDIT!

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Page struct {
	Uuid             *string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Data             *string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Page) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *Page) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Page)(nil), "press.models.Page")
}

var fileDescriptor1 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x48, 0x4c, 0x4f,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x28, 0x4a, 0x2d, 0x2e, 0xd6, 0xcb, 0xcd,
	0x4f, 0x49, 0xcd, 0x29, 0x56, 0xd2, 0xe3, 0x62, 0x09, 0x00, 0xca, 0x09, 0x09, 0x71, 0xb1, 0x94,
	0x96, 0x66, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x20, 0xb1, 0x94, 0xc4,
	0x92, 0x44, 0x09, 0x26, 0x88, 0x18, 0x88, 0xed, 0xc4, 0x17, 0x05, 0xd1, 0xaf, 0x0f, 0xd1, 0x0f,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xcd, 0x0c, 0x38, 0x5a, 0x00, 0x00, 0x00,
}
