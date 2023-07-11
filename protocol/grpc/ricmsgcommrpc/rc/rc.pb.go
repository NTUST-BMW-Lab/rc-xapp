// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rc.proto

package rc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RICControlCellTypeEnum int32

const (
	RICControlCellTypeEnum_RIC_CONTROL_CELL_UNKWON RICControlCellTypeEnum = 0
	RICControlCellTypeEnum_RIC_CONTROL_NR_CELL     RICControlCellTypeEnum = 1
	RICControlCellTypeEnum_RIC_CONTROL_EUTRAN_CELL RICControlCellTypeEnum = 2
)

var RICControlCellTypeEnum_name = map[int32]string{
	0: "RIC_CONTROL_CELL_UNKWON",
	1: "RIC_CONTROL_NR_CELL",
	2: "RIC_CONTROL_EUTRAN_CELL",
}
var RICControlCellTypeEnum_value = map[string]int32{
	"RIC_CONTROL_CELL_UNKWON": 0,
	"RIC_CONTROL_NR_CELL":     1,
	"RIC_CONTROL_EUTRAN_CELL": 2,
}

func (x RICControlCellTypeEnum) String() string {
	return proto.EnumName(RICControlCellTypeEnum_name, int32(x))
}
func (RICControlCellTypeEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{0}
}

type RICControlAckEnum int32

const (
	RICControlAckEnum_RIC_CONTROL_ACK_UNKWON RICControlAckEnum = 0
	RICControlAckEnum_RIC_CONTROL_NO_ACK     RICControlAckEnum = 1
	RICControlAckEnum_RIC_CONTROL_ACK        RICControlAckEnum = 2
	RICControlAckEnum_RIC_CONTROL_NACK       RICControlAckEnum = 3
)

var RICControlAckEnum_name = map[int32]string{
	0: "RIC_CONTROL_ACK_UNKWON",
	1: "RIC_CONTROL_NO_ACK",
	2: "RIC_CONTROL_ACK",
	3: "RIC_CONTROL_NACK",
}
var RICControlAckEnum_value = map[string]int32{
	"RIC_CONTROL_ACK_UNKWON": 0,
	"RIC_CONTROL_NO_ACK":     1,
	"RIC_CONTROL_ACK":        2,
	"RIC_CONTROL_NACK":       3,
}

func (x RICControlAckEnum) String() string {
	return proto.EnumName(RICControlAckEnum_name, int32(x))
}
func (RICControlAckEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{1}
}

