// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query.proto

package query

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QueryRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Date                 int32    `protobuf:"varint,2,opt,name=date,proto3" json:"date,omitempty"`
	Status               int32    `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	HotelId              string   `protobuf:"bytes,4,opt,name=hotelId,proto3" json:"hotelId,omitempty"`
	Adviser              string   `protobuf:"bytes,5,opt,name=Adviser,proto3" json:"Adviser,omitempty"`
	HrId                 string   `protobuf:"bytes,6,opt,name=hrId,proto3" json:"hrId,omitempty"`
	PtId                 string   `protobuf:"bytes,7,opt,name=ptId,proto3" json:"ptId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *QueryRequest) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *QueryRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *QueryRequest) GetHotelId() string {
	if m != nil {
		return m.HotelId
	}
	return ""
}

func (m *QueryRequest) GetAdviser() string {
	if m != nil {
		return m.Adviser
	}
	return ""
}

func (m *QueryRequest) GetHrId() string {
	if m != nil {
		return m.HrId
	}
	return ""
}

func (m *QueryRequest) GetPtId() string {
	if m != nil {
		return m.PtId
	}
	return ""
}

type QueryReply struct {
	OrderOrigins         string   `protobuf:"bytes,1,opt,name=orderOrigins,proto3" json:"orderOrigins,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryReply) Reset()         { *m = QueryReply{} }
func (m *QueryReply) String() string { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()    {}
func (*QueryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
}

func (m *QueryReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReply.Unmarshal(m, b)
}
func (m *QueryReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReply.Marshal(b, m, deterministic)
}
func (m *QueryReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReply.Merge(m, src)
}
func (m *QueryReply) XXX_Size() int {
	return xxx_messageInfo_QueryReply.Size(m)
}
func (m *QueryReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReply proto.InternalMessageInfo

func (m *QueryReply) GetOrderOrigins() string {
	if m != nil {
		return m.OrderOrigins
	}
	return ""
}

type QueryPTRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	PtId                 string   `protobuf:"bytes,2,opt,name=ptId,proto3" json:"ptId,omitempty"`
	RegistrationChannel  string   `protobuf:"bytes,3,opt,name=registrationChannel,proto3" json:"registrationChannel,omitempty"`
	PtStatus             int32    `protobuf:"varint,4,opt,name=ptStatus,proto3" json:"ptStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryPTRequest) Reset()         { *m = QueryPTRequest{} }
func (m *QueryPTRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPTRequest) ProtoMessage()    {}
func (*QueryPTRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{2}
}

func (m *QueryPTRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryPTRequest.Unmarshal(m, b)
}
func (m *QueryPTRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryPTRequest.Marshal(b, m, deterministic)
}
func (m *QueryPTRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPTRequest.Merge(m, src)
}
func (m *QueryPTRequest) XXX_Size() int {
	return xxx_messageInfo_QueryPTRequest.Size(m)
}
func (m *QueryPTRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPTRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPTRequest proto.InternalMessageInfo

func (m *QueryPTRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *QueryPTRequest) GetPtId() string {
	if m != nil {
		return m.PtId
	}
	return ""
}

func (m *QueryPTRequest) GetRegistrationChannel() string {
	if m != nil {
		return m.RegistrationChannel
	}
	return ""
}

func (m *QueryPTRequest) GetPtStatus() int32 {
	if m != nil {
		return m.PtStatus
	}
	return 0
}

type PT struct {
	PtId                 string   `protobuf:"bytes,1,opt,name=ptId,proto3" json:"ptId,omitempty"`
	AdviserId            string   `protobuf:"bytes,2,opt,name=adviserId,proto3" json:"adviserId,omitempty"`
	AgentId              string   `protobuf:"bytes,3,opt,name=agentId,proto3" json:"agentId,omitempty"`
	ApplyTime            int32    `protobuf:"varint,4,opt,name=applyTime,proto3" json:"applyTime,omitempty"`
	SignTime             int32    `protobuf:"varint,5,opt,name=signTime,proto3" json:"signTime,omitempty"`
	OrderId              string   `protobuf:"bytes,6,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Id                   string   `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PT) Reset()         { *m = PT{} }
func (m *PT) String() string { return proto.CompactTextString(m) }
func (*PT) ProtoMessage()    {}
func (*PT) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{3}
}

func (m *PT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PT.Unmarshal(m, b)
}
func (m *PT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PT.Marshal(b, m, deterministic)
}
func (m *PT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PT.Merge(m, src)
}
func (m *PT) XXX_Size() int {
	return xxx_messageInfo_PT.Size(m)
}
func (m *PT) XXX_DiscardUnknown() {
	xxx_messageInfo_PT.DiscardUnknown(m)
}

var xxx_messageInfo_PT proto.InternalMessageInfo

func (m *PT) GetPtId() string {
	if m != nil {
		return m.PtId
	}
	return ""
}

func (m *PT) GetAdviserId() string {
	if m != nil {
		return m.AdviserId
	}
	return ""
}

func (m *PT) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *PT) GetApplyTime() int32 {
	if m != nil {
		return m.ApplyTime
	}
	return 0
}

