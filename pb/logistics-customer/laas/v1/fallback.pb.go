// Definitions of all message used only in the availability fallback

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: laas/v1/fallback.proto

package v1

import (
	v1 "github.com/deliveryhero/logistics-customer/domain/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Snapshot vendor representation.
type SnapshotVendor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlobalEntityId       string                        `protobuf:"bytes,1,opt,name=global_entity_id,json=globalEntityId,proto3" json:"global_entity_id,omitempty"`
	VendorId             string                        `protobuf:"bytes,2,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	CustomerTypes        []v1.CustomerType             `protobuf:"varint,3,rep,packed,name=customer_types,json=customerTypes,proto3,enum=domain.v1.CustomerType" json:"customer_types,omitempty"`
	VerticalType         *wrapperspb.StringValue       `protobuf:"bytes,4,opt,name=vertical_type,json=verticalType,proto3" json:"vertical_type,omitempty"`
	VerticalParent       *wrapperspb.StringValue       `protobuf:"bytes,5,opt,name=vertical_parent,json=verticalParent,proto3" json:"vertical_parent,omitempty"`
	DeliveryTypes        []v1.DeliveryType             `protobuf:"varint,6,rep,packed,name=delivery_types,json=deliveryTypes,proto3,enum=domain.v1.DeliveryType" json:"delivery_types,omitempty"`
	FilterAttributes     []*VendorFilterAttribute      `protobuf:"bytes,7,rep,name=filter_attributes,json=filterAttributes,proto3" json:"filter_attributes,omitempty"`
	VendorSchedule       *VendorSchedule               `protobuf:"bytes,8,opt,name=vendor_schedule,json=vendorSchedule,proto3" json:"vendor_schedule,omitempty"`
	Private              *wrapperspb.BoolValue         `protobuf:"bytes,9,opt,name=private,proto3" json:"private,omitempty"`
	TestVendor           *wrapperspb.BoolValue         `protobuf:"bytes,10,opt,name=test_vendor,json=testVendor,proto3" json:"test_vendor,omitempty"`
	AllowPreorder        *wrapperspb.BoolValue         `protobuf:"bytes,11,opt,name=allow_preorder,json=allowPreorder,proto3" json:"allow_preorder,omitempty"`
	VerticalTypeIds      []string                      `protobuf:"bytes,12,rep,name=vertical_type_ids,json=verticalTypeIds,proto3" json:"vertical_type_ids,omitempty"`
	PricingResponse      *FallbackPricingResponse      `protobuf:"bytes,13,opt,name=pricing_response,json=pricingResponse,proto3" json:"pricing_response,omitempty"`
	TimeEstimateResponse *FallbackTimeEstimateResponse `protobuf:"bytes,14,opt,name=time_estimate_response,json=timeEstimateResponse,proto3" json:"time_estimate_response,omitempty"`
	CellIds              []uint64                      `protobuf:"varint,15,rep,packed,name=cell_ids,json=cellIds,proto3" json:"cell_ids,omitempty"`
	Location             *v1.Coordinate                `protobuf:"bytes,16,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *SnapshotVendor) Reset() {
	*x = SnapshotVendor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laas_v1_fallback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotVendor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotVendor) ProtoMessage() {}

