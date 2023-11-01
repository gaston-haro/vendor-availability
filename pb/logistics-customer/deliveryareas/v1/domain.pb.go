// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: deliveryareas/v1/domain.proto

package v1

import (
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

// Operators allowed within the Vendor Filter DSL.
type VendorFilterOperator int32

const (
	VendorFilterOperator_VENDOR_FILTER_OPERATOR_INVALID VendorFilterOperator = 0
	VendorFilterOperator_VENDOR_FILTER_OPERATOR_IS      VendorFilterOperator = 1
	VendorFilterOperator_VENDOR_FILTER_OPERATOR_NOT     VendorFilterOperator = 2
	VendorFilterOperator_VENDOR_FILTER_OPERATOR_IN      VendorFilterOperator = 3
	VendorFilterOperator_VENDOR_FILTER_OPERATOR_NOT_ALL VendorFilterOperator = 4
	VendorFilterOperator_VENDOR_FILTER_OPERATOR_ALL     VendorFilterOperator = 5
	VendorFilterOperator_VENDOR_FILTER_OPERATOR_NONE_OF VendorFilterOperator = 6
)

// Enum value maps for VendorFilterOperator.
var (
	VendorFilterOperator_name = map[int32]string{
		0: "VENDOR_FILTER_OPERATOR_INVALID",
		1: "VENDOR_FILTER_OPERATOR_IS",
		2: "VENDOR_FILTER_OPERATOR_NOT",
		3: "VENDOR_FILTER_OPERATOR_IN",
		4: "VENDOR_FILTER_OPERATOR_NOT_ALL",
		5: "VENDOR_FILTER_OPERATOR_ALL",
		6: "VENDOR_FILTER_OPERATOR_NONE_OF",
	}
	VendorFilterOperator_value = map[string]int32{
		"VENDOR_FILTER_OPERATOR_INVALID": 0,
		"VENDOR_FILTER_OPERATOR_IS":      1,
		"VENDOR_FILTER_OPERATOR_NOT":     2,
		"VENDOR_FILTER_OPERATOR_IN":      3,
		"VENDOR_FILTER_OPERATOR_NOT_ALL": 4,
		"VENDOR_FILTER_OPERATOR_ALL":     5,
		"VENDOR_FILTER_OPERATOR_NONE_OF": 6,
	}
)

func (x VendorFilterOperator) Enum() *VendorFilterOperator {
	p := new(VendorFilterOperator)
	*p = x
	return p
}

func (x VendorFilterOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VendorFilterOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_deliveryareas_v1_domain_proto_enumTypes[0].Descriptor()
}

func (VendorFilterOperator) Type() protoreflect.EnumType {
	return &file_deliveryareas_v1_domain_proto_enumTypes[0]
}

func (x VendorFilterOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VendorFilterOperator.Descriptor instead.
func (VendorFilterOperator) EnumDescriptor() ([]byte, []int) {
	return file_deliveryareas_v1_domain_proto_rawDescGZIP(), []int{0}
}

// Fields that can be used to filter vendors with.
type VendorFilterField int32

const (
	VendorFilterField_VENDOR_FILTER_FIELD_INVALID        VendorFilterField = 0
	VendorFilterField_VENDOR_FILTER_FIELD_HALAL          VendorFilterField = 1
	VendorFilterField_VENDOR_FILTER_FIELD_CHAIN_IDS      VendorFilterField = 2
	VendorFilterField_VENDOR_FILTER_FIELD_VERTICAL_TYPES VendorFilterField = 3
	VendorFilterField_VENDOR_FILTER_FIELD_DELIVERY_TYPES VendorFilterField = 4
	VendorFilterField_VENDOR_FILTER_FIELD_TAGS           VendorFilterField = 5
)

// Enum value maps for VendorFilterField.
var (
	VendorFilterField_name = map[int32]string{
		0: "VENDOR_FILTER_FIELD_INVALID",
		1: "VENDOR_FILTER_FIELD_HALAL",
		2: "VENDOR_FILTER_FIELD_CHAIN_IDS",
		3: "VENDOR_FILTER_FIELD_VERTICAL_TYPES",
		4: "VENDOR_FILTER_FIELD_DELIVERY_TYPES",
		5: "VENDOR_FILTER_FIELD_TAGS",
	}
	VendorFilterField_value = map[string]int32{
		"VENDOR_FILTER_FIELD_INVALID":        0,
		"VENDOR_FILTER_FIELD_HALAL":          1,
		"VENDOR_FILTER_FIELD_CHAIN_IDS":      2,
		"VENDOR_FILTER_FIELD_VERTICAL_TYPES": 3,
		"VENDOR_FILTER_FIELD_DELIVERY_TYPES": 4,
		"VENDOR_FILTER_FIELD_TAGS":           5,
	}
)

func (x VendorFilterField) Enum() *VendorFilterField {
	p := new(VendorFilterField)
	*p = x
	return p
}

func (x VendorFilterField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VendorFilterField) Descriptor() protoreflect.EnumDescriptor {
	return file_deliveryareas_v1_domain_proto_enumTypes[1].Descriptor()
}

func (VendorFilterField) Type() protoreflect.EnumType {
	return &file_deliveryareas_v1_domain_proto_enumTypes[1]
}

func (x VendorFilterField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VendorFilterField.Descriptor instead.
func (VendorFilterField) EnumDescriptor() ([]byte, []int) {
	return file_deliveryareas_v1_domain_proto_rawDescGZIP(), []int{1}
}

// Set of expressions used to filter vendors with.
// Multiple filter expressions are chained together with Logical OR operation
type VendorFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expressions []*VendorFilterExpression `protobuf:"bytes,1,rep,name=expressions,proto3" json:"expressions,omitempty"`
}

func (x *VendorFilters) Reset() {
	*x = VendorFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliveryareas_v1_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VendorFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VendorFilters) ProtoMessage() {}

func (x *VendorFilters) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryareas_v1_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VendorFilters.ProtoReflect.Descriptor instead.
func (*VendorFilters) Descriptor() ([]byte, []int) {
	return file_deliveryareas_v1_domain_proto_rawDescGZIP(), []int{0}
}

func (x *VendorFilters) GetExpressions() []*VendorFilterExpression {
	if x != nil {
		return x.Expressions
	}
	return nil
}

// VendorFilterExpression represents a list of filtering conditions that are chained together with Logical AND operation.
type VendorFilterExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conditions []*VendorFilterCondition `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *VendorFilterExpression) Reset() {
	*x = VendorFilterExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliveryareas_v1_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VendorFilterExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VendorFilterExpression) ProtoMessage() {}

func (x *VendorFilterExpression) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryareas_v1_domain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VendorFilterExpression.ProtoReflect.Descriptor instead.
func (*VendorFilterExpression) Descriptor() ([]byte, []int) {
	return file_deliveryareas_v1_domain_proto_rawDescGZIP(), []int{1}
}

func (x *VendorFilterExpression) GetConditions() []*VendorFilterCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

// VendorFilterClause encapsulates a single condition to filter vendors with.
type VendorFilterCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator VendorFilterOperator `protobuf:"varint,1,opt,name=operator,proto3,enum=deliveryareas.v1.VendorFilterOperator" json:"operator,omitempty"`
	Field    VendorFilterField    `protobuf:"varint,2,opt,name=field,proto3,enum=deliveryareas.v1.VendorFilterField" json:"field,omitempty"`
	Values   []string             `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *VendorFilterCondition) Reset() {
	*x = VendorFilterCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliveryareas_v1_domain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VendorFilterCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VendorFilterCondition) ProtoMessage() {}

func (x *VendorFilterCondition) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryareas_v1_domain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VendorFilterCondition.ProtoReflect.Descriptor instead.
func (*VendorFilterCondition) Descriptor() ([]byte, []int) {
	return file_deliveryareas_v1_domain_proto_rawDescGZIP(), []int{2}
}

func (x *VendorFilterCondition) GetOperator() VendorFilterOperator {
	if x != nil {
		return x.Operator
	}
	return VendorFilterOperator_VENDOR_FILTER_OPERATOR_INVALID
}

func (x *VendorFilterCondition) GetField() VendorFilterField {
	if x != nil {
		return x.Field
	}
	return VendorFilterField_VENDOR_FILTER_FIELD_INVALID
}

func (x *VendorFilterCondition) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_deliveryareas_v1_domain_proto protoreflect.FileDescriptor

var file_deliveryareas_v1_domain_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x22, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x61, 0x72, 0x65, 0x61, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x73, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65,
	0x0a, 0x0d, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x54, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x61,
	0x72, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6b, 0x0a, 0x16, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x51, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x61, 0x72,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x15, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x08,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01,
	0x20, 0x00, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x45, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x92, 0x01, 0x08, 0x08, 0x01, 0x22, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2a, 0xbb, 0x02, 0x0a, 0x14,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x1e, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x19, 0x56, 0x45, 0x4e, 0x44,
	0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x4f, 0x52, 0x5f, 0x49, 0x53, 0x10, 0x01, 0x1a, 0x07, 0x82, 0xb5, 0x18, 0x03, 0x24, 0x69, 0x73,
	0x12, 0x28, 0x0a, 0x1a, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45,
	0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x10, 0x02,
	0x1a, 0x08, 0x82, 0xb5, 0x18, 0x04, 0x24, 0x6e, 0x6f, 0x74, 0x12, 0x26, 0x0a, 0x19, 0x56, 0x45,
	0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52,
	0x41, 0x54, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x10, 0x03, 0x1a, 0x07, 0x82, 0xb5, 0x18, 0x03, 0x24,
	0x69, 0x6e, 0x12, 0x2d, 0x0a, 0x1e, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4c,
	0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x04, 0x1a, 0x09, 0x82, 0xb5, 0x18, 0x05, 0x24, 0x6e, 0x61, 0x6c,
	0x6c, 0x12, 0x28, 0x0a, 0x1a, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x54,
	0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x41, 0x4c, 0x4c, 0x10,
	0x05, 0x1a, 0x08, 0x82, 0xb5, 0x18, 0x04, 0x24, 0x61, 0x6c, 0x6c, 0x12, 0x2c, 0x0a, 0x1e, 0x56,
	0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x4f, 0x46, 0x10, 0x06, 0x1a,
	0x08, 0x82, 0xb5, 0x18, 0x04, 0x24, 0x6e, 0x69, 0x6e, 0x2a, 0xad, 0x02, 0x0a, 0x11, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1f, 0x0a, 0x1b, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52,
	0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x28, 0x0a, 0x19, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45,
	0x52, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x48, 0x41, 0x4c, 0x41, 0x4c, 0x10, 0x01, 0x1a,
	0x09, 0x82, 0xb5, 0x18, 0x05, 0x68, 0x61, 0x6c, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x1d, 0x56, 0x45,
	0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x5f, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x02, 0x1a, 0x0a, 0x82,
	0xb5, 0x18, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x22, 0x56, 0x45, 0x4e,
	0x44, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x5f, 0x56, 0x45, 0x52, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x53, 0x10,
	0x03, 0x1a, 0x12, 0x82, 0xb5, 0x18, 0x0e, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x22, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f,
	0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x44, 0x45, 0x4c,
	0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x53, 0x10, 0x04, 0x1a, 0x12, 0x82,
	0xb5, 0x18, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x26, 0x0a, 0x18, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x54,
	0x45, 0x52, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x41, 0x47, 0x53, 0x10, 0x05, 0x1a,
	0x08, 0x82, 0xb5, 0x18, 0x04, 0x74, 0x61, 0x67, 0x73, 0x42, 0x82, 0x01, 0x0a, 0x34, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x68, 0x65, 0x72, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deliveryareas_v1_domain_proto_rawDescOnce sync.Once
	file_deliveryareas_v1_domain_proto_rawDescData = file_deliveryareas_v1_domain_proto_rawDesc
)

