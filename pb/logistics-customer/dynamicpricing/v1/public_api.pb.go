// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: dynamicpricing/v1/public_api.proto

package v1

import (
	v1 "github.com/deliveryhero/logistics-customer/domain/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DPS multiple vendor request payload.
type GetMultipleVendorsFeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlobalEntityId string       `protobuf:"bytes,1,opt,name=global_entity_id,json=globalEntityId,proto3" json:"global_entity_id,omitempty"`
	Customer       *v1.Customer `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	Vendors        []*v1.Vendor `protobuf:"bytes,3,rep,name=vendors,proto3" json:"vendors,omitempty"`
}

func (x *GetMultipleVendorsFeeRequest) Reset() {
	*x = GetMultipleVendorsFeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamicpricing_v1_public_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultipleVendorsFeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultipleVendorsFeeRequest) ProtoMessage() {}

func (x *GetMultipleVendorsFeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dynamicpricing_v1_public_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultipleVendorsFeeRequest.ProtoReflect.Descriptor instead.
func (*GetMultipleVendorsFeeRequest) Descriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_public_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetMultipleVendorsFeeRequest) GetGlobalEntityId() string {
	if x != nil {
		return x.GlobalEntityId
	}
	return ""
}

func (x *GetMultipleVendorsFeeRequest) GetCustomer() *v1.Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *GetMultipleVendorsFeeRequest) GetVendors() []*v1.Vendor {
	if x != nil {
		return x.Vendors
	}
	return nil
}

// DPS multiple vendor response payload.
type GetMultipleVendorsFeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VendorFees []*VendorFee `protobuf:"bytes,1,rep,name=vendor_fees,json=vendorFees,proto3" json:"vendor_fees,omitempty"`
}

func (x *GetMultipleVendorsFeeResponse) Reset() {
	*x = GetMultipleVendorsFeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamicpricing_v1_public_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultipleVendorsFeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultipleVendorsFeeResponse) ProtoMessage() {}

func (x *GetMultipleVendorsFeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dynamicpricing_v1_public_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultipleVendorsFeeResponse.ProtoReflect.Descriptor instead.
func (*GetMultipleVendorsFeeResponse) Descriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_public_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetMultipleVendorsFeeResponse) GetVendorFees() []*VendorFee {
	if x != nil {
		return x.VendorFees
	}
	return nil
}

// DPS single vendor request payload.
type GetSingleVendorFeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlobalEntityId string       `protobuf:"bytes,1,opt,name=global_entity_id,json=globalEntityId,proto3" json:"global_entity_id,omitempty"`
	Vendor         *v1.Vendor   `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Customer       *v1.Customer `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
	Order          *v1.Order    `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetSingleVendorFeeRequest) Reset() {
	*x = GetSingleVendorFeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamicpricing_v1_public_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSingleVendorFeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSingleVendorFeeRequest) ProtoMessage() {}

func (x *GetSingleVendorFeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dynamicpricing_v1_public_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSingleVendorFeeRequest.ProtoReflect.Descriptor instead.
func (*GetSingleVendorFeeRequest) Descriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_public_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetSingleVendorFeeRequest) GetGlobalEntityId() string {
	if x != nil {
		return x.GlobalEntityId
	}
	return ""
}

func (x *GetSingleVendorFeeRequest) GetVendor() *v1.Vendor {
	if x != nil {
		return x.Vendor
	}
	return nil
}

func (x *GetSingleVendorFeeRequest) GetCustomer() *v1.Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *GetSingleVendorFeeRequest) GetOrder() *v1.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

