// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/service.proto

package go_tokenservice_grpc

import (
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

type RBroadcastMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rid uint64 `protobuf:"varint,1,opt,name=rid,proto3" json:"rid,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RBroadcastMessage) Reset() {
	*x = RBroadcastMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RBroadcastMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RBroadcastMessage) ProtoMessage() {}

func (x *RBroadcastMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RBroadcastMessage.ProtoReflect.Descriptor instead.
func (*RBroadcastMessage) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *RBroadcastMessage) GetRid() uint64 {
	if x != nil {
		return x.Rid
	}
	return 0
}

func (x *RBroadcastMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RBroadcastReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rid  uint64 `protobuf:"varint,1,opt,name=rid,proto3" json:"rid,omitempty"`
	Wts  uint64 `protobuf:"varint,2,opt,name=wts,proto3" json:"wts,omitempty"`
	Id   string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Low  uint64 `protobuf:"varint,6,opt,name=low,proto3" json:"low,omitempty"`
	High uint64 `protobuf:"varint,7,opt,name=high,proto3" json:"high,omitempty"`
	Mid  uint64 `protobuf:"varint,8,opt,name=mid,proto3" json:"mid,omitempty"`
	Val  uint64 `protobuf:"varint,9,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *RBroadcastReturn) Reset() {
	*x = RBroadcastReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RBroadcastReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RBroadcastReturn) ProtoMessage() {}

func (x *RBroadcastReturn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RBroadcastReturn.ProtoReflect.Descriptor instead.
func (*RBroadcastReturn) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *RBroadcastReturn) GetRid() uint64 {
	if x != nil {
		return x.Rid
	}
	return 0
}

func (x *RBroadcastReturn) GetWts() uint64 {
	if x != nil {
		return x.Wts
	}
	return 0
}

func (x *RBroadcastReturn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RBroadcastReturn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RBroadcastReturn) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *RBroadcastReturn) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *RBroadcastReturn) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *RBroadcastReturn) GetVal() uint64 {
	if x != nil {
		return x.Val
	}
	return 0
}

type WBroadcastReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack string `protobuf:"bytes,1,opt,name=ack,proto3" json:"ack,omitempty"`
	Rid uint64 `protobuf:"varint,2,opt,name=rid,proto3" json:"rid,omitempty"`
}

func (x *WBroadcastReturn) Reset() {
	*x = WBroadcastReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WBroadcastReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WBroadcastReturn) ProtoMessage() {}

func (x *WBroadcastReturn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WBroadcastReturn.ProtoReflect.Descriptor instead.
func (*WBroadcastReturn) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *WBroadcastReturn) GetAck() string {
	if x != nil {
		return x.Ack
	}
	return ""
}

func (x *WBroadcastReturn) GetRid() uint64 {
	if x != nil {
		return x.Rid
	}
	return 0
}

type WBroadcastMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rid  uint64 `protobuf:"varint,1,opt,name=rid,proto3" json:"rid,omitempty"`
	Wts  uint64 `protobuf:"varint,2,opt,name=wts,proto3" json:"wts,omitempty"`
	Id   string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Low  uint64 `protobuf:"varint,6,opt,name=low,proto3" json:"low,omitempty"`
	High uint64 `protobuf:"varint,7,opt,name=high,proto3" json:"high,omitempty"`
	Mid  uint64 `protobuf:"varint,8,opt,name=mid,proto3" json:"mid,omitempty"`
	Val  uint64 `protobuf:"varint,9,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WBroadcastMessage) Reset() {
	*x = WBroadcastMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WBroadcastMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WBroadcastMessage) ProtoMessage() {}

func (x *WBroadcastMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WBroadcastMessage.ProtoReflect.Descriptor instead.
func (*WBroadcastMessage) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *WBroadcastMessage) GetRid() uint64 {
	if x != nil {
		return x.Rid
	}
	return 0
}

func (x *WBroadcastMessage) GetWts() uint64 {
	if x != nil {
		return x.Wts
	}
	return 0
}

func (x *WBroadcastMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WBroadcastMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WBroadcastMessage) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *WBroadcastMessage) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *WBroadcastMessage) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *WBroadcastMessage) GetVal() uint64 {
	if x != nil {
		return x.Val
	}
	return 0
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Low  uint64 `protobuf:"varint,3,opt,name=low,proto3" json:"low,omitempty"`
	High uint64 `protobuf:"varint,4,opt,name=high,proto3" json:"high,omitempty"`
	Mid  uint64 `protobuf:"varint,5,opt,name=mid,proto3" json:"mid,omitempty"`
	Wts  uint64 `protobuf:"varint,6,opt,name=wts,proto3" json:"wts,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *WriteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WriteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WriteRequest) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *WriteRequest) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *WriteRequest) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *WriteRequest) GetWts() uint64 {
	if x != nil {
		return x.Wts
	}
	return 0
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *IdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack string `protobuf:"bytes,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *AckResponse) Reset() {
	*x = AckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckResponse) ProtoMessage() {}

func (x *AckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckResponse.ProtoReflect.Descriptor instead.
func (*AckResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{6}
}

func (x *AckResponse) GetAck() string {
	if x != nil {
		return x.Ack
	}
	return ""
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x11,
	0x52, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x72, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x10, 0x52, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x77, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6c,
	0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x36, 0x0a, 0x10, 0x57, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x63, 0x6b,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72,
	0x69, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x11, 0x57, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x77, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6c,
	0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x7c, 0x0a, 0x0c, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6c, 0x6f, 0x77,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x74, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x77, 0x74, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0b, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x32, 0xc9, 0x02, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x44, 0x72, 0x6f, 0x70, 0x12,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x6e,
	0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x12, 0x3c, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x65, 0x12,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x42, 0x37, 0x5a, 0x35, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x6f, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData = file_proto_service_proto_rawDesc
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_proto_rawDescData)
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_service_proto_goTypes = []interface{}{
	(*RBroadcastMessage)(nil), // 0: proto.RBroadcastMessage
	(*RBroadcastReturn)(nil),  // 1: proto.RBroadcastReturn
	(*WBroadcastReturn)(nil),  // 2: proto.WBroadcastReturn
	(*WBroadcastMessage)(nil), // 3: proto.WBroadcastMessage
	(*WriteRequest)(nil),      // 4: proto.WriteRequest
	(*IdRequest)(nil),         // 5: proto.IdRequest
	(*AckResponse)(nil),       // 6: proto.AckResponse
}
var file_proto_service_proto_depIdxs = []int32{
	5, // 0: proto.TokenService.Create:input_type -> proto.IdRequest
	5, // 1: proto.TokenService.Drop:input_type -> proto.IdRequest
	4, // 2: proto.TokenService.Write:input_type -> proto.WriteRequest
	5, // 3: proto.TokenService.Read:input_type -> proto.IdRequest
	3, // 4: proto.TokenService.WriteOne:input_type -> proto.WBroadcastMessage
	0, // 5: proto.TokenService.ReadOne:input_type -> proto.RBroadcastMessage
	6, // 6: proto.TokenService.Create:output_type -> proto.AckResponse
	6, // 7: proto.TokenService.Drop:output_type -> proto.AckResponse
	6, // 8: proto.TokenService.Write:output_type -> proto.AckResponse
	6, // 9: proto.TokenService.Read:output_type -> proto.AckResponse
	2, // 10: proto.TokenService.WriteOne:output_type -> proto.WBroadcastReturn
	1, // 11: proto.TokenService.ReadOne:output_type -> proto.RBroadcastReturn
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RBroadcastMessage); i {
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
		file_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RBroadcastReturn); i {
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
		file_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WBroadcastReturn); i {
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
		file_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WBroadcastMessage); i {
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
		file_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
		file_proto_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckResponse); i {
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
			RawDescriptor: file_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_rawDesc = nil
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}
