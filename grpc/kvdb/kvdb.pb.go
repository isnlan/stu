// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kvdb.proto

package kvdb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SetKV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetKV) Reset()         { *m = SetKV{} }
func (m *SetKV) String() string { return proto.CompactTextString(m) }
func (*SetKV) ProtoMessage()    {}
func (*SetKV) Descriptor() ([]byte, []int) {
	return fileDescriptor_143269afec78af68, []int{0}
}

func (m *SetKV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetKV.Unmarshal(m, b)
}
func (m *SetKV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetKV.Marshal(b, m, deterministic)
}
func (m *SetKV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetKV.Merge(m, src)
}
func (m *SetKV) XXX_Size() int {
	return xxx_messageInfo_SetKV.Size(m)
}
func (m *SetKV) XXX_DiscardUnknown() {
	xxx_messageInfo_SetKV.DiscardUnknown(m)
}

var xxx_messageInfo_SetKV proto.InternalMessageInfo

func (m *SetKV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetKV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_143269afec78af68, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SetKV)(nil), "SetKV")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("kvdb.proto", fileDescriptor_143269afec78af68) }

var fileDescriptor_143269afec78af68 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x2e, 0x4b, 0x49,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe7, 0x62, 0x0d, 0x4e, 0x2d, 0xf1, 0x0e, 0x13,
	0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85,
	0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0xc0, 0x62, 0x10, 0x8e, 0x12, 0x3b,
	0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x91, 0x2c, 0x17, 0x8b, 0x77, 0x99, 0x8b, 0x93, 0x90,
	0x28, 0x17, 0x73, 0x70, 0x6a, 0x89, 0x10, 0x9b, 0x1e, 0xd8, 0x1c, 0x29, 0x36, 0x3d, 0xb0, 0xb4,
	0x13, 0x7b, 0x14, 0xab, 0x3e, 0xc8, 0x9e, 0x24, 0x36, 0xb0, 0x45, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xca, 0xe7, 0x6e, 0x4c, 0x76, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KvDBClient is the client API for KvDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KvDBClient interface {
	Set(ctx context.Context, in *SetKV, opts ...grpc.CallOption) (*Empty, error)
}

type kvDBClient struct {
	cc *grpc.ClientConn
}

func NewKvDBClient(cc *grpc.ClientConn) KvDBClient {
	return &kvDBClient{cc}
}

func (c *kvDBClient) Set(ctx context.Context, in *SetKV, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/KvDB/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KvDBServer is the server API for KvDB service.
type KvDBServer interface {
	Set(context.Context, *SetKV) (*Empty, error)
}

// UnimplementedKvDBServer can be embedded to have forward compatible implementations.
type UnimplementedKvDBServer struct {
}

func (*UnimplementedKvDBServer) Set(ctx context.Context, req *SetKV) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}

func RegisterKvDBServer(s *grpc.Server, srv KvDBServer) {
	s.RegisterService(&_KvDB_serviceDesc, srv)
}

func _KvDB_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetKV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvDBServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KvDB/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvDBServer).Set(ctx, req.(*SetKV))
	}
	return interceptor(ctx, in, info, handler)
}

var _KvDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "KvDB",
	HandlerType: (*KvDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _KvDB_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kvdb.proto",
}
