// Code generated by protoc-gen-go.
// source: anonstore.proto
// DO NOT EDIT!

/*
Package censys_definitions is a generated protocol buffer package.

It is generated from these files:
	anonstore.proto
	caa.proto
	certificate.proto
	common.proto
	ct.proto
	hoststore.proto
	protocols.proto
	pubkey.proto
	rpc.proto
	search.proto
	zlint.proto

It has these top-level messages:
	AnonymousRecord
	AnonymousDelta
	ExternalCertificate
	CAATagValue
	CAARecord
	CAALookup
	Path
	RootStoreStatus
	CertificateValidation
	MozillaSalesForceStatus
	CertificateRevocation
	CertificateAudit
	Certificate
	Metadatum
	UserdataAtom
	ASAtom
	CTServerStatus
	CTStatus
	SCT
	WHOISAtom
	LocationAtom
	ProtocolAtom
	AnonymousKey
	Record
	Delta
	RSACryptographicKey
	DSACryptographicKey
	ECCCryptographicKey
	CryptographicKey
	MinScanId
	MozillaOneCRLEntry
	Command
	AnonymousStoreStatistics
	StatisticsPair
	StoreStatistics
	ServerStatistics
	PruneStatistics
	CommandReply
	HostQuery
	HostQueryResponse
	AnonymousQuery
	AnonymousQueryResponse
	UserDataRequest
	RootStoreQuery
	RootStoreReply
	LintResult
	ZLint
	Lints
*/
package censys_definitions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AnonymousDelta_DeltaType int32

const (
	AnonymousDelta_DT_RESERVED AnonymousDelta_DeltaType = 0
	AnonymousDelta_DT_UPDATE   AnonymousDelta_DeltaType = 1
	AnonymousDelta_DT_DELETE   AnonymousDelta_DeltaType = 2
)

var AnonymousDelta_DeltaType_name = map[int32]string{
	0: "DT_RESERVED",
	1: "DT_UPDATE",
	2: "DT_DELETE",
}
var AnonymousDelta_DeltaType_value = map[string]int32{
	"DT_RESERVED": 0,
	"DT_UPDATE":   1,
	"DT_DELETE":   2,
}