func (x *SnapshotVendor) ProtoReflect() protoreflect.Message {
	mi := &file_laas_v1_fallback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotVendor.ProtoReflect.Descriptor instead.
func (*SnapshotVendor) Descriptor() ([]byte, []int) {
	return file_laas_v1_fallback_proto_rawDescGZIP(), []int{0}
}

func (x *SnapshotVendor) GetGlobalEntityId() string {
	if x != nil {
		return x.GlobalEntityId
	}
	return ""
}

func (x *SnapshotVendor) GetVendorId() string {
	if x != nil {
		return x.VendorId
	}
	return ""
}

func (x *SnapshotVendor) GetCustomerTypes() []v1.CustomerType {
	if x != nil {
		return x.CustomerTypes
	}
	return nil
}

func (x *SnapshotVendor) GetVerticalType() *wrapperspb.StringValue {
	if x != nil {
		return x.VerticalType
	}
	return nil
}

func (x *SnapshotVendor) GetVerticalParent() *wrapperspb.StringValue {
	if x != nil {
		return x.VerticalParent
	}
	return nil
}

func (x *SnapshotVendor) GetDeliveryTypes() []v1.DeliveryType {
	if x != nil {
		return x.DeliveryTypes
	}
	return nil
}

func (x *SnapshotVendor) GetFilterAttributes() []*VendorFilterAttribute {
	if x != nil {
		return x.FilterAttributes
	}
	return nil
}

func (x *SnapshotVendor) GetVendorSchedule() *VendorSchedule {
	if x != nil {
		return x.VendorSchedule
	}
	return nil
}

func (x *SnapshotVendor) GetPrivate() *wrapperspb.BoolValue {
	if x != nil {
		return x.Private
	}
	return nil
}

func (x *SnapshotVendor) GetTestVendor() *wrapperspb.BoolValue {
	if x != nil {
		return x.TestVendor
	}
	return nil
}

func (x *SnapshotVendor) GetAllowPreorder() *wrapperspb.BoolValue {
	if x != nil {
		return x.AllowPreorder
	}
	return nil
}

func (x *SnapshotVendor) GetVerticalTypeIds() []string {
	if x != nil {
		return x.VerticalTypeIds
	}
	return nil
}

func (x *SnapshotVendor) GetPricingResponse() *FallbackPricingResponse {
	if x != nil {
		return x.PricingResponse
	}
	return nil
}

func (x *SnapshotVendor) GetTimeEstimateResponse() *FallbackTimeEstimateResponse {
	if x != nil {
		return x.TimeEstimateResponse
	}
	return nil
}

func (x *SnapshotVendor) GetCellIds() []uint64 {
	if x != nil {
		return x.CellIds
	}
	return nil
}

func (x *SnapshotVendor) GetLocation() *v1.Coordinate {
	if x != nil {
		return x.Location
	}
	return nil
}

type SnapshotIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlobalEntityId  string            `protobuf:"bytes,1,opt,name=global_entity_id,json=globalEntityId,proto3" json:"global_entity_id,omitempty"`
	CellId          uint64            `protobuf:"varint,2,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	SnapshotVendors []*SnapshotVendor `protobuf:"bytes,3,rep,name=snapshot_vendors,json=snapshotVendors,proto3" json:"snapshot_vendors,omitempty"`
}

func (x *SnapshotIndex) Reset() {
	*x = SnapshotIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laas_v1_fallback_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotIndex) ProtoMessage() {}

func (x *SnapshotIndex) ProtoReflect() protoreflect.Message {
	mi := &file_laas_v1_fallback_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotIndex.ProtoReflect.Descriptor instead.
func (*SnapshotIndex) Descriptor() ([]byte, []int) {
	return file_laas_v1_fallback_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotIndex) GetGlobalEntityId() string {
	if x != nil {
		return x.GlobalEntityId
	}
	return ""
}

func (x *SnapshotIndex) GetCellId() uint64 {
	if x != nil {
		return x.CellId
	}
	return 0
}

func (x *SnapshotIndex) GetSnapshotVendors() []*SnapshotVendor {
	if x != nil {
		return x.SnapshotVendors
	}
	return nil
}

var File_laas_v1_fallback_proto protoreflect.FileDescriptor

var file_laas_v1_fallback_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x61, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x61, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x13, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x6c, 0x61, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6c, 0x61, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x07, 0x0a, 0x0e, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x10,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a,
	0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x4b, 0x0a,
	0x11, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x61, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x10, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0f, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x0e, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x07,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12,
	0x41, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x72, 0x65, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x65, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x73, 0x12, 0x4b,
	0x0a, 0x10, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x61, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x16, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x61,
	0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x69,
	0x6d, 0x65, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x14, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x65, 0x6c, 0x6c,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x63, 0x65, 0x6c, 0x6c,
	0x49, 0x64, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x96, 0x01, 0x0a, 0x0d, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x10, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x0f,
	0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x42,
	0x72, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x68,
	0x65, 0x72, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x6c, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d,
	0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x68, 0x65, 0x72, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x6c, 0x61, 0x61, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_laas_v1_fallback_proto_rawDescOnce sync.Once
	file_laas_v1_fallback_proto_rawDescData = file_laas_v1_fallback_proto_rawDesc
)

func file_laas_v1_fallback_proto_rawDescGZIP() []byte {
	file_laas_v1_fallback_proto_rawDescOnce.Do(func() {
		file_laas_v1_fallback_proto_rawDescData = protoimpl.X.CompressGZIP(file_laas_v1_fallback_proto_rawDescData)
	})
	return file_laas_v1_fallback_proto_rawDescData
}

var file_laas_v1_fallback_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_laas_v1_fallback_proto_goTypes = []interface{}{
	(*SnapshotVendor)(nil),               // 0: laas.v1.SnapshotVendor
	(*SnapshotIndex)(nil),                // 1: laas.v1.SnapshotIndex
	(v1.CustomerType)(0),                 // 2: domain.v1.CustomerType
	(*wrapperspb.StringValue)(nil),       // 3: google.protobuf.StringValue
	(v1.DeliveryType)(0),                 // 4: domain.v1.DeliveryType
	(*VendorFilterAttribute)(nil),        // 5: laas.v1.VendorFilterAttribute
	(*VendorSchedule)(nil),               // 6: laas.v1.VendorSchedule
	(*wrapperspb.BoolValue)(nil),         // 7: google.protobuf.BoolValue
	(*FallbackPricingResponse)(nil),      // 8: laas.v1.FallbackPricingResponse
	(*FallbackTimeEstimateResponse)(nil), // 9: laas.v1.FallbackTimeEstimateResponse
	(*v1.Coordinate)(nil),                // 10: domain.v1.Coordinate
}
var file_laas_v1_fallback_proto_depIdxs = []int32{
	2,  // 0: laas.v1.SnapshotVendor.customer_types:type_name -> domain.v1.CustomerType
	3,  // 1: laas.v1.SnapshotVendor.vertical_type:type_name -> google.protobuf.StringValue
	3,  // 2: laas.v1.SnapshotVendor.vertical_parent:type_name -> google.protobuf.StringValue
	4,  // 3: laas.v1.SnapshotVendor.delivery_types:type_name -> domain.v1.DeliveryType
	5,  // 4: laas.v1.SnapshotVendor.filter_attributes:type_name -> laas.v1.VendorFilterAttribute
	6,  // 5: laas.v1.SnapshotVendor.vendor_schedule:type_name -> laas.v1.VendorSchedule
	7,  // 6: laas.v1.SnapshotVendor.private:type_name -> google.protobuf.BoolValue
	7,  // 7: laas.v1.SnapshotVendor.test_vendor:type_name -> google.protobuf.BoolValue
	7,  // 8: laas.v1.SnapshotVendor.allow_preorder:type_name -> google.protobuf.BoolValue
	8,  // 9: laas.v1.SnapshotVendor.pricing_response:type_name -> laas.v1.FallbackPricingResponse
	9,  // 10: laas.v1.SnapshotVendor.time_estimate_response:type_name -> laas.v1.FallbackTimeEstimateResponse
	10, // 11: laas.v1.SnapshotVendor.location:type_name -> domain.v1.Coordinate
	0,  // 12: laas.v1.SnapshotIndex.snapshot_vendors:type_name -> laas.v1.SnapshotVendor
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_laas_v1_fallback_proto_init() }
func file_laas_v1_fallback_proto_init() {
	if File_laas_v1_fallback_proto != nil {
		return
	}
	file_laas_v1_domain_proto_init()
	file_laas_v1_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_laas_v1_fallback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotVendor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_laas_v1_fallback_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotIndex); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_laas_v1_fallback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_laas_v1_fallback_proto_goTypes,
		DependencyIndexes: file_laas_v1_fallback_proto_depIdxs,
		MessageInfos:      file_laas_v1_fallback_proto_msgTypes,
	}.Build()
	File_laas_v1_fallback_proto = out.File
	file_laas_v1_fallback_proto_rawDesc = nil
	file_laas_v1_fallback_proto_goTypes = nil
	file_laas_v1_fallback_proto_depIdxs = nil
}
