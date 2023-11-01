// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: dynamicpricing/v1/workspace_optimization.proto

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

// Possible status codes for the optimization response.
type StatusCode int32

const (
	StatusCode_STATUS_CODE_INVALID                   StatusCode = 0
	StatusCode_STATUS_CODE_SUCCESS                   StatusCode = 1
	StatusCode_STATUS_CODE_ERROR_UNKNOWN             StatusCode = 2
	StatusCode_STATUS_CODE_ERROR_NOT_ENOUGH_ORDERS   StatusCode = 3
	StatusCode_STATUS_CODE_ERROR_INVALID_PAYLOAD     StatusCode = 4
	StatusCode_STATUS_CODE_ERROR_BQ_EXCEPTION        StatusCode = 5
	StatusCode_STATUS_CODE_ERROR_TIME_OUT            StatusCode = 6
	StatusCode_STATUS_CODE_ERROR_NO_DATA_IN_TABLE    StatusCode = 7
	StatusCode_STATUS_CODE_ERROR_BQ_NOT_SYNCHRONIZED StatusCode = 8
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "STATUS_CODE_INVALID",
		1: "STATUS_CODE_SUCCESS",
		2: "STATUS_CODE_ERROR_UNKNOWN",
		3: "STATUS_CODE_ERROR_NOT_ENOUGH_ORDERS",
		4: "STATUS_CODE_ERROR_INVALID_PAYLOAD",
		5: "STATUS_CODE_ERROR_BQ_EXCEPTION",
		6: "STATUS_CODE_ERROR_TIME_OUT",
		7: "STATUS_CODE_ERROR_NO_DATA_IN_TABLE",
		8: "STATUS_CODE_ERROR_BQ_NOT_SYNCHRONIZED",
	}
	StatusCode_value = map[string]int32{
		"STATUS_CODE_INVALID":                   0,
		"STATUS_CODE_SUCCESS":                   1,
		"STATUS_CODE_ERROR_UNKNOWN":             2,
		"STATUS_CODE_ERROR_NOT_ENOUGH_ORDERS":   3,
		"STATUS_CODE_ERROR_INVALID_PAYLOAD":     4,
		"STATUS_CODE_ERROR_BQ_EXCEPTION":        5,
		"STATUS_CODE_ERROR_TIME_OUT":            6,
		"STATUS_CODE_ERROR_NO_DATA_IN_TABLE":    7,
		"STATUS_CODE_ERROR_BQ_NOT_SYNCHRONIZED": 8,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_dynamicpricing_v1_workspace_optimization_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_dynamicpricing_v1_workspace_optimization_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_workspace_optimization_proto_rawDescGZIP(), []int{0}
}

// DPS order share based request optimization payload.
type OrderShareOptimizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptimizationId int64  `protobuf:"varint,1,opt,name=optimization_id,json=optimizationId,proto3" json:"optimization_id,omitempty"`
	GlobalEntityId string `protobuf:"bytes,2,opt,name=global_entity_id,json=globalEntityId,proto3" json:"global_entity_id,omitempty"`
	// Unique identifier of logistics environment, it is used to identify each different deployment of the same application.
	// Eg: a4 and at points to the Austria environment but both addresses different use-cases in the same country.
	LogisticsEnvironmentId string `protobuf:"bytes,3,opt,name=logistics_environment_id,json=logisticsEnvironmentId,proto3" json:"logistics_environment_id,omitempty"`
	// Date must be in the format YYYY-MM-DD
	StartDate string `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// Date must be in the format YYYY-MM-DD
	EndDate             string                                             `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	VendorIds           []string                                           `protobuf:"bytes,6,rep,name=vendor_ids,json=vendorIds,proto3" json:"vendor_ids,omitempty"`
	TravelTimeComponent *OrderShareOptimizationRequest_TravelTimeComponent `protobuf:"bytes,7,opt,name=travel_time_component,json=travelTimeComponent,proto3" json:"travel_time_component,omitempty"`
}

func (x *OrderShareOptimizationRequest) Reset() {
	*x = OrderShareOptimizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderShareOptimizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderShareOptimizationRequest) ProtoMessage() {}

func (x *OrderShareOptimizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderShareOptimizationRequest.ProtoReflect.Descriptor instead.
func (*OrderShareOptimizationRequest) Descriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_workspace_optimization_proto_rawDescGZIP(), []int{0}
}

