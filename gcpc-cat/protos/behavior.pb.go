// Code generated by protoc-gen-go. DO NOT EDIT.
// source: behavior.proto

package protos

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8d850be096d032f, []int{0}
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

type Record struct {
	// 当前微服务名称
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// 区块链网络ID,取值包括"fabric"、"kbaas"
	NetworkId string `protobuf:"bytes,2,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// 区块链网络中的指定链
	Channel string `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	// 合约名称
	Contract string `protobuf:"bytes,4,opt,name=contract,proto3" json:"contract,omitempty"`
	// 代理数字身份
	Proxy string `protobuf:"bytes,5,opt,name=proxy,proto3" json:"proxy,omitempty"`
	// 数字身份唯一标识
	Bduid string `protobuf:"bytes,6,opt,name=bduid,proto3" json:"bduid,omitempty"`
	// 操作名称，特指调用智能合约的方法名称
	// 例如调用 "AuthService.ApplyAuthorization"
	Operation string `protobuf:"bytes,7,opt,name=operation,proto3" json:"operation,omitempty"`
	// 操作的资源列表
	Resources []*Resource `protobuf:"bytes,8,rep,name=resources,proto3" json:"resources,omitempty"`
	// 交易Id
	TxId                 string   `protobuf:"bytes,9,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	SpanId               string   `protobuf:"bytes,10,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	TraceId              string   `protobuf:"bytes,11,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Timestamp            int64    `protobuf:"varint,12,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8d850be096d032f, []int{1}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Record) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *Record) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Record) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *Record) GetProxy() string {
	if m != nil {
		return m.Proxy
	}
	return ""
}

func (m *Record) GetBduid() string {
	if m != nil {
		return m.Bduid
	}
	return ""
}

func (m *Record) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *Record) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Record) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *Record) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

func (m *Record) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Resource struct {
	// 资源ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 资源类型
	// 例如"DataService.DataID", "SchemaService.SchemaID
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8d850be096d032f, []int{2}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "protos.Empty")
	proto.RegisterType((*Record)(nil), "protos.Record")
	proto.RegisterType((*Resource)(nil), "protos.Resource")
}

func init() {
	proto.RegisterFile("behavior.proto", fileDescriptor_b8d850be096d032f)
}

var fileDescriptor_b8d850be096d032f = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0xcb, 0x4e, 0x02, 0x31,
	0x18, 0x85, 0x65, 0x60, 0x6e, 0x3f, 0x4a, 0xcc, 0xaf, 0x86, 0x4a, 0x34, 0x21, 0xac, 0x70, 0x33,
	0x0b, 0xdc, 0xb8, 0x26, 0x71, 0x31, 0x89, 0xab, 0x79, 0x01, 0x53, 0xda, 0x06, 0x1a, 0x64, 0x3a,
	0x69, 0x0b, 0xc2, 0x9b, 0xfa, 0x38, 0xa6, 0x17, 0x60, 0x35, 0xf3, 0x9d, 0xd3, 0xdb, 0xf9, 0x0f,
	0x8c, 0x56, 0x62, 0x43, 0x0f, 0x52, 0xe9, 0xaa, 0xd3, 0xca, 0x2a, 0xcc, 0xfc, 0xc7, 0xcc, 0x72,
	0x48, 0x3f, 0x77, 0x9d, 0x3d, 0xcd, 0xfe, 0x12, 0xc8, 0x1a, 0xc1, 0x94, 0xe6, 0x48, 0x20, 0x37,
	0x42, 0x1f, 0x24, 0x13, 0xa4, 0x37, 0xed, 0xcd, 0xcb, 0xe6, 0x8c, 0xf8, 0x0a, 0xd0, 0x0a, 0xfb,
	0xab, 0xf4, 0xf6, 0x5b, 0x72, 0x92, 0x78, 0xb3, 0x8c, 0x4a, 0xed, 0x37, 0xb2, 0x0d, 0x6d, 0x5b,
	0xf1, 0x43, 0xfa, 0x61, 0x63, 0x44, 0x9c, 0x40, 0xc1, 0x54, 0x6b, 0x35, 0x65, 0x96, 0x0c, 0xbc,
	0x75, 0x61, 0x7c, 0x84, 0xb4, 0xd3, 0xea, 0x78, 0x22, 0xa9, 0x37, 0x02, 0x38, 0x75, 0xc5, 0xf7,
	0x92, 0x93, 0x2c, 0xa8, 0x1e, 0xf0, 0x05, 0x4a, 0xd5, 0x09, 0x4d, 0xad, 0x54, 0x2d, 0xc9, 0xc3,
	0xfd, 0x17, 0x01, 0x2b, 0x28, 0xb5, 0x30, 0x6a, 0xaf, 0x99, 0x30, 0xa4, 0x98, 0xf6, 0xe7, 0xc3,
	0xc5, 0x7d, 0xc8, 0x6b, 0xaa, 0x26, 0x1a, 0xcd, 0x75, 0x09, 0x3e, 0x40, 0x6a, 0x8f, 0x2e, 0x49,
	0xe9, 0x4f, 0x1a, 0xd8, 0x63, 0xcd, 0x71, 0x0c, 0xb9, 0xe9, 0x68, 0xeb, 0x64, 0xf0, 0x72, 0xe6,
	0xb0, 0xe6, 0xf8, 0x0c, 0x85, 0x7b, 0xb0, 0x70, 0xce, 0x30, 0xc4, 0xf3, 0x5c, 0xfb, 0x67, 0x59,
	0xb9, 0x13, 0xc6, 0xd2, 0x5d, 0x47, 0x6e, 0xa7, 0xbd, 0x79, 0xbf, 0xb9, 0x0a, 0xb3, 0x0a, 0x8a,
	0xf3, 0xed, 0x38, 0x82, 0x44, 0xf2, 0x38, 0xd6, 0x44, 0x72, 0x44, 0x18, 0xd8, 0x53, 0x27, 0xe2,
	0x2c, 0xfd, 0xff, 0xe2, 0x03, 0x86, 0xcb, 0xd8, 0xd6, 0x97, 0x5a, 0xe3, 0x5b, 0x28, 0x86, 0x0b,
	0x1c, 0x5d, 0xc3, 0xb8, 0xa2, 0x26, 0x77, 0x67, 0x0e, 0x15, 0xde, 0x2c, 0xc7, 0xf0, 0xa4, 0xf4,
	0xba, 0xda, 0xb2, 0x0d, 0x95, 0x6d, 0xc5, 0xa8, 0x8d, 0x0b, 0x56, 0xa1, 0xee, 0xf7, 0xff, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x69, 0x5c, 0xc3, 0xce, 0x07, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BehaviorLogClient is the client API for BehaviorLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BehaviorLogClient interface {
	Recode(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Empty, error)
}

type behaviorLogClient struct {
	cc grpc.ClientConnInterface
}

func NewBehaviorLogClient(cc grpc.ClientConnInterface) BehaviorLogClient {
	return &behaviorLogClient{cc}
}

func (c *behaviorLogClient) Recode(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.BehaviorLog/Recode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BehaviorLogServer is the server API for BehaviorLog service.
type BehaviorLogServer interface {
	Recode(context.Context, *Record) (*Empty, error)
}

// UnimplementedBehaviorLogServer can be embedded to have forward compatible implementations.
type UnimplementedBehaviorLogServer struct {
}

func (*UnimplementedBehaviorLogServer) Recode(ctx context.Context, req *Record) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recode not implemented")
}

func RegisterBehaviorLogServer(s *grpc.Server, srv BehaviorLogServer) {
	s.RegisterService(&_BehaviorLog_serviceDesc, srv)
}

func _BehaviorLog_Recode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BehaviorLogServer).Recode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BehaviorLog/Recode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BehaviorLogServer).Recode(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

var _BehaviorLog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.BehaviorLog",
	HandlerType: (*BehaviorLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recode",
			Handler:    _BehaviorLog_Recode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "behavior.proto",
}
