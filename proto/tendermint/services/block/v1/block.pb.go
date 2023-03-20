// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/services/block/v1/block.proto

package v1

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	types1 "github.com/tendermint/tendermint/abci/types"
	types "github.com/tendermint/tendermint/proto/tendermint/types"
	io "io"
	math "math"
	math_bits "math/bits"
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

type GetLatestHeightRequest struct {
}

func (m *GetLatestHeightRequest) Reset()         { *m = GetLatestHeightRequest{} }
func (m *GetLatestHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetLatestHeightRequest) ProtoMessage()    {}
func (*GetLatestHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48acf20d1015667, []int{0}
}
func (m *GetLatestHeightRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestHeightRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestHeightRequest.Merge(m, src)
}
func (m *GetLatestHeightRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestHeightRequest proto.InternalMessageInfo

// GetLatestHeightResponse is a lightweight indicator of the current height of
// the blockchain.
type GetLatestHeightResponse struct {
	// The height of the latest committed block.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *GetLatestHeightResponse) Reset()         { *m = GetLatestHeightResponse{} }
func (m *GetLatestHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetLatestHeightResponse) ProtoMessage()    {}
func (*GetLatestHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48acf20d1015667, []int{1}
}
func (m *GetLatestHeightResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestHeightResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestHeightResponse.Merge(m, src)
}
func (m *GetLatestHeightResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestHeightResponse proto.InternalMessageInfo

func (m *GetLatestHeightResponse) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetBlockRequest struct {
	// The height of the block to get. If empty or set to 0, the block
	// associated with the latest height will be returned.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *GetBlockRequest) Reset()         { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()    {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48acf20d1015667, []int{2}
}
func (m *GetBlockRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockRequest.Merge(m, src)
}
func (m *GetBlockRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockRequest proto.InternalMessageInfo

func (m *GetBlockRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetBlockResponse struct {
	// The height of the block being returned.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// The block ID for the block we are returning.
	BlockId *types.BlockID `protobuf:"bytes,2,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	// Block data for the requested height.
	Block *types.Block `protobuf:"bytes,3,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *GetBlockResponse) Reset()         { *m = GetBlockResponse{} }
func (m *GetBlockResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockResponse) ProtoMessage()    {}
func (*GetBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48acf20d1015667, []int{3}
}
func (m *GetBlockResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockResponse.Merge(m, src)
}
func (m *GetBlockResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockResponse proto.InternalMessageInfo

func (m *GetBlockResponse) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetBlockResponse) GetBlockId() *types.BlockID {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetBlockResponse) GetBlock() *types.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type GetBlockResultsRequest struct {
	// The height of the block results to get. If empty or set to 0, the block
	// results associated with the latest height will be returned.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *GetBlockResultsRequest) Reset()         { *m = GetBlockResultsRequest{} }
func (m *GetBlockResultsRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockResultsRequest) ProtoMessage()    {}
func (*GetBlockResultsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48acf20d1015667, []int{4}
}
func (m *GetBlockResultsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockResultsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockResultsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockResultsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockResultsRequest.Merge(m, src)
}
func (m *GetBlockResultsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockResultsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockResultsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockResultsRequest proto.InternalMessageInfo

func (m *GetBlockResultsRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetBlockResultsResponse struct {
	// The height of the block whose results are returned in this structure.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// All events produced by the ABCI BeginBlock call for the block.
	BeginBlockEvents []*types1.Event `protobuf:"bytes,2,rep,name=begin_block_events,json=beginBlockEvents,proto3" json:"begin_block_events,omitempty"`
	// All transaction results produced by block execution.
	TxResults []*types1.ResponseDeliverTx `protobuf:"bytes,3,rep,name=tx_results,json=txResults,proto3" json:"tx_results,omitempty"`
	// Validator updates during block execution.
	ValidatorUpdates []*types1.ValidatorUpdate `protobuf:"bytes,4,rep,name=validator_updates,json=validatorUpdates,proto3" json:"validator_updates,omitempty"`
	// Consensus parameter updates during block execution.
	ConsensusParamUpdates *types1.ConsensusParams `protobuf:"bytes,5,opt,name=consensus_param_updates,json=consensusParamUpdates,proto3" json:"consensus_param_updates,omitempty"`
	// All events produced by the ABCI EndBlock call.
	EndBlockEvents []*types1.Event `protobuf:"bytes,6,rep,name=end_block_events,json=endBlockEvents,proto3" json:"end_block_events,omitempty"`
}

func (m *GetBlockResultsResponse) Reset()         { *m = GetBlockResultsResponse{} }
func (m *GetBlockResultsResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockResultsResponse) ProtoMessage()    {}
func (*GetBlockResultsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48acf20d1015667, []int{5}
}
func (m *GetBlockResultsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockResultsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockResultsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockResultsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockResultsResponse.Merge(m, src)
}
func (m *GetBlockResultsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockResultsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockResultsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockResultsResponse proto.InternalMessageInfo

func (m *GetBlockResultsResponse) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetBlockResultsResponse) GetBeginBlockEvents() []*types1.Event {
	if m != nil {
		return m.BeginBlockEvents
	}
	return nil
}

func (m *GetBlockResultsResponse) GetTxResults() []*types1.ResponseDeliverTx {
	if m != nil {
		return m.TxResults
	}
	return nil
}

func (m *GetBlockResultsResponse) GetValidatorUpdates() []*types1.ValidatorUpdate {
	if m != nil {
		return m.ValidatorUpdates
	}
	return nil
}

func (m *GetBlockResultsResponse) GetConsensusParamUpdates() *types1.ConsensusParams {
	if m != nil {
		return m.ConsensusParamUpdates
	}
	return nil
}

func (m *GetBlockResultsResponse) GetEndBlockEvents() []*types1.Event {
	if m != nil {
		return m.EndBlockEvents
	}
	return nil
}

func init() {
	proto.RegisterType((*GetLatestHeightRequest)(nil), "tendermint.services.block.v1.GetLatestHeightRequest")
	proto.RegisterType((*GetLatestHeightResponse)(nil), "tendermint.services.block.v1.GetLatestHeightResponse")
	proto.RegisterType((*GetBlockRequest)(nil), "tendermint.services.block.v1.GetBlockRequest")
	proto.RegisterType((*GetBlockResponse)(nil), "tendermint.services.block.v1.GetBlockResponse")
	proto.RegisterType((*GetBlockResultsRequest)(nil), "tendermint.services.block.v1.GetBlockResultsRequest")
	proto.RegisterType((*GetBlockResultsResponse)(nil), "tendermint.services.block.v1.GetBlockResultsResponse")
}

func init() {
	proto.RegisterFile("tendermint/services/block/v1/block.proto", fileDescriptor_d48acf20d1015667)
}

var fileDescriptor_d48acf20d1015667 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x9b, 0x8d, 0x5b, 0x75, 0x16, 0xb4, 0x06, 0x6c, 0xe3, 0x2a, 0xa1, 0xe4, 0xaa, 0x5e,
	0x98, 0xd8, 0xd5, 0x07, 0xd0, 0xdd, 0xca, 0xba, 0xa0, 0x20, 0x83, 0x8a, 0x08, 0x12, 0xf2, 0xe7,
	0xd0, 0x0e, 0xb6, 0x93, 0x98, 0x39, 0x09, 0xf5, 0x29, 0xf4, 0x25, 0x7c, 0x17, 0x2f, 0xf7, 0xd2,
	0x4b, 0x69, 0x5f, 0x44, 0x72, 0xa6, 0xb1, 0x29, 0xa5, 0xed, 0xdd, 0xe4, 0xcc, 0xef, 0xfb, 0xce,
	0x39, 0x5f, 0x18, 0x36, 0x40, 0x90, 0x09, 0xe4, 0x33, 0x21, 0xd1, 0x57, 0x90, 0x97, 0x22, 0x06,
	0xe5, 0x47, 0xd3, 0x34, 0xfe, 0xea, 0x97, 0x43, 0x7d, 0xf0, 0xb2, 0x3c, 0xc5, 0xd4, 0x7a, 0xb4,
	0x26, 0xbd, 0x9a, 0xf4, 0x34, 0x50, 0x0e, 0x4f, 0x1f, 0x36, 0x7c, 0xc2, 0x28, 0x16, 0x3e, 0x7e,
	0xcf, 0x40, 0x69, 0xe9, 0x69, 0x43, 0xaa, 0xeb, 0x07, 0x6e, 0x1b, 0x6d, 0x5d, 0x9b, 0x75, 0x2f,
	0x01, 0xdf, 0x84, 0x08, 0x0a, 0x5f, 0x83, 0x18, 0x4f, 0x90, 0xc3, 0xb7, 0x02, 0x14, 0xba, 0x43,
	0xd6, 0xdb, 0xba, 0x51, 0x59, 0x2a, 0x15, 0x58, 0x5d, 0xd6, 0x9e, 0x50, 0xc5, 0x36, 0xfa, 0xc6,
	0xc0, 0xe4, 0xab, 0x2f, 0xf7, 0x31, 0xbb, 0x7b, 0x09, 0x78, 0x5e, 0xd9, 0xaf, 0x5c, 0x76, 0xa2,
	0x3f, 0x0c, 0xd6, 0x59, 0xb3, 0xfb, 0x7d, 0xad, 0xe7, 0xec, 0x16, 0xcd, 0x1c, 0x88, 0xc4, 0x3e,
	0xea, 0x1b, 0x83, 0x93, 0xb3, 0x07, 0x5e, 0x23, 0x2e, 0xbd, 0x2d, 0x59, 0x5d, 0x8d, 0xf8, 0x4d,
	0x42, 0xaf, 0x12, 0xeb, 0x09, 0x3b, 0xa6, 0xa3, 0x6d, 0x92, 0xa4, 0xb7, 0x43, 0xc2, 0x35, 0xe5,
	0x3e, 0xa5, 0x24, 0xea, 0x81, 0x8a, 0x29, 0xaa, 0x43, 0x3b, 0xfc, 0x32, 0x29, 0xa2, 0x4d, 0xc9,
	0x81, 0x55, 0x46, 0xcc, 0x8a, 0x60, 0x2c, 0x64, 0xa0, 0x17, 0x82, 0x12, 0x24, 0x2a, 0xfb, 0xa8,
	0x6f, 0x0e, 0x4e, 0xce, 0xba, 0xcd, 0x09, 0xab, 0xbf, 0xec, 0xbd, 0xaa, 0xae, 0x79, 0x87, 0x14,
	0xd4, 0x86, 0x0a, 0xca, 0x7a, 0xc9, 0x18, 0xce, 0x83, 0x5c, 0xf7, 0xb4, 0x4d, 0x52, 0xbb, 0x5b,
	0xea, 0x7a, 0x98, 0x11, 0x4c, 0x45, 0x09, 0xf9, 0xfb, 0x39, 0xbf, 0x8d, 0xf3, 0xd5, 0xa0, 0xd6,
	0x5b, 0x76, 0xaf, 0x0c, 0xa7, 0x22, 0x09, 0x31, 0xcd, 0x83, 0x22, 0x4b, 0xaa, 0xff, 0x6c, 0xdf,
	0x20, 0xa7, 0xfe, 0x96, 0xd3, 0xc7, 0x9a, 0xfc, 0x40, 0x20, 0xef, 0x94, 0x9b, 0x05, 0x65, 0x7d,
	0x62, 0xbd, 0xb8, 0xea, 0x25, 0x55, 0xa1, 0x82, 0x2c, 0xcc, 0xc3, 0xd9, 0x7f, 0xd3, 0x63, 0x8a,
	0x7f, 0xdb, 0xf4, 0xa2, 0xe6, 0xdf, 0x55, 0xb8, 0xe2, 0xf7, 0xe3, 0x8d, 0x42, 0xed, 0xfc, 0x82,
	0x75, 0x40, 0x26, 0x9b, 0x79, 0xb5, 0xf7, 0xe6, 0x75, 0x07, 0x64, 0xd2, 0x48, 0xeb, 0xfc, 0xcb,
	0xef, 0x85, 0x63, 0x5c, 0x2f, 0x1c, 0xe3, 0xef, 0xc2, 0x31, 0x7e, 0x2e, 0x9d, 0xd6, 0xf5, 0xd2,
	0x69, 0xfd, 0x59, 0x3a, 0xad, 0xcf, 0x17, 0x63, 0x81, 0x93, 0x22, 0xf2, 0xe2, 0x74, 0xe6, 0x37,
	0x9f, 0xc9, 0xfa, 0x48, 0xaf, 0xc4, 0xdf, 0xf7, 0x8a, 0xa3, 0x36, 0x31, 0xcf, 0xfe, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x08, 0xff, 0xfd, 0x76, 0xec, 0x03, 0x00, 0x00,
}

func (m *GetLatestHeightRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestHeightRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestHeightRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetLatestHeightResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestHeightResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestHeightResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetBlockRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetBlockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockId != nil {
		{
			size, err := m.BlockId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetBlockResultsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockResultsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockResultsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetBlockResultsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockResultsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockResultsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EndBlockEvents) > 0 {
		for iNdEx := len(m.EndBlockEvents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EndBlockEvents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlock(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.ConsensusParamUpdates != nil {
		{
			size, err := m.ConsensusParamUpdates.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ValidatorUpdates) > 0 {
		for iNdEx := len(m.ValidatorUpdates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorUpdates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlock(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.TxResults) > 0 {
		for iNdEx := len(m.TxResults) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TxResults[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlock(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.BeginBlockEvents) > 0 {
		for iNdEx := len(m.BeginBlockEvents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BeginBlockEvents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlock(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Height != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetLatestHeightRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetLatestHeightResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlock(uint64(m.Height))
	}
	return n
}

func (m *GetBlockRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlock(uint64(m.Height))
	}
	return n
}

func (m *GetBlockResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlock(uint64(m.Height))
	}
	if m.BlockId != nil {
		l = m.BlockId.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func (m *GetBlockResultsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlock(uint64(m.Height))
	}
	return n
}

func (m *GetBlockResultsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlock(uint64(m.Height))
	}
	if len(m.BeginBlockEvents) > 0 {
		for _, e := range m.BeginBlockEvents {
			l = e.Size()
			n += 1 + l + sovBlock(uint64(l))
		}
	}
	if len(m.TxResults) > 0 {
		for _, e := range m.TxResults {
			l = e.Size()
			n += 1 + l + sovBlock(uint64(l))
		}
	}
	if len(m.ValidatorUpdates) > 0 {
		for _, e := range m.ValidatorUpdates {
			l = e.Size()
			n += 1 + l + sovBlock(uint64(l))
		}
	}
	if m.ConsensusParamUpdates != nil {
		l = m.ConsensusParamUpdates.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if len(m.EndBlockEvents) > 0 {
		for _, e := range m.EndBlockEvents {
			l = e.Size()
			n += 1 + l + sovBlock(uint64(l))
		}
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetLatestHeightRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetLatestHeightRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetLatestHeightResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetLatestHeightResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetBlockRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetBlockRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetBlockResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetBlockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockId == nil {
				m.BlockId = &types.BlockID{}
			}
			if err := m.BlockId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &types.Block{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetBlockResultsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetBlockResultsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockResultsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetBlockResultsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetBlockResultsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockResultsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginBlockEvents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeginBlockEvents = append(m.BeginBlockEvents, &types1.Event{})
			if err := m.BeginBlockEvents[len(m.BeginBlockEvents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxResults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxResults = append(m.TxResults, &types1.ResponseDeliverTx{})
			if err := m.TxResults[len(m.TxResults)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorUpdates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorUpdates = append(m.ValidatorUpdates, &types1.ValidatorUpdate{})
			if err := m.ValidatorUpdates[len(m.ValidatorUpdates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusParamUpdates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsensusParamUpdates == nil {
				m.ConsensusParamUpdates = &types1.ConsensusParams{}
			}
			if err := m.ConsensusParamUpdates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockEvents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndBlockEvents = append(m.EndBlockEvents, &types1.Event{})
			if err := m.EndBlockEvents[len(m.EndBlockEvents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)