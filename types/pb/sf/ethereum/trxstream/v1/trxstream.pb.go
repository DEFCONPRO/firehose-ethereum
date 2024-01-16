// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.0
// source: sf/ethereum/trxstream/v1/trxstream.proto

package pbtrxstream

import (
	v2 "github.com/streamingfast/firehose-ethereum/types/pb/sf/ethereum/type/v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionState_Transition int32

const (
	TransactionState_TRANS_INIT                   TransactionState_Transition = 0
	TransactionState_TRANS_POOLED                 TransactionState_Transition = 1
	TransactionState_TRANS_MINED                  TransactionState_Transition = 2
	TransactionState_TRANS_FORKED                 TransactionState_Transition = 3
	TransactionState_TRANS_CONFIRMED              TransactionState_Transition = 4
	TransactionState_TRANS_REPLACED               TransactionState_Transition = 5
	TransactionState_TRANS_SPECULATIVELY_EXECUTED TransactionState_Transition = 6 // makes speculative traces available on a PENDING transaction. May not be emitted if the transaction is seen a block before
)

// Enum value maps for TransactionState_Transition.
var (
	TransactionState_Transition_name = map[int32]string{
		0: "TRANS_INIT",
		1: "TRANS_POOLED",
		2: "TRANS_MINED",
		3: "TRANS_FORKED",
		4: "TRANS_CONFIRMED",
		5: "TRANS_REPLACED",
		6: "TRANS_SPECULATIVELY_EXECUTED",
	}
	TransactionState_Transition_value = map[string]int32{
		"TRANS_INIT":                   0,
		"TRANS_POOLED":                 1,
		"TRANS_MINED":                  2,
		"TRANS_FORKED":                 3,
		"TRANS_CONFIRMED":              4,
		"TRANS_REPLACED":               5,
		"TRANS_SPECULATIVELY_EXECUTED": 6,
	}
)

func (x TransactionState_Transition) Enum() *TransactionState_Transition {
	p := new(TransactionState_Transition)
	*p = x
	return p
}

func (x TransactionState_Transition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionState_Transition) Descriptor() protoreflect.EnumDescriptor {
	return file_sf_ethereum_trxstream_v1_trxstream_proto_enumTypes[0].Descriptor()
}

func (TransactionState_Transition) Type() protoreflect.EnumType {
	return &file_sf_ethereum_trxstream_v1_trxstream_proto_enumTypes[0]
}

func (x TransactionState_Transition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionState_Transition.Descriptor instead.
func (TransactionState_Transition) EnumDescriptor() ([]byte, []int) {
	return file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescGZIP(), []int{1, 0}
}

type TransactionState_State int32

const (
	TransactionState_STATE_UNKNOWN  TransactionState_State = 0
	TransactionState_STATE_PENDING  TransactionState_State = 1
	TransactionState_STATE_IN_BLOCK TransactionState_State = 2
	TransactionState_STATE_REPLACED TransactionState_State = 3
)

// Enum value maps for TransactionState_State.
var (
	TransactionState_State_name = map[int32]string{
		0: "STATE_UNKNOWN",
		1: "STATE_PENDING",
		2: "STATE_IN_BLOCK",
		3: "STATE_REPLACED",
	}
	TransactionState_State_value = map[string]int32{
		"STATE_UNKNOWN":  0,
		"STATE_PENDING":  1,
		"STATE_IN_BLOCK": 2,
		"STATE_REPLACED": 3,
	}
)

func (x TransactionState_State) Enum() *TransactionState_State {
	p := new(TransactionState_State)
	*p = x
	return p
}

func (x TransactionState_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionState_State) Descriptor() protoreflect.EnumDescriptor {
	return file_sf_ethereum_trxstream_v1_trxstream_proto_enumTypes[1].Descriptor()
}

func (TransactionState_State) Type() protoreflect.EnumType {
	return &file_sf_ethereum_trxstream_v1_trxstream_proto_enumTypes[1]
}

func (x TransactionState_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionState_State.Descriptor instead.
func (TransactionState_State) EnumDescriptor() ([]byte, []int) {
	return file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescGZIP(), []int{1, 1}
}

type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_trxstream_v1_trxstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_trxstream_v1_trxstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescGZIP(), []int{0}
}

type TransactionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousState     TransactionState_State      `protobuf:"varint,1,opt,name=previous_state,json=previousState,proto3,enum=sf.ethereum.trxstream.v1.TransactionState_State" json:"previous_state,omitempty"`
	CurrentState      TransactionState_State      `protobuf:"varint,2,opt,name=current_state,json=currentState,proto3,enum=sf.ethereum.trxstream.v1.TransactionState_State" json:"current_state,omitempty"`
	Transition        TransactionState_Transition `protobuf:"varint,10,opt,name=transition,proto3,enum=sf.ethereum.trxstream.v1.TransactionState_Transition" json:"transition,omitempty"`
	Hash              []byte                      `protobuf:"bytes,11,opt,name=hash,proto3" json:"hash,omitempty"`
	Trx               *Transaction                `protobuf:"bytes,3,opt,name=trx,proto3" json:"trx,omitempty"`
	BlockHeader       *v2.BlockHeader             `protobuf:"bytes,4,opt,name=block_header,json=blockHeader,proto3" json:"block_header,omitempty"`
	TransactionTraces *v2.TransactionTrace        `protobuf:"bytes,5,opt,name=transaction_traces,json=transactionTraces,proto3" json:"transaction_traces,omitempty"`
	Confirmation      uint64                      `protobuf:"varint,6,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
	HeadBlockHeader   *v2.BlockHeader             `protobuf:"bytes,7,opt,name=head_block_header,json=headBlockHeader,proto3" json:"head_block_header,omitempty"`
	ReplacedByHash    []byte                      `protobuf:"bytes,8,opt,name=replaced_by_hash,json=replacedByHash,proto3" json:"replaced_by_hash,omitempty"`
	PendingFirstSeen  *timestamppb.Timestamp      `protobuf:"bytes,12,opt,name=pending_first_seen,json=pendingFirstSeen,proto3" json:"pending_first_seen,omitempty"`
	PendingLastSeen   *timestamppb.Timestamp      `protobuf:"bytes,13,opt,name=pending_last_seen,json=pendingLastSeen,proto3" json:"pending_last_seen,omitempty"`
}

func (x *TransactionState) Reset() {
	*x = TransactionState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_trxstream_v1_trxstream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionState) ProtoMessage() {}

func (x *TransactionState) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_trxstream_v1_trxstream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionState.ProtoReflect.Descriptor instead.
func (*TransactionState) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionState) GetPreviousState() TransactionState_State {
	if x != nil {
		return x.PreviousState
	}
	return TransactionState_STATE_UNKNOWN
}

func (x *TransactionState) GetCurrentState() TransactionState_State {
	if x != nil {
		return x.CurrentState
	}
	return TransactionState_STATE_UNKNOWN
}

func (x *TransactionState) GetTransition() TransactionState_Transition {
	if x != nil {
		return x.Transition
	}
	return TransactionState_TRANS_INIT
}

func (x *TransactionState) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *TransactionState) GetTrx() *Transaction {
	if x != nil {
		return x.Trx
	}
	return nil
}

func (x *TransactionState) GetBlockHeader() *v2.BlockHeader {
	if x != nil {
		return x.BlockHeader
	}
	return nil
}

func (x *TransactionState) GetTransactionTraces() *v2.TransactionTrace {
	if x != nil {
		return x.TransactionTraces
	}
	return nil
}

func (x *TransactionState) GetConfirmation() uint64 {
	if x != nil {
		return x.Confirmation
	}
	return 0
}

func (x *TransactionState) GetHeadBlockHeader() *v2.BlockHeader {
	if x != nil {
		return x.HeadBlockHeader
	}
	return nil
}

func (x *TransactionState) GetReplacedByHash() []byte {
	if x != nil {
		return x.ReplacedByHash
	}
	return nil
}

func (x *TransactionState) GetPendingFirstSeen() *timestamppb.Timestamp {
	if x != nil {
		return x.PendingFirstSeen
	}
	return nil
}

func (x *TransactionState) GetPendingLastSeen() *timestamppb.Timestamp {
	if x != nil {
		return x.PendingLastSeen
	}
	return nil
}

// A Transaction not yet in block
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// consensus
	To       []byte     `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Nonce    uint64     `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	GasPrice *v2.BigInt `protobuf:"bytes,3,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	GasLimit uint64     `protobuf:"varint,4,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	Value    *v2.BigInt `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"` // amount of ETH transfered, in addition to used_gas * gas_price, sometimes referred to as `Amount`
	Input    []byte     `protobuf:"bytes,6,opt,name=input,proto3" json:"input,omitempty"`
	V        []byte     `protobuf:"bytes,7,opt,name=v,proto3" json:"v,omitempty"` // signature values
	R        []byte     `protobuf:"bytes,8,opt,name=r,proto3" json:"r,omitempty"`
	S        []byte     `protobuf:"bytes,9,opt,name=s,proto3" json:"s,omitempty"`
	// meta
	Hash []byte `protobuf:"bytes,21,opt,name=hash,proto3" json:"hash,omitempty"`
	From []byte `protobuf:"bytes,22,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_trxstream_v1_trxstream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_trxstream_v1_trxstream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescGZIP(), []int{2}
}

func (x *Transaction) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Transaction) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Transaction) GetGasPrice() *v2.BigInt {
	if x != nil {
		return x.GasPrice
	}
	return nil
}

func (x *Transaction) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *Transaction) GetValue() *v2.BigInt {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Transaction) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Transaction) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

func (x *Transaction) GetR() []byte {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *Transaction) GetS() []byte {
	if x != nil {
		return x.S
	}
	return nil
}

func (x *Transaction) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Transaction) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

var File_sf_ethereum_trxstream_v1_trxstream_proto protoreflect.FileDescriptor

var file_sf_ethereum_trxstream_v1_trxstream_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x66, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x74, 0x72,
	0x78, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x78, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x66, 0x2e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x78, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x66, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa5, 0x08, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x57, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x78, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x55, 0x0a, 0x0d, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x30, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74,
	0x72, 0x78, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x55, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x74, 0x72, 0x78, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x37, 0x0a, 0x03, 0x74,
	0x72, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x78, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x03, 0x74, 0x72, 0x78, 0x12, 0x43, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x66, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x12, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x11, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x11, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x0f, 0x68, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x72, 0x65, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x64, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x48, 0x0a, 0x12, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x10, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x53, 0x65, 0x65, 0x6e, 0x12, 0x46, 0x0a, 0x11, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x22, 0x9c, 0x01,
	0x0a, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x0a,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x5f, 0x4d, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x4b, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x52, 0x4d, 0x45, 0x44, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x5f,
	0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x44, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x55, 0x4c, 0x41, 0x54, 0x49, 0x56, 0x45, 0x4c,
	0x59, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x44, 0x10, 0x06, 0x22, 0x55, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x02, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45,
	0x44, 0x10, 0x03, 0x22, 0xa5, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x67, 0x61, 0x73,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73,
	0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x76, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x01, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x32, 0x7a, 0x0a, 0x11, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x65, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x2c, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74,
	0x72, 0x78, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x78,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x42, 0x5f, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66,
	0x61, 0x73, 0x74, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x2d, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2d, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f,
	0x74, 0x72, 0x78, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x74,
	0x72, 0x78, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescOnce sync.Once
	file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescData = file_sf_ethereum_trxstream_v1_trxstream_proto_rawDesc
)

func file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescGZIP() []byte {
	file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescOnce.Do(func() {
		file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescData)
	})
	return file_sf_ethereum_trxstream_v1_trxstream_proto_rawDescData
}

var file_sf_ethereum_trxstream_v1_trxstream_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sf_ethereum_trxstream_v1_trxstream_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sf_ethereum_trxstream_v1_trxstream_proto_goTypes = []interface{}{
	(TransactionState_Transition)(0), // 0: sf.ethereum.trxstream.v1.TransactionState.Transition
	(TransactionState_State)(0),      // 1: sf.ethereum.trxstream.v1.TransactionState.State
	(*TransactionRequest)(nil),       // 2: sf.ethereum.trxstream.v1.TransactionRequest
	(*TransactionState)(nil),         // 3: sf.ethereum.trxstream.v1.TransactionState
	(*Transaction)(nil),              // 4: sf.ethereum.trxstream.v1.Transaction
	(*v2.BlockHeader)(nil),           // 5: sf.ethereum.type.v2.BlockHeader
	(*v2.TransactionTrace)(nil),      // 6: sf.ethereum.type.v2.TransactionTrace
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
	(*v2.BigInt)(nil),                // 8: sf.ethereum.type.v2.BigInt
}
var file_sf_ethereum_trxstream_v1_trxstream_proto_depIdxs = []int32{
	1,  // 0: sf.ethereum.trxstream.v1.TransactionState.previous_state:type_name -> sf.ethereum.trxstream.v1.TransactionState.State
	1,  // 1: sf.ethereum.trxstream.v1.TransactionState.current_state:type_name -> sf.ethereum.trxstream.v1.TransactionState.State
	0,  // 2: sf.ethereum.trxstream.v1.TransactionState.transition:type_name -> sf.ethereum.trxstream.v1.TransactionState.Transition
	4,  // 3: sf.ethereum.trxstream.v1.TransactionState.trx:type_name -> sf.ethereum.trxstream.v1.Transaction
	5,  // 4: sf.ethereum.trxstream.v1.TransactionState.block_header:type_name -> sf.ethereum.type.v2.BlockHeader
	6,  // 5: sf.ethereum.trxstream.v1.TransactionState.transaction_traces:type_name -> sf.ethereum.type.v2.TransactionTrace
	5,  // 6: sf.ethereum.trxstream.v1.TransactionState.head_block_header:type_name -> sf.ethereum.type.v2.BlockHeader
	7,  // 7: sf.ethereum.trxstream.v1.TransactionState.pending_first_seen:type_name -> google.protobuf.Timestamp
	7,  // 8: sf.ethereum.trxstream.v1.TransactionState.pending_last_seen:type_name -> google.protobuf.Timestamp
	8,  // 9: sf.ethereum.trxstream.v1.Transaction.gas_price:type_name -> sf.ethereum.type.v2.BigInt
	8,  // 10: sf.ethereum.trxstream.v1.Transaction.value:type_name -> sf.ethereum.type.v2.BigInt
	2,  // 11: sf.ethereum.trxstream.v1.TransactionStream.Transactions:input_type -> sf.ethereum.trxstream.v1.TransactionRequest
	4,  // 12: sf.ethereum.trxstream.v1.TransactionStream.Transactions:output_type -> sf.ethereum.trxstream.v1.Transaction
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_sf_ethereum_trxstream_v1_trxstream_proto_init() }
func file_sf_ethereum_trxstream_v1_trxstream_proto_init() {
	if File_sf_ethereum_trxstream_v1_trxstream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_ethereum_trxstream_v1_trxstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
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
		file_sf_ethereum_trxstream_v1_trxstream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionState); i {
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
		file_sf_ethereum_trxstream_v1_trxstream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			RawDescriptor: file_sf_ethereum_trxstream_v1_trxstream_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sf_ethereum_trxstream_v1_trxstream_proto_goTypes,
		DependencyIndexes: file_sf_ethereum_trxstream_v1_trxstream_proto_depIdxs,
		EnumInfos:         file_sf_ethereum_trxstream_v1_trxstream_proto_enumTypes,
		MessageInfos:      file_sf_ethereum_trxstream_v1_trxstream_proto_msgTypes,
	}.Build()
	File_sf_ethereum_trxstream_v1_trxstream_proto = out.File
	file_sf_ethereum_trxstream_v1_trxstream_proto_rawDesc = nil
	file_sf_ethereum_trxstream_v1_trxstream_proto_goTypes = nil
	file_sf_ethereum_trxstream_v1_trxstream_proto_depIdxs = nil
}
