// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: teleport/auditlog/v1/auditlog.proto

package auditlogv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Order controls the ordering of objects returned based on the timestamp field.
type Order int32

const (
	Order_ORDER_DESCENDING_UNSPECIFIED Order = 0
	Order_ORDER_ASCENDING              Order = 1
)

// Enum value maps for Order.
var (
	Order_name = map[int32]string{
		0: "ORDER_DESCENDING_UNSPECIFIED",
		1: "ORDER_ASCENDING",
	}
	Order_value = map[string]int32{
		"ORDER_DESCENDING_UNSPECIFIED": 0,
		"ORDER_ASCENDING":              1,
	}
)

func (x Order) Enum() *Order {
	p := new(Order)
	*p = x
	return p
}

func (x Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_auditlog_v1_auditlog_proto_enumTypes[0].Descriptor()
}

func (Order) Type() protoreflect.EnumType {
	return &file_teleport_auditlog_v1_auditlog_proto_enumTypes[0]
}

func (x Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order.Descriptor instead.
func (Order) EnumDescriptor() ([]byte, []int) {
	return file_teleport_auditlog_v1_auditlog_proto_rawDescGZIP(), []int{0}
}

// StreamUnstructuredSessionEventsRequest is a request containing data needed to fetch a session recording.
type StreamUnstructuredSessionEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// session_id is the ID for a given session in an UUIDv4 format.
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// start_index is the index of the event to resume the stream after.
	// A start_index of 0 creates a new stream.
	StartIndex int32 `protobuf:"varint,2,opt,name=start_index,json=startIndex,proto3" json:"start_index,omitempty"`
}

func (x *StreamUnstructuredSessionEventsRequest) Reset() {
	*x = StreamUnstructuredSessionEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamUnstructuredSessionEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamUnstructuredSessionEventsRequest) ProtoMessage() {}

func (x *StreamUnstructuredSessionEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamUnstructuredSessionEventsRequest.ProtoReflect.Descriptor instead.
func (*StreamUnstructuredSessionEventsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_auditlog_v1_auditlog_proto_rawDescGZIP(), []int{0}
}

func (x *StreamUnstructuredSessionEventsRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *StreamUnstructuredSessionEventsRequest) GetStartIndex() int32 {
	if x != nil {
		return x.StartIndex
	}
	return 0
}

// GetUnstructuredEventsRequest is a request with the needed data to fetch events.
type GetUnstructuredEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// namespace, if not set, defaults to 'default'.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// start_date is the oldest date of returned events.
	StartDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// end_date is the newest date of returned events.
	EndDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// event_types is optional, if not set, returns all events.
	EventTypes []string `protobuf:"bytes,4,rep,name=event_types,json=eventTypes,proto3" json:"event_types,omitempty"`
	// limit is the maximum amount of events returned.
	Limit int32 `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	// start_key is used to resume a query in order to enable pagination.
	// If the previous response had LastKey set then this should be
	// set to its value. Otherwise leave empty.
	StartKey string `protobuf:"bytes,6,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	// order specifies an ascending or descending order of events.
	// A value of 0 means a descending order and a value of 1 means an ascending order.
	Order Order `protobuf:"varint,7,opt,name=order,proto3,enum=teleport.auditlog.v1.Order" json:"order,omitempty"`
}

func (x *GetUnstructuredEventsRequest) Reset() {
	*x = GetUnstructuredEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnstructuredEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnstructuredEventsRequest) ProtoMessage() {}

func (x *GetUnstructuredEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnstructuredEventsRequest.ProtoReflect.Descriptor instead.
func (*GetUnstructuredEventsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_auditlog_v1_auditlog_proto_rawDescGZIP(), []int{1}
}

func (x *GetUnstructuredEventsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetUnstructuredEventsRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *GetUnstructuredEventsRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *GetUnstructuredEventsRequest) GetEventTypes() []string {
	if x != nil {
		return x.EventTypes
	}
	return nil
}

func (x *GetUnstructuredEventsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUnstructuredEventsRequest) GetStartKey() string {
	if x != nil {
		return x.StartKey
	}
	return ""
}

func (x *GetUnstructuredEventsRequest) GetOrder() Order {
	if x != nil {
		return x.Order
	}
	return Order_ORDER_DESCENDING_UNSPECIFIED
}

// EventsUnstructured represents a list of events.AuditEvent in an unstructured format.
type EventsUnstructured struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// items is a list of unstructured formatted audit events.
	Items []*EventUnstructured `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// last_key is the key of the last event if the returned set did not contain all events found i.e limit <
	// actual amount. this is the key clients can supply in another API request to continue fetching
	// events from the previous last position.
	LastKey string `protobuf:"bytes,2,opt,name=last_key,json=lastKey,proto3" json:"last_key,omitempty"`
}

func (x *EventsUnstructured) Reset() {
	*x = EventsUnstructured{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsUnstructured) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsUnstructured) ProtoMessage() {}

func (x *EventsUnstructured) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsUnstructured.ProtoReflect.Descriptor instead.
func (*EventsUnstructured) Descriptor() ([]byte, []int) {
	return file_teleport_auditlog_v1_auditlog_proto_rawDescGZIP(), []int{2}
}

func (x *EventsUnstructured) GetItems() []*EventUnstructured {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *EventsUnstructured) GetLastKey() string {
	if x != nil {
		return x.LastKey
	}
	return ""
}

// ExportUnstructuredEventsRequest is a request with the needed data to export events.
type ExportUnstructuredEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start_date is the oldest date of returned events.
	StartDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// end_date is the newest date of returned events.
	EndDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// cursor is an optional mechanism to resume interrupted streams.
	Cursor string `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ExportUnstructuredEventsRequest) Reset() {
	*x = ExportUnstructuredEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportUnstructuredEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportUnstructuredEventsRequest) ProtoMessage() {}

