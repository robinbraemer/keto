// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: keto/acl/v1alpha1/acl.proto

package acl

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// RelationTuple relates an Object with a Subject.
//
// While a tuple reflects a relationship between Object
// and Subject, they do not completely define the effective ACLs.
type RelationTuple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace this relation tuple lives in.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The object related by this tuple.
	// It is naturally in the namespace of the tuple.
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	// The relation between an Object and a Subject.
	Relation string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	// The subject related by this tuple.
	// A Subject either represents a concrete subject id or
	// a SubjectSet that expands to more Subjects.
	Subject *Subject `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *RelationTuple) Reset() {
	*x = RelationTuple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keto_acl_v1alpha1_acl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationTuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationTuple) ProtoMessage() {}

func (x *RelationTuple) ProtoReflect() protoreflect.Message {
	mi := &file_keto_acl_v1alpha1_acl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationTuple.ProtoReflect.Descriptor instead.
func (*RelationTuple) Descriptor() ([]byte, []int) {
	return file_keto_acl_v1alpha1_acl_proto_rawDescGZIP(), []int{0}
}

func (x *RelationTuple) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RelationTuple) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *RelationTuple) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *RelationTuple) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

// Object represents a "resource/digital object" in a namespace.
// A RelationTuple relates a Subject to an Object.
type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace of the object.
	// This effectively is the namespace of the whole RelationTuple.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The object id.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keto_acl_v1alpha1_acl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_keto_acl_v1alpha1_acl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_keto_acl_v1alpha1_acl_proto_rawDescGZIP(), []int{1}
}

func (x *Object) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Object) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Subject is either a concrete subject id or
// a subject set expanding to more Subjects.
type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The reference of this abstract subject.
	//
	// Types that are assignable to Ref:
	//	*Subject_Id
	//	*Subject_Set
	Ref isSubject_Ref `protobuf_oneof:"ref"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keto_acl_v1alpha1_acl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_keto_acl_v1alpha1_acl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_keto_acl_v1alpha1_acl_proto_rawDescGZIP(), []int{2}
}

func (m *Subject) GetRef() isSubject_Ref {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (x *Subject) GetId() string {
	if x, ok := x.GetRef().(*Subject_Id); ok {
		return x.Id
	}
	return ""
}

func (x *Subject) GetSet() *SubjectSet {
	if x, ok := x.GetRef().(*Subject_Set); ok {
		return x.Set
	}
	return nil
}

type isSubject_Ref interface {
	isSubject_Ref()
}

type Subject_Id struct {
	// A concrete id of the subject.
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type Subject_Set struct {
	// A subject set that expands to more Subjects
	// (used for inheritance).
	Set *SubjectSet `protobuf:"bytes,2,opt,name=set,proto3,oneof"`
}

func (*Subject_Id) isSubject_Ref() {}

func (*Subject_Set) isSubject_Ref() {}

// SubjectSet refers to all subjects which have
// the same `relation` to an `object`.
// It is also used for inheritance.
type SubjectSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace of the object and relation
	// referenced in this subject set.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The object selected by the subjects.
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	// The relation to the object by the subjects.
	Relation string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *SubjectSet) Reset() {
	*x = SubjectSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keto_acl_v1alpha1_acl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubjectSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectSet) ProtoMessage() {}

func (x *SubjectSet) ProtoReflect() protoreflect.Message {
	mi := &file_keto_acl_v1alpha1_acl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectSet.ProtoReflect.Descriptor instead.
func (*SubjectSet) Descriptor() ([]byte, []int) {
	return file_keto_acl_v1alpha1_acl_proto_rawDescGZIP(), []int{3}
}

func (x *SubjectSet) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SubjectSet) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *SubjectSet) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

var File_keto_acl_v1alpha1_acl_proto protoreflect.FileDescriptor

var file_keto_acl_v1alpha1_acl_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6b,
	0x65, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x22, 0x97, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70,
	0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x6c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x36, 0x0a, 0x06, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x55, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x31, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b,
	0x65, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x03, 0x73,
	0x65, 0x74, 0x42, 0x05, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x22, 0x5e, 0x0a, 0x0a, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x85, 0x01, 0x0a, 0x18, 0x73, 0x68,
	0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x08, 0x41, 0x63, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x74, 0x6f,
	0x2f, 0x61, 0x63, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x63,
	0x6c, 0xaa, 0x02, 0x15, 0x4f, 0x72, 0x79, 0x2e, 0x4b, 0x65, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6c,
	0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x15, 0x4f, 0x72, 0x79, 0x5c,
	0x4b, 0x65, 0x74, 0x6f, 0x5c, 0x41, 0x63, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keto_acl_v1alpha1_acl_proto_rawDescOnce sync.Once
	file_keto_acl_v1alpha1_acl_proto_rawDescData = file_keto_acl_v1alpha1_acl_proto_rawDesc
)

func file_keto_acl_v1alpha1_acl_proto_rawDescGZIP() []byte {
	file_keto_acl_v1alpha1_acl_proto_rawDescOnce.Do(func() {
		file_keto_acl_v1alpha1_acl_proto_rawDescData = protoimpl.X.CompressGZIP(file_keto_acl_v1alpha1_acl_proto_rawDescData)
	})
	return file_keto_acl_v1alpha1_acl_proto_rawDescData
}

var file_keto_acl_v1alpha1_acl_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_keto_acl_v1alpha1_acl_proto_goTypes = []interface{}{
	(*RelationTuple)(nil), // 0: keto.acl.v1alpha1.RelationTuple
	(*Object)(nil),        // 1: keto.acl.v1alpha1.Object
	(*Subject)(nil),       // 2: keto.acl.v1alpha1.Subject
	(*SubjectSet)(nil),    // 3: keto.acl.v1alpha1.SubjectSet
}
var file_keto_acl_v1alpha1_acl_proto_depIdxs = []int32{
	2, // 0: keto.acl.v1alpha1.RelationTuple.subject:type_name -> keto.acl.v1alpha1.Subject
	3, // 1: keto.acl.v1alpha1.Subject.set:type_name -> keto.acl.v1alpha1.SubjectSet
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_keto_acl_v1alpha1_acl_proto_init() }
func file_keto_acl_v1alpha1_acl_proto_init() {
	if File_keto_acl_v1alpha1_acl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keto_acl_v1alpha1_acl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationTuple); i {
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
		file_keto_acl_v1alpha1_acl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_keto_acl_v1alpha1_acl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
		file_keto_acl_v1alpha1_acl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubjectSet); i {
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
	file_keto_acl_v1alpha1_acl_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Subject_Id)(nil),
		(*Subject_Set)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_keto_acl_v1alpha1_acl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_keto_acl_v1alpha1_acl_proto_goTypes,
		DependencyIndexes: file_keto_acl_v1alpha1_acl_proto_depIdxs,
		MessageInfos:      file_keto_acl_v1alpha1_acl_proto_msgTypes,
	}.Build()
	File_keto_acl_v1alpha1_acl_proto = out.File
	file_keto_acl_v1alpha1_acl_proto_rawDesc = nil
	file_keto_acl_v1alpha1_acl_proto_goTypes = nil
	file_keto_acl_v1alpha1_acl_proto_depIdxs = nil
}
