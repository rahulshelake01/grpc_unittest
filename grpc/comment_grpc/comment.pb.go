// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comment.proto

package comment_grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Comment struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Articleid            int64    `protobuf:"varint,3,opt,name=articleid,proto3" json:"articleid,omitempty"`
	Commentcreationdate  string   `protobuf:"bytes,4,opt,name=commentcreationdate,proto3" json:"commentcreationdate,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{0}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Comment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Comment) GetArticleid() int64 {
	if m != nil {
		return m.Articleid
	}
	return 0
}

func (m *Comment) GetCommentcreationdate() string {
	if m != nil {
		return m.Commentcreationdate
	}
	return ""
}

func (m *Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CommentList struct {
	Comments             []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CommentList) Reset()         { *m = CommentList{} }
func (m *CommentList) String() string { return proto.CompactTextString(m) }
func (*CommentList) ProtoMessage()    {}
func (*CommentList) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{1}
}

func (m *CommentList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentList.Unmarshal(m, b)
}
func (m *CommentList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentList.Marshal(b, m, deterministic)
}
func (m *CommentList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentList.Merge(m, src)
}
func (m *CommentList) XXX_Size() int {
	return xxx_messageInfo_CommentList.Size(m)
}
func (m *CommentList) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentList.DiscardUnknown(m)
}

var xxx_messageInfo_CommentList proto.InternalMessageInfo

func (m *CommentList) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

type AddComment struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Articleid            int64    `protobuf:"varint,2,opt,name=articleid,proto3" json:"articleid,omitempty"`
	Commentcreationdate  string   `protobuf:"bytes,3,opt,name=commentcreationdate,proto3" json:"commentcreationdate,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddComment) Reset()         { *m = AddComment{} }
func (m *AddComment) String() string { return proto.CompactTextString(m) }
func (*AddComment) ProtoMessage()    {}
func (*AddComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{2}
}

func (m *AddComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddComment.Unmarshal(m, b)
}
func (m *AddComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddComment.Marshal(b, m, deterministic)
}
func (m *AddComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddComment.Merge(m, src)
}
func (m *AddComment) XXX_Size() int {
	return xxx_messageInfo_AddComment.Size(m)
}
func (m *AddComment) XXX_DiscardUnknown() {
	xxx_messageInfo_AddComment.DiscardUnknown(m)
}

var xxx_messageInfo_AddComment proto.InternalMessageInfo

func (m *AddComment) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AddComment) GetArticleid() int64 {
	if m != nil {
		return m.Articleid
	}
	return 0
}

func (m *AddComment) GetCommentcreationdate() string {
	if m != nil {
		return m.Commentcreationdate
	}
	return ""
}

func (m *AddComment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CommentResponse struct {
	Status               int32      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Success              bool       `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message              string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Comments             []*Comment `protobuf:"bytes,4,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CommentResponse) Reset()         { *m = CommentResponse{} }
func (m *CommentResponse) String() string { return proto.CompactTextString(m) }
func (*CommentResponse) ProtoMessage()    {}
func (*CommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{3}
}

func (m *CommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentResponse.Unmarshal(m, b)
}
func (m *CommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentResponse.Marshal(b, m, deterministic)
}
func (m *CommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentResponse.Merge(m, src)
}
func (m *CommentResponse) XXX_Size() int {
	return xxx_messageInfo_CommentResponse.Size(m)
}
func (m *CommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommentResponse proto.InternalMessageInfo

func (m *CommentResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CommentResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CommentResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CommentResponse) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

type SearchComment struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchComment) Reset()         { *m = SearchComment{} }
func (m *SearchComment) String() string { return proto.CompactTextString(m) }
func (*SearchComment) ProtoMessage()    {}
func (*SearchComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{4}
}

func (m *SearchComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchComment.Unmarshal(m, b)
}
func (m *SearchComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchComment.Marshal(b, m, deterministic)
}
func (m *SearchComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchComment.Merge(m, src)
}
func (m *SearchComment) XXX_Size() int {
	return xxx_messageInfo_SearchComment.Size(m)
}
func (m *SearchComment) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchComment.DiscardUnknown(m)
}

var xxx_messageInfo_SearchComment proto.InternalMessageInfo

func (m *SearchComment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SearchArticleComment struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchArticleComment) Reset()         { *m = SearchArticleComment{} }
func (m *SearchArticleComment) String() string { return proto.CompactTextString(m) }
func (*SearchArticleComment) ProtoMessage()    {}
func (*SearchArticleComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{5}
}

func (m *SearchArticleComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchArticleComment.Unmarshal(m, b)
}
func (m *SearchArticleComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchArticleComment.Marshal(b, m, deterministic)
}
func (m *SearchArticleComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchArticleComment.Merge(m, src)
}
func (m *SearchArticleComment) XXX_Size() int {
	return xxx_messageInfo_SearchArticleComment.Size(m)
}
func (m *SearchArticleComment) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchArticleComment.DiscardUnknown(m)
}