func (x *ExportUnstructuredEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportUnstructuredEventsRequest.ProtoReflect.Descriptor instead.
func (*ExportUnstructuredEventsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_auditlog_v1_auditlog_proto_rawDescGZIP(), []int{3}
}

func (x *ExportUnstructuredEventsRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *ExportUnstructuredEventsRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *ExportUnstructuredEventsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// ExportEventUnstructured is the stream item of the ExportUnstructuredEvents method.
type ExportEventUnstructured struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// event is the unstructured representation of the event payload.
	Event *EventUnstructured `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	// cursor is the cursor to resume the stream after this point.
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ExportEventUnstructured) Reset() {
	*x = ExportEventUnstructured{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportEventUnstructured) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportEventUnstructured) ProtoMessage() {}

func (x *ExportEventUnstructured) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportEventUnstructured.ProtoReflect.Descriptor instead.
func (*ExportEventUnstructured) Descriptor() ([]byte, []int) {
	return file_teleport_auditlog_v1_auditlog_proto_rawDescGZIP(), []int{4}
}

func (x *ExportEventUnstructured) GetEvent() *EventUnstructured {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *ExportEventUnstructured) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// EventUnstructured represents a single events.AuditEvent in an unstructured format.
type EventUnstructured struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type is the type of the event.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// id is the unique ID of the event.
	// If the underlying event defines an ID, it will be used, otherwise
	// it is a SHA256 hash of the event payload.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// time is the time when the event was generated.
	Time *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	// index is the index of the event.
	Index int64 `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	// unstructured is the unstructured representation of the event payload.
	Unstructured *structpb.Struct `protobuf:"bytes,5,opt,name=unstructured,proto3" json:"unstructured,omitempty"`
}

func (x *EventUnstructured) Reset() {
	*x = EventUnstructured{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventUnstructured) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventUnstructured) ProtoMessage() {}

func (x *EventUnstructured) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_auditlog_v1_auditlog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventUnstructured.ProtoReflect.Descriptor instead.
func (*EventUnstructured) Descriptor() ([]byte, []int) {
	return file_teleport_auditlog_v1_auditlog_proto_rawDescGZIP(), []int{5}
}

func (x *EventUnstructured) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EventUnstructured) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventUnstructured) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *EventUnstructured) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *EventUnstructured) GetUnstructured() *structpb.Struct {
	if x != nil {
		return x.Unstructured
	}
	return nil
}

var File_teleport_auditlog_v1_auditlog_proto protoreflect.FileDescriptor

var file_teleport_auditlog_v1_auditlog_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x26, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x22, 0xb5, 0x02, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x6e, 0x0a, 0x12,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x64, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x22, 0xab, 0x01, 0x0a,
	0x1f, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x70, 0x0a, 0x17, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x64, 0x12, 0x3d, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0xba, 0x01, 0x0a,
	0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x3b, 0x0a, 0x0c,
	0x75, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x75, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x2a, 0x3e, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x43,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x53,
	0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x32, 0x9a, 0x03, 0x0a, 0x0f, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x01,
	0x0a, 0x1f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x3c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x30, 0x01, 0x12, 0x75, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x32, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x64, 0x12, 0x82, 0x01, 0x0a, 0x18, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x35,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x64, 0x30, 0x01, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2f, 0x76,
	0x31, 0x3b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_auditlog_v1_auditlog_proto_rawDescOnce sync.Once
	file_teleport_auditlog_v1_auditlog_proto_rawDescData = file_teleport_auditlog_v1_auditlog_proto_rawDesc
)

func file_teleport_auditlog_v1_auditlog_proto_rawDescGZIP() []byte {
	file_teleport_auditlog_v1_auditlog_proto_rawDescOnce.Do(func() {
		file_teleport_auditlog_v1_auditlog_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_auditlog_v1_auditlog_proto_rawDescData)
	})
	return file_teleport_auditlog_v1_auditlog_proto_rawDescData
}

