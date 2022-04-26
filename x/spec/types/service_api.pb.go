// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spec/service_api.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ServiceApi struct {
	Name          string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ComputeUnits  uint64         `protobuf:"varint,2,opt,name=computeUnits,proto3" json:"computeUnits,omitempty"`
	Enabled       bool           `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ApiInterfaces []ApiInterface `protobuf:"bytes,4,rep,name=apiInterfaces,proto3" json:"apiInterfaces"`
}

func (m *ServiceApi) Reset()         { *m = ServiceApi{} }
func (m *ServiceApi) String() string { return proto.CompactTextString(m) }
func (*ServiceApi) ProtoMessage()    {}
func (*ServiceApi) Descriptor() ([]byte, []int) {
	return fileDescriptor_3323a3ad252c5ed4, []int{0}
}
func (m *ServiceApi) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceApi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceApi.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceApi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceApi.Merge(m, src)
}
func (m *ServiceApi) XXX_Size() int {
	return m.Size()
}
func (m *ServiceApi) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceApi.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceApi proto.InternalMessageInfo

func (m *ServiceApi) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceApi) GetComputeUnits() uint64 {
	if m != nil {
		return m.ComputeUnits
	}
	return 0
}

func (m *ServiceApi) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *ServiceApi) GetApiInterfaces() []ApiInterface {
	if m != nil {
		return m.ApiInterfaces
	}
	return nil
}

type ApiInterface struct {
	Interface         string `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	Type              string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ExtraComputeUnits uint64 `protobuf:"varint,3,opt,name=extraComputeUnits,proto3" json:"extraComputeUnits,omitempty"`
}

func (m *ApiInterface) Reset()         { *m = ApiInterface{} }
func (m *ApiInterface) String() string { return proto.CompactTextString(m) }
func (*ApiInterface) ProtoMessage()    {}
func (*ApiInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_3323a3ad252c5ed4, []int{1}
}
func (m *ApiInterface) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApiInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApiInterface.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApiInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiInterface.Merge(m, src)
}
func (m *ApiInterface) XXX_Size() int {
	return m.Size()
}
func (m *ApiInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiInterface.DiscardUnknown(m)
}

var xxx_messageInfo_ApiInterface proto.InternalMessageInfo

func (m *ApiInterface) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *ApiInterface) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ApiInterface) GetExtraComputeUnits() uint64 {
	if m != nil {
		return m.ExtraComputeUnits
	}
	return 0
}

func init() {
	proto.RegisterType((*ServiceApi)(nil), "lavanet.lava.spec.ServiceApi")
	proto.RegisterType((*ApiInterface)(nil), "lavanet.lava.spec.ApiInterface")
}

func init() { proto.RegisterFile("spec/service_api.proto", fileDescriptor_3323a3ad252c5ed4) }

var fileDescriptor_3323a3ad252c5ed4 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x2e, 0x48, 0x4d,
	0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x8d, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0xcc, 0x49, 0x2c, 0x4b, 0xcc, 0x4b, 0x2d, 0xd1, 0x03, 0xd1, 0x7a, 0x20,
	0x45, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x59, 0x7d, 0x10, 0x0b, 0xa2, 0x50, 0x69, 0x35,
	0x23, 0x17, 0x57, 0x30, 0x44, 0xbb, 0x63, 0x41, 0xa6, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e,
	0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0xa4, 0xc4, 0xc5, 0x93, 0x9c, 0x9f,
	0x5b, 0x50, 0x5a, 0x92, 0x1a, 0x9a, 0x97, 0x59, 0x52, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x12,
	0x84, 0x22, 0x26, 0x24, 0xc1, 0xc5, 0x9e, 0x9a, 0x97, 0x98, 0x94, 0x93, 0x9a, 0x22, 0xc1, 0xac,
	0xc0, 0xa8, 0xc1, 0x11, 0x04, 0xe3, 0x0a, 0x79, 0x73, 0xf1, 0x26, 0x16, 0x64, 0x7a, 0xe6, 0x95,
	0xa4, 0x16, 0xa5, 0x25, 0x26, 0xa7, 0x16, 0x4b, 0xb0, 0x28, 0x30, 0x6b, 0x70, 0x1b, 0xc9, 0xeb,
	0x61, 0xb8, 0x50, 0xcf, 0x11, 0x49, 0x9d, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0xa8, 0x7a,
	0x95, 0xf2, 0xb8, 0x78, 0x90, 0x15, 0x09, 0xc9, 0x70, 0x71, 0x66, 0xc2, 0x38, 0x50, 0x37, 0x23,
	0x04, 0x40, 0x9e, 0x29, 0xa9, 0x2c, 0x48, 0x05, 0x3b, 0x98, 0x33, 0x08, 0xcc, 0x16, 0xd2, 0xe1,
	0x12, 0x4c, 0xad, 0x28, 0x29, 0x4a, 0x74, 0x46, 0xf6, 0x11, 0x33, 0xd8, 0x47, 0x98, 0x12, 0x4e,
	0x4e, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94,
	0x4a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xd4, 0x37, 0x60, 0x5a,
	0xbf, 0x42, 0x1f, 0x1c, 0x2d, 0x20, 0x0b, 0x8b, 0x93, 0xd8, 0xc0, 0x01, 0x6d, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x53, 0xf6, 0x24, 0x3c, 0xab, 0x01, 0x00, 0x00,
}

