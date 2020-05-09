// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package server_hello_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

type HelloResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloResponse) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

type FeedRequest struct {
	Meat                 string   `protobuf:"bytes,1,opt,name=meat,proto3" json:"meat,omitempty"`
	Weight               int64    `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedRequest) Reset()         { *m = FeedRequest{} }
func (m *FeedRequest) String() string { return proto.CompactTextString(m) }
func (*FeedRequest) ProtoMessage()    {}
func (*FeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{2}
}

func (m *FeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedRequest.Unmarshal(m, b)
}
func (m *FeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedRequest.Marshal(b, m, deterministic)
}
func (m *FeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedRequest.Merge(m, src)
}
func (m *FeedRequest) XXX_Size() int {
	return xxx_messageInfo_FeedRequest.Size(m)
}
func (m *FeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FeedRequest proto.InternalMessageInfo

func (m *FeedRequest) GetMeat() string {
	if m != nil {
		return m.Meat
	}
	return ""
}

func (m *FeedRequest) GetWeight() int64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type FeedResponse struct {
	HowlCount            int64    `protobuf:"varint,1,opt,name=howlCount,proto3" json:"howlCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedResponse) Reset()         { *m = FeedResponse{} }
func (m *FeedResponse) String() string { return proto.CompactTextString(m) }
func (*FeedResponse) ProtoMessage()    {}
func (*FeedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{3}
}

func (m *FeedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedResponse.Unmarshal(m, b)
}
func (m *FeedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedResponse.Marshal(b, m, deterministic)
}
func (m *FeedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedResponse.Merge(m, src)
}
func (m *FeedResponse) XXX_Size() int {
	return xxx_messageInfo_FeedResponse.Size(m)
}
func (m *FeedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FeedResponse proto.InternalMessageInfo

func (m *FeedResponse) GetHowlCount() int64 {
	if m != nil {
		return m.HowlCount
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "server_hello_proto.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "server_hello_proto.HelloResponse")
	proto.RegisterType((*FeedRequest)(nil), "server_hello_proto.FeedRequest")
	proto.RegisterType((*FeedResponse)(nil), "server_hello_proto.FeedResponse")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x8a,
	0x07, 0x8b, 0xc5, 0x83, 0xc5, 0x94, 0x4c, 0xb8, 0x78, 0x3c, 0x40, 0xdc, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x30, 0x5b, 0x48, 0x80, 0x8b, 0x39, 0x31, 0x3d, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83,
	0x39, 0x08, 0xc4, 0x54, 0x32, 0xe5, 0xe2, 0x85, 0xea, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x25,
	0x52, 0x9b, 0x25, 0x17, 0xb7, 0x5b, 0x6a, 0x6a, 0x0a, 0x92, 0x5d, 0xb9, 0xa9, 0x89, 0x25, 0x30,
	0x4d, 0x20, 0xb6, 0x90, 0x18, 0x17, 0x5b, 0x79, 0x6a, 0x66, 0x7a, 0x46, 0x09, 0x54, 0x1f, 0x94,
	0xa7, 0xa4, 0xc3, 0xc5, 0x03, 0xd1, 0x0a, 0xb5, 0x50, 0x86, 0x8b, 0x33, 0x23, 0xbf, 0x3c, 0xc7,
	0x39, 0xbf, 0x34, 0x0f, 0x62, 0x00, 0x73, 0x10, 0x42, 0xc0, 0xa8, 0x99, 0x89, 0x8b, 0x27, 0x24,
	0x33, 0x3d, 0xb5, 0x28, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x90, 0x8b, 0x0b, 0xec,
	0x60, 0xb0, 0xa0, 0x90, 0x82, 0x1e, 0x66, 0x48, 0xe8, 0x21, 0x07, 0x83, 0x94, 0x22, 0x1e, 0x15,
	0x50, 0x17, 0xf8, 0x71, 0x71, 0x82, 0x5c, 0x04, 0x31, 0x51, 0x1e, 0x9b, 0x7a, 0x24, 0xbf, 0x4a,
	0x29, 0xe0, 0x56, 0x00, 0x35, 0x2f, 0x80, 0x8b, 0xdd, 0x39, 0x23, 0x31, 0x2f, 0x2f, 0x35, 0x87,
	0x0a, 0xa6, 0x69, 0x30, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0xc5, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xa5, 0x91, 0x48, 0xe0, 0x05, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TigerServiceClient is the client API for TigerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TigerServiceClient interface {
	HelloTiger(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	FeedTiger(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error)
	Channel(ctx context.Context, opts ...grpc.CallOption) (TigerService_ChannelClient, error)
}

type tigerServiceClient struct {
	cc *grpc.ClientConn
}

func NewTigerServiceClient(cc *grpc.ClientConn) TigerServiceClient {
	return &tigerServiceClient{cc}
}

func (c *tigerServiceClient) HelloTiger(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/server_hello_proto.TigerService/HelloTiger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tigerServiceClient) FeedTiger(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	out := new(FeedResponse)
	err := c.cc.Invoke(ctx, "/server_hello_proto.TigerService/FeedTiger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tigerServiceClient) Channel(ctx context.Context, opts ...grpc.CallOption) (TigerService_ChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TigerService_serviceDesc.Streams[0], "/server_hello_proto.TigerService/Channel", opts...)
	if err != nil {
		return nil, err
	}
	x := &tigerServiceChannelClient{stream}
	return x, nil
}

type TigerService_ChannelClient interface {
	Send(*FeedRequest) error
	Recv() (*FeedResponse, error)
	grpc.ClientStream
}

type tigerServiceChannelClient struct {
	grpc.ClientStream
}

func (x *tigerServiceChannelClient) Send(m *FeedRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tigerServiceChannelClient) Recv() (*FeedResponse, error) {
	m := new(FeedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TigerServiceServer is the server API for TigerService service.
type TigerServiceServer interface {
	HelloTiger(context.Context, *HelloRequest) (*HelloResponse, error)
	FeedTiger(context.Context, *FeedRequest) (*FeedResponse, error)
	Channel(TigerService_ChannelServer) error
}

func RegisterTigerServiceServer(s *grpc.Server, srv TigerServiceServer) {
	s.RegisterService(&_TigerService_serviceDesc, srv)
}

func _TigerService_HelloTiger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TigerServiceServer).HelloTiger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server_hello_proto.TigerService/HelloTiger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TigerServiceServer).HelloTiger(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TigerService_FeedTiger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TigerServiceServer).FeedTiger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server_hello_proto.TigerService/FeedTiger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TigerServiceServer).FeedTiger(ctx, req.(*FeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TigerService_Channel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TigerServiceServer).Channel(&tigerServiceChannelServer{stream})
}

type TigerService_ChannelServer interface {
	Send(*FeedResponse) error
	Recv() (*FeedRequest, error)
	grpc.ServerStream
}

type tigerServiceChannelServer struct {
	grpc.ServerStream
}

func (x *tigerServiceChannelServer) Send(m *FeedResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tigerServiceChannelServer) Recv() (*FeedRequest, error) {
	m := new(FeedRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TigerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server_hello_proto.TigerService",
	HandlerType: (*TigerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloTiger",
			Handler:    _TigerService_HelloTiger_Handler,
		},
		{
			MethodName: "FeedTiger",
			Handler:    _TigerService_FeedTiger_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Channel",
			Handler:       _TigerService_Channel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}