// DPS single vendor response payload.
type GetSingleVendorFeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VendorFee *VendorFee `protobuf:"bytes,1,opt,name=vendor_fee,json=vendorFee,proto3" json:"vendor_fee,omitempty"`
	Customer  *Customer  `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *GetSingleVendorFeeResponse) Reset() {
	*x = GetSingleVendorFeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamicpricing_v1_public_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSingleVendorFeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSingleVendorFeeResponse) ProtoMessage() {}

func (x *GetSingleVendorFeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dynamicpricing_v1_public_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSingleVendorFeeResponse.ProtoReflect.Descriptor instead.
func (*GetSingleVendorFeeResponse) Descriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_public_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetSingleVendorFeeResponse) GetVendorFee() *VendorFee {
	if x != nil {
		return x.VendorFee
	}
	return nil
}

func (x *GetSingleVendorFeeResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

var File_dynamicpricing_v1_public_api_proto protoreflect.FileDescriptor

var file_dynamicpricing_v1_public_api_proto_rawDesc = []byte{
	0x0a, 0x22, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x73, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x10,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x39, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x07, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x73, 0x22, 0x5e, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x66, 0x65, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x46, 0x65, 0x65, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x65, 0x65,
	0x73, 0x22, 0xe6, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x10, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x92, 0x01, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x65,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x65, 0x65, 0x52, 0x09, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x46, 0x65, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x32,
	0xfa, 0x01, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x50, 0x49, 0x12, 0x7a, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x73, 0x46, 0x65, 0x65, 0x12, 0x2f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x46, 0x65, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x46, 0x65,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x65, 0x65, 0x12,
	0x2c, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x87, 0x01, 0x0a,
	0x35, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x68, 0x65, 0x72,
	0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x70,
	0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x68, 0x65, 0x72,
	0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2d, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dynamicpricing_v1_public_api_proto_rawDescOnce sync.Once
	file_dynamicpricing_v1_public_api_proto_rawDescData = file_dynamicpricing_v1_public_api_proto_rawDesc
)

func file_dynamicpricing_v1_public_api_proto_rawDescGZIP() []byte {
	file_dynamicpricing_v1_public_api_proto_rawDescOnce.Do(func() {
		file_dynamicpricing_v1_public_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_dynamicpricing_v1_public_api_proto_rawDescData)
	})
	return file_dynamicpricing_v1_public_api_proto_rawDescData
}

var file_dynamicpricing_v1_public_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dynamicpricing_v1_public_api_proto_goTypes = []interface{}{
	(*GetMultipleVendorsFeeRequest)(nil),  // 0: dynamicpricing.v1.GetMultipleVendorsFeeRequest
	(*GetMultipleVendorsFeeResponse)(nil), // 1: dynamicpricing.v1.GetMultipleVendorsFeeResponse
	(*GetSingleVendorFeeRequest)(nil),     // 2: dynamicpricing.v1.GetSingleVendorFeeRequest
	(*GetSingleVendorFeeResponse)(nil),    // 3: dynamicpricing.v1.GetSingleVendorFeeResponse
	(*v1.Customer)(nil),                   // 4: domain.v1.Customer
	(*v1.Vendor)(nil),                     // 5: domain.v1.Vendor
	(*VendorFee)(nil),                     // 6: dynamicpricing.v1.VendorFee
	(*v1.Order)(nil),                      // 7: domain.v1.Order
	(*Customer)(nil),                      // 8: dynamicpricing.v1.Customer
}
var file_dynamicpricing_v1_public_api_proto_depIdxs = []int32{
	4,  // 0: dynamicpricing.v1.GetMultipleVendorsFeeRequest.customer:type_name -> domain.v1.Customer
	5,  // 1: dynamicpricing.v1.GetMultipleVendorsFeeRequest.vendors:type_name -> domain.v1.Vendor
	6,  // 2: dynamicpricing.v1.GetMultipleVendorsFeeResponse.vendor_fees:type_name -> dynamicpricing.v1.VendorFee
	5,  // 3: dynamicpricing.v1.GetSingleVendorFeeRequest.vendor:type_name -> domain.v1.Vendor
	4,  // 4: dynamicpricing.v1.GetSingleVendorFeeRequest.customer:type_name -> domain.v1.Customer
	7,  // 5: dynamicpricing.v1.GetSingleVendorFeeRequest.order:type_name -> domain.v1.Order
	6,  // 6: dynamicpricing.v1.GetSingleVendorFeeResponse.vendor_fee:type_name -> dynamicpricing.v1.VendorFee
	8,  // 7: dynamicpricing.v1.GetSingleVendorFeeResponse.customer:type_name -> dynamicpricing.v1.Customer
	0,  // 8: dynamicpricing.v1.PublicAPI.GetMultipleVendorsFee:input_type -> dynamicpricing.v1.GetMultipleVendorsFeeRequest
	2,  // 9: dynamicpricing.v1.PublicAPI.GetSingleVendorFee:input_type -> dynamicpricing.v1.GetSingleVendorFeeRequest
	1,  // 10: dynamicpricing.v1.PublicAPI.GetMultipleVendorsFee:output_type -> dynamicpricing.v1.GetMultipleVendorsFeeResponse
	3,  // 11: dynamicpricing.v1.PublicAPI.GetSingleVendorFee:output_type -> dynamicpricing.v1.GetSingleVendorFeeResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_dynamicpricing_v1_public_api_proto_init() }
func file_dynamicpricing_v1_public_api_proto_init() {
	if File_dynamicpricing_v1_public_api_proto != nil {
		return
	}
	file_dynamicpricing_v1_response_dto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dynamicpricing_v1_public_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultipleVendorsFeeRequest); i {
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
		file_dynamicpricing_v1_public_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultipleVendorsFeeResponse); i {
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
		file_dynamicpricing_v1_public_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSingleVendorFeeRequest); i {
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
		file_dynamicpricing_v1_public_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSingleVendorFeeResponse); i {
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
			RawDescriptor: file_dynamicpricing_v1_public_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dynamicpricing_v1_public_api_proto_goTypes,
		DependencyIndexes: file_dynamicpricing_v1_public_api_proto_depIdxs,
		MessageInfos:      file_dynamicpricing_v1_public_api_proto_msgTypes,
	}.Build()
	File_dynamicpricing_v1_public_api_proto = out.File
	file_dynamicpricing_v1_public_api_proto_rawDesc = nil
	file_dynamicpricing_v1_public_api_proto_goTypes = nil
	file_dynamicpricing_v1_public_api_proto_depIdxs = nil
}