func (this *ServiceApi) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceApi)
	if !ok {
		that2, ok := that.(ServiceApi)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.ComputeUnits != that1.ComputeUnits {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if len(this.ApiInterfaces) != len(that1.ApiInterfaces) {
		return false
	}
	for i := range this.ApiInterfaces {
		if !this.ApiInterfaces[i].Equal(&that1.ApiInterfaces[i]) {
			return false
		}
	}
	return true
}
func (this *ApiInterface) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ApiInterface)
	if !ok {
		that2, ok := that.(ApiInterface)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Interface != that1.Interface {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.ExtraComputeUnits != that1.ExtraComputeUnits {
		return false
	}
	return true
}
func (m *ServiceApi) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceApi) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceApi) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ApiInterfaces) > 0 {
		for iNdEx := len(m.ApiInterfaces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApiInterfaces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintServiceApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.ComputeUnits != 0 {
		i = encodeVarintServiceApi(dAtA, i, uint64(m.ComputeUnits))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintServiceApi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ApiInterface) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApiInterface) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApiInterface) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExtraComputeUnits != 0 {
		i = encodeVarintServiceApi(dAtA, i, uint64(m.ExtraComputeUnits))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintServiceApi(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Interface) > 0 {
		i -= len(m.Interface)
		copy(dAtA[i:], m.Interface)
		i = encodeVarintServiceApi(dAtA, i, uint64(len(m.Interface)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintServiceApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovServiceApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ServiceApi) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovServiceApi(uint64(l))
	}
	if m.ComputeUnits != 0 {
		n += 1 + sovServiceApi(uint64(m.ComputeUnits))
	}
	if m.Enabled {
		n += 2
	}
	if len(m.ApiInterfaces) > 0 {
		for _, e := range m.ApiInterfaces {
			l = e.Size()
			n += 1 + l + sovServiceApi(uint64(l))
		}
	}
	return n
}

func (m *ApiInterface) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Interface)
	if l > 0 {
		n += 1 + l + sovServiceApi(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovServiceApi(uint64(l))
	}
	if m.ExtraComputeUnits != 0 {
		n += 1 + sovServiceApi(uint64(m.ExtraComputeUnits))
	}
	return n
}

func sovServiceApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozServiceApi(x uint64) (n int) {
	return sovServiceApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ServiceApi) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceApi
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
			return fmt.Errorf("proto: ServiceApi: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceApi: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
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
				return ErrInvalidLengthServiceApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeUnits", wireType)
			}
			m.ComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiInterfaces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
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
				return ErrInvalidLengthServiceApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServiceApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiInterfaces = append(m.ApiInterfaces, ApiInterface{})
			if err := m.ApiInterfaces[len(m.ApiInterfaces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServiceApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServiceApi
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
func (m *ApiInterface) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceApi
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
			return fmt.Errorf("proto: ApiInterface: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApiInterface: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interface", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
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
				return ErrInvalidLengthServiceApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Interface = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
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
				return ErrInvalidLengthServiceApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraComputeUnits", wireType)
			}
			m.ExtraComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExtraComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipServiceApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServiceApi
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
func skipServiceApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServiceApi
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
					return 0, ErrIntOverflowServiceApi
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
					return 0, ErrIntOverflowServiceApi
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
				return 0, ErrInvalidLengthServiceApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupServiceApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthServiceApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthServiceApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServiceApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupServiceApi = fmt.Errorf("proto: unexpected end of group")
)