func file_deliveryareas_v1_domain_proto_rawDescGZIP() []byte {
	file_deliveryareas_v1_domain_proto_rawDescOnce.Do(func() {
		file_deliveryareas_v1_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_deliveryareas_v1_domain_proto_rawDescData)
	})
	return file_deliveryareas_v1_domain_proto_rawDescData
}

var file_deliveryareas_v1_domain_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_deliveryareas_v1_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_deliveryareas_v1_domain_proto_goTypes = []interface{}{
	(VendorFilterOperator)(0),      // 0: deliveryareas.v1.VendorFilterOperator
	(VendorFilterField)(0),         // 1: deliveryareas.v1.VendorFilterField
	(*VendorFilters)(nil),          // 2: deliveryareas.v1.VendorFilters
	(*VendorFilterExpression)(nil), // 3: deliveryareas.v1.VendorFilterExpression
	(*VendorFilterCondition)(nil),  // 4: deliveryareas.v1.VendorFilterCondition
}
var file_deliveryareas_v1_domain_proto_depIdxs = []int32{
	3, // 0: deliveryareas.v1.VendorFilters.expressions:type_name -> deliveryareas.v1.VendorFilterExpression
	4, // 1: deliveryareas.v1.VendorFilterExpression.conditions:type_name -> deliveryareas.v1.VendorFilterCondition
	0, // 2: deliveryareas.v1.VendorFilterCondition.operator:type_name -> deliveryareas.v1.VendorFilterOperator
	1, // 3: deliveryareas.v1.VendorFilterCondition.field:type_name -> deliveryareas.v1.VendorFilterField
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_deliveryareas_v1_domain_proto_init() }
func file_deliveryareas_v1_domain_proto_init() {
	if File_deliveryareas_v1_domain_proto != nil {
		return
	}
	file_deliveryareas_v1_dsl_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_deliveryareas_v1_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VendorFilters); i {
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
		file_deliveryareas_v1_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VendorFilterExpression); i {
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
		file_deliveryareas_v1_domain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VendorFilterCondition); i {
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
			RawDescriptor: file_deliveryareas_v1_domain_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deliveryareas_v1_domain_proto_goTypes,
		DependencyIndexes: file_deliveryareas_v1_domain_proto_depIdxs,
		EnumInfos:         file_deliveryareas_v1_domain_proto_enumTypes,
		MessageInfos:      file_deliveryareas_v1_domain_proto_msgTypes,
	}.Build()
	File_deliveryareas_v1_domain_proto = out.File
	file_deliveryareas_v1_domain_proto_rawDesc = nil
	file_deliveryareas_v1_domain_proto_goTypes = nil
	file_deliveryareas_v1_domain_proto_depIdxs = nil
}
