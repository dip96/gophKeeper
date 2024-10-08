// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: internal/api/proto/V1/users/user_service.proto

package users

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

type Platform int32

const (
	Platform_UNKNOWN Platform = 0
	Platform_WINDOWS Platform = 1
	Platform_LINUX   Platform = 2
	Platform_MACOS   Platform = 3
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0: "UNKNOWN",
		1: "WINDOWS",
		2: "LINUX",
		3: "MACOS",
	}
	Platform_value = map[string]int32{
		"UNKNOWN": 0,
		"WINDOWS": 1,
		"LINUX":   2,
		"MACOS":   3,
	}
)

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}

func (x Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_api_proto_V1_users_user_service_proto_enumTypes[0].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_internal_api_proto_V1_users_user_service_proto_enumTypes[0]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_users_user_service_proto_rawDescGZIP(), []int{0}
}

type UserRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Otp      string   `protobuf:"bytes,3,opt,name=otp,proto3" json:"otp,omitempty"`
	Platform Platform `protobuf:"varint,5,opt,name=platform,proto3,enum=v1.users.Platform" json:"platform,omitempty"` // определение платформы
}

func (x *UserRegistrationRequest) Reset() {
	*x = UserRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_users_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegistrationRequest) ProtoMessage() {}

func (x *UserRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_users_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegistrationRequest.ProtoReflect.Descriptor instead.
func (*UserRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_users_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegistrationRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserRegistrationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserRegistrationRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *UserRegistrationRequest) GetPlatform() Platform {
	if x != nil {
		return x.Platform
	}
	return Platform_UNKNOWN
}

type UserRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserRegistrationResponse) Reset() {
	*x = UserRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_users_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegistrationResponse) ProtoMessage() {}

func (x *UserRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_users_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegistrationResponse.ProtoReflect.Descriptor instead.
func (*UserRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_users_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserRegistrationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_users_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_users_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_users_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_V1_users_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_V1_users_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_V1_users_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_internal_api_proto_V1_users_user_service_proto protoreflect.FileDescriptor

var file_internal_api_proto_V1_users_user_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x56, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x17, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x76,
	0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x34, 0x0a, 0x18, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x40, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x3f, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2a, 0x3a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4e,
	0x55, 0x58, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x43, 0x4f, 0x53, 0x10, 0x03, 0x32,
	0x9e, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x55, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x16, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x6f, 0x70, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x56, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_api_proto_V1_users_user_service_proto_rawDescOnce sync.Once
	file_internal_api_proto_V1_users_user_service_proto_rawDescData = file_internal_api_proto_V1_users_user_service_proto_rawDesc
)

func file_internal_api_proto_V1_users_user_service_proto_rawDescGZIP() []byte {
	file_internal_api_proto_V1_users_user_service_proto_rawDescOnce.Do(func() {
		file_internal_api_proto_V1_users_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_api_proto_V1_users_user_service_proto_rawDescData)
	})
	return file_internal_api_proto_V1_users_user_service_proto_rawDescData
}

var file_internal_api_proto_V1_users_user_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_api_proto_V1_users_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_api_proto_V1_users_user_service_proto_goTypes = []any{
	(Platform)(0),                    // 0: v1.users.Platform
	(*UserRegistrationRequest)(nil),  // 1: v1.users.UserRegistrationRequest
	(*UserRegistrationResponse)(nil), // 2: v1.users.UserRegistrationResponse
	(*LoginRequest)(nil),             // 3: v1.users.LoginRequest
	(*LoginResponse)(nil),            // 4: v1.users.LoginResponse
}
var file_internal_api_proto_V1_users_user_service_proto_depIdxs = []int32{
	0, // 0: v1.users.UserRegistrationRequest.platform:type_name -> v1.users.Platform
	1, // 1: v1.users.UserService.Registration:input_type -> v1.users.UserRegistrationRequest
	3, // 2: v1.users.UserService.Login:input_type -> v1.users.LoginRequest
	2, // 3: v1.users.UserService.Registration:output_type -> v1.users.UserRegistrationResponse
	4, // 4: v1.users.UserService.Login:output_type -> v1.users.LoginResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_api_proto_V1_users_user_service_proto_init() }
func file_internal_api_proto_V1_users_user_service_proto_init() {
	if File_internal_api_proto_V1_users_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_api_proto_V1_users_user_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserRegistrationRequest); i {
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
		file_internal_api_proto_V1_users_user_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserRegistrationResponse); i {
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
		file_internal_api_proto_V1_users_user_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LoginRequest); i {
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
		file_internal_api_proto_V1_users_user_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LoginResponse); i {
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
			RawDescriptor: file_internal_api_proto_V1_users_user_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_api_proto_V1_users_user_service_proto_goTypes,
		DependencyIndexes: file_internal_api_proto_V1_users_user_service_proto_depIdxs,
		EnumInfos:         file_internal_api_proto_V1_users_user_service_proto_enumTypes,
		MessageInfos:      file_internal_api_proto_V1_users_user_service_proto_msgTypes,
	}.Build()
	File_internal_api_proto_V1_users_user_service_proto = out.File
	file_internal_api_proto_V1_users_user_service_proto_rawDesc = nil
	file_internal_api_proto_V1_users_user_service_proto_goTypes = nil
	file_internal_api_proto_V1_users_user_service_proto_depIdxs = nil
}