type RICE2APHeader struct {
	RanFuncId            int64    `protobuf:"varint,1,opt,name=RanFuncId" json:"RanFuncId,omitempty"`
	RICRequestorID       int64    `protobuf:"varint,2,opt,name=RICRequestorID" json:"RICRequestorID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RICE2APHeader) Reset()         { *m = RICE2APHeader{} }
func (m *RICE2APHeader) String() string { return proto.CompactTextString(m) }
func (*RICE2APHeader) ProtoMessage()    {}
func (*RICE2APHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{0}
}
func (m *RICE2APHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICE2APHeader.Unmarshal(m, b)
}
func (m *RICE2APHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICE2APHeader.Marshal(b, m, deterministic)
}
func (dst *RICE2APHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICE2APHeader.Merge(dst, src)
}
func (m *RICE2APHeader) XXX_Size() int {
	return xxx_messageInfo_RICE2APHeader.Size(m)
}
func (m *RICE2APHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RICE2APHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RICE2APHeader proto.InternalMessageInfo

func (m *RICE2APHeader) GetRanFuncId() int64 {
	if m != nil {
		return m.RanFuncId
	}
	return 0
}

func (m *RICE2APHeader) GetRICRequestorID() int64 {
	if m != nil {
		return m.RICRequestorID
	}
	return 0
}

type RICControlHeader struct {
	ControlStyle         int64    `protobuf:"varint,1,opt,name=ControlStyle" json:"ControlStyle,omitempty"`
	ControlActionId      int64    `protobuf:"varint,2,opt,name=ControlActionId" json:"ControlActionId,omitempty"`
	UEID                 *UeId    `protobuf:"bytes,3,opt,name=UEID" json:"UEID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RICControlHeader) Reset()         { *m = RICControlHeader{} }
func (m *RICControlHeader) String() string { return proto.CompactTextString(m) }
func (*RICControlHeader) ProtoMessage()    {}
func (*RICControlHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{1}
}
func (m *RICControlHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICControlHeader.Unmarshal(m, b)
}
func (m *RICControlHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICControlHeader.Marshal(b, m, deterministic)
}
func (dst *RICControlHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICControlHeader.Merge(dst, src)
}
func (m *RICControlHeader) XXX_Size() int {
	return xxx_messageInfo_RICControlHeader.Size(m)
}
func (m *RICControlHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RICControlHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RICControlHeader proto.InternalMessageInfo

func (m *RICControlHeader) GetControlStyle() int64 {
	if m != nil {
		return m.ControlStyle
	}
	return 0
}

func (m *RICControlHeader) GetControlActionId() int64 {
	if m != nil {
		return m.ControlActionId
	}
	return 0
}

func (m *RICControlHeader) GetUEID() *UeId {
	if m != nil {
		return m.UEID
	}
	return nil
}

type UeId struct {
	GnbUEID              *GNBUEID `protobuf:"bytes,1,opt,name=GnbUEID" json:"GnbUEID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UeId) Reset()         { *m = UeId{} }
func (m *UeId) String() string { return proto.CompactTextString(m) }
func (*UeId) ProtoMessage()    {}
func (*UeId) Descriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{2}
}
func (m *UeId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UeId.Unmarshal(m, b)
}
func (m *UeId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UeId.Marshal(b, m, deterministic)
}
func (dst *UeId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UeId.Merge(dst, src)
}
func (m *UeId) XXX_Size() int {
	return xxx_messageInfo_UeId.Size(m)
}
func (m *UeId) XXX_DiscardUnknown() {
	xxx_messageInfo_UeId.DiscardUnknown(m)
}

var xxx_messageInfo_UeId proto.InternalMessageInfo

func (m *UeId) GetGnbUEID() *GNBUEID {
	if m != nil {
		return m.GnbUEID
	}
	return nil
}

type GNBUEID struct {
	AmfUENGAPID          int64    `protobuf:"varint,1,opt,name=amfUENGAPID" json:"amfUENGAPID,omitempty"`
	Guami                *Guami   `protobuf:"bytes,2,opt,name=guami" json:"guami,omitempty"`
	GNBCUUEF1APID        []int64  `protobuf:"varint,3,rep,packed,name=gNBCUUEF1APID" json:"gNBCUUEF1APID,omitempty"`
	GNBCUCPUEE1APID      []int64  `protobuf:"varint,4,rep,packed,name=gNBCUCPUEE1APID" json:"gNBCUCPUEE1APID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GNBUEID) Reset()         { *m = GNBUEID{} }
func (m *GNBUEID) String() string { return proto.CompactTextString(m) }
func (*GNBUEID) ProtoMessage()    {}
func (*GNBUEID) Descriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{3}
}
func (m *GNBUEID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GNBUEID.Unmarshal(m, b)
}
func (m *GNBUEID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GNBUEID.Marshal(b, m, deterministic)
}
func (dst *GNBUEID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GNBUEID.Merge(dst, src)
}
func (m *GNBUEID) XXX_Size() int {
	return xxx_messageInfo_GNBUEID.Size(m)
}
func (m *GNBUEID) XXX_DiscardUnknown() {
	xxx_messageInfo_GNBUEID.DiscardUnknown(m)
}

var xxx_messageInfo_GNBUEID proto.InternalMessageInfo

func (m *GNBUEID) GetAmfUENGAPID() int64 {
	if m != nil {
		return m.AmfUENGAPID
	}
	return 0
}

func (m *GNBUEID) GetGuami() *Guami {
	if m != nil {
		return m.Guami
	}
	return nil
}

func (m *GNBUEID) GetGNBCUUEF1APID() []int64 {
	if m != nil {
		return m.GNBCUUEF1APID
	}
	return nil
}

func (m *GNBUEID) GetGNBCUCPUEE1APID() []int64 {
	if m != nil {
		return m.GNBCUCPUEE1APID
	}
	return nil
}

type Guami struct {
	PLMNIdentity         string   `protobuf:"bytes,1,opt,name=pLMNIdentity" json:"pLMNIdentity,omitempty"`
	AMFRegionID          string   `protobuf:"bytes,2,opt,name=aMFRegionID" json:"aMFRegionID,omitempty"`
	AMFSetID             string   `protobuf:"bytes,3,opt,name=aMFSetID" json:"aMFSetID,omitempty"`
	AMFPointer           string   `protobuf:"bytes,4,opt,name=aMFPointer" json:"aMFPointer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Guami) Reset()         { *m = Guami{} }
func (m *Guami) String() string { return proto.CompactTextString(m) }
func (*Guami) ProtoMessage()    {}
func (*Guami) Descriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{4}
}
func (m *Guami) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Guami.Unmarshal(m, b)
}
func (m *Guami) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Guami.Marshal(b, m, deterministic)
}
func (dst *Guami) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Guami.Merge(dst, src)
}
func (m *Guami) XXX_Size() int {
	return xxx_messageInfo_Guami.Size(m)
}
func (m *Guami) XXX_DiscardUnknown() {
	xxx_messageInfo_Guami.DiscardUnknown(m)
}

var xxx_messageInfo_Guami proto.InternalMessageInfo

func (m *Guami) GetPLMNIdentity() string {
	if m != nil {
		return m.PLMNIdentity
	}
	return ""
}

func (m *Guami) GetAMFRegionID() string {
	if m != nil {
		return m.AMFRegionID
	}
	return ""
}

func (m *Guami) GetAMFSetID() string {
	if m != nil {
		return m.AMFSetID
	}
	return ""
}

func (m *Guami) GetAMFPointer() string {
	if m != nil {
		return m.AMFPointer
	}
	return ""
}

type RICControlMessage struct {
	RICControlCellTypeVal RICControlCellTypeEnum `protobuf:"varint,1,opt,name=RICControlCellTypeVal,enum=rc.RICControlCellTypeEnum" json:"RICControlCellTypeVal,omitempty"`
	TargetCellID          string                 `protobuf:"bytes,2,opt,name=TargetCellID" json:"TargetCellID,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *RICControlMessage) Reset()         { *m = RICControlMessage{} }
func (m *RICControlMessage) String() string { return proto.CompactTextString(m) }
func (*RICControlMessage) ProtoMessage()    {}
func (*RICControlMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{5}
}
func (m *RICControlMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICControlMessage.Unmarshal(m, b)
}
func (m *RICControlMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICControlMessage.Marshal(b, m, deterministic)
}
func (dst *RICControlMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICControlMessage.Merge(dst, src)
}
func (m *RICControlMessage) XXX_Size() int {
	return xxx_messageInfo_RICControlMessage.Size(m)
}
func (m *RICControlMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RICControlMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RICControlMessage proto.InternalMessageInfo

func (m *RICControlMessage) GetRICControlCellTypeVal() RICControlCellTypeEnum {
	if m != nil {
		return m.RICControlCellTypeVal
	}
	return RICControlCellTypeEnum_RIC_CONTROL_CELL_UNKWON
}

func (m *RICControlMessage) GetTargetCellID() string {
	if m != nil {
		return m.TargetCellID
	}
	return ""
}

// RicControl GRPC Req
type RicControlGrpcReq struct {
	E2NodeID              string             `protobuf:"bytes,1,opt,name=e2NodeID" json:"e2NodeID,omitempty"`
	PlmnID                string             `protobuf:"bytes,2,opt,name=plmnID" json:"plmnID,omitempty"`
	RanName               string             `protobuf:"bytes,3,opt,name=ranName" json:"ranName,omitempty"`
	RICE2APHeaderData     *RICE2APHeader     `protobuf:"bytes,4,opt,name=RICE2APHeaderData" json:"RICE2APHeaderData,omitempty"`
	RICControlHeaderData  *RICControlHeader  `protobuf:"bytes,5,opt,name=RICControlHeaderData" json:"RICControlHeaderData,omitempty"`
	RICControlMessageData *RICControlMessage `protobuf:"bytes,6,opt,name=RICControlMessageData" json:"RICControlMessageData,omitempty"`
	RICControlAckReqVal   RICControlAckEnum  `protobuf:"varint,7,opt,name=RICControlAckReqVal,enum=rc.RICControlAckEnum" json:"RICControlAckReqVal,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}           `json:"-"`
	XXX_unrecognized      []byte             `json:"-"`
	XXX_sizecache         int32              `json:"-"`
}

func (m *RicControlGrpcReq) Reset()         { *m = RicControlGrpcReq{} }
func (m *RicControlGrpcReq) String() string { return proto.CompactTextString(m) }
func (*RicControlGrpcReq) ProtoMessage()    {}
func (*RicControlGrpcReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{6}
}
func (m *RicControlGrpcReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RicControlGrpcReq.Unmarshal(m, b)
}
func (m *RicControlGrpcReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RicControlGrpcReq.Marshal(b, m, deterministic)
}
func (dst *RicControlGrpcReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RicControlGrpcReq.Merge(dst, src)
}
func (m *RicControlGrpcReq) XXX_Size() int {
	return xxx_messageInfo_RicControlGrpcReq.Size(m)
}
func (m *RicControlGrpcReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RicControlGrpcReq.DiscardUnknown(m)
}

var xxx_messageInfo_RicControlGrpcReq proto.InternalMessageInfo

func (m *RicControlGrpcReq) GetE2NodeID() string {
	if m != nil {
		return m.E2NodeID
	}
	return ""
}

func (m *RicControlGrpcReq) GetPlmnID() string {
	if m != nil {
		return m.PlmnID
	}
	return ""
}

func (m *RicControlGrpcReq) GetRanName() string {
	if m != nil {
		return m.RanName
	}
	return ""
}

func (m *RicControlGrpcReq) GetRICE2APHeaderData() *RICE2APHeader {
	if m != nil {
		return m.RICE2APHeaderData
	}
	return nil
}

func (m *RicControlGrpcReq) GetRICControlHeaderData() *RICControlHeader {
	if m != nil {
		return m.RICControlHeaderData
	}
	return nil
}

func (m *RicControlGrpcReq) GetRICControlMessageData() *RICControlMessage {
	if m != nil {
		return m.RICControlMessageData
	}
	return nil
}

func (m *RicControlGrpcReq) GetRICControlAckReqVal() RICControlAckEnum {
	if m != nil {
		return m.RICControlAckReqVal
	}
	return RICControlAckEnum_RIC_CONTROL_ACK_UNKWON
}

// RicControlGrpc Rsp
type RicControlGrpcRsp struct {
	RspCode              int32    `protobuf:"varint,1,opt,name=rspCode" json:"rspCode,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RicControlGrpcRsp) Reset()         { *m = RicControlGrpcRsp{} }
func (m *RicControlGrpcRsp) String() string { return proto.CompactTextString(m) }
func (*RicControlGrpcRsp) ProtoMessage()    {}
func (*RicControlGrpcRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{7}
}
func (m *RicControlGrpcRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RicControlGrpcRsp.Unmarshal(m, b)
}
func (m *RicControlGrpcRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RicControlGrpcRsp.Marshal(b, m, deterministic)
}
func (dst *RicControlGrpcRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RicControlGrpcRsp.Merge(dst, src)
}
func (m *RicControlGrpcRsp) XXX_Size() int {
	return xxx_messageInfo_RicControlGrpcRsp.Size(m)
}
func (m *RicControlGrpcRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RicControlGrpcRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RicControlGrpcRsp proto.InternalMessageInfo

func (m *RicControlGrpcRsp) GetRspCode() int32 {
	if m != nil {
		return m.RspCode
	}
	return 0
}

func (m *RicControlGrpcRsp) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Member struct {
	PlmnId               string   `protobuf:"bytes,1,opt,name=plmnId" json:"plmnId,omitempty"`
	Sst                  string   `protobuf:"bytes,2,opt,name=sst" json:"sst,omitempty"`
	Sd                   string   `protobuf:"bytes,3,opt,name=sd" json:"sd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{8}
}
func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (dst *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(dst, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetPlmnId() string {
	if m != nil {
		return m.PlmnId
	}
	return ""
}

func (m *Member) GetSst() string {
	if m != nil {
		return m.Sst
	}
	return ""
}

func (m *Member) GetSd() string {
	if m != nil {
		return m.Sd
	}
	return ""
}

type RrmPolicy struct {
	Member               []*Member `protobuf:"bytes,1,rep,name=member" json:"member,omitempty"`
	MinPRB               int64     `protobuf:"varint,2,opt,name=minPRB" json:"minPRB,omitempty"`
	MaxPRB               int64     `protobuf:"varint,3,opt,name=maxPRB" json:"maxPRB,omitempty"`
	DedPRB               int64     `protobuf:"varint,4,opt,name=dedPRB" json:"dedPRB,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RrmPolicy) Reset()         { *m = RrmPolicy{} }
func (m *RrmPolicy) String() string { return proto.CompactTextString(m) }
func (*RrmPolicy) ProtoMessage()    {}
func (*RrmPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{9}
}
func (m *RrmPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RrmPolicy.Unmarshal(m, b)
}
func (m *RrmPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RrmPolicy.Marshal(b, m, deterministic)
}
func (dst *RrmPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RrmPolicy.Merge(dst, src)
}
func (m *RrmPolicy) XXX_Size() int {
	return xxx_messageInfo_RrmPolicy.Size(m)
}
func (m *RrmPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_RrmPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_RrmPolicy proto.InternalMessageInfo

func (m *RrmPolicy) GetMember() []*Member {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *RrmPolicy) GetMinPRB() int64 {
	if m != nil {
		return m.MinPRB
	}
	return 0
}

func (m *RrmPolicy) GetMaxPRB() int64 {
	if m != nil {
		return m.MaxPRB
	}
	return 0
}

func (m *RrmPolicy) GetDedPRB() int64 {
	if m != nil {
		return m.DedPRB
	}
	return 0
}

type RICControlRequest_RRMPolicy struct {
	RanName              string       `protobuf:"bytes,1,opt,name=ranName" json:"ranName,omitempty"`
	RrmPolicy            []*RrmPolicy `protobuf:"bytes,2,rep,name=rrmPolicy" json:"rrmPolicy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RICControlRequest_RRMPolicy) Reset()         { *m = RICControlRequest_RRMPolicy{} }
func (m *RICControlRequest_RRMPolicy) String() string { return proto.CompactTextString(m) }
func (*RICControlRequest_RRMPolicy) ProtoMessage()    {}
func (*RICControlRequest_RRMPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_rc_37af976f3cd13100, []int{10}
}
func (m *RICControlRequest_RRMPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICControlRequest_RRMPolicy.Unmarshal(m, b)
}
func (m *RICControlRequest_RRMPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICControlRequest_RRMPolicy.Marshal(b, m, deterministic)
}
func (dst *RICControlRequest_RRMPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICControlRequest_RRMPolicy.Merge(dst, src)
}
func (m *RICControlRequest_RRMPolicy) XXX_Size() int {
	return xxx_messageInfo_RICControlRequest_RRMPolicy.Size(m)
}
func (m *RICControlRequest_RRMPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_RICControlRequest_RRMPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_RICControlRequest_RRMPolicy proto.InternalMessageInfo

func (m *RICControlRequest_RRMPolicy) GetRanName() string {
	if m != nil {
		return m.RanName
	}
	return ""
}

func (m *RICControlRequest_RRMPolicy) GetRrmPolicy() []*RrmPolicy {
	if m != nil {
		return m.RrmPolicy
	}
	return nil
}

func init() {
	proto.RegisterType((*RICE2APHeader)(nil), "rc.RICE2APHeader")
	proto.RegisterType((*RICControlHeader)(nil), "rc.RICControlHeader")
	proto.RegisterType((*UeId)(nil), "rc.UeId")
	proto.RegisterType((*GNBUEID)(nil), "rc.gNBUEID")
	proto.RegisterType((*Guami)(nil), "rc.Guami")
	proto.RegisterType((*RICControlMessage)(nil), "rc.RICControlMessage")
	proto.RegisterType((*RicControlGrpcReq)(nil), "rc.RicControlGrpcReq")
	proto.RegisterType((*RicControlGrpcRsp)(nil), "rc.RicControlGrpcRsp")
	proto.RegisterType((*Member)(nil), "rc.Member")
	proto.RegisterType((*RrmPolicy)(nil), "rc.RrmPolicy")
	proto.RegisterType((*RICControlRequest_RRMPolicy)(nil), "rc.RICControlRequest_RRMPolicy")
	proto.RegisterEnum("rc.RICControlCellTypeEnum", RICControlCellTypeEnum_name, RICControlCellTypeEnum_value)
	proto.RegisterEnum("rc.RICControlAckEnum", RICControlAckEnum_name, RICControlAckEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MsgComm service

type MsgCommClient interface {
	// gRPC call to Send RICControlReqServiceGrpc
	SendRICControlReqServiceGrpc(ctx context.Context, in *RicControlGrpcReq, opts ...grpc.CallOption) (*RicControlGrpcRsp, error)
	SendRRMPolicyServiceGrpc(ctx context.Context, in *RICControlRequest_RRMPolicy, opts ...grpc.CallOption) (*RicControlGrpcRsp, error)
}

type msgCommClient struct {
	cc *grpc.ClientConn
}

func NewMsgCommClient(cc *grpc.ClientConn) MsgCommClient {
	return &msgCommClient{cc}
}

func (c *msgCommClient) SendRICControlReqServiceGrpc(ctx context.Context, in *RicControlGrpcReq, opts ...grpc.CallOption) (*RicControlGrpcRsp, error) {
	out := new(RicControlGrpcRsp)
	err := grpc.Invoke(ctx, "/rc.MsgComm/SendRICControlReqServiceGrpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgCommClient) SendRRMPolicyServiceGrpc(ctx context.Context, in *RICControlRequest_RRMPolicy, opts ...grpc.CallOption) (*RicControlGrpcRsp, error) {
	out := new(RicControlGrpcRsp)
	err := grpc.Invoke(ctx, "/rc.MsgComm/SendRRMPolicyServiceGrpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MsgComm service

type MsgCommServer interface {
	// gRPC call to Send RICControlReqServiceGrpc
	SendRICControlReqServiceGrpc(context.Context, *RicControlGrpcReq) (*RicControlGrpcRsp, error)
	SendRRMPolicyServiceGrpc(context.Context, *RICControlRequest_RRMPolicy) (*RicControlGrpcRsp, error)
}

func RegisterMsgCommServer(s *grpc.Server, srv MsgCommServer) {
	s.RegisterService(&_MsgComm_serviceDesc, srv)
}

func _MsgComm_SendRICControlReqServiceGrpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RicControlGrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgCommServer).SendRICControlReqServiceGrpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rc.MsgComm/SendRICControlReqServiceGrpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgCommServer).SendRICControlReqServiceGrpc(ctx, req.(*RicControlGrpcReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgComm_SendRRMPolicyServiceGrpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RICControlRequest_RRMPolicy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgCommServer).SendRRMPolicyServiceGrpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rc.MsgComm/SendRRMPolicyServiceGrpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgCommServer).SendRRMPolicyServiceGrpc(ctx, req.(*RICControlRequest_RRMPolicy))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgComm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rc.MsgComm",
	HandlerType: (*MsgCommServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRICControlReqServiceGrpc",
			Handler:    _MsgComm_SendRICControlReqServiceGrpc_Handler,
		},
		{
			MethodName: "SendRRMPolicyServiceGrpc",
			Handler:    _MsgComm_SendRRMPolicyServiceGrpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rc.proto",
}

func init() { proto.RegisterFile("rc.proto", fileDescriptor_rc_37af976f3cd13100) }

var fileDescriptor_rc_37af976f3cd13100 = []byte{
	// 837 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x95, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0xc7, 0x6b, 0x4c, 0x20, 0x1c, 0x9a, 0x2c, 0x99, 0xcd, 0x66, 0x51, 0x36, 0xea, 0x22, 0xab,
	0xad, 0xd0, 0x56, 0x8d, 0x54, 0xfa, 0x00, 0x15, 0x71, 0x0c, 0x6b, 0x05, 0x0c, 0x3a, 0x89, 0xdb,
	0xcb, 0xc8, 0xb1, 0xa7, 0xc8, 0x0a, 0xfe, 0xc8, 0xd8, 0x69, 0x9b, 0x9b, 0x5e, 0xf4, 0xae, 0xef,
	0xd0, 0x37, 0xe8, 0x4b, 0x56, 0x73, 0x3c, 0x60, 0xec, 0xd0, 0x3b, 0xce, 0xef, 0x9c, 0xf3, 0x9f,
	0x39, 0x1f, 0x63, 0xe0, 0x50, 0xf8, 0x97, 0xa9, 0x48, 0xf2, 0x84, 0x35, 0x84, 0x6f, 0xb8, 0x70,
	0x84, 0xb6, 0x69, 0x8d, 0xc6, 0xcb, 0xcf, 0xdc, 0x0b, 0xb8, 0x60, 0x17, 0xd0, 0x41, 0x2f, 0x9e,
	0x3c, 0xc7, 0xbe, 0x1d, 0xf4, 0xb5, 0x81, 0x36, 0xd4, 0xb1, 0x04, 0xec, 0x5b, 0x38, 0x46, 0xdb,
	0x44, 0xfe, 0xf4, 0xcc, 0xb3, 0x3c, 0x11, 0xf6, 0x75, 0xbf, 0x41, 0x21, 0x35, 0x6a, 0xfc, 0x09,
	0x3d, 0xb4, 0x4d, 0x33, 0x89, 0x73, 0x91, 0xac, 0x95, 0xb2, 0x01, 0x5f, 0x2a, 0x70, 0x9b, 0xbf,
	0xac, 0xb9, 0x12, 0xaf, 0x30, 0x36, 0x84, 0x37, 0xca, 0x1e, 0xfb, 0x79, 0x98, 0xc4, 0x76, 0xa0,
	0x0e, 0xa8, 0x63, 0x76, 0x01, 0x4d, 0xd7, 0xb2, 0xaf, 0xfb, 0xfa, 0x40, 0x1b, 0x76, 0x47, 0x87,
	0x97, 0xc2, 0xbf, 0x74, 0xb9, 0x1d, 0x20, 0x51, 0xe3, 0x7b, 0x68, 0x4a, 0x8b, 0x7d, 0x03, 0xed,
	0x69, 0xfc, 0x40, 0x81, 0x1a, 0x05, 0x76, 0x65, 0xe0, 0xca, 0xb9, 0x92, 0x08, 0x37, 0x3e, 0xe3,
	0x1f, 0x0d, 0xda, 0x0a, 0xb2, 0x01, 0x74, 0xbd, 0xe8, 0x57, 0xd7, 0x72, 0xa6, 0xe3, 0xa5, 0x4a,
	0xd3, 0x71, 0x17, 0xb1, 0x8f, 0x70, 0xb0, 0x7a, 0xf6, 0xa2, 0x90, 0xae, 0xd6, 0x1d, 0x75, 0xa4,
	0xe4, 0x54, 0x02, 0x2c, 0x38, 0xfb, 0x1a, 0x8e, 0x56, 0xce, 0x95, 0xe9, 0xba, 0xd6, 0xe4, 0x07,
	0x12, 0xd1, 0x07, 0xfa, 0x50, 0xc7, 0x2a, 0x94, 0xb5, 0x12, 0x30, 0x97, 0xae, 0x65, 0x15, 0x71,
	0x4d, 0x8a, 0xab, 0x63, 0xe3, 0x6f, 0x0d, 0x0e, 0xe8, 0x00, 0xd9, 0xc3, 0x74, 0x36, 0x77, 0xec,
	0x80, 0xc7, 0x79, 0x98, 0xbf, 0xd0, 0xed, 0x3a, 0x58, 0x61, 0x54, 0xc0, 0x7c, 0x82, 0x7c, 0x25,
	0x1b, 0x55, 0x0c, 0xa8, 0x83, 0xbb, 0x88, 0x9d, 0xc3, 0xa1, 0x37, 0x9f, 0xdc, 0xf2, 0x5c, 0xf5,
	0xaf, 0x83, 0x5b, 0x9b, 0x7d, 0x05, 0xe0, 0xcd, 0x27, 0xcb, 0x24, 0x8c, 0x73, 0x2e, 0xfa, 0x4d,
	0xf2, 0xee, 0x10, 0x79, 0x97, 0x93, 0x72, 0xb4, 0x73, 0x9e, 0x65, 0xde, 0x8a, 0xb3, 0x25, 0xbc,
	0x2b, 0xa1, 0xc9, 0xd7, 0xeb, 0xbb, 0x97, 0x94, 0xff, 0xec, 0xad, 0xe9, 0x82, 0xc7, 0xa3, 0x73,
	0xd9, 0xa2, 0xd7, 0x01, 0x56, 0xfc, 0x1c, 0xe1, 0xfe, 0x44, 0x59, 0xe9, 0x9d, 0x27, 0x56, 0x3c,
	0x97, 0x70, 0x5b, 0x46, 0x85, 0x19, 0x7f, 0xe9, 0x70, 0x82, 0xa1, 0xaf, 0xb2, 0xa7, 0x22, 0xf5,
	0x91, 0x3f, 0xc9, 0xea, 0xf8, 0xc8, 0x49, 0x02, 0xae, 0xa6, 0xd7, 0xc1, 0xad, 0xcd, 0xce, 0xa0,
	0x95, 0xae, 0xa3, 0xb2, 0x2d, 0xca, 0x62, 0x7d, 0x68, 0x0b, 0x2f, 0x76, 0xbc, 0x88, 0xab, 0x86,
	0x6c, 0x4c, 0xf6, 0x13, 0x95, 0x5b, 0x3e, 0x90, 0x6b, 0x2f, 0xf7, 0xa8, 0x2d, 0xdd, 0xd1, 0x89,
	0xaa, 0xaa, 0x74, 0xe2, 0xeb, 0x58, 0xf6, 0x19, 0x4e, 0xeb, 0x4f, 0x81, 0x34, 0x0e, 0x48, 0xe3,
	0xb4, 0xda, 0x19, 0x25, 0xb3, 0x37, 0x83, 0xdd, 0xec, 0x36, 0x59, 0x75, 0x9e, 0xa4, 0x5a, 0x24,
	0xf5, 0xae, 0x2a, 0xa5, 0x02, 0x70, 0x7f, 0x0e, 0x9b, 0xc2, 0xdb, 0xd2, 0x31, 0xf6, 0x1f, 0x91,
	0x3f, 0xc9, 0x79, 0xb5, 0x69, 0x5e, 0x35, 0xa9, 0xb1, 0xff, 0x48, 0xa3, 0xda, 0x97, 0x61, 0x2c,
	0x5e, 0xcd, 0x20, 0x4b, 0xa9, 0x9f, 0x59, 0x6a, 0x26, 0x41, 0xf1, 0xcc, 0x0f, 0x70, 0x63, 0xca,
	0xed, 0x0c, 0x78, 0xe6, 0x8b, 0x30, 0x95, 0x0f, 0x79, 0xb3, 0x9d, 0x3b, 0xc8, 0xb8, 0x82, 0xd6,
	0x9c, 0x47, 0x0f, 0x5c, 0x6c, 0xa7, 0x15, 0xa8, 0x39, 0x2a, 0x8b, 0xf5, 0x40, 0xcf, 0xb2, 0x5c,
	0xe5, 0xca, 0x9f, 0xec, 0x18, 0x1a, 0x59, 0xa0, 0x46, 0xd7, 0xc8, 0x02, 0xe3, 0x77, 0xe8, 0xa0,
	0x88, 0x96, 0xc9, 0x3a, 0xf4, 0x5f, 0x98, 0x01, 0xad, 0x88, 0x04, 0xfb, 0xda, 0x40, 0x1f, 0x76,
	0x47, 0x20, 0xab, 0x2b, 0x8e, 0x40, 0xe5, 0x91, 0x47, 0x45, 0x61, 0xbc, 0xc4, 0x2b, 0xf5, 0xbd,
	0x51, 0x16, 0x71, 0xef, 0x0f, 0xc9, 0x75, 0xc5, 0xc9, 0x92, 0x3c, 0xe0, 0x81, 0xe4, 0xcd, 0x82,
	0x17, 0x96, 0x11, 0xc0, 0x87, 0xb2, 0x49, 0xea, 0x8b, 0x78, 0x8f, 0x38, 0x57, 0x57, 0xd9, 0xd9,
	0x33, 0xad, 0xba, 0x67, 0xdf, 0x41, 0x47, 0x6c, 0x6e, 0xdc, 0x6f, 0xd0, 0x3d, 0x8f, 0x68, 0x0a,
	0x1b, 0x88, 0xa5, 0xff, 0xd3, 0x23, 0x9c, 0xed, 0x7f, 0x4d, 0xec, 0x03, 0xbc, 0x47, 0xdb, 0xbc,
	0x37, 0x17, 0xce, 0x1d, 0x2e, 0x66, 0xf7, 0xa6, 0x35, 0x9b, 0xdd, 0xbb, 0xce, 0xcd, 0x2f, 0x0b,
	0xa7, 0xf7, 0x05, 0x7b, 0x4f, 0x33, 0xdf, 0x3a, 0x1d, 0x24, 0x7f, 0x4f, 0xab, 0x67, 0x59, 0xee,
	0x1d, 0x8e, 0x9d, 0xc2, 0xd9, 0xf8, 0x24, 0x76, 0x1f, 0xbc, 0x5a, 0x05, 0x76, 0x4e, 0x37, 0xd8,
	0x66, 0x8c, 0xcd, 0x9b, 0xf2, 0x98, 0x33, 0x60, 0x95, 0x63, 0x16, 0xd2, 0xdd, 0xd3, 0xd8, 0x5b,
	0x78, 0x53, 0xcb, 0xe9, 0x35, 0xd8, 0x29, 0xfd, 0x53, 0x94, 0xc1, 0x92, 0xea, 0xa3, 0x7f, 0x35,
	0x68, 0xcf, 0xb3, 0x95, 0x99, 0x44, 0x11, 0x9b, 0xc1, 0xc5, 0x2d, 0x8f, 0x83, 0x4a, 0x5b, 0x6f,
	0xb9, 0xf8, 0x2d, 0xf4, 0xb9, 0x5c, 0x37, 0x56, 0x2c, 0x6b, 0xfd, 0x33, 0x70, 0xbe, 0x0f, 0x67,
	0x29, 0x43, 0xe8, 0x93, 0xda, 0x66, 0x24, 0xbb, 0x4a, 0x1f, 0xab, 0x6b, 0xff, 0x6a, 0x7c, 0xff,
	0xa3, 0xf9, 0xd0, 0xa2, 0xff, 0xd3, 0x1f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x96, 0xd8, 0xc0,
	0x01, 0x5b, 0x07, 0x00, 0x00,
}
