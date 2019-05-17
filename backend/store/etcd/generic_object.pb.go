// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: generic_object.proto

package etcd

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Asset defines an asset agents install as a dependency for a check.
type GenericObject struct {
	// URL is the location of the asset
	revision uint32 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	// Metadata contains the name, namespace, labels and annotations of the asset
	ObjectMeta           `protobuf:"bytes,2,opt,name=metadata,embedded=metadata" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericObject) Reset()         { *m = GenericObject{} }
func (m *GenericObject) String() string { return proto.CompactTextString(m) }
func (*GenericObject) ProtoMessage()    {}
func (*GenericObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_generic_object_f8e756e839b521aa, []int{0}
}
func (m *GenericObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenericObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenericObject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GenericObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericObject.Merge(dst, src)
}
func (m *GenericObject) XXX_Size() int {
	return m.Size()
}
func (m *GenericObject) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericObject.DiscardUnknown(m)
}

var xxx_messageInfo_GenericObject proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenericObject)(nil), "sensu.etcd.GenericObject")
}
func (this *GenericObject) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenericObject)
	if !ok {
		that2, ok := that.(GenericObject)
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
	if this.revision != that1.revision {
		return false
	}
	if !this.ObjectMeta.Equal(&that1.ObjectMeta) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type GenericObjectFace interface {
	Proto() github_com_golang_protobuf_proto.Message
	Getrevision() uint32
	GetObjectMeta() ObjectMeta
}

func (this *GenericObject) Proto() github_com_golang_protobuf_proto.Message {
	return this
}

func (this *GenericObject) TestProto() github_com_golang_protobuf_proto.Message {
	return NewGenericObjectFromFace(this)
}

func (this *GenericObject) Getrevision() uint32 {
	return this.revision
}

func (this *GenericObject) GetObjectMeta() ObjectMeta {
	return this.ObjectMeta
}

func NewGenericObjectFromFace(that GenericObjectFace) *GenericObject {
	this := &GenericObject{}
	this.revision = that.Getrevision()
	this.ObjectMeta = that.GetObjectMeta()
	return this
}

func (m *GenericObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenericObject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.revision != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGenericObject(dAtA, i, uint64(m.revision))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenericObject(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintGenericObject(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedGenericObject(r randyGenericObject, easy bool) *GenericObject {
	this := &GenericObject{}
	this.revision = uint32(r.Uint32())
	v1 := NewPopulatedObjectMeta(r, easy)
	this.ObjectMeta = *v1
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedGenericObject(r, 3)
	}
	return this
}

type randyGenericObject interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneGenericObject(r randyGenericObject) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringGenericObject(r randyGenericObject) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneGenericObject(r)
	}
	return string(tmps)
}
func randUnrecognizedGenericObject(r randyGenericObject, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldGenericObject(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldGenericObject(dAtA []byte, r randyGenericObject, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateGenericObject(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateGenericObject(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateGenericObject(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateGenericObject(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateGenericObject(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateGenericObject(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateGenericObject(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *GenericObject) Size() (n int) {
	var l int
	_ = l
	if m.revision != 0 {
		n += 1 + sovGenericObject(uint64(m.revision))
	}
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenericObject(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGenericObject(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenericObject(x uint64) (n int) {
	return sovGenericObject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenericObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenericObject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenericObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenericObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field revision", wireType)
			}
			m.revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenericObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.revision |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenericObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenericObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenericObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenericObject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenericObject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenericObject
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
					return 0, ErrIntOverflowGenericObject
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenericObject
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenericObject
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenericObject
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenericObject(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenericObject = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenericObject   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("generic_object.proto", fileDescriptor_generic_object_f8e756e839b521aa)
}

var fileDescriptor_generic_object_f8e756e839b521aa = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x4f, 0xcd, 0x4b,
	0x2d, 0xca, 0x4c, 0x8e, 0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x2a, 0x4e, 0xcd, 0x2b, 0x2e, 0xd5, 0x4b, 0x2d, 0x49, 0x4e, 0x91, 0xd2, 0x4d, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b,
	0x49, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x55, 0x8a, 0x2b, 0x37, 0xb5, 0x24,
	0x11, 0xc2, 0x56, 0x9a, 0xc5, 0xc8, 0xc5, 0xeb, 0x0e, 0x31, 0xdf, 0x1f, 0x6c, 0xbc, 0x90, 0x06,
	0x17, 0x47, 0x51, 0x6a, 0x59, 0x66, 0x71, 0x66, 0x7e, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xaf,
	0x13, 0xcf, 0xa3, 0x7b, 0xf2, 0x70, 0xb1, 0x20, 0x38, 0x4b, 0x28, 0x94, 0x8b, 0x03, 0x64, 0x52,
	0x4a, 0x62, 0x49, 0xa2, 0x04, 0x93, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xa4, 0x1e, 0xc4, 0x55, 0xc9,
	0xf9, 0x45, 0xa9, 0x7a, 0x65, 0x46, 0x7a, 0x10, 0x23, 0x7d, 0x53, 0x4b, 0x12, 0x9d, 0xe4, 0x4e,
	0xdc, 0x93, 0x67, 0xb8, 0x70, 0x4f, 0x9e, 0xf1, 0xd5, 0x3d, 0x79, 0x21, 0x98, 0x36, 0x9d, 0xfc,
	0xdc, 0xcc, 0x92, 0xd4, 0xdc, 0x82, 0x92, 0xca, 0x20, 0xb8, 0x51, 0x56, 0x1c, 0x1d, 0x0b, 0xe4,
	0x19, 0x56, 0x2c, 0x90, 0x67, 0x74, 0x52, 0xfa, 0xf1, 0x50, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6,
	0x1d, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x19, 0x8f, 0xe5, 0x18, 0xa2, 0x58, 0x40, 0x7e, 0x4f, 0x62, 0x03, 0xfb, 0xc3, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x74, 0x61, 0xb9, 0x5e, 0x26, 0x01, 0x00, 0x00,
}