func (x *OrderShareOptimizationRequest) GetOptimizationId() int64 {
	if x != nil {
		return x.OptimizationId
	}
	return 0
}

func (x *OrderShareOptimizationRequest) GetGlobalEntityId() string {
	if x != nil {
		return x.GlobalEntityId
	}
	return ""
}

func (x *OrderShareOptimizationRequest) GetLogisticsEnvironmentId() string {
	if x != nil {
		return x.LogisticsEnvironmentId
	}
	return ""
}

func (x *OrderShareOptimizationRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *OrderShareOptimizationRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *OrderShareOptimizationRequest) GetVendorIds() []string {
	if x != nil {
		return x.VendorIds
	}
	return nil
}

func (x *OrderShareOptimizationRequest) GetTravelTimeComponent() *OrderShareOptimizationRequest_TravelTimeComponent {
	if x != nil {
		return x.TravelTimeComponent
	}
	return nil
}

// DPS order share tier.
type OrderShareTier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CumulativeOrderShare int32 `protobuf:"varint,1,opt,name=cumulative_order_share,json=cumulativeOrderShare,proto3" json:"cumulative_order_share,omitempty"`
	// Deprecated: Marked as deprecated in dynamicpricing/v1/workspace_optimization.proto.
	CumulativeOrderShareDouble float64 `protobuf:"fixed64,2,opt,name=cumulative_order_share_double,json=cumulativeOrderShareDouble,proto3" json:"cumulative_order_share_double,omitempty"`
}

func (x *OrderShareTier) Reset() {
	*x = OrderShareTier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderShareTier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderShareTier) ProtoMessage() {}

func (x *OrderShareTier) ProtoReflect() protoreflect.Message {
	mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderShareTier.ProtoReflect.Descriptor instead.
func (*OrderShareTier) Descriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_workspace_optimization_proto_rawDescGZIP(), []int{1}
}

func (x *OrderShareTier) GetCumulativeOrderShare() int32 {
	if x != nil {
		return x.CumulativeOrderShare
	}
	return 0
}

// Deprecated: Marked as deprecated in dynamicpricing/v1/workspace_optimization.proto.
func (x *OrderShareTier) GetCumulativeOrderShareDouble() float64 {
	if x != nil {
		return x.CumulativeOrderShareDouble
	}
	return 0
}

// DPS order share based response optimization payload.
type OrderShareOptimizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptimizationId int64  `protobuf:"varint,1,opt,name=optimization_id,json=optimizationId,proto3" json:"optimization_id,omitempty"`
	GlobalEntityId string `protobuf:"bytes,2,opt,name=global_entity_id,json=globalEntityId,proto3" json:"global_entity_id,omitempty"`
	// Unique identifier of logistics environment, it is used to identify each different deployment of the same application.
	// Eg: a4 and at points to the Austria environment but both addresses different use-cases in the same country.
	LogisticsEnvironmentId string               `protobuf:"bytes,3,opt,name=logistics_environment_id,json=logisticsEnvironmentId,proto3" json:"logistics_environment_id,omitempty"`
	Status                 *OptimizationStatus  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	TravelTimeComponent    *TravelTimeComponent `protobuf:"bytes,5,opt,name=travel_time_component,json=travelTimeComponent,proto3" json:"travel_time_component,omitempty"`
}

func (x *OrderShareOptimizationResponse) Reset() {
	*x = OrderShareOptimizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderShareOptimizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderShareOptimizationResponse) ProtoMessage() {}

func (x *OrderShareOptimizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderShareOptimizationResponse.ProtoReflect.Descriptor instead.
func (*OrderShareOptimizationResponse) Descriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_workspace_optimization_proto_rawDescGZIP(), []int{2}
}

func (x *OrderShareOptimizationResponse) GetOptimizationId() int64 {
	if x != nil {
		return x.OptimizationId
	}
	return 0
}

func (x *OrderShareOptimizationResponse) GetGlobalEntityId() string {
	if x != nil {
		return x.GlobalEntityId
	}
	return ""
}

func (x *OrderShareOptimizationResponse) GetLogisticsEnvironmentId() string {
	if x != nil {
		return x.LogisticsEnvironmentId
	}
	return ""
}

