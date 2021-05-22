// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.6.1
// source: ca.proto

package pb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "github.com/gogo/googleapis/google/api"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CertProfileInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label   string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Profile string `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *CertProfileInfoRequest) Reset() {
	*x = CertProfileInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertProfileInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertProfileInfoRequest) ProtoMessage() {}

func (x *CertProfileInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertProfileInfoRequest.ProtoReflect.Descriptor instead.
func (*CertProfileInfoRequest) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{0}
}

func (x *CertProfileInfoRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *CertProfileInfoRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

// CertProfileInfo is the response for an Profile Info API request
type CertProfileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer  string       `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Profile *CertProfile `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *CertProfileInfo) Reset() {
	*x = CertProfileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertProfileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertProfileInfo) ProtoMessage() {}

func (x *CertProfileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertProfileInfo.ProtoReflect.Descriptor instead.
func (*CertProfileInfo) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{1}
}

func (x *CertProfileInfo) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *CertProfileInfo) GetProfile() *CertProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

// CertificateBundle provides certificate and its issuers
type CertificateBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Certificate provides the certificate in PEM format
	Certificate string `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// Intermediates provides the intermediate CA certificates bundle in PEM format
	Intermediates string `protobuf:"bytes,2,opt,name=intermediates,proto3" json:"intermediates,omitempty"`
	// Root provides the Root CA certifica in PEM format
	Root string `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
}

func (x *CertificateBundle) Reset() {
	*x = CertificateBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateBundle) ProtoMessage() {}

func (x *CertificateBundle) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateBundle.ProtoReflect.Descriptor instead.
func (*CertificateBundle) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{2}
}

func (x *CertificateBundle) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *CertificateBundle) GetIntermediates() string {
	if x != nil {
		return x.Intermediates
	}
	return ""
}

func (x *CertificateBundle) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

// IssuerInfo provides Issuer information
type IssuerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Certificate provides the certificate in PEM format
	Certificate string `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// Intermediates provides the intermediate CA certificates bundle in PEM format
	Intermediates string `protobuf:"bytes,2,opt,name=intermediates,proto3" json:"intermediates,omitempty"`
	// Root provides the Root CA certificate in PEM format
	Root string `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	// Label specifies the Issuer's label
	Label string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *IssuerInfo) Reset() {
	*x = IssuerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssuerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssuerInfo) ProtoMessage() {}

func (x *IssuerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssuerInfo.ProtoReflect.Descriptor instead.
func (*IssuerInfo) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{3}
}

func (x *IssuerInfo) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *IssuerInfo) GetIntermediates() string {
	if x != nil {
		return x.Intermediates
	}
	return ""
}

func (x *IssuerInfo) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *IssuerInfo) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// IssuersInfoResponse provides response for Issuers Info request
type IssuersInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuers []*IssuerInfo `protobuf:"bytes,1,rep,name=issuers,proto3" json:"issuers,omitempty"`
}

func (x *IssuersInfoResponse) Reset() {
	*x = IssuersInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssuersInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssuersInfoResponse) ProtoMessage() {}

func (x *IssuersInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssuersInfoResponse.ProtoReflect.Descriptor instead.
func (*IssuersInfoResponse) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{4}
}

func (x *IssuersInfoResponse) GetIssuers() []*IssuerInfo {
	if x != nil {
		return x.Issuers
	}
	return nil
}

// SignCertificateRequest specifies certificate request
type SignCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RequestFormat provides the certificate request format: CSR, CMS
	RequestFormat EncodingFormat `protobuf:"varint,1,opt,name=request_format,json=requestFormat,proto3,enum=pb.EncodingFormat" json:"request_format,omitempty"`
	// Request provides the certificate request
	Request string `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	// Profile specifies the certificate profile: client, server, spiffe
	Profile string `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	// IssuerLabel specifies which Issuer to be appointed to sign the request
	IssuerLabel string `protobuf:"bytes,4,opt,name=issuer_label,json=issuerLabel,proto3" json:"issuer_label,omitempty"`
	// WithBundle specifies whether to include an "optimal" certificate bundle along with the certificate
	WithBundle bool `protobuf:"varint,5,opt,name=with_bundle,json=withBundle,proto3" json:"with_bundle,omitempty"`
	// Token provides the authorization token for the request
	Token string `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	// San specifies Subject Alternative Names
	San []string `protobuf:"bytes,7,rep,name=san,proto3" json:"san,omitempty"`
}

