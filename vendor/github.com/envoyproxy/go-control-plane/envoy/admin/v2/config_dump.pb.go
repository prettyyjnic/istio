// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v2/config_dump.proto

/*
	Package envoy_admin_v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/admin/v2/config_dump.proto

	It has these top-level messages:
		ConfigDump
		RouteConfigDump
*/
package envoy_admin_v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import envoy_api_v21 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The /config_dump admin endpoint uses this wrapper message to maintain and serve arbitrary
// configuration information from any component in Envoy.
// TODO(jsedgwick) In the future, we may want to formalize this further with an RPC for config_dump,
// and perhaps even with an RPC per config type. That strategy across all endpoints will allow for
// more flexibility w.r.t. protocol, serialization, parameters, etc.
type ConfigDump struct {
	// This map is serialized and dumped in its entirety at the /config_dump endpoint.
	//
	// Keys should be a short descriptor of the config object they map to. For example, Envoy's HTTP
	// routing subsystem might use "routes" as the key for its config, for which it uses the message
	// RouteConfigDump as defined below. In the future, the key will also be used to filter the output
	// of the /config_dump endpoint.
	Configs map[string]google_protobuf.Any `protobuf:"bytes,1,rep,name=configs" json:"configs" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ConfigDump) Reset()                    { *m = ConfigDump{} }
func (m *ConfigDump) String() string            { return proto.CompactTextString(m) }
func (*ConfigDump) ProtoMessage()               {}
func (*ConfigDump) Descriptor() ([]byte, []int) { return fileDescriptorConfigDump, []int{0} }

func (m *ConfigDump) GetConfigs() map[string]google_protobuf.Any {
	if m != nil {
		return m.Configs
	}
	return nil
}

// Envoy's RDS implementation fills this message with all currently loaded routes, as described by
// their RouteConnfiguration objects. Static routes configured in the bootstrap configuration are
// separated from those configured dynamically via RDS. This message is available at the
// /config_dump admin endpoint.
type RouteConfigDump struct {
	StaticRouteConfigs  []envoy_api_v21.RouteConfiguration `protobuf:"bytes,1,rep,name=static_route_configs,json=staticRouteConfigs" json:"static_route_configs"`
	DynamicRouteConfigs []envoy_api_v21.RouteConfiguration `protobuf:"bytes,2,rep,name=dynamic_route_configs,json=dynamicRouteConfigs" json:"dynamic_route_configs"`
}

func (m *RouteConfigDump) Reset()                    { *m = RouteConfigDump{} }
func (m *RouteConfigDump) String() string            { return proto.CompactTextString(m) }
func (*RouteConfigDump) ProtoMessage()               {}
func (*RouteConfigDump) Descriptor() ([]byte, []int) { return fileDescriptorConfigDump, []int{1} }

func (m *RouteConfigDump) GetStaticRouteConfigs() []envoy_api_v21.RouteConfiguration {
	if m != nil {
		return m.StaticRouteConfigs
	}
	return nil
}

func (m *RouteConfigDump) GetDynamicRouteConfigs() []envoy_api_v21.RouteConfiguration {
	if m != nil {
		return m.DynamicRouteConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigDump)(nil), "envoy.admin.v2.ConfigDump")
	proto.RegisterType((*RouteConfigDump)(nil), "envoy.admin.v2.RouteConfigDump")
}
func (m *ConfigDump) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigDump) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Configs) > 0 {
		for k, _ := range m.Configs {
			dAtA[i] = 0xa
			i++
			v := m.Configs[k]
			msgSize := 0
			if (&v) != nil {
				msgSize = (&v).Size()
				msgSize += 1 + sovConfigDump(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovConfigDump(uint64(len(k))) + msgSize
			i = encodeVarintConfigDump(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfigDump(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintConfigDump(dAtA, i, uint64((&v).Size()))
			n1, err := (&v).MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n1
		}
	}
	return i, nil
}

func (m *RouteConfigDump) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RouteConfigDump) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StaticRouteConfigs) > 0 {
		for _, msg := range m.StaticRouteConfigs {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfigDump(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.DynamicRouteConfigs) > 0 {
		for _, msg := range m.DynamicRouteConfigs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintConfigDump(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintConfigDump(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ConfigDump) Size() (n int) {
	var l int
	_ = l
	if len(m.Configs) > 0 {
		for k, v := range m.Configs {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovConfigDump(uint64(len(k))) + 1 + l + sovConfigDump(uint64(l))
			n += mapEntrySize + 1 + sovConfigDump(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *RouteConfigDump) Size() (n int) {
	var l int
	_ = l
	if len(m.StaticRouteConfigs) > 0 {
		for _, e := range m.StaticRouteConfigs {
			l = e.Size()
			n += 1 + l + sovConfigDump(uint64(l))
		}
	}
	if len(m.DynamicRouteConfigs) > 0 {
		for _, e := range m.DynamicRouteConfigs {
			l = e.Size()
			n += 1 + l + sovConfigDump(uint64(l))
		}
	}
	return n
}

func sovConfigDump(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfigDump(x uint64) (n int) {
	return sovConfigDump(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConfigDump) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfigDump
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
			return fmt.Errorf("proto: ConfigDump: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigDump: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Configs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigDump
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
				return ErrInvalidLengthConfigDump
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Configs == nil {
				m.Configs = make(map[string]google_protobuf.Any)
			}
			var mapkey string
			mapvalue := &google_protobuf.Any{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfigDump
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfigDump
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthConfigDump
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfigDump
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthConfigDump
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthConfigDump
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &google_protobuf.Any{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipConfigDump(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthConfigDump
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Configs[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfigDump(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfigDump
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
func (m *RouteConfigDump) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfigDump
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
			return fmt.Errorf("proto: RouteConfigDump: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RouteConfigDump: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StaticRouteConfigs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigDump
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
				return ErrInvalidLengthConfigDump
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StaticRouteConfigs = append(m.StaticRouteConfigs, envoy_api_v21.RouteConfiguration{})
			if err := m.StaticRouteConfigs[len(m.StaticRouteConfigs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DynamicRouteConfigs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigDump
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
				return ErrInvalidLengthConfigDump
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DynamicRouteConfigs = append(m.DynamicRouteConfigs, envoy_api_v21.RouteConfiguration{})
			if err := m.DynamicRouteConfigs[len(m.DynamicRouteConfigs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfigDump(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfigDump
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
func skipConfigDump(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfigDump
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
					return 0, ErrIntOverflowConfigDump
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
					return 0, ErrIntOverflowConfigDump
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
				return 0, ErrInvalidLengthConfigDump
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfigDump
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
				next, err := skipConfigDump(dAtA[start:])
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
	ErrInvalidLengthConfigDump = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfigDump   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/admin/v2/config_dump.proto", fileDescriptorConfigDump) }

var fileDescriptorConfigDump = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x7f, 0xb7, 0x3f, 0x20, 0xdc, 0x0a, 0x90, 0x29, 0xa8, 0x64, 0x08, 0x51, 0x17, 0x22,
	0x06, 0x5b, 0x0a, 0x0b, 0x62, 0xa3, 0x80, 0x58, 0x51, 0x26, 0xc4, 0x12, 0xb9, 0x4d, 0x1a, 0x59,
	0x34, 0x76, 0x94, 0xd8, 0x91, 0xfc, 0x44, 0x3c, 0x0a, 0x1d, 0x79, 0x02, 0x84, 0xf2, 0x24, 0x28,
	0x76, 0x22, 0x52, 0x31, 0xb1, 0x1d, 0xdd, 0x73, 0xee, 0x77, 0xcf, 0x85, 0x5e, 0xc2, 0x2b, 0xa1,
	0x09, 0x8d, 0x33, 0xc6, 0x49, 0x15, 0x90, 0xa5, 0xe0, 0x2b, 0x96, 0x46, 0xb1, 0xca, 0x72, 0x9c,
	0x17, 0x42, 0x0a, 0x74, 0x60, 0x12, 0xd8, 0x24, 0x70, 0x15, 0x38, 0x67, 0xa9, 0x10, 0xe9, 0x3a,
	0x21, 0xc6, 0x5d, 0xa8, 0x15, 0xa1, 0x5c, 0xdb, 0xa8, 0x73, 0xda, 0xc2, 0x72, 0xd6, 0xa0, 0x8a,
	0xb8, 0x6c, 0xe7, 0x93, 0x54, 0xa4, 0xc2, 0x48, 0xd2, 0x28, 0x3b, 0x9d, 0xbd, 0x01, 0x08, 0xef,
	0xcc, 0xb9, 0x7b, 0x95, 0xe5, 0xe8, 0x11, 0xee, 0xd9, 0xe3, 0xe5, 0x14, 0x78, 0x43, 0x7f, 0x14,
	0x5c, 0xe0, 0xed, 0xcb, 0xf8, 0x27, 0xdc, 0xca, 0xf2, 0x81, 0xcb, 0x42, 0xcf, 0xff, 0x6f, 0x3e,
	0xcf, 0xff, 0x85, 0xdd, 0xb6, 0xf3, 0x04, 0xc7, 0x7d, 0x1b, 0x1d, 0xc1, 0xe1, 0x6b, 0xa2, 0xa7,
	0xc0, 0x03, 0xfe, 0x7e, 0xd8, 0x48, 0x74, 0x09, 0x77, 0x2a, 0xba, 0x56, 0xc9, 0x74, 0xe0, 0x01,
	0x7f, 0x14, 0x4c, 0xb0, 0x7d, 0x09, 0x77, 0x2f, 0xe1, 0x5b, 0xae, 0x43, 0x1b, 0xb9, 0x19, 0x5c,
	0x83, 0xd9, 0x3b, 0x80, 0x87, 0xa1, 0x50, 0x32, 0xe9, 0xd5, 0x7d, 0x86, 0x93, 0x52, 0x52, 0xc9,
	0x96, 0x51, 0xd1, 0x38, 0xd1, 0x76, 0x77, 0xaf, 0xeb, 0x9e, 0xb3, 0xa6, 0x79, 0x6f, 0x59, 0x15,
	0x54, 0x32, 0xc1, 0xdb, 0xd2, 0xc8, 0x32, 0x7a, 0x7e, 0x89, 0x5e, 0xe0, 0x49, 0xac, 0x39, 0xcd,
	0x7e, 0xa1, 0x07, 0x7f, 0x42, 0x1f, 0xb7, 0x90, 0x3e, 0x7b, 0x3e, 0xde, 0xd4, 0x2e, 0xf8, 0xa8,
	0x5d, 0xf0, 0x55, 0xbb, 0x60, 0xb1, 0x6b, 0x1e, 0xbe, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x2a,
	0x10, 0xc8, 0xf0, 0x05, 0x02, 0x00, 0x00,
}
