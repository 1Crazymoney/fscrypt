// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/metadata.proto

package metadata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Specifies the method in which an outside secret is obtained for a Protector
type SourceType int32

const (
	SourceType_default           SourceType = 0
	SourceType_pam_passphrase    SourceType = 1
	SourceType_custom_passphrase SourceType = 2
	SourceType_raw_key           SourceType = 3
)

var SourceType_name = map[int32]string{
	0: "default",
	1: "pam_passphrase",
	2: "custom_passphrase",
	3: "raw_key",
}
var SourceType_value = map[string]int32{
	"default":           0,
	"pam_passphrase":    1,
	"custom_passphrase": 2,
	"raw_key":           3,
}

func (x SourceType) String() string {
	return proto.EnumName(SourceType_name, int32(x))
}
func (SourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_metadata_ff272c42a9f0f3d3, []int{0}
}

// Type of encryption; should match declarations of unix.FSCRYPT_MODE
type EncryptionOptions_Mode int32

const (
	EncryptionOptions_default     EncryptionOptions_Mode = 0
	EncryptionOptions_AES_256_XTS EncryptionOptions_Mode = 1
	EncryptionOptions_AES_256_GCM EncryptionOptions_Mode = 2
	EncryptionOptions_AES_256_CBC EncryptionOptions_Mode = 3
	EncryptionOptions_AES_256_CTS EncryptionOptions_Mode = 4
	EncryptionOptions_AES_128_CBC EncryptionOptions_Mode = 5
	EncryptionOptions_AES_128_CTS EncryptionOptions_Mode = 6
	EncryptionOptions_Adiantum    EncryptionOptions_Mode = 9
)

var EncryptionOptions_Mode_name = map[int32]string{
	0: "default",
	1: "AES_256_XTS",
	2: "AES_256_GCM",
	3: "AES_256_CBC",
	4: "AES_256_CTS",
	5: "AES_128_CBC",
	6: "AES_128_CTS",
	9: "Adiantum",
}
var EncryptionOptions_Mode_value = map[string]int32{
	"default":     0,
	"AES_256_XTS": 1,
	"AES_256_GCM": 2,
	"AES_256_CBC": 3,
	"AES_256_CTS": 4,
	"AES_128_CBC": 5,
	"AES_128_CTS": 6,
	"Adiantum":    9,
}

func (x EncryptionOptions_Mode) String() string {
	return proto.EnumName(EncryptionOptions_Mode_name, int32(x))
}
func (EncryptionOptions_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_metadata_ff272c42a9f0f3d3, []int{3, 0}
}

// Cost parameters to be used in our hashing functions.
type HashingCosts struct {
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Memory               int64    `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Parallelism          int64    `protobuf:"varint,4,opt,name=parallelism,proto3" json:"parallelism,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashingCosts) Reset()         { *m = HashingCosts{} }
func (m *HashingCosts) String() string { return proto.CompactTextString(m) }
func (*HashingCosts) ProtoMessage()    {}
func (*HashingCosts) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_ff272c42a9f0f3d3, []int{0}
}
func (m *HashingCosts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashingCosts.Unmarshal(m, b)
}
func (m *HashingCosts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashingCosts.Marshal(b, m, deterministic)
}
func (dst *HashingCosts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashingCosts.Merge(dst, src)
}
func (m *HashingCosts) XXX_Size() int {
	return xxx_messageInfo_HashingCosts.Size(m)
}
func (m *HashingCosts) XXX_DiscardUnknown() {
	xxx_messageInfo_HashingCosts.DiscardUnknown(m)
}

var xxx_messageInfo_HashingCosts proto.InternalMessageInfo

func (m *HashingCosts) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *HashingCosts) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *HashingCosts) GetParallelism() int64 {
	if m != nil {
		return m.Parallelism
	}
	return 0
}

