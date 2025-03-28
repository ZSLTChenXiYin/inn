// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: captcha.proto

package captcha

import (
	context "context"
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

type EmailCaptchaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email []string `protobuf:"bytes,1,rep,name=email,proto3" json:"email,omitempty"`
}

func (x *EmailCaptchaRequest) Reset() {
	*x = EmailCaptchaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailCaptchaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailCaptchaRequest) ProtoMessage() {}

func (x *EmailCaptchaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailCaptchaRequest.ProtoReflect.Descriptor instead.
func (*EmailCaptchaRequest) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{0}
}

func (x *EmailCaptchaRequest) GetEmail() []string {
	if x != nil {
		return x.Email
	}
	return nil
}

type EmailCaptchaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success []bool `protobuf:"varint,1,rep,packed,name=success,proto3" json:"success,omitempty"`
}

func (x *EmailCaptchaResponse) Reset() {
	*x = EmailCaptchaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailCaptchaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailCaptchaResponse) ProtoMessage() {}

func (x *EmailCaptchaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailCaptchaResponse.ProtoReflect.Descriptor instead.
func (*EmailCaptchaResponse) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{1}
}

func (x *EmailCaptchaResponse) GetSuccess() []bool {
	if x != nil {
		return x.Success
	}
	return nil
}

type EmailValidateList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Captcha string `protobuf:"bytes,2,opt,name=captcha,proto3" json:"captcha,omitempty"`
}

func (x *EmailValidateList) Reset() {
	*x = EmailValidateList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailValidateList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailValidateList) ProtoMessage() {}

func (x *EmailValidateList) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailValidateList.ProtoReflect.Descriptor instead.
func (*EmailValidateList) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{2}
}

func (x *EmailValidateList) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailValidateList) GetCaptcha() string {
	if x != nil {
		return x.Captcha
	}
	return ""
}

type EmailValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lists []*EmailValidateList `protobuf:"bytes,1,rep,name=lists,proto3" json:"lists,omitempty"`
}

func (x *EmailValidateRequest) Reset() {
	*x = EmailValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailValidateRequest) ProtoMessage() {}

func (x *EmailValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailValidateRequest.ProtoReflect.Descriptor instead.
func (*EmailValidateRequest) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{3}
}

func (x *EmailValidateRequest) GetLists() []*EmailValidateList {
	if x != nil {
		return x.Lists
	}
	return nil
}

type EmailValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success []bool `protobuf:"varint,1,rep,packed,name=success,proto3" json:"success,omitempty"`
}

func (x *EmailValidateResponse) Reset() {
	*x = EmailValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailValidateResponse) ProtoMessage() {}

func (x *EmailValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailValidateResponse.ProtoReflect.Descriptor instead.
func (*EmailValidateResponse) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{4}
}

func (x *EmailValidateResponse) GetSuccess() []bool {
	if x != nil {
		return x.Success
	}
	return nil
}

type EmailRefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email []string `protobuf:"bytes,1,rep,name=email,proto3" json:"email,omitempty"`
}

func (x *EmailRefreshRequest) Reset() {
	*x = EmailRefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailRefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRefreshRequest) ProtoMessage() {}

func (x *EmailRefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRefreshRequest.ProtoReflect.Descriptor instead.
func (*EmailRefreshRequest) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{5}
}

func (x *EmailRefreshRequest) GetEmail() []string {
	if x != nil {
		return x.Email
	}
	return nil
}

type EmailRefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success []bool `protobuf:"varint,1,rep,packed,name=success,proto3" json:"success,omitempty"`
}

func (x *EmailRefreshResponse) Reset() {
	*x = EmailRefreshResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailRefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRefreshResponse) ProtoMessage() {}

func (x *EmailRefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRefreshResponse.ProtoReflect.Descriptor instead.
func (*EmailRefreshResponse) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{6}
}

func (x *EmailRefreshResponse) GetSuccess() []bool {
	if x != nil {
		return x.Success
	}
	return nil
}

var File_captcha_proto protoreflect.FileDescriptor

var file_captcha_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x22, 0x2b, 0x0a, 0x13, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x30, 0x0a, 0x14, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x43, 0x0a, 0x11, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x22, 0x48, 0x0a, 0x14,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x22, 0x31, 0x0a, 0x15, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2b, 0x0a, 0x13, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x30, 0x0a, 0x14, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xfa, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x1c, 0x2e, 0x63, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_captcha_proto_rawDescOnce sync.Once
	file_captcha_proto_rawDescData = file_captcha_proto_rawDesc
)

func file_captcha_proto_rawDescGZIP() []byte {
	file_captcha_proto_rawDescOnce.Do(func() {
		file_captcha_proto_rawDescData = protoimpl.X.CompressGZIP(file_captcha_proto_rawDescData)
	})
	return file_captcha_proto_rawDescData
}

var file_captcha_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_captcha_proto_goTypes = []interface{}{
	(*EmailCaptchaRequest)(nil),   // 0: captcha.EmailCaptchaRequest
	(*EmailCaptchaResponse)(nil),  // 1: captcha.EmailCaptchaResponse
	(*EmailValidateList)(nil),     // 2: captcha.EmailValidateList
	(*EmailValidateRequest)(nil),  // 3: captcha.EmailValidateRequest
	(*EmailValidateResponse)(nil), // 4: captcha.EmailValidateResponse
	(*EmailRefreshRequest)(nil),   // 5: captcha.EmailRefreshRequest
	(*EmailRefreshResponse)(nil),  // 6: captcha.EmailRefreshResponse
}
var file_captcha_proto_depIdxs = []int32{
	2, // 0: captcha.EmailValidateRequest.lists:type_name -> captcha.EmailValidateList
	0, // 1: captcha.CaptchaService.EmailCaptcha:input_type -> captcha.EmailCaptchaRequest
	3, // 2: captcha.CaptchaService.EmailValidate:input_type -> captcha.EmailValidateRequest
	5, // 3: captcha.CaptchaService.EmailRefresh:input_type -> captcha.EmailRefreshRequest
	1, // 4: captcha.CaptchaService.EmailCaptcha:output_type -> captcha.EmailCaptchaResponse
	4, // 5: captcha.CaptchaService.EmailValidate:output_type -> captcha.EmailValidateResponse
	6, // 6: captcha.CaptchaService.EmailRefresh:output_type -> captcha.EmailRefreshResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_captcha_proto_init() }
func file_captcha_proto_init() {
	if File_captcha_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_captcha_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailCaptchaRequest); i {
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
		file_captcha_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailCaptchaResponse); i {
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
		file_captcha_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailValidateList); i {
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
		file_captcha_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailValidateRequest); i {
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
		file_captcha_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailValidateResponse); i {
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
		file_captcha_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailRefreshRequest); i {
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
		file_captcha_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailRefreshResponse); i {
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
			RawDescriptor: file_captcha_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_captcha_proto_goTypes,
		DependencyIndexes: file_captcha_proto_depIdxs,
		MessageInfos:      file_captcha_proto_msgTypes,
	}.Build()
	File_captcha_proto = out.File
	file_captcha_proto_rawDesc = nil
	file_captcha_proto_goTypes = nil
	file_captcha_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type CaptchaService interface {
	EmailCaptcha(ctx context.Context, req *EmailCaptchaRequest) (res *EmailCaptchaResponse, err error)
	EmailValidate(ctx context.Context, req *EmailValidateRequest) (res *EmailValidateResponse, err error)
	EmailRefresh(ctx context.Context, req *EmailRefreshRequest) (res *EmailRefreshResponse, err error)
}
