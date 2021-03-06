// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/carrier_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for [CarrierConstantService.GetCarrierConstant][google.ads.googleads.v1.services.CarrierConstantService.GetCarrierConstant].
type GetCarrierConstantRequest struct {
	// Resource name of the carrier constant to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCarrierConstantRequest) Reset()         { *m = GetCarrierConstantRequest{} }
func (m *GetCarrierConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetCarrierConstantRequest) ProtoMessage()    {}
func (*GetCarrierConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad8e2a96eee0407f, []int{0}
}

func (m *GetCarrierConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCarrierConstantRequest.Unmarshal(m, b)
}
func (m *GetCarrierConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCarrierConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetCarrierConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCarrierConstantRequest.Merge(m, src)
}
func (m *GetCarrierConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetCarrierConstantRequest.Size(m)
}
func (m *GetCarrierConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCarrierConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCarrierConstantRequest proto.InternalMessageInfo

func (m *GetCarrierConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCarrierConstantRequest)(nil), "google.ads.googleads.v1.services.GetCarrierConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/carrier_constant_service.proto", fileDescriptor_ad8e2a96eee0407f)
}

var fileDescriptor_ad8e2a96eee0407f = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x26, 0xb9, 0x70, 0xe1, 0x86, 0xeb, 0x26, 0x0b, 0xad, 0xa9, 0x8b, 0x52, 0x8b, 0x14, 0x17,
	0x33, 0xa4, 0x6e, 0x64, 0x8a, 0x68, 0xda, 0x45, 0x5d, 0x49, 0xa9, 0xd0, 0x85, 0x04, 0xca, 0x98,
	0x0c, 0x21, 0xd0, 0xcc, 0xd4, 0x39, 0x69, 0x37, 0xe2, 0xa6, 0xaf, 0xe0, 0x1b, 0xb8, 0x74, 0xef,
	0x4b, 0x74, 0xeb, 0x2b, 0xb8, 0x12, 0x7c, 0x07, 0x49, 0x27, 0x93, 0xda, 0xda, 0xd0, 0xdd, 0xc7,
	0x9c, 0xef, 0xe7, 0x9c, 0x2f, 0xb1, 0x2e, 0x23, 0x21, 0xa2, 0x31, 0xc3, 0x34, 0x04, 0xac, 0x60,
	0x86, 0x66, 0x2e, 0x06, 0x26, 0x67, 0x71, 0xc0, 0x00, 0x07, 0x54, 0xca, 0x98, 0xc9, 0x51, 0x20,
	0x38, 0xa4, 0x94, 0xa7, 0xa3, 0x7c, 0x82, 0x26, 0x52, 0xa4, 0xc2, 0xae, 0x29, 0x15, 0xa2, 0x21,
	0xa0, 0xc2, 0x00, 0xcd, 0x5c, 0xa4, 0x0d, 0x9c, 0xf3, 0xb2, 0x08, 0xc9, 0x40, 0x4c, 0xe5, 0xb6,
	0x0c, 0xe5, 0xed, 0x1c, 0x69, 0xe5, 0x24, 0xc6, 0x94, 0x73, 0x91, 0xd2, 0x34, 0x16, 0x1c, 0xf2,
	0xe9, 0xc1, 0x8f, 0x69, 0x30, 0x8e, 0x99, 0x96, 0xd5, 0xaf, 0xac, 0xc3, 0x1e, 0x4b, 0xbb, 0xca,
	0xb3, 0x9b, 0x5b, 0x0e, 0xd8, 0xc3, 0x94, 0x41, 0x6a, 0x1f, 0x5b, 0x7b, 0x3a, 0x77, 0xc4, 0x69,
	0xc2, 0x2a, 0x46, 0xcd, 0x68, 0xfe, 0x1b, 0xfc, 0xd7, 0x8f, 0x37, 0x34, 0x61, 0xad, 0x2f, 0xc3,
	0xda, 0xdf, 0xd0, 0xdf, 0xaa, 0x73, 0xec, 0x37, 0xc3, 0xb2, 0x7f, 0xbb, 0xdb, 0x6d, 0xb4, 0xab,
	0x07, 0x54, 0xba, 0x93, 0xd3, 0x2a, 0x15, 0x17, 0x15, 0xa1, 0x0d, 0x69, 0x1d, 0xcd, 0xdf, 0x3f,
	0x9e, 0xcd, 0xa6, 0x7d, 0x92, 0x35, 0xf9, 0xb8, 0x76, 0xd2, 0x45, 0xb0, 0xce, 0x05, 0x7c, 0xfa,
	0xe4, 0x54, 0x17, 0x5e, 0x65, 0x65, 0x9d, 0xa3, 0x49, 0x0c, 0x28, 0x10, 0x49, 0x67, 0x6e, 0x5a,
	0x8d, 0x40, 0x24, 0x3b, 0x6f, 0xe8, 0x54, 0xb7, 0xb7, 0xd2, 0xcf, 0x7a, 0xef, 0x1b, 0x77, 0xd7,
	0xb9, 0x41, 0x24, 0xc6, 0x94, 0x47, 0x48, 0xc8, 0x08, 0x47, 0x8c, 0x2f, 0xbf, 0x0a, 0x5e, 0x45,
	0x96, 0xff, 0x6c, 0x6d, 0x0d, 0x5e, 0xcc, 0x3f, 0x3d, 0xcf, 0x7b, 0x35, 0x6b, 0x3d, 0x65, 0xe8,
	0x85, 0x80, 0x14, 0xcc, 0xd0, 0xd0, 0x45, 0x79, 0x30, 0x2c, 0x34, 0xc5, 0xf7, 0x42, 0xf0, 0x0b,
	0x8a, 0x3f, 0x74, 0x7d, 0x4d, 0xf9, 0x34, 0x1b, 0xea, 0x9d, 0x10, 0x2f, 0x04, 0x42, 0x0a, 0x12,
	0x21, 0x43, 0x97, 0x10, 0x4d, 0xbb, 0xff, 0xbb, 0xdc, 0xf3, 0xec, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x13, 0xa6, 0x17, 0x6a, 0x13, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CarrierConstantServiceClient is the client API for CarrierConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CarrierConstantServiceClient interface {
	// Returns the requested carrier constant in full detail.
	GetCarrierConstant(ctx context.Context, in *GetCarrierConstantRequest, opts ...grpc.CallOption) (*resources.CarrierConstant, error)
}

type carrierConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarrierConstantServiceClient(cc grpc.ClientConnInterface) CarrierConstantServiceClient {
	return &carrierConstantServiceClient{cc}
}

func (c *carrierConstantServiceClient) GetCarrierConstant(ctx context.Context, in *GetCarrierConstantRequest, opts ...grpc.CallOption) (*resources.CarrierConstant, error) {
	out := new(resources.CarrierConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CarrierConstantService/GetCarrierConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarrierConstantServiceServer is the server API for CarrierConstantService service.
type CarrierConstantServiceServer interface {
	// Returns the requested carrier constant in full detail.
	GetCarrierConstant(context.Context, *GetCarrierConstantRequest) (*resources.CarrierConstant, error)
}

// UnimplementedCarrierConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCarrierConstantServiceServer struct {
}

func (*UnimplementedCarrierConstantServiceServer) GetCarrierConstant(ctx context.Context, req *GetCarrierConstantRequest) (*resources.CarrierConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarrierConstant not implemented")
}

func RegisterCarrierConstantServiceServer(s *grpc.Server, srv CarrierConstantServiceServer) {
	s.RegisterService(&_CarrierConstantService_serviceDesc, srv)
}

func _CarrierConstantService_GetCarrierConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarrierConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierConstantServiceServer).GetCarrierConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CarrierConstantService/GetCarrierConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierConstantServiceServer).GetCarrierConstant(ctx, req.(*GetCarrierConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CarrierConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CarrierConstantService",
	HandlerType: (*CarrierConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCarrierConstant",
			Handler:    _CarrierConstantService_GetCarrierConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/carrier_constant_service.proto",
}