var file_teleport_auditlog_v1_auditlog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_teleport_auditlog_v1_auditlog_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_teleport_auditlog_v1_auditlog_proto_goTypes = []any{
	(Order)(0), // 0: teleport.auditlog.v1.Order
	(*StreamUnstructuredSessionEventsRequest)(nil), // 1: teleport.auditlog.v1.StreamUnstructuredSessionEventsRequest
	(*GetUnstructuredEventsRequest)(nil),           // 2: teleport.auditlog.v1.GetUnstructuredEventsRequest
	(*EventsUnstructured)(nil),                     // 3: teleport.auditlog.v1.EventsUnstructured
	(*ExportUnstructuredEventsRequest)(nil),        // 4: teleport.auditlog.v1.ExportUnstructuredEventsRequest
	(*ExportEventUnstructured)(nil),                // 5: teleport.auditlog.v1.ExportEventUnstructured
	(*EventUnstructured)(nil),                      // 6: teleport.auditlog.v1.EventUnstructured
	(*timestamppb.Timestamp)(nil),                  // 7: google.protobuf.Timestamp
	(*structpb.Struct)(nil),                        // 8: google.protobuf.Struct
}
var file_teleport_auditlog_v1_auditlog_proto_depIdxs = []int32{
	7,  // 0: teleport.auditlog.v1.GetUnstructuredEventsRequest.start_date:type_name -> google.protobuf.Timestamp
	7,  // 1: teleport.auditlog.v1.GetUnstructuredEventsRequest.end_date:type_name -> google.protobuf.Timestamp
	0,  // 2: teleport.auditlog.v1.GetUnstructuredEventsRequest.order:type_name -> teleport.auditlog.v1.Order
	6,  // 3: teleport.auditlog.v1.EventsUnstructured.items:type_name -> teleport.auditlog.v1.EventUnstructured
	7,  // 4: teleport.auditlog.v1.ExportUnstructuredEventsRequest.start_date:type_name -> google.protobuf.Timestamp
	7,  // 5: teleport.auditlog.v1.ExportUnstructuredEventsRequest.end_date:type_name -> google.protobuf.Timestamp
	6,  // 6: teleport.auditlog.v1.ExportEventUnstructured.event:type_name -> teleport.auditlog.v1.EventUnstructured
	7,  // 7: teleport.auditlog.v1.EventUnstructured.time:type_name -> google.protobuf.Timestamp
	8,  // 8: teleport.auditlog.v1.EventUnstructured.unstructured:type_name -> google.protobuf.Struct
	1,  // 9: teleport.auditlog.v1.AuditLogService.StreamUnstructuredSessionEvents:input_type -> teleport.auditlog.v1.StreamUnstructuredSessionEventsRequest
	2,  // 10: teleport.auditlog.v1.AuditLogService.GetUnstructuredEvents:input_type -> teleport.auditlog.v1.GetUnstructuredEventsRequest
	4,  // 11: teleport.auditlog.v1.AuditLogService.ExportUnstructuredEvents:input_type -> teleport.auditlog.v1.ExportUnstructuredEventsRequest
	6,  // 12: teleport.auditlog.v1.AuditLogService.StreamUnstructuredSessionEvents:output_type -> teleport.auditlog.v1.EventUnstructured
	3,  // 13: teleport.auditlog.v1.AuditLogService.GetUnstructuredEvents:output_type -> teleport.auditlog.v1.EventsUnstructured
	5,  // 14: teleport.auditlog.v1.AuditLogService.ExportUnstructuredEvents:output_type -> teleport.auditlog.v1.ExportEventUnstructured
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_teleport_auditlog_v1_auditlog_proto_init() }
func file_teleport_auditlog_v1_auditlog_proto_init() {
	if File_teleport_auditlog_v1_auditlog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_auditlog_v1_auditlog_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StreamUnstructuredSessionEventsRequest); i {
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
		file_teleport_auditlog_v1_auditlog_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetUnstructuredEventsRequest); i {
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
		file_teleport_auditlog_v1_auditlog_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EventsUnstructured); i {
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
		file_teleport_auditlog_v1_auditlog_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ExportUnstructuredEventsRequest); i {
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
		file_teleport_auditlog_v1_auditlog_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ExportEventUnstructured); i {
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
		file_teleport_auditlog_v1_auditlog_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*EventUnstructured); i {
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
			RawDescriptor: file_teleport_auditlog_v1_auditlog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_auditlog_v1_auditlog_proto_goTypes,
		DependencyIndexes: file_teleport_auditlog_v1_auditlog_proto_depIdxs,
		EnumInfos:         file_teleport_auditlog_v1_auditlog_proto_enumTypes,
		MessageInfos:      file_teleport_auditlog_v1_auditlog_proto_msgTypes,
	}.Build()
	File_teleport_auditlog_v1_auditlog_proto = out.File
	file_teleport_auditlog_v1_auditlog_proto_rawDesc = nil
	file_teleport_auditlog_v1_auditlog_proto_goTypes = nil
	file_teleport_auditlog_v1_auditlog_proto_depIdxs = nil
}
