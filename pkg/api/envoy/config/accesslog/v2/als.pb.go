// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/accesslog/v2/als.proto

package envoy_config_accesslog_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Configuration for the built-in *envoy.http_grpc_access_log*
// :ref:`AccessLog <envoy_api_msg_config.filter.accesslog.v2.AccessLog>`. This configuration will
// populate :ref:`StreamAccessLogsMessage.http_logs
// <envoy_api_field_service.accesslog.v2.StreamAccessLogsMessage.http_logs>`.
// [#extension: envoy.access_loggers.http_grpc]
type HttpGrpcAccessLogConfig struct {
	CommonConfig *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// Additional request headers to log in :ref:`HTTPRequestProperties.request_headers
	// <envoy_api_field_data.accesslog.v2.HTTPRequestProperties.request_headers>`.
	AdditionalRequestHeadersToLog []string `protobuf:"bytes,2,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	// Additional response headers to log in :ref:`HTTPResponseProperties.response_headers
	// <envoy_api_field_data.accesslog.v2.HTTPResponseProperties.response_headers>`.
	AdditionalResponseHeadersToLog []string `protobuf:"bytes,3,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	// Additional response trailers to log in :ref:`HTTPResponseProperties.response_trailers
	// <envoy_api_field_data.accesslog.v2.HTTPResponseProperties.response_trailers>`.
	AdditionalResponseTrailersToLog []string `protobuf:"bytes,4,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *HttpGrpcAccessLogConfig) Reset()         { *m = HttpGrpcAccessLogConfig{} }
func (m *HttpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*HttpGrpcAccessLogConfig) ProtoMessage()    {}
func (*HttpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7b431652a309a2e, []int{0}
}
func (m *HttpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpGrpcAccessLogConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpGrpcAccessLogConfig.Merge(m, src)
}
func (m *HttpGrpcAccessLogConfig) XXX_Size() int {
	return m.Size()
}
func (m *HttpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *HttpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

// Configuration for the built-in *envoy.tcp_grpc_access_log* type. This configuration will
// populate *StreamAccessLogsMessage.tcp_logs*.
// [#extension: envoy.access_loggers.tcp_grpc]
type TcpGrpcAccessLogConfig struct {
	CommonConfig         *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TcpGrpcAccessLogConfig) Reset()         { *m = TcpGrpcAccessLogConfig{} }
func (m *TcpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*TcpGrpcAccessLogConfig) ProtoMessage()    {}
func (*TcpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7b431652a309a2e, []int{1}
}
func (m *TcpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TcpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TcpGrpcAccessLogConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TcpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpGrpcAccessLogConfig.Merge(m, src)
}
func (m *TcpGrpcAccessLogConfig) XXX_Size() int {
	return m.Size()
}
func (m *TcpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *TcpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

// Common configuration for gRPC access logs.
// [#next-free-field: 6]
type CommonGrpcAccessLogConfig struct {
	// The friendly name of the access log to be returned in :ref:`StreamAccessLogsMessage.Identifier
	// <envoy_api_msg_service.accesslog.v2.StreamAccessLogsMessage.Identifier>`. This allows the
	// access log server to differentiate between different access logs coming from the same Envoy.
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// The gRPC service for the access log service.
	GrpcService *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	// Interval for flushing access logs to the gRPC stream. Logger will flush requests every time
	// this interval is elapsed, or when batch size limit is hit, whichever comes first. Defaults to
	// 1 second.
	BufferFlushInterval *types.Duration `protobuf:"bytes,3,opt,name=buffer_flush_interval,json=bufferFlushInterval,proto3" json:"buffer_flush_interval,omitempty"`
	// Soft size limit in bytes for access log entries buffer. Logger will buffer requests until
	// this limit it hit, or every time flush interval is elapsed, whichever comes first. Setting it
	// to zero effectively disables the batching. Defaults to 16384.
	BufferSizeBytes *types.UInt32Value `protobuf:"bytes,4,opt,name=buffer_size_bytes,json=bufferSizeBytes,proto3" json:"buffer_size_bytes,omitempty"`
	// Additional filter state objects to log in :ref:`filter_state_objects
	// <envoy_api_field_data.accesslog.v2.AccessLogCommon.filter_state_objects>`.
	// Logger will call `FilterState::Object::serializeAsProto` to serialize the filter state object.
	FilterStateObjectsToLog []string `protobuf:"bytes,5,rep,name=filter_state_objects_to_log,json=filterStateObjectsToLog,proto3" json:"filter_state_objects_to_log,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *CommonGrpcAccessLogConfig) Reset()         { *m = CommonGrpcAccessLogConfig{} }
func (m *CommonGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*CommonGrpcAccessLogConfig) ProtoMessage()    {}
func (*CommonGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7b431652a309a2e, []int{2}
}
func (m *CommonGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommonGrpcAccessLogConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommonGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonGrpcAccessLogConfig.Merge(m, src)
}
func (m *CommonGrpcAccessLogConfig) XXX_Size() int {
	return m.Size()
}
func (m *CommonGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonGrpcAccessLogConfig proto.InternalMessageInfo

func (m *CommonGrpcAccessLogConfig) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *CommonGrpcAccessLogConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferFlushInterval() *types.Duration {
	if m != nil {
		return m.BufferFlushInterval
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferSizeBytes() *types.UInt32Value {
	if m != nil {
		return m.BufferSizeBytes
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetFilterStateObjectsToLog() []string {
	if m != nil {
		return m.FilterStateObjectsToLog
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.HttpGrpcAccessLogConfig")
	proto.RegisterType((*TcpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.TcpGrpcAccessLogConfig")
	proto.RegisterType((*CommonGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.CommonGrpcAccessLogConfig")
}

func init() {
	proto.RegisterFile("envoy/config/accesslog/v2/als.proto", fileDescriptor_e7b431652a309a2e)
}

var fileDescriptor_e7b431652a309a2e = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xcf, 0x4e, 0x14, 0x31,
	0x18, 0xc0, 0x9d, 0x01, 0x04, 0x0a, 0x46, 0x1d, 0xa3, 0x2c, 0xa8, 0x2b, 0x2e, 0x26, 0xa0, 0x87,
	0x99, 0x64, 0xf1, 0xe8, 0x85, 0xc1, 0xe8, 0xa2, 0x44, 0xc9, 0x82, 0x7a, 0x9c, 0x74, 0x67, 0xbe,
	0x29, 0x35, 0xdd, 0x69, 0x6d, 0x3b, 0x23, 0x90, 0xf8, 0x02, 0x5e, 0xbc, 0xfa, 0x0c, 0x3e, 0x89,
	0x47, 0x1f, 0xc1, 0xec, 0xdd, 0x8b, 0x47, 0x0e, 0xc6, 0xb4, 0x9d, 0x85, 0x45, 0xd8, 0xb3, 0xb7,
	0xdd, 0xe9, 0xef, 0xfb, 0xf5, 0xfb, 0x57, 0xb4, 0x02, 0x45, 0xc5, 0x0f, 0xa3, 0x94, 0x17, 0x39,
	0x25, 0x11, 0x4e, 0x53, 0x50, 0x8a, 0x71, 0x12, 0x55, 0xed, 0x08, 0x33, 0x15, 0x0a, 0xc9, 0x35,
	0x0f, 0x16, 0x2d, 0x14, 0x3a, 0x28, 0x3c, 0x81, 0xc2, 0xaa, 0xbd, 0xf4, 0xc0, 0xc5, 0x63, 0x41,
	0x4d, 0x48, 0xca, 0x25, 0x44, 0x44, 0x8a, 0x34, 0x51, 0x20, 0x2b, 0x9a, 0x82, 0x13, 0x2c, 0x35,
	0x09, 0xe7, 0x84, 0x41, 0x64, 0xff, 0xf5, 0xca, 0x3c, 0xca, 0x4a, 0x89, 0x35, 0xe5, 0xc5, 0xb8,
	0xf3, 0x8f, 0x12, 0x0b, 0x01, 0x52, 0x0d, 0xcf, 0xcb, 0x4c, 0xe0, 0x08, 0x17, 0x05, 0xd7, 0x36,
	0x4c, 0x45, 0x7d, 0x4a, 0x24, 0xd6, 0x43, 0xff, 0x42, 0x85, 0x19, 0xcd, 0xb0, 0x86, 0x68, 0xf8,
	0xc3, 0x1d, 0xb4, 0x06, 0x3e, 0x5a, 0xe8, 0x68, 0x2d, 0x9e, 0x4b, 0x91, 0x6e, 0xd8, 0xbc, 0xb7,
	0x39, 0xd9, 0xb4, 0x75, 0x04, 0x29, 0xba, 0x92, 0xf2, 0x7e, 0x9f, 0x17, 0x89, 0x2b, 0xac, 0xe1,
	0x2d, 0x7b, 0x6b, 0x73, 0xed, 0xc7, 0xe1, 0xd8, 0x6a, 0xc3, 0x4d, 0xcb, 0x5f, 0x20, 0x8b, 0x67,
	0x8e, 0xe3, 0xa9, 0xcf, 0x9e, 0x7f, 0xcd, 0xeb, 0xce, 0x3b, 0x69, 0x7d, 0x49, 0x07, 0xdd, 0xc7,
	0x59, 0x46, 0x4d, 0xd2, 0x98, 0x25, 0x12, 0x3e, 0x94, 0xa0, 0x74, 0xb2, 0x0f, 0x38, 0x03, 0xa9,
	0x12, 0xcd, 0x13, 0xc6, 0x49, 0xc3, 0x5f, 0x9e, 0x58, 0x9b, 0xed, 0xde, 0x3d, 0x05, 0xbb, 0x8e,
	0xeb, 0x38, 0x6c, 0x8f, 0x6f, 0x73, 0x12, 0xbc, 0x40, 0xad, 0x33, 0x26, 0x25, 0x78, 0xa1, 0xe0,
	0x5f, 0xd5, 0x84, 0x55, 0x35, 0x47, 0x55, 0x0e, 0x3c, 0xe3, 0xda, 0x46, 0x2b, 0x17, 0xb9, 0xb4,
	0xc4, 0x94, 0x8d, 0xc8, 0x26, 0xad, 0xec, 0xde, 0x79, 0xd9, 0x5e, 0x0d, 0x5a, 0x5b, 0xeb, 0x13,
	0xba, 0xb5, 0x97, 0xfe, 0xb7, 0x16, 0xb7, 0x7e, 0xf9, 0x68, 0x71, 0x6c, 0x54, 0xd0, 0x42, 0x33,
	0x8c, 0x93, 0xa4, 0xc0, 0x7d, 0xb0, 0xb7, 0xcf, 0xc6, 0xd3, 0xc7, 0xf1, 0xa4, 0xf4, 0x97, 0xbd,
	0xee, 0x34, 0xe3, 0xe4, 0x15, 0xee, 0x43, 0xf0, 0x12, 0xcd, 0x8f, 0x2e, 0x6d, 0xc3, 0xb7, 0x59,
	0x36, 0xeb, 0x2c, 0xb1, 0xa0, 0x26, 0x31, 0xb3, 0xdb, 0xa1, 0xb9, 0x61, 0xd7, 0x51, 0x23, 0xf9,
	0xcc, 0x91, 0xd3, 0xcf, 0xc1, 0x3b, 0x74, 0xb3, 0x57, 0xe6, 0x39, 0xc8, 0x24, 0x67, 0xa5, 0xda,
	0x4f, 0x68, 0xa1, 0x41, 0x56, 0x98, 0x35, 0x26, 0xac, 0x75, 0x31, 0x74, 0xbb, 0x1e, 0x0e, 0x77,
	0x3d, 0x7c, 0x5a, 0xbf, 0x05, 0x2b, 0xfc, 0xe6, 0xf9, 0x8f, 0x2e, 0x75, 0x6f, 0x38, 0xc3, 0x33,
	0x23, 0xd8, 0xaa, 0xe3, 0x83, 0x0e, 0xba, 0x5e, 0x8b, 0x15, 0x3d, 0x82, 0xa4, 0x77, 0xa8, 0x41,
	0x35, 0x26, 0xad, 0xf4, 0xce, 0x39, 0xe9, 0x9b, 0xad, 0x42, 0xaf, 0xb7, 0xdf, 0x62, 0x56, 0x42,
	0xf7, 0xaa, 0x0b, 0xdb, 0xa5, 0x47, 0x10, 0x9b, 0xa0, 0xe0, 0x09, 0xba, 0x9d, 0x53, 0xa6, 0x8d,
	0x49, 0x63, 0x0d, 0x09, 0xef, 0xbd, 0x87, 0x54, 0x9f, 0x8c, 0x7d, 0xca, 0x8e, 0x7d, 0xc1, 0x21,
	0xbb, 0x86, 0x78, 0xed, 0x00, 0x3b, 0xee, 0x58, 0x7c, 0x1f, 0x34, 0xbd, 0x1f, 0x83, 0xa6, 0xf7,
	0x73, 0xd0, 0xf4, 0x7e, 0x7f, 0xfd, 0xf3, 0x65, 0xea, 0x61, 0xb0, 0xea, 0x5a, 0x05, 0x07, 0x1a,
	0x0a, 0x65, 0x1e, 0x68, 0x3d, 0x54, 0xe3, 0x23, 0xe6, 0x1d, 0x9b, 0x0e, 0x85, 0xd5, 0x3a, 0x5a,
	0xa5, 0xdc, 0xb5, 0x55, 0x48, 0x7e, 0x70, 0x38, 0x7e, 0x0f, 0xe2, 0x99, 0x0d, 0xa6, 0x76, 0x4c,
	0x35, 0x3b, 0x5e, 0xef, 0xb2, 0x2d, 0x6b, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xda,
	0x61, 0xc5, 0xad, 0x04, 0x00, 0x00,
}

func (m *HttpGrpcAccessLogConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpGrpcAccessLogConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpGrpcAccessLogConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.AdditionalResponseTrailersToLog) > 0 {
		for iNdEx := len(m.AdditionalResponseTrailersToLog) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AdditionalResponseTrailersToLog[iNdEx])
			copy(dAtA[i:], m.AdditionalResponseTrailersToLog[iNdEx])
			i = encodeVarintAls(dAtA, i, uint64(len(m.AdditionalResponseTrailersToLog[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AdditionalResponseHeadersToLog) > 0 {
		for iNdEx := len(m.AdditionalResponseHeadersToLog) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AdditionalResponseHeadersToLog[iNdEx])
			copy(dAtA[i:], m.AdditionalResponseHeadersToLog[iNdEx])
			i = encodeVarintAls(dAtA, i, uint64(len(m.AdditionalResponseHeadersToLog[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AdditionalRequestHeadersToLog) > 0 {
		for iNdEx := len(m.AdditionalRequestHeadersToLog) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AdditionalRequestHeadersToLog[iNdEx])
			copy(dAtA[i:], m.AdditionalRequestHeadersToLog[iNdEx])
			i = encodeVarintAls(dAtA, i, uint64(len(m.AdditionalRequestHeadersToLog[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CommonConfig != nil {
		{
			size, err := m.CommonConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAls(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TcpGrpcAccessLogConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TcpGrpcAccessLogConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TcpGrpcAccessLogConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CommonConfig != nil {
		{
			size, err := m.CommonConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAls(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommonGrpcAccessLogConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonGrpcAccessLogConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonGrpcAccessLogConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.FilterStateObjectsToLog) > 0 {
		for iNdEx := len(m.FilterStateObjectsToLog) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FilterStateObjectsToLog[iNdEx])
			copy(dAtA[i:], m.FilterStateObjectsToLog[iNdEx])
			i = encodeVarintAls(dAtA, i, uint64(len(m.FilterStateObjectsToLog[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.BufferSizeBytes != nil {
		{
			size, err := m.BufferSizeBytes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAls(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.BufferFlushInterval != nil {
		{
			size, err := m.BufferFlushInterval.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAls(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.GrpcService != nil {
		{
			size, err := m.GrpcService.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAls(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.LogName) > 0 {
		i -= len(m.LogName)
		copy(dAtA[i:], m.LogName)
		i = encodeVarintAls(dAtA, i, uint64(len(m.LogName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAls(dAtA []byte, offset int, v uint64) int {
	offset -= sovAls(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HttpGrpcAccessLogConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommonConfig != nil {
		l = m.CommonConfig.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if len(m.AdditionalRequestHeadersToLog) > 0 {
		for _, s := range m.AdditionalRequestHeadersToLog {
			l = len(s)
			n += 1 + l + sovAls(uint64(l))
		}
	}
	if len(m.AdditionalResponseHeadersToLog) > 0 {
		for _, s := range m.AdditionalResponseHeadersToLog {
			l = len(s)
			n += 1 + l + sovAls(uint64(l))
		}
	}
	if len(m.AdditionalResponseTrailersToLog) > 0 {
		for _, s := range m.AdditionalResponseTrailersToLog {
			l = len(s)
			n += 1 + l + sovAls(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TcpGrpcAccessLogConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommonConfig != nil {
		l = m.CommonConfig.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CommonGrpcAccessLogConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LogName)
	if l > 0 {
		n += 1 + l + sovAls(uint64(l))
	}
	if m.GrpcService != nil {
		l = m.GrpcService.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if m.BufferFlushInterval != nil {
		l = m.BufferFlushInterval.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if m.BufferSizeBytes != nil {
		l = m.BufferSizeBytes.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if len(m.FilterStateObjectsToLog) > 0 {
		for _, s := range m.FilterStateObjectsToLog {
			l = len(s)
			n += 1 + l + sovAls(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAls(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAls(x uint64) (n int) {
	return sovAls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HttpGrpcAccessLogConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
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
			return fmt.Errorf("proto: HttpGrpcAccessLogConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpGrpcAccessLogConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CommonConfig == nil {
				m.CommonConfig = &CommonGrpcAccessLogConfig{}
			}
			if err := m.CommonConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalRequestHeadersToLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalRequestHeadersToLog = append(m.AdditionalRequestHeadersToLog, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalResponseHeadersToLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalResponseHeadersToLog = append(m.AdditionalResponseHeadersToLog, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalResponseTrailersToLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalResponseTrailersToLog = append(m.AdditionalResponseTrailersToLog, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAls
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
func (m *TcpGrpcAccessLogConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
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
			return fmt.Errorf("proto: TcpGrpcAccessLogConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TcpGrpcAccessLogConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CommonConfig == nil {
				m.CommonConfig = &CommonGrpcAccessLogConfig{}
			}
			if err := m.CommonConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAls
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
func (m *CommonGrpcAccessLogConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
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
			return fmt.Errorf("proto: CommonGrpcAccessLogConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonGrpcAccessLogConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrpcService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GrpcService == nil {
				m.GrpcService = &core.GrpcService{}
			}
			if err := m.GrpcService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferFlushInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BufferFlushInterval == nil {
				m.BufferFlushInterval = &types.Duration{}
			}
			if err := m.BufferFlushInterval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferSizeBytes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BufferSizeBytes == nil {
				m.BufferSizeBytes = &types.UInt32Value{}
			}
			if err := m.BufferSizeBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilterStateObjectsToLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilterStateObjectsToLog = append(m.FilterStateObjectsToLog, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAls
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
func skipAls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAls
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
					return 0, ErrIntOverflowAls
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
					return 0, ErrIntOverflowAls
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
				return 0, ErrInvalidLengthAls
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAls
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAls
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAls        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAls          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAls = fmt.Errorf("proto: unexpected end of group")
)