func (x *OrderShareOptimizationResponse) GetStatus() *OptimizationStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *OrderShareOptimizationResponse) GetTravelTimeComponent() *TravelTimeComponent {
	if x != nil {
		return x.TravelTimeComponent
	}
	return nil
}

// DPS order share based optimized travel time configuration.
type OptimizationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=dynamicpricing.v1.StatusCode" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OptimizationStatus) Reset() {
	*x = OptimizationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptimizationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptimizationStatus) ProtoMessage() {}

func (x *OptimizationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptimizationStatus.ProtoReflect.Descriptor instead.
func (*OptimizationStatus) Descriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_workspace_optimization_proto_rawDescGZIP(), []int{3}
}

func (x *OptimizationStatus) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_STATUS_CODE_INVALID
}

func (x *OptimizationStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// DPS order share based optimized travel time configuration.
type TravelTimeComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TravelTimeTiers []*TravelTimeTier `protobuf:"bytes,1,rep,name=travel_time_tiers,json=travelTimeTiers,proto3" json:"travel_time_tiers,omitempty"`
}

func (x *TravelTimeComponent) Reset() {
	*x = TravelTimeComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelTimeComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelTimeComponent) ProtoMessage() {}

func (x *TravelTimeComponent) ProtoReflect() protoreflect.Message {
	mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelTimeComponent.ProtoReflect.Descriptor instead.
func (*TravelTimeComponent) Descriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_workspace_optimization_proto_rawDescGZIP(), []int{4}
}

func (x *TravelTimeComponent) GetTravelTimeTiers() []*TravelTimeTier {
	if x != nil {
		return x.TravelTimeTiers
	}
	return nil
}

// DPS order share based optimized travel time tier.
type TravelTimeTier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TravelTimeThresholdSeconds int32 `protobuf:"varint,1,opt,name=travel_time_threshold_seconds,json=travelTimeThresholdSeconds,proto3" json:"travel_time_threshold_seconds,omitempty"`
}

func (x *TravelTimeTier) Reset() {
	*x = TravelTimeTier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelTimeTier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelTimeTier) ProtoMessage() {}

func (x *TravelTimeTier) ProtoReflect() protoreflect.Message {
	mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelTimeTier.ProtoReflect.Descriptor instead.
func (*TravelTimeTier) Descriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_workspace_optimization_proto_rawDescGZIP(), []int{5}
}

func (x *TravelTimeTier) GetTravelTimeThresholdSeconds() int32 {
	if x != nil {
		return x.TravelTimeThresholdSeconds
	}
	return 0
}

// DPS order share input tiers.
type OrderShareOptimizationRequest_TravelTimeComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderShareTiers []*OrderShareTier `protobuf:"bytes,1,rep,name=order_share_tiers,json=orderShareTiers,proto3" json:"order_share_tiers,omitempty"`
}

func (x *OrderShareOptimizationRequest_TravelTimeComponent) Reset() {
	*x = OrderShareOptimizationRequest_TravelTimeComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderShareOptimizationRequest_TravelTimeComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderShareOptimizationRequest_TravelTimeComponent) ProtoMessage() {}

