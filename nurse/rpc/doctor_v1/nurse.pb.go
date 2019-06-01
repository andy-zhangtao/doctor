// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nurse.proto

package doctor_v1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DoctorRegister struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Hostname             string   `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Os                   string   `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	Cpu                  int32    `protobuf:"varint,5,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory               int32    `protobuf:"varint,6,opt,name=memory,proto3" json:"memory,omitempty"`
	Disk                 int32    `protobuf:"varint,7,opt,name=disk,proto3" json:"disk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoctorRegister) Reset()         { *m = DoctorRegister{} }
func (m *DoctorRegister) String() string { return proto.CompactTextString(m) }
func (*DoctorRegister) ProtoMessage()    {}
func (*DoctorRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa1e12dd1855b39, []int{0}
}

func (m *DoctorRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoctorRegister.Unmarshal(m, b)
}
func (m *DoctorRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoctorRegister.Marshal(b, m, deterministic)
}
func (m *DoctorRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoctorRegister.Merge(m, src)
}
func (m *DoctorRegister) XXX_Size() int {
	return xxx_messageInfo_DoctorRegister.Size(m)
}
func (m *DoctorRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_DoctorRegister.DiscardUnknown(m)
}

var xxx_messageInfo_DoctorRegister proto.InternalMessageInfo

func (m *DoctorRegister) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *DoctorRegister) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DoctorRegister) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *DoctorRegister) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *DoctorRegister) GetCpu() int32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *DoctorRegister) GetMemory() int32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *DoctorRegister) GetDisk() int32 {
	if m != nil {
		return m.Disk
	}
	return 0
}

type Reply struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa1e12dd1855b39, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*DoctorRegister)(nil), "doctor_v1.DoctorRegister")
	proto.RegisterType((*Reply)(nil), "doctor_v1.Reply")
}

func init() { proto.RegisterFile("nurse.proto", fileDescriptor_eaa1e12dd1855b39) }

var fileDescriptor_eaa1e12dd1855b39 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x4d, 0x77, 0x5b, 0x77, 0x47, 0x58, 0x4a, 0x0e, 0x92, 0xf6, 0x54, 0x7a, 0xea, 0xa9,
	0xa0, 0x9e, 0xbc, 0x0b, 0xde, 0x3c, 0xe4, 0x05, 0x44, 0xdb, 0x50, 0x43, 0x4d, 0x27, 0x24, 0xa9,
	0xd0, 0x77, 0xf1, 0x61, 0x25, 0x63, 0x09, 0x7a, 0xfb, 0xfe, 0x2f, 0xfc, 0xc3, 0x4c, 0xe0, 0x66,
	0x59, 0x9d, 0x57, 0xbd, 0x75, 0x18, 0x90, 0x9f, 0x47, 0x1c, 0x02, 0xba, 0xd7, 0xaf, 0xbb, 0xf6,
	0x9b, 0xc1, 0xe5, 0x89, 0x92, 0x54, 0x93, 0xf6, 0x41, 0x39, 0x7e, 0x81, 0x4c, 0x5b, 0xc1, 0x1a,
	0xd6, 0x9d, 0x65, 0xa6, 0x2d, 0x2f, 0xe1, 0x30, 0xab, 0x4d, 0x64, 0x24, 0x22, 0xf2, 0x1a, 0x4e,
	0x1f, 0xe8, 0xc3, 0xf2, 0x66, 0x94, 0x38, 0x90, 0x4e, 0x39, 0xb6, 0xd1, 0x8b, 0xe3, 0x6f, 0x1b,
	0x7d, 0x6c, 0x0f, 0x76, 0x15, 0x79, 0xc3, 0xba, 0x5c, 0x46, 0xe4, 0xb7, 0x50, 0x18, 0x65, 0xd0,
	0x6d, 0xa2, 0x20, 0xb9, 0x27, 0xce, 0xe1, 0x38, 0x6a, 0x3f, 0x8b, 0x6b, 0xb2, 0xc4, 0x6d, 0x05,
	0xb9, 0x54, 0xf6, 0x73, 0x8b, 0x63, 0x8c, 0x9f, 0xf6, 0xad, 0x22, 0xde, 0x3f, 0x03, 0xc8, 0x97,
	0xb4, 0xf4, 0x23, 0x9c, 0x12, 0x57, 0x7d, 0xba, 0xaf, 0xff, 0x7f, 0x5b, 0x5d, 0xfe, 0x79, 0xa2,
	0xc1, 0xed, 0xd5, 0x7b, 0x41, 0x9f, 0xf2, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x89, 0xea,
	0x4c, 0x23, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RNRegisterClient is the client API for RNRegister service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RNRegisterClient interface {
	Register(ctx context.Context, in *DoctorRegister, opts ...grpc.CallOption) (*Reply, error)
}

type rNRegisterClient struct {
	cc *grpc.ClientConn
}

func NewRNRegisterClient(cc *grpc.ClientConn) RNRegisterClient {
	return &rNRegisterClient{cc}
}

func (c *rNRegisterClient) Register(ctx context.Context, in *DoctorRegister, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/doctor_v1.RNRegister/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RNRegisterServer is the server API for RNRegister service.
type RNRegisterServer interface {
	Register(context.Context, *DoctorRegister) (*Reply, error)
}

func RegisterRNRegisterServer(s *grpc.Server, srv RNRegisterServer) {
	s.RegisterService(&_RNRegister_serviceDesc, srv)
}

func _RNRegister_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RNRegisterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor_v1.RNRegister/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RNRegisterServer).Register(ctx, req.(*DoctorRegister))
	}
	return interceptor(ctx, in, info, handler)
}

var _RNRegister_serviceDesc = grpc.ServiceDesc{
	ServiceName: "doctor_v1.RNRegister",
	HandlerType: (*RNRegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RNRegister_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nurse.proto",
}
