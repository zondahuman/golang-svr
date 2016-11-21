// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Test
*/
package pb

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

type Test struct {
	Label string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	Type  int32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Reps  string `protobuf:"bytes,3,opt,name=reps" json:"reps,omitempty"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Test) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Test) GetReps() string {
	if m != nil {
		return m.Reps
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "pb.Test")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe1, 0x62, 0x09, 0x49,
	0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0xcd, 0x49, 0x4c, 0x4a, 0xcd, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0x70, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0x14, 0x18,
	0x35, 0x58, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x51, 0x6a, 0x41, 0xb1, 0x04, 0x33, 0x58, 0x21, 0x98,
	0x9d, 0xc4, 0x06, 0x36, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xe1, 0x6a, 0xd8, 0x5e,
	0x00, 0x00, 0x00,
}
