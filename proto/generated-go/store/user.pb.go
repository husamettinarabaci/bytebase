// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: store/user.proto

package store

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

// MFAConfig is the MFA configuration for a user.
type MFAConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The otp_secret is the secret key used to validate the OTP code.
	OtpSecret string `protobuf:"bytes,1,opt,name=otp_secret,json=otpSecret,proto3" json:"otp_secret,omitempty"`
	// The temp_otp_secret is the temporary secret key used to validate the OTP code and will replace the otp_secret in two phase commits.
	TempOtpSecret string `protobuf:"bytes,2,opt,name=temp_otp_secret,json=tempOtpSecret,proto3" json:"temp_otp_secret,omitempty"`
	// The recovery_codes are the codes that can be used to recover the account.
	RecoveryCodes []string `protobuf:"bytes,3,rep,name=recovery_codes,json=recoveryCodes,proto3" json:"recovery_codes,omitempty"`
	// The temp_recovery_codes are the temporary codes that will replace the recovery_codes in two phase commits.
	TempRecoveryCodes []string `protobuf:"bytes,4,rep,name=temp_recovery_codes,json=tempRecoveryCodes,proto3" json:"temp_recovery_codes,omitempty"`
}

func (x *MFAConfig) Reset() {
	*x = MFAConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MFAConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MFAConfig) ProtoMessage() {}

func (x *MFAConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MFAConfig.ProtoReflect.Descriptor instead.
func (*MFAConfig) Descriptor() ([]byte, []int) {
	return file_store_user_proto_rawDescGZIP(), []int{0}
}

func (x *MFAConfig) GetOtpSecret() string {
	if x != nil {
		return x.OtpSecret
	}
	return ""
}

func (x *MFAConfig) GetTempOtpSecret() string {
	if x != nil {
		return x.TempOtpSecret
	}
	return ""
}

func (x *MFAConfig) GetRecoveryCodes() []string {
	if x != nil {
		return x.RecoveryCodes
	}
	return nil
}

func (x *MFAConfig) GetTempRecoveryCodes() []string {
	if x != nil {
		return x.TempRecoveryCodes
	}
	return nil
}

var File_store_user_proto protoreflect.FileDescriptor

var file_store_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x09, 0x4d, 0x46, 0x41, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x74, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12,
	0x26, 0x0a, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x6f, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x4f, 0x74,
	0x70, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2e,
	0x0a, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x74, 0x65, 0x6d,
	0x70, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x42, 0x14,
	0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_user_proto_rawDescOnce sync.Once
	file_store_user_proto_rawDescData = file_store_user_proto_rawDesc
)

func file_store_user_proto_rawDescGZIP() []byte {
	file_store_user_proto_rawDescOnce.Do(func() {
		file_store_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_user_proto_rawDescData)
	})
	return file_store_user_proto_rawDescData
}

var file_store_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_store_user_proto_goTypes = []interface{}{
	(*MFAConfig)(nil), // 0: bytebase.store.MFAConfig
}
var file_store_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_store_user_proto_init() }
func file_store_user_proto_init() {
	if File_store_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MFAConfig); i {
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
			RawDescriptor: file_store_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_user_proto_goTypes,
		DependencyIndexes: file_store_user_proto_depIdxs,
		MessageInfos:      file_store_user_proto_msgTypes,
	}.Build()
	File_store_user_proto = out.File
	file_store_user_proto_rawDesc = nil
	file_store_user_proto_goTypes = nil
	file_store_user_proto_depIdxs = nil
}