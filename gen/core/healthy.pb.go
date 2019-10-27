// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/healthy.proto

package healthv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CheckResponse_ServingStatus int32

const (
	CheckResponse_UNKNOWN         CheckResponse_ServingStatus = 0
	CheckResponse_SERVING         CheckResponse_ServingStatus = 1
	CheckResponse_NOT_SERVING     CheckResponse_ServingStatus = 2
	CheckResponse_SERVICE_UNKNOWN CheckResponse_ServingStatus = 3
)

var CheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
	3: "SERVICE_UNKNOWN",
}

var CheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":         0,
	"SERVING":         1,
	"NOT_SERVING":     2,
	"SERVICE_UNKNOWN": 3,
}

func (x CheckResponse_ServingStatus) String() string {
	return proto.EnumName(CheckResponse_ServingStatus_name, int32(x))
}

func (CheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9ddb1deb3b492237, []int{1, 0}
}

type CheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ddb1deb3b492237, []int{0}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type CheckResponse struct {
	Status               CheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=grpc.health.v1.CheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ddb1deb3b492237, []int{1}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetStatus() CheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return CheckResponse_UNKNOWN
}

func init() {
	proto.RegisterEnum("grpc.health.v1.CheckResponse_ServingStatus", CheckResponse_ServingStatus_name, CheckResponse_ServingStatus_value)
	proto.RegisterType((*CheckRequest)(nil), "grpc.health.v1.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "grpc.health.v1.CheckResponse")
}

func init() { proto.RegisterFile("core/healthy.proto", fileDescriptor_9ddb1deb3b492237) }

var fileDescriptor_9ddb1deb3b492237 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0x2f, 0x4a,
	0xd5, 0xcf, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0xa8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x4b, 0x2f, 0x2a, 0x48, 0xd6, 0x83, 0x88, 0xe9, 0x95, 0x19, 0x4a, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7,
	0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6,
	0xe7, 0x15, 0x43, 0x54, 0x4b, 0xe9, 0x80, 0xa9, 0x64, 0xdd, 0xf4, 0xd4, 0x3c, 0xdd, 0xe2, 0xf2,
	0xc4, 0xf4, 0xf4, 0xd4, 0x22, 0xfd, 0xfc, 0x02, 0xb0, 0x0a, 0x4c, 0xd5, 0x4a, 0x1a, 0x5c, 0x3c,
	0xce, 0x19, 0xa9, 0xc9, 0xd9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12, 0x5c, 0xec,
	0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae,
	0xd2, 0x52, 0x46, 0x2e, 0x5e, 0xa8, 0xd2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x67, 0x2e,
	0xb6, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0xb0, 0x52, 0x3e, 0x23, 0x6d, 0x3d, 0x54, 0x87, 0xea,
	0xa1, 0x28, 0xd7, 0x0b, 0x06, 0x19, 0x94, 0x97, 0x1e, 0x0c, 0xd6, 0x12, 0x04, 0xd5, 0xaa, 0xe4,
	0xcf, 0xc5, 0x8b, 0x22, 0x21, 0xc4, 0xcd, 0xc5, 0x1e, 0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x27,
	0xc0, 0x00, 0xe2, 0x04, 0xbb, 0x06, 0x85, 0x79, 0xfa, 0xb9, 0x0b, 0x30, 0x0a, 0xf1, 0x73, 0x71,
	0xfb, 0xf9, 0x87, 0xc4, 0xc3, 0x04, 0x98, 0x84, 0x84, 0xb9, 0xf8, 0xc1, 0x1c, 0x67, 0xd7, 0x78,
	0x98, 0x16, 0x66, 0xa3, 0xed, 0x8c, 0x5c, 0x9c, 0x1e, 0x60, 0x27, 0x38, 0x06, 0x78, 0x0a, 0x25,
	0x71, 0xb1, 0x82, 0x5d, 0x21, 0x24, 0x83, 0xc3, 0x71, 0x60, 0x6f, 0x4b, 0xc9, 0xe2, 0x75, 0xba,
	0x92, 0x4c, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xc4, 0x84, 0x44, 0xf4, 0xcb, 0x0c, 0xf5, 0x91, 0x62,
	0x28, 0x19, 0x6c, 0xb4, 0x1b, 0x17, 0x6b, 0x78, 0x62, 0x49, 0x72, 0x06, 0x45, 0x76, 0x18, 0x30,
	0x3a, 0x45, 0x80, 0x62, 0x3f, 0x17, 0x4d, 0x95, 0x13, 0x0f, 0xc4, 0x33, 0x95, 0x01, 0xa0, 0xf8,
	0x0a, 0x60, 0x8c, 0xe2, 0x80, 0x48, 0x95, 0x19, 0x2e, 0x62, 0x62, 0x76, 0xf7, 0x88, 0x58, 0xc5,
	0xc4, 0xe7, 0x0e, 0xd2, 0x00, 0x51, 0xa5, 0x17, 0x66, 0x78, 0x8a, 0x49, 0x00, 0x24, 0x10, 0x13,
	0x03, 0x11, 0x89, 0x89, 0x09, 0x33, 0x4c, 0x62, 0x03, 0x47, 0xb6, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x7c, 0x35, 0xfe, 0x15, 0x5e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthAPIClient is the client API for HealthAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthAPIClient interface {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// Performs a watch for the serving status of the requested service.
	// The server will immediately send back a message indicating the current
	// serving status.  It will then subsequently send a new message whenever
	// the service's serving status changes.
	//
	// If the requested service is unknown when the call is received, the
	// server will send a message setting the serving status to
	// SERVICE_UNKNOWN but will *not* terminate the call.  If at some
	// future point, the serving status of the service becomes known, the
	// server will send a new message with the service's serving status.
	//
	// If the call terminates with status UNIMPLEMENTED, then clients
	// should assume this method is not supported and should not retry the
	// call.  If the call terminates with any other status (including OK),
	// clients should retry the call with appropriate exponential backoff.
	Watch(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (HealthAPI_WatchClient, error)
}

type healthAPIClient struct {
	cc *grpc.ClientConn
}

func NewHealthAPIClient(cc *grpc.ClientConn) HealthAPIClient {
	return &healthAPIClient{cc}
}

func (c *healthAPIClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/grpc.health.v1.HealthAPI/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAPIClient) Watch(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (HealthAPI_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HealthAPI_serviceDesc.Streams[0], "/grpc.health.v1.HealthAPI/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthAPIWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HealthAPI_WatchClient interface {
	Recv() (*CheckResponse, error)
	grpc.ClientStream
}

type healthAPIWatchClient struct {
	grpc.ClientStream
}

func (x *healthAPIWatchClient) Recv() (*CheckResponse, error) {
	m := new(CheckResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HealthAPIServer is the server API for HealthAPI service.
type HealthAPIServer interface {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// Performs a watch for the serving status of the requested service.
	// The server will immediately send back a message indicating the current
	// serving status.  It will then subsequently send a new message whenever
	// the service's serving status changes.
	//
	// If the requested service is unknown when the call is received, the
	// server will send a message setting the serving status to
	// SERVICE_UNKNOWN but will *not* terminate the call.  If at some
	// future point, the serving status of the service becomes known, the
	// server will send a new message with the service's serving status.
	//
	// If the call terminates with status UNIMPLEMENTED, then clients
	// should assume this method is not supported and should not retry the
	// call.  If the call terminates with any other status (including OK),
	// clients should retry the call with appropriate exponential backoff.
	Watch(*CheckRequest, HealthAPI_WatchServer) error
}

// UnimplementedHealthAPIServer can be embedded to have forward compatible implementations.
type UnimplementedHealthAPIServer struct {
}

func (*UnimplementedHealthAPIServer) Check(ctx context.Context, req *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedHealthAPIServer) Watch(req *CheckRequest, srv HealthAPI_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterHealthAPIServer(s *grpc.Server, srv HealthAPIServer) {
	s.RegisterService(&_HealthAPI_serviceDesc, srv)
}

func _HealthAPI_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAPIServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.health.v1.HealthAPI/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAPIServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAPI_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CheckRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HealthAPIServer).Watch(m, &healthAPIWatchServer{stream})
}

type HealthAPI_WatchServer interface {
	Send(*CheckResponse) error
	grpc.ServerStream
}

type healthAPIWatchServer struct {
	grpc.ServerStream
}

func (x *healthAPIWatchServer) Send(m *CheckResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _HealthAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.health.v1.HealthAPI",
	HandlerType: (*HealthAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _HealthAPI_Check_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _HealthAPI_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "core/healthy.proto",
}
