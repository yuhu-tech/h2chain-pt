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

type QueryWhat struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	HotelId              string   `protobuf:"bytes,2,opt,name=hotelId,proto3" json:"hotelId,omitempty"`
	AdviseId             string   `protobuf:"bytes,3,opt,name=adviseId,proto3" json:"adviseId,omitempty"`
	Date                 int32    `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
	Duration             int32    `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	Count                int32    `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	CountMale            int32    `protobuf:"varint,7,opt,name=countMale,proto3" json:"countMale,omitempty"`
	CountFemale          int32    `protobuf:"varint,8,opt,name=countFemale,proto3" json:"countFemale,omitempty"`
	Status               int32    `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	IsFloat              int32    `protobuf:"varint,10,opt,name=isFloat,proto3" json:"isFloat,omitempty"`
	HourlySalary         int32    `protobuf:"varint,11,opt,name=hourlySalary,proto3" json:"hourlySalary,omitempty"`
	WorkContent          string   `protobuf:"bytes,12,opt,name=workContent,proto3" json:"workContent,omitempty"`
	Attention            string   `protobuf:"bytes,13,opt,name=attention,proto3" json:"attention,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryWhat) Reset()         { *m = QueryWhat{} }
func (m *QueryWhat) String() string { return proto.CompactTextString(m) }
func (*QueryWhat) ProtoMessage()    {}
func (*QueryWhat) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}

func (m *QueryWhat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryWhat.Unmarshal(m, b)
}
func (m *QueryWhat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryWhat.Marshal(b, m, deterministic)
}
func (m *QueryWhat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWhat.Merge(m, src)
}
func (m *QueryWhat) XXX_Size() int {
	return xxx_messageInfo_QueryWhat.Size(m)
}
func (m *QueryWhat) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWhat.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWhat proto.InternalMessageInfo

func (m *QueryWhat) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *QueryWhat) GetHotelId() string {
	if m != nil {
		return m.HotelId
	}
	return ""
}

func (m *QueryWhat) GetAdviseId() string {
	if m != nil {
		return m.AdviseId
	}
	return ""
}

func (m *QueryWhat) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *QueryWhat) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *QueryWhat) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *QueryWhat) GetCountMale() int32 {
	if m != nil {
		return m.CountMale
	}
	return 0
}

func (m *QueryWhat) GetCountFemale() int32 {
	if m != nil {
		return m.CountFemale
	}
	return 0
}

func (m *QueryWhat) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *QueryWhat) GetIsFloat() int32 {
	if m != nil {
		return m.IsFloat
	}
	return 0
}

func (m *QueryWhat) GetHourlySalary() int32 {
	if m != nil {
		return m.HourlySalary
	}
	return 0
}

func (m *QueryWhat) GetWorkContent() string {
	if m != nil {
		return m.WorkContent
	}
	return ""
}

func (m *QueryWhat) GetAttention() string {
	if m != nil {
		return m.Attention
	}
	return ""
}

type QueryRequest struct {
	OrderId              string     `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Date                 int32      `protobuf:"varint,2,opt,name=date,proto3" json:"date,omitempty"`
	Status               int32      `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	QueryValue           *QueryWhat `protobuf:"bytes,4,opt,name=queryValue,proto3" json:"queryValue,omitempty"`
	QueryKey             string     `protobuf:"bytes,5,opt,name=queryKey,proto3" json:"queryKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
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

func (m *QueryRequest) GetQueryValue() *QueryWhat {
	if m != nil {
		return m.QueryValue
	}
	return nil
}

func (m *QueryRequest) GetQueryKey() string {
	if m != nil {
		return m.QueryKey
	}
	return ""
}

