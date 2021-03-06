// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/query/generated/proto/admin/namespace.proto

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package admin

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import namespace1 "github.com/m3db/m3/src/dbnode/generated/proto/namespace"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NamespaceGetResponse struct {
	Registry *namespace1.Registry `protobuf:"bytes,1,opt,name=registry" json:"registry,omitempty"`
}

func (m *NamespaceGetResponse) Reset()                    { *m = NamespaceGetResponse{} }
func (m *NamespaceGetResponse) String() string            { return proto.CompactTextString(m) }
func (*NamespaceGetResponse) ProtoMessage()               {}
func (*NamespaceGetResponse) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{0} }

func (m *NamespaceGetResponse) GetRegistry() *namespace1.Registry {
	if m != nil {
		return m.Registry
	}
	return nil
}

type NamespaceAddRequest struct {
	Name    string                       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Options *namespace1.NamespaceOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
}

func (m *NamespaceAddRequest) Reset()                    { *m = NamespaceAddRequest{} }
func (m *NamespaceAddRequest) String() string            { return proto.CompactTextString(m) }
func (*NamespaceAddRequest) ProtoMessage()               {}
func (*NamespaceAddRequest) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{1} }

func (m *NamespaceAddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamespaceAddRequest) GetOptions() *namespace1.NamespaceOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*NamespaceGetResponse)(nil), "admin.NamespaceGetResponse")
	proto.RegisterType((*NamespaceAddRequest)(nil), "admin.NamespaceAddRequest")
}
func (m *NamespaceGetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceGetResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Registry != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.Registry.Size()))
		n1, err := m.Registry.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *NamespaceAddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceAddRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Options != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.Options.Size()))
		n2, err := m.Options.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintNamespace(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NamespaceGetResponse) Size() (n int) {
	var l int
	_ = l
	if m.Registry != nil {
		l = m.Registry.Size()
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func (m *NamespaceAddRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	if m.Options != nil {
		l = m.Options.Size()
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func sovNamespace(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNamespace(x uint64) (n int) {
	return sovNamespace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NamespaceGetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: NamespaceGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Registry == nil {
				m.Registry = &namespace1.Registry{}
			}
			if err := m.Registry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func (m *NamespaceAddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: NamespaceAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Options == nil {
				m.Options = &namespace1.NamespaceOptions{}
			}
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
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
func skipNamespace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNamespace
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
					return 0, ErrIntOverflowNamespace
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
					return 0, ErrIntOverflowNamespace
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
				return 0, ErrInvalidLengthNamespace
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNamespace
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
				next, err := skipNamespace(dAtA[start:])
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
	ErrInvalidLengthNamespace = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNamespace   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/query/generated/proto/admin/namespace.proto", fileDescriptorNamespace)
}

var fileDescriptorNamespace = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xcf, 0x3f, 0x4a, 0x04, 0x31,
	0x14, 0x06, 0x70, 0x23, 0xfe, 0x8d, 0x8d, 0x64, 0x2d, 0x16, 0x85, 0x20, 0x5b, 0x59, 0xcd, 0x03,
	0x07, 0x0f, 0xe0, 0x36, 0xd3, 0x29, 0xe4, 0x04, 0x26, 0x93, 0xc7, 0x98, 0x22, 0x7f, 0x36, 0xc9,
	0x14, 0x7b, 0x0b, 0x8f, 0x65, 0xe9, 0x11, 0x64, 0xbc, 0x88, 0x18, 0x9d, 0x28, 0x8a, 0x5d, 0xf8,
	0xf2, 0x7d, 0x3f, 0x78, 0x74, 0x3d, 0x98, 0xfc, 0x38, 0xaa, 0xa6, 0xf7, 0x16, 0x6c, 0xab, 0x15,
	0xd8, 0x16, 0x52, 0xec, 0x61, 0x33, 0x62, 0xdc, 0xc2, 0x80, 0x0e, 0xa3, 0xcc, 0xa8, 0x21, 0x44,
	0x9f, 0x3d, 0x48, 0x6d, 0x8d, 0x03, 0x27, 0x2d, 0xa6, 0x20, 0x7b, 0x6c, 0x4a, 0xca, 0xf6, 0x4b,
	0x7c, 0xde, 0xfd, 0x43, 0x69, 0xe5, 0xbc, 0xc6, 0x3f, 0x56, 0x55, 0x7e, 0x7b, 0xab, 0x8e, 0x9e,
	0xdd, 0xcd, 0x51, 0x87, 0x59, 0x60, 0x0a, 0xde, 0x25, 0x64, 0x40, 0x8f, 0x22, 0x0e, 0x26, 0xe5,
	0xb8, 0x5d, 0x92, 0x4b, 0x72, 0x75, 0x72, 0xbd, 0x68, 0xbe, 0xb7, 0xe2, 0xeb, 0x4b, 0xd4, 0xd2,
	0xea, 0x81, 0x2e, 0x2a, 0x74, 0xab, 0xb5, 0xc0, 0xcd, 0x88, 0x29, 0x33, 0x46, 0xf7, 0x3e, 0x66,
	0xc5, 0x38, 0x16, 0xe5, 0xcd, 0x6e, 0xe8, 0xa1, 0x0f, 0xd9, 0x78, 0x97, 0x96, 0xbb, 0x85, 0xbe,
	0xf8, 0x41, 0x57, 0xe4, 0xfe, 0xb3, 0x22, 0xe6, 0xee, 0xfa, 0xf4, 0x79, 0xe2, 0xe4, 0x65, 0xe2,
	0xe4, 0x75, 0xe2, 0xe4, 0xe9, 0x8d, 0xef, 0xa8, 0x83, 0x72, 0x43, 0xfb, 0x1e, 0x00, 0x00, 0xff,
	0xff, 0xa1, 0x66, 0xec, 0x92, 0x59, 0x01, 0x00, 0x00,
}
