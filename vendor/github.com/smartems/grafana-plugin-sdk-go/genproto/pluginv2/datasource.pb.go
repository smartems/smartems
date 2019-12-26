// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datasource.proto

package pluginv2

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

type DatasourceRequest struct {
	TimeRange            *TimeRange         `protobuf:"bytes,1,opt,name=timeRange,proto3" json:"timeRange,omitempty"`
	Datasource           *DatasourceInfo    `protobuf:"bytes,2,opt,name=datasource,proto3" json:"datasource,omitempty"`
	Queries              []*DatasourceQuery `protobuf:"bytes,3,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DatasourceRequest) Reset()         { *m = DatasourceRequest{} }
func (m *DatasourceRequest) String() string { return proto.CompactTextString(m) }
func (*DatasourceRequest) ProtoMessage()    {}
func (*DatasourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{0}
}

func (m *DatasourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasourceRequest.Unmarshal(m, b)
}
func (m *DatasourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasourceRequest.Marshal(b, m, deterministic)
}
func (m *DatasourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasourceRequest.Merge(m, src)
}
func (m *DatasourceRequest) XXX_Size() int {
	return xxx_messageInfo_DatasourceRequest.Size(m)
}
func (m *DatasourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DatasourceRequest proto.InternalMessageInfo

func (m *DatasourceRequest) GetTimeRange() *TimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

func (m *DatasourceRequest) GetDatasource() *DatasourceInfo {
	if m != nil {
		return m.Datasource
	}
	return nil
}

func (m *DatasourceRequest) GetQueries() []*DatasourceQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

type DatasourceQuery struct {
	RefId                string   `protobuf:"bytes,1,opt,name=refId,proto3" json:"refId,omitempty"`
	MaxDataPoints        int64    `protobuf:"varint,2,opt,name=maxDataPoints,proto3" json:"maxDataPoints,omitempty"`
	IntervalMs           int64    `protobuf:"varint,3,opt,name=intervalMs,proto3" json:"intervalMs,omitempty"`
	ModelJson            string   `protobuf:"bytes,4,opt,name=modelJson,proto3" json:"modelJson,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasourceQuery) Reset()         { *m = DatasourceQuery{} }
func (m *DatasourceQuery) String() string { return proto.CompactTextString(m) }
func (*DatasourceQuery) ProtoMessage()    {}
func (*DatasourceQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{1}
}

func (m *DatasourceQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasourceQuery.Unmarshal(m, b)
}
func (m *DatasourceQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasourceQuery.Marshal(b, m, deterministic)
}
func (m *DatasourceQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasourceQuery.Merge(m, src)
}
func (m *DatasourceQuery) XXX_Size() int {
	return xxx_messageInfo_DatasourceQuery.Size(m)
}
func (m *DatasourceQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasourceQuery.DiscardUnknown(m)
}

var xxx_messageInfo_DatasourceQuery proto.InternalMessageInfo

func (m *DatasourceQuery) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *DatasourceQuery) GetMaxDataPoints() int64 {
	if m != nil {
		return m.MaxDataPoints
	}
	return 0
}

func (m *DatasourceQuery) GetIntervalMs() int64 {
	if m != nil {
		return m.IntervalMs
	}
	return 0
}

func (m *DatasourceQuery) GetModelJson() string {
	if m != nil {
		return m.ModelJson
	}
	return ""
}

type DatasourceResponse struct {
	Results              []*DatasourceQueryResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *DatasourceResponse) Reset()         { *m = DatasourceResponse{} }
func (m *DatasourceResponse) String() string { return proto.CompactTextString(m) }
func (*DatasourceResponse) ProtoMessage()    {}
func (*DatasourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{2}
}