type Order struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	HotelId              string   `protobuf:"bytes,2,opt,name=hotelId,proto3" json:"hotelId,omitempty"`
	AdviserId            string   `protobuf:"bytes,3,opt,name=adviserId,proto3" json:"adviserId,omitempty"`
	Date                 int32    `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
	Duration             int32    `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	Count                int32    `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	CountMale            int32    `protobuf:"varint,7,opt,name=countMale,proto3" json:"countMale,omitempty"`
	CountFemale          int32    `protobuf:"varint,8,opt,name=countFemale,proto3" json:"countFemale,omitempty"`
	Status               int32    `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	IsFloat              int32    `protobuf:"varint,10,opt,name=isFloat,proto3" json:"isFloat,omitempty"`
	HourlySalary         int32    `protobuf:"varint,11,opt,name=hourlySalary,proto3" json:"hourlySalary,omitempty"`
	WorkContent          string   `protobuf:"bytes,12,opt,name=workContent,proto3" json:"workContent,omitempty"`
	Attention            string   `protobuf:"bytes,13,opt,name=attention,proto3" json:"attention,omitempty"`
	Job                  string   `protobuf:"bytes,14,opt,name=job,proto3" json:"job,omitempty"`
	Mode                 int32    `protobuf:"varint,15,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{2}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Order) GetHotelId() string {
	if m != nil {
		return m.HotelId
	}
	return ""
}

func (m *Order) GetAdviserId() string {
	if m != nil {
		return m.AdviserId
	}
	return ""
}

func (m *Order) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Order) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Order) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Order) GetCountMale() int32 {
	if m != nil {
		return m.CountMale
	}
	return 0
}

func (m *Order) GetCountFemale() int32 {
	if m != nil {
		return m.CountFemale
	}
	return 0
}

func (m *Order) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Order) GetIsFloat() int32 {
	if m != nil {
		return m.IsFloat
	}
	return 0
}

func (m *Order) GetHourlySalary() int32 {
	if m != nil {
		return m.HourlySalary
	}
	return 0
}

func (m *Order) GetWorkContent() string {
	if m != nil {
		return m.WorkContent
	}
	return ""
}

func (m *Order) GetAttention() string {
	if m != nil {
		return m.Attention
	}
	return ""
}

func (m *Order) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *Order) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

type OrderHotelModify struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Revision             int32    `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
	TimeStamp            int32    `protobuf:"varint,3,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	Count                int32    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	CountMale            int32    `protobuf:"varint,5,opt,name=countMale,proto3" json:"countMale,omitempty"`
	CountYet             int32    `protobuf:"varint,6,opt,name=countYet,proto3" json:"countYet,omitempty"`
	CountMaleYet         int32    `protobuf:"varint,7,opt,name=countMale_yet,json=countMaleYet,proto3" json:"countMale_yet,omitempty"`
	Date                 int32    `protobuf:"varint,8,opt,name=date,proto3" json:"date,omitempty"`
	Duration             int32    `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
	Mode                 int32    `protobuf:"varint,10,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderHotelModify) Reset()         { *m = OrderHotelModify{} }
func (m *OrderHotelModify) String() string { return proto.CompactTextString(m) }
func (*OrderHotelModify) ProtoMessage()    {}
func (*OrderHotelModify) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{3}
}

func (m *OrderHotelModify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderHotelModify.Unmarshal(m, b)
}
func (m *OrderHotelModify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderHotelModify.Marshal(b, m, deterministic)
}
func (m *OrderHotelModify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderHotelModify.Merge(m, src)
}
func (m *OrderHotelModify) XXX_Size() int {
	return xxx_messageInfo_OrderHotelModify.Size(m)
}
func (m *OrderHotelModify) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderHotelModify.DiscardUnknown(m)
}

var xxx_messageInfo_OrderHotelModify proto.InternalMessageInfo

func (m *OrderHotelModify) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrderHotelModify) GetRevision() int32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *OrderHotelModify) GetTimeStamp() int32 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *OrderHotelModify) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *OrderHotelModify) GetCountMale() int32 {
	if m != nil {
		return m.CountMale
	}
	return 0
}

func (m *OrderHotelModify) GetCountYet() int32 {
	if m != nil {
		return m.CountYet
	}
	return 0
}

func (m *OrderHotelModify) GetCountMaleYet() int32 {
	if m != nil {
		return m.CountMaleYet
	}
	return 0
}

func (m *OrderHotelModify) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *OrderHotelModify) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *OrderHotelModify) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

type OrderAdviserModify struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Revision             int32    `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
	TimeStamp            int32    `protobuf:"varint,3,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	IsFloat              int32    `protobuf:"varint,4,opt,name=isFloat,proto3" json:"isFloat,omitempty"`
	Count                int32    `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	CountMale            int32    `protobuf:"varint,6,opt,name=countMale,proto3" json:"countMale,omitempty"`
	HourlySalary         int32    `protobuf:"varint,7,opt,name=hourlySalary,proto3" json:"hourlySalary,omitempty"`
	WorkContent          string   `protobuf:"bytes,8,opt,name=workContent,proto3" json:"workContent,omitempty"`
	Attention            string   `protobuf:"bytes,9,opt,name=attention,proto3" json:"attention,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderAdviserModify) Reset()         { *m = OrderAdviserModify{} }
func (m *OrderAdviserModify) String() string { return proto.CompactTextString(m) }
func (*OrderAdviserModify) ProtoMessage()    {}
func (*OrderAdviserModify) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{4}
}

func (m *OrderAdviserModify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderAdviserModify.Unmarshal(m, b)
}
func (m *OrderAdviserModify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderAdviserModify.Marshal(b, m, deterministic)
}
func (m *OrderAdviserModify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderAdviserModify.Merge(m, src)
}
func (m *OrderAdviserModify) XXX_Size() int {
	return xxx_messageInfo_OrderAdviserModify.Size(m)
}
func (m *OrderAdviserModify) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderAdviserModify.DiscardUnknown(m)
}

var xxx_messageInfo_OrderAdviserModify proto.InternalMessageInfo

func (m *OrderAdviserModify) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrderAdviserModify) GetRevision() int32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *OrderAdviserModify) GetTimeStamp() int32 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *OrderAdviserModify) GetIsFloat() int32 {
	if m != nil {
		return m.IsFloat
	}
	return 0
}

func (m *OrderAdviserModify) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *OrderAdviserModify) GetCountMale() int32 {
	if m != nil {
		return m.CountMale
	}
	return 0
}

func (m *OrderAdviserModify) GetHourlySalary() int32 {
	if m != nil {
		return m.HourlySalary
	}
	return 0
}

func (m *OrderAdviserModify) GetWorkContent() string {
	if m != nil {
		return m.WorkContent
	}
	return ""
}

func (m *OrderAdviserModify) GetAttention() string {
	if m != nil {
		return m.Attention
	}
	return ""
}

type OrderCandidate struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AdviserId            string   `protobuf:"bytes,2,opt,name=adviserId,proto3" json:"adviserId,omitempty"`
	AgentId              string   `protobuf:"bytes,3,opt,name=agentId,proto3" json:"agentId,omitempty"`
	PtId                 string   `protobuf:"bytes,4,opt,name=ptId,proto3" json:"ptId,omitempty"`
	ApplyTime            int32    `protobuf:"varint,5,opt,name=applyTime,proto3" json:"applyTime,omitempty"`
	SignInTime           int32    `protobuf:"varint,6,opt,name=signInTime,proto3" json:"signInTime,omitempty"`
	PtStatus             int32    `protobuf:"varint,7,opt,name=ptStatus,proto3" json:"ptStatus,omitempty"`
	PtPerformance        int32    `protobuf:"varint,8,opt,name=ptPerformance,proto3" json:"ptPerformance,omitempty"`
	ObjectReason         int32    `protobuf:"varint,9,opt,name=objectReason,proto3" json:"objectReason,omitempty"`
	RegistrationChannel  string   `protobuf:"bytes,10,opt,name=registrationChannel,proto3" json:"registrationChannel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderCandidate) Reset()         { *m = OrderCandidate{} }
