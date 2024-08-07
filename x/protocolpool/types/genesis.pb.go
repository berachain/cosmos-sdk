// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/protocolpool/v1/genesis.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the protocolpool module's genesis state.
type GenesisState struct {
	// ContinuousFund defines the continuous funds at genesis.
	ContinuousFund []*ContinuousFund `protobuf:"bytes,1,rep,name=continuous_fund,json=continuousFund,proto3" json:"continuous_fund,omitempty"`
	// Budget defines the budget proposals at genesis.
	Budget []*Budget `protobuf:"bytes,2,rep,name=budget,proto3" json:"budget,omitempty"`
	// last_balance contains the amount of tokens yet to be distributed, will be zero if
	// there are no funds to distribute.
	LastBalance cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=last_balance,json=lastBalance,proto3,customtype=cosmossdk.io/math.Int" json:"last_balance"`
	// distributions contains the list of distributions to be made to continuous
	// funds and budgets. It contains time in order to distribute to non-expired
	// funds only.
	Distributions []*Distribution `protobuf:"bytes,4,rep,name=distributions,proto3" json:"distributions,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_72560a99455b4146, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetContinuousFund() []*ContinuousFund {
	if m != nil {
		return m.ContinuousFund
	}
	return nil
}

func (m *GenesisState) GetBudget() []*Budget {
	if m != nil {
		return m.Budget
	}
	return nil
}

func (m *GenesisState) GetDistributions() []*Distribution {
	if m != nil {
		return m.Distributions
	}
	return nil
}

type Distribution struct {
	Amount cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
	Time   *time.Time            `protobuf:"bytes,6,opt,name=time,proto3,stdtime" json:"time,omitempty"`
}

func (m *Distribution) Reset()         { *m = Distribution{} }
func (m *Distribution) String() string { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()    {}
func (*Distribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_72560a99455b4146, []int{1}
}
func (m *Distribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Distribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Distribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Distribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Distribution.Merge(m, src)
}
func (m *Distribution) XXX_Size() int {
	return m.Size()
}
func (m *Distribution) XXX_DiscardUnknown() {
	xxx_messageInfo_Distribution.DiscardUnknown(m)
}

var xxx_messageInfo_Distribution proto.InternalMessageInfo

func (m *Distribution) GetTime() *time.Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.protocolpool.v1.GenesisState")
	proto.RegisterType((*Distribution)(nil), "cosmos.protocolpool.v1.Distribution")
}

func init() {
	proto.RegisterFile("cosmos/protocolpool/v1/genesis.proto", fileDescriptor_72560a99455b4146)
}

var fileDescriptor_72560a99455b4146 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0xca, 0xd3, 0x40,
	0x14, 0xc5, 0x33, 0x5f, 0x4b, 0xc0, 0x69, 0x55, 0x08, 0x2a, 0x31, 0x8b, 0xa4, 0x96, 0x22, 0x05,
	0x71, 0x42, 0xab, 0xb8, 0x71, 0x97, 0x8a, 0x52, 0x17, 0x0a, 0xd1, 0x95, 0x9b, 0x92, 0x3f, 0xd3,
	0x18, 0x4c, 0xe6, 0x86, 0xce, 0x4c, 0xd1, 0x47, 0x70, 0xd7, 0x77, 0xd1, 0x87, 0xe8, 0xb2, 0xb8,
	0x12, 0x17, 0x55, 0xda, 0x17, 0x91, 0xce, 0xa4, 0x92, 0x80, 0xdd, 0x7c, 0xbb, 0xc9, 0xbd, 0xbf,
	0x73, 0x72, 0xe0, 0x5c, 0x3c, 0x4a, 0x80, 0x97, 0xc0, 0xfd, 0x6a, 0x05, 0x02, 0x12, 0x28, 0x2a,
	0x80, 0xc2, 0x5f, 0x4f, 0xfc, 0x8c, 0x32, 0xca, 0x73, 0x4e, 0xd4, 0xdc, 0xba, 0xa7, 0x29, 0xd2,
	0xa4, 0xc8, 0x7a, 0xe2, 0x0c, 0x2f, 0xa8, 0xc5, 0x97, 0x8a, 0xd6, 0xb4, 0x73, 0x27, 0x83, 0x0c,
	0xd4, 0xd3, 0x3f, 0xbd, 0xea, 0xe9, 0x7d, 0xad, 0x5c, 0xe8, 0x45, 0xd3, 0xde, 0xf1, 0x32, 0x80,
	0xac, 0xa0, 0xda, 0x34, 0x96, 0x4b, 0x5f, 0xe4, 0x25, 0xe5, 0x22, 0x2a, 0x2b, 0x0d, 0x0c, 0xbf,
	0x5d, 0xe1, 0xfe, 0x2b, 0x9d, 0xef, 0x9d, 0x88, 0x04, 0xb5, 0xde, 0xe2, 0xdb, 0x09, 0x30, 0x91,
	0x33, 0x09, 0x92, 0x2f, 0x96, 0x92, 0xa5, 0x36, 0x1a, 0x74, 0xc6, 0xbd, 0xe9, 0x43, 0xf2, 0xff,
	0xe0, 0x64, 0xf6, 0x0f, 0x7f, 0x29, 0x59, 0x1a, 0xde, 0x4a, 0x5a, 0xdf, 0xd6, 0x33, 0x6c, 0xc6,
	0x32, 0xcd, 0xa8, 0xb0, 0xaf, 0x94, 0x8f, 0x7b, 0xc9, 0x27, 0x50, 0x54, 0x58, 0xd3, 0xd6, 0x1b,
	0xdc, 0x2f, 0x22, 0x2e, 0x16, 0x71, 0x54, 0x44, 0x2c, 0xa1, 0x76, 0x67, 0x80, 0xc6, 0x37, 0x82,
	0x47, 0xdb, 0xbd, 0x67, 0xfc, 0xda, 0x7b, 0x77, 0xb5, 0x09, 0x4f, 0x3f, 0x91, 0x1c, 0xfc, 0x32,
	0x12, 0x1f, 0xc9, 0x9c, 0x89, 0x1f, 0xdf, 0x1f, 0xe3, 0xda, 0x7d, 0xce, 0x44, 0xd8, 0x3b, 0x19,
	0x04, 0x5a, 0x6f, 0xbd, 0xc6, 0x37, 0xd3, 0x9c, 0x8b, 0x55, 0x1e, 0x4b, 0x91, 0x03, 0xe3, 0x76,
	0x57, 0xc5, 0x19, 0x5d, 0x8a, 0xf3, 0xa2, 0x01, 0x87, 0x6d, 0xe9, 0xf0, 0x2b, 0xc2, 0xfd, 0xe6,
	0xde, 0x9a, 0x61, 0x33, 0x2a, 0x41, 0x32, 0x71, 0x9d, 0x98, 0xb5, 0xd4, 0x7a, 0x8a, 0xbb, 0xa7,
	0x7a, 0x6c, 0x73, 0x80, 0xc6, 0xbd, 0xa9, 0x43, 0x74, 0x77, 0xe4, 0xdc, 0x1d, 0x79, 0x7f, 0xee,
	0x2e, 0xe8, 0x6e, 0x7e, 0x7b, 0x28, 0x54, 0x74, 0xf0, 0x7c, 0x7b, 0x70, 0xd1, 0xee, 0xe0, 0xa2,
	0x3f, 0x07, 0x17, 0x6d, 0x8e, 0xae, 0xb1, 0x3b, 0xba, 0xc6, 0xcf, 0xa3, 0x6b, 0x7c, 0x78, 0xd0,
	0xfa, 0xf9, 0xe7, 0xf6, 0x65, 0xa9, 0xb3, 0x8a, 0x4d, 0x35, 0x7b, 0xf2, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x0c, 0x77, 0xdd, 0xf7, 0xbb, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Distributions) > 0 {
		for iNdEx := len(m.Distributions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Distributions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.LastBalance.Size()
		i -= size
		if _, err := m.LastBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Budget) > 0 {
		for iNdEx := len(m.Budget) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Budget[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ContinuousFund) > 0 {
		for iNdEx := len(m.ContinuousFund) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContinuousFund[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Distribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Distribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Distribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Time != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Time):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintGenesis(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ContinuousFund) > 0 {
		for _, e := range m.ContinuousFund {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Budget) > 0 {
		for _, e := range m.Budget {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.LastBalance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Distributions) > 0 {
		for _, e := range m.Distributions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Distribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.Time != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Time)
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContinuousFund", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContinuousFund = append(m.ContinuousFund, &ContinuousFund{})
			if err := m.ContinuousFund[len(m.ContinuousFund)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Budget", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Budget = append(m.Budget, &Budget{})
			if err := m.Budget[len(m.Budget)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distributions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Distributions = append(m.Distributions, &Distribution{})
			if err := m.Distributions[len(m.Distributions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Distribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Distribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Distribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Time == nil {
				m.Time = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
