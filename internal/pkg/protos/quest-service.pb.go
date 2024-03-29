// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quest-service.proto

package questapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Status struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_16f61ec904d51425, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Quest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	AdditionalInfo       string   `protobuf:"bytes,4,opt,name=additionalInfo,proto3" json:"additionalInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Quest) Reset()         { *m = Quest{} }
func (m *Quest) String() string { return proto.CompactTextString(m) }
func (*Quest) ProtoMessage()    {}
func (*Quest) Descriptor() ([]byte, []int) {
	return fileDescriptor_16f61ec904d51425, []int{1}
}

func (m *Quest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quest.Unmarshal(m, b)
}
func (m *Quest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quest.Marshal(b, m, deterministic)
}
func (m *Quest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quest.Merge(m, src)
}
func (m *Quest) XXX_Size() int {
	return xxx_messageInfo_Quest.Size(m)
}
func (m *Quest) XXX_DiscardUnknown() {
	xxx_messageInfo_Quest.DiscardUnknown(m)
}

var xxx_messageInfo_Quest proto.InternalMessageInfo

func (m *Quest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Quest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Quest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Quest) GetAdditionalInfo() string {
	if m != nil {
		return m.AdditionalInfo
	}
	return ""
}

type QuestList struct {
	Quests               []*Quest `protobuf:"bytes,1,rep,name=quests,proto3" json:"quests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuestList) Reset()         { *m = QuestList{} }
func (m *QuestList) String() string { return proto.CompactTextString(m) }
func (*QuestList) ProtoMessage()    {}
func (*QuestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_16f61ec904d51425, []int{2}
}

func (m *QuestList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestList.Unmarshal(m, b)
}
func (m *QuestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestList.Marshal(b, m, deterministic)
}
func (m *QuestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestList.Merge(m, src)
}
func (m *QuestList) XXX_Size() int {
	return xxx_messageInfo_QuestList.Size(m)
}
func (m *QuestList) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestList.DiscardUnknown(m)
}

var xxx_messageInfo_QuestList proto.InternalMessageInfo

func (m *QuestList) GetQuests() []*Quest {
	if m != nil {
		return m.Quests
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "questapi.Status")
	proto.RegisterType((*Quest)(nil), "questapi.Quest")
	proto.RegisterType((*QuestList)(nil), "questapi.QuestList")
}

func init() { proto.RegisterFile("quest-service.proto", fileDescriptor_16f61ec904d51425) }

