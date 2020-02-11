// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/label_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [LabelService.GetLabel][google.ads.googleads.v2.services.LabelService.GetLabel].
type GetLabelRequest struct {
	// The resource name of the label to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLabelRequest) Reset()         { *m = GetLabelRequest{} }
func (m *GetLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetLabelRequest) ProtoMessage()    {}
func (*GetLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa07dafc79d19310, []int{0}
}

func (m *GetLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLabelRequest.Unmarshal(m, b)
}
func (m *GetLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLabelRequest.Marshal(b, m, deterministic)
}
func (m *GetLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLabelRequest.Merge(m, src)
}
func (m *GetLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetLabelRequest.Size(m)
}
func (m *GetLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLabelRequest proto.InternalMessageInfo

func (m *GetLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [LabelService.MutateLabels][google.ads.googleads.v2.services.LabelService.MutateLabels].
type MutateLabelsRequest struct {
	// ID of the customer whose labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on labels.
	Operations []*LabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateLabelsRequest) Reset()         { *m = MutateLabelsRequest{} }
func (m *MutateLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateLabelsRequest) ProtoMessage()    {}
func (*MutateLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa07dafc79d19310, []int{1}
}

func (m *MutateLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateLabelsRequest.Unmarshal(m, b)
}
func (m *MutateLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateLabelsRequest.Marshal(b, m, deterministic)
}
func (m *MutateLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateLabelsRequest.Merge(m, src)
}
func (m *MutateLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateLabelsRequest.Size(m)
}
func (m *MutateLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateLabelsRequest proto.InternalMessageInfo

func (m *MutateLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateLabelsRequest) GetOperations() []*LabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove, update) on a label.
type LabelOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*LabelOperation_Create
	//	*LabelOperation_Update
	//	*LabelOperation_Remove
	Operation            isLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *LabelOperation) Reset()         { *m = LabelOperation{} }
func (m *LabelOperation) String() string { return proto.CompactTextString(m) }
func (*LabelOperation) ProtoMessage()    {}
func (*LabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa07dafc79d19310, []int{2}
}

func (m *LabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelOperation.Unmarshal(m, b)
}
func (m *LabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelOperation.Marshal(b, m, deterministic)
}
func (m *LabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelOperation.Merge(m, src)
}
func (m *LabelOperation) XXX_Size() int {
	return xxx_messageInfo_LabelOperation.Size(m)
}
func (m *LabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_LabelOperation proto.InternalMessageInfo

func (m *LabelOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isLabelOperation_Operation interface {
	isLabelOperation_Operation()
}

type LabelOperation_Create struct {
	Create *resources.Label `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type LabelOperation_Update struct {
	Update *resources.Label `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type LabelOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*LabelOperation_Create) isLabelOperation_Operation() {}

func (*LabelOperation_Update) isLabelOperation_Operation() {}

func (*LabelOperation_Remove) isLabelOperation_Operation() {}

func (m *LabelOperation) GetOperation() isLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *LabelOperation) GetCreate() *resources.Label {
	if x, ok := m.GetOperation().(*LabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *LabelOperation) GetUpdate() *resources.Label {
	if x, ok := m.GetOperation().(*LabelOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *LabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*LabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LabelOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LabelOperation_Create)(nil),
		(*LabelOperation_Update)(nil),
		(*LabelOperation_Remove)(nil),
	}
}

// Response message for a labels mutate.
type MutateLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MutateLabelsResponse) Reset()         { *m = MutateLabelsResponse{} }
func (m *MutateLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateLabelsResponse) ProtoMessage()    {}
func (*MutateLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa07dafc79d19310, []int{3}
}

func (m *MutateLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateLabelsResponse.Unmarshal(m, b)
}
func (m *MutateLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateLabelsResponse.Marshal(b, m, deterministic)
}
func (m *MutateLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateLabelsResponse.Merge(m, src)
}
func (m *MutateLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateLabelsResponse.Size(m)
}
func (m *MutateLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateLabelsResponse proto.InternalMessageInfo

func (m *MutateLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateLabelsResponse) GetResults() []*MutateLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for a label mutate.
type MutateLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateLabelResult) Reset()         { *m = MutateLabelResult{} }
func (m *MutateLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateLabelResult) ProtoMessage()    {}
func (*MutateLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa07dafc79d19310, []int{4}
}

func (m *MutateLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateLabelResult.Unmarshal(m, b)
}
func (m *MutateLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateLabelResult.Marshal(b, m, deterministic)
}
func (m *MutateLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateLabelResult.Merge(m, src)
}
func (m *MutateLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateLabelResult.Size(m)
}
func (m *MutateLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateLabelResult proto.InternalMessageInfo

func (m *MutateLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetLabelRequest)(nil), "google.ads.googleads.v2.services.GetLabelRequest")
	proto.RegisterType((*MutateLabelsRequest)(nil), "google.ads.googleads.v2.services.MutateLabelsRequest")
	proto.RegisterType((*LabelOperation)(nil), "google.ads.googleads.v2.services.LabelOperation")
	proto.RegisterType((*MutateLabelsResponse)(nil), "google.ads.googleads.v2.services.MutateLabelsResponse")
	proto.RegisterType((*MutateLabelResult)(nil), "google.ads.googleads.v2.services.MutateLabelResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/label_service.proto", fileDescriptor_fa07dafc79d19310)
}

var fileDescriptor_fa07dafc79d19310 = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0x36, 0x5b, 0xa9, 0xed, 0x64, 0x6d, 0xe9, 0x54, 0x31, 0xac, 0x82, 0x4b, 0x2c, 0xb8, 0xac,
	0x98, 0xd4, 0x54, 0x8b, 0xa4, 0xf4, 0xb0, 0x0b, 0xb6, 0x15, 0xac, 0x2d, 0x29, 0xf4, 0x20, 0x0b,
	0x61, 0x9a, 0x4c, 0x97, 0xd0, 0x24, 0x13, 0x67, 0x26, 0x0b, 0xa5, 0xf4, 0xe2, 0x5f, 0xf0, 0xe4,
	0xd5, 0xa3, 0xe0, 0xd1, 0x3f, 0xd1, 0xab, 0xe0, 0x2f, 0xf0, 0xe4, 0x2f, 0xf0, 0x20, 0x28, 0x33,
	0x93, 0xd9, 0xee, 0x56, 0xca, 0xb6, 0xb7, 0x97, 0x37, 0xdf, 0xf7, 0xbd, 0x6f, 0xde, 0xcb, 0x1b,
	0xf0, 0xbc, 0x4f, 0x48, 0x3f, 0xc5, 0x2e, 0x8a, 0x99, 0xab, 0x42, 0x11, 0x0d, 0x3c, 0x97, 0x61,
	0x3a, 0x48, 0x22, 0xcc, 0xdc, 0x14, 0x1d, 0xe0, 0x34, 0xac, 0x3e, 0x9d, 0x82, 0x12, 0x4e, 0x60,
	0x53, 0x41, 0x1d, 0x14, 0x33, 0x67, 0xc8, 0x72, 0x06, 0x9e, 0xa3, 0x59, 0x8d, 0xa7, 0x97, 0xe9,
	0x52, 0xcc, 0x48, 0x49, 0x87, 0xc2, 0x4a, 0xb0, 0xf1, 0x40, 0xc3, 0x8b, 0xc4, 0x45, 0x79, 0x4e,
	0x38, 0xe2, 0x09, 0xc9, 0x59, 0x75, 0x5a, 0x95, 0x73, 0xe5, 0xd7, 0x41, 0x79, 0xe8, 0x1e, 0x26,
	0x38, 0x8d, 0xc3, 0x0c, 0xb1, 0xa3, 0x0a, 0x71, 0xaf, 0x42, 0xd0, 0x22, 0x72, 0x19, 0x47, 0xbc,
	0x64, 0x17, 0x0e, 0x84, 0x70, 0x94, 0x26, 0x38, 0xe7, 0xea, 0xc0, 0x5e, 0x05, 0xf3, 0x9b, 0x98,
	0xbf, 0x11, 0x1e, 0x02, 0xfc, 0xbe, 0xc4, 0x8c, 0xc3, 0x47, 0xe0, 0xb6, 0x76, 0x17, 0xe6, 0x28,
	0xc3, 0x96, 0xd1, 0x34, 0x5a, 0xb3, 0x41, 0x5d, 0x27, 0xdf, 0xa2, 0x0c, 0xdb, 0x3f, 0x0c, 0xb0,
	0xb8, 0x5d, 0x72, 0xc4, 0xb1, 0xe4, 0x32, 0x4d, 0x7e, 0x08, 0xcc, 0xa8, 0x64, 0x9c, 0x64, 0x98,
	0x86, 0x49, 0x5c, 0x51, 0x81, 0x4e, 0xbd, 0x8e, 0xe1, 0x2e, 0x00, 0xa4, 0xc0, 0x54, 0x5d, 0xcc,
	0xaa, 0x35, 0xa7, 0x5a, 0xa6, 0xb7, 0xec, 0x4c, 0x6a, 0xa4, 0x23, 0xab, 0xec, 0x68, 0x62, 0x30,
	0xa2, 0x01, 0x1f, 0x83, 0xf9, 0x02, 0x51, 0x9e, 0xa0, 0x34, 0x3c, 0x44, 0x49, 0x5a, 0x52, 0x6c,
	0x4d, 0x35, 0x8d, 0xd6, 0x4c, 0x30, 0x57, 0xa5, 0x37, 0x54, 0x56, 0x5c, 0x6c, 0x80, 0xd2, 0x24,
	0x46, 0x1c, 0x87, 0x24, 0x4f, 0x8f, 0xad, 0x9b, 0x12, 0x56, 0xd7, 0xc9, 0x9d, 0x3c, 0x3d, 0xb6,
	0xff, 0x18, 0x60, 0x6e, 0xbc, 0x18, 0x5c, 0x03, 0x66, 0x59, 0x48, 0x96, 0x68, 0xb5, 0x64, 0x99,
	0x5e, 0x43, 0x7b, 0xd6, 0xd3, 0x70, 0x36, 0xc4, 0x34, 0xb6, 0x11, 0x3b, 0x0a, 0x80, 0x82, 0x8b,
	0x18, 0x76, 0xc1, 0x74, 0x44, 0x31, 0xe2, 0xaa, 0x8d, 0xa6, 0xd7, 0xba, 0xf4, 0xae, 0xc3, 0x5f,
	0x42, 0x5d, 0x76, 0xeb, 0x46, 0x50, 0x31, 0x85, 0x86, 0x52, 0xb4, 0x6a, 0xd7, 0xd7, 0x50, 0x4c,
	0x68, 0x81, 0x69, 0x8a, 0x33, 0x32, 0x50, 0xcd, 0x99, 0x15, 0x27, 0xea, 0xbb, 0x6b, 0x82, 0xd9,
	0x61, 0x37, 0xed, 0xaf, 0x06, 0xb8, 0x33, 0x3e, 0x57, 0x56, 0x90, 0x9c, 0x61, 0xb8, 0x01, 0xee,
	0x5e, 0xe8, 0x72, 0x88, 0x29, 0x25, 0x54, 0xca, 0x99, 0x1e, 0xd4, 0x96, 0x68, 0x11, 0x39, 0x7b,
	0xf2, 0xd7, 0x0b, 0x16, 0xc7, 0xfb, 0xff, 0x4a, 0xc0, 0xe1, 0x36, 0xb8, 0x45, 0x31, 0x2b, 0x53,
	0xae, 0x87, 0xbf, 0x32, 0x79, 0xf8, 0x23, 0x86, 0x02, 0xc9, 0x0d, 0xb4, 0x86, 0xfd, 0x12, 0x2c,
	0xfc, 0x77, 0x7a, 0xa5, 0x3f, 0xd8, 0xfb, 0x5b, 0x03, 0x75, 0x49, 0xda, 0x53, 0x65, 0xe0, 0x27,
	0x03, 0xcc, 0xe8, 0x5d, 0x80, 0xcf, 0x26, 0xbb, 0xba, 0xb0, 0x37, 0x8d, 0x2b, 0x4f, 0xc5, 0x5e,
	0xfe, 0xf0, 0xfd, 0xe7, 0xc7, 0x5a, 0x1b, 0xb6, 0xc4, 0x4b, 0x70, 0x32, 0x66, 0x75, 0x5d, 0xaf,
	0x0a, 0x73, 0xdb, 0xea, 0x69, 0x60, 0x6e, 0xfb, 0x14, 0x7e, 0x33, 0x40, 0x7d, 0x74, 0x2c, 0xf0,
	0xc5, 0xb5, 0xba, 0xa6, 0xd7, 0xb3, 0xb1, 0x7a, 0x5d, 0x9a, 0x9a, 0xbe, 0xbd, 0x2a, 0x1d, 0x2f,
	0xdb, 0x4f, 0x84, 0xe3, 0x73, 0x8b, 0x27, 0x23, 0xbb, 0xbe, 0xde, 0x3e, 0xad, 0x0c, 0xfb, 0x99,
	0x94, 0xf0, 0x8d, 0x76, 0xe3, 0xfe, 0x59, 0xc7, 0x3a, 0x2f, 0x53, 0x45, 0x45, 0xc2, 0x9c, 0x88,
	0x64, 0xdd, 0xdf, 0x06, 0x58, 0x8a, 0x48, 0x36, 0xd1, 0x52, 0x77, 0x61, 0x74, 0x4e, 0xbb, 0x62,
	0xdf, 0x76, 0x8d, 0x77, 0x5b, 0x15, 0xad, 0x4f, 0x52, 0x94, 0xf7, 0x1d, 0x42, 0xfb, 0x6e, 0x1f,
	0xe7, 0x72, 0x1b, 0xdd, 0xf3, 0x42, 0x97, 0xbf, 0xe8, 0x6b, 0x3a, 0xf8, 0x5c, 0x9b, 0xda, 0xec,
	0x74, 0xbe, 0xd4, 0x9a, 0x9b, 0x4a, 0xb0, 0x13, 0x33, 0x47, 0x85, 0x22, 0xda, 0xf7, 0x9c, 0xaa,
	0x30, 0x3b, 0xd3, 0x90, 0x5e, 0x27, 0x66, 0xbd, 0x21, 0xa4, 0xb7, 0xef, 0xf5, 0x34, 0xe4, 0x57,
	0x6d, 0x49, 0xe5, 0x7d, 0xbf, 0x13, 0x33, 0xdf, 0x1f, 0x82, 0x7c, 0x7f, 0xdf, 0xf3, 0x7d, 0x0d,
	0x3b, 0x98, 0x96, 0x3e, 0x57, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x64, 0xc1, 0x87, 0x25, 0x78,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LabelServiceClient is the client API for LabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LabelServiceClient interface {
	// Returns the requested label in full detail.
	GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*resources.Label, error)
	// Creates, updates, or removes labels. Operation statuses are returned.
	MutateLabels(ctx context.Context, in *MutateLabelsRequest, opts ...grpc.CallOption) (*MutateLabelsResponse, error)
}

type labelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLabelServiceClient(cc grpc.ClientConnInterface) LabelServiceClient {
	return &labelServiceClient{cc}
}

func (c *labelServiceClient) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*resources.Label, error) {
	out := new(resources.Label)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.LabelService/GetLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) MutateLabels(ctx context.Context, in *MutateLabelsRequest, opts ...grpc.CallOption) (*MutateLabelsResponse, error) {
	out := new(MutateLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.LabelService/MutateLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabelServiceServer is the server API for LabelService service.
type LabelServiceServer interface {
	// Returns the requested label in full detail.
	GetLabel(context.Context, *GetLabelRequest) (*resources.Label, error)
	// Creates, updates, or removes labels. Operation statuses are returned.
	MutateLabels(context.Context, *MutateLabelsRequest) (*MutateLabelsResponse, error)
}

// UnimplementedLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLabelServiceServer struct {
}

func (*UnimplementedLabelServiceServer) GetLabel(ctx context.Context, req *GetLabelRequest) (*resources.Label, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetLabel not implemented")
}
func (*UnimplementedLabelServiceServer) MutateLabels(ctx context.Context, req *MutateLabelsRequest) (*MutateLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateLabels not implemented")
}

func RegisterLabelServiceServer(s *grpc.Server, srv LabelServiceServer) {
	s.RegisterService(&_LabelService_serviceDesc, srv)
}

func _LabelService_GetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).GetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.LabelService/GetLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).GetLabel(ctx, req.(*GetLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_MutateLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).MutateLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.LabelService/MutateLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).MutateLabels(ctx, req.(*MutateLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.LabelService",
	HandlerType: (*LabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLabel",
			Handler:    _LabelService_GetLabel_Handler,
		},
		{
			MethodName: "MutateLabels",
			Handler:    _LabelService_MutateLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/label_service.proto",
}