func (x *SignCertificateRequest) Reset() {
	*x = SignCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignCertificateRequest) ProtoMessage() {}

func (x *SignCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignCertificateRequest.ProtoReflect.Descriptor instead.
func (*SignCertificateRequest) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{5}
}

func (x *SignCertificateRequest) GetRequestFormat() EncodingFormat {
	if x != nil {
		return x.RequestFormat
	}
	return EncodingFormat_PEM
}

func (x *SignCertificateRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *SignCertificateRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *SignCertificateRequest) GetIssuerLabel() string {
	if x != nil {
		return x.IssuerLabel
	}
	return ""
}

func (x *SignCertificateRequest) GetWithBundle() bool {
	if x != nil {
		return x.WithBundle
	}
	return false
}

func (x *SignCertificateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SignCertificateRequest) GetSan() []string {
	if x != nil {
		return x.San
	}
	return nil
}

var File_ca_proto protoreflect.FileDescriptor

var file_ca_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a,
	0x70, 0x6b, 0x69, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x16, 0x43, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0x54, 0x0a, 0x0f, 0x43, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x6f, 0x0a, 0x11, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x22, 0x7e, 0x0a, 0x0a, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x3f, 0x0a, 0x13, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x62, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x22, 0xf3, 0x01, 0x0a, 0x16, 0x53, 0x69, 0x67, 0x6e,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x0d,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x62, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x61, 0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x73, 0x61, 0x6e, 0x32, 0xa8, 0x02,
	0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x61, 0x2f, 0x63, 0x73, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x12, 0x5f, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x2f, 0x63, 0x73, 0x72, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x52, 0x0a, 0x07, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6b, 0x73, 0x70, 0x61, 0x6e, 0x64, 0x2f, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ca_proto_rawDescOnce sync.Once
	file_ca_proto_rawDescData = file_ca_proto_rawDesc
)

func file_ca_proto_rawDescGZIP() []byte {
	file_ca_proto_rawDescOnce.Do(func() {
		file_ca_proto_rawDescData = protoimpl.X.CompressGZIP(file_ca_proto_rawDescData)
	})
	return file_ca_proto_rawDescData
}

var file_ca_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ca_proto_goTypes = []interface{}{
	(*CertProfileInfoRequest)(nil), // 0: pb.CertProfileInfoRequest
	(*CertProfileInfo)(nil),        // 1: pb.CertProfileInfo
	(*CertificateBundle)(nil),      // 2: pb.CertificateBundle
	(*IssuerInfo)(nil),             // 3: pb.IssuerInfo
	(*IssuersInfoResponse)(nil),    // 4: pb.IssuersInfoResponse
	(*SignCertificateRequest)(nil), // 5: pb.SignCertificateRequest
	(*CertProfile)(nil),            // 6: pb.CertProfile
	(EncodingFormat)(0),            // 7: pb.EncodingFormat
	(*empty.Empty)(nil),            // 8: google.protobuf.Empty
}
var file_ca_proto_depIdxs = []int32{
	6, // 0: pb.CertProfileInfo.profile:type_name -> pb.CertProfile
	3, // 1: pb.IssuersInfoResponse.issuers:type_name -> pb.IssuerInfo
	7, // 2: pb.SignCertificateRequest.request_format:type_name -> pb.EncodingFormat
	0, // 3: pb.AuthorityService.ProfileInfo:input_type -> pb.CertProfileInfoRequest
	5, // 4: pb.AuthorityService.SignCertificate:input_type -> pb.SignCertificateRequest
	8, // 5: pb.AuthorityService.Issuers:input_type -> google.protobuf.Empty
	1, // 6: pb.AuthorityService.ProfileInfo:output_type -> pb.CertProfileInfo
	2, // 7: pb.AuthorityService.SignCertificate:output_type -> pb.CertificateBundle
	4, // 8: pb.AuthorityService.Issuers:output_type -> pb.IssuersInfoResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ca_proto_init() }
func file_ca_proto_init() {
	if File_ca_proto != nil {
		return
	}
	file_pkix_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ca_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertProfileInfoRequest); i {
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
		file_ca_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertProfileInfo); i {
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
		file_ca_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateBundle); i {
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
		file_ca_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssuerInfo); i {
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
		file_ca_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssuersInfoResponse); i {
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
		file_ca_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignCertificateRequest); i {
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
			RawDescriptor: file_ca_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ca_proto_goTypes,
		DependencyIndexes: file_ca_proto_depIdxs,
		MessageInfos:      file_ca_proto_msgTypes,
	}.Build()
	File_ca_proto = out.File
	file_ca_proto_rawDesc = nil
	file_ca_proto_goTypes = nil
	file_ca_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthorityServiceClient is the client API for AuthorityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorityServiceClient interface {
	// ProfileInfo returns the certificate profile info
	ProfileInfo(ctx context.Context, in *CertProfileInfoRequest, opts ...grpc.CallOption) (*CertProfileInfo, error)
	// SignCertificate returns the certificate
	SignCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*CertificateBundle, error)
	// Issuers returns the issuing CAs
	Issuers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IssuersInfoResponse, error)
}

type authorityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorityServiceClient(cc grpc.ClientConnInterface) AuthorityServiceClient {
	return &authorityServiceClient{cc}
}

func (c *authorityServiceClient) ProfileInfo(ctx context.Context, in *CertProfileInfoRequest, opts ...grpc.CallOption) (*CertProfileInfo, error) {
	out := new(CertProfileInfo)
	err := c.cc.Invoke(ctx, "/pb.AuthorityService/ProfileInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) SignCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*CertificateBundle, error) {
	out := new(CertificateBundle)
	err := c.cc.Invoke(ctx, "/pb.AuthorityService/SignCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) Issuers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IssuersInfoResponse, error) {
	out := new(IssuersInfoResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthorityService/Issuers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorityServiceServer is the server API for AuthorityService service.
type AuthorityServiceServer interface {
	// ProfileInfo returns the certificate profile info
	ProfileInfo(context.Context, *CertProfileInfoRequest) (*CertProfileInfo, error)
	// SignCertificate returns the certificate
	SignCertificate(context.Context, *SignCertificateRequest) (*CertificateBundle, error)
	// Issuers returns the issuing CAs
	Issuers(context.Context, *empty.Empty) (*IssuersInfoResponse, error)
}

// UnimplementedAuthorityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorityServiceServer struct {
}

func (*UnimplementedAuthorityServiceServer) ProfileInfo(context.Context, *CertProfileInfoRequest) (*CertProfileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProfileInfo not implemented")
}
func (*UnimplementedAuthorityServiceServer) SignCertificate(context.Context, *SignCertificateRequest) (*CertificateBundle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignCertificate not implemented")
}
func (*UnimplementedAuthorityServiceServer) Issuers(context.Context, *empty.Empty) (*IssuersInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issuers not implemented")
}

func RegisterAuthorityServiceServer(s *grpc.Server, srv AuthorityServiceServer) {
	s.RegisterService(&_AuthorityService_serviceDesc, srv)
}

func _AuthorityService_ProfileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertProfileInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).ProfileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityService/ProfileInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).ProfileInfo(ctx, req.(*CertProfileInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_SignCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).SignCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityService/SignCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).SignCertificate(ctx, req.(*SignCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_Issuers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).Issuers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityService/Issuers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).Issuers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthorityService",
	HandlerType: (*AuthorityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProfileInfo",
			Handler:    _AuthorityService_ProfileInfo_Handler,
		},
		{
			MethodName: "SignCertificate",
			Handler:    _AuthorityService_SignCertificate_Handler,
		},
		{
			MethodName: "Issuers",
			Handler:    _AuthorityService_Issuers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ca.proto",
}