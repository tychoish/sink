// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_criterion_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for [CampaignCriterionService.GetCampaignCriterion][google.ads.googleads.v1.services.CampaignCriterionService.GetCampaignCriterion].
type GetCampaignCriterionRequest struct {
	// The resource name of the criterion to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignCriterionRequest) Reset()         { *m = GetCampaignCriterionRequest{} }
func (m *GetCampaignCriterionRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignCriterionRequest) ProtoMessage()    {}
func (*GetCampaignCriterionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a3c2b9bfd5dd6f9, []int{0}
}

func (m *GetCampaignCriterionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignCriterionRequest.Unmarshal(m, b)
}
func (m *GetCampaignCriterionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignCriterionRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignCriterionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignCriterionRequest.Merge(m, src)
}
func (m *GetCampaignCriterionRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignCriterionRequest.Size(m)
}
func (m *GetCampaignCriterionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignCriterionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignCriterionRequest proto.InternalMessageInfo

func (m *GetCampaignCriterionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignCriterionService.MutateCampaignCriteria][google.ads.googleads.v1.services.CampaignCriterionService.MutateCampaignCriteria].
type MutateCampaignCriteriaRequest struct {
	// The ID of the customer whose criteria are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual criteria.
	Operations []*CampaignCriterionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateCampaignCriteriaRequest) Reset()         { *m = MutateCampaignCriteriaRequest{} }
func (m *MutateCampaignCriteriaRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignCriteriaRequest) ProtoMessage()    {}
func (*MutateCampaignCriteriaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a3c2b9bfd5dd6f9, []int{1}
}

func (m *MutateCampaignCriteriaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignCriteriaRequest.Unmarshal(m, b)
}
func (m *MutateCampaignCriteriaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignCriteriaRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignCriteriaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignCriteriaRequest.Merge(m, src)
}
func (m *MutateCampaignCriteriaRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignCriteriaRequest.Size(m)
}
func (m *MutateCampaignCriteriaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignCriteriaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignCriteriaRequest proto.InternalMessageInfo

func (m *MutateCampaignCriteriaRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignCriteriaRequest) GetOperations() []*CampaignCriterionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignCriteriaRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignCriteriaRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a campaign criterion.
type CampaignCriterionOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignCriterionOperation_Create
	//	*CampaignCriterionOperation_Update
	//	*CampaignCriterionOperation_Remove
	Operation            isCampaignCriterionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *CampaignCriterionOperation) Reset()         { *m = CampaignCriterionOperation{} }
func (m *CampaignCriterionOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterionOperation) ProtoMessage()    {}
func (*CampaignCriterionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a3c2b9bfd5dd6f9, []int{2}
}

func (m *CampaignCriterionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterionOperation.Unmarshal(m, b)
}
func (m *CampaignCriterionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterionOperation.Marshal(b, m, deterministic)
}
func (m *CampaignCriterionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterionOperation.Merge(m, src)
}
func (m *CampaignCriterionOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterionOperation.Size(m)
}
func (m *CampaignCriterionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterionOperation proto.InternalMessageInfo

func (m *CampaignCriterionOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignCriterionOperation_Operation interface {
	isCampaignCriterionOperation_Operation()
}

type CampaignCriterionOperation_Create struct {
	Create *resources.CampaignCriterion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignCriterionOperation_Update struct {
	Update *resources.CampaignCriterion `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignCriterionOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignCriterionOperation_Create) isCampaignCriterionOperation_Operation() {}

func (*CampaignCriterionOperation_Update) isCampaignCriterionOperation_Operation() {}

func (*CampaignCriterionOperation_Remove) isCampaignCriterionOperation_Operation() {}

func (m *CampaignCriterionOperation) GetOperation() isCampaignCriterionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignCriterionOperation) GetCreate() *resources.CampaignCriterion {
	if x, ok := m.GetOperation().(*CampaignCriterionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignCriterionOperation) GetUpdate() *resources.CampaignCriterion {
	if x, ok := m.GetOperation().(*CampaignCriterionOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignCriterionOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignCriterionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterionOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterionOperation_Create)(nil),
		(*CampaignCriterionOperation_Update)(nil),
		(*CampaignCriterionOperation_Remove)(nil),
	}
}

// Response message for campaign criterion mutate.
type MutateCampaignCriteriaResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignCriterionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *MutateCampaignCriteriaResponse) Reset()         { *m = MutateCampaignCriteriaResponse{} }
func (m *MutateCampaignCriteriaResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignCriteriaResponse) ProtoMessage()    {}
func (*MutateCampaignCriteriaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a3c2b9bfd5dd6f9, []int{3}
}

func (m *MutateCampaignCriteriaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignCriteriaResponse.Unmarshal(m, b)
}
func (m *MutateCampaignCriteriaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignCriteriaResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignCriteriaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignCriteriaResponse.Merge(m, src)
}
func (m *MutateCampaignCriteriaResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignCriteriaResponse.Size(m)
}
func (m *MutateCampaignCriteriaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignCriteriaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignCriteriaResponse proto.InternalMessageInfo

func (m *MutateCampaignCriteriaResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignCriteriaResponse) GetResults() []*MutateCampaignCriterionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateCampaignCriterionResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignCriterionResult) Reset()         { *m = MutateCampaignCriterionResult{} }
func (m *MutateCampaignCriterionResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignCriterionResult) ProtoMessage()    {}
func (*MutateCampaignCriterionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a3c2b9bfd5dd6f9, []int{4}
}

func (m *MutateCampaignCriterionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignCriterionResult.Unmarshal(m, b)
}
func (m *MutateCampaignCriterionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignCriterionResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignCriterionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignCriterionResult.Merge(m, src)
}
func (m *MutateCampaignCriterionResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignCriterionResult.Size(m)
}
func (m *MutateCampaignCriterionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignCriterionResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignCriterionResult proto.InternalMessageInfo

func (m *MutateCampaignCriterionResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignCriterionRequest)(nil), "google.ads.googleads.v1.services.GetCampaignCriterionRequest")
	proto.RegisterType((*MutateCampaignCriteriaRequest)(nil), "google.ads.googleads.v1.services.MutateCampaignCriteriaRequest")
	proto.RegisterType((*CampaignCriterionOperation)(nil), "google.ads.googleads.v1.services.CampaignCriterionOperation")
	proto.RegisterType((*MutateCampaignCriteriaResponse)(nil), "google.ads.googleads.v1.services.MutateCampaignCriteriaResponse")
	proto.RegisterType((*MutateCampaignCriterionResult)(nil), "google.ads.googleads.v1.services.MutateCampaignCriterionResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_criterion_service.proto", fileDescriptor_3a3c2b9bfd5dd6f9)
}

var fileDescriptor_3a3c2b9bfd5dd6f9 = []byte{
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcb, 0x6e, 0xd3, 0x4a,
	0x18, 0xc7, 0x8f, 0x9d, 0xa3, 0x9e, 0xd3, 0x49, 0xcf, 0x41, 0x1a, 0x6e, 0x56, 0x4a, 0x21, 0x32,
	0x95, 0xa8, 0xb2, 0xb0, 0x95, 0x50, 0x36, 0x2e, 0x15, 0x4d, 0x0a, 0x6d, 0x59, 0xf4, 0x22, 0x57,
	0xaa, 0x04, 0x8a, 0x64, 0x4d, 0xed, 0xa9, 0x65, 0xd5, 0xf6, 0x98, 0x99, 0x71, 0xa4, 0xaa, 0xea,
	0x06, 0xf1, 0x06, 0xbc, 0x00, 0x62, 0xc9, 0x6b, 0xb0, 0x40, 0xea, 0x96, 0x35, 0x0b, 0x24, 0x56,
	0x88, 0x87, 0x40, 0xe3, 0xf1, 0xa4, 0x97, 0xc4, 0x04, 0x95, 0xdd, 0x97, 0x99, 0xbf, 0x7f, 0xdf,
	0x75, 0xbe, 0x80, 0x6e, 0x48, 0x48, 0x18, 0x63, 0x1b, 0x05, 0xcc, 0x96, 0xa6, 0xb0, 0x06, 0x6d,
	0x9b, 0x61, 0x3a, 0x88, 0x7c, 0xcc, 0x6c, 0x1f, 0x25, 0x19, 0x8a, 0xc2, 0xd4, 0xf3, 0x69, 0xc4,
	0x31, 0x8d, 0x48, 0xea, 0x95, 0x77, 0x56, 0x46, 0x09, 0x27, 0xb0, 0x29, 0xbf, 0xb3, 0x50, 0xc0,
	0xac, 0x21, 0xc2, 0x1a, 0xb4, 0x2d, 0x85, 0x68, 0x38, 0x55, 0x4e, 0x28, 0x66, 0x24, 0xa7, 0xe3,
	0xbd, 0x48, 0x7a, 0xe3, 0x8e, 0xfa, 0x36, 0x8b, 0x6c, 0x94, 0xa6, 0x84, 0x23, 0x1e, 0x91, 0x94,
	0x95, 0xb7, 0xa5, 0x6f, 0xbb, 0xf8, 0xb5, 0x9f, 0x1f, 0xd8, 0x07, 0x11, 0x8e, 0x03, 0x2f, 0x41,
	0xec, 0xb0, 0x54, 0xdc, 0x2e, 0x15, 0x34, 0xf3, 0x6d, 0xc6, 0x11, 0xcf, 0xd9, 0xa5, 0x0b, 0x01,
	0xf6, 0xe3, 0x08, 0xa7, 0x5c, 0x5e, 0x98, 0x3d, 0x30, 0xbb, 0x8e, 0xf9, 0x6a, 0x19, 0xd0, 0xaa,
	0x8a, 0xc7, 0xc5, 0xaf, 0x72, 0xcc, 0x38, 0xbc, 0x0f, 0xfe, 0x53, 0x61, 0x7b, 0x29, 0x4a, 0xb0,
	0xa1, 0x35, 0xb5, 0x85, 0x69, 0x77, 0x46, 0x1d, 0x6e, 0xa1, 0x04, 0x9b, 0x3f, 0x34, 0x30, 0xb7,
	0x99, 0x73, 0xc4, 0xf1, 0x25, 0x0e, 0x52, 0x98, 0x7b, 0xa0, 0xee, 0xe7, 0x8c, 0x93, 0x04, 0x53,
	0x2f, 0x0a, 0x4a, 0x08, 0x50, 0x47, 0xcf, 0x03, 0xd8, 0x07, 0x80, 0x64, 0x98, 0xca, 0x74, 0x0d,
	0xbd, 0x59, 0x5b, 0xa8, 0x77, 0x1e, 0x5b, 0x93, 0x6a, 0x6d, 0x8d, 0xc4, 0xbd, 0xad, 0x20, 0xee,
	0x39, 0x1e, 0x7c, 0x00, 0xae, 0x65, 0x88, 0xf2, 0x08, 0xc5, 0xde, 0x01, 0x8a, 0xe2, 0x9c, 0x62,
	0xa3, 0xd6, 0xd4, 0x16, 0xfe, 0x75, 0xff, 0x2f, 0x8f, 0xd7, 0xe4, 0xa9, 0x48, 0x77, 0x80, 0xe2,
	0x28, 0x40, 0x1c, 0x7b, 0x24, 0x8d, 0x8f, 0x8c, 0xbf, 0x0b, 0xd9, 0x8c, 0x3a, 0xdc, 0x4e, 0xe3,
	0x23, 0xf3, 0x9d, 0x0e, 0x1a, 0xd5, 0x8e, 0xe1, 0x12, 0xa8, 0xe7, 0x59, 0x41, 0x10, 0x8d, 0x29,
	0x08, 0xf5, 0x4e, 0x43, 0xe5, 0xa2, 0x7a, 0x67, 0xad, 0x89, 0xde, 0x6d, 0x22, 0x76, 0xe8, 0x02,
	0x29, 0x17, 0x36, 0xdc, 0x02, 0x53, 0x3e, 0xc5, 0x88, 0xcb, 0x42, 0xd7, 0x3b, 0x8b, 0x95, 0x35,
	0x18, 0x4e, 0xd3, 0x68, 0x11, 0x36, 0xfe, 0x72, 0x4b, 0x8a, 0xe0, 0x49, 0xba, 0xa1, 0xff, 0x19,
	0x4f, 0x52, 0xa0, 0x01, 0xa6, 0x28, 0x4e, 0xc8, 0x40, 0x16, 0x70, 0x5a, 0xdc, 0xc8, 0xdf, 0xbd,
	0x3a, 0x98, 0x1e, 0x56, 0xdc, 0xfc, 0xa8, 0x81, 0xbb, 0x55, 0x13, 0xc1, 0x32, 0x92, 0x32, 0x0c,
	0xd7, 0xc0, 0xcd, 0x4b, 0x3d, 0xf1, 0x30, 0xa5, 0x84, 0x16, 0xe0, 0x7a, 0x07, 0xaa, 0x40, 0x69,
	0xe6, 0x5b, 0xbb, 0xc5, 0x28, 0xbb, 0xd7, 0x2f, 0x76, 0xeb, 0x99, 0x90, 0xc3, 0x17, 0xe0, 0x1f,
	0x8a, 0x59, 0x1e, 0x73, 0x35, 0x36, 0x4f, 0x26, 0x8f, 0xcd, 0xd8, 0xd0, 0xc4, 0xd0, 0x0b, 0x8e,
	0xab, 0x78, 0xe6, 0xd3, 0x8a, 0xb1, 0x56, 0xca, 0xdf, 0x7a, 0x1d, 0x9d, 0xaf, 0x35, 0x60, 0x8c,
	0x00, 0x76, 0x65, 0x28, 0xf0, 0x93, 0x06, 0x6e, 0x8c, 0x7b, 0x7f, 0x70, 0x79, 0x72, 0x16, 0xbf,
	0x78, 0xb7, 0x8d, 0x2b, 0xf5, 0xd9, 0x74, 0x5e, 0x7f, 0xfe, 0xf6, 0x56, 0x5f, 0x84, 0x1d, 0xb1,
	0xae, 0x8e, 0x2f, 0xa4, 0xb6, 0xac, 0x1e, 0x2b, 0xb3, 0x5b, 0xc3, 0xfd, 0xa5, 0x9a, 0x6a, 0xb7,
	0x4e, 0xe0, 0x17, 0x0d, 0xdc, 0x1a, 0xdf, 0x72, 0x78, 0xd5, 0x8e, 0xa8, 0xf5, 0xd1, 0x58, 0xb9,
	0x3a, 0x40, 0x4e, 0x9b, 0xb9, 0x52, 0x64, 0xe6, 0x98, 0x8f, 0x44, 0x66, 0x67, 0xa9, 0x1c, 0x9f,
	0xdb, 0x4a, 0xcb, 0xad, 0x93, 0x91, 0xc4, 0x9c, 0xa4, 0xc0, 0x3a, 0x5a, 0xab, 0x31, 0x7b, 0xda,
	0x35, 0xce, 0x5c, 0x97, 0x56, 0x16, 0x31, 0xcb, 0x27, 0x49, 0xef, 0x8d, 0x0e, 0xe6, 0x7d, 0x92,
	0x4c, 0x0c, 0xb3, 0x37, 0x57, 0x35, 0x09, 0x3b, 0x62, 0x2f, 0xec, 0x68, 0x2f, 0x37, 0x4a, 0x44,
	0x48, 0x62, 0x94, 0x86, 0x16, 0xa1, 0xa1, 0x1d, 0xe2, 0xb4, 0xd8, 0x1a, 0xf6, 0x99, 0xd3, 0xea,
	0x7f, 0xb0, 0x25, 0x65, 0xbc, 0xd7, 0x6b, 0xeb, 0xdd, 0xee, 0x07, 0xbd, 0xb9, 0x2e, 0x81, 0xdd,
	0x80, 0x59, 0xd2, 0x14, 0xd6, 0x5e, 0xdb, 0x2a, 0x1d, 0xb3, 0x53, 0x25, 0xe9, 0x77, 0x03, 0xd6,
	0x1f, 0x4a, 0xfa, 0x7b, 0xed, 0xbe, 0x92, 0x7c, 0xd7, 0xe7, 0xe5, 0xb9, 0xe3, 0x74, 0x03, 0xe6,
	0x38, 0x43, 0x91, 0xe3, 0xec, 0xb5, 0x1d, 0x47, 0xc9, 0xf6, 0xa7, 0x8a, 0x38, 0x1f, 0xfe, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0x0e, 0xde, 0xef, 0x03, 0x68, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignCriterionServiceClient is the client API for CampaignCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignCriterionServiceClient interface {
	// Returns the requested criterion in full detail.
	GetCampaignCriterion(ctx context.Context, in *GetCampaignCriterionRequest, opts ...grpc.CallOption) (*resources.CampaignCriterion, error)
	// Creates, updates, or removes criteria. Operation statuses are returned.
	MutateCampaignCriteria(ctx context.Context, in *MutateCampaignCriteriaRequest, opts ...grpc.CallOption) (*MutateCampaignCriteriaResponse, error)
}

type campaignCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignCriterionServiceClient(cc grpc.ClientConnInterface) CampaignCriterionServiceClient {
	return &campaignCriterionServiceClient{cc}
}

func (c *campaignCriterionServiceClient) GetCampaignCriterion(ctx context.Context, in *GetCampaignCriterionRequest, opts ...grpc.CallOption) (*resources.CampaignCriterion, error) {
	out := new(resources.CampaignCriterion)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignCriterionService/GetCampaignCriterion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignCriterionServiceClient) MutateCampaignCriteria(ctx context.Context, in *MutateCampaignCriteriaRequest, opts ...grpc.CallOption) (*MutateCampaignCriteriaResponse, error) {
	out := new(MutateCampaignCriteriaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignCriterionService/MutateCampaignCriteria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignCriterionServiceServer is the server API for CampaignCriterionService service.
type CampaignCriterionServiceServer interface {
	// Returns the requested criterion in full detail.
	GetCampaignCriterion(context.Context, *GetCampaignCriterionRequest) (*resources.CampaignCriterion, error)
	// Creates, updates, or removes criteria. Operation statuses are returned.
	MutateCampaignCriteria(context.Context, *MutateCampaignCriteriaRequest) (*MutateCampaignCriteriaResponse, error)
}

// UnimplementedCampaignCriterionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignCriterionServiceServer struct {
}

func (*UnimplementedCampaignCriterionServiceServer) GetCampaignCriterion(ctx context.Context, req *GetCampaignCriterionRequest) (*resources.CampaignCriterion, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCampaignCriterion not implemented")
}
func (*UnimplementedCampaignCriterionServiceServer) MutateCampaignCriteria(ctx context.Context, req *MutateCampaignCriteriaRequest) (*MutateCampaignCriteriaResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignCriteria not implemented")
}

func RegisterCampaignCriterionServiceServer(s *grpc.Server, srv CampaignCriterionServiceServer) {
	s.RegisterService(&_CampaignCriterionService_serviceDesc, srv)
}

func _CampaignCriterionService_GetCampaignCriterion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignCriterionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignCriterionServiceServer).GetCampaignCriterion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignCriterionService/GetCampaignCriterion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignCriterionServiceServer).GetCampaignCriterion(ctx, req.(*GetCampaignCriterionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignCriterionService_MutateCampaignCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignCriterionServiceServer).MutateCampaignCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignCriterionService/MutateCampaignCriteria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignCriterionServiceServer).MutateCampaignCriteria(ctx, req.(*MutateCampaignCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignCriterionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignCriterionService",
	HandlerType: (*CampaignCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignCriterion",
			Handler:    _CampaignCriterionService_GetCampaignCriterion_Handler,
		},
		{
			MethodName: "MutateCampaignCriteria",
			Handler:    _CampaignCriterionService_MutateCampaignCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_criterion_service.proto",
}
