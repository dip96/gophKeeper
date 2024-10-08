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

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
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

func (x *TextDataRequest) GetId() string {
	if x != nil {
		return x.Id
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

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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

func (x *GetTextDataRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTextDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
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

func (x *GetTextDataResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EditTextDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
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

func (x *EditTextDataRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTextDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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

func (x *DeleteTextDataRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllTextDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllTextDataRequest) Reset() {
	*x = GetAllTextDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTextDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTextDataRequest) ProtoMessage() {}

func (x *GetAllTextDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTextDataRequest.ProtoReflect.Descriptor instead.
func (*GetAllTextDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_text_data_text_data_proto_rawDescGZIP(), []int{6}
}

type GetAllTextDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GetTextDataResponse `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetAllTextDataResponse) Reset() {
	*x = GetAllTextDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTextDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTextDataResponse) ProtoMessage() {}

func (x *GetAllTextDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTextDataResponse.ProtoReflect.Descriptor instead.
func (*GetAllTextDataResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_text_data_text_data_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllTextDataResponse) GetItems() []*GetTextDataResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_internal_api_proto_V1_text_data_text_data_proto protoreflect.FileDescriptor

var file_internal_api_proto_V1_text_data_text_data_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x56, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x35, 0x0a, 0x0f, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x13, 0x45, 0x64, 0x69, 0x74, 0x54, 0x65, 0x78,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x51, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x78, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xbb, 0x03, 0x0a, 0x0f, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x53, 0x61, 0x76,
	0x65, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c,
	0x45, 0x64, 0x69, 0x74, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x64, 0x69, 0x74,
	0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54,
	0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x55, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x23, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65,
	0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x6f, 0x70, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x56, 0x31, 0x2f, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_internal_api_proto_V1_text_data_text_data_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internal_api_proto_V1_text_data_text_data_proto_goTypes = []any{
	(*TextDataRequest)(nil),        // 0: v1.text_data.TextDataRequest
	(*TextDataResponse)(nil),       // 1: v1.text_data.TextDataResponse
	(*GetTextDataRequest)(nil),     // 2: v1.text_data.GetTextDataRequest
	(*GetTextDataResponse)(nil),    // 3: v1.text_data.GetTextDataResponse
	(*EditTextDataRequest)(nil),    // 4: v1.text_data.EditTextDataRequest
	(*DeleteTextDataRequest)(nil),  // 5: v1.text_data.DeleteTextDataRequest
	(*GetAllTextDataRequest)(nil),  // 6: v1.text_data.GetAllTextDataRequest
	(*GetAllTextDataResponse)(nil), // 7: v1.text_data.GetAllTextDataResponse
}
var file_internal_api_proto_V1_text_data_text_data_proto_depIdxs = []int32{
	3, // 0: v1.text_data.GetAllTextDataResponse.items:type_name -> v1.text_data.GetTextDataResponse
	0, // 1: v1.text_data.TextDataService.SaveTextData:input_type -> v1.text_data.TextDataRequest
	2, // 2: v1.text_data.TextDataService.GetTextData:input_type -> v1.text_data.GetTextDataRequest
	4, // 3: v1.text_data.TextDataService.EditTextData:input_type -> v1.text_data.EditTextDataRequest
	5, // 4: v1.text_data.TextDataService.DeleteTextData:input_type -> v1.text_data.DeleteTextDataRequest
	6, // 5: v1.text_data.TextDataService.GetAllTextData:input_type -> v1.text_data.GetAllTextDataRequest
	1, // 6: v1.text_data.TextDataService.SaveTextData:output_type -> v1.text_data.TextDataResponse
	3, // 7: v1.text_data.TextDataService.GetTextData:output_type -> v1.text_data.GetTextDataResponse
	1, // 8: v1.text_data.TextDataService.EditTextData:output_type -> v1.text_data.TextDataResponse
	1, // 9: v1.text_data.TextDataService.DeleteTextData:output_type -> v1.text_data.TextDataResponse
	7, // 10: v1.text_data.TextDataService.GetAllTextData:output_type -> v1.text_data.GetAllTextDataResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
		file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTextDataRequest); i {
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
		file_internal_api_proto_V1_text_data_text_data_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTextDataResponse); i {
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
			NumMessages:   8,
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