// This structure is used for our authenticated wrapping/unwrapping of keys.
type WrappedKeyData struct {
	IV                   []byte   `protobuf:"bytes,1,opt,name=IV,proto3" json:"IV,omitempty"`
	EncryptedKey         []byte   `protobuf:"bytes,2,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
	Hmac                 []byte   `protobuf:"bytes,3,opt,name=hmac,proto3" json:"hmac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WrappedKeyData) Reset()         { *m = WrappedKeyData{} }
func (m *WrappedKeyData) String() string { return proto.CompactTextString(m) }
func (*WrappedKeyData) ProtoMessage()    {}
func (*WrappedKeyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_ff272c42a9f0f3d3, []int{1}
}
func (m *WrappedKeyData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WrappedKeyData.Unmarshal(m, b)
}
func (m *WrappedKeyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WrappedKeyData.Marshal(b, m, deterministic)
}
func (dst *WrappedKeyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedKeyData.Merge(dst, src)
}
func (m *WrappedKeyData) XXX_Size() int {
	return xxx_messageInfo_WrappedKeyData.Size(m)
}
func (m *WrappedKeyData) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedKeyData.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedKeyData proto.InternalMessageInfo

func (m *WrappedKeyData) GetIV() []byte {
	if m != nil {
		return m.IV
	}
	return nil
}

func (m *WrappedKeyData) GetEncryptedKey() []byte {
	if m != nil {
		return m.EncryptedKey
	}
	return nil
}

func (m *WrappedKeyData) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

// The associated data for each protector
type ProtectorData struct {
	ProtectorDescriptor string     `protobuf:"bytes,1,opt,name=protector_descriptor,json=protectorDescriptor,proto3" json:"protector_descriptor,omitempty"`
	Source              SourceType `protobuf:"varint,2,opt,name=source,proto3,enum=metadata.SourceType" json:"source,omitempty"`
	// These are only used by some of the protector types
	Name                 string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Costs                *HashingCosts   `protobuf:"bytes,4,opt,name=costs,proto3" json:"costs,omitempty"`
	Salt                 []byte          `protobuf:"bytes,5,opt,name=salt,proto3" json:"salt,omitempty"`
	Uid                  int64           `protobuf:"varint,6,opt,name=uid,proto3" json:"uid,omitempty"`
	WrappedKey           *WrappedKeyData `protobuf:"bytes,7,opt,name=wrapped_key,json=wrappedKey,proto3" json:"wrapped_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProtectorData) Reset()         { *m = ProtectorData{} }
func (m *ProtectorData) String() string { return proto.CompactTextString(m) }
func (*ProtectorData) ProtoMessage()    {}
func (*ProtectorData) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_ff272c42a9f0f3d3, []int{2}
}
func (m *ProtectorData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtectorData.Unmarshal(m, b)
}
func (m *ProtectorData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtectorData.Marshal(b, m, deterministic)
}
func (dst *ProtectorData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtectorData.Merge(dst, src)
}
func (m *ProtectorData) XXX_Size() int {
	return xxx_messageInfo_ProtectorData.Size(m)
}
func (m *ProtectorData) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtectorData.DiscardUnknown(m)
}

var xxx_messageInfo_ProtectorData proto.InternalMessageInfo

func (m *ProtectorData) GetProtectorDescriptor() string {
	if m != nil {
		return m.ProtectorDescriptor
	}
	return ""
}

func (m *ProtectorData) GetSource() SourceType {
	if m != nil {
		return m.Source
	}
	return SourceType_default
}

func (m *ProtectorData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtectorData) GetCosts() *HashingCosts {
	if m != nil {
		return m.Costs
	}
	return nil
}

func (m *ProtectorData) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *ProtectorData) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ProtectorData) GetWrappedKey() *WrappedKeyData {
	if m != nil {
		return m.WrappedKey
	}
	return nil
}

// Encryption policy specifics, corresponds to the fscrypt_policy struct
type EncryptionOptions struct {
	Padding              int64                  `protobuf:"varint,1,opt,name=padding,proto3" json:"padding,omitempty"`
	Contents             EncryptionOptions_Mode `protobuf:"varint,2,opt,name=contents,proto3,enum=metadata.EncryptionOptions_Mode" json:"contents,omitempty"`
	Filenames            EncryptionOptions_Mode `protobuf:"varint,3,opt,name=filenames,proto3,enum=metadata.EncryptionOptions_Mode" json:"filenames,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EncryptionOptions) Reset()         { *m = EncryptionOptions{} }
func (m *EncryptionOptions) String() string { return proto.CompactTextString(m) }
func (*EncryptionOptions) ProtoMessage()    {}
func (*EncryptionOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_ff272c42a9f0f3d3, []int{3}
}
func (m *EncryptionOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptionOptions.Unmarshal(m, b)
}
func (m *EncryptionOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptionOptions.Marshal(b, m, deterministic)
}
func (dst *EncryptionOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptionOptions.Merge(dst, src)
}
func (m *EncryptionOptions) XXX_Size() int {
	return xxx_messageInfo_EncryptionOptions.Size(m)
}
func (m *EncryptionOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptionOptions.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptionOptions proto.InternalMessageInfo

func (m *EncryptionOptions) GetPadding() int64 {
	if m != nil {
		return m.Padding
	}
	return 0
}

func (m *EncryptionOptions) GetContents() EncryptionOptions_Mode {
	if m != nil {
		return m.Contents
	}
	return EncryptionOptions_default
}

func (m *EncryptionOptions) GetFilenames() EncryptionOptions_Mode {
	if m != nil {
		return m.Filenames
	}
	return EncryptionOptions_default
}

type WrappedPolicyKey struct {
	ProtectorDescriptor  string          `protobuf:"bytes,1,opt,name=protector_descriptor,json=protectorDescriptor,proto3" json:"protector_descriptor,omitempty"`
	WrappedKey           *WrappedKeyData `protobuf:"bytes,2,opt,name=wrapped_key,json=wrappedKey,proto3" json:"wrapped_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WrappedPolicyKey) Reset()         { *m = WrappedPolicyKey{} }
func (m *WrappedPolicyKey) String() string { return proto.CompactTextString(m) }
func (*WrappedPolicyKey) ProtoMessage()    {}
func (*WrappedPolicyKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_ff272c42a9f0f3d3, []int{4}
}
func (m *WrappedPolicyKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WrappedPolicyKey.Unmarshal(m, b)
}
func (m *WrappedPolicyKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WrappedPolicyKey.Marshal(b, m, deterministic)
}
func (dst *WrappedPolicyKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedPolicyKey.Merge(dst, src)
}
func (m *WrappedPolicyKey) XXX_Size() int {
	return xxx_messageInfo_WrappedPolicyKey.Size(m)
}
func (m *WrappedPolicyKey) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedPolicyKey.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedPolicyKey proto.InternalMessageInfo

func (m *WrappedPolicyKey) GetProtectorDescriptor() string {
	if m != nil {
		return m.ProtectorDescriptor
	}
	return ""
}

func (m *WrappedPolicyKey) GetWrappedKey() *WrappedKeyData {
	if m != nil {
		return m.WrappedKey
	}
	return nil
}

// The associated data for each policy
type PolicyData struct {
	KeyDescriptor        string              `protobuf:"bytes,1,opt,name=key_descriptor,json=keyDescriptor,proto3" json:"key_descriptor,omitempty"`
	Options              *EncryptionOptions  `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	WrappedPolicyKeys    []*WrappedPolicyKey `protobuf:"bytes,3,rep,name=wrapped_policy_keys,json=wrappedPolicyKeys,proto3" json:"wrapped_policy_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PolicyData) Reset()         { *m = PolicyData{} }
func (m *PolicyData) String() string { return proto.CompactTextString(m) }
func (*PolicyData) ProtoMessage()    {}
func (*PolicyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_ff272c42a9f0f3d3, []int{5}
}
func (m *PolicyData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyData.Unmarshal(m, b)
}
func (m *PolicyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyData.Marshal(b, m, deterministic)
}
func (dst *PolicyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyData.Merge(dst, src)
}
func (m *PolicyData) XXX_Size() int {
	return xxx_messageInfo_PolicyData.Size(m)
}
func (m *PolicyData) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyData.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyData proto.InternalMessageInfo

func (m *PolicyData) GetKeyDescriptor() string {
	if m != nil {
		return m.KeyDescriptor
	}
	return ""
}

func (m *PolicyData) GetOptions() *EncryptionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *PolicyData) GetWrappedPolicyKeys() []*WrappedPolicyKey {
	if m != nil {
		return m.WrappedPolicyKeys
	}
	return nil
}

// Data stored in the config file
type Config struct {
	Source                    SourceType         `protobuf:"varint,1,opt,name=source,proto3,enum=metadata.SourceType" json:"source,omitempty"`
	HashCosts                 *HashingCosts      `protobuf:"bytes,2,opt,name=hash_costs,json=hashCosts,proto3" json:"hash_costs,omitempty"`
	Compatibility             string             `protobuf:"bytes,3,opt,name=compatibility,proto3" json:"compatibility,omitempty"`
	Options                   *EncryptionOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	UseFsKeyringForV1Policies bool               `protobuf:"varint,5,opt,name=use_fs_keyring_for_v1_policies,json=useFsKeyringForV1Policies,proto3" json:"use_fs_keyring_for_v1_policies,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}           `json:"-"`
	XXX_unrecognized          []byte             `json:"-"`
	XXX_sizecache             int32              `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_ff272c42a9f0f3d3, []int{6}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetSource() SourceType {
	if m != nil {
		return m.Source
	}
	return SourceType_default
}

func (m *Config) GetHashCosts() *HashingCosts {
	if m != nil {
		return m.HashCosts
	}
	return nil
}

func (m *Config) GetCompatibility() string {
	if m != nil {
		return m.Compatibility
	}
	return ""
}

func (m *Config) GetOptions() *EncryptionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Config) GetUseFsKeyringForV1Policies() bool {
	if m != nil {
		return m.UseFsKeyringForV1Policies
	}
	return false
}

func init() {
	proto.RegisterType((*HashingCosts)(nil), "metadata.HashingCosts")
	proto.RegisterType((*WrappedKeyData)(nil), "metadata.WrappedKeyData")
	proto.RegisterType((*ProtectorData)(nil), "metadata.ProtectorData")
	proto.RegisterType((*EncryptionOptions)(nil), "metadata.EncryptionOptions")
	proto.RegisterType((*WrappedPolicyKey)(nil), "metadata.WrappedPolicyKey")
	proto.RegisterType((*PolicyData)(nil), "metadata.PolicyData")
	proto.RegisterType((*Config)(nil), "metadata.Config")
	proto.RegisterEnum("metadata.SourceType", SourceType_name, SourceType_value)
	proto.RegisterEnum("metadata.EncryptionOptions_Mode", EncryptionOptions_Mode_name, EncryptionOptions_Mode_value)
}

func init() { proto.RegisterFile("metadata/metadata.proto", fileDescriptor_metadata_ff272c42a9f0f3d3) }

var fileDescriptor_metadata_ff272c42a9f0f3d3 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6b, 0x13, 0x41,
	0x10, 0xc7, 0xbd, 0x4b, 0x9a, 0x1f, 0x93, 0x1f, 0x5e, 0xb7, 0xb5, 0x9e, 0x0a, 0x12, 0xa2, 0x42,
	0x91, 0x52, 0x49, 0xa4, 0xa2, 0x20, 0x42, 0x4d, 0x5b, 0xad, 0xa5, 0x58, 0x2f, 0xa1, 0x2a, 0x08,
	0xc7, 0xf6, 0x6e, 0x93, 0x2c, 0xbd, 0xbb, 0x5d, 0x76, 0x37, 0x86, 0x7b, 0xf3, 0xcd, 0x27, 0x9f,
	0xfc, 0x5b, 0xf4, 0xef, 0x93, 0xdd, 0xcb, 0xcf, 0x16, 0x4a, 0xeb, 0xcb, 0x31, 0xfb, 0xdd, 0xd9,
	0x99, 0xd9, 0xcf, 0xce, 0x1c, 0xdc, 0x8d, 0x89, 0xc2, 0x21, 0x56, 0xf8, 0xd9, 0xd4, 0xd8, 0xe6,
	0x82, 0x29, 0x86, 0x4a, 0xd3, 0x75, 0xf3, 0x1b, 0x54, 0xdf, 0x63, 0x39, 0xa4, 0xc9, 0xa0, 0xc3,
	0xa4, 0x92, 0x08, 0x41, 0x5e, 0xd1, 0x98, 0xb8, 0x76, 0xc3, 0xda, 0xcc, 0x79, 0xc6, 0x46, 0x1b,
	0x50, 0x88, 0x49, 0xcc, 0x44, 0xea, 0xe6, 0x8c, 0x3a, 0x59, 0xa1, 0x06, 0x54, 0x38, 0x16, 0x38,
	0x8a, 0x48, 0x44, 0x65, 0xec, 0xe6, 0xcd, 0xe6, 0xa2, 0xd4, 0xfc, 0x0a, 0xf5, 0xcf, 0x02, 0x73,
	0x4e, 0xc2, 0x23, 0x92, 0xee, 0x61, 0x85, 0x51, 0x1d, 0xec, 0xc3, 0x53, 0xd7, 0x6a, 0x58, 0x9b,
	0x55, 0xcf, 0x3e, 0x3c, 0x45, 0x8f, 0xa0, 0x46, 0x92, 0x40, 0xa4, 0x5c, 0x91, 0xd0, 0x3f, 0x27,
	0xa9, 0x49, 0x5c, 0xf5, 0xaa, 0x33, 0xf1, 0x88, 0xa4, 0xba, 0xa8, 0x61, 0x8c, 0x03, 0x93, 0xbe,
	0xea, 0x19, 0xbb, 0xf9, 0xdb, 0x86, 0xda, 0x89, 0x60, 0x8a, 0x04, 0x8a, 0x09, 0x13, 0xba, 0x05,
	0xeb, 0x7c, 0x2a, 0xf8, 0x21, 0x91, 0x81, 0xa0, 0x5c, 0x31, 0x61, 0x92, 0x95, 0xbd, 0xb5, 0xd9,
	0xde, 0xde, 0x6c, 0x0b, 0x6d, 0x41, 0x41, 0xb2, 0x91, 0x08, 0xb2, 0xfb, 0xd6, 0xdb, 0xeb, 0xdb,
	0x33, 0x50, 0x5d, 0xa3, 0xf7, 0x52, 0x4e, 0xbc, 0x89, 0x8f, 0x2e, 0x23, 0xc1, 0x31, 0x31, 0x65,
	0x94, 0x3d, 0x63, 0xa3, 0x2d, 0x58, 0x09, 0x34, 0x38, 0x73, 0xfb, 0x4a, 0x7b, 0x63, 0x1e, 0x60,
	0x11, 0xab, 0x97, 0x39, 0xe9, 0x08, 0x12, 0x47, 0xca, 0x5d, 0xc9, 0x2e, 0xa2, 0x6d, 0xe4, 0x40,
	0x6e, 0x44, 0x43, 0xb7, 0x60, 0xe8, 0x69, 0x13, 0xbd, 0x82, 0xca, 0x38, 0xa3, 0x66, 0x88, 0x14,
	0x4d, 0x64, 0x77, 0x1e, 0x79, 0x19, 0xa9, 0x07, 0xe3, 0xd9, 0xba, 0xf9, 0xc7, 0x86, 0xd5, 0xfd,
	0x0c, 0x1d, 0x65, 0xc9, 0x47, 0xf3, 0x95, 0xc8, 0x85, 0x22, 0xc7, 0x61, 0x48, 0x93, 0x81, 0x81,
	0x91, 0xf3, 0xa6, 0x4b, 0xf4, 0x1a, 0x4a, 0x01, 0x4b, 0x14, 0x49, 0x94, 0x9c, 0x20, 0x68, 0xcc,
	0xf3, 0x5c, 0x0a, 0xb4, 0x7d, 0xcc, 0x42, 0xe2, 0xcd, 0x4e, 0xa0, 0x37, 0x50, 0xee, 0xd3, 0x88,
	0x68, 0x10, 0xd2, 0x50, 0xb9, 0xce, 0xf1, 0xf9, 0x91, 0xe6, 0x4f, 0x0b, 0xf2, 0x5a, 0x43, 0x15,
	0x28, 0x86, 0xa4, 0x8f, 0x47, 0x91, 0x72, 0x6e, 0xa1, 0xdb, 0x50, 0xd9, 0xdd, 0xef, 0xfa, 0xed,
	0x9d, 0x17, 0xfe, 0x97, 0x5e, 0xd7, 0xb1, 0x16, 0x85, 0x77, 0x9d, 0x63, 0xc7, 0x5e, 0x14, 0x3a,
	0x6f, 0x3b, 0x4e, 0x6e, 0x49, 0xe8, 0x75, 0x9d, 0xfc, 0x54, 0x68, 0xb5, 0x5f, 0x1a, 0x8f, 0x95,
	0x25, 0xa1, 0xd7, 0x75, 0x0a, 0xa8, 0x0a, 0xa5, 0xdd, 0x90, 0xe2, 0x44, 0x8d, 0x62, 0xa7, 0xdc,
	0xfc, 0x61, 0x81, 0x33, 0xc1, 0x7a, 0xc2, 0x22, 0x1a, 0xa4, 0xba, 0xed, 0xfe, 0xa3, 0xa1, 0x2e,
	0x3c, 0x9d, 0x7d, 0x83, 0xa7, 0xfb, 0x6b, 0x01, 0x64, 0xb9, 0x4d, 0x37, 0x3f, 0x81, 0xfa, 0x39,
	0x49, 0x2f, 0xa7, 0xad, 0x9d, 0x93, 0x74, 0x21, 0xe1, 0x0e, 0x14, 0x59, 0x46, 0x77, 0x92, 0xec,
	0xc1, 0x15, 0x0f, 0xe0, 0x4d, 0x7d, 0xd1, 0x07, 0x58, 0x9b, 0xd6, 0xc9, 0x4d, 0x4e, 0x5d, 0xae,
	0x7e, 0xc3, 0xdc, 0x66, 0xa5, 0x7d, 0xff, 0x52, 0xbd, 0x33, 0x26, 0xde, 0xea, 0xf8, 0x82, 0x22,
	0x9b, 0xbf, 0x6c, 0x28, 0x74, 0x58, 0xd2, 0xa7, 0x83, 0x85, 0x79, 0xb2, 0xae, 0x31, 0x4f, 0x3b,
	0x00, 0x43, 0x2c, 0x87, 0x7e, 0x36, 0x40, 0xf6, 0x95, 0x03, 0x54, 0xd6, 0x9e, 0xd9, 0x2f, 0xea,
	0x31, 0xd4, 0x02, 0x16, 0x73, 0xac, 0xe8, 0x19, 0x8d, 0xa8, 0x4a, 0x27, 0xf3, 0xb8, 0x2c, 0x2e,
	0x82, 0xc9, 0xdf, 0x00, 0xcc, 0x2e, 0x3c, 0x1c, 0x49, 0xe2, 0xf7, 0xa5, 0x06, 0x22, 0x68, 0x32,
	0xf0, 0xfb, 0x4c, 0xf8, 0xdf, 0x5b, 0x19, 0x26, 0x4a, 0xa4, 0x99, 0xdd, 0x92, 0x77, 0x6f, 0x24,
	0xc9, 0x81, 0x3c, 0xca, 0x7c, 0x0e, 0x98, 0x38, 0x6d, 0x9d, 0x4c, 0x1c, 0x9e, 0x7e, 0x02, 0x98,
	0x5f, 0x76, 0xb9, 0xb5, 0x11, 0xd4, 0x39, 0x8e, 0x7d, 0x8e, 0xa5, 0xe4, 0x43, 0x81, 0x25, 0x71,
	0x2c, 0x74, 0x07, 0x56, 0x83, 0x91, 0x54, 0x6c, 0x49, 0xb6, 0xf5, 0x39, 0x81, 0xc7, 0xba, 0x0a,
	0x27, 0x77, 0x56, 0x30, 0xbf, 0xed, 0xe7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x31, 0x82,
	0x34, 0xd1, 0x05, 0x00, 0x00,
}