var fileDescriptor_16f61ec904d51425 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x49, 0x42, 0x03, 0xbd, 0xa0, 0x82, 0x5c, 0x84, 0xac, 0xb0, 0x44, 0x19, 0xa0, 0x0b,
	0xae, 0x14, 0xd8, 0x18, 0x81, 0x01, 0x89, 0x85, 0xf4, 0x09, 0xdc, 0xe6, 0x1a, 0x59, 0x4a, 0xeb,
	0x10, 0x3b, 0x20, 0x9e, 0x8b, 0x17, 0x44, 0x3e, 0x13, 0x81, 0xb2, 0xb1, 0xf9, 0xbe, 0xbb, 0xff,
	0x3f, 0xdf, 0x0f, 0xf3, 0xb7, 0x1e, 0x8d, 0xbd, 0x31, 0xd8, 0xbd, 0xab, 0x0d, 0x8a, 0xb6, 0xd3,
	0x56, 0xb3, 0x63, 0x82, 0xb2, 0x55, 0xe9, 0x65, 0xad, 0x75, 0xdd, 0xe0, 0x92, 0xf8, 0xba, 0xdf,
	0x2e, 0x71, 0xd7, 0xda, 0x4f, 0x3f, 0x96, 0xe7, 0x10, 0xaf, 0xac, 0xb4, 0xbd, 0x61, 0x1c, 0x8e,
	0x76, 0x68, 0x8c, 0xac, 0x91, 0x07, 0x59, 0xb0, 0x98, 0x96, 0x43, 0x99, 0x7f, 0xc0, 0xe4, 0xd5,
	0x99, 0xb1, 0x19, 0x84, 0xaa, 0xa2, 0x6e, 0x54, 0x86, 0xaa, 0x62, 0xe7, 0x30, 0xb1, 0xca, 0x36,
	0xc8, 0x43, 0x12, 0xf8, 0x82, 0x65, 0x90, 0x54, 0x68, 0x36, 0x9d, 0x6a, 0xad, 0xd2, 0x7b, 0x1e,
	0x51, 0xef, 0x2f, 0x62, 0x57, 0x30, 0x93, 0x55, 0xa5, 0xdc, 0x5b, 0x36, 0xcf, 0xfb, 0xad, 0xe6,
	0x87, 0x34, 0x34, 0xa2, 0xf9, 0x1d, 0x4c, 0x69, 0xf1, 0x8b, 0x32, 0x96, 0x5d, 0x43, 0x4c, 0x27,
	0x19, 0x1e, 0x64, 0xd1, 0x22, 0x29, 0x4e, 0xc5, 0x70, 0xa1, 0xa0, 0xa1, 0xf2, 0xa7, 0x5d, 0x7c,
	0x05, 0x70, 0x42, 0x64, 0xe5, 0x03, 0x61, 0xf7, 0x00, 0xce, 0x81, 0x98, 0x61, 0x17, 0xc2, 0xe7,
	0x21, 0x86, 0x3c, 0xc4, 0x93, 0xcb, 0x23, 0x9d, 0x8f, 0xfc, 0x9c, 0x24, 0x3f, 0x60, 0x05, 0x24,
	0x0f, 0x1d, 0x4a, 0x8b, 0x3e, 0x82, 0xf1, 0xd6, 0xf4, 0xec, 0x17, 0xf8, 0x20, 0xbd, 0xe6, 0x11,
	0x1b, 0xfc, 0x8f, 0x66, 0x1d, 0xd3, 0x77, 0x6e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xce, 0x1b,
	0xc3, 0x18, 0xcd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuestServiceClient is the client API for QuestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuestServiceClient interface {
	ListQuests(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*QuestList, error)
	CreateQuest(ctx context.Context, in *Quest, opts ...grpc.CallOption) (*Status, error)
	DeleteQuest(ctx context.Context, in *Quest, opts ...grpc.CallOption) (*Status, error)
}

type questServiceClient struct {
	cc *grpc.ClientConn
}

func NewQuestServiceClient(cc *grpc.ClientConn) QuestServiceClient {
	return &questServiceClient{cc}
}

func (c *questServiceClient) ListQuests(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*QuestList, error) {
	out := new(QuestList)
	err := c.cc.Invoke(ctx, "/questapi.QuestService/ListQuests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questServiceClient) CreateQuest(ctx context.Context, in *Quest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/questapi.QuestService/CreateQuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questServiceClient) DeleteQuest(ctx context.Context, in *Quest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/questapi.QuestService/DeleteQuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestServiceServer is the server API for QuestService service.
type QuestServiceServer interface {
	ListQuests(context.Context, *empty.Empty) (*QuestList, error)
	CreateQuest(context.Context, *Quest) (*Status, error)
	DeleteQuest(context.Context, *Quest) (*Status, error)
}

func RegisterQuestServiceServer(s *grpc.Server, srv QuestServiceServer) {
	s.RegisterService(&_QuestService_serviceDesc, srv)
}

func _QuestService_ListQuests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestServiceServer).ListQuests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/questapi.QuestService/ListQuests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestServiceServer).ListQuests(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestService_CreateQuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Quest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestServiceServer).CreateQuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/questapi.QuestService/CreateQuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestServiceServer).CreateQuest(ctx, req.(*Quest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestService_DeleteQuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Quest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestServiceServer).DeleteQuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/questapi.QuestService/DeleteQuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestServiceServer).DeleteQuest(ctx, req.(*Quest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "questapi.QuestService",
	HandlerType: (*QuestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListQuests",
			Handler:    _QuestService_ListQuests_Handler,
		},
		{
			MethodName: "CreateQuest",
			Handler:    _QuestService_CreateQuest_Handler,
		},
		{
			MethodName: "DeleteQuest",
			Handler:    _QuestService_DeleteQuest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quest-service.proto",
}
