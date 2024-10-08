// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: internal/api/proto/V1/binary_data/binary_data.proto

package binary_data

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

type BinaryDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Notes   string `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
	EntryId string `protobuf:"bytes,3,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *BinaryDataRequest) Reset() {
	*x = BinaryDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryDataRequest) ProtoMessage() {}

func (x *BinaryDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryDataRequest.ProtoReflect.Descriptor instead.
func (*BinaryDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescGZIP(), []int{0}
}

func (x *BinaryDataRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *BinaryDataRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *BinaryDataRequest) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

type BinaryDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BinaryDataResponse) Reset() {
	*x = BinaryDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryDataResponse) ProtoMessage() {}

func (x *BinaryDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryDataResponse.ProtoReflect.Descriptor instead.
func (*BinaryDataResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescGZIP(), []int{1}
}

func (x *BinaryDataResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetBinaryDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *GetBinaryDataRequest) Reset() {
	*x = GetBinaryDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBinaryDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBinaryDataRequest) ProtoMessage() {}

func (x *GetBinaryDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBinaryDataRequest.ProtoReflect.Descriptor instead.
func (*GetBinaryDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescGZIP(), []int{2}
}

func (x *GetBinaryDataRequest) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

type GetBinaryDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Notes   string `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
	EntryId string `protobuf:"bytes,3,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *GetBinaryDataResponse) Reset() {
	*x = GetBinaryDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBinaryDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBinaryDataResponse) ProtoMessage() {}

func (x *GetBinaryDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBinaryDataResponse.ProtoReflect.Descriptor instead.
func (*GetBinaryDataResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescGZIP(), []int{3}
}

func (x *GetBinaryDataResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GetBinaryDataResponse) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *GetBinaryDataResponse) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

type EditBinaryDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Notes   string `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
	EntryId string `protobuf:"bytes,3,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *EditBinaryDataRequest) Reset() {
	*x = EditBinaryDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditBinaryDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditBinaryDataRequest) ProtoMessage() {}

func (x *EditBinaryDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditBinaryDataRequest.ProtoReflect.Descriptor instead.
func (*EditBinaryDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescGZIP(), []int{4}
}

func (x *EditBinaryDataRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *EditBinaryDataRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *EditBinaryDataRequest) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

type DeleteBinaryDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *DeleteBinaryDataRequest) Reset() {
	*x = DeleteBinaryDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBinaryDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBinaryDataRequest) ProtoMessage() {}

func (x *DeleteBinaryDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBinaryDataRequest.ProtoReflect.Descriptor instead.
func (*DeleteBinaryDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBinaryDataRequest) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

var File_internal_api_proto_V1_binary_data_binary_data_proto protoreflect.FileDescriptor

var file_internal_api_proto_V1_binary_data_binary_data_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x56, 0x31, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x76, 0x31, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x58, 0x0a, 0x11, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22,
	0x2e, 0x0a, 0x12, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x31, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x49, 0x64, 0x22, 0x5c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64,
	0x22, 0x5c, 0x0a, 0x15, 0x45, 0x64, 0x69, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22, 0x34,
	0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x64, 0x32, 0x88, 0x03, 0x0a, 0x11, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0e, 0x53, 0x61,
	0x76, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x2e, 0x76,
	0x31, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x45, 0x64, 0x69, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x25, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x27, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31,
	0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x24, 0x5a, 0x22, 0x67, 0x6f, 0x70, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x56, 0x31, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescOnce sync.Once
	file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescData = file_internal_api_proto_V1_binary_data_binary_data_proto_rawDesc
)

func file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescGZIP() []byte {
	file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescOnce.Do(func() {
		file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescData)
	})
	return file_internal_api_proto_V1_binary_data_binary_data_proto_rawDescData
}

var file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_api_proto_V1_binary_data_binary_data_proto_goTypes = []any{
	(*BinaryDataRequest)(nil),       // 0: v1.binary_data.BinaryDataRequest
	(*BinaryDataResponse)(nil),      // 1: v1.binary_data.BinaryDataResponse
	(*GetBinaryDataRequest)(nil),    // 2: v1.binary_data.GetBinaryDataRequest
	(*GetBinaryDataResponse)(nil),   // 3: v1.binary_data.GetBinaryDataResponse
	(*EditBinaryDataRequest)(nil),   // 4: v1.binary_data.EditBinaryDataRequest
	(*DeleteBinaryDataRequest)(nil), // 5: v1.binary_data.DeleteBinaryDataRequest
}
var file_internal_api_proto_V1_binary_data_binary_data_proto_depIdxs = []int32{
	0, // 0: v1.binary_data.BinaryDataService.SaveBinaryData:input_type -> v1.binary_data.BinaryDataRequest
	2, // 1: v1.binary_data.BinaryDataService.GetBinaryData:input_type -> v1.binary_data.GetBinaryDataRequest
	4, // 2: v1.binary_data.BinaryDataService.EditBinaryData:input_type -> v1.binary_data.EditBinaryDataRequest
	5, // 3: v1.binary_data.BinaryDataService.DeleteBinaryData:input_type -> v1.binary_data.DeleteBinaryDataRequest
	1, // 4: v1.binary_data.BinaryDataService.SaveBinaryData:output_type -> v1.binary_data.BinaryDataResponse
	3, // 5: v1.binary_data.BinaryDataService.GetBinaryData:output_type -> v1.binary_data.GetBinaryDataResponse
	1, // 6: v1.binary_data.BinaryDataService.EditBinaryData:output_type -> v1.binary_data.BinaryDataResponse
	1, // 7: v1.binary_data.BinaryDataService.DeleteBinaryData:output_type -> v1.binary_data.BinaryDataResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_api_proto_V1_binary_data_binary_data_proto_init() }
func file_internal_api_proto_V1_binary_data_binary_data_proto_init() {
	if File_internal_api_proto_V1_binary_data_binary_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BinaryDataRequest); i {
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
		file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BinaryDataResponse); i {
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
		file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetBinaryDataRequest); i {
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
		file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetBinaryDataResponse); i {
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
		file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*EditBinaryDataRequest); i {
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
		file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteBinaryDataRequest); i {
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
			RawDescriptor: file_internal_api_proto_V1_binary_data_binary_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_api_proto_V1_binary_data_binary_data_proto_goTypes,
		DependencyIndexes: file_internal_api_proto_V1_binary_data_binary_data_proto_depIdxs,
		MessageInfos:      file_internal_api_proto_V1_binary_data_binary_data_proto_msgTypes,
	}.Build()
	File_internal_api_proto_V1_binary_data_binary_data_proto = out.File
	file_internal_api_proto_V1_binary_data_binary_data_proto_rawDesc = nil
	file_internal_api_proto_V1_binary_data_binary_data_proto_goTypes = nil
	file_internal_api_proto_V1_binary_data_binary_data_proto_depIdxs = nil
}
