// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: msgforward.proto

package transmit

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SeTuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Desc        string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	OriginalUrl string `protobuf:"bytes,3,opt,name=originalUrl,proto3" json:"originalUrl,omitempty"`
	Url         string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	PicBase64   string `protobuf:"bytes,5,opt,name=picBase64,proto3" json:"picBase64,omitempty"`
	PicMd5      string `protobuf:"bytes,6,opt,name=picMd5,proto3" json:"picMd5,omitempty"`
}

func (x *SeTuRequest) Reset() {
	*x = SeTuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgforward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeTuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeTuRequest) ProtoMessage() {}

func (x *SeTuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_msgforward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeTuRequest.ProtoReflect.Descriptor instead.
func (*SeTuRequest) Descriptor() ([]byte, []int) {
	return file_msgforward_proto_rawDescGZIP(), []int{0}
}

func (x *SeTuRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SeTuRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *SeTuRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

func (x *SeTuRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SeTuRequest) GetPicBase64() string {
	if x != nil {
		return x.PicBase64
	}
	return ""
}

func (x *SeTuRequest) GetPicMd5() string {
	if x != nil {
		return x.PicMd5
	}
	return ""
}

type SeTuReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrMessage string `protobuf:"bytes,1,opt,name=errMessage,proto3" json:"errMessage,omitempty"`
}

func (x *SeTuReply) Reset() {
	*x = SeTuReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgforward_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeTuReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeTuReply) ProtoMessage() {}

func (x *SeTuReply) ProtoReflect() protoreflect.Message {
	mi := &file_msgforward_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeTuReply.ProtoReflect.Descriptor instead.
func (*SeTuReply) Descriptor() ([]byte, []int) {
	return file_msgforward_proto_rawDescGZIP(), []int{1}
}

func (x *SeTuReply) GetErrMessage() string {
	if x != nil {
		return x.ErrMessage
	}
	return ""
}

var File_msgforward_proto protoreflect.FileDescriptor

var file_msgforward_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x73, 0x67, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x22, 0xa1, 0x01, 0x0a,
	0x0b, 0x53, 0x65, 0x54, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x69,
	0x63, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x69, 0x63, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x69, 0x63, 0x4d,
	0x64, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x69, 0x63, 0x4d, 0x64, 0x35,
	0x22, 0x2b, 0x0a, 0x09, 0x53, 0x65, 0x54, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x65, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x47, 0x0a,
	0x0b, 0x53, 0x65, 0x74, 0x75, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x08,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x75, 0x54, 0x75, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6d, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x54, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x54, 0x75, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msgforward_proto_rawDescOnce sync.Once
	file_msgforward_proto_rawDescData = file_msgforward_proto_rawDesc
)

func file_msgforward_proto_rawDescGZIP() []byte {
	file_msgforward_proto_rawDescOnce.Do(func() {
		file_msgforward_proto_rawDescData = protoimpl.X.CompressGZIP(file_msgforward_proto_rawDescData)
	})
	return file_msgforward_proto_rawDescData
}

var file_msgforward_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msgforward_proto_goTypes = []interface{}{
	(*SeTuRequest)(nil), // 0: transmit.SeTuRequest
	(*SeTuReply)(nil),   // 1: transmit.SeTuReply
}
var file_msgforward_proto_depIdxs = []int32{
	0, // 0: transmit.SetuCourier.SendSuTu:input_type -> transmit.SeTuRequest
	1, // 1: transmit.SetuCourier.SendSuTu:output_type -> transmit.SeTuReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msgforward_proto_init() }
func file_msgforward_proto_init() {
	if File_msgforward_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msgforward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeTuRequest); i {
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
		file_msgforward_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeTuReply); i {
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
			RawDescriptor: file_msgforward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msgforward_proto_goTypes,
		DependencyIndexes: file_msgforward_proto_depIdxs,
		MessageInfos:      file_msgforward_proto_msgTypes,
	}.Build()
	File_msgforward_proto = out.File
	file_msgforward_proto_rawDesc = nil
	file_msgforward_proto_goTypes = nil
	file_msgforward_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SetuCourierClient is the client API for SetuCourier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SetuCourierClient interface {
	SendSuTu(ctx context.Context, in *SeTuRequest, opts ...grpc.CallOption) (*SeTuReply, error)
}

type setuCourierClient struct {
	cc grpc.ClientConnInterface
}

func NewSetuCourierClient(cc grpc.ClientConnInterface) SetuCourierClient {
	return &setuCourierClient{cc}
}

func (c *setuCourierClient) SendSuTu(ctx context.Context, in *SeTuRequest, opts ...grpc.CallOption) (*SeTuReply, error) {
	out := new(SeTuReply)
	err := c.cc.Invoke(ctx, "/transmit.SetuCourier/SendSuTu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SetuCourierServer is the server API for SetuCourier service.
type SetuCourierServer interface {
	SendSuTu(context.Context, *SeTuRequest) (*SeTuReply, error)
}

// UnimplementedSetuCourierServer can be embedded to have forward compatible implementations.
type UnimplementedSetuCourierServer struct {
}

func (*UnimplementedSetuCourierServer) SendSuTu(context.Context, *SeTuRequest) (*SeTuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSuTu not implemented")
}

func RegisterSetuCourierServer(s *grpc.Server, srv SetuCourierServer) {
	s.RegisterService(&_SetuCourier_serviceDesc, srv)
}

func _SetuCourier_SendSuTu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeTuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetuCourierServer).SendSuTu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transmit.SetuCourier/SendSuTu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetuCourierServer).SendSuTu(ctx, req.(*SeTuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SetuCourier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transmit.SetuCourier",
	HandlerType: (*SetuCourierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSuTu",
			Handler:    _SetuCourier_SendSuTu_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msgforward.proto",
}
