// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package wire

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InternalTransaction struct {
	Index                uint64   `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Amount               uint64   `protobuf:"varint,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Receiver             string   `protobuf:"bytes,3,opt,name=Receiver,proto3" json:"Receiver,omitempty"`
	UntilBlock           uint64   `protobuf:"varint,4,opt,name=UntilBlock,proto3" json:"UntilBlock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalTransaction) Reset()         { *m = InternalTransaction{} }
func (m *InternalTransaction) String() string { return proto.CompactTextString(m) }
func (*InternalTransaction) ProtoMessage()    {}
func (*InternalTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *InternalTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalTransaction.Unmarshal(m, b)
}
func (m *InternalTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalTransaction.Marshal(b, m, deterministic)
}
func (m *InternalTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalTransaction.Merge(m, src)
}
func (m *InternalTransaction) XXX_Size() int {
	return xxx_messageInfo_InternalTransaction.Size(m)
}
func (m *InternalTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_InternalTransaction proto.InternalMessageInfo

func (m *InternalTransaction) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *InternalTransaction) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *InternalTransaction) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *InternalTransaction) GetUntilBlock() uint64 {
	if m != nil {
		return m.UntilBlock
	}
	return 0
}

type Event struct {
	Index                uint64                 `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Creator              string                 `protobuf:"bytes,2,opt,name=Creator,proto3" json:"Creator,omitempty"`
	Parents              [][]byte               `protobuf:"bytes,3,rep,name=Parents,proto3" json:"Parents,omitempty"`
	LamportTime          uint64                 `protobuf:"varint,4,opt,name=LamportTime,proto3" json:"LamportTime,omitempty"`
	InternalTransactions []*InternalTransaction `protobuf:"bytes,5,rep,name=InternalTransactions,proto3" json:"InternalTransactions,omitempty"`
	ExternalTransactions [][]byte               `protobuf:"bytes,6,rep,name=ExternalTransactions,proto3" json:"ExternalTransactions,omitempty"`
	Sign                 string                 `protobuf:"bytes,7,opt,name=Sign,proto3" json:"Sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Event) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Event) GetParents() [][]byte {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *Event) GetLamportTime() uint64 {
	if m != nil {
		return m.LamportTime
	}
	return 0
}

func (m *Event) GetInternalTransactions() []*InternalTransaction {
	if m != nil {
		return m.InternalTransactions
	}
	return nil
}

func (m *Event) GetExternalTransactions() [][]byte {
	if m != nil {
		return m.ExternalTransactions
	}
	return nil
}

func (m *Event) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func init() {
	proto.RegisterType((*InternalTransaction)(nil), "wire.InternalTransaction")
	proto.RegisterType((*Event)(nil), "wire.Event")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0xe6, 0x4f, 0xe9, 0x85, 0xe9, 0x88, 0x90, 0x61, 0x40, 0x56, 0xa7, 0x4c, 0x19,
	0xca, 0x13, 0x00, 0xea, 0x50, 0x09, 0x24, 0x64, 0xca, 0x03, 0x98, 0x70, 0x42, 0x16, 0xc9, 0xb9,
	0x72, 0x4c, 0xe9, 0xc6, 0xce, 0x53, 0xa3, 0x98, 0x16, 0x75, 0x30, 0x9b, 0x7f, 0xdf, 0xf9, 0xfc,
	0xdd, 0xe7, 0x83, 0x92, 0xb6, 0xc4, 0xbe, 0xd9, 0x38, 0xeb, 0x2d, 0x66, 0x9f, 0xc6, 0xd1, 0xfc,
	0x0b, 0xce, 0x56, 0xec, 0xc9, 0xb1, 0xee, 0xd6, 0x4e, 0xf3, 0xa0, 0x5b, 0x6f, 0x2c, 0x63, 0x05,
	0xf9, 0x8a, 0x5f, 0x69, 0x27, 0x12, 0x99, 0xd4, 0x99, 0xfa, 0x05, 0x3c, 0x87, 0xe2, 0xa6, 0xb7,
	0x1f, 0xec, 0xc5, 0x24, 0xc8, 0x7b, 0xc2, 0x4b, 0x38, 0x51, 0xd4, 0x92, 0xd9, 0x92, 0x13, 0xa9,
	0x4c, 0xea, 0x99, 0xfa, 0x63, 0xbc, 0x02, 0x78, 0x66, 0x6f, 0xba, 0xdb, 0xce, 0xb6, 0xef, 0x22,
	0x0b, 0x7d, 0x47, 0xca, 0xfc, 0x7b, 0x02, 0xf9, 0x72, 0x1c, 0xeb, 0x1f, 0x4f, 0x01, 0xd3, 0x3b,
	0x47, 0xda, 0x5b, 0x17, 0x4c, 0x67, 0xea, 0x80, 0x63, 0xe5, 0x51, 0x3b, 0x62, 0x3f, 0x88, 0x54,
	0xa6, 0xf5, 0xa9, 0x3a, 0x20, 0x4a, 0x28, 0xef, 0x75, 0xbf, 0xb1, 0xce, 0xaf, 0x4d, 0x4f, 0x7b,
	0xd3, 0x63, 0x09, 0x1f, 0xa0, 0x8a, 0xc4, 0x1e, 0x44, 0x2e, 0xd3, 0xba, 0x5c, 0x5c, 0x34, 0xe3,
	0xdf, 0x34, 0x91, 0x1b, 0x2a, 0xda, 0x86, 0x0b, 0xa8, 0x96, 0xbb, 0xc8, 0x73, 0x45, 0x98, 0x2b,
	0x5a, 0x43, 0x84, 0xec, 0xc9, 0xbc, 0xb1, 0x98, 0x86, 0x54, 0xe1, 0xfc, 0x52, 0x84, 0xd5, 0x5c,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x8e, 0x2c, 0x01, 0xa9, 0x01, 0x00, 0x00,
}