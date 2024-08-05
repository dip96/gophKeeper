// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: internal/api/proto/V1/text_data/text_data.proto

package text_data

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

type TextDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text    string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	EntryId string `protobuf:"bytes,2,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *TextDataRequest) Reset() {
	*x = TextDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextDataRequest) ProtoMessage() {}

func (x *TextDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextDataRequest.ProtoReflect.Descriptor instead.
func (*TextDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_text_data_text_data_proto_rawDescGZIP(), []int{0}
}

func (x *TextDataRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TextDataRequest) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

type TextDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TextDataResponse) Reset() {
	*x = TextDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextDataResponse) ProtoMessage() {}

func (x *TextDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextDataResponse.ProtoReflect.Descriptor instead.
func (*TextDataResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_text_data_text_data_proto_rawDescGZIP(), []int{1}
}

func (x *TextDataResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetTextDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *GetTextDataRequest) Reset() {
	*x = GetTextDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextDataRequest) ProtoMessage() {}

func (x *GetTextDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextDataRequest.ProtoReflect.Descriptor instead.
func (*GetTextDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_text_data_text_data_proto_rawDescGZIP(), []int{2}
}

func (x *GetTextDataRequest) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

type GetTextDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text    string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	EntryId string `protobuf:"bytes,2,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *GetTextDataResponse) Reset() {
	*x = GetTextDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextDataResponse) ProtoMessage() {}

func (x *GetTextDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextDataResponse.ProtoReflect.Descriptor instead.
func (*GetTextDataResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_text_data_text_data_proto_rawDescGZIP(), []int{3}
}

func (x *GetTextDataResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetTextDataResponse) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

type EditTextDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text    string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	EntryId string `protobuf:"bytes,2,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *EditTextDataRequest) Reset() {
	*x = EditTextDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditTextDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditTextDataRequest) ProtoMessage() {}

func (x *EditTextDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditTextDataRequest.ProtoReflect.Descriptor instead.
func (*EditTextDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_text_data_text_data_proto_rawDescGZIP(), []int{4}
}

func (x *EditTextDataRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *EditTextDataRequest) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

type DeleteTextDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *DeleteTextDataRequest) Reset() {
	*x = DeleteTextDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTextDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTextDataRequest) ProtoMessage() {}

func (x *DeleteTextDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTextDataRequest.ProtoReflect.Descriptor instead.
func (*DeleteTextDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_text_data_text_data_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTextDataRequest) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

var File_internal_api_proto_V1_text_data_text_data_proto protoreflect.FileDescriptor

var file_internal_api_proto_V1_text_data_text_data_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x56, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x40, 0x0a, 0x0f, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x64, 0x22, 0x2c, 0x0a, 0x10, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x2f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64,
	0x22, 0x44, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x13, 0x45, 0x64, 0x69, 0x74, 0x54, 0x65,
	0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64,
	0x32, 0xde, 0x02, 0x0a, 0x0f, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x54,
	0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x2e,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x6f, 0x70, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x56, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_api_proto_V1_text_data_text_data_proto_rawDescOnce sync.Once
	file_internal_api_proto_V1_text_data_text_data_proto_rawDescData = file_internal_api_proto_V1_text_data_text_data_proto_rawDesc
)

func file_internal_api_proto_V1_text_data_text_data_proto_rawDescGZIP() []byte {
	file_internal_api_proto_V1_text_data_text_data_proto_rawDescOnce.Do(func() {
		file_internal_api_proto_V1_text_data_text_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_api_proto_V1_text_data_text_data_proto_rawDescData)
	})
	return file_internal_api_proto_V1_text_data_text_data_proto_rawDescData
}

var file_internal_api_proto_V1_text_data_text_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_api_proto_V1_text_data_text_data_proto_goTypes = []any{
	(*TextDataRequest)(nil),       // 0: v1.text_data.TextDataRequest
	(*TextDataResponse)(nil),      // 1: v1.text_data.TextDataResponse
	(*GetTextDataRequest)(nil),    // 2: v1.text_data.GetTextDataRequest
	(*GetTextDataResponse)(nil),   // 3: v1.text_data.GetTextDataResponse
	(*EditTextDataRequest)(nil),   // 4: v1.text_data.EditTextDataRequest
	(*DeleteTextDataRequest)(nil), // 5: v1.text_data.DeleteTextDataRequest
}
var file_internal_api_proto_V1_text_data_text_data_proto_depIdxs = []int32{
	0, // 0: v1.text_data.TextDataService.SaveTextData:input_type -> v1.text_data.TextDataRequest
	2, // 1: v1.text_data.TextDataService.GetTextData:input_type -> v1.text_data.GetTextDataRequest
	4, // 2: v1.text_data.TextDataService.EditTextData:input_type -> v1.text_data.EditTextDataRequest
	5, // 3: v1.text_data.TextDataService.DeleteTextData:input_type -> v1.text_data.DeleteTextDataRequest
	1, // 4: v1.text_data.TextDataService.SaveTextData:output_type -> v1.text_data.TextDataResponse
	3, // 5: v1.text_data.TextDataService.GetTextData:output_type -> v1.text_data.GetTextDataResponse
	1, // 6: v1.text_data.TextDataService.EditTextData:output_type -> v1.text_data.TextDataResponse
	1, // 7: v1.text_data.TextDataService.DeleteTextData:output_type -> v1.text_data.TextDataResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_api_proto_V1_text_data_text_data_proto_init() }
func file_internal_api_proto_V1_text_data_text_data_proto_init() {
	if File_internal_api_proto_V1_text_data_text_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TextDataRequest); i {
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
		file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TextDataResponse); i {
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
		file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetTextDataRequest); i {
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
		file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetTextDataResponse); i {
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
		file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*EditTextDataRequest); i {
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
		file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTextDataRequest); i {
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
			RawDescriptor: file_internal_api_proto_V1_text_data_text_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_api_proto_V1_text_data_text_data_proto_goTypes,
		DependencyIndexes: file_internal_api_proto_V1_text_data_text_data_proto_depIdxs,
		MessageInfos:      file_internal_api_proto_V1_text_data_text_data_proto_msgTypes,
	}.Build()
	File_internal_api_proto_V1_text_data_text_data_proto = out.File
	file_internal_api_proto_V1_text_data_text_data_proto_rawDesc = nil
	file_internal_api_proto_V1_text_data_text_data_proto_goTypes = nil
	file_internal_api_proto_V1_text_data_text_data_proto_depIdxs = nil
}