func (m *OrderCandidate) String() string { return proto.CompactTextString(m) }
func (*OrderCandidate) ProtoMessage()    {}
func (*OrderCandidate) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{5}
}

func (m *OrderCandidate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderCandidate.Unmarshal(m, b)
}
func (m *OrderCandidate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderCandidate.Marshal(b, m, deterministic)
}
func (m *OrderCandidate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderCandidate.Merge(m, src)
}
func (m *OrderCandidate) XXX_Size() int {
	return xxx_messageInfo_OrderCandidate.Size(m)
}
func (m *OrderCandidate) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderCandidate.DiscardUnknown(m)
}

var xxx_messageInfo_OrderCandidate proto.InternalMessageInfo

func (m *OrderCandidate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrderCandidate) GetAdviserId() string {
	if m != nil {
		return m.AdviserId
	}
	return ""
}

func (m *OrderCandidate) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *OrderCandidate) GetPtId() string {
	if m != nil {
		return m.PtId
	}
	return ""
}

func (m *OrderCandidate) GetApplyTime() int32 {
	if m != nil {
		return m.ApplyTime
	}
	return 0
}

func (m *OrderCandidate) GetSignInTime() int32 {
	if m != nil {
		return m.SignInTime
	}
	return 0
}

func (m *OrderCandidate) GetPtStatus() int32 {
	if m != nil {
		return m.PtStatus
	}
	return 0
}

