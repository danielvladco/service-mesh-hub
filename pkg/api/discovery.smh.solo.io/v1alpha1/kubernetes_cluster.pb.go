// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/kubernetes_cluster.proto

package v1alpha1

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Representation of a Kubernetes cluster that has been registered in Service Mesh Hub.
type KubernetesClusterSpec struct {
	// pointer to secret which contains the kubeconfig with information to connect to the remote cluster.
	SecretRef *types.ResourceRef `protobuf:"bytes,1,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// context to use within the kubeconfig pointed to by the above reference
	Context string `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	// version of kubernetes
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// cloud provider, empty if unknown
	Cloud string `protobuf:"bytes,4,opt,name=cloud,proto3" json:"cloud,omitempty"`
	// namespace to use when writing Service Mesh Hub resources to this cluster
	WriteNamespace       string   `protobuf:"bytes,5,opt,name=write_namespace,json=writeNamespace,proto3" json:"write_namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubernetesClusterSpec) Reset()         { *m = KubernetesClusterSpec{} }
func (m *KubernetesClusterSpec) String() string { return proto.CompactTextString(m) }
func (*KubernetesClusterSpec) ProtoMessage()    {}
func (*KubernetesClusterSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_cebb6fc6a7744031, []int{0}
}
func (m *KubernetesClusterSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesClusterSpec.Unmarshal(m, b)
}
func (m *KubernetesClusterSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesClusterSpec.Marshal(b, m, deterministic)
}
func (m *KubernetesClusterSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesClusterSpec.Merge(m, src)
}
func (m *KubernetesClusterSpec) XXX_Size() int {
	return xxx_messageInfo_KubernetesClusterSpec.Size(m)
}
func (m *KubernetesClusterSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesClusterSpec.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesClusterSpec proto.InternalMessageInfo

func (m *KubernetesClusterSpec) GetSecretRef() *types.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return nil
}

func (m *KubernetesClusterSpec) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *KubernetesClusterSpec) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *KubernetesClusterSpec) GetCloud() string {
	if m != nil {
		return m.Cloud
	}
	return ""
}

func (m *KubernetesClusterSpec) GetWriteNamespace() string {
	if m != nil {
		return m.WriteNamespace
	}
	return ""
}

func init() {
	proto.RegisterType((*KubernetesClusterSpec)(nil), "discovery.smh.solo.io.KubernetesClusterSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/kubernetes_cluster.proto", fileDescriptor_cebb6fc6a7744031)
}

var fileDescriptor_cebb6fc6a7744031 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x14, 0x84, 0x89, 0x5a, 0xa5, 0x2b, 0x28, 0x04, 0x0b, 0x41, 0x10, 0x8a, 0x17, 0x7b, 0xb0, 0x1b,
	0xaa, 0x57, 0x4f, 0x7a, 0x53, 0xf1, 0x90, 0xde, 0xbc, 0x94, 0x64, 0x3b, 0x69, 0x96, 0x26, 0x79,
	0xcb, 0xdb, 0xdd, 0xa8, 0xff, 0xd0, 0x9f, 0x25, 0x4d, 0x6c, 0xac, 0xe0, 0xc1, 0xe3, 0xcc, 0xec,
	0x7e, 0xc3, 0x1b, 0x31, 0x5f, 0x69, 0x57, 0xf8, 0x4c, 0x2a, 0xaa, 0x62, 0x4b, 0x25, 0x4d, 0x35,
	0xc5, 0x16, 0xdc, 0x68, 0x85, 0x69, 0x05, 0x5b, 0x4c, 0x0b, 0x9f, 0xc5, 0xa9, 0xd1, 0xf1, 0x52,
	0x5b, 0x45, 0x0d, 0xf8, 0x23, 0x6e, 0x66, 0x69, 0x69, 0x8a, 0x74, 0x16, 0xaf, 0x7d, 0x06, 0xae,
	0xe1, 0x60, 0x17, 0xaa, 0xf4, 0xd6, 0x81, 0xa5, 0x61, 0x72, 0x14, 0x8e, 0xfa, 0xc7, 0xd2, 0x56,
	0x85, 0xdc, 0x70, 0xa5, 0xa6, 0xf3, 0xeb, 0x3f, 0xc1, 0x8a, 0x18, 0x3f, 0x4c, 0x46, 0xde, 0x41,
	0x2e, 0x3f, 0x03, 0x31, 0x7a, 0xea, 0x1b, 0x1e, 0xba, 0x82, 0xb9, 0x81, 0x0a, 0xef, 0x84, 0xb0,
	0x50, 0x0c, 0xb7, 0x60, 0xe4, 0x51, 0x30, 0x0e, 0x26, 0xc7, 0x37, 0x17, 0x72, 0xc3, 0xd9, 0xad,
	0x93, 0x09, 0x2c, 0x79, 0x56, 0x48, 0x90, 0x27, 0xc3, 0xee, 0x43, 0x82, 0x3c, 0x8c, 0xc4, 0x91,
	0xa2, 0xda, 0xe1, 0xdd, 0x45, 0x7b, 0xe3, 0x60, 0x32, 0x4c, 0xb6, 0x72, 0x93, 0x34, 0x60, 0xab,
	0xa9, 0x8e, 0xf6, 0xbb, 0xe4, 0x5b, 0x86, 0x67, 0x62, 0xa0, 0x4a, 0xf2, 0xcb, 0xe8, 0xa0, 0xf5,
	0x3b, 0x11, 0x5e, 0x89, 0xd3, 0x37, 0xd6, 0x0e, 0x8b, 0x3a, 0xad, 0x60, 0x4d, 0xaa, 0x10, 0x0d,
	0xda, 0xfc, 0xa4, 0xb5, 0x5f, 0xb6, 0xee, 0xfd, 0xf3, 0xeb, 0xe3, 0x7f, 0x66, 0x36, 0xeb, 0xd5,
	0xef, 0xa9, 0x77, 0xcf, 0xe9, 0x27, 0xca, 0x0e, 0xdb, 0x7d, 0x6e, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x91, 0x7a, 0x7f, 0xfb, 0xbb, 0x01, 0x00, 0x00,
}