func (x *OrderShareOptimizationRequest_TravelTimeComponent) ProtoReflect() protoreflect.Message {
	mi := &file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderShareOptimizationRequest_TravelTimeComponent.ProtoReflect.Descriptor instead.
func (*OrderShareOptimizationRequest_TravelTimeComponent) Descriptor() ([]byte, []int) {
	return file_dynamicpricing_v1_workspace_optimization_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OrderShareOptimizationRequest_TravelTimeComponent) GetOrderShareTiers() []*OrderShareTier {
	if x != nil {
		return x.OrderShareTiers
	}
	return nil
}

var File_dynamicpricing_v1_workspace_optimization_proto protoreflect.FileDescriptor

var file_dynamicpricing_v1_workspace_optimization_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x05, 0x0a,
	0x1d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6d,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x46, 0x0a, 0x10, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xfa, 0x42, 0x19, 0x72,
	0x17, 0x32, 0x15, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x5d, 0x7b, 0x32, 0x2c, 0x33, 0x7d, 0x5f, 0x5b,
	0x41, 0x2d, 0x5a, 0x5d, 0x7b, 0x32, 0x7d, 0x24, 0x52, 0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x57, 0x0a, 0x18, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xfa, 0x42, 0x1a, 0x72,
	0x18, 0x32, 0x16, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x7b, 0x31, 0x2c, 0x32, 0x7d, 0x5b, 0x30,
	0x2d, 0x39, 0x5d, 0x7b, 0x30, 0x2c, 0x31, 0x7d, 0x24, 0x52, 0x16, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x55, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xfa, 0x42, 0x33, 0x72, 0x31, 0x32, 0x2f, 0x5c, 0x64,
	0x7b, 0x34, 0x7d, 0x2d, 0x28, 0x30, 0x5b, 0x31, 0x2d, 0x39, 0x5d, 0x7c, 0x31, 0x5b, 0x30, 0x31,
	0x32, 0x5d, 0x29, 0x2d, 0x28, 0x30, 0x5b, 0x31, 0x2d, 0x39, 0x5d, 0x7c, 0x5b, 0x31, 0x32, 0x5d,
	0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7c, 0x33, 0x5b, 0x30, 0x31, 0x5d, 0x29, 0x24, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xfa, 0x42, 0x33, 0x72,
	0x31, 0x32, 0x2f, 0x5c, 0x64, 0x7b, 0x34, 0x7d, 0x2d, 0x28, 0x30, 0x5b, 0x31, 0x2d, 0x39, 0x5d,
	0x7c, 0x31, 0x5b, 0x30, 0x31, 0x32, 0x5d, 0x29, 0x2d, 0x28, 0x30, 0x5b, 0x31, 0x2d, 0x39, 0x5d,
	0x7c, 0x5b, 0x31, 0x32, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7c, 0x33, 0x5b, 0x30, 0x31, 0x5d,
	0x29, 0x24, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0a, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x49, 0x64, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72,
	0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x13, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x1a, 0x6e, 0x0a, 0x13, 0x54, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x12, 0x57, 0x0a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f,
	0x74, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x54, 0x69, 0x65, 0x72, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x54, 0x69, 0x65, 0x72, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x0e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x54, 0x69, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x16,
	0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x1a, 0x04, 0x18, 0x64, 0x20, 0x00, 0x52, 0x14, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x5c, 0x0a,
	0x1d, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x42, 0x19, 0xfa, 0x42, 0x14, 0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x59, 0x40, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x18, 0x01, 0x52,
	0x1a, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x22, 0x98, 0x03, 0x0a, 0x1e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x46, 0x0a, 0x10, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xfa, 0x42, 0x19, 0x72,
	0x17, 0x32, 0x15, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x5d, 0x7b, 0x32, 0x2c, 0x33, 0x7d, 0x5f, 0x5b,
	0x41, 0x2d, 0x5a, 0x5d, 0x7b, 0x32, 0x7d, 0x24, 0x52, 0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x57, 0x0a, 0x18, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xfa, 0x42, 0x1a, 0x72,
	0x18, 0x32, 0x16, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x7b, 0x31, 0x2c, 0x32, 0x7d, 0x5b, 0x30,
	0x2d, 0x39, 0x5d, 0x7b, 0x30, 0x2c, 0x31, 0x7d, 0x24, 0x52, 0x16, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5a, 0x0a, 0x15, 0x74, 0x72,
	0x61, 0x76, 0x65, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x52, 0x13, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x22, 0x61, 0x0a, 0x12, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x64, 0x0a, 0x13, 0x54, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x12, 0x4d, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x74, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x65, 0x72, 0x52, 0x0f,
	0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x65, 0x72, 0x73, 0x22,
	0x53, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x65,
	0x72, 0x12, 0x41, 0x0a, 0x1d, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1a, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x2a, 0xc4, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x02, 0x12, 0x27, 0x0a, 0x23, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e,
	0x4f, 0x55, 0x47, 0x48, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x53, 0x10, 0x03, 0x12, 0x25, 0x0a,
	0x21, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x59, 0x4c, 0x4f,
	0x41, 0x44, 0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x42, 0x51, 0x5f, 0x45, 0x58, 0x43,
	0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x06, 0x12, 0x26, 0x0a, 0x22, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4e, 0x4f,
	0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x49, 0x4e, 0x5f, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x07,
	0x12, 0x29, 0x0a, 0x25, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x42, 0x51, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x59, 0x4e,
	0x43, 0x48, 0x52, 0x4f, 0x4e, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x08, 0x42, 0x93, 0x01, 0x0a, 0x35,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x68, 0x65, 0x72, 0x6f,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x1a, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x68, 0x65, 0x72, 0x6f, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dynamicpricing_v1_workspace_optimization_proto_rawDescOnce sync.Once
	file_dynamicpricing_v1_workspace_optimization_proto_rawDescData = file_dynamicpricing_v1_workspace_optimization_proto_rawDesc
)