func (m *DatasourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasourceResponse.Unmarshal(m, b)
}
func (m *DatasourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasourceResponse.Marshal(b, m, deterministic)
}
func (m *DatasourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasourceResponse.Merge(m, src)
}
func (m *DatasourceResponse) XXX_Size() int {
	return xxx_messageInfo_DatasourceResponse.Size(m)
}
func (m *DatasourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DatasourceResponse proto.InternalMessageInfo

func (m *DatasourceResponse) GetResults() []*DatasourceQueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type DatasourceQueryResult struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	RefId                string   `protobuf:"bytes,2,opt,name=refId,proto3" json:"refId,omitempty"`
	MetaJson             string   `protobuf:"bytes,3,opt,name=metaJson,proto3" json:"metaJson,omitempty"`
	Dataframes           [][]byte `protobuf:"bytes,4,rep,name=dataframes,proto3" json:"dataframes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasourceQueryResult) Reset()         { *m = DatasourceQueryResult{} }
func (m *DatasourceQueryResult) String() string { return proto.CompactTextString(m) }
func (*DatasourceQueryResult) ProtoMessage()    {}
func (*DatasourceQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{3}
}

func (m *DatasourceQueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasourceQueryResult.Unmarshal(m, b)
}
func (m *DatasourceQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasourceQueryResult.Marshal(b, m, deterministic)
}
func (m *DatasourceQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasourceQueryResult.Merge(m, src)
}
func (m *DatasourceQueryResult) XXX_Size() int {
	return xxx_messageInfo_DatasourceQueryResult.Size(m)
}
func (m *DatasourceQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasourceQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_DatasourceQueryResult proto.InternalMessageInfo

func (m *DatasourceQueryResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DatasourceQueryResult) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *DatasourceQueryResult) GetMetaJson() string {
	if m != nil {
		return m.MetaJson
	}
	return ""
}

func (m *DatasourceQueryResult) GetDataframes() [][]byte {
	if m != nil {
		return m.Dataframes
	}
	return nil
}

func init() {
	proto.RegisterType((*DatasourceRequest)(nil), "pluginv2.DatasourceRequest")
	proto.RegisterType((*DatasourceQuery)(nil), "pluginv2.DatasourceQuery")
	proto.RegisterType((*DatasourceResponse)(nil), "pluginv2.DatasourceResponse")
	proto.RegisterType((*DatasourceQueryResult)(nil), "pluginv2.DatasourceQueryResult")
}

func init() { proto.RegisterFile("datasource.proto", fileDescriptor_bb096a9d85d590d2) }

var fileDescriptor_bb096a9d85d590d2 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0x33, 0x06, 0x02, 0x0f, 0x8c, 0x58, 0x35, 0x99, 0x48, 0x94, 0x2c, 0x1e, 0x38, 0x91,
	0x38, 0x2e, 0x7a, 0x35, 0x5e, 0x30, 0x31, 0x62, 0x63, 0xbc, 0x57, 0x78, 0x90, 0x25, 0x6b, 0x0b,
	0x6d, 0x47, 0xf4, 0xe4, 0x07, 0xf0, 0xf3, 0xf8, 0xfd, 0xcc, 0x3a, 0x47, 0xa7, 0x99, 0xc7, 0xf7,
	0x7f, 0xbf, 0xd7, 0xfe, 0xff, 0xaf, 0x85, 0xde, 0x82, 0x19, 0xa6, 0x65, 0xaa, 0xe6, 0x38, 0x5e,
	0x2b, 0x69, 0x24, 0x69, 0xad, 0x93, 0x74, 0x15, 0x8b, 0x6d, 0xd4, 0xef, 0xce, 0x25, 0xe7, 0x52,
	0xe4, 0x7a, 0xf8, 0xe5, 0xc1, 0xe1, 0xdd, 0x0e, 0xa6, 0xb8, 0x49, 0x51, 0x1b, 0x72, 0x05, 0x6d,
	0x13, 0x73, 0xa4, 0x4c, 0xac, 0x30, 0xf0, 0x86, 0xde, 0xa8, 0x13, 0x1d, 0x8d, 0x8b, 0x13, 0xc6,
	0xcf, 0x45, 0x8b, 0x3a, 0x8a, 0x5c, 0x03, 0xb8, 0x4b, 0x83, 0x9a, 0x9d, 0x09, 0xdc, 0x8c, 0xbb,
	0x63, 0x2a, 0x96, 0x92, 0x96, 0x58, 0x32, 0x81, 0xe6, 0x26, 0x45, 0x15, 0xa3, 0x0e, 0xfc, 0xa1,
	0x3f, 0xea, 0x44, 0xa7, 0x55, 0x63, 0x4f, 0x29, 0xaa, 0x77, 0x5a, 0x90, 0xe1, 0xa7, 0x07, 0x07,
	0x7f, 0x9a, 0xe4, 0x18, 0x1a, 0x0a, 0x97, 0xd3, 0x85, 0x75, 0xdc, 0xa6, 0x79, 0x41, 0x2e, 0x61,
	0x9f, 0xb3, 0xb7, 0x8c, 0x9d, 0xc9, 0x58, 0x18, 0x6d, 0xbd, 0xf9, 0xf4, 0xb7, 0x48, 0xce, 0x01,
	0x62, 0x61, 0x50, 0x6d, 0x59, 0xf2, 0x90, 0xf9, 0xc8, 0x90, 0x92, 0x42, 0x06, 0xd0, 0xe6, 0x72,
	0x81, 0xc9, 0xbd, 0x96, 0x22, 0xa8, 0xdb, 0xf3, 0x9d, 0x10, 0x3e, 0x02, 0x29, 0x2f, 0x51, 0xaf,
	0xa5, 0xd0, 0x48, 0x6e, 0xa0, 0xa9, 0x50, 0xa7, 0x89, 0xd1, 0x81, 0x67, 0x83, 0x5d, 0xfc, 0x1f,
	0xcc, 0x72, 0xb4, 0xe0, 0xc3, 0x0f, 0x38, 0xa9, 0x24, 0xb2, 0x8c, 0xa8, 0x94, 0x54, 0x45, 0x46,
	0x5b, 0xb8, 0xe4, 0xb5, 0x72, 0xf2, 0x3e, 0xb4, 0x38, 0x1a, 0x66, 0x2d, 0xfb, 0xb6, 0xb1, 0xab,
	0xb3, 0xbc, 0xd9, 0x13, 0x2c, 0x15, 0xe3, 0xa8, 0x83, 0xfa, 0xd0, 0x1f, 0x75, 0x69, 0x49, 0x89,
	0x5e, 0xa0, 0xe7, 0x0c, 0xcc, 0xac, 0x6b, 0x72, 0x0b, 0x8d, 0x7c, 0xd1, 0x67, 0x55, 0x39, 0x7e,
	0xfe, 0x4e, 0x7f, 0x50, 0xdd, 0xcc, 0x77, 0xf2, 0xba, 0x67, 0xbf, 0xdd, 0xe4, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x64, 0x89, 0x5d, 0xe6, 0xa2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatasourcePluginClient is the client API for DatasourcePlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatasourcePluginClient interface {
	Query(ctx context.Context, in *DatasourceRequest, opts ...grpc.CallOption) (*DatasourceResponse, error)
}

type datasourcePluginClient struct {
	cc *grpc.ClientConn
}

func NewDatasourcePluginClient(cc *grpc.ClientConn) DatasourcePluginClient {
	return &datasourcePluginClient{cc}
}

func (c *datasourcePluginClient) Query(ctx context.Context, in *DatasourceRequest, opts ...grpc.CallOption) (*DatasourceResponse, error) {
	out := new(DatasourceResponse)
	err := c.cc.Invoke(ctx, "/pluginv2.DatasourcePlugin/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasourcePluginServer is the server API for DatasourcePlugin service.
type DatasourcePluginServer interface {
	Query(context.Context, *DatasourceRequest) (*DatasourceResponse, error)
}

// UnimplementedDatasourcePluginServer can be embedded to have forward compatible implementations.
type UnimplementedDatasourcePluginServer struct {
}

func (*UnimplementedDatasourcePluginServer) Query(ctx context.Context, req *DatasourceRequest) (*DatasourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterDatasourcePluginServer(s *grpc.Server, srv DatasourcePluginServer) {
	s.RegisterService(&_DatasourcePlugin_serviceDesc, srv)
}

func _DatasourcePlugin_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatasourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourcePluginServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pluginv2.DatasourcePlugin/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourcePluginServer).Query(ctx, req.(*DatasourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatasourcePlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pluginv2.DatasourcePlugin",
	HandlerType: (*DatasourcePluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _DatasourcePlugin_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datasource.proto",
}