func (m *OrderCandidate) GetPtPerformance() int32 {
	if m != nil {
		return m.PtPerformance
	}
	return 0
}

func (m *OrderCandidate) GetObjectReason() int32 {
	if m != nil {
		return m.ObjectReason
	}
	return 0
}

func (m *OrderCandidate) GetRegistrationChannel() string {
	if m != nil {
		return m.RegistrationChannel
	}
	return ""
}

type QueryReply struct {
	Orders               []*Order              `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	OrderHotelModifies   []*OrderHotelModify   `protobuf:"bytes,2,rep,name=orderHotelModifies,proto3" json:"orderHotelModifies,omitempty"`
	OrderAdviserModifies []*OrderAdviserModify `protobuf:"bytes,3,rep,name=orderAdviserModifies,proto3" json:"orderAdviserModifies,omitempty"`
	OrderCandidates      []*OrderCandidate     `protobuf:"bytes,4,rep,name=orderCandidates,proto3" json:"orderCandidates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *QueryReply) Reset()         { *m = QueryReply{} }
func (m *QueryReply) String() string { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()    {}
func (*QueryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{6}
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

func (m *QueryReply) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func (m *QueryReply) GetOrderHotelModifies() []*OrderHotelModify {
	if m != nil {
		return m.OrderHotelModifies
	}
	return nil
}

func (m *QueryReply) GetOrderAdviserModifies() []*OrderAdviserModify {
	if m != nil {
		return m.OrderAdviserModifies
	}
	return nil
}

func (m *QueryReply) GetOrderCandidates() []*OrderCandidate {
	if m != nil {
		return m.OrderCandidates
	}
	return nil
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
	return fileDescriptor_5c6ac9b241082464, []int{7}
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
	return fileDescriptor_5c6ac9b241082464, []int{8}
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
	return fileDescriptor_5c6ac9b241082464, []int{9}
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
	proto.RegisterType((*QueryWhat)(nil), "QueryWhat")
	proto.RegisterType((*QueryRequest)(nil), "QueryRequest")
	proto.RegisterType((*Order)(nil), "Order")
	proto.RegisterType((*OrderHotelModify)(nil), "OrderHotelModify")
	proto.RegisterType((*OrderAdviserModify)(nil), "OrderAdviserModify")
	proto.RegisterType((*OrderCandidate)(nil), "OrderCandidate")
	proto.RegisterType((*QueryReply)(nil), "QueryReply")
	proto.RegisterType((*QueryPTRequest)(nil), "QueryPTRequest")
	proto.RegisterType((*PT)(nil), "PT")
	proto.RegisterType((*QueryPTReply)(nil), "QueryPTReply")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	// 829 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x96, 0xcb, 0x6e, 0xfb, 0x44,
	0x14, 0xc6, 0xeb, 0x4b, 0x2e, 0x3e, 0xb9, 0x95, 0x69, 0x85, 0x46, 0x11, 0x2a, 0x91, 0x61, 0x51,
	0x75, 0x61, 0x50, 0x58, 0xb1, 0xac, 0x2a, 0x15, 0x22, 0x54, 0x35, 0xb8, 0x11, 0x88, 0x15, 0x9a,
	0xd6, 0xd3, 0xc6, 0xc5, 0xf1, 0xb8, 0xf6, 0xa4, 0xc8, 0xef, 0xc0, 0x82, 0x0d, 0x2b, 0x78, 0x0c,
	0x9e, 0x80, 0x57, 0x62, 0xc1, 0x16, 0xcd, 0xb1, 0x3d, 0xb1, 0x4d, 0x43, 0x51, 0xc5, 0xf2, 0xbf,
	0x9b, 0xef, 0x3b, 0x13, 0xdb, 0xf3, 0x9d, 0xdf, 0xb1, 0x03, 0x83, 0xa7, 0x2d, 0x4f, 0x73, 0x2f,
	0x49, 0x85, 0x14, 0xee, 0x5f, 0x26, 0x38, 0x5f, 0x2b, 0xfd, 0xed, 0x9a, 0x49, 0x42, 0xa1, 0x27,
	0xd2, 0x80, 0xa7, 0x8b, 0x80, 0x1a, 0x33, 0xe3, 0xd4, 0xf1, 0x2b, 0xa9, 0x2a, 0x6b, 0x21, 0x79,
	0xb4, 0x08, 0xa8, 0x59, 0x54, 0x4a, 0x49, 0xa6, 0xd0, 0x67, 0xc1, 0x73, 0x98, 0xf1, 0x45, 0x40,
	0x2d, 0x2c, 0x69, 0x4d, 0x08, 0xd8, 0x01, 0x93, 0x9c, 0xda, 0x33, 0xe3, 0xb4, 0xe3, 0xe3, 0x5a,
	0xed, 0x0f, 0xb6, 0x29, 0x93, 0xa1, 0x88, 0x69, 0x07, 0x7d, 0xad, 0xc9, 0x31, 0x74, 0xee, 0xc4,
	0x36, 0x96, 0xb4, 0x8b, 0x85, 0x42, 0x90, 0x0f, 0xc0, 0xc1, 0xc5, 0x15, 0x8b, 0x38, 0xed, 0x61,
	0x65, 0x67, 0x90, 0x19, 0x0c, 0x50, 0x5c, 0xf2, 0x8d, 0xaa, 0xf7, 0xb1, 0x5e, 0xb7, 0xc8, 0xfb,
	0xd0, 0xcd, 0x24, 0x93, 0xdb, 0x8c, 0x3a, 0x58, 0x2c, 0x95, 0x3a, 0x53, 0x98, 0x5d, 0x46, 0x82,
	0x49, 0x0a, 0x58, 0xa8, 0x24, 0x71, 0x61, 0xb8, 0x16, 0xdb, 0x34, 0xca, 0x6f, 0x58, 0xc4, 0xd2,
	0x9c, 0x0e, 0xb0, 0xdc, 0xf0, 0xd4, 0x7d, 0x7f, 0x14, 0xe9, 0x0f, 0x17, 0x22, 0x96, 0x3c, 0x96,
	0x74, 0x88, 0x47, 0xaf, 0x5b, 0xea, 0xb9, 0x99, 0x54, 0x2b, 0x75, 0xd4, 0x11, 0xd6, 0x77, 0x86,
	0xfb, 0x9b, 0x01, 0x43, 0x4c, 0xde, 0xe7, 0x4f, 0x5b, 0x9e, 0xfd, 0x5b, 0xf8, 0x55, 0x8c, 0x66,
	0x2d, 0xc6, 0xdd, 0xa1, 0xac, 0xc6, 0xa1, 0xce, 0x00, 0xb0, 0xbf, 0xdf, 0xb0, 0x68, 0x5b, 0x04,
	0x3f, 0x98, 0x83, 0xa7, 0x5b, 0xec, 0xd7, 0xaa, 0xaa, 0x15, 0xa8, 0xbe, 0xe2, 0x39, 0xb6, 0xc2,
	0xf1, 0xb5, 0x76, 0x7f, 0xb5, 0xa0, 0x73, 0xad, 0xee, 0xff, 0x26, 0x28, 0xd4, 0xd1, 0x11, 0x82,
	0x54, 0x53, 0xb1, 0x33, 0xde, 0x61, 0xb1, 0x1f, 0x0b, 0x72, 0x08, 0xd6, 0xa3, 0xb8, 0xa5, 0x63,
	0xf4, 0xd5, 0x52, 0xa5, 0xb5, 0x11, 0x01, 0xa7, 0x93, 0x22, 0x2d, 0xb5, 0x76, 0x7f, 0x31, 0xe1,
	0x10, 0xbb, 0xf3, 0xa5, 0x0a, 0xfc, 0x4a, 0x04, 0xe1, 0x7d, 0x4e, 0xc6, 0x60, 0x86, 0x55, 0x8f,
	0xcc, 0x10, 0x27, 0x33, 0xe5, 0xcf, 0x61, 0xa6, 0xee, 0x53, 0xa0, 0xa3, 0xb5, 0x7a, 0x08, 0x19,
	0x6e, 0xf8, 0x8d, 0x64, 0x9b, 0xa4, 0x24, 0x68, 0x67, 0xec, 0x02, 0xb7, 0xf7, 0x06, 0xde, 0x69,
	0x07, 0x3e, 0x85, 0x3e, 0x8a, 0xef, 0x78, 0xd5, 0x27, 0xad, 0xc9, 0x47, 0x30, 0xd2, 0x1b, 0xbf,
	0xcf, 0xb9, 0x2c, 0xdb, 0x35, 0xd4, 0xa6, 0xda, 0x54, 0x51, 0xd1, 0xdf, 0x43, 0x85, 0xd3, 0xa2,
	0xa2, 0xca, 0x05, 0x6a, 0xb9, 0xfc, 0x6c, 0x02, 0xc1, 0x5c, 0xce, 0x0b, 0xd8, 0xfe, 0xf7, 0x64,
	0x6a, 0x70, 0xd8, 0x4d, 0x38, 0x74, 0x66, 0x9d, 0xbd, 0x99, 0x75, 0xdb, 0x99, 0xb5, 0x81, 0xea,
	0xbd, 0x0e, 0x54, 0xff, 0x15, 0xa0, 0x9c, 0xf6, 0x7b, 0xe6, 0x0f, 0x13, 0xc6, 0x18, 0xc9, 0x05,
	0x8b, 0x83, 0x10, 0x53, 0x6d, 0xc7, 0xd1, 0x98, 0x56, 0xb3, 0x3d, 0xad, 0x14, 0x7a, 0xec, 0x81,
	0xc7, 0x52, 0x4f, 0x72, 0x25, 0x55, 0x07, 0x12, 0x65, 0xdb, 0x68, 0xe3, 0x1a, 0xaf, 0x95, 0x24,
	0x51, 0xbe, 0x0a, 0x37, 0x1a, 0x12, 0x6d, 0x90, 0x13, 0x80, 0x2c, 0x7c, 0x88, 0x17, 0x31, 0x96,
	0x8b, 0x3c, 0x6a, 0x8e, 0x6a, 0x4c, 0x22, 0x6f, 0x8a, 0xa9, 0x2c, 0xc2, 0xd0, 0x9a, 0x7c, 0x0c,
	0xa3, 0x44, 0x2e, 0x79, 0x7a, 0x2f, 0xd2, 0x0d, 0x8b, 0xef, 0x2a, 0x50, 0x9a, 0xa6, 0x8a, 0x54,
	0xdc, 0x3e, 0xf2, 0x3b, 0xe9, 0x73, 0x96, 0x69, 0x6a, 0x1a, 0x1e, 0xf9, 0x14, 0x8e, 0x52, 0xfe,
	0x10, 0x66, 0xb2, 0x20, 0xe9, 0x62, 0xcd, 0xe2, 0x98, 0x47, 0x08, 0x92, 0xe3, 0xbf, 0x54, 0x72,
	0xff, 0x34, 0x00, 0xca, 0x97, 0x75, 0x12, 0xe5, 0xe4, 0x04, 0xba, 0xf8, 0x0e, 0xcc, 0xa8, 0x31,
	0xb3, 0x4e, 0x07, 0xf3, 0xae, 0x87, 0x09, 0xfb, 0xa5, 0x4b, 0xce, 0x81, 0x88, 0xe6, 0x74, 0x86,
	0x3c, 0xa3, 0x26, 0xee, 0x7d, 0xcf, 0x6b, 0x0f, 0xae, 0xff, 0xc2, 0x66, 0xf2, 0x05, 0x1c, 0x8b,
	0x36, 0xc8, 0xea, 0x22, 0x16, 0x5e, 0xe4, 0xc8, 0xfb, 0x27, 0xe5, 0xfe, 0x8b, 0x3f, 0x20, 0x9f,
	0xc3, 0x44, 0x34, 0xda, 0x9f, 0x51, 0x1b, 0xaf, 0x31, 0xf1, 0x9a, 0x58, 0xf8, 0xed, 0x7d, 0xee,
	0x4f, 0x06, 0x8c, 0xf1, 0xd4, 0xcb, 0xd5, 0x7f, 0xfa, 0x48, 0x21, 0x0c, 0x66, 0x0d, 0x86, 0x3d,
	0x41, 0x5b, 0x7b, 0x83, 0x6e, 0x00, 0x60, 0x37, 0x01, 0x70, 0x7f, 0x37, 0xc0, 0x5c, 0xae, 0xf4,
	0x8d, 0x8c, 0x16, 0x75, 0x6f, 0x21, 0xb8, 0x41, 0xab, 0xdd, 0xa6, 0x75, 0x0a, 0x7d, 0xc5, 0x66,
	0x0d, 0x65, 0xad, 0xeb, 0x41, 0x74, 0x9b, 0x41, 0x14, 0xd3, 0xd5, 0xab, 0xa6, 0xcb, 0xfd, 0xa4,
	0xfc, 0xce, 0xab, 0x10, 0x15, 0x3c, 0x1f, 0x42, 0x1f, 0xb7, 0x2e, 0x65, 0x85, 0x8f, 0xe5, 0x2d,
	0x57, 0xbe, 0x36, 0xe7, 0x51, 0xc9, 0x5a, 0xf1, 0xf9, 0x3d, 0x6b, 0xa8, 0x91, 0x57, 0xff, 0xcf,
	0x30, 0x1d, 0x78, 0x3b, 0x2a, 0xdd, 0x03, 0x32, 0xd7, 0xfd, 0xba, 0xbe, 0x2f, 0xf6, 0x4f, 0xbc,
	0x66, 0x03, 0xa7, 0x23, 0xaf, 0xfe, 0x30, 0xee, 0xc1, 0x6d, 0x17, 0xff, 0x08, 0x7e, 0xf6, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x5d, 0x98, 0xc4, 0x17, 0x0a, 0x00, 0x00,
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