var xxx_messageInfo_SearchArticleComment proto.InternalMessageInfo

func (m *SearchArticleComment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{6}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Comment)(nil), "comment_grpc.Comment")
	proto.RegisterType((*CommentList)(nil), "comment_grpc.CommentList")
	proto.RegisterType((*AddComment)(nil), "comment_grpc.AddComment")
	proto.RegisterType((*CommentResponse)(nil), "comment_grpc.CommentResponse")
	proto.RegisterType((*SearchComment)(nil), "comment_grpc.SearchComment")
	proto.RegisterType((*SearchArticleComment)(nil), "comment_grpc.SearchArticleComment")
	proto.RegisterType((*Void)(nil), "comment_grpc.void")
}

func init() { proto.RegisterFile("comment.proto", fileDescriptor_749aee09ea917828) }

var fileDescriptor_749aee09ea917828 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x4e, 0xc2, 0x30,
	0x1c, 0xb6, 0xdb, 0x80, 0xf1, 0x43, 0x34, 0xa9, 0x7f, 0xd2, 0xa0, 0x46, 0xb2, 0x83, 0xe1, 0x44,
	0x14, 0x5f, 0x40, 0xf4, 0xe0, 0x45, 0x2f, 0xf3, 0xe8, 0xc1, 0xcc, 0xae, 0xc1, 0x46, 0xd6, 0x92,
	0xb5, 0xf8, 0x1c, 0xc6, 0x17, 0xf0, 0x61, 0x7c, 0x31, 0xb3, 0xae, 0x05, 0x9b, 0x60, 0xd8, 0xf1,
	0x6b, 0xbf, 0x7e, 0x7c, 0x7f, 0x18, 0xf4, 0xa9, 0x2c, 0x0a, 0x26, 0xf4, 0x78, 0x51, 0x4a, 0x2d,
	0xf1, 0xae, 0x85, 0x2f, 0xb3, 0x72, 0x41, 0x93, 0x6f, 0x04, 0x9d, 0xbb, 0xfa, 0x00, 0x0f, 0x20,
	0x16, 0x9c, 0xbe, 0x8b, 0xac, 0x60, 0x04, 0x0d, 0xd1, 0xa8, 0x9b, 0xae, 0x30, 0xde, 0x83, 0x80,
	0xe7, 0x24, 0x18, 0xa2, 0x51, 0x98, 0x06, 0x3c, 0xc7, 0xa7, 0xd0, 0xcd, 0x4a, 0xcd, 0xe9, 0x9c,
	0xf1, 0x9c, 0x84, 0xe6, 0x78, 0x7d, 0x80, 0x2f, 0xe1, 0xc0, 0xfe, 0x0a, 0x2d, 0x59, 0xa6, 0xb9,
	0x14, 0x79, 0xa6, 0x19, 0x89, 0x8c, 0xe8, 0xa6, 0x2b, 0x4c, 0xa0, 0x43, 0xa5, 0xd0, 0x4c, 0x68,
	0xd2, 0x32, 0x2c, 0x07, 0x93, 0x1b, 0xe8, 0x59, 0x83, 0x0f, 0x5c, 0x69, 0x7c, 0x05, 0xb1, 0x7d,
	0xaf, 0x08, 0x1a, 0x86, 0xa3, 0xde, 0xe4, 0x68, 0xfc, 0x37, 0xd1, 0xd8, 0x92, 0xd3, 0x15, 0x2d,
	0xf9, 0x42, 0x00, 0xd3, 0x3c, 0x6f, 0x12, 0xd3, 0x8b, 0x15, 0x34, 0x8c, 0x15, 0x36, 0x8a, 0x15,
	0xf9, 0xb1, 0x3e, 0x11, 0xec, 0x3b, 0xab, 0x4c, 0x2d, 0xa4, 0x50, 0x0c, 0x1f, 0x43, 0x5b, 0xe9,
	0x4c, 0x2f, 0x95, 0xf1, 0xd5, 0x4a, 0x2d, 0xaa, 0x54, 0xd4, 0x92, 0x52, 0xa6, 0x94, 0xf1, 0x14,
	0xa7, 0x0e, 0x56, 0x37, 0x05, 0x53, 0x2a, 0x9b, 0x39, 0x17, 0x0e, 0x7a, 0x3d, 0x45, 0xcd, 0x7a,
	0x3a, 0x87, 0xfe, 0x13, 0xcb, 0x4a, 0xfa, 0xe6, 0x9a, 0xaa, 0x47, 0x47, 0x6e, 0xf4, 0xe4, 0x02,
	0x0e, 0x6b, 0xc2, 0xb4, 0xae, 0xe4, 0x3f, 0x5e, 0x1b, 0xa2, 0x0f, 0xc9, 0xf3, 0xc9, 0x4f, 0x00,
	0xb1, 0xe5, 0x28, 0x7c, 0xbb, 0xda, 0x71, 0x5e, 0xed, 0x88, 0x7d, 0x37, 0x15, 0x7f, 0x70, 0xb6,
	0xd9, 0xa1, 0xad, 0x27, 0xd9, 0xc1, 0xf7, 0x66, 0x48, 0x4b, 0xc2, 0xc4, 0xa7, 0xaf, 0x27, 0xde,
	0x2e, 0xf4, 0xe8, 0xa2, 0x3a, 0xad, 0x13, 0xff, 0x85, 0xd7, 0xc3, 0x76, 0xb9, 0x67, 0x57, 0x8c,
	0xfd, 0xaf, 0x38, 0xd5, 0x64, 0x93, 0xaa, 0x5f, 0xde, 0x56, 0xf1, 0xd7, 0xb6, 0xf9, 0x6e, 0xaf,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x6b, 0xf6, 0x2e, 0xc8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommentsClient is the client API for Comments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentsClient interface {
	Commentlist(ctx context.Context, in *Void, opts ...grpc.CallOption) (*CommentResponse, error)
	Addcomment(ctx context.Context, in *AddComment, opts ...grpc.CallOption) (*CommentResponse, error)
	Searchcomment(ctx context.Context, in *SearchComment, opts ...grpc.CallOption) (*CommentResponse, error)
	Searcharticlecomment(ctx context.Context, in *SearchArticleComment, opts ...grpc.CallOption) (*CommentResponse, error)
}

