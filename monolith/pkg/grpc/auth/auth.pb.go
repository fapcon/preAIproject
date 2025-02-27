// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: auth.proto

package auth

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

type VerifyEmailIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash  string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *VerifyEmailIn) Reset() {
	*x = VerifyEmailIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyEmailIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailIn) ProtoMessage() {}

func (x *VerifyEmailIn) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailIn.ProtoReflect.Descriptor instead.
func (*VerifyEmailIn) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyEmailIn) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *VerifyEmailIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type VerifyEmailOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorCode int32 `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *VerifyEmailOut) Reset() {
	*x = VerifyEmailOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyEmailOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailOut) ProtoMessage() {}

func (x *VerifyEmailOut) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailOut.ProtoReflect.Descriptor instead.
func (*VerifyEmailOut) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyEmailOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *VerifyEmailOut) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

type SendPhoneCodeIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *SendPhoneCodeIn) Reset() {
	*x = SendPhoneCodeIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPhoneCodeIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPhoneCodeIn) ProtoMessage() {}

func (x *SendPhoneCodeIn) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPhoneCodeIn.ProtoReflect.Descriptor instead.
func (*SendPhoneCodeIn) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SendPhoneCodeIn) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type SendPhoneCodeOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone     string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code      int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	ErrorCode int32  `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *SendPhoneCodeOut) Reset() {
	*x = SendPhoneCodeOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPhoneCodeOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPhoneCodeOut) ProtoMessage() {}

func (x *SendPhoneCodeOut) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPhoneCodeOut.ProtoReflect.Descriptor instead.
func (*SendPhoneCodeOut) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SendPhoneCodeOut) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SendPhoneCodeOut) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SendPhoneCodeOut) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

type AuthorizeIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthorizeIn) Reset() {
	*x = AuthorizeIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeIn) ProtoMessage() {}

func (x *AuthorizeIn) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeIn.ProtoReflect.Descriptor instead.
func (*AuthorizeIn) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorizeIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthorizeIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthorizeOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ErrorCode    int32  `protobuf:"varint,4,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *AuthorizeOut) Reset() {
	*x = AuthorizeOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeOut) ProtoMessage() {}

func (x *AuthorizeOut) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeOut.ProtoReflect.Descriptor instead.
func (*AuthorizeOut) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

func (x *AuthorizeOut) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AuthorizeOut) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthorizeOut) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AuthorizeOut) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

type RegisterIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email          string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone          string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password       string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	IdempotencyKey string `protobuf:"bytes,4,opt,name=idempotency_key,json=idempotencyKey,proto3" json:"idempotency_key,omitempty"`
	Field          int32  `protobuf:"varint,5,opt,name=Field,proto3" json:"Field,omitempty"`
}

func (x *RegisterIn) Reset() {
	*x = RegisterIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterIn) ProtoMessage() {}

func (x *RegisterIn) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterIn.ProtoReflect.Descriptor instead.
func (*RegisterIn) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterIn) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RegisterIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterIn) GetIdempotencyKey() string {
	if x != nil {
		return x.IdempotencyKey
	}
	return ""
}

func (x *RegisterIn) GetField() int32 {
	if x != nil {
		return x.Field
	}
	return 0
}

type RegisterOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ErrorCode int32 `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *RegisterOut) Reset() {
	*x = RegisterOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOut) ProtoMessage() {}

func (x *RegisterOut) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOut.ProtoReflect.Descriptor instead.
func (*RegisterOut) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterOut) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RegisterOut) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

type AuthorizeEmailIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email          string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password       string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	RetypePassword string `protobuf:"bytes,3,opt,name=retype_password,json=retypePassword,proto3" json:"retype_password,omitempty"`
}

func (x *AuthorizeEmailIn) Reset() {
	*x = AuthorizeEmailIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeEmailIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeEmailIn) ProtoMessage() {}

func (x *AuthorizeEmailIn) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeEmailIn.ProtoReflect.Descriptor instead.
func (*AuthorizeEmailIn) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{8}
}

func (x *AuthorizeEmailIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthorizeEmailIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthorizeEmailIn) GetRetypePassword() string {
	if x != nil {
		return x.RetypePassword
	}
	return ""
}

type AuthorizeRefreshIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AuthorizeRefreshIn) Reset() {
	*x = AuthorizeRefreshIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeRefreshIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeRefreshIn) ProtoMessage() {}

func (x *AuthorizeRefreshIn) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeRefreshIn.ProtoReflect.Descriptor instead.
func (*AuthorizeRefreshIn) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{9}
}

func (x *AuthorizeRefreshIn) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AuthorizePhoneIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code  int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *AuthorizePhoneIn) Reset() {
	*x = AuthorizePhoneIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizePhoneIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizePhoneIn) ProtoMessage() {}

func (x *AuthorizePhoneIn) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizePhoneIn.ProtoReflect.Descriptor instead.
func (*AuthorizePhoneIn) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{10}
}

func (x *AuthorizePhoneIn) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AuthorizePhoneIn) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type SocialCallbackIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *SocialCallbackIn) Reset() {
	*x = SocialCallbackIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialCallbackIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialCallbackIn) ProtoMessage() {}

func (x *SocialCallbackIn) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialCallbackIn.ProtoReflect.Descriptor instead.
func (*SocialCallbackIn) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{11}
}

func (x *SocialCallbackIn) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SocialCallbackIn) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type SocialGetRedirectUrlIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *SocialGetRedirectUrlIn) Reset() {
	*x = SocialGetRedirectUrlIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialGetRedirectUrlIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialGetRedirectUrlIn) ProtoMessage() {}

func (x *SocialGetRedirectUrlIn) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialGetRedirectUrlIn.ProtoReflect.Descriptor instead.
func (*SocialGetRedirectUrlIn) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{12}
}

