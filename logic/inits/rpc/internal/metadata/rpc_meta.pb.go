// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc_meta.proto

package metadata

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

// Message type.
type RpcMeta_Type int32

const (
	RpcMeta_REQUEST  RpcMeta_Type = 0
	RpcMeta_RESPONSE RpcMeta_Type = 1
)

var RpcMeta_Type_name = map[int32]string{
	0: "REQUEST",
	1: "RESPONSE",
}

var RpcMeta_Type_value = map[string]int32{
	"REQUEST":  0,
	"RESPONSE": 1,
}

func (x RpcMeta_Type) Enum() *RpcMeta_Type {
	p := new(RpcMeta_Type)
	*p = x
	return p
}

func (x RpcMeta_Type) String() string {
	return proto.EnumName(RpcMeta_Type_name, int32(x))
}

func (x *RpcMeta_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcMeta_Type_value, data, "RpcMeta_Type")
	if err != nil {
		return err
	}
	*x = RpcMeta_Type(value)
	return nil
}

func (RpcMeta_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be8c9fa80a702ee0, []int{0, 0}
}

type RpcMeta struct {
	Type *RpcMeta_Type `protobuf:"varint,1,req,name=type,enum=metadata.RpcMeta_Type" json:"type,omitempty"`
	// Message sequence id.
	SequenceId *uint64 `protobuf:"varint,2,req,name=sequence_id,json=sequenceId" json:"sequence_id,omitempty"`
	// Method full name.
	// For example: "test.HelloService.GreetMethod"
	Method *string `protobuf:"bytes,100,opt,name=method" json:"method,omitempty"`
	// Server timeout in milli-seconds.
	ServerTimeout *int64 `protobuf:"varint,101,opt,name=server_timeout,json=serverTimeout" json:"server_timeout,omitempty"`
	// Set as true if the call is failed.
	Failed *bool `protobuf:"varint,200,opt,name=failed" json:"failed,omitempty"`
	// The error code if the call is failed.
	ErrorCode *int32 `protobuf:"varint,201,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	// The error reason if the call is failed.
	Reason               *string  `protobuf:"bytes,202,opt,name=reason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcMeta) Reset()         { *m = RpcMeta{} }
func (m *RpcMeta) String() string { return proto.CompactTextString(m) }
func (*RpcMeta) ProtoMessage()    {}
func (*RpcMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_be8c9fa80a702ee0, []int{0}
}

func (m *RpcMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcMeta.Unmarshal(m, b)
}
func (m *RpcMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcMeta.Marshal(b, m, deterministic)
}
func (m *RpcMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcMeta.Merge(m, src)
}
func (m *RpcMeta) XXX_Size() int {
	return xxx_messageInfo_RpcMeta.Size(m)
}
func (m *RpcMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcMeta.DiscardUnknown(m)
}

var xxx_messageInfo_RpcMeta proto.InternalMessageInfo

func (m *RpcMeta) GetType() RpcMeta_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return RpcMeta_REQUEST
}

func (m *RpcMeta) GetSequenceId() uint64 {
	if m != nil && m.SequenceId != nil {
		return *m.SequenceId
	}
	return 0
}

func (m *RpcMeta) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *RpcMeta) GetServerTimeout() int64 {
	if m != nil && m.ServerTimeout != nil {
		return *m.ServerTimeout
	}
	return 0
}

func (m *RpcMeta) GetFailed() bool {
	if m != nil && m.Failed != nil {
		return *m.Failed
	}
	return false
}

func (m *RpcMeta) GetErrorCode() int32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *RpcMeta) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}

func init() {
	proto.RegisterEnum("metadata.RpcMeta_Type", RpcMeta_Type_name, RpcMeta_Type_value)
	proto.RegisterType((*RpcMeta)(nil), "metadata.RpcMeta")
}

func init() { proto.RegisterFile("rpc_meta.proto", fileDescriptor_be8c9fa80a702ee0) }

var fileDescriptor_be8c9fa80a702ee0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x80, 0x71, 0x08, 0x6d, 0x7a, 0x85, 0xa8, 0xf2, 0x50, 0x3c, 0x81, 0xa9, 0x84, 0x64, 0x31,
	0x64, 0xe0, 0x15, 0x50, 0x06, 0x06, 0xfe, 0x2e, 0x61, 0x8e, 0xac, 0xf8, 0x10, 0x91, 0x48, 0x6d,
	0x1c, 0x17, 0xa9, 0xaf, 0xc0, 0x93, 0x01, 0x4f, 0x85, 0xd2, 0x24, 0x9b, 0xbf, 0xcf, 0xdf, 0xe9,
	0x0e, 0x52, 0xef, 0xea, 0xaa, 0xa5, 0xa0, 0x33, 0xe7, 0x6d, 0xb0, 0x3c, 0xe9, 0xdf, 0x46, 0x07,
	0xbd, 0xf9, 0x8e, 0x60, 0x8e, 0xae, 0x7e, 0xa0, 0xa0, 0xf9, 0x0d, 0xc4, 0x61, 0xef, 0x48, 0x30,
	0x19, 0xa9, 0xf4, 0x76, 0x9d, 0x4d, 0x51, 0x36, 0x06, 0x59, 0xb9, 0x77, 0x84, 0x87, 0x86, 0x5f,
	0xc2, 0xb2, 0xa3, 0xcf, 0x1d, 0x6d, 0x6b, 0xaa, 0x1a, 0x23, 0x22, 0x19, 0xa9, 0x18, 0x61, 0x52,
	0xf7, 0x86, 0xaf, 0x61, 0xd6, 0x52, 0x78, 0xb7, 0x46, 0x18, 0xc9, 0xd4, 0x02, 0x47, 0xe2, 0xd7,
	0x90, 0x76, 0xe4, 0xbf, 0xc8, 0x57, 0xa1, 0x69, 0xc9, 0xee, 0x82, 0x20, 0xc9, 0xd4, 0x31, 0x9e,
	0x0d, 0xb6, 0x1c, 0x24, 0x3f, 0x87, 0xd9, 0x9b, 0x6e, 0x3e, 0xc8, 0x88, 0x1f, 0x26, 0x99, 0x4a,
	0x70, 0x44, 0x7e, 0x01, 0x40, 0xde, 0x5b, 0x5f, 0xd5, 0xd6, 0x90, 0xf8, 0xed, 0x3f, 0x4f, 0x70,
	0x71, 0x50, 0x77, 0xd6, 0x50, 0x3f, 0xe8, 0x49, 0x77, 0x76, 0x2b, 0xfe, 0xd8, 0xb0, 0x78, 0xc0,
	0xcd, 0x15, 0xc4, 0xfd, 0xfd, 0x7c, 0x09, 0x73, 0xcc, 0x5f, 0x5e, 0xf3, 0xa2, 0x5c, 0x1d, 0xf1,
	0x53, 0x48, 0x30, 0x2f, 0x9e, 0x9f, 0x1e, 0x8b, 0x7c, 0xc5, 0xfe, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xa2, 0x48, 0xac, 0x08, 0x27, 0x01, 0x00, 0x00,
}