func file_dynamicpricing_v1_workspace_optimization_proto_rawDescGZIP() []byte {
	file_dynamicpricing_v1_workspace_optimization_proto_rawDescOnce.Do(func() {
		file_dynamicpricing_v1_workspace_optimization_proto_rawDescData = protoimpl.X.CompressGZIP(file_dynamicpricing_v1_workspace_optimization_proto_rawDescData)
	})
	return file_dynamicpricing_v1_workspace_optimization_proto_rawDescData
}

var file_dynamicpricing_v1_workspace_optimization_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dynamicpricing_v1_workspace_optimization_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_dynamicpricing_v1_workspace_optimization_proto_goTypes = []interface{}{
	(StatusCode)(0),                                           // 0: dynamicpricing.v1.StatusCode
	(*OrderShareOptimizationRequest)(nil),                     // 1: dynamicpricing.v1.OrderShareOptimizationRequest
	(*OrderShareTier)(nil),                                    // 2: dynamicpricing.v1.OrderShareTier
	(*OrderShareOptimizationResponse)(nil),                    // 3: dynamicpricing.v1.OrderShareOptimizationResponse
	(*OptimizationStatus)(nil),                                // 4: dynamicpricing.v1.OptimizationStatus
	(*TravelTimeComponent)(nil),                               // 5: dynamicpricing.v1.TravelTimeComponent
	(*TravelTimeTier)(nil),                                    // 6: dynamicpricing.v1.TravelTimeTier
	(*OrderShareOptimizationRequest_TravelTimeComponent)(nil), // 7: dynamicpricing.v1.OrderShareOptimizationRequest.TravelTimeComponent
}
var file_dynamicpricing_v1_workspace_optimization_proto_depIdxs = []int32{
	7, // 0: dynamicpricing.v1.OrderShareOptimizationRequest.travel_time_component:type_name -> dynamicpricing.v1.OrderShareOptimizationRequest.TravelTimeComponent
	4, // 1: dynamicpricing.v1.OrderShareOptimizationResponse.status:type_name -> dynamicpricing.v1.OptimizationStatus
	5, // 2: dynamicpricing.v1.OrderShareOptimizationResponse.travel_time_component:type_name -> dynamicpricing.v1.TravelTimeComponent
	0, // 3: dynamicpricing.v1.OptimizationStatus.code:type_name -> dynamicpricing.v1.StatusCode
	6, // 4: dynamicpricing.v1.TravelTimeComponent.travel_time_tiers:type_name -> dynamicpricing.v1.TravelTimeTier
	2, // 5: dynamicpricing.v1.OrderShareOptimizationRequest.TravelTimeComponent.order_share_tiers:type_name -> dynamicpricing.v1.OrderShareTier
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_dynamicpricing_v1_workspace_optimization_proto_init() }
func file_dynamicpricing_v1_workspace_optimization_proto_init() {
	if File_dynamicpricing_v1_workspace_optimization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderShareOptimizationRequest); i {
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
		file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderShareTier); i {
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
		file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderShareOptimizationResponse); i {
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
		file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptimizationStatus); i {
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
		file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelTimeComponent); i {
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
		file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelTimeTier); i {
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
		file_dynamicpricing_v1_workspace_optimization_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderShareOptimizationRequest_TravelTimeComponent); i {
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
			RawDescriptor: file_dynamicpricing_v1_workspace_optimization_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dynamicpricing_v1_workspace_optimization_proto_goTypes,
		DependencyIndexes: file_dynamicpricing_v1_workspace_optimization_proto_depIdxs,
		EnumInfos:         file_dynamicpricing_v1_workspace_optimization_proto_enumTypes,
		MessageInfos:      file_dynamicpricing_v1_workspace_optimization_proto_msgTypes,
	}.Build()
	File_dynamicpricing_v1_workspace_optimization_proto = out.File
	file_dynamicpricing_v1_workspace_optimization_proto_rawDesc = nil
	file_dynamicpricing_v1_workspace_optimization_proto_goTypes = nil
	file_dynamicpricing_v1_workspace_optimization_proto_depIdxs = nil
}