func (x *SocialGetRedirectUrlIn) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type SocialGetRedirectUrlOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url       string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	ErrorCode int32  `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *SocialGetRedirectUrlOut) Reset() {
	*x = SocialGetRedirectUrlOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialGetRedirectUrlOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialGetRedirectUrlOut) ProtoMessage() {}

func (x *SocialGetRedirectUrlOut) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialGetRedirectUrlOut.ProtoReflect.Descriptor instead.
func (*SocialGetRedirectUrlOut) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{13}
}

func (x *SocialGetRedirectUrlOut) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SocialGetRedirectUrlOut) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x39, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x49, 0x0a,
	0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x75, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x22, 0x5b, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x3f,
	0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x8e, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x4f, 0x75, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x93, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6d, 0x70,
	0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x44, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x4f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x6d, 0x0a, 0x10,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x10, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x49, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x42, 0x0a, 0x10, 0x53, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x16,
	0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x55, 0x72, 0x6c, 0x49, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x22, 0x4a, 0x0a, 0x17, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x4f, 0x75, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x9d,
	0x04, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x52,
	0x50, 0x43, 0x12, 0x31, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x1a, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x1a,
	0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x49, 0x6e, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x49, 0x6e, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x53, 0x65, 0x6e,
	0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x1a,
	0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x53, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49,
	0x6e, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x14, 0x53, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x12,
	0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x49, 0x6e, 0x1a, 0x1d, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x42, 0x48,
	0x5a, 0x46, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x67, 0x69, 0x74, 0x2e, 0x6b, 0x61, 0x74,
	0x61, 0x2e, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x79, 0x2f, 0x65, 0x61, 0x7a, 0x7a, 0x79, 0x65,
	0x61, 0x72, 0x6e, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6d, 0x6f, 0x6e,
	0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_auth_proto_goTypes = []interface{}{
	(*VerifyEmailIn)(nil),           // 0: auth.VerifyEmailIn
	(*VerifyEmailOut)(nil),          // 1: auth.VerifyEmailOut
	(*SendPhoneCodeIn)(nil),         // 2: auth.SendPhoneCodeIn
	(*SendPhoneCodeOut)(nil),        // 3: auth.SendPhoneCodeOut
	(*AuthorizeIn)(nil),             // 4: auth.AuthorizeIn
	(*AuthorizeOut)(nil),            // 5: auth.AuthorizeOut
	(*RegisterIn)(nil),              // 6: auth.RegisterIn
	(*RegisterOut)(nil),             // 7: auth.RegisterOut
	(*AuthorizeEmailIn)(nil),        // 8: auth.AuthorizeEmailIn
	(*AuthorizeRefreshIn)(nil),      // 9: auth.AuthorizeRefreshIn
	(*AuthorizePhoneIn)(nil),        // 10: auth.AuthorizePhoneIn
	(*SocialCallbackIn)(nil),        // 11: auth.SocialCallbackIn
	(*SocialGetRedirectUrlIn)(nil),  // 12: auth.SocialGetRedirectUrlIn
	(*SocialGetRedirectUrlOut)(nil), // 13: auth.SocialGetRedirectUrlOut
}
var file_auth_proto_depIdxs = []int32{
	6,  // 0: auth.AuthServiceGRPC.Register:input_type -> auth.RegisterIn
	8,  // 1: auth.AuthServiceGRPC.AuthorizeEmail:input_type -> auth.AuthorizeEmailIn
	9,  // 2: auth.AuthServiceGRPC.AuthorizeRefresh:input_type -> auth.AuthorizeRefreshIn
	10, // 3: auth.AuthServiceGRPC.AuthorizePhone:input_type -> auth.AuthorizePhoneIn
	2,  // 4: auth.AuthServiceGRPC.SendPhoneCode:input_type -> auth.SendPhoneCodeIn
	0,  // 5: auth.AuthServiceGRPC.VerifyEmail:input_type -> auth.VerifyEmailIn
	11, // 6: auth.AuthServiceGRPC.SocialCallback:input_type -> auth.SocialCallbackIn
	12, // 7: auth.AuthServiceGRPC.SocialGetRedirectURL:input_type -> auth.SocialGetRedirectUrlIn
	7,  // 8: auth.AuthServiceGRPC.Register:output_type -> auth.RegisterOut
	5,  // 9: auth.AuthServiceGRPC.AuthorizeEmail:output_type -> auth.AuthorizeOut
	5,  // 10: auth.AuthServiceGRPC.AuthorizeRefresh:output_type -> auth.AuthorizeOut
	5,  // 11: auth.AuthServiceGRPC.AuthorizePhone:output_type -> auth.AuthorizeOut
	3,  // 12: auth.AuthServiceGRPC.SendPhoneCode:output_type -> auth.SendPhoneCodeOut
	1,  // 13: auth.AuthServiceGRPC.VerifyEmail:output_type -> auth.VerifyEmailOut
	5,  // 14: auth.AuthServiceGRPC.SocialCallback:output_type -> auth.AuthorizeOut
	13, // 15: auth.AuthServiceGRPC.SocialGetRedirectURL:output_type -> auth.SocialGetRedirectUrlOut
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyEmailIn); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyEmailOut); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendPhoneCodeIn); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendPhoneCodeOut); i {
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
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeIn); i {
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
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeOut); i {
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
		file_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterIn); i {
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
		file_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterOut); i {
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
		file_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeEmailIn); i {
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
		file_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeRefreshIn); i {
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
		file_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizePhoneIn); i {
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
		file_auth_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialCallbackIn); i {
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
		file_auth_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialGetRedirectUrlIn); i {
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
		file_auth_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialGetRedirectUrlOut); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