func (m *PT) GetSignTime() int32 {
	if m != nil {
		return m.SignTime
	}
	return 0
}

func (m *PT) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *PT) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryPTReply struct {
	OrderPts             []*PT    `protobuf:"bytes,1,rep,name=orderPts,proto3" json:"orderPts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryPTReply) Reset()         { *m = QueryPTReply{} }
func (m *QueryPTReply) String() string { return proto.CompactTextString(m) }
func (*QueryPTReply) ProtoMessage()    {}
func (*QueryPTReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{4}
}

func (m *QueryPTReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryPTReply.Unmarshal(m, b)
}
func (m *QueryPTReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryPTReply.Marshal(b, m, deterministic)
}
func (m *QueryPTReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPTReply.Merge(m, src)
}
func (m *QueryPTReply) XXX_Size() int {
	return xxx_messageInfo_QueryPTReply.Size(m)
}
func (m *QueryPTReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPTReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPTReply proto.InternalMessageInfo

func (m *QueryPTReply) GetOrderPts() []*PT {
	if m != nil {
		return m.OrderPts
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "QueryRequest")
	proto.RegisterType((*QueryReply)(nil), "QueryReply")
	proto.RegisterType((*QueryPTRequest)(nil), "QueryPTRequest")
	proto.RegisterType((*PT)(nil), "PT")
	proto.RegisterType((*QueryPTReply)(nil), "QueryPTReply")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcb, 0x6a, 0xeb, 0x30,
	0x10, 0x8d, 0x9d, 0xc4, 0x49, 0x26, 0x8f, 0x0b, 0xba, 0x70, 0x31, 0xe1, 0x42, 0x83, 0x56, 0xa1,
	0x0b, 0x37, 0xa4, 0x5f, 0x50, 0xba, 0xca, 0x2a, 0xae, 0xeb, 0x1f, 0x70, 0xb1, 0xea, 0x08, 0x5c,
	0xdb, 0x91, 0x94, 0x42, 0xfe, 0xa1, 0x1f, 0xd3, 0x45, 0x3f, 0xb0, 0x8c, 0x24, 0xbf, 0xa0, 0x8b,
	0xee, 0x74, 0xce, 0x48, 0x73, 0x8e, 0xce, 0x0c, 0xcc, 0xcf, 0x17, 0x26, 0xae, 0x41, 0x25, 0x4a,
	0x55, 0xd2, 0x4f, 0x07, 0x16, 0x4f, 0x88, 0x23, 0x76, 0xbe, 0x30, 0xa9, 0x88, 0x0f, 0x93, 0x52,
	0xa4, 0x4c, 0x1c, 0x52, 0xdf, 0xd9, 0x38, 0xdb, 0x59, 0x54, 0x43, 0x42, 0x60, 0x94, 0x26, 0x8a,
	0xf9, 0xee, 0xc6, 0xd9, 0x8e, 0x23, 0x7d, 0x26, 0xff, 0xc0, 0x93, 0x2a, 0x51, 0x17, 0xe9, 0x0f,
	0x35, 0x6b, 0x11, 0x76, 0x39, 0x95, 0x8a, 0xe5, 0x87, 0xd4, 0x1f, 0x99, 0x2e, 0x16, 0x62, 0xe5,
	0x21, 0x7d, 0xe7, 0x92, 0x09, 0x7f, 0x6c, 0x2a, 0x16, 0x62, 0xff, 0x13, 0xca, 0x7a, 0x9a, 0xd6,
	0x67, 0xe4, 0x2a, 0x75, 0x48, 0xfd, 0x89, 0xe1, 0xf0, 0x4c, 0x77, 0x00, 0xd6, 0x71, 0x95, 0x5f,
	0x09, 0x85, 0x85, 0x36, 0x78, 0x14, 0x3c, 0xe3, 0x85, 0xb4, 0xa6, 0x7b, 0x1c, 0xfd, 0x70, 0x60,
	0xa5, 0x9f, 0x84, 0xf1, 0xaf, 0xbe, 0xa9, 0x25, 0xdd, 0x56, 0x92, 0xec, 0xe0, 0xaf, 0x60, 0x19,
	0x97, 0x4a, 0x24, 0x8a, 0x97, 0xc5, 0xe3, 0x29, 0x29, 0x0a, 0x96, 0xeb, 0x3f, 0xcf, 0xa2, 0x9f,
	0x4a, 0x64, 0x0d, 0xd3, 0x4a, 0x3d, 0x9b, 0x68, 0x46, 0x3a, 0x9a, 0x06, 0xd3, 0x2f, 0x07, 0xdc,
	0x30, 0x6e, 0x84, 0x9c, 0x8e, 0xd0, 0x7f, 0x98, 0x25, 0x26, 0x8e, 0xc6, 0x41, 0x4b, 0xa0, 0xe9,
	0x24, 0x63, 0x05, 0x3e, 0x32, 0xd2, 0x35, 0xd4, 0xef, 0xaa, 0x2a, 0xbf, 0xc6, 0xfc, 0x8d, 0x59,
	0xbd, 0x96, 0x40, 0x33, 0x92, 0x67, 0x85, 0x2e, 0x8e, 0x8d, 0x99, 0x1a, 0x77, 0x83, 0xf0, 0xfa,
	0x41, 0xac, 0xc0, 0xe5, 0x75, 0xf2, 0x2e, 0x4f, 0xe9, 0x9d, 0xdd, 0x14, 0x0c, 0x11, 0x93, 0xbf,
	0x81, 0xa9, 0xbe, 0x1a, 0x2a, 0x4c, 0x7d, 0xb8, 0x9d, 0xef, 0x87, 0x41, 0x18, 0x47, 0x0d, 0xb9,
	0xcf, 0xed, 0xa0, 0x8e, 0x48, 0x90, 0xdb, 0x1e, 0x5a, 0x06, 0xdd, 0xad, 0x5b, 0xcf, 0x83, 0x76,
	0xa4, 0x74, 0x40, 0xf6, 0xcd, 0xbc, 0x8e, 0xaf, 0xe6, 0xfe, 0x9f, 0xa0, 0x3f, 0xc0, 0xf5, 0x32,
	0xe8, 0x9a, 0xa1, 0x83, 0x17, 0x4f, 0x2f, 0xf4, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x23,
	0x40, 0x47, 0xfc, 0xdf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryOrderClient is the client API for QueryOrder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryOrderClient interface {
	// hotel/adviser/pt query order
	QueryOrder(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error)
	// hotel/adviser query order's pt
	QueryPTOfOrder(ctx context.Context, in *QueryPTRequest, opts ...grpc.CallOption) (*QueryPTReply, error)
}

type queryOrderClient struct {
	cc *grpc.ClientConn
}

func NewQueryOrderClient(cc *grpc.ClientConn) QueryOrderClient {
	return &queryOrderClient{cc}
}

func (c *queryOrderClient) QueryOrder(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error) {
	out := new(QueryReply)
	err := c.cc.Invoke(ctx, "/QueryOrder/QueryOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryOrderClient) QueryPTOfOrder(ctx context.Context, in *QueryPTRequest, opts ...grpc.CallOption) (*QueryPTReply, error) {
	out := new(QueryPTReply)
	err := c.cc.Invoke(ctx, "/QueryOrder/QueryPTOfOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryOrderServer is the server API for QueryOrder service.
type QueryOrderServer interface {
	// hotel/adviser/pt query order
	QueryOrder(context.Context, *QueryRequest) (*QueryReply, error)
	// hotel/adviser query order's pt
	QueryPTOfOrder(context.Context, *QueryPTRequest) (*QueryPTReply, error)
}

// UnimplementedQueryOrderServer can be embedded to have forward compatible implementations.
type UnimplementedQueryOrderServer struct {
}

func (*UnimplementedQueryOrderServer) QueryOrder(ctx context.Context, req *QueryRequest) (*QueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrder not implemented")
}
func (*UnimplementedQueryOrderServer) QueryPTOfOrder(ctx context.Context, req *QueryPTRequest) (*QueryPTReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPTOfOrder not implemented")
}

func RegisterQueryOrderServer(s *grpc.Server, srv QueryOrderServer) {
	s.RegisterService(&_QueryOrder_serviceDesc, srv)
}

func _QueryOrder_QueryOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryOrderServer).QueryOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QueryOrder/QueryOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryOrderServer).QueryOrder(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryOrder_QueryPTOfOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryOrderServer).QueryPTOfOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QueryOrder/QueryPTOfOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryOrderServer).QueryPTOfOrder(ctx, req.(*QueryPTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryOrder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "QueryOrder",
	HandlerType: (*QueryOrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryOrder",
			Handler:    _QueryOrder_QueryOrder_Handler,
		},
		{
			MethodName: "QueryPTOfOrder",
			Handler:    _QueryOrder_QueryPTOfOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query.proto",
}
