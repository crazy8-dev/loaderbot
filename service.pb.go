// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package loaderbot

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

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusResponse struct {
	Busy                 bool     `protobuf:"varint,1,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetBusy() bool {
	if m != nil {
		return m.Busy
	}
	return false
}

type RunConfigRequest struct {
	Config               []byte   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	AttackerName         string   `protobuf:"bytes,2,opt,name=attackerName,proto3" json:"attackerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunConfigRequest) Reset()         { *m = RunConfigRequest{} }
func (m *RunConfigRequest) String() string { return proto.CompactTextString(m) }
func (*RunConfigRequest) ProtoMessage()    {}
func (*RunConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *RunConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunConfigRequest.Unmarshal(m, b)
}
func (m *RunConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunConfigRequest.Marshal(b, m, deterministic)
}
func (m *RunConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunConfigRequest.Merge(m, src)
}
func (m *RunConfigRequest) XXX_Size() int {
	return xxx_messageInfo_RunConfigRequest.Size(m)
}
func (m *RunConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunConfigRequest proto.InternalMessageInfo

func (m *RunConfigRequest) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *RunConfigRequest) GetAttackerName() string {
	if m != nil {
		return m.AttackerName
	}
	return ""
}

type ResultsResponse struct {
	ResultsChunk         []byte   `protobuf:"bytes,1,opt,name=resultsChunk,proto3" json:"resultsChunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResultsResponse) Reset()         { *m = ResultsResponse{} }
func (m *ResultsResponse) String() string { return proto.CompactTextString(m) }
func (*ResultsResponse) ProtoMessage()    {}
func (*ResultsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *ResultsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultsResponse.Unmarshal(m, b)
}
func (m *ResultsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultsResponse.Marshal(b, m, deterministic)
}
func (m *ResultsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultsResponse.Merge(m, src)
}
func (m *ResultsResponse) XXX_Size() int {
	return xxx_messageInfo_ResultsResponse.Size(m)
}
func (m *ResultsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResultsResponse proto.InternalMessageInfo

func (m *ResultsResponse) GetResultsChunk() []byte {
	if m != nil {
		return m.ResultsChunk
	}
	return nil
}

type ShutdownNodeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownNodeRequest) Reset()         { *m = ShutdownNodeRequest{} }
func (m *ShutdownNodeRequest) String() string { return proto.CompactTextString(m) }
func (*ShutdownNodeRequest) ProtoMessage()    {}
func (*ShutdownNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *ShutdownNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownNodeRequest.Unmarshal(m, b)
}
func (m *ShutdownNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownNodeRequest.Marshal(b, m, deterministic)
}
func (m *ShutdownNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownNodeRequest.Merge(m, src)
}
func (m *ShutdownNodeRequest) XXX_Size() int {
	return xxx_messageInfo_ShutdownNodeRequest.Size(m)
}
func (m *ShutdownNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownNodeRequest proto.InternalMessageInfo

type ShutdownNodeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownNodeResponse) Reset()         { *m = ShutdownNodeResponse{} }
func (m *ShutdownNodeResponse) String() string { return proto.CompactTextString(m) }
func (*ShutdownNodeResponse) ProtoMessage()    {}
func (*ShutdownNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *ShutdownNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownNodeResponse.Unmarshal(m, b)
}
func (m *ShutdownNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownNodeResponse.Marshal(b, m, deterministic)
}
func (m *ShutdownNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownNodeResponse.Merge(m, src)
}
func (m *ShutdownNodeResponse) XXX_Size() int {
	return xxx_messageInfo_ShutdownNodeResponse.Size(m)
}
func (m *ShutdownNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownNodeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StatusRequest)(nil), "loaderbot.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "loaderbot.StatusResponse")
	proto.RegisterType((*RunConfigRequest)(nil), "loaderbot.RunConfigRequest")
	proto.RegisterType((*ResultsResponse)(nil), "loaderbot.ResultsResponse")
	proto.RegisterType((*ShutdownNodeRequest)(nil), "loaderbot.ShutdownNodeRequest")
	proto.RegisterType((*ShutdownNodeResponse)(nil), "loaderbot.ShutdownNodeResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4e, 0xc3, 0x30,
	0x10, 0x84, 0x1b, 0x40, 0x11, 0x5d, 0x12, 0x8a, 0x16, 0xa8, 0x42, 0x90, 0xa0, 0xb2, 0x38, 0xf4,
	0x14, 0x21, 0x10, 0x27, 0x0e, 0x48, 0xed, 0x15, 0x45, 0xc2, 0xbd, 0x71, 0xcb, 0xcf, 0x42, 0xab,
	0x16, 0xbb, 0xc4, 0x36, 0x88, 0xb7, 0xe5, 0x51, 0x90, 0x9c, 0x34, 0xc4, 0xfc, 0xdc, 0x92, 0x6f,
	0xc7, 0xb3, 0x9e, 0x31, 0x84, 0x8a, 0xaa, 0xb7, 0x45, 0x41, 0xc9, 0xba, 0x92, 0x5a, 0x62, 0x7f,
	0x25, 0xb3, 0x92, 0xaa, 0x5c, 0x6a, 0x36, 0x80, 0x70, 0xa6, 0x33, 0x6d, 0x14, 0xa7, 0x57, 0x43,
	0x4a, 0xb3, 0x0b, 0xd8, 0xdf, 0x00, 0xb5, 0x96, 0x42, 0x11, 0x22, 0xec, 0xe4, 0x46, 0x7d, 0x44,
	0xde, 0xc8, 0x1b, 0xef, 0x72, 0xfb, 0xcd, 0x52, 0x38, 0xe0, 0x46, 0x4c, 0xa5, 0x78, 0x5a, 0x3c,
	0x37, 0x27, 0x71, 0x08, 0x7e, 0x61, 0x81, 0x55, 0x06, 0xbc, 0xf9, 0x43, 0x06, 0x41, 0xa6, 0x75,
	0x56, 0x2c, 0xa9, 0x4a, 0xb3, 0x17, 0x8a, 0xb6, 0x46, 0xde, 0xb8, 0xcf, 0x1d, 0xc6, 0x6e, 0x60,
	0xc0, 0x49, 0x99, 0x95, 0xfe, 0x5e, 0xcb, 0x20, 0xa8, 0x6a, 0x34, 0x9d, 0x1b, 0xb1, 0x6c, 0x4c,
	0x1d, 0xc6, 0x8e, 0xe1, 0x70, 0x36, 0x37, 0xba, 0x94, 0xef, 0x22, 0x95, 0x25, 0x6d, 0x32, 0x0c,
	0xe1, 0xc8, 0xc5, 0xb5, 0xe5, 0xd5, 0xa7, 0x07, 0xfe, 0xbd, 0x8d, 0x8e, 0x77, 0xe0, 0xd7, 0x31,
	0x31, 0x4a, 0xda, 0x36, 0x12, 0xa7, 0x8a, 0xf8, 0xe4, 0x8f, 0x49, 0xed, 0xc4, 0x7a, 0xf8, 0x00,
	0x41, 0x77, 0x07, 0x9e, 0x75, 0xc5, 0xbf, 0xef, 0x14, 0x9f, 0xff, 0x3b, 0x6f, 0x2d, 0x27, 0xb0,
	0xcd, 0x8d, 0xc0, 0xd3, 0x8e, 0xf2, 0x67, 0xc9, 0x71, 0xdc, 0x1d, 0xba, 0x8d, 0xb1, 0xde, 0xa5,
	0x37, 0x09, 0x1f, 0xf7, 0x92, 0xdb, 0x56, 0x92, 0xfb, 0xf6, 0xc1, 0xaf, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xb3, 0x18, 0x5f, 0xc8, 0x01, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoaderClient is the client API for Loader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoaderClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	ShutdownNode(ctx context.Context, in *ShutdownNodeRequest, opts ...grpc.CallOption) (*ShutdownNodeResponse, error)
	Run(ctx context.Context, in *RunConfigRequest, opts ...grpc.CallOption) (Loader_RunClient, error)
}

type loaderClient struct {
	cc *grpc.ClientConn
}

func NewLoaderClient(cc *grpc.ClientConn) LoaderClient {
	return &loaderClient{cc}
}

func (c *loaderClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/loaderbot.Loader/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) ShutdownNode(ctx context.Context, in *ShutdownNodeRequest, opts ...grpc.CallOption) (*ShutdownNodeResponse, error) {
	out := new(ShutdownNodeResponse)
	err := c.cc.Invoke(ctx, "/loaderbot.Loader/ShutdownNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) Run(ctx context.Context, in *RunConfigRequest, opts ...grpc.CallOption) (Loader_RunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Loader_serviceDesc.Streams[0], "/loaderbot.Loader/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &loaderRunClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Loader_RunClient interface {
	Recv() (*ResultsResponse, error)
	grpc.ClientStream
}

type loaderRunClient struct {
	grpc.ClientStream
}

func (x *loaderRunClient) Recv() (*ResultsResponse, error) {
	m := new(ResultsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoaderServer is the server API for Loader service.
type LoaderServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	ShutdownNode(context.Context, *ShutdownNodeRequest) (*ShutdownNodeResponse, error)
	Run(*RunConfigRequest, Loader_RunServer) error
}

// UnimplementedLoaderServer can be embedded to have forward compatible implementations.
type UnimplementedLoaderServer struct {
}

func (*UnimplementedLoaderServer) Status(ctx context.Context, req *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedLoaderServer) ShutdownNode(ctx context.Context, req *ShutdownNodeRequest) (*ShutdownNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShutdownNode not implemented")
}
func (*UnimplementedLoaderServer) Run(req *RunConfigRequest, srv Loader_RunServer) error {
	return status.Errorf(codes.Unimplemented, "method Run not implemented")
}

func RegisterLoaderServer(s *grpc.Server, srv LoaderServer) {
	s.RegisterService(&_Loader_serviceDesc, srv)
}

func _Loader_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loaderbot.Loader/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_ShutdownNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).ShutdownNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loaderbot.Loader/ShutdownNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).ShutdownNode(ctx, req.(*ShutdownNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RunConfigRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LoaderServer).Run(m, &loaderRunServer{stream})
}

type Loader_RunServer interface {
	Send(*ResultsResponse) error
	grpc.ServerStream
}

type loaderRunServer struct {
	grpc.ServerStream
}

func (x *loaderRunServer) Send(m *ResultsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Loader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loaderbot.Loader",
	HandlerType: (*LoaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Loader_Status_Handler,
		},
		{
			MethodName: "ShutdownNode",
			Handler:    _Loader_ShutdownNode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _Loader_Run_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
