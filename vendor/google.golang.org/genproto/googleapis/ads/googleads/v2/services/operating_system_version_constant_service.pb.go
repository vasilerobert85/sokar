// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/operating_system_version_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for
// [OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant][google.ads.googleads.v2.services.OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant].
type GetOperatingSystemVersionConstantRequest struct {
	// Resource name of the OS version to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperatingSystemVersionConstantRequest) Reset() {
	*m = GetOperatingSystemVersionConstantRequest{}
}
func (m *GetOperatingSystemVersionConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperatingSystemVersionConstantRequest) ProtoMessage()    {}
func (*GetOperatingSystemVersionConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02bc5000ca7d978c, []int{0}
}

func (m *GetOperatingSystemVersionConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Unmarshal(m, b)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Merge(m, src)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Size(m)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperatingSystemVersionConstantRequest proto.InternalMessageInfo

func (m *GetOperatingSystemVersionConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetOperatingSystemVersionConstantRequest)(nil), "google.ads.googleads.v2.services.GetOperatingSystemVersionConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/operating_system_version_constant_service.proto", fileDescriptor_02bc5000ca7d978c)
}

var fileDescriptor_02bc5000ca7d978c = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x6a, 0xdb, 0x40,
	0x1c, 0xc7, 0x91, 0x0a, 0x85, 0x8a, 0x76, 0xd1, 0x52, 0xa3, 0x76, 0x70, 0x5d, 0x17, 0x8c, 0x87,
	0x3b, 0x50, 0x29, 0x85, 0x0b, 0x1e, 0xe4, 0x0c, 0x4e, 0x32, 0xc4, 0xc6, 0x06, 0x0d, 0x41, 0x20,
	0x2e, 0xd2, 0x21, 0x04, 0xd6, 0x9d, 0xa2, 0xdf, 0x59, 0x10, 0x42, 0x96, 0xbc, 0x41, 0xc8, 0x1b,
	0x64, 0xcc, 0x13, 0xe4, 0x19, 0xbc, 0xe6, 0x15, 0x32, 0x65, 0xcb, 0x1b, 0x04, 0xf9, 0x74, 0x72,
	0x32, 0xf8, 0xcf, 0xf6, 0xe5, 0xee, 0x7b, 0xdf, 0xcf, 0xef, 0xcf, 0x59, 0x93, 0x44, 0x88, 0x64,
	0xce, 0x30, 0x8d, 0x01, 0x2b, 0x59, 0xa9, 0xd2, 0xc5, 0xc0, 0x8a, 0x32, 0x8d, 0x18, 0x60, 0x91,
	0xb3, 0x82, 0xca, 0x94, 0x27, 0x21, 0x5c, 0x82, 0x64, 0x59, 0x58, 0xb2, 0x02, 0x52, 0xc1, 0xc3,
	0x48, 0x70, 0x90, 0x94, 0xcb, 0xb0, 0xb6, 0xa2, 0xbc, 0x10, 0x52, 0xd8, 0x6d, 0x15, 0x83, 0x68,
	0x0c, 0xa8, 0x49, 0x44, 0xa5, 0x8b, 0x74, 0xa2, 0x73, 0xbc, 0x89, 0x59, 0x30, 0x10, 0x8b, 0x62,
	0x2f, 0xa8, 0x82, 0x39, 0x3f, 0x75, 0x54, 0x9e, 0x62, 0xca, 0xb9, 0x90, 0x54, 0xa6, 0x82, 0x43,
	0x7d, 0xfb, 0xfd, 0xdd, 0x6d, 0x34, 0x4f, 0x99, 0x7e, 0xd6, 0x19, 0x5b, 0xbd, 0x11, 0x93, 0x63,
	0x0d, 0x99, 0xad, 0x18, 0xbe, 0x42, 0x1c, 0xd6, 0x84, 0x29, 0xbb, 0x58, 0x30, 0x90, 0xf6, 0x6f,
	0xeb, 0x9b, 0xae, 0x2b, 0xe4, 0x34, 0x63, 0x2d, 0xa3, 0x6d, 0xf4, 0xbe, 0x4c, 0xbf, 0xea, 0xc3,
	0x53, 0x9a, 0x31, 0xf7, 0xd1, 0xb4, 0xfe, 0x6c, 0x8f, 0x9b, 0xa9, 0xee, 0xed, 0x57, 0xc3, 0xfa,
	0xb5, 0x93, 0x6d, 0x9f, 0xa0, 0x5d, 0x53, 0x44, 0xfb, 0x36, 0xe0, 0x78, 0x1b, 0xb3, 0x9a, 0x79,
	0xa3, 0xed, 0x49, 0x9d, 0xc1, 0xcd, 0xd3, 0xf3, 0x9d, 0xf9, 0xdf, 0xfe, 0x57, 0x6d, 0xe9, 0xea,
	0xc3, 0x38, 0x06, 0x62, 0xeb, 0x53, 0xc0, 0xfd, 0x6b, 0xe7, 0xc7, 0xd2, 0x6b, 0xad, 0xc1, 0xb5,
	0xca, 0x53, 0x40, 0x91, 0xc8, 0x86, 0xb7, 0xa6, 0xd5, 0x8d, 0x44, 0xb6, 0xb3, 0xe1, 0x61, 0x7f,
	0xaf, 0x01, 0x4f, 0xaa, 0x05, 0x4f, 0x8c, 0xb3, 0xa3, 0x3a, 0x2f, 0x11, 0x73, 0xca, 0x13, 0x24,
	0x8a, 0x04, 0x27, 0x8c, 0xaf, 0xd6, 0x8f, 0xd7, 0x15, 0x6c, 0xfe, 0xf7, 0x07, 0x5a, 0xdc, 0x9b,
	0x9f, 0x46, 0x9e, 0xf7, 0x60, 0xb6, 0x47, 0x2a, 0xd0, 0x8b, 0x01, 0x29, 0x59, 0x29, 0xdf, 0x45,
	0x35, 0x18, 0x96, 0xda, 0x12, 0x78, 0x31, 0x04, 0x8d, 0x25, 0xf0, 0xdd, 0x40, 0x5b, 0x5e, 0xcc,
	0xae, 0x3a, 0x27, 0xc4, 0x8b, 0x81, 0x90, 0xc6, 0x44, 0x88, 0xef, 0x12, 0xa2, 0x6d, 0xe7, 0x9f,
	0x57, 0x75, 0xfe, 0x7d, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xec, 0x57, 0x60, 0xed, 0x9e, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperatingSystemVersionConstantServiceClient is the client API for OperatingSystemVersionConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperatingSystemVersionConstantServiceClient interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error)
}

type operatingSystemVersionConstantServiceClient struct {
	cc *grpc.ClientConn
}

func NewOperatingSystemVersionConstantServiceClient(cc *grpc.ClientConn) OperatingSystemVersionConstantServiceClient {
	return &operatingSystemVersionConstantServiceClient{cc}
}

func (c *operatingSystemVersionConstantServiceClient) GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error) {
	out := new(resources.OperatingSystemVersionConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatingSystemVersionConstantServiceServer is the server API for OperatingSystemVersionConstantService service.
type OperatingSystemVersionConstantServiceServer interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(context.Context, *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error)
}

func RegisterOperatingSystemVersionConstantServiceServer(s *grpc.Server, srv OperatingSystemVersionConstantServiceServer) {
	s.RegisterService(&_OperatingSystemVersionConstantService_serviceDesc, srv)
}

func _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatingSystemVersionConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, req.(*GetOperatingSystemVersionConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OperatingSystemVersionConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.OperatingSystemVersionConstantService",
	HandlerType: (*OperatingSystemVersionConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperatingSystemVersionConstant",
			Handler:    _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/operating_system_version_constant_service.proto",
}
