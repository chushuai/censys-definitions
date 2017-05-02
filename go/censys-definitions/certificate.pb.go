// Code generated by protoc-gen-go.
// source: certificate.proto
// DO NOT EDIT!

package censys_definitions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CertificateType int32

const (
	CertificateType_CERTIFICATE_TYPE_RESERVED     CertificateType = 0
	CertificateType_CERTIFICATE_TYPE_UNKNOWN      CertificateType = 1
	CertificateType_CERTIFICATE_TYPE_LEAF         CertificateType = 2
	CertificateType_CERTIFICATE_TYPE_INTERMEDIATE CertificateType = 3
	CertificateType_CERTIFICATE_TYPE_ROOT         CertificateType = 4
)

var CertificateType_name = map[int32]string{
	0: "CERTIFICATE_TYPE_RESERVED",
	1: "CERTIFICATE_TYPE_UNKNOWN",
	2: "CERTIFICATE_TYPE_LEAF",
	3: "CERTIFICATE_TYPE_INTERMEDIATE",
	4: "CERTIFICATE_TYPE_ROOT",
}
var CertificateType_value = map[string]int32{
	"CERTIFICATE_TYPE_RESERVED":     0,
	"CERTIFICATE_TYPE_UNKNOWN":      1,
	"CERTIFICATE_TYPE_LEAF":         2,
	"CERTIFICATE_TYPE_INTERMEDIATE": 3,
	"CERTIFICATE_TYPE_ROOT":         4,
}

