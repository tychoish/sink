// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1beta1/search.proto

package datacatalog

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The different types of resources that can be returned in search.
type SearchResultType int32

const (
	// Default unknown type.
	SearchResultType_SEARCH_RESULT_TYPE_UNSPECIFIED SearchResultType = 0
	// An [Entry][google.cloud.datacatalog.v1beta1.Entry].
	SearchResultType_ENTRY SearchResultType = 1
	// A [TagTemplate][google.cloud.datacatalog.v1beta1.TagTemplate].
	SearchResultType_TAG_TEMPLATE SearchResultType = 2
	// An [EntryGroup][google.cloud.datacatalog.v1beta1.EntryGroup].
	SearchResultType_ENTRY_GROUP SearchResultType = 3
)

var SearchResultType_name = map[int32]string{
	0: "SEARCH_RESULT_TYPE_UNSPECIFIED",
	1: "ENTRY",
	2: "TAG_TEMPLATE",
	3: "ENTRY_GROUP",
}

var SearchResultType_value = map[string]int32{
	"SEARCH_RESULT_TYPE_UNSPECIFIED": 0,
	"ENTRY":                          1,
	"TAG_TEMPLATE":                   2,
	"ENTRY_GROUP":                    3,
}

func (x SearchResultType) String() string {
	return proto.EnumName(SearchResultType_name, int32(x))
}

func (SearchResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_075a004a2049c613, []int{0}
}