func (x AnonymousDelta_DeltaType) String() string {
	return proto.EnumName(AnonymousDelta_DeltaType_name, int32(x))
}
func (AnonymousDelta_DeltaType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type AnonymousDelta_DeltaScope int32

const (
	AnonymousDelta_SCOPE_RESERVED  AnonymousDelta_DeltaScope = 0
	AnonymousDelta_SCOPE_NO_CHANGE AnonymousDelta_DeltaScope = 1
	AnonymousDelta_SCOPE_NEW       AnonymousDelta_DeltaScope = 2
	AnonymousDelta_SCOPE_UPDATE    AnonymousDelta_DeltaScope = 3
)

var AnonymousDelta_DeltaScope_name = map[int32]string{
	0: "SCOPE_RESERVED",
	1: "SCOPE_NO_CHANGE",
	2: "SCOPE_NEW",
	3: "SCOPE_UPDATE",
}
var AnonymousDelta_DeltaScope_value = map[string]int32{
	"SCOPE_RESERVED":  0,
	"SCOPE_NO_CHANGE": 1,
	"SCOPE_NEW":       2,
	"SCOPE_UPDATE":    3,
}

func (x AnonymousDelta_DeltaScope) String() string {
	return proto.EnumName(AnonymousDelta_DeltaScope_name, int32(x))
}
func (AnonymousDelta_DeltaScope) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

type AnonymousRecord struct {
	Sha256Fp  []byte        `protobuf:"bytes,1,opt,name=sha256fp,proto3" json:"sha256fp,omitempty"`
	Timestamp int64         `protobuf:"fixed64,2,opt,name=timestamp" json:"timestamp,omitempty"`
	ScanId    uint32        `protobuf:"varint,3,opt,name=scan_id,json=scanId" json:"scan_id,omitempty"`
	Exported  bool          `protobuf:"varint,4,opt,name=exported" json:"exported,omitempty"`
	Userdata  *UserdataAtom `protobuf:"bytes,5,opt,name=userdata" json:"userdata,omitempty"`
	// Types that are valid to be assigned to OneofData:
	//	*AnonymousRecord_Data
	//	*AnonymousRecord_RawData
	//	*AnonymousRecord_Certificate
	//	*AnonymousRecord_Key
	//	*AnonymousRecord_As
	OneofData isAnonymousRecord_OneofData `protobuf_oneof:"oneof_data"`
	Metadata  []*Metadatum                `protobuf:"bytes,14,rep,name=metadata" json:"metadata,omitempty"`
	Tags      []string                    `protobuf:"bytes,15,rep,name=tags" json:"tags,omitempty"`
	UpdatedAt uint32                      `protobuf:"fixed32,16,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	AddedAt   uint32                      `protobuf:"fixed32,17,opt,name=added_at,json=addedAt" json:"added_at,omitempty"`
}

func (m *AnonymousRecord) Reset()                    { *m = AnonymousRecord{} }
func (m *AnonymousRecord) String() string            { return proto.CompactTextString(m) }
func (*AnonymousRecord) ProtoMessage()               {}
func (*AnonymousRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isAnonymousRecord_OneofData interface {
	isAnonymousRecord_OneofData()
}

type AnonymousRecord_Data struct {
	Data string `protobuf:"bytes,6,opt,name=data,oneof"`
}
type AnonymousRecord_RawData struct {
	RawData []byte `protobuf:"bytes,7,opt,name=raw_data,json=rawData,proto3,oneof"`
}
type AnonymousRecord_Certificate struct {
	Certificate *Certificate `protobuf:"bytes,8,opt,name=certificate,oneof"`
}
type AnonymousRecord_Key struct {
	Key *CryptographicKey `protobuf:"bytes,9,opt,name=key,oneof"`
}
type AnonymousRecord_As struct {
	As *ASAtom `protobuf:"bytes,10,opt,name=as,oneof"`
}

func (*AnonymousRecord_Data) isAnonymousRecord_OneofData()        {}
func (*AnonymousRecord_RawData) isAnonymousRecord_OneofData()     {}
func (*AnonymousRecord_Certificate) isAnonymousRecord_OneofData() {}
func (*AnonymousRecord_Key) isAnonymousRecord_OneofData()         {}
func (*AnonymousRecord_As) isAnonymousRecord_OneofData()          {}

func (m *AnonymousRecord) GetOneofData() isAnonymousRecord_OneofData {
	if m != nil {
		return m.OneofData
	}
	return nil
}

func (m *AnonymousRecord) GetSha256Fp() []byte {
	if m != nil {
		return m.Sha256Fp
	}
	return nil
}

func (m *AnonymousRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AnonymousRecord) GetScanId() uint32 {
	if m != nil {
		return m.ScanId
	}
	return 0
}

func (m *AnonymousRecord) GetExported() bool {
	if m != nil {
		return m.Exported
	}
	return false
}

func (m *AnonymousRecord) GetUserdata() *UserdataAtom {
	if m != nil {
		return m.Userdata
	}
	return nil
}

func (m *AnonymousRecord) GetData() string {
	if x, ok := m.GetOneofData().(*AnonymousRecord_Data); ok {
		return x.Data
	}
	return ""
}

func (m *AnonymousRecord) GetRawData() []byte {
	if x, ok := m.GetOneofData().(*AnonymousRecord_RawData); ok {
		return x.RawData
	}
	return nil
}

func (m *AnonymousRecord) GetCertificate() *Certificate {
	if x, ok := m.GetOneofData().(*AnonymousRecord_Certificate); ok {
		return x.Certificate
	}
	return nil
}

func (m *AnonymousRecord) GetKey() *CryptographicKey {
	if x, ok := m.GetOneofData().(*AnonymousRecord_Key); ok {
		return x.Key
	}
	return nil
}

func (m *AnonymousRecord) GetAs() *ASAtom {
	if x, ok := m.GetOneofData().(*AnonymousRecord_As); ok {
		return x.As
	}
	return nil
}

func (m *AnonymousRecord) GetMetadata() []*Metadatum {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *AnonymousRecord) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *AnonymousRecord) GetUpdatedAt() uint32 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *AnonymousRecord) GetAddedAt() uint32 {
	if m != nil {
		return m.AddedAt
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AnonymousRecord) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AnonymousRecord_OneofMarshaler, _AnonymousRecord_OneofUnmarshaler, _AnonymousRecord_OneofSizer, []interface{}{
		(*AnonymousRecord_Data)(nil),
		(*AnonymousRecord_RawData)(nil),
		(*AnonymousRecord_Certificate)(nil),
		(*AnonymousRecord_Key)(nil),
		(*AnonymousRecord_As)(nil),
	}
}

func _AnonymousRecord_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AnonymousRecord)
	// oneof_data
	switch x := m.OneofData.(type) {
	case *AnonymousRecord_Data:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Data)
	case *AnonymousRecord_RawData:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RawData)
	case *AnonymousRecord_Certificate:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Certificate); err != nil {
			return err
		}
	case *AnonymousRecord_Key:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Key); err != nil {
			return err
		}
	case *AnonymousRecord_As:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.As); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AnonymousRecord.OneofData has unexpected type %T", x)
	}
	return nil
}

func _AnonymousRecord_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AnonymousRecord)
	switch tag {
	case 6: // oneof_data.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OneofData = &AnonymousRecord_Data{x}
		return true, err
	case 7: // oneof_data.raw_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.OneofData = &AnonymousRecord_RawData{x}
		return true, err
	case 8: // oneof_data.certificate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Certificate)
		err := b.DecodeMessage(msg)
		m.OneofData = &AnonymousRecord_Certificate{msg}
		return true, err
	case 9: // oneof_data.key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CryptographicKey)
		err := b.DecodeMessage(msg)
		m.OneofData = &AnonymousRecord_Key{msg}
		return true, err
	case 10: // oneof_data.as
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ASAtom)
		err := b.DecodeMessage(msg)
		m.OneofData = &AnonymousRecord_As{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AnonymousRecord_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AnonymousRecord)
	// oneof_data
	switch x := m.OneofData.(type) {
	case *AnonymousRecord_Data:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Data)))
		n += len(x.Data)
	case *AnonymousRecord_RawData:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RawData)))
		n += len(x.RawData)
	case *AnonymousRecord_Certificate:
		s := proto.Size(x.Certificate)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AnonymousRecord_Key:
		s := proto.Size(x.Key)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AnonymousRecord_As:
		s := proto.Size(x.As)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AnonymousDelta struct {
	DeltaType  AnonymousDelta_DeltaType  `protobuf:"varint,1,opt,name=delta_type,json=deltaType,enum=zsearch.AnonymousDelta_DeltaType" json:"delta_type,omitempty"`
	DeltaScope AnonymousDelta_DeltaScope `protobuf:"varint,2,opt,name=delta_scope,json=deltaScope,enum=zsearch.AnonymousDelta_DeltaScope" json:"delta_scope,omitempty"`
	Record     *AnonymousRecord          `protobuf:"bytes,3,opt,name=record" json:"record,omitempty"`
}

func (m *AnonymousDelta) Reset()                    { *m = AnonymousDelta{} }
func (m *AnonymousDelta) String() string            { return proto.CompactTextString(m) }
func (*AnonymousDelta) ProtoMessage()               {}
func (*AnonymousDelta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AnonymousDelta) GetDeltaType() AnonymousDelta_DeltaType {
	if m != nil {
		return m.DeltaType
	}
	return AnonymousDelta_DT_RESERVED
}

func (m *AnonymousDelta) GetDeltaScope() AnonymousDelta_DeltaScope {
	if m != nil {
		return m.DeltaScope
	}
	return AnonymousDelta_SCOPE_RESERVED
}

func (m *AnonymousDelta) GetRecord() *AnonymousRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type ExternalCertificate struct {
	Source          CertificateSource        `protobuf:"varint,1,opt,name=source,enum=zsearch.CertificateSource" json:"source,omitempty"`
	AnonymousRecord *AnonymousRecord         `protobuf:"bytes,2,opt,name=anonymous_record,json=anonymousRecord" json:"anonymous_record,omitempty"`
	CtServer        CTServer                 `protobuf:"varint,3,opt,name=ct_server,json=ctServer,enum=zsearch.CTServer" json:"ct_server,omitempty"`
	CtStatus        *CTServerStatus          `protobuf:"bytes,4,opt,name=ct_status,json=ctStatus" json:"ct_status,omitempty"`
	NssStatus       *MozillaSalesForceStatus `protobuf:"bytes,5,opt,name=nss_status,json=nssStatus" json:"nss_status,omitempty"`
	TbsHash         []byte                   `protobuf:"bytes,6,opt,name=tbsHash,proto3" json:"tbsHash,omitempty"`
}

func (m *ExternalCertificate) Reset()                    { *m = ExternalCertificate{} }
func (m *ExternalCertificate) String() string            { return proto.CompactTextString(m) }
func (*ExternalCertificate) ProtoMessage()               {}
func (*ExternalCertificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ExternalCertificate) GetSource() CertificateSource {
	if m != nil {
		return m.Source
	}
	return CertificateSource_CERTIFICATE_SOURCE_RESERVED
}

func (m *ExternalCertificate) GetAnonymousRecord() *AnonymousRecord {
	if m != nil {
		return m.AnonymousRecord
	}
	return nil
}

func (m *ExternalCertificate) GetCtServer() CTServer {
	if m != nil {
		return m.CtServer
	}
	return CTServer_CT_SERVER_RESERVED
}

func (m *ExternalCertificate) GetCtStatus() *CTServerStatus {
	if m != nil {
		return m.CtStatus
	}
	return nil
}

func (m *ExternalCertificate) GetNssStatus() *MozillaSalesForceStatus {
	if m != nil {
		return m.NssStatus
	}
	return nil
}

func (m *ExternalCertificate) GetTbsHash() []byte {
	if m != nil {
		return m.TbsHash
	}
	return nil
}

func init() {
	proto.RegisterType((*AnonymousRecord)(nil), "zsearch.AnonymousRecord")
	proto.RegisterType((*AnonymousDelta)(nil), "zsearch.AnonymousDelta")
	proto.RegisterType((*ExternalCertificate)(nil), "zsearch.ExternalCertificate")
	proto.RegisterEnum("zsearch.AnonymousDelta_DeltaType", AnonymousDelta_DeltaType_name, AnonymousDelta_DeltaType_value)
	proto.RegisterEnum("zsearch.AnonymousDelta_DeltaScope", AnonymousDelta_DeltaScope_name, AnonymousDelta_DeltaScope_value)
}

func init() { proto.RegisterFile("anonstore.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xda, 0x4a,
	0x10, 0xc5, 0x86, 0x0b, 0xf6, 0x40, 0xc0, 0xd9, 0xe4, 0x2a, 0x0e, 0xf7, 0x43, 0x0e, 0x4f, 0x7e,
	0xb9, 0xe8, 0xd6, 0xfd, 0x50, 0xd5, 0x97, 0x96, 0x80, 0x5b, 0xaa, 0xb6, 0x49, 0xb4, 0x26, 0xe9,
	0xa3, 0xb5, 0xb1, 0x37, 0x01, 0x05, 0x7b, 0xad, 0xdd, 0xa5, 0x09, 0x79, 0xec, 0x1f, 0xec, 0x2f,
	0xe8, 0x7f, 0xa9, 0x58, 0x1b, 0x43, 0x9b, 0xaa, 0x7d, 0xb1, 0xe6, 0xcc, 0x39, 0x33, 0xb3, 0x3b,
	0x3a, 0x6b, 0xe8, 0x90, 0x94, 0xa5, 0x42, 0x32, 0x4e, 0xfb, 0x19, 0x67, 0x92, 0xa1, 0xc6, 0xbd,
	0xa0, 0x84, 0x47, 0xd3, 0x6e, 0x2b, 0x62, 0x49, 0xc2, 0xd2, 0x3c, 0xdd, 0x6d, 0x65, 0x8b, 0xcb,
	0x1b, 0xba, 0x2c, 0x90, 0x11, 0xc9, 0x22, 0xda, 0x8d, 0x28, 0x97, 0xb3, 0xab, 0x59, 0x44, 0x64,
	0xd1, 0xa1, 0xf7, 0xb9, 0x06, 0x9d, 0x41, 0xca, 0xd2, 0x65, 0xc2, 0x16, 0x02, 0xd3, 0x88, 0xf1,
	0x18, 0x75, 0xc1, 0x10, 0x53, 0xe2, 0x3d, 0x7d, 0x76, 0x95, 0xd9, 0x9a, 0xa3, 0xb9, 0x2d, 0x5c,
	0x62, 0xf4, 0x37, 0x98, 0x72, 0x96, 0x50, 0x21, 0x49, 0x92, 0xd9, 0xba, 0xa3, 0xb9, 0x16, 0xde,
	0x24, 0xd0, 0x01, 0x34, 0x44, 0x44, 0xd2, 0x70, 0x16, 0xdb, 0x55, 0x47, 0x73, 0x77, 0x70, 0x7d,
	0x05, 0xdf, 0xc6, 0xe8, 0x5f, 0x30, 0xe8, 0x5d, 0xc6, 0xb8, 0xa4, 0xb1, 0x5d, 0x73, 0x34, 0xd7,
	0x38, 0xd6, 0x6d, 0x0d, 0x97, 0x39, 0xf4, 0x08, 0x8c, 0x85, 0xa0, 0x3c, 0x26, 0x92, 0xd8, 0x7f,
	0x38, 0x9a, 0xdb, 0xf4, 0xfe, 0xec, 0x17, 0x77, 0xeb, 0x9f, 0x17, 0xc4, 0x40, 0xb2, 0x04, 0x97,
	0x32, 0xb4, 0x0f, 0x35, 0x25, 0xaf, 0x3b, 0x9a, 0x6b, 0x8e, 0x2b, 0x58, 0x21, 0xf4, 0x17, 0x18,
	0x9c, 0xdc, 0x86, 0x8a, 0x69, 0xac, 0xce, 0x3e, 0xae, 0xe0, 0x06, 0x27, 0xb7, 0xa3, 0x15, 0xf9,
	0x1c, 0x9a, 0x5b, 0x1b, 0xb0, 0x0d, 0x35, 0x68, 0xbf, 0x1c, 0x34, 0xdc, 0x70, 0xe3, 0x0a, 0xde,
	0x96, 0xa2, 0xff, 0xa0, 0x7a, 0x43, 0x97, 0xb6, 0xa9, 0x2a, 0x0e, 0x37, 0x15, 0x7c, 0x99, 0x49,
	0x76, 0xcd, 0x49, 0x36, 0x9d, 0x45, 0xef, 0xe8, 0x72, 0x5c, 0xc1, 0x2b, 0x1d, 0x3a, 0x02, 0x9d,
	0x08, 0x1b, 0x94, 0xba, 0x53, 0xaa, 0x07, 0xc1, 0xea, 0x0a, 0xe3, 0x0a, 0xd6, 0x89, 0x40, 0x7d,
	0x30, 0x12, 0x2a, 0x89, 0x3a, 0x68, 0xdb, 0xa9, 0xba, 0x4d, 0x0f, 0x95, 0xc2, 0x0f, 0x39, 0xb1,
	0x48, 0x70, 0xa9, 0x41, 0x08, 0x6a, 0x92, 0x5c, 0x0b, 0xbb, 0xe3, 0x54, 0x5d, 0x13, 0xab, 0x18,
	0xfd, 0x03, 0xb0, 0xc8, 0x62, 0x22, 0x69, 0x1c, 0x12, 0x69, 0x5b, 0x8e, 0xe6, 0x36, 0xb0, 0x59,
	0x64, 0x06, 0x12, 0x1d, 0x82, 0x41, 0xe2, 0x38, 0x27, 0x77, 0x15, 0xd9, 0x50, 0x78, 0x20, 0x8f,
	0x5b, 0x00, 0x2c, 0xa5, 0xec, 0x4a, 0x2d, 0xaa, 0xf7, 0x55, 0x87, 0x76, 0x69, 0x82, 0x11, 0x9d,
	0x4b, 0x82, 0x5e, 0x01, 0xc4, 0xab, 0x20, 0x94, 0xcb, 0x8c, 0x2a, 0x17, 0xb4, 0xbd, 0xa3, 0xcd,
	0x4d, 0xbe, 0x13, 0xf7, 0xd5, 0x77, 0xb2, 0xcc, 0x28, 0x36, 0xe3, 0x75, 0x88, 0x86, 0xd0, 0xcc,
	0x3b, 0x88, 0x88, 0x65, 0x54, 0x79, 0xa5, 0xed, 0xf5, 0x7e, 0xd9, 0x22, 0x58, 0x29, 0x71, 0x3e,
	0x58, 0xc5, 0xe8, 0x7f, 0xa8, 0x73, 0x65, 0x4a, 0xe5, 0xa7, 0xa6, 0x67, 0x3f, 0xac, 0xcf, 0x4d,
	0x8b, 0x0b, 0x5d, 0xef, 0x05, 0x98, 0xe5, 0x71, 0x50, 0x07, 0x9a, 0xa3, 0x49, 0x88, 0xfd, 0xc0,
	0xc7, 0x17, 0xfe, 0xc8, 0xaa, 0xa0, 0x1d, 0x30, 0x47, 0x93, 0xf0, 0xfc, 0x6c, 0x34, 0x98, 0xf8,
	0x96, 0x56, 0xc0, 0x91, 0xff, 0xde, 0x9f, 0xf8, 0x96, 0xde, 0xbb, 0x00, 0xd8, 0x9c, 0x03, 0x21,
	0x68, 0x07, 0xc3, 0xd3, 0x33, 0x7f, 0xbb, 0x7e, 0x0f, 0x3a, 0x79, 0xee, 0xe4, 0x34, 0x1c, 0x8e,
	0x07, 0x27, 0x6f, 0x8a, 0x2e, 0x45, 0xd2, 0xff, 0x68, 0xe9, 0xc8, 0x82, 0x56, 0x0e, 0x8b, 0x31,
	0xd5, 0xde, 0x17, 0x1d, 0xf6, 0xfc, 0x3b, 0x49, 0x79, 0x4a, 0xe6, 0x5b, 0x26, 0x43, 0x1e, 0xd4,
	0x05, 0x5b, 0xf0, 0x68, 0xbd, 0xe0, 0xee, 0xcf, 0xac, 0x18, 0x28, 0x05, 0x2e, 0x94, 0x68, 0x08,
	0x16, 0x59, 0x5f, 0x3d, 0x2c, 0x76, 0xa3, 0xff, 0x66, 0x37, 0x1d, 0xf2, 0xc3, 0x0b, 0xef, 0x83,
	0x19, 0xc9, 0x50, 0x50, 0xfe, 0x89, 0x72, 0xb5, 0xd9, 0xb6, 0xb7, 0xbb, 0x99, 0x3d, 0x09, 0x14,
	0x81, 0x8d, 0x48, 0xe6, 0x11, 0x7a, 0x92, 0xeb, 0x25, 0x91, 0x0b, 0xa1, 0xde, 0x6f, 0xd3, 0x3b,
	0x78, 0xa0, 0x0f, 0x14, 0xad, 0xaa, 0x54, 0x84, 0x5e, 0x02, 0xa4, 0x42, 0xac, 0xcb, 0xf2, 0x67,
	0xed, 0x6c, 0x4c, 0xce, 0xee, 0x67, 0xf3, 0x39, 0x09, 0xc8, 0x9c, 0x8a, 0xd7, 0x8c, 0x47, 0xb4,
	0xa8, 0x37, 0x53, 0x21, 0x8a, 0x06, 0x36, 0x34, 0xe4, 0xa5, 0x18, 0x13, 0x31, 0x55, 0xaf, 0xbc,
	0x85, 0xd7, 0xf0, 0xb2, 0xae, 0xfe, 0x5e, 0x8f, 0xbf, 0x05, 0x00, 0x00, 0xff, 0xff, 0x07, 0xb6,
	0x5a, 0xca, 0x12, 0x05, 0x00, 0x00,
}