func (x CertificateType) String() string {
	return proto.EnumName(CertificateType_name, int32(x))
}
func (CertificateType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type CertificateSource int32

const (
	CertificateSource_CERTIFICATE_SOURCE_RESERVED           CertificateSource = 0
	CertificateSource_CERTIFICATE_SOURCE_UNKNOWN            CertificateSource = 1
	CertificateSource_CERTIFICATE_SOURCE_SCAN               CertificateSource = 2
	CertificateSource_CERTIFICATE_SOURCE_CT                 CertificateSource = 3
	CertificateSource_CERTIFICATE_SOURCE_MOZILLA_SALESFORCE CertificateSource = 4
	CertificateSource_CERTIFICATE_SOURCE_RESEARCH           CertificateSource = 5
	CertificateSource_CERTIFICATE_SOURCE_RAPID7             CertificateSource = 6
	CertificateSource_CERTIFICATE_SOURCE_HUBBLE             CertificateSource = 7
	CertificateSource_CERTIFICATE_SOURCE_CT_CHAIN           CertificateSource = 8
)

var CertificateSource_name = map[int32]string{
	0: "CERTIFICATE_SOURCE_RESERVED",
	1: "CERTIFICATE_SOURCE_UNKNOWN",
	2: "CERTIFICATE_SOURCE_SCAN",
	3: "CERTIFICATE_SOURCE_CT",
	4: "CERTIFICATE_SOURCE_MOZILLA_SALESFORCE",
	5: "CERTIFICATE_SOURCE_RESEARCH",
	6: "CERTIFICATE_SOURCE_RAPID7",
	7: "CERTIFICATE_SOURCE_HUBBLE",
	8: "CERTIFICATE_SOURCE_CT_CHAIN",
}
var CertificateSource_value = map[string]int32{
	"CERTIFICATE_SOURCE_RESERVED":           0,
	"CERTIFICATE_SOURCE_UNKNOWN":            1,
	"CERTIFICATE_SOURCE_SCAN":               2,
	"CERTIFICATE_SOURCE_CT":                 3,
	"CERTIFICATE_SOURCE_MOZILLA_SALESFORCE": 4,
	"CERTIFICATE_SOURCE_RESEARCH":           5,
	"CERTIFICATE_SOURCE_RAPID7":             6,
	"CERTIFICATE_SOURCE_HUBBLE":             7,
	"CERTIFICATE_SOURCE_CT_CHAIN":           8,
}

func (x CertificateSource) String() string {
	return proto.EnumName(CertificateSource_name, int32(x))
}
func (CertificateSource) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type CertificateParseStatus int32

const (
	CertificateParseStatus_CERTIFICATE_PARSE_STATUS_RESERVED   CertificateParseStatus = 0
	CertificateParseStatus_CERTIFICATE_PARSE_STATUS_UNKNOWN    CertificateParseStatus = 1
	CertificateParseStatus_CERTIFICATE_PARSE_STATUS_SUCCESS    CertificateParseStatus = 2
	CertificateParseStatus_CERTIFICATE_PARSE_STATUS_FAIL       CertificateParseStatus = 3
	CertificateParseStatus_CERTIFICATE_PARSE_STATUS_NOT_PARSED CertificateParseStatus = 4
)

var CertificateParseStatus_name = map[int32]string{
	0: "CERTIFICATE_PARSE_STATUS_RESERVED",
	1: "CERTIFICATE_PARSE_STATUS_UNKNOWN",
	2: "CERTIFICATE_PARSE_STATUS_SUCCESS",
	3: "CERTIFICATE_PARSE_STATUS_FAIL",
	4: "CERTIFICATE_PARSE_STATUS_NOT_PARSED",
}
var CertificateParseStatus_value = map[string]int32{
	"CERTIFICATE_PARSE_STATUS_RESERVED":   0,
	"CERTIFICATE_PARSE_STATUS_UNKNOWN":    1,
	"CERTIFICATE_PARSE_STATUS_SUCCESS":    2,
	"CERTIFICATE_PARSE_STATUS_FAIL":       3,
	"CERTIFICATE_PARSE_STATUS_NOT_PARSED": 4,
}

func (x CertificateParseStatus) String() string {
	return proto.EnumName(CertificateParseStatus_name, int32(x))
}
func (CertificateParseStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type CertificateRevocationReason int32

const (
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_RESERVED CertificateRevocationReason = 0
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_UNKNOWN  CertificateRevocationReason = 1
	// RFC 5280
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_UNSPECIFIED            CertificateRevocationReason = 2
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_KEY_COMPROMISE         CertificateRevocationReason = 3
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_CA_COMPROMISE          CertificateRevocationReason = 4
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_AFFILIATION_CHANGED    CertificateRevocationReason = 5
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_SUPERSEDED             CertificateRevocationReason = 6
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_CESSATION_OF_OPERATION CertificateRevocationReason = 7
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_CERTIFICATE_HOLD       CertificateRevocationReason = 8
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_REMOVE_FROM_CRL        CertificateRevocationReason = 9
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_PRIVILEGE_WITHDRAWN    CertificateRevocationReason = 10
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_AA_COMPROMISE          CertificateRevocationReason = 11
)

var CertificateRevocationReason_name = map[int32]string{
	0:  "CERTIFICATE_REVOCATION_REASON_RESERVED",
	1:  "CERTIFICATE_REVOCATION_REASON_UNKNOWN",
	2:  "CERTIFICATE_REVOCATION_REASON_UNSPECIFIED",
	3:  "CERTIFICATE_REVOCATION_REASON_KEY_COMPROMISE",
	4:  "CERTIFICATE_REVOCATION_REASON_CA_COMPROMISE",
	5:  "CERTIFICATE_REVOCATION_REASON_AFFILIATION_CHANGED",
	6:  "CERTIFICATE_REVOCATION_REASON_SUPERSEDED",
	7:  "CERTIFICATE_REVOCATION_REASON_CESSATION_OF_OPERATION",
	8:  "CERTIFICATE_REVOCATION_REASON_CERTIFICATE_HOLD",
	9:  "CERTIFICATE_REVOCATION_REASON_REMOVE_FROM_CRL",
	10: "CERTIFICATE_REVOCATION_REASON_PRIVILEGE_WITHDRAWN",
	11: "CERTIFICATE_REVOCATION_REASON_AA_COMPROMISE",
}
var CertificateRevocationReason_value = map[string]int32{
	"CERTIFICATE_REVOCATION_REASON_RESERVED":               0,
	"CERTIFICATE_REVOCATION_REASON_UNKNOWN":                1,
	"CERTIFICATE_REVOCATION_REASON_UNSPECIFIED":            2,
	"CERTIFICATE_REVOCATION_REASON_KEY_COMPROMISE":         3,
	"CERTIFICATE_REVOCATION_REASON_CA_COMPROMISE":          4,
	"CERTIFICATE_REVOCATION_REASON_AFFILIATION_CHANGED":    5,
	"CERTIFICATE_REVOCATION_REASON_SUPERSEDED":             6,
	"CERTIFICATE_REVOCATION_REASON_CESSATION_OF_OPERATION": 7,
	"CERTIFICATE_REVOCATION_REASON_CERTIFICATE_HOLD":       8,
	"CERTIFICATE_REVOCATION_REASON_REMOVE_FROM_CRL":        9,
	"CERTIFICATE_REVOCATION_REASON_PRIVILEGE_WITHDRAWN":    10,
	"CERTIFICATE_REVOCATION_REASON_AA_COMPROMISE":          11,
}

func (x CertificateRevocationReason) String() string {
	return proto.EnumName(CertificateRevocationReason_name, int32(x))
}
func (CertificateRevocationReason) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type RootStoreStatus struct {
	// (trusted path || whitelisted) && other certificate validation (e.g., expiration)
	Valid bool `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
	// was certificate ever valid in the past?
	WasValid bool `protobuf:"varint,2,opt,name=was_valid,json=wasValid" json:"was_valid,omitempty"`
	// does there exist a path from the certificate to the root store
	TrustedPath    bool `protobuf:"varint,3,opt,name=trusted_path,json=trustedPath" json:"trusted_path,omitempty"`
	WasTrustedPath bool `protobuf:"varint,4,opt,name=was_trusted_path,json=wasTrustedPath" json:"was_trusted_path,omitempty"`
	// is the certificate explicitly blacklisted by the browser (e.g., OneCRL)
	Blacklisted bool            `protobuf:"varint,5,opt,name=blacklisted" json:"blacklisted,omitempty"`
	Whitelisted bool            `protobuf:"varint,6,opt,name=whitelisted" json:"whitelisted,omitempty"`
	Type        CertificateType `protobuf:"varint,7,opt,name=type,enum=zsearch.CertificateType" json:"type,omitempty"`
	// contains sha256fp of path to root store
	Path [][]byte `protobuf:"bytes,8,rep,name=path,proto3" json:"path,omitempty"`
}

func (m *RootStoreStatus) Reset()                    { *m = RootStoreStatus{} }
func (m *RootStoreStatus) String() string            { return proto.CompactTextString(m) }
func (*RootStoreStatus) ProtoMessage()               {}
func (*RootStoreStatus) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *RootStoreStatus) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *RootStoreStatus) GetWasValid() bool {
	if m != nil {
		return m.WasValid
	}
	return false
}

func (m *RootStoreStatus) GetTrustedPath() bool {
	if m != nil {
		return m.TrustedPath
	}
	return false
}

func (m *RootStoreStatus) GetWasTrustedPath() bool {
	if m != nil {
		return m.WasTrustedPath
	}
	return false
}

func (m *RootStoreStatus) GetBlacklisted() bool {
	if m != nil {
		return m.Blacklisted
	}
	return false
}

func (m *RootStoreStatus) GetWhitelisted() bool {
	if m != nil {
		return m.Whitelisted
	}
	return false
}

func (m *RootStoreStatus) GetType() CertificateType {
	if m != nil {
		return m.Type
	}
	return CertificateType_CERTIFICATE_TYPE_RESERVED
}

func (m *RootStoreStatus) GetPath() [][]byte {
	if m != nil {
		return m.Path
	}
	return nil
}

type CertificateValidation struct {
	Nss       *RootStoreStatus `protobuf:"bytes,1,opt,name=nss" json:"nss,omitempty"`
	Microsoft *RootStoreStatus `protobuf:"bytes,2,opt,name=microsoft" json:"microsoft,omitempty"`
	Apple     *RootStoreStatus `protobuf:"bytes,3,opt,name=apple" json:"apple,omitempty"`
	Java      *RootStoreStatus `protobuf:"bytes,4,opt,name=java" json:"java,omitempty"`
	Android   *RootStoreStatus `protobuf:"bytes,5,opt,name=android" json:"android,omitempty"`
	// also track for Google CT servers so we know what to push
	GoogleCtPrimary *RootStoreStatus `protobuf:"bytes,10,opt,name=google_ct_primary,json=googleCtPrimary" json:"google_ct_primary,omitempty"`
}

func (m *CertificateValidation) Reset()                    { *m = CertificateValidation{} }
func (m *CertificateValidation) String() string            { return proto.CompactTextString(m) }
func (*CertificateValidation) ProtoMessage()               {}
func (*CertificateValidation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CertificateValidation) GetNss() *RootStoreStatus {
	if m != nil {
		return m.Nss
	}
	return nil
}

func (m *CertificateValidation) GetMicrosoft() *RootStoreStatus {
	if m != nil {
		return m.Microsoft
	}
	return nil
}

func (m *CertificateValidation) GetApple() *RootStoreStatus {
	if m != nil {
		return m.Apple
	}
	return nil
}

func (m *CertificateValidation) GetJava() *RootStoreStatus {
	if m != nil {
		return m.Java
	}
	return nil
}

func (m *CertificateValidation) GetAndroid() *RootStoreStatus {
	if m != nil {
		return m.Android
	}
	return nil
}

func (m *CertificateValidation) GetGoogleCtPrimary() *RootStoreStatus {
	if m != nil {
		return m.GoogleCtPrimary
	}
	return nil
}

type MozillaSalesForceStatus struct {
	CurrentIn                      bool   `protobuf:"varint,1,opt,name=current_in,json=currentIn" json:"current_in,omitempty"`
	WasIn                          bool   `protobuf:"varint,2,opt,name=was_in,json=wasIn" json:"was_in,omitempty"`
	OwnerName                      string `protobuf:"bytes,3,opt,name=owner_name,json=ownerName" json:"owner_name,omitempty"`
	ParentName                     string `protobuf:"bytes,4,opt,name=parent_name,json=parentName" json:"parent_name,omitempty"`
	CertificateName                string `protobuf:"bytes,5,opt,name=certificate_name,json=certificateName" json:"certificate_name,omitempty"`
	CertificatePolicy              string `protobuf:"bytes,6,opt,name=certificate_policy,json=certificatePolicy" json:"certificate_policy,omitempty"`
	CertificationPracticeStatement string `protobuf:"bytes,7,opt,name=certification_practice_statement,json=certificationPracticeStatement" json:"certification_practice_statement,omitempty"`
	CpSameAsParent                 bool   `protobuf:"varint,8,opt,name=cp_same_as_parent,json=cpSameAsParent" json:"cp_same_as_parent,omitempty"`
	AuditSameAsParent              bool   `protobuf:"varint,9,opt,name=audit_same_as_parent,json=auditSameAsParent" json:"audit_same_as_parent,omitempty"`
	StandardAudit                  string `protobuf:"bytes,10,opt,name=standard_audit,json=standardAudit" json:"standard_audit,omitempty"`
	BrAudit                        string `protobuf:"bytes,11,opt,name=br_audit,json=brAudit" json:"br_audit,omitempty"`
	Auditor                        string `protobuf:"bytes,12,opt,name=auditor" json:"auditor,omitempty"`
}

func (m *MozillaSalesForceStatus) Reset()                    { *m = MozillaSalesForceStatus{} }
func (m *MozillaSalesForceStatus) String() string            { return proto.CompactTextString(m) }
func (*MozillaSalesForceStatus) ProtoMessage()               {}
func (*MozillaSalesForceStatus) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *MozillaSalesForceStatus) GetCurrentIn() bool {
	if m != nil {
		return m.CurrentIn
	}
	return false
}

func (m *MozillaSalesForceStatus) GetWasIn() bool {
	if m != nil {
		return m.WasIn
	}
	return false
}

func (m *MozillaSalesForceStatus) GetOwnerName() string {
	if m != nil {
		return m.OwnerName
	}
	return ""
}

func (m *MozillaSalesForceStatus) GetParentName() string {
	if m != nil {
		return m.ParentName
	}
	return ""
}

func (m *MozillaSalesForceStatus) GetCertificateName() string {
	if m != nil {
		return m.CertificateName
	}
	return ""
}

func (m *MozillaSalesForceStatus) GetCertificatePolicy() string {
	if m != nil {
		return m.CertificatePolicy
	}
	return ""
}

func (m *MozillaSalesForceStatus) GetCertificationPracticeStatement() string {
	if m != nil {
		return m.CertificationPracticeStatement
	}
	return ""
}

func (m *MozillaSalesForceStatus) GetCpSameAsParent() bool {
	if m != nil {
		return m.CpSameAsParent
	}
	return false
}

func (m *MozillaSalesForceStatus) GetAuditSameAsParent() bool {
	if m != nil {
		return m.AuditSameAsParent
	}
	return false
}

func (m *MozillaSalesForceStatus) GetStandardAudit() string {
	if m != nil {
		return m.StandardAudit
	}
	return ""
}

func (m *MozillaSalesForceStatus) GetBrAudit() string {
	if m != nil {
		return m.BrAudit
	}
	return ""
}

func (m *MozillaSalesForceStatus) GetAuditor() string {
	if m != nil {
		return m.Auditor
	}
	return ""
}

type CertificateRevocation struct {
	Revoked bool                        `protobuf:"varint,1,opt,name=revoked" json:"revoked,omitempty"`
	Reason  CertificateRevocationReason `protobuf:"varint,2,opt,name=reason,enum=zsearch.CertificateRevocationReason" json:"reason,omitempty"`
}

func (m *CertificateRevocation) Reset()                    { *m = CertificateRevocation{} }
func (m *CertificateRevocation) String() string            { return proto.CompactTextString(m) }
func (*CertificateRevocation) ProtoMessage()               {}
func (*CertificateRevocation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *CertificateRevocation) GetRevoked() bool {
	if m != nil {
		return m.Revoked
	}
	return false
}

func (m *CertificateRevocation) GetReason() CertificateRevocationReason {
	if m != nil {
		return m.Reason
	}
	return CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_RESERVED
}

type CertificateAudit struct {
	Mozilla *MozillaSalesForceStatus `protobuf:"bytes,1,opt,name=mozilla" json:"mozilla,omitempty"`
}

func (m *CertificateAudit) Reset()                    { *m = CertificateAudit{} }
func (m *CertificateAudit) String() string            { return proto.CompactTextString(m) }
func (*CertificateAudit) ProtoMessage()               {}
func (*CertificateAudit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *CertificateAudit) GetMozilla() *MozillaSalesForceStatus {
	if m != nil {
		return m.Mozilla
	}
	return nil
}

// Database Records
type Certificate struct {
	Sha1Fp       []byte                 `protobuf:"bytes,1,opt,name=sha1fp,proto3" json:"sha1fp,omitempty"`
	Sha256Fp     []byte                 `protobuf:"bytes,2,opt,name=sha256fp,proto3" json:"sha256fp,omitempty"`
	Raw          []byte                 `protobuf:"bytes,3,opt,name=raw,proto3" json:"raw,omitempty"`
	Parsed       string                 `protobuf:"bytes,4,opt,name=parsed" json:"parsed,omitempty"`
	ParseStatus  CertificateParseStatus `protobuf:"varint,44,opt,name=parse_status,json=parseStatus,enum=zsearch.CertificateParseStatus" json:"parse_status,omitempty"`
	ParseVersion uint32                 `protobuf:"varint,39,opt,name=parse_version,json=parseVersion" json:"parse_version,omitempty"`
	ParseError   string                 `protobuf:"bytes,47,opt,name=parse_error,json=parseError" json:"parse_error,omitempty"`
	Parents      [][]byte               `protobuf:"bytes,5,rep,name=parents,proto3" json:"parents,omitempty"`
	// the chain that we received with the certificate and pass
	// to the certificate processing daemon.
	PresentedChain       [][]byte               `protobuf:"bytes,45,rep,name=presented_chain,json=presentedChain,proto3" json:"presented_chain,omitempty"`
	Source               CertificateSource      `protobuf:"varint,28,opt,name=source,enum=zsearch.CertificateSource" json:"source,omitempty"`
	SeenInScan           bool                   `protobuf:"varint,29,opt,name=seen_in_scan,json=seenInScan" json:"seen_in_scan,omitempty"`
	PostProcessed        bool                   `protobuf:"varint,26,opt,name=post_processed,json=postProcessed" json:"post_processed,omitempty"`
	PostProcessTimestamp uint32                 `protobuf:"varint,37,opt,name=post_process_timestamp,json=postProcessTimestamp" json:"post_process_timestamp,omitempty"`
	Validation           *CertificateValidation `protobuf:"bytes,35,opt,name=validation" json:"validation,omitempty"`
	Ct                   *CTStatus              `protobuf:"bytes,30,opt,name=ct" json:"ct,omitempty"`
	Zlint                *ZLint                 `protobuf:"bytes,38,opt,name=zlint" json:"zlint,omitempty"`
	Revocation           *CertificateRevocation `protobuf:"bytes,43,opt,name=revocation" json:"revocation,omitempty"`
	Audit                *CertificateAudit      `protobuf:"bytes,46,opt,name=audit" json:"audit,omitempty"`
	// store wheter record is precert so that CT synchronization
	// daemon can quickly decide what to do without reparsing the
	// raw certificate
	IsPrecert bool `protobuf:"varint,32,opt,name=is_precert,json=isPrecert" json:"is_precert,omitempty"`
	// store time range that certificate is valid so that cert
	// dump job can mark certificates as expired
	NotValidAfter         uint32                   `protobuf:"varint,41,opt,name=not_valid_after,json=notValidAfter" json:"not_valid_after,omitempty"`
	NotValidBefore        uint32                   `protobuf:"varint,42,opt,name=not_valid_before,json=notValidBefore" json:"not_valid_before,omitempty"`
	InNss                 bool                     `protobuf:"varint,6,opt,name=in_nss,json=inNss" json:"in_nss,omitempty"`
	InMicrosoft           bool                     `protobuf:"varint,7,opt,name=in_microsoft,json=inMicrosoft" json:"in_microsoft,omitempty"`
	InApple               bool                     `protobuf:"varint,8,opt,name=in_apple,json=inApple" json:"in_apple,omitempty"`
	ValidationTimestamp   uint32                   `protobuf:"varint,10,opt,name=validation_timestamp,json=validationTimestamp" json:"validation_timestamp,omitempty"`
	ValidNss              bool                     `protobuf:"varint,11,opt,name=valid_nss,json=validNss" json:"valid_nss,omitempty"`
	ValidMicrosoft        bool                     `protobuf:"varint,12,opt,name=valid_microsoft,json=validMicrosoft" json:"valid_microsoft,omitempty"`
	ValidApple            bool                     `protobuf:"varint,13,opt,name=valid_apple,json=validApple" json:"valid_apple,omitempty"`
	WasValidNss           bool                     `protobuf:"varint,14,opt,name=was_valid_nss,json=wasValidNss" json:"was_valid_nss,omitempty"`
	WasValidMicrosoft     bool                     `protobuf:"varint,15,opt,name=was_valid_microsoft,json=wasValidMicrosoft" json:"was_valid_microsoft,omitempty"`
	WasValidApple         bool                     `protobuf:"varint,16,opt,name=was_valid_apple,json=wasValidApple" json:"was_valid_apple,omitempty"`
	WasInNss              bool                     `protobuf:"varint,17,opt,name=was_in_nss,json=wasInNss" json:"was_in_nss,omitempty"`
	WasInMicrosoft        bool                     `protobuf:"varint,18,opt,name=was_in_microsoft,json=wasInMicrosoft" json:"was_in_microsoft,omitempty"`
	WasInApple            bool                     `protobuf:"varint,19,opt,name=was_in_apple,json=wasInApple" json:"was_in_apple,omitempty"`
	CurrentValidNss       bool                     `protobuf:"varint,20,opt,name=current_valid_nss,json=currentValidNss" json:"current_valid_nss,omitempty"`
	CurrentValidMicrosoft bool                     `protobuf:"varint,21,opt,name=current_valid_microsoft,json=currentValidMicrosoft" json:"current_valid_microsoft,omitempty"`
	CurrentValidApple     bool                     `protobuf:"varint,22,opt,name=current_valid_apple,json=currentValidApple" json:"current_valid_apple,omitempty"`
	CurrentInNss          bool                     `protobuf:"varint,23,opt,name=current_in_nss,json=currentInNss" json:"current_in_nss,omitempty"`
	CurrentInMicrosoft    bool                     `protobuf:"varint,24,opt,name=current_in_microsoft,json=currentInMicrosoft" json:"current_in_microsoft,omitempty"`
	CurrentInApple        bool                     `protobuf:"varint,25,opt,name=current_in_apple,json=currentInApple" json:"current_in_apple,omitempty"`
	NssAudit              *MozillaSalesForceStatus `protobuf:"bytes,31,opt,name=nss_audit,json=nssAudit" json:"nss_audit,omitempty"`
	ShouldPostProcess     bool                     `protobuf:"varint,27,opt,name=should_post_process,json=shouldPostProcess" json:"should_post_process,omitempty"`
	DoNotPostProcess      bool                     `protobuf:"varint,36,opt,name=do_not_post_process,json=doNotPostProcess" json:"do_not_post_process,omitempty"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *Certificate) GetSha1Fp() []byte {
	if m != nil {
		return m.Sha1Fp
	}
	return nil
}

func (m *Certificate) GetSha256Fp() []byte {
	if m != nil {
		return m.Sha256Fp
	}
	return nil
}

func (m *Certificate) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Certificate) GetParsed() string {
	if m != nil {
		return m.Parsed
	}
	return ""
}

func (m *Certificate) GetParseStatus() CertificateParseStatus {
	if m != nil {
		return m.ParseStatus
	}
	return CertificateParseStatus_CERTIFICATE_PARSE_STATUS_RESERVED
}

func (m *Certificate) GetParseVersion() uint32 {
	if m != nil {
		return m.ParseVersion
	}
	return 0
}

func (m *Certificate) GetParseError() string {
	if m != nil {
		return m.ParseError
	}
	return ""
}

func (m *Certificate) GetParents() [][]byte {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *Certificate) GetPresentedChain() [][]byte {
	if m != nil {
		return m.PresentedChain
	}
	return nil
}

func (m *Certificate) GetSource() CertificateSource {
	if m != nil {
		return m.Source
	}
	return CertificateSource_CERTIFICATE_SOURCE_RESERVED
}

func (m *Certificate) GetSeenInScan() bool {
	if m != nil {
		return m.SeenInScan
	}
	return false
}

func (m *Certificate) GetPostProcessed() bool {
	if m != nil {
		return m.PostProcessed
	}
	return false
}

func (m *Certificate) GetPostProcessTimestamp() uint32 {
	if m != nil {
		return m.PostProcessTimestamp
	}
	return 0
}

func (m *Certificate) GetValidation() *CertificateValidation {
	if m != nil {
		return m.Validation
	}
	return nil
}

func (m *Certificate) GetCt() *CTStatus {
	if m != nil {
		return m.Ct
	}
	return nil
}

func (m *Certificate) GetZlint() *ZLint {
	if m != nil {
		return m.Zlint
	}
	return nil
}

func (m *Certificate) GetRevocation() *CertificateRevocation {
	if m != nil {
		return m.Revocation
	}
	return nil
}

func (m *Certificate) GetAudit() *CertificateAudit {
	if m != nil {
		return m.Audit
	}
	return nil
}

func (m *Certificate) GetIsPrecert() bool {
	if m != nil {
		return m.IsPrecert
	}
	return false
}

func (m *Certificate) GetNotValidAfter() uint32 {
	if m != nil {
		return m.NotValidAfter
	}
	return 0
}

func (m *Certificate) GetNotValidBefore() uint32 {
	if m != nil {
		return m.NotValidBefore
	}
	return 0
}

func (m *Certificate) GetInNss() bool {
	if m != nil {
		return m.InNss
	}
	return false
}

func (m *Certificate) GetInMicrosoft() bool {
	if m != nil {
		return m.InMicrosoft
	}
	return false
}

func (m *Certificate) GetInApple() bool {
	if m != nil {
		return m.InApple
	}
	return false
}

func (m *Certificate) GetValidationTimestamp() uint32 {
	if m != nil {
		return m.ValidationTimestamp
	}
	return 0
}

func (m *Certificate) GetValidNss() bool {
	if m != nil {
		return m.ValidNss
	}
	return false
}

func (m *Certificate) GetValidMicrosoft() bool {
	if m != nil {
		return m.ValidMicrosoft
	}
	return false
}

func (m *Certificate) GetValidApple() bool {
	if m != nil {
		return m.ValidApple
	}
	return false
}

func (m *Certificate) GetWasValidNss() bool {
	if m != nil {
		return m.WasValidNss
	}
	return false
}

func (m *Certificate) GetWasValidMicrosoft() bool {
	if m != nil {
		return m.WasValidMicrosoft
	}
	return false
}

func (m *Certificate) GetWasValidApple() bool {
	if m != nil {
		return m.WasValidApple
	}
	return false
}

func (m *Certificate) GetWasInNss() bool {
	if m != nil {
		return m.WasInNss
	}
	return false
}

func (m *Certificate) GetWasInMicrosoft() bool {
	if m != nil {
		return m.WasInMicrosoft
	}
	return false
}

func (m *Certificate) GetWasInApple() bool {
	if m != nil {
		return m.WasInApple
	}
	return false
}

func (m *Certificate) GetCurrentValidNss() bool {
	if m != nil {
		return m.CurrentValidNss
	}
	return false
}

func (m *Certificate) GetCurrentValidMicrosoft() bool {
	if m != nil {
		return m.CurrentValidMicrosoft
	}
	return false
}

func (m *Certificate) GetCurrentValidApple() bool {
	if m != nil {
		return m.CurrentValidApple
	}
	return false
}

func (m *Certificate) GetCurrentInNss() bool {
	if m != nil {
		return m.CurrentInNss
	}
	return false
}

func (m *Certificate) GetCurrentInMicrosoft() bool {
	if m != nil {
		return m.CurrentInMicrosoft
	}
	return false
}

func (m *Certificate) GetCurrentInApple() bool {
	if m != nil {
		return m.CurrentInApple
	}
	return false
}

func (m *Certificate) GetNssAudit() *MozillaSalesForceStatus {
	if m != nil {
		return m.NssAudit
	}
	return nil
}

func (m *Certificate) GetShouldPostProcess() bool {
	if m != nil {
		return m.ShouldPostProcess
	}
	return false
}

func (m *Certificate) GetDoNotPostProcess() bool {
	if m != nil {
		return m.DoNotPostProcess
	}
	return false
}

func init() {
	proto.RegisterType((*RootStoreStatus)(nil), "zsearch.RootStoreStatus")
	proto.RegisterType((*CertificateValidation)(nil), "zsearch.CertificateValidation")
	proto.RegisterType((*MozillaSalesForceStatus)(nil), "zsearch.MozillaSalesForceStatus")
	proto.RegisterType((*CertificateRevocation)(nil), "zsearch.CertificateRevocation")
	proto.RegisterType((*CertificateAudit)(nil), "zsearch.CertificateAudit")
	proto.RegisterType((*Certificate)(nil), "zsearch.Certificate")
	proto.RegisterEnum("zsearch.CertificateType", CertificateType_name, CertificateType_value)
	proto.RegisterEnum("zsearch.CertificateSource", CertificateSource_name, CertificateSource_value)
	proto.RegisterEnum("zsearch.CertificateParseStatus", CertificateParseStatus_name, CertificateParseStatus_value)
	proto.RegisterEnum("zsearch.CertificateRevocationReason", CertificateRevocationReason_name, CertificateRevocationReason_value)
}

func init() { proto.RegisterFile("certificate.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x98, 0xdf, 0x52, 0xe3, 0xc8,
	0x15, 0xc6, 0x83, 0xc1, 0x60, 0x1f, 0x1b, 0x5b, 0x34, 0x30, 0x23, 0x60, 0x19, 0x3c, 0x0c, 0xcc,
	0x80, 0x87, 0x61, 0x82, 0x33, 0xbb, 0x49, 0x6d, 0xa5, 0x52, 0x25, 0x64, 0x79, 0x51, 0xad, 0xb1,
	0x5d, 0x2d, 0xc3, 0xd4, 0xee, 0x8d, 0xaa, 0x91, 0x9b, 0x45, 0x19, 0x5b, 0x52, 0xa9, 0x05, 0xd4,
	0xcc, 0x8b, 0xe4, 0x01, 0x72, 0x91, 0x47, 0xc8, 0x83, 0xe4, 0x22, 0x0f, 0x91, 0xdb, 0x3c, 0x40,
	0xaa, 0xbb, 0x65, 0x49, 0x06, 0x63, 0x72, 0x85, 0xfb, 0x9c, 0xdf, 0xd7, 0xfa, 0xfa, 0xf4, 0xdf,
	0x02, 0x56, 0x1c, 0x1a, 0x46, 0xee, 0xb5, 0xeb, 0x90, 0x88, 0x1e, 0x07, 0xa1, 0x1f, 0xf9, 0x68,
	0xe9, 0x1b, 0xa3, 0x24, 0x74, 0x6e, 0x36, 0xcb, 0x8e, 0x3f, 0x1a, 0xf9, 0x9e, 0x0c, 0x6f, 0x96,
	0xbe, 0x0d, 0x5d, 0x2f, 0x8a, 0x1b, 0x05, 0x27, 0xfe, 0xb5, 0xfb, 0xb7, 0x1c, 0x54, 0xb1, 0xef,
	0x47, 0x56, 0xe4, 0x87, 0xd4, 0x8a, 0x48, 0x74, 0xcb, 0xd0, 0x1a, 0xe4, 0xef, 0xc8, 0xd0, 0x1d,
	0xa8, 0x73, 0xb5, 0xb9, 0x83, 0x02, 0x96, 0x0d, 0xb4, 0x05, 0xc5, 0x7b, 0xc2, 0x6c, 0x99, 0xc9,
	0x89, 0x4c, 0xe1, 0x9e, 0xb0, 0x4b, 0x91, 0x7c, 0x0d, 0xe5, 0x28, 0xbc, 0x65, 0x11, 0x1d, 0xd8,
	0x01, 0x89, 0x6e, 0xd4, 0x79, 0x91, 0x2f, 0xc5, 0xb1, 0x1e, 0x89, 0x6e, 0xd0, 0x01, 0x28, 0x5c,
	0x3f, 0x81, 0x2d, 0x08, 0xac, 0x72, 0x4f, 0x58, 0x3f, 0x43, 0xd6, 0xa0, 0x74, 0x35, 0x24, 0xce,
	0x97, 0xa1, 0xcb, 0x43, 0x6a, 0x5e, 0xf6, 0x95, 0x09, 0x71, 0xe2, 0xfe, 0xc6, 0x8d, 0x68, 0x4c,
	0x2c, 0x4a, 0x22, 0x13, 0x42, 0x47, 0xb0, 0x10, 0x7d, 0x0d, 0xa8, 0xba, 0x54, 0x9b, 0x3b, 0xa8,
	0x34, 0xd4, 0xe3, 0xb8, 0x28, 0xc7, 0x7a, 0x5a, 0xaf, 0xfe, 0xd7, 0x80, 0x62, 0x41, 0x21, 0x04,
	0x0b, 0xc2, 0x4f, 0xa1, 0x36, 0x7f, 0x50, 0xc6, 0xe2, 0xf7, 0xee, 0xbf, 0x72, 0xb0, 0x9e, 0xa1,
	0xc5, 0x38, 0x49, 0xe4, 0xfa, 0x1e, 0xaa, 0xc3, 0xbc, 0xc7, 0x98, 0xa8, 0x4e, 0x29, 0xd3, 0xf5,
	0x83, 0x32, 0x62, 0x0e, 0xa1, 0x1f, 0xa0, 0x38, 0x72, 0x9d, 0xd0, 0x67, 0xfe, 0x75, 0x24, 0xaa,
	0x36, 0x4b, 0x91, 0xa2, 0xe8, 0x18, 0xf2, 0x24, 0x08, 0x86, 0x54, 0x54, 0x72, 0x96, 0x46, 0x62,
	0x7c, 0xbc, 0x7f, 0x25, 0x77, 0x44, 0x54, 0x74, 0x16, 0x2e, 0x28, 0xd4, 0x80, 0x25, 0xe2, 0x0d,
	0x42, 0xdf, 0x95, 0xd5, 0x9d, 0x25, 0x18, 0x83, 0xa8, 0x09, 0x2b, 0xbf, 0xf9, 0xfe, 0x6f, 0x43,
	0x6a, 0x3b, 0x91, 0x1d, 0x84, 0xee, 0x88, 0x84, 0x5f, 0x55, 0x78, 0x46, 0x5d, 0x95, 0x12, 0x3d,
	0xea, 0x49, 0xc1, 0xee, 0x7f, 0xe6, 0xe1, 0xe5, 0xb9, 0xff, 0xcd, 0x1d, 0x0e, 0x89, 0x45, 0x86,
	0x94, 0xb5, 0xfc, 0xd0, 0x19, 0xaf, 0xbb, 0x6d, 0x00, 0xe7, 0x36, 0x0c, 0xa9, 0x17, 0xd9, 0xae,
	0x17, 0x2f, 0xbe, 0x62, 0x1c, 0x31, 0x3d, 0xb4, 0x0e, 0x8b, 0x7c, 0x01, 0xb9, 0x5e, 0xbc, 0xfa,
	0xf2, 0xf7, 0x84, 0x99, 0x1e, 0x57, 0xf9, 0xf7, 0x1e, 0x0d, 0x6d, 0x8f, 0x8c, 0x64, 0xb9, 0x8a,
	0xb8, 0x28, 0x22, 0x1d, 0x32, 0xa2, 0x68, 0x07, 0x4a, 0x01, 0x11, 0x7d, 0x8a, 0xfc, 0x82, 0xc8,
	0x83, 0x0c, 0x09, 0xe0, 0x10, 0x94, 0xcc, 0x26, 0x92, 0x54, 0x5e, 0x50, 0xd5, 0x4c, 0x5c, 0xa0,
	0x1f, 0x00, 0x65, 0xd1, 0xc0, 0x1f, 0xba, 0xce, 0x57, 0xb1, 0xfa, 0x8a, 0x38, 0xbb, 0x13, 0x7b,
	0x22, 0x81, 0xce, 0xa0, 0x96, 0x06, 0x5d, 0xdf, 0xb3, 0x83, 0x90, 0x38, 0x91, 0xeb, 0x50, 0x9b,
	0x45, 0x24, 0xa2, 0x23, 0xea, 0x45, 0x62, 0x7d, 0x16, 0xf1, 0xab, 0x09, 0xae, 0x17, 0x63, 0xd6,
	0x98, 0x42, 0x87, 0xb0, 0xe2, 0x04, 0x36, 0x23, 0x23, 0x6a, 0x13, 0x66, 0x4b, 0xf3, 0x6a, 0x41,
	0x6e, 0x1e, 0x27, 0xb0, 0xc8, 0x88, 0x6a, 0xac, 0x27, 0xa2, 0xe8, 0x23, 0xac, 0x91, 0xdb, 0x81,
	0x1b, 0x3d, 0xa4, 0x8b, 0x82, 0x5e, 0x11, 0xb9, 0x09, 0xc1, 0x3e, 0x54, 0x58, 0x44, 0xbc, 0x01,
	0x09, 0x07, 0xb6, 0xc8, 0x8a, 0x49, 0x2d, 0xe2, 0xe5, 0x71, 0x54, 0xe3, 0x41, 0xb4, 0x01, 0x85,
	0xab, 0x30, 0x06, 0x4a, 0x02, 0x58, 0xba, 0x0a, 0x65, 0x4a, 0x85, 0x25, 0x11, 0xf7, 0x43, 0xb5,
	0x2c, 0x33, 0x71, 0x73, 0xd7, 0x9f, 0xd8, 0x42, 0x98, 0xde, 0xf9, 0x72, 0x84, 0x5c, 0x12, 0xd2,
	0x3b, 0xff, 0x0b, 0x1d, 0x1f, 0x32, 0xe3, 0x26, 0xfa, 0x33, 0x2c, 0x86, 0x94, 0x30, 0x5f, 0xce,
	0x72, 0xa5, 0xb1, 0x37, 0x6d, 0xeb, 0xa6, 0x3d, 0x61, 0xc1, 0xe2, 0x58, 0xb3, 0xdb, 0x01, 0x25,
	0x83, 0x49, 0x7b, 0x3f, 0xc2, 0xd2, 0x48, 0xae, 0xb8, 0x78, 0xcb, 0xd6, 0x92, 0x2e, 0x9f, 0x58,
	0x89, 0x78, 0x2c, 0xd8, 0xfd, 0x7b, 0x05, 0x4a, 0x99, 0x0e, 0xd1, 0x0b, 0x58, 0x64, 0x37, 0xe4,
	0xe4, 0x3a, 0x10, 0x5d, 0x95, 0x71, 0xdc, 0x42, 0x9b, 0x50, 0x60, 0x37, 0xa4, 0xf1, 0xfd, 0x0f,
	0xd7, 0x81, 0xf0, 0x5d, 0xc6, 0x49, 0x1b, 0x29, 0x30, 0x1f, 0x92, 0x7b, 0xb1, 0x32, 0xcb, 0x98,
	0xff, 0xe4, 0xbd, 0x04, 0x24, 0x64, 0x74, 0x10, 0x2f, 0xc7, 0xb8, 0x85, 0x4e, 0xa1, 0x2c, 0x7e,
	0x89, 0xf5, 0x71, 0xcb, 0xd4, 0x23, 0x51, 0x81, 0x9d, 0x69, 0x15, 0xe8, 0x71, 0x2e, 0x76, 0x5b,
	0x0a, 0xd2, 0x06, 0x7a, 0x03, 0xcb, 0xb2, 0x8f, 0x3b, 0x1a, 0x32, 0xd7, 0xf7, 0xd4, 0x77, 0xb5,
	0xb9, 0x83, 0x65, 0x2c, 0x3b, 0xbe, 0x94, 0xb1, 0x78, 0x53, 0x30, 0x6a, 0xd3, 0x30, 0xf4, 0x43,
	0xf5, 0x63, 0xb2, 0x29, 0x18, 0x35, 0x78, 0x84, 0xcf, 0x8f, 0x5c, 0x37, 0x4c, 0xcd, 0x8b, 0x33,
	0x71, 0xdc, 0x44, 0xef, 0xa0, 0x1a, 0x84, 0x94, 0x51, 0x8f, 0x1f, 0xe2, 0xce, 0x0d, 0x71, 0x3d,
	0xf5, 0x83, 0x20, 0x2a, 0x49, 0x58, 0xe7, 0x51, 0xd4, 0x80, 0x45, 0xe6, 0xdf, 0x86, 0x0e, 0x55,
	0xbf, 0x13, 0xc3, 0xd8, 0x9c, 0x36, 0x0c, 0x4b, 0x10, 0x38, 0x26, 0x51, 0x0d, 0xca, 0x8c, 0x52,
	0xcf, 0x76, 0x3d, 0x9b, 0x39, 0xc4, 0x53, 0xb7, 0xc5, 0xda, 0x00, 0x1e, 0x33, 0x3d, 0xcb, 0x21,
	0x1e, 0x5f, 0xad, 0x81, 0xcf, 0xf8, 0x01, 0xe4, 0x3b, 0x94, 0xf1, 0x12, 0x6e, 0x0a, 0x66, 0x99,
	0x47, 0x7b, 0xe3, 0x20, 0xfa, 0x04, 0x2f, 0xb2, 0x98, 0x1d, 0xb9, 0x23, 0xca, 0x22, 0x32, 0x0a,
	0xd4, 0x7d, 0x51, 0x8e, 0xb5, 0x0c, 0xde, 0x1f, 0xe7, 0xd0, 0x5f, 0x00, 0xee, 0x92, 0x63, 0x5e,
	0x7d, 0x23, 0x16, 0xcb, 0xab, 0x69, 0xb6, 0xd3, 0xcb, 0x00, 0x67, 0x14, 0xe8, 0x35, 0xe4, 0x9c,
	0x48, 0x7d, 0x25, 0x74, 0x2b, 0xa9, 0xae, 0x1f, 0xcf, 0x53, 0xce, 0x89, 0xd0, 0x1e, 0xe4, 0xc5,
	0x45, 0xac, 0xbe, 0x15, 0x54, 0x25, 0xa1, 0x7e, 0x6d, 0xbb, 0x5e, 0x84, 0x65, 0x92, 0x1b, 0x09,
	0x93, 0x25, 0xae, 0xbe, 0x7f, 0xda, 0x48, 0x66, 0x23, 0x64, 0x14, 0xe8, 0x23, 0xe4, 0xe5, 0x4e,
	0x3d, 0x16, 0xd2, 0x8d, 0x69, 0x52, 0xb1, 0x39, 0xb0, 0xe4, 0xf8, 0x21, 0xea, 0x32, 0x3b, 0x08,
	0x29, 0x3f, 0x88, 0xd4, 0x9a, 0x3c, 0x7a, 0x5d, 0xd6, 0x93, 0x01, 0xf4, 0x16, 0xaa, 0x9e, 0x1f,
	0xc9, 0xbb, 0xdf, 0x26, 0xd7, 0x11, 0x0d, 0xd5, 0x43, 0x51, 0xc7, 0x65, 0xcf, 0x8f, 0x44, 0x31,
	0x34, 0x1e, 0xe4, 0x77, 0x7c, 0xca, 0x5d, 0xd1, 0x6b, 0x3f, 0xa4, 0x6a, 0x5d, 0x80, 0x95, 0x31,
	0x78, 0x2a, 0xa2, 0xfc, 0x30, 0x77, 0x3d, 0x9b, 0x5f, 0xa3, 0xf2, 0xf2, 0xce, 0xbb, 0x5e, 0x87,
	0x31, 0xb4, 0x0f, 0x65, 0xd7, 0xb3, 0xd3, 0x1b, 0x93, 0x1f, 0x8f, 0x85, 0xd3, 0x9c, 0x3a, 0x87,
	0x4b, 0xae, 0x77, 0x9e, 0xdc, 0x8e, 0xdb, 0x50, 0x70, 0x3d, 0x5b, 0x5e, 0x90, 0x85, 0x04, 0x59,
	0x72, 0x3d, 0x4d, 0x5c, 0x86, 0x27, 0xb0, 0x96, 0xce, 0x4a, 0x66, 0xee, 0x41, 0x58, 0x59, 0x4d,
	0x73, 0xe9, 0xd4, 0xef, 0x40, 0x51, 0xba, 0xe6, 0x96, 0x4a, 0x49, 0x97, 0x05, 0x11, 0xe4, 0xce,
	0xde, 0x43, 0x55, 0x02, 0xa9, 0xb9, 0x72, 0x82, 0x55, 0x44, 0x2a, 0xf5, 0xf7, 0x06, 0x4a, 0x71,
	0xad, 0x84, 0xc5, 0xe5, 0x04, 0x94, 0xab, 0x45, 0xba, 0xdc, 0x85, 0xe5, 0xe4, 0x41, 0x25, 0x3e,
	0x5b, 0x89, 0x9f, 0x31, 0xf1, 0xa3, 0x8a, 0x7f, 0xb5, 0x01, 0xab, 0x29, 0x93, 0x7e, 0xb9, 0x9a,
	0x74, 0xb8, 0x32, 0xa6, 0xd3, 0x8f, 0xd7, 0xa1, 0x9a, 0x6a, 0xa4, 0x01, 0x25, 0xe1, 0x97, 0xc7,
	0xbc, 0xf4, 0x50, 0x03, 0x90, 0x77, 0xaa, 0x30, 0xb0, 0x92, 0x8e, 0x5b, 0xdc, 0xad, 0xdc, 0xc1,
	0x91, 0x7c, 0xb6, 0x4d, 0xcc, 0x0a, 0x4a, 0x07, 0x2e, 0xb8, 0xf4, 0xdb, 0x7b, 0x50, 0x8e, 0x69,
	0xf9, 0xe1, 0xd5, 0x74, 0xe4, 0x82, 0x94, 0x5f, 0xad, 0xc3, 0xca, 0xf8, 0xa2, 0x4f, 0x47, 0xbf,
	0x26, 0x46, 0x5f, 0x8d, 0x13, 0x49, 0x05, 0x7e, 0x84, 0x97, 0x93, 0x6c, 0x6a, 0x63, 0x3d, 0xe9,
	0x7c, 0x3d, 0xab, 0x4a, 0xdd, 0x34, 0x60, 0x75, 0x52, 0x2b, 0x4d, 0xbd, 0x48, 0xab, 0x97, 0xd5,
	0x49, 0x6f, 0x7b, 0x50, 0x49, 0x1f, 0x21, 0xc2, 0xd8, 0x4b, 0x61, 0xac, 0x9c, 0x3c, 0x44, 0xb8,
	0xab, 0x4f, 0xb0, 0x96, 0xa1, 0x52, 0x4b, 0x6a, 0xd2, 0x35, 0x4a, 0xf8, 0xd4, 0xcf, 0x11, 0x28,
	0x19, 0x95, 0x34, 0xb3, 0x91, 0xd6, 0x32, 0x51, 0x48, 0x27, 0x1a, 0x14, 0x3d, 0xc6, 0xe2, 0x2b,
	0x77, 0xe7, 0xff, 0xbb, 0xb9, 0xe4, 0xe4, 0x79, 0x8c, 0xc9, 0xab, 0xaf, 0x01, 0xab, 0xec, 0xc6,
	0xbf, 0x1d, 0x0e, 0xec, 0xec, 0x69, 0xa8, 0x6e, 0xa5, 0x05, 0x90, 0xe9, 0x5e, 0x7a, 0x1a, 0xa2,
	0x13, 0x58, 0x1d, 0xf8, 0x36, 0xdf, 0xc6, 0x13, 0x9a, 0xbd, 0x44, 0xa3, 0x0c, 0xfc, 0x8e, 0x1f,
	0x65, 0x24, 0xf5, 0x7f, 0xcc, 0x41, 0xf5, 0xc1, 0xc3, 0x1a, 0x6d, 0xc3, 0x86, 0x6e, 0xe0, 0xbe,
	0xd9, 0x32, 0x75, 0xad, 0x6f, 0xd8, 0xfd, 0x5f, 0x7a, 0x86, 0x8d, 0x0d, 0xcb, 0xc0, 0x97, 0x46,
	0x53, 0xf9, 0x1d, 0xfa, 0x0e, 0xd4, 0x47, 0xe9, 0x8b, 0xce, 0xcf, 0x9d, 0xee, 0xe7, 0x8e, 0x32,
	0x87, 0x36, 0x60, 0xfd, 0x51, 0xb6, 0x6d, 0x68, 0x2d, 0x25, 0x87, 0x5e, 0xc3, 0xf6, 0xa3, 0x94,
	0xd9, 0xe9, 0x1b, 0xf8, 0xdc, 0x68, 0x9a, 0x5a, 0xdf, 0x50, 0xe6, 0xa7, 0xaa, 0x71, 0xb7, 0xdb,
	0x57, 0x16, 0xea, 0xff, 0xcc, 0xc1, 0xca, 0xa3, 0xeb, 0x07, 0xed, 0xc0, 0x56, 0x56, 0x60, 0x75,
	0x2f, 0xb0, 0x3e, 0xe1, 0xf6, 0x15, 0x6c, 0x4e, 0x01, 0x52, 0xbf, 0x5b, 0xf0, 0x72, 0x4a, 0xde,
	0xd2, 0xb5, 0x8e, 0x92, 0x7b, 0x68, 0x27, 0x4e, 0xea, 0x7d, 0x65, 0x1e, 0x1d, 0xc2, 0xfe, 0x94,
	0xd4, 0x79, 0xf7, 0x57, 0xb3, 0xdd, 0xd6, 0x6c, 0x4b, 0x6b, 0x1b, 0x56, 0xab, 0x8b, 0x75, 0x43,
	0x59, 0x98, 0xe1, 0x51, 0xc3, 0xfa, 0x99, 0x92, 0x7f, 0x58, 0xf0, 0x31, 0xa0, 0xf5, 0xcc, 0xe6,
	0x1f, 0x95, 0xc5, 0x27, 0xd2, 0x67, 0x17, 0xa7, 0xa7, 0x6d, 0x43, 0x59, 0x7a, 0xa2, 0x7b, 0xbd,
	0x6f, 0xeb, 0x67, 0x9a, 0xd9, 0x51, 0x0a, 0xf5, 0x7f, 0xcf, 0xc1, 0x8b, 0xe9, 0xef, 0x0f, 0xb4,
	0x0f, 0xaf, 0xb3, 0xda, 0x9e, 0x86, 0x2d, 0xc3, 0xb6, 0xfa, 0x5a, 0xff, 0xc2, 0xca, 0x16, 0x71,
	0x0f, 0x6a, 0x4f, 0x62, 0x69, 0x29, 0x67, 0x51, 0xd6, 0x85, 0xae, 0x1b, 0x96, 0xf5, 0x78, 0x15,
	0x4c, 0x50, 0x2d, 0xcd, 0x6c, 0x2b, 0xf3, 0xe8, 0x1d, 0xbc, 0x79, 0x12, 0xe9, 0x74, 0xfb, 0x32,
	0xd0, 0x54, 0x16, 0xea, 0xff, 0x5d, 0x80, 0xad, 0x19, 0x6f, 0x4b, 0x54, 0x87, 0xb7, 0xd9, 0x8e,
	0xb0, 0x71, 0xd9, 0xd5, 0xb5, 0xbe, 0xd9, 0xed, 0xd8, 0xd8, 0xd0, 0x2c, 0xf1, 0x27, 0x19, 0xe3,
	0x83, 0x09, 0x7d, 0xcc, 0xa6, 0x03, 0xfd, 0x00, 0x87, 0xcf, 0xa1, 0x56, 0xcf, 0xd0, 0xcd, 0x96,
	0x69, 0x34, 0x95, 0x1c, 0xfa, 0x3d, 0x1c, 0xcd, 0xc6, 0x7f, 0x36, 0x7e, 0xb1, 0xf5, 0xee, 0x79,
	0x0f, 0x77, 0xcf, 0x4d, 0x8b, 0x6f, 0x83, 0x8f, 0xf0, 0x7e, 0xb6, 0x42, 0xd7, 0xb2, 0x82, 0x05,
	0xf4, 0x3d, 0x9c, 0xcc, 0x16, 0x68, 0xad, 0x96, 0xd9, 0x36, 0x65, 0x48, 0x3f, 0xd3, 0x3a, 0x3f,
	0x19, 0x4d, 0x25, 0x8f, 0x8e, 0xe0, 0x60, 0xb6, 0xcc, 0xba, 0xe8, 0x19, 0xbc, 0xd6, 0x46, 0x53,
	0x59, 0x44, 0x7f, 0x82, 0x4f, 0xcf, 0xb8, 0x32, 0x2c, 0x4b, 0x06, 0xba, 0x2d, 0xbb, 0xdb, 0x33,
	0xb0, 0x68, 0x28, 0x4b, 0xa8, 0x01, 0xc7, 0xcf, 0x29, 0xd3, 0xec, 0x59, 0xb7, 0xdd, 0x54, 0x0a,
	0xe8, 0x04, 0x3e, 0x3c, 0x37, 0x77, 0xe7, 0xdd, 0x4b, 0xc3, 0x6e, 0xe1, 0xee, 0xb9, 0xad, 0xe3,
	0xb6, 0x52, 0x7c, 0xbe, 0x0a, 0x3d, 0x6c, 0x5e, 0x9a, 0x6d, 0xe3, 0x27, 0xc3, 0xfe, 0x6c, 0xf6,
	0xcf, 0x9a, 0x58, 0xfb, 0xdc, 0x51, 0xe0, 0xf9, 0x6a, 0x6b, 0x13, 0xd5, 0x2e, 0x5d, 0x2d, 0x8a,
	0x7f, 0xc0, 0xfc, 0xe1, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda, 0xa8, 0xad, 0x18, 0xc3, 0x11,
	0x00, 0x00,
}
