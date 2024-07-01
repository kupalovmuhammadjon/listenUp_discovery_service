// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: user_interactions.proto

package user_interactions

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

type InteractEpisode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PodcastId       string `protobuf:"bytes,2,opt,name=podcast_id,json=podcastId,proto3" json:"podcast_id,omitempty"`
	EpisodeId       string `protobuf:"bytes,3,opt,name=episode_id,json=episodeId,proto3" json:"episode_id,omitempty"`
	InteractionType string `protobuf:"bytes,4,opt,name=interaction_type,json=interactionType,proto3" json:"interaction_type,omitempty"`
}

func (x *InteractEpisode) Reset() {
	*x = InteractEpisode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_interactions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractEpisode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractEpisode) ProtoMessage() {}

func (x *InteractEpisode) ProtoReflect() protoreflect.Message {
	mi := &file_user_interactions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractEpisode.ProtoReflect.Descriptor instead.
func (*InteractEpisode) Descriptor() ([]byte, []int) {
	return file_user_interactions_proto_rawDescGZIP(), []int{0}
}

func (x *InteractEpisode) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *InteractEpisode) GetPodcastId() string {
	if x != nil {
		return x.PodcastId
	}
	return ""
}

func (x *InteractEpisode) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

func (x *InteractEpisode) GetInteractionType() string {
	if x != nil {
		return x.InteractionType
	}
	return ""
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_interactions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_user_interactions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_user_interactions_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteLike struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PodcastId string `protobuf:"bytes,2,opt,name=podcast_id,json=podcastId,proto3" json:"podcast_id,omitempty"`
	EpisodeId string `protobuf:"bytes,3,opt,name=episode_id,json=episodeId,proto3" json:"episode_id,omitempty"`
}

func (x *DeleteLike) Reset() {
	*x = DeleteLike{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_interactions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLike) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLike) ProtoMessage() {}

func (x *DeleteLike) ProtoReflect() protoreflect.Message {
	mi := &file_user_interactions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLike.ProtoReflect.Descriptor instead.
func (*DeleteLike) Descriptor() ([]byte, []int) {
	return file_user_interactions_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteLike) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteLike) GetPodcastId() string {
	if x != nil {
		return x.PodcastId
	}
	return ""
}

func (x *DeleteLike) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_interactions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_user_interactions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_user_interactions_proto_rawDescGZIP(), []int{3}
}

func (x *Success) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_user_interactions_proto protoreflect.FileDescriptor

var file_user_interactions_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x70, 0x69, 0x73, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x63, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x69, 0x6b, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32,
	0xb2, 0x01, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x14, 0x4c, 0x69, 0x6b, 0x65, 0x45, 0x70, 0x69,
	0x73, 0x6f, 0x64, 0x65, 0x4f, 0x66, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x10, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x1a,
	0x03, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4c, 0x69, 0x6b, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4f,
	0x66, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x69, 0x6b, 0x65, 0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x70, 0x69, 0x73, 0x6f,
	0x64, 0x65, 0x4f, 0x66, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x1a, 0x03, 0x2e,
	0x49, 0x44, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_interactions_proto_rawDescOnce sync.Once
	file_user_interactions_proto_rawDescData = file_user_interactions_proto_rawDesc
)

func file_user_interactions_proto_rawDescGZIP() []byte {
	file_user_interactions_proto_rawDescOnce.Do(func() {
		file_user_interactions_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_interactions_proto_rawDescData)
	})
	return file_user_interactions_proto_rawDescData
}

var file_user_interactions_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_interactions_proto_goTypes = []interface{}{
	(*InteractEpisode)(nil), // 0: InteractEpisode
	(*ID)(nil),              // 1: ID
	(*DeleteLike)(nil),      // 2: DeleteLike
	(*Success)(nil),         // 3: Success
}
var file_user_interactions_proto_depIdxs = []int32{
	0, // 0: user_interactions.LikeEpisodeOfPodcast:input_type -> InteractEpisode
	2, // 1: user_interactions.DeleteLikeFromEpisodeOfPodcast:input_type -> DeleteLike
	0, // 2: user_interactions.ListenEpisodeOfPodcast:input_type -> InteractEpisode
	1, // 3: user_interactions.LikeEpisodeOfPodcast:output_type -> ID
	3, // 4: user_interactions.DeleteLikeFromEpisodeOfPodcast:output_type -> Success
	1, // 5: user_interactions.ListenEpisodeOfPodcast:output_type -> ID
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_interactions_proto_init() }
func file_user_interactions_proto_init() {
	if File_user_interactions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_interactions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InteractEpisode); i {
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
		file_user_interactions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_user_interactions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLike); i {
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
		file_user_interactions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
			RawDescriptor: file_user_interactions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_interactions_proto_goTypes,
		DependencyIndexes: file_user_interactions_proto_depIdxs,
		MessageInfos:      file_user_interactions_proto_msgTypes,
	}.Build()
	File_user_interactions_proto = out.File
	file_user_interactions_proto_rawDesc = nil
	file_user_interactions_proto_goTypes = nil
	file_user_interactions_proto_depIdxs = nil
}
