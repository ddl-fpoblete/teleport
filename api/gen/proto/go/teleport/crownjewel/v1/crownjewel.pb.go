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
// source: teleport/crownjewel/v1/crownjewel.proto

package crownjewelv1

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	v11 "github.com/gravitational/teleport/api/gen/proto/go/teleport/label/v1"
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

// CrownJewel represents a Crown Jewel resource.
// Crown Jewel is a resource that represents a set of resources that are
// considered critical to the organization. Access Graph uses Crown Jewel to
// generate audit events if access to a resource has changed.
type CrownJewel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The kind of resource represented.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Mandatory field for all resources. Not populated for this resource type.
	SubKind string `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	// The version of the resource being represented.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Common metadata that all resources share.
	Metadata *v1.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec is the crown jewel spec.
	Spec *CrownJewelSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *CrownJewel) Reset() {
	*x = CrownJewel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrownJewel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrownJewel) ProtoMessage() {}

func (x *CrownJewel) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrownJewel.ProtoReflect.Descriptor instead.
func (*CrownJewel) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_proto_rawDescGZIP(), []int{0}
}

func (x *CrownJewel) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CrownJewel) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *CrownJewel) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CrownJewel) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CrownJewel) GetSpec() *CrownJewelSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// CrownJewelSpec is the specification of a Crown Jewel.
type CrownJewelSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TeleportMatchers is a list of teleport matchers.
	// DEPRECATED: Use query instead.
	TeleportMatchers []*TeleportMatcher `protobuf:"bytes,1,rep,name=teleport_matchers,json=teleportMatchers,proto3" json:"teleport_matchers,omitempty"`
	// AWSMatchers is a list of AWS matchers.
	// DEPRECATED: Use query instead.
	AwsMatchers []*AWSMatcher `protobuf:"bytes,2,rep,name=aws_matchers,json=awsMatchers,proto3" json:"aws_matchers,omitempty"`
	// Query is a Access Graph query to match resources.
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *CrownJewelSpec) Reset() {
	*x = CrownJewelSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrownJewelSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrownJewelSpec) ProtoMessage() {}

func (x *CrownJewelSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrownJewelSpec.ProtoReflect.Descriptor instead.
func (*CrownJewelSpec) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_proto_rawDescGZIP(), []int{1}
}

func (x *CrownJewelSpec) GetTeleportMatchers() []*TeleportMatcher {
	if x != nil {
		return x.TeleportMatchers
	}
	return nil
}

func (x *CrownJewelSpec) GetAwsMatchers() []*AWSMatcher {
	if x != nil {
		return x.AwsMatchers
	}
	return nil
}

func (x *CrownJewelSpec) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// TeleportMatcher represents a matcher for Teleport resources.
type TeleportMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kind is the kind of the resource: ssh, k8s, db, etc
	// Multiple kinds can be provided to match multiple kinds.
	Kinds []string `protobuf:"bytes,2,rep,name=kinds,proto3" json:"kinds,omitempty"`
	// Labels is a set of labels.
	Labels []*v11.Label `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	// Names are the name of resources. When the name is provided, it will match
	// resources with the same name.
	Names []string `protobuf:"bytes,4,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *TeleportMatcher) Reset() {
	*x = TeleportMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeleportMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeleportMatcher) ProtoMessage() {}

func (x *TeleportMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeleportMatcher.ProtoReflect.Descriptor instead.
func (*TeleportMatcher) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_proto_rawDescGZIP(), []int{2}
}

func (x *TeleportMatcher) GetKinds() []string {
	if x != nil {
		return x.Kinds
	}
	return nil
}

func (x *TeleportMatcher) GetLabels() []*v11.Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TeleportMatcher) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

// AWSMatcher represents a matcher for AWS resources.
// Those matchers are used only by Access Graph. Teleport related matchers are
// defined in the TeleportMatcher.
type AWSMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types are AWS database types to match, "ec2", "rds", "s3", etc
	Types []string `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
	// Regions are AWS regions to query for resources.
	Regions []string `protobuf:"bytes,2,rep,name=regions,proto3" json:"regions,omitempty"`
	// Tags are AWS resource Tags to match.
	// labels is a set of labels.
	Tags []*AWSTag `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	// ARNs are AWS resources ARN to match.
	Arns []string `protobuf:"bytes,5,rep,name=arns,proto3" json:"arns,omitempty"`
}

func (x *AWSMatcher) Reset() {
	*x = AWSMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSMatcher) ProtoMessage() {}

func (x *AWSMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSMatcher.ProtoReflect.Descriptor instead.
func (*AWSMatcher) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_proto_rawDescGZIP(), []int{3}
}

func (x *AWSMatcher) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *AWSMatcher) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *AWSMatcher) GetTags() []*AWSTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *AWSMatcher) GetArns() []string {
	if x != nil {
		return x.Arns
	}
	return nil
}

// AWSTag is a tag that is attached to an AWS resource.
type AWSTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is the key of the tag.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Value is the value of the tag.
	Values []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *AWSTag) Reset() {
	*x = AWSTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSTag) ProtoMessage() {}

func (x *AWSTag) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSTag.ProtoReflect.Descriptor instead.
func (*AWSTag) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_proto_rawDescGZIP(), []int{4}
}

func (x *AWSTag) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AWSTag) GetValues() []*wrapperspb.StringValue {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_teleport_crownjewel_v1_crownjewel_proto protoreflect.FileDescriptor

var file_teleport_crownjewel_v1_crownjewel_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x72, 0x6f, 0x77, 0x6e,
	0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65,
	0x77, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x4b, 0x69, 0x6e,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63,
	0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x6f,
	0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x22, 0xc3, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x54, 0x0a, 0x11, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x6e,
	0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x45, 0x0a, 0x0c, 0x61, 0x77,
	0x73, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77,
	0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x0b, 0x61, 0x77, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x7b, 0x0a, 0x0f, 0x54, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x69,
	0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73,
	0x12, 0x30, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x0a, 0x41, 0x57, 0x53, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f,
	0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x54, 0x61,
	0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x6e, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x6e, 0x73, 0x4a, 0x04, 0x08, 0x04, 0x10,
	0x05, 0x52, 0x03, 0x61, 0x72, 0x6e, 0x22, 0x50, 0x0a, 0x06, 0x41, 0x57, 0x53, 0x54, 0x61, 0x67,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x58, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77,
	0x65, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_crownjewel_v1_crownjewel_proto_rawDescOnce sync.Once
	file_teleport_crownjewel_v1_crownjewel_proto_rawDescData = file_teleport_crownjewel_v1_crownjewel_proto_rawDesc
)

func file_teleport_crownjewel_v1_crownjewel_proto_rawDescGZIP() []byte {
	file_teleport_crownjewel_v1_crownjewel_proto_rawDescOnce.Do(func() {
		file_teleport_crownjewel_v1_crownjewel_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_crownjewel_v1_crownjewel_proto_rawDescData)
	})
	return file_teleport_crownjewel_v1_crownjewel_proto_rawDescData
}

var file_teleport_crownjewel_v1_crownjewel_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_teleport_crownjewel_v1_crownjewel_proto_goTypes = []any{
	(*CrownJewel)(nil),             // 0: teleport.crownjewel.v1.CrownJewel
	(*CrownJewelSpec)(nil),         // 1: teleport.crownjewel.v1.CrownJewelSpec
	(*TeleportMatcher)(nil),        // 2: teleport.crownjewel.v1.TeleportMatcher
	(*AWSMatcher)(nil),             // 3: teleport.crownjewel.v1.AWSMatcher
	(*AWSTag)(nil),                 // 4: teleport.crownjewel.v1.AWSTag
	(*v1.Metadata)(nil),            // 5: teleport.header.v1.Metadata
	(*v11.Label)(nil),              // 6: teleport.label.v1.Label
	(*wrapperspb.StringValue)(nil), // 7: google.protobuf.StringValue
}
var file_teleport_crownjewel_v1_crownjewel_proto_depIdxs = []int32{
	5, // 0: teleport.crownjewel.v1.CrownJewel.metadata:type_name -> teleport.header.v1.Metadata
	1, // 1: teleport.crownjewel.v1.CrownJewel.spec:type_name -> teleport.crownjewel.v1.CrownJewelSpec
	2, // 2: teleport.crownjewel.v1.CrownJewelSpec.teleport_matchers:type_name -> teleport.crownjewel.v1.TeleportMatcher
	3, // 3: teleport.crownjewel.v1.CrownJewelSpec.aws_matchers:type_name -> teleport.crownjewel.v1.AWSMatcher
	6, // 4: teleport.crownjewel.v1.TeleportMatcher.labels:type_name -> teleport.label.v1.Label
	4, // 5: teleport.crownjewel.v1.AWSMatcher.tags:type_name -> teleport.crownjewel.v1.AWSTag
	7, // 6: teleport.crownjewel.v1.AWSTag.values:type_name -> google.protobuf.StringValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_teleport_crownjewel_v1_crownjewel_proto_init() }
func file_teleport_crownjewel_v1_crownjewel_proto_init() {
	if File_teleport_crownjewel_v1_crownjewel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CrownJewel); i {
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
		file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CrownJewelSpec); i {
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
		file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TeleportMatcher); i {
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
		file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AWSMatcher); i {
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
		file_teleport_crownjewel_v1_crownjewel_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AWSTag); i {
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
			RawDescriptor: file_teleport_crownjewel_v1_crownjewel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_crownjewel_v1_crownjewel_proto_goTypes,
		DependencyIndexes: file_teleport_crownjewel_v1_crownjewel_proto_depIdxs,
		MessageInfos:      file_teleport_crownjewel_v1_crownjewel_proto_msgTypes,
	}.Build()
	File_teleport_crownjewel_v1_crownjewel_proto = out.File
	file_teleport_crownjewel_v1_crownjewel_proto_rawDesc = nil
	file_teleport_crownjewel_v1_crownjewel_proto_goTypes = nil
	file_teleport_crownjewel_v1_crownjewel_proto_depIdxs = nil
}