type commentsClient struct {
	cc *grpc.ClientConn
}

func NewCommentsClient(cc *grpc.ClientConn) CommentsClient {
	return &commentsClient{cc}
}

func (c *commentsClient) Commentlist(ctx context.Context, in *Void, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/comment_grpc.Comments/Commentlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsClient) Addcomment(ctx context.Context, in *AddComment, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/comment_grpc.Comments/Addcomment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsClient) Searchcomment(ctx context.Context, in *SearchComment, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/comment_grpc.Comments/Searchcomment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsClient) Searcharticlecomment(ctx context.Context, in *SearchArticleComment, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/comment_grpc.Comments/Searcharticlecomment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentsServer is the server API for Comments service.
type CommentsServer interface {
	Commentlist(context.Context, *Void) (*CommentResponse, error)
	Addcomment(context.Context, *AddComment) (*CommentResponse, error)
	Searchcomment(context.Context, *SearchComment) (*CommentResponse, error)
	Searcharticlecomment(context.Context, *SearchArticleComment) (*CommentResponse, error)
}

func RegisterCommentsServer(s *grpc.Server, srv CommentsServer) {
	s.RegisterService(&_Comments_serviceDesc, srv)
}

func _Comments_Commentlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).Commentlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment_grpc.Comments/Commentlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).Commentlist(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comments_Addcomment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).Addcomment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment_grpc.Comments/Addcomment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).Addcomment(ctx, req.(*AddComment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comments_Searchcomment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).Searchcomment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment_grpc.Comments/Searchcomment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).Searchcomment(ctx, req.(*SearchComment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comments_Searcharticlecomment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArticleComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).Searcharticlecomment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment_grpc.Comments/Searcharticlecomment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).Searcharticlecomment(ctx, req.(*SearchArticleComment))
	}
	return interceptor(ctx, in, info, handler)
}

var _Comments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comment_grpc.Comments",
	HandlerType: (*CommentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Commentlist",
			Handler:    _Comments_Commentlist_Handler,
		},
		{
			MethodName: "Addcomment",
			Handler:    _Comments_Addcomment_Handler,
		},
		{
			MethodName: "Searchcomment",
			Handler:    _Comments_Searchcomment_Handler,
		},
		{
			MethodName: "Searcharticlecomment",
			Handler:    _Comments_Searcharticlecomment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
