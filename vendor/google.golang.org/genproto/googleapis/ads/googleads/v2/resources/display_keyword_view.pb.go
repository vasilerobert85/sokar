// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/display_keyword_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// A display keyword view.
type DisplayKeywordView struct {
	// The resource name of the display keyword view.
	// Display Keyword view resource names have the form:
	//
	// `customers/{customer_id}/displayKeywordViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisplayKeywordView) Reset()         { *m = DisplayKeywordView{} }
func (m *DisplayKeywordView) String() string { return proto.CompactTextString(m) }
func (*DisplayKeywordView) ProtoMessage()    {}
func (*DisplayKeywordView) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbae9f6689a25c70, []int{0}
}

func (m *DisplayKeywordView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisplayKeywordView.Unmarshal(m, b)
}
func (m *DisplayKeywordView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisplayKeywordView.Marshal(b, m, deterministic)
}
func (m *DisplayKeywordView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisplayKeywordView.Merge(m, src)
}
func (m *DisplayKeywordView) XXX_Size() int {
	return xxx_messageInfo_DisplayKeywordView.Size(m)
}
func (m *DisplayKeywordView) XXX_DiscardUnknown() {
	xxx_messageInfo_DisplayKeywordView.DiscardUnknown(m)
}

var xxx_messageInfo_DisplayKeywordView proto.InternalMessageInfo

func (m *DisplayKeywordView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*DisplayKeywordView)(nil), "google.ads.googleads.v2.resources.DisplayKeywordView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/display_keyword_view.proto", fileDescriptor_bbae9f6689a25c70)
}

var fileDescriptor_bbae9f6689a25c70 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4a, 0xec, 0x30,
	0x14, 0xc6, 0x69, 0x2f, 0x5c, 0xb0, 0xe8, 0xa6, 0x1b, 0x45, 0x5c, 0x38, 0xca, 0x80, 0xab, 0x04,
	0xea, 0xca, 0xe8, 0x26, 0x83, 0x30, 0xa0, 0x20, 0xc3, 0x2c, 0xba, 0x90, 0x42, 0x89, 0x93, 0x10,
	0x82, 0x6d, 0x4e, 0xc9, 0xa9, 0x2d, 0xb3, 0xf7, 0x49, 0x5c, 0xfa, 0x28, 0x3e, 0x8a, 0x4f, 0x21,
	0x9d, 0x98, 0x6c, 0x04, 0xdd, 0x7d, 0x24, 0xbf, 0xef, 0x0f, 0x27, 0xbb, 0xd1, 0x00, 0xba, 0x51,
	0x54, 0x48, 0xa4, 0x5e, 0x4e, 0x6a, 0x28, 0xa8, 0x53, 0x08, 0x2f, 0x6e, 0xa3, 0x90, 0x4a, 0x83,
	0x5d, 0x23, 0xb6, 0xf5, 0xb3, 0xda, 0x8e, 0xe0, 0x64, 0x3d, 0x18, 0x35, 0x92, 0xce, 0x41, 0x0f,
	0xf9, 0xcc, 0x5b, 0x88, 0x90, 0x48, 0xa2, 0x9b, 0x0c, 0x05, 0x89, 0xee, 0xe3, 0x93, 0x50, 0xd0,
	0x19, 0x2a, 0xac, 0x85, 0x5e, 0xf4, 0x06, 0x2c, 0xfa, 0x80, 0xb3, 0xab, 0x2c, 0xbf, 0xf5, 0xf1,
	0xf7, 0x3e, 0xbd, 0x34, 0x6a, 0xcc, 0xcf, 0xb3, 0x83, 0x10, 0x50, 0x5b, 0xd1, 0xaa, 0xa3, 0xe4,
	0x34, 0xb9, 0xd8, 0x5b, 0xef, 0x87, 0xc7, 0x07, 0xd1, 0xaa, 0xc5, 0x6b, 0x9a, 0xcd, 0x37, 0xd0,
	0x92, 0x3f, 0x27, 0x2c, 0x0e, 0x7f, 0x56, 0xac, 0xa6, 0xf6, 0x55, 0xf2, 0x78, 0xf7, 0xed, 0xd6,
	0xd0, 0x08, 0xab, 0x09, 0x38, 0x4d, 0xb5, 0xb2, 0xbb, 0x6d, 0xe1, 0x1c, 0x9d, 0xc1, 0x5f, 0xae,
	0x73, 0x1d, 0xd5, 0x5b, 0xfa, 0x6f, 0xc9, 0xf9, 0x7b, 0x3a, 0x5b, 0xfa, 0x48, 0x2e, 0x91, 0x78,
	0x39, 0xa9, 0xb2, 0x20, 0xeb, 0x40, 0x7e, 0x04, 0xa6, 0xe2, 0x12, 0xab, 0xc8, 0x54, 0x65, 0x51,
	0x45, 0xe6, 0x33, 0x9d, 0xfb, 0x0f, 0xc6, 0xb8, 0x44, 0xc6, 0x22, 0xc5, 0x58, 0x59, 0x30, 0x16,
	0xb9, 0xa7, 0xff, 0xbb, 0xb1, 0x97, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xf7, 0x89, 0xaa,
	0xc9, 0x01, 0x00, 0x00,
}