// A result that appears in the response of a search request. Each result
// captures details of one entry that matches the search.
type SearchCatalogResult struct {
	// Type of the search result. This field can be used to determine which Get
	// method to call to fetch the full resource.
	SearchResultType SearchResultType `protobuf:"varint,1,opt,name=search_result_type,json=searchResultType,proto3,enum=google.cloud.datacatalog.v1beta1.SearchResultType" json:"search_result_type,omitempty"`
	// Sub-type of the search result. This is a dot-delimited description of the
	// resource's full type, and is the same as the value callers would provide in
	// the "type" search facet.  Examples: `entry.table`, `entry.dataStream`,
	// `tagTemplate`.
	SearchResultSubtype string `protobuf:"bytes,2,opt,name=search_result_subtype,json=searchResultSubtype,proto3" json:"search_result_subtype,omitempty"`
	// The relative resource name of the resource in URL format.
	// Examples:
	//
	//  * `projects/{project_id}/locations/{location_id}/entryGroups/{entry_group_id}/entries/{entry_id}`
	//  * `projects/{project_id}/tagTemplates/{tag_template_id}`
	RelativeResourceName string `protobuf:"bytes,3,opt,name=relative_resource_name,json=relativeResourceName,proto3" json:"relative_resource_name,omitempty"`
	// The full name of the cloud resource the entry belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name.
	// Example:
	//
	//  * `//bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId`
	LinkedResource       string   `protobuf:"bytes,4,opt,name=linked_resource,json=linkedResource,proto3" json:"linked_resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchCatalogResult) Reset()         { *m = SearchCatalogResult{} }
func (m *SearchCatalogResult) String() string { return proto.CompactTextString(m) }
func (*SearchCatalogResult) ProtoMessage()    {}
func (*SearchCatalogResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_075a004a2049c613, []int{0}
}

func (m *SearchCatalogResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchCatalogResult.Unmarshal(m, b)
}
func (m *SearchCatalogResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchCatalogResult.Marshal(b, m, deterministic)
}
func (m *SearchCatalogResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchCatalogResult.Merge(m, src)
}
func (m *SearchCatalogResult) XXX_Size() int {
	return xxx_messageInfo_SearchCatalogResult.Size(m)
}
func (m *SearchCatalogResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchCatalogResult.DiscardUnknown(m)
}

var xxx_messageInfo_SearchCatalogResult proto.InternalMessageInfo

func (m *SearchCatalogResult) GetSearchResultType() SearchResultType {
	if m != nil {
		return m.SearchResultType
	}
	return SearchResultType_SEARCH_RESULT_TYPE_UNSPECIFIED
}

func (m *SearchCatalogResult) GetSearchResultSubtype() string {
	if m != nil {
		return m.SearchResultSubtype
	}
	return ""
}

func (m *SearchCatalogResult) GetRelativeResourceName() string {
	if m != nil {
		return m.RelativeResourceName
	}
	return ""
}

func (m *SearchCatalogResult) GetLinkedResource() string {
	if m != nil {
		return m.LinkedResource
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.datacatalog.v1beta1.SearchResultType", SearchResultType_name, SearchResultType_value)
	proto.RegisterType((*SearchCatalogResult)(nil), "google.cloud.datacatalog.v1beta1.SearchCatalogResult")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1beta1/search.proto", fileDescriptor_075a004a2049c613)
}

var fileDescriptor_075a004a2049c613 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0xaf, 0xd2, 0x30,
	0x18, 0xc5, 0x1d, 0x57, 0x4d, 0x6e, 0x35, 0xf7, 0x2e, 0x45, 0x0d, 0xe1, 0x41, 0x09, 0x31, 0x91,
	0x98, 0xb8, 0x06, 0xf4, 0xcd, 0x27, 0xc4, 0x8a, 0x44, 0xc4, 0xa5, 0x1b, 0x0f, 0xf8, 0x52, 0xbb,
	0xed, 0x63, 0x2c, 0x76, 0xeb, 0xb2, 0x75, 0x44, 0xfe, 0x05, 0xff, 0x62, 0x1f, 0x8d, 0xed, 0x30,
	0x48, 0x62, 0x78, 0xdc, 0x39, 0xbf, 0x73, 0xbe, 0x6f, 0x6d, 0xd1, 0xab, 0x54, 0xa9, 0x54, 0x02,
	0x89, 0xa5, 0x6a, 0x12, 0x92, 0x08, 0x2d, 0x62, 0xa1, 0x85, 0x54, 0x29, 0xd9, 0x8f, 0x23, 0xd0,
	0x62, 0x4c, 0x6a, 0x10, 0x55, 0xbc, 0xf3, 0xca, 0x4a, 0x69, 0x85, 0x07, 0x16, 0xf7, 0x0c, 0xee,
	0x9d, 0xe0, 0x5e, 0x8b, 0xf7, 0x9f, 0xb5, 0x85, 0xa2, 0xcc, 0xc8, 0x36, 0x03, 0x99, 0xf0, 0x08,
	0x76, 0x62, 0x9f, 0xa9, 0xca, 0x56, 0xf4, 0x2f, 0x4f, 0x8c, 0x55, 0x9e, 0xab, 0xa2, 0xc5, 0x8f,
	0x7d, 0xe6, 0x2b, 0x6a, 0xb6, 0x44, 0x67, 0x39, 0xd4, 0x5a, 0xe4, 0xa5, 0x05, 0x86, 0x3f, 0x3b,
	0xa8, 0x1b, 0x98, 0x1d, 0x67, 0xb6, 0x87, 0x41, 0xdd, 0x48, 0x8d, 0xbf, 0x21, 0x6c, 0x57, 0xe7,
	0x95, 0x11, 0xb8, 0x3e, 0x94, 0xd0, 0x73, 0x06, 0xce, 0xe8, 0x66, 0x32, 0xf1, 0x2e, 0xfd, 0x87,
	0x67, 0x2b, 0x6d, 0x57, 0x78, 0x28, 0x81, 0xb9, 0xf5, 0x99, 0x82, 0x27, 0xe8, 0xf1, 0xbf, 0x13,
	0xea, 0x26, 0x32, 0x43, 0x3a, 0x03, 0x67, 0x74, 0xcd, 0xba, 0xa7, 0x81, 0xc0, 0x5a, 0xf8, 0x0d,
	0x7a, 0x52, 0x81, 0x14, 0x3a, 0xdb, 0xc3, 0x9f, 0x94, 0x6a, 0xaa, 0x18, 0x78, 0x21, 0x72, 0xe8,
	0x5d, 0x99, 0xd0, 0xa3, 0xa3, 0xcb, 0x5a, 0x73, 0x25, 0x72, 0xc0, 0x2f, 0xd0, 0xad, 0xcc, 0x8a,
	0xef, 0x90, 0xfc, 0xcd, 0xf4, 0xee, 0x1a, 0xfc, 0xc6, 0xca, 0x47, 0xf8, 0x65, 0x82, 0xdc, 0xf3,
	0xc5, 0xf1, 0x10, 0x3d, 0x0d, 0xe8, 0x94, 0xcd, 0x3e, 0x72, 0x46, 0x83, 0xf5, 0x32, 0xe4, 0xe1,
	0xc6, 0xa7, 0x7c, 0xbd, 0x0a, 0x7c, 0x3a, 0x5b, 0x7c, 0x58, 0xd0, 0xf7, 0xee, 0x1d, 0x7c, 0x8d,
	0xee, 0xd1, 0x55, 0xc8, 0x36, 0xae, 0x83, 0x5d, 0xf4, 0x30, 0x9c, 0xce, 0x79, 0x48, 0x3f, 0xfb,
	0xcb, 0x69, 0x48, 0xdd, 0x0e, 0xbe, 0x45, 0x0f, 0x8c, 0xc9, 0xe7, 0xec, 0xcb, 0xda, 0x77, 0xaf,
	0xde, 0xfd, 0x40, 0xcf, 0x63, 0x95, 0x5f, 0x3c, 0x43, 0xdf, 0xf9, 0xfa, 0xa9, 0x65, 0x52, 0x25,
	0x45, 0x91, 0x7a, 0xaa, 0x4a, 0x49, 0x0a, 0x85, 0xb9, 0x38, 0x62, 0x2d, 0x51, 0x66, 0xf5, 0xff,
	0xdf, 0xc2, 0xdb, 0x13, 0xed, 0x97, 0xe3, 0x44, 0xf7, 0x4d, 0xf4, 0xf5, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x90, 0xae, 0x3d, 0x1e, 0xb7, 0x02, 0x00, 0x00,
}