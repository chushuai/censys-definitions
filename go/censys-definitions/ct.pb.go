// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ct.proto

package censys_definitions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CTPushStatus int32

const (
	CTPushStatus_CT_PUSH_STATUS_RESERVED       CTPushStatus = 0
	CTPushStatus_CT_PUSH_STATUS_UNKNOWN        CTPushStatus = 1
	CTPushStatus_CT_PUSH_STATUS_SUCCESS        CTPushStatus = 2
	CTPushStatus_CT_PUSH_STATUS_UNKNOWN_ERROR  CTPushStatus = 3
	CTPushStatus_CT_PUSH_STATUS_INVALID_ROOT   CTPushStatus = 4
	CTPushStatus_CT_PUSH_STATUS_ALREADY_EXISTS CTPushStatus = 5
	CTPushStatus_CT_PUSH_STATUS_WILL_NOT_PUSH  CTPushStatus = 6
)

var CTPushStatus_name = map[int32]string{
	0: "CT_PUSH_STATUS_RESERVED",
	1: "CT_PUSH_STATUS_UNKNOWN",
	2: "CT_PUSH_STATUS_SUCCESS",
	3: "CT_PUSH_STATUS_UNKNOWN_ERROR",
	4: "CT_PUSH_STATUS_INVALID_ROOT",
	5: "CT_PUSH_STATUS_ALREADY_EXISTS",
	6: "CT_PUSH_STATUS_WILL_NOT_PUSH",
}
var CTPushStatus_value = map[string]int32{
	"CT_PUSH_STATUS_RESERVED":       0,
	"CT_PUSH_STATUS_UNKNOWN":        1,
	"CT_PUSH_STATUS_SUCCESS":        2,
	"CT_PUSH_STATUS_UNKNOWN_ERROR":  3,
	"CT_PUSH_STATUS_INVALID_ROOT":   4,
	"CT_PUSH_STATUS_ALREADY_EXISTS": 5,
	"CT_PUSH_STATUS_WILL_NOT_PUSH":  6,
}

func (x CTPushStatus) String() string {
	return proto.EnumName(CTPushStatus_name, int32(x))
}
func (CTPushStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type CTServer int32

const (
	CTServer_CT_SERVER_RESERVED CTServer = 0
	// Censys
	CTServer_CT_SERVER_CENSYS_PRODUCTION  CTServer = 1
	CTServer_CT_SERVER_CENSYS_DEVELOPMENT CTServer = 2
	// Google
	CTServer_CT_SERVER_GOOGLE_PILOT      CTServer = 11
	CTServer_CT_SERVER_GOOGLE_ROCKETEER  CTServer = 12
	CTServer_CT_SERVER_GOOGLE_SUBMARINER CTServer = 13
	CTServer_CT_SERVER_GOOGLE_TESTTUBE   CTServer = 14
	CTServer_CT_SERVER_GOOGLE_ICARUS     CTServer = 15
	CTServer_CT_SERVER_GOOGLE_SKYDIVER   CTServer = 16
	CTServer_CT_SERVER_GOOGLE_DAEDALUS   CTServer = 17
	CTServer_CT_SERVER_GOOGLE_AVIATOR    CTServer = 10
	// Google Argon
	CTServer_CT_SERVER_GOOGLE_ARGON2017 CTServer = 50
	CTServer_CT_SERVER_GOOGLE_ARGON2018 CTServer = 51
	CTServer_CT_SERVER_GOOGLE_ARGON2019 CTServer = 52
	CTServer_CT_SERVER_GOOGLE_ARGON2020 CTServer = 53
	CTServer_CT_SERVER_GOOGLE_ARGON2021 CTServer = 54
	// Symantec
	CTServer_CT_SERVER_SYMANTEC_WS_CT     CTServer = 23
	CTServer_CT_SERVER_SYMANTEC_WS_VEGA   CTServer = 24
	CTServer_CT_SERVER_SYMANTEC_WS_DENEB  CTServer = 32
	CTServer_CT_SERVER_SYMANTEC_WS_SIRIUS CTServer = 37
	// Comodo
	CTServer_CT_SERVER_COMODO_DODO    CTServer = 35
	CTServer_CT_SERVER_COMODO_MAMMOTH CTServer = 36
	CTServer_CT_SERVER_COMODO_SABRE   CTServer = 41
	// Wosign
	CTServer_CT_SERVER_WOSIGN_CTLOG CTServer = 25
	CTServer_CT_SERVER_WOSIGN_CT    CTServer = 26
	// Venafi
	CTServer_CT_SERVER_VENAFI_API_CTLOG      CTServer = 31
	CTServer_CT_SERVER_VENAFI_API_CTLOG_GEN2 CTServer = 39
	// GDCA
	CTServer_CT_SERVER_GDCA_CT    CTServer = 28
	CTServer_CT_SERVER_GDCA_CTLOG CTServer = 34
	// Izenpe
	CTServer_CT_SERVER_IZENPE_COM_CT CTServer = 21
	CTServer_CT_SERVER_IZENPE_EUS_CT CTServer = 22
	// Digicert
	CTServer_CT_SERVER_DIGICERT_CT1 CTServer = 20
	CTServer_CT_SERVER_DIGICERT_CT2 CTServer = 40
	// Other
	CTServer_CT_SERVER_CNNIC_CTSERVER                CTServer = 27
	CTServer_CT_SERVER_STARTSSL_CT                   CTServer = 29
	CTServer_CT_SERVER_CERTLY_LOG                    CTServer = 30
	CTServer_CT_SERVER_NORDU_CT_PLAUSIBLE            CTServer = 33
	CTServer_CT_SERVER_CERTIFICATETRANSPARENCY_CN_CT CTServer = 38
	CTServer_CT_SERVER_SHECA_CT                      CTServer = 42
	CTServer_CT_SERVER_LETSENCRYPT_CT_CLICKY         CTServer = 44
)

var CTServer_name = map[int32]string{
	0:  "CT_SERVER_RESERVED",
	1:  "CT_SERVER_CENSYS_PRODUCTION",
	2:  "CT_SERVER_CENSYS_DEVELOPMENT",
	11: "CT_SERVER_GOOGLE_PILOT",
	12: "CT_SERVER_GOOGLE_ROCKETEER",
	13: "CT_SERVER_GOOGLE_SUBMARINER",
	14: "CT_SERVER_GOOGLE_TESTTUBE",
	15: "CT_SERVER_GOOGLE_ICARUS",
	16: "CT_SERVER_GOOGLE_SKYDIVER",
	17: "CT_SERVER_GOOGLE_DAEDALUS",
	10: "CT_SERVER_GOOGLE_AVIATOR",
	50: "CT_SERVER_GOOGLE_ARGON2017",
	51: "CT_SERVER_GOOGLE_ARGON2018",
	52: "CT_SERVER_GOOGLE_ARGON2019",
	53: "CT_SERVER_GOOGLE_ARGON2020",
	54: "CT_SERVER_GOOGLE_ARGON2021",
	23: "CT_SERVER_SYMANTEC_WS_CT",
	24: "CT_SERVER_SYMANTEC_WS_VEGA",
	32: "CT_SERVER_SYMANTEC_WS_DENEB",
	37: "CT_SERVER_SYMANTEC_WS_SIRIUS",
	35: "CT_SERVER_COMODO_DODO",
	36: "CT_SERVER_COMODO_MAMMOTH",
	41: "CT_SERVER_COMODO_SABRE",
	25: "CT_SERVER_WOSIGN_CTLOG",
	26: "CT_SERVER_WOSIGN_CT",
	31: "CT_SERVER_VENAFI_API_CTLOG",
	39: "CT_SERVER_VENAFI_API_CTLOG_GEN2",
	28: "CT_SERVER_GDCA_CT",
	34: "CT_SERVER_GDCA_CTLOG",
	21: "CT_SERVER_IZENPE_COM_CT",
	22: "CT_SERVER_IZENPE_EUS_CT",
	20: "CT_SERVER_DIGICERT_CT1",
	40: "CT_SERVER_DIGICERT_CT2",
	27: "CT_SERVER_CNNIC_CTSERVER",
	29: "CT_SERVER_STARTSSL_CT",
	30: "CT_SERVER_CERTLY_LOG",
	33: "CT_SERVER_NORDU_CT_PLAUSIBLE",
	38: "CT_SERVER_CERTIFICATETRANSPARENCY_CN_CT",
	42: "CT_SERVER_SHECA_CT",
	44: "CT_SERVER_LETSENCRYPT_CT_CLICKY",
}
var CTServer_value = map[string]int32{
	"CT_SERVER_RESERVED":                      0,
	"CT_SERVER_CENSYS_PRODUCTION":             1,
	"CT_SERVER_CENSYS_DEVELOPMENT":            2,
	"CT_SERVER_GOOGLE_PILOT":                  11,
	"CT_SERVER_GOOGLE_ROCKETEER":              12,
	"CT_SERVER_GOOGLE_SUBMARINER":             13,
	"CT_SERVER_GOOGLE_TESTTUBE":               14,
	"CT_SERVER_GOOGLE_ICARUS":                 15,
	"CT_SERVER_GOOGLE_SKYDIVER":               16,
	"CT_SERVER_GOOGLE_DAEDALUS":               17,
	"CT_SERVER_GOOGLE_AVIATOR":                10,
	"CT_SERVER_GOOGLE_ARGON2017":              50,
	"CT_SERVER_GOOGLE_ARGON2018":              51,
	"CT_SERVER_GOOGLE_ARGON2019":              52,
	"CT_SERVER_GOOGLE_ARGON2020":              53,
	"CT_SERVER_GOOGLE_ARGON2021":              54,
	"CT_SERVER_SYMANTEC_WS_CT":                23,
	"CT_SERVER_SYMANTEC_WS_VEGA":              24,
	"CT_SERVER_SYMANTEC_WS_DENEB":             32,
	"CT_SERVER_SYMANTEC_WS_SIRIUS":            37,
	"CT_SERVER_COMODO_DODO":                   35,
	"CT_SERVER_COMODO_MAMMOTH":                36,
	"CT_SERVER_COMODO_SABRE":                  41,
	"CT_SERVER_WOSIGN_CTLOG":                  25,
	"CT_SERVER_WOSIGN_CT":                     26,
	"CT_SERVER_VENAFI_API_CTLOG":              31,
	"CT_SERVER_VENAFI_API_CTLOG_GEN2":         39,
	"CT_SERVER_GDCA_CT":                       28,
	"CT_SERVER_GDCA_CTLOG":                    34,
	"CT_SERVER_IZENPE_COM_CT":                 21,
	"CT_SERVER_IZENPE_EUS_CT":                 22,
	"CT_SERVER_DIGICERT_CT1":                  20,
	"CT_SERVER_DIGICERT_CT2":                  40,
	"CT_SERVER_CNNIC_CTSERVER":                27,
	"CT_SERVER_STARTSSL_CT":                   29,
	"CT_SERVER_CERTLY_LOG":                    30,
	"CT_SERVER_NORDU_CT_PLAUSIBLE":            33,
	"CT_SERVER_CERTIFICATETRANSPARENCY_CN_CT": 38,
	"CT_SERVER_SHECA_CT":                      42,
	"CT_SERVER_LETSENCRYPT_CT_CLICKY":         44,
}

func (x CTServer) String() string {
	return proto.EnumName(CTServer_name, int32(x))
}
func (CTServer) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type CTServerStatus struct {
	Index int64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	// parsed out timestamp inside the SCT
	CtTimestamp int64 `protobuf:"varint,2,opt,name=ct_timestamp,json=ctTimestamp" json:"ct_timestamp,omitempty"`
	// when we found the certificate in CT
	PullTimestamp int64 `protobuf:"varint,3,opt,name=pull_timestamp,json=pullTimestamp" json:"pull_timestamp,omitempty"`
	// when we pushed the record to CT
	PushTimestamp int64        `protobuf:"varint,4,opt,name=push_timestamp,json=pushTimestamp" json:"push_timestamp,omitempty"`
	PushStatus    CTPushStatus `protobuf:"varint,5,opt,name=push_status,json=pushStatus,enum=zsearch.CTPushStatus" json:"push_status,omitempty"`
	Sct           []byte       `protobuf:"bytes,6,opt,name=sct,proto3" json:"sct,omitempty"`
	PushError     string       `protobuf:"bytes,7,opt,name=push_error,json=pushError" json:"push_error,omitempty"`
}

func (m *CTServerStatus) Reset()                    { *m = CTServerStatus{} }
func (m *CTServerStatus) String() string            { return proto.CompactTextString(m) }
func (*CTServerStatus) ProtoMessage()               {}
func (*CTServerStatus) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *CTServerStatus) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *CTServerStatus) GetCtTimestamp() int64 {
	if m != nil {
		return m.CtTimestamp
	}
	return 0
}

func (m *CTServerStatus) GetPullTimestamp() int64 {
	if m != nil {
		return m.PullTimestamp
	}
	return 0
}

func (m *CTServerStatus) GetPushTimestamp() int64 {
	if m != nil {
		return m.PushTimestamp
	}
	return 0
}

func (m *CTServerStatus) GetPushStatus() CTPushStatus {
	if m != nil {
		return m.PushStatus
	}
	return CTPushStatus_CT_PUSH_STATUS_RESERVED
}

func (m *CTServerStatus) GetSct() []byte {
	if m != nil {
		return m.Sct
	}
	return nil
}

func (m *CTServerStatus) GetPushError() string {
	if m != nil {
		return m.PushError
	}
	return ""
}

type CTStatus struct {
	CensysDev                   *CTServerStatus `protobuf:"bytes,1,opt,name=censys_dev,json=censysDev" json:"censys_dev,omitempty"`
	Censys                      *CTServerStatus `protobuf:"bytes,2,opt,name=censys" json:"censys,omitempty"`
	GoogleAviator               *CTServerStatus `protobuf:"bytes,10,opt,name=google_aviator,json=googleAviator" json:"google_aviator,omitempty"`
	GooglePilot                 *CTServerStatus `protobuf:"bytes,11,opt,name=google_pilot,json=googlePilot" json:"google_pilot,omitempty"`
	GoogleRocketeer             *CTServerStatus `protobuf:"bytes,12,opt,name=google_rocketeer,json=googleRocketeer" json:"google_rocketeer,omitempty"`
	GoogleSubmariner            *CTServerStatus `protobuf:"bytes,13,opt,name=google_submariner,json=googleSubmariner" json:"google_submariner,omitempty"`
	GoogleTesttube              *CTServerStatus `protobuf:"bytes,14,opt,name=google_testtube,json=googleTesttube" json:"google_testtube,omitempty"`
	GoogleIcarus                *CTServerStatus `protobuf:"bytes,15,opt,name=google_icarus,json=googleIcarus" json:"google_icarus,omitempty"`
	GoogleSkydiver              *CTServerStatus `protobuf:"bytes,16,opt,name=google_skydiver,json=googleSkydiver" json:"google_skydiver,omitempty"`
	GoogleDaedalus              *CTServerStatus `protobuf:"bytes,17,opt,name=google_daedalus,json=googleDaedalus" json:"google_daedalus,omitempty"`
	DigicertCt1                 *CTServerStatus `protobuf:"bytes,20,opt,name=digicert_ct1,json=digicertCt1" json:"digicert_ct1,omitempty"`
	DigicertCt2                 *CTServerStatus `protobuf:"bytes,40,opt,name=digicert_ct2,json=digicertCt2" json:"digicert_ct2,omitempty"`
	IzenpeComCt                 *CTServerStatus `protobuf:"bytes,21,opt,name=izenpe_com_ct,json=izenpeComCt" json:"izenpe_com_ct,omitempty"`
	IzenpeEusCt                 *CTServerStatus `protobuf:"bytes,22,opt,name=izenpe_eus_ct,json=izenpeEusCt" json:"izenpe_eus_ct,omitempty"`
	SymantecWsCt                *CTServerStatus `protobuf:"bytes,23,opt,name=symantec_ws_ct,json=symantecWsCt" json:"symantec_ws_ct,omitempty"`
	SymantecWsVega              *CTServerStatus `protobuf:"bytes,24,opt,name=symantec_ws_vega,json=symantecWsVega" json:"symantec_ws_vega,omitempty"`
	SymantecWsSirius            *CTServerStatus `protobuf:"bytes,37,opt,name=symantec_ws_sirius,json=symantecWsSirius" json:"symantec_ws_sirius,omitempty"`
	SymantecWsDeneb             *CTServerStatus `protobuf:"bytes,32,opt,name=symantec_ws_deneb,json=symantecWsDeneb" json:"symantec_ws_deneb,omitempty"`
	ComodoDodo                  *CTServerStatus `protobuf:"bytes,34,opt,name=comodo_dodo,json=comodoDodo" json:"comodo_dodo,omitempty"`
	ComodoMammoth               *CTServerStatus `protobuf:"bytes,35,opt,name=comodo_mammoth,json=comodoMammoth" json:"comodo_mammoth,omitempty"`
	ComodoSabre                 *CTServerStatus `protobuf:"bytes,41,opt,name=comodo_sabre,json=comodoSabre" json:"comodo_sabre,omitempty"`
	WosignCtlog                 *CTServerStatus `protobuf:"bytes,25,opt,name=wosign_ctlog,json=wosignCtlog" json:"wosign_ctlog,omitempty"`
	WosignCt                    *CTServerStatus `protobuf:"bytes,26,opt,name=wosign_ct,json=wosignCt" json:"wosign_ct,omitempty"`
	GdcaCt                      *CTServerStatus `protobuf:"bytes,28,opt,name=gdca_ct,json=gdcaCt" json:"gdca_ct,omitempty"`
	GdcaCtlog                   *CTServerStatus `protobuf:"bytes,36,opt,name=gdca_ctlog,json=gdcaCtlog" json:"gdca_ctlog,omitempty"`
	VenafiApiCtlog              *CTServerStatus `protobuf:"bytes,31,opt,name=venafi_api_ctlog,json=venafiApiCtlog" json:"venafi_api_ctlog,omitempty"`
	VenafiApiCtlogGen2          *CTServerStatus `protobuf:"bytes,39,opt,name=venafi_api_ctlog_gen2,json=venafiApiCtlogGen2" json:"venafi_api_ctlog_gen2,omitempty"`
	NorduCtPlausible            *CTServerStatus `protobuf:"bytes,33,opt,name=nordu_ct_plausible,json=norduCtPlausible" json:"nordu_ct_plausible,omitempty"`
	CnnicCtserver               *CTServerStatus `protobuf:"bytes,27,opt,name=cnnic_ctserver,json=cnnicCtserver" json:"cnnic_ctserver,omitempty"`
	StartsslCt                  *CTServerStatus `protobuf:"bytes,29,opt,name=startssl_ct,json=startsslCt" json:"startssl_ct,omitempty"`
	CertlyLog                   *CTServerStatus `protobuf:"bytes,30,opt,name=certly_log,json=certlyLog" json:"certly_log,omitempty"`
	ShecaCt                     *CTServerStatus `protobuf:"bytes,42,opt,name=sheca_ct,json=shecaCt" json:"sheca_ct,omitempty"`
	CertificatetransparencyCnCt *CTServerStatus `protobuf:"bytes,38,opt,name=certificatetransparency_cn_ct,json=certificatetransparencyCnCt" json:"certificatetransparency_cn_ct,omitempty"`
	LetsencryptCtClicky         *CTServerStatus `protobuf:"bytes,43,opt,name=letsencrypt_ct_clicky,json=letsencryptCtClicky" json:"letsencrypt_ct_clicky,omitempty"`
	GoogleArgon2017             *CTServerStatus `protobuf:"bytes,50,opt,name=google_argon2017,json=googleArgon2017" json:"google_argon2017,omitempty"`
	GoogleArgon2018             *CTServerStatus `protobuf:"bytes,51,opt,name=google_argon2018,json=googleArgon2018" json:"google_argon2018,omitempty"`
	GoogleArgon2019             *CTServerStatus `protobuf:"bytes,52,opt,name=google_argon2019,json=googleArgon2019" json:"google_argon2019,omitempty"`
	GoogleArgon2020             *CTServerStatus `protobuf:"bytes,53,opt,name=google_argon2020,json=googleArgon2020" json:"google_argon2020,omitempty"`
	GoogleArgon2021             *CTServerStatus `protobuf:"bytes,54,opt,name=google_argon2021,json=googleArgon2021" json:"google_argon2021,omitempty"`
}

func (m *CTStatus) Reset()                    { *m = CTStatus{} }
func (m *CTStatus) String() string            { return proto.CompactTextString(m) }
func (*CTStatus) ProtoMessage()               {}
func (*CTStatus) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *CTStatus) GetCensysDev() *CTServerStatus {
	if m != nil {
		return m.CensysDev
	}
	return nil
}

func (m *CTStatus) GetCensys() *CTServerStatus {
	if m != nil {
		return m.Censys
	}
	return nil
}

func (m *CTStatus) GetGoogleAviator() *CTServerStatus {
	if m != nil {
		return m.GoogleAviator
	}
	return nil
}

func (m *CTStatus) GetGooglePilot() *CTServerStatus {
	if m != nil {
		return m.GooglePilot
	}
	return nil
}

func (m *CTStatus) GetGoogleRocketeer() *CTServerStatus {
	if m != nil {
		return m.GoogleRocketeer
	}
	return nil
}

func (m *CTStatus) GetGoogleSubmariner() *CTServerStatus {
	if m != nil {
		return m.GoogleSubmariner
	}
	return nil
}

func (m *CTStatus) GetGoogleTesttube() *CTServerStatus {
	if m != nil {
		return m.GoogleTesttube
	}
	return nil
}

func (m *CTStatus) GetGoogleIcarus() *CTServerStatus {
	if m != nil {
		return m.GoogleIcarus
	}
	return nil
}

func (m *CTStatus) GetGoogleSkydiver() *CTServerStatus {
	if m != nil {
		return m.GoogleSkydiver
	}
	return nil
}

func (m *CTStatus) GetGoogleDaedalus() *CTServerStatus {
	if m != nil {
		return m.GoogleDaedalus
	}
	return nil
}

func (m *CTStatus) GetDigicertCt1() *CTServerStatus {
	if m != nil {
		return m.DigicertCt1
	}
	return nil
}

func (m *CTStatus) GetDigicertCt2() *CTServerStatus {
	if m != nil {
		return m.DigicertCt2
	}
	return nil
}

func (m *CTStatus) GetIzenpeComCt() *CTServerStatus {
	if m != nil {
		return m.IzenpeComCt
	}
	return nil
}

func (m *CTStatus) GetIzenpeEusCt() *CTServerStatus {
	if m != nil {
		return m.IzenpeEusCt
	}
	return nil
}

func (m *CTStatus) GetSymantecWsCt() *CTServerStatus {
	if m != nil {
		return m.SymantecWsCt
	}
	return nil
}

func (m *CTStatus) GetSymantecWsVega() *CTServerStatus {
	if m != nil {
		return m.SymantecWsVega
	}
	return nil
}

func (m *CTStatus) GetSymantecWsSirius() *CTServerStatus {
	if m != nil {
		return m.SymantecWsSirius
	}
	return nil
}

func (m *CTStatus) GetSymantecWsDeneb() *CTServerStatus {
	if m != nil {
		return m.SymantecWsDeneb
	}
	return nil
}

func (m *CTStatus) GetComodoDodo() *CTServerStatus {
	if m != nil {
		return m.ComodoDodo
	}
	return nil
}

func (m *CTStatus) GetComodoMammoth() *CTServerStatus {
	if m != nil {
		return m.ComodoMammoth
	}
	return nil
}

func (m *CTStatus) GetComodoSabre() *CTServerStatus {
	if m != nil {
		return m.ComodoSabre
	}
	return nil
}

func (m *CTStatus) GetWosignCtlog() *CTServerStatus {
	if m != nil {
		return m.WosignCtlog
	}
	return nil
}

func (m *CTStatus) GetWosignCt() *CTServerStatus {
	if m != nil {
		return m.WosignCt
	}
	return nil
}

func (m *CTStatus) GetGdcaCt() *CTServerStatus {
	if m != nil {
		return m.GdcaCt
	}
	return nil
}

func (m *CTStatus) GetGdcaCtlog() *CTServerStatus {
	if m != nil {
		return m.GdcaCtlog
	}
	return nil
}

func (m *CTStatus) GetVenafiApiCtlog() *CTServerStatus {
	if m != nil {
		return m.VenafiApiCtlog
	}
	return nil
}

func (m *CTStatus) GetVenafiApiCtlogGen2() *CTServerStatus {
	if m != nil {
		return m.VenafiApiCtlogGen2
	}
	return nil
}

func (m *CTStatus) GetNorduCtPlausible() *CTServerStatus {
	if m != nil {
		return m.NorduCtPlausible
	}
	return nil
}

func (m *CTStatus) GetCnnicCtserver() *CTServerStatus {
	if m != nil {
		return m.CnnicCtserver
	}
	return nil
}

func (m *CTStatus) GetStartsslCt() *CTServerStatus {
	if m != nil {
		return m.StartsslCt
	}
	return nil
}

func (m *CTStatus) GetCertlyLog() *CTServerStatus {
	if m != nil {
		return m.CertlyLog
	}
	return nil
}

func (m *CTStatus) GetShecaCt() *CTServerStatus {
	if m != nil {
		return m.ShecaCt
	}
	return nil
}

func (m *CTStatus) GetCertificatetransparencyCnCt() *CTServerStatus {
	if m != nil {
		return m.CertificatetransparencyCnCt
	}
	return nil
}

func (m *CTStatus) GetLetsencryptCtClicky() *CTServerStatus {
	if m != nil {
		return m.LetsencryptCtClicky
	}
	return nil
}

func (m *CTStatus) GetGoogleArgon2017() *CTServerStatus {
	if m != nil {
		return m.GoogleArgon2017
	}
	return nil
}

func (m *CTStatus) GetGoogleArgon2018() *CTServerStatus {
	if m != nil {
		return m.GoogleArgon2018
	}
	return nil
}

func (m *CTStatus) GetGoogleArgon2019() *CTServerStatus {
	if m != nil {
		return m.GoogleArgon2019
	}
	return nil
}

func (m *CTStatus) GetGoogleArgon2020() *CTServerStatus {
	if m != nil {
		return m.GoogleArgon2020
	}
	return nil
}

func (m *CTStatus) GetGoogleArgon2021() *CTServerStatus {
	if m != nil {
		return m.GoogleArgon2021
	}
	return nil
}

type SCT struct {
	Sha256Fp []byte          `protobuf:"bytes,1,opt,name=sha256fp,proto3" json:"sha256fp,omitempty"`
	Server   CTServer        `protobuf:"varint,2,opt,name=server,enum=zsearch.CTServer" json:"server,omitempty"`
	Status   *CTServerStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *SCT) Reset()                    { *m = SCT{} }
func (m *SCT) String() string            { return proto.CompactTextString(m) }
func (*SCT) ProtoMessage()               {}
func (*SCT) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *SCT) GetSha256Fp() []byte {
	if m != nil {
		return m.Sha256Fp
	}
	return nil
}

func (m *SCT) GetServer() CTServer {
	if m != nil {
		return m.Server
	}
	return CTServer_CT_SERVER_RESERVED
}

func (m *SCT) GetStatus() *CTServerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*CTServerStatus)(nil), "zsearch.CTServerStatus")
	proto.RegisterType((*CTStatus)(nil), "zsearch.CTStatus")
	proto.RegisterType((*SCT)(nil), "zsearch.SCT")
	proto.RegisterEnum("zsearch.CTPushStatus", CTPushStatus_name, CTPushStatus_value)
	proto.RegisterEnum("zsearch.CTServer", CTServer_name, CTServer_value)
}

func init() { proto.RegisterFile("ct.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x98, 0x7f, 0x53, 0xdc, 0xb8,
	0x19, 0xc7, 0x0b, 0x5c, 0x48, 0x10, 0x3f, 0x22, 0x94, 0x10, 0x14, 0x08, 0x09, 0x21, 0x97, 0x0b,
	0xc9, 0x75, 0x52, 0xd8, 0x24, 0x34, 0xe9, 0xaf, 0xa9, 0x91, 0x75, 0x1b, 0x97, 0xc5, 0xde, 0x91,
	0xbc, 0x50, 0xda, 0x3f, 0x34, 0xc6, 0xab, 0x2c, 0x9e, 0xec, 0xda, 0x3b, 0xb6, 0x96, 0x2b, 0x37,
	0xd3, 0x99, 0xbe, 0xbd, 0xce, 0xf4, 0x9d, 0xf4, 0x4d, 0xdc, 0xc8, 0xf6, 0x26, 0x36, 0x6c, 0x6c,
	0xf8, 0x6f, 0xad, 0xe7, 0xfb, 0xf9, 0x4a, 0x7a, 0x9e, 0xc7, 0xb2, 0x66, 0xc1, 0x1d, 0x5f, 0xbd,
	0x1e, 0xc6, 0x91, 0x8a, 0xd0, 0xed, 0x5f, 0x12, 0xe9, 0xc5, 0xfe, 0xd9, 0xd6, 0x7f, 0xa6, 0xc1,
	0x12, 0x71, 0xb9, 0x8c, 0xcf, 0x65, 0xcc, 0x95, 0xa7, 0x46, 0x09, 0xba, 0x0f, 0x6e, 0x05, 0x61,
	0x57, 0xfe, 0x0b, 0x4f, 0x6d, 0x4e, 0x6d, 0xcf, 0xb0, 0xec, 0x01, 0x3d, 0x05, 0x0b, 0xbe, 0x12,
	0x2a, 0x18, 0xc8, 0x44, 0x79, 0x83, 0x21, 0x9e, 0x4e, 0x83, 0xf3, 0xbe, 0x72, 0xc7, 0x43, 0xe8,
	0x39, 0x58, 0x1a, 0x8e, 0xfa, 0xfd, 0x82, 0x68, 0x26, 0x15, 0x2d, 0xea, 0xd1, 0x4b, 0xb2, 0xe4,
	0xac, 0x20, 0xfb, 0x6e, 0x2c, 0x4b, 0xce, 0xbe, 0xca, 0xf6, 0xc0, 0x7c, 0x2a, 0x4b, 0xd2, 0x55,
	0xe1, 0x5b, 0x9b, 0x53, 0xdb, 0x4b, 0x8d, 0x95, 0xd7, 0xf9, 0xc2, 0x5f, 0x13, 0xb7, 0x3d, 0x4a,
	0xce, 0xb2, 0x25, 0x33, 0x30, 0xfc, 0xf2, 0x1b, 0x41, 0x30, 0x93, 0xf8, 0x0a, 0xcf, 0x6e, 0x4e,
	0x6d, 0x2f, 0x30, 0xfd, 0x13, 0x6d, 0x80, 0x34, 0x2e, 0x64, 0x1c, 0x47, 0x31, 0xbe, 0xbd, 0x39,
	0xb5, 0x3d, 0xc7, 0xe6, 0xf4, 0x08, 0xd5, 0x03, 0x5b, 0xff, 0xbd, 0x07, 0xee, 0x10, 0x37, 0xa7,
	0xf7, 0x00, 0xf0, 0x65, 0x98, 0x5c, 0x24, 0xa2, 0x2b, 0xcf, 0xd3, 0x0c, 0xcc, 0x37, 0x56, 0x0b,
	0x93, 0x16, 0x33, 0xc5, 0xe6, 0x32, 0xa9, 0x29, 0xcf, 0xd1, 0xef, 0xc0, 0x6c, 0xf6, 0x90, 0x26,
	0xa6, 0x82, 0xc9, 0x65, 0xe8, 0x2f, 0x60, 0xa9, 0x17, 0x45, 0xbd, 0xbe, 0x14, 0xde, 0x79, 0xe0,
	0xa9, 0x28, 0xc6, 0xa0, 0x1a, 0x5c, 0xcc, 0xe4, 0x46, 0xa6, 0x46, 0x7f, 0x00, 0x0b, 0x39, 0x3f,
	0x0c, 0xfa, 0x91, 0xc2, 0xf3, 0xd5, 0xf4, 0x7c, 0x26, 0x6e, 0x6b, 0x2d, 0xda, 0x07, 0x30, 0x67,
	0xe3, 0xc8, 0xff, 0x2c, 0x95, 0x94, 0x31, 0x5e, 0xa8, 0xe6, 0xef, 0x66, 0x00, 0x1b, 0xeb, 0x91,
	0x09, 0x96, 0x73, 0x8f, 0x64, 0x74, 0x3a, 0xf0, 0xe2, 0x20, 0x94, 0x31, 0x5e, 0xac, 0x36, 0xc9,
	0x67, 0xe5, 0x5f, 0x00, 0xf4, 0x57, 0x90, 0x1b, 0x0b, 0x25, 0x13, 0xa5, 0x46, 0xa7, 0x12, 0x2f,
	0x55, 0x7b, 0xe4, 0x59, 0x73, 0x73, 0x39, 0xfa, 0x13, 0xc8, 0x13, 0x23, 0x02, 0xdf, 0x8b, 0x47,
	0x09, 0xbe, 0x5b, 0xcd, 0xe7, 0x59, 0xb3, 0x52, 0x71, 0x61, 0xfe, 0xe4, 0xf3, 0x45, 0x37, 0x38,
	0x97, 0x31, 0x86, 0xd7, 0x9a, 0x9f, 0xe7, 0xf2, 0x82, 0x43, 0xd7, 0x93, 0x5d, 0xaf, 0x3f, 0x4a,
	0xf0, 0xf2, 0xb5, 0x1c, 0xcc, 0x5c, 0xae, 0x2b, 0xd9, 0x0d, 0x7a, 0x81, 0x2f, 0x63, 0x25, 0x7c,
	0xb5, 0x8b, 0xef, 0xd7, 0x54, 0x72, 0x2c, 0x26, 0x6a, 0xf7, 0x12, 0xdb, 0xc0, 0xdb, 0xd7, 0x66,
	0x1b, 0xe8, 0x8f, 0x60, 0x31, 0xf8, 0x45, 0x86, 0x43, 0x29, 0xfc, 0x68, 0x20, 0x7c, 0x85, 0x57,
	0x6a, 0xe0, 0x4c, 0x4d, 0xa2, 0x01, 0x51, 0x05, 0x58, 0x8e, 0x12, 0x0d, 0x3f, 0xb8, 0x16, 0x4c,
	0x47, 0x09, 0x51, 0xe8, 0xcf, 0x60, 0x29, 0xb9, 0x18, 0x78, 0xa1, 0x92, 0xbe, 0xf8, 0x39, 0xa5,
	0x57, 0x6b, 0x8a, 0x36, 0x96, 0x1f, 0x6b, 0xdc, 0x00, 0xb0, 0x88, 0x9f, 0xcb, 0x9e, 0x87, 0x71,
	0x4d, 0xce, 0xbf, 0x1a, 0x1c, 0xc9, 0x9e, 0x87, 0x28, 0x40, 0x45, 0x8b, 0x24, 0x88, 0x83, 0x51,
	0x82, 0x9f, 0xd7, 0xb4, 0xef, 0x57, 0x13, 0x9e, 0x02, 0x88, 0x80, 0xe5, 0xa2, 0x4d, 0x57, 0x86,
	0xf2, 0x14, 0x6f, 0xd6, 0xbc, 0x49, 0x5f, 0x5d, 0x4c, 0xad, 0x47, 0xef, 0xc1, 0xbc, 0x1f, 0x0d,
	0xa2, 0x6e, 0x24, 0xba, 0x51, 0x37, 0xc2, 0x5b, 0xd5, 0x38, 0xc8, 0xb4, 0x66, 0xd4, 0x8d, 0xf4,
	0x19, 0x92, 0x93, 0x03, 0x6f, 0x30, 0x88, 0xd4, 0x19, 0x7e, 0x56, 0x73, 0x86, 0x64, 0xf2, 0xc3,
	0x4c, 0xad, 0xbb, 0x27, 0xe7, 0x13, 0xef, 0x34, 0x96, 0xf8, 0x65, 0x4d, 0x0d, 0x33, 0x31, 0xd7,
	0x5a, 0xcd, 0xfe, 0x1c, 0x25, 0x41, 0x2f, 0x14, 0xbe, 0xea, 0x47, 0x3d, 0xfc, 0xb0, 0x86, 0xcd,
	0xc4, 0x44, 0x6b, 0xd1, 0x5b, 0x30, 0xf7, 0x85, 0xc5, 0x6b, 0xd5, 0xe0, 0x9d, 0x31, 0x88, 0x76,
	0xc0, 0xed, 0x5e, 0xd7, 0xf7, 0x34, 0xf3, 0xa8, 0xe6, 0x8c, 0xd5, 0x3a, 0xa2, 0xf4, 0x61, 0x9e,
	0x13, 0x7a, 0x85, 0xdf, 0xd7, 0x1c, 0xe6, 0x19, 0xa4, 0xd7, 0x67, 0x00, 0x78, 0x2e, 0x43, 0xef,
	0x53, 0x20, 0xbc, 0x61, 0x90, 0xd3, 0x4f, 0x6a, 0x1a, 0x2c, 0x03, 0x8c, 0x61, 0x90, 0x59, 0xfc,
	0x0d, 0xac, 0x5c, 0xb6, 0x10, 0x3d, 0x19, 0x36, 0xf0, 0x8b, 0x6a, 0x1f, 0x54, 0xf6, 0x69, 0xca,
	0xb0, 0xa1, 0x9b, 0x35, 0x8c, 0xe2, 0xee, 0x48, 0xf8, 0x4a, 0x0c, 0xfb, 0xde, 0x28, 0x09, 0x4e,
	0xfb, 0x12, 0x3f, 0xad, 0x69, 0xd6, 0x14, 0x21, 0xaa, 0x3d, 0x06, 0xd2, 0x6e, 0x09, 0xc3, 0xc0,
	0x17, 0xbe, 0x4a, 0x52, 0x25, 0x5e, 0xaf, 0xeb, 0x16, 0x2d, 0x27, 0xb9, 0x5a, 0xf7, 0x69, 0xa2,
	0xbc, 0x58, 0x25, 0x49, 0x5f, 0xd7, 0x60, 0xa3, 0xa6, 0x4f, 0xc7, 0xda, 0xac, 0x0e, 0xfa, 0xd0,
	0xe9, 0x5f, 0x08, 0x9d, 0xc9, 0xc7, 0xb5, 0x1f, 0x55, 0x2d, 0x6d, 0x45, 0x3d, 0xd4, 0x00, 0x77,
	0x92, 0x33, 0x99, 0x95, 0xfc, 0x55, 0x35, 0x75, 0x3b, 0x15, 0x12, 0x85, 0xfe, 0x09, 0x36, 0xb4,
	0x41, 0xf0, 0x29, 0xf0, 0x3d, 0x25, 0x55, 0xec, 0x85, 0xc9, 0xd0, 0x8b, 0x65, 0xe8, 0x5f, 0x08,
	0x3f, 0xed, 0xb7, 0x1f, 0xaa, 0x8d, 0xd6, 0xbf, 0x41, 0x13, 0xdd, 0x82, 0x07, 0x60, 0xa5, 0x2f,
	0x55, 0x22, 0x43, 0x3f, 0xbe, 0x18, 0xea, 0x13, 0x57, 0xf8, 0xfd, 0xc0, 0xff, 0x7c, 0x81, 0x7f,
	0xac, 0x36, 0xbd, 0x57, 0xa0, 0x88, 0x22, 0x29, 0x53, 0xf8, 0x0a, 0x7b, 0x71, 0x2f, 0x0a, 0x1b,
	0x3b, 0xbb, 0xbf, 0xc7, 0x8d, 0x6b, 0x7d, 0x85, 0x8d, 0xb1, 0x7e, 0x82, 0xc7, 0x7b, 0xfc, 0xe6,
	0x46, 0x1e, 0xef, 0x27, 0x78, 0x7c, 0xc0, 0x6f, 0x6f, 0xe4, 0xf1, 0xe1, 0xaa, 0x47, 0x63, 0x07,
	0xbf, 0xbb, 0x89, 0x47, 0x63, 0x67, 0x82, 0xc7, 0x2e, 0xde, 0xbb, 0x91, 0xc7, 0xee, 0xd6, 0xbf,
	0xc1, 0x0c, 0x27, 0x2e, 0x5a, 0xd3, 0x8d, 0xe3, 0x35, 0xde, 0xed, 0x7d, 0x1a, 0xa6, 0x77, 0xb8,
	0x05, 0xf6, 0xe5, 0x19, 0xbd, 0x04, 0xb3, 0x79, 0xfb, 0x4f, 0xa7, 0x57, 0xca, 0xe5, 0x2b, 0xe6,
	0x2c, 0x17, 0xe8, 0x4b, 0x5d, 0x7e, 0xfb, 0x9c, 0xa9, 0x39, 0x70, 0x32, 0xd9, 0xab, 0xff, 0x4f,
	0x81, 0x85, 0xe2, 0xc5, 0x14, 0xad, 0x83, 0x55, 0xe2, 0x8a, 0x76, 0x87, 0x7f, 0x14, 0xdc, 0x35,
	0xdc, 0x0e, 0x17, 0x8c, 0x72, 0xca, 0x8e, 0xa8, 0x09, 0x7f, 0x83, 0xd6, 0xc0, 0x83, 0x4b, 0xc1,
	0x8e, 0x7d, 0x60, 0x3b, 0xc7, 0x36, 0x9c, 0x9a, 0x10, 0xe3, 0x1d, 0x42, 0x28, 0xe7, 0x70, 0x1a,
	0x6d, 0x82, 0x47, 0x93, 0x39, 0x41, 0x19, 0x73, 0x18, 0x9c, 0x41, 0x4f, 0xc0, 0xfa, 0x25, 0x85,
	0x65, 0x1f, 0x19, 0x2d, 0xcb, 0x14, 0xcc, 0x71, 0x5c, 0xf8, 0x1d, 0x7a, 0x0a, 0x36, 0x2e, 0x09,
	0x8c, 0x16, 0xa3, 0x86, 0x79, 0x22, 0xe8, 0xdf, 0x2d, 0xee, 0x72, 0x78, 0x6b, 0xc2, 0x2c, 0xc7,
	0x56, 0xab, 0x25, 0x6c, 0x27, 0x1b, 0x84, 0xb3, 0xaf, 0xfe, 0x37, 0x97, 0x5e, 0x9c, 0xb3, 0x5c,
	0x3d, 0x00, 0x88, 0xb8, 0x22, 0xdd, 0x1b, 0x2b, 0x6e, 0x32, 0x5b, 0x4a, 0x3e, 0x4e, 0xa8, 0xcd,
	0x4f, 0xb8, 0x68, 0x33, 0xc7, 0xec, 0x10, 0xd7, 0x72, 0xf4, 0x4e, 0xb3, 0x79, 0xca, 0x02, 0x93,
	0x1e, 0xd1, 0x96, 0xd3, 0x3e, 0xa4, 0xb6, 0x0b, 0xa7, 0xf3, 0x5c, 0xe4, 0x8a, 0xa6, 0xe3, 0x34,
	0x5b, 0x54, 0xb4, 0xad, 0x96, 0xe3, 0xc2, 0x79, 0xf4, 0x18, 0xac, 0x5d, 0x89, 0x31, 0x87, 0x1c,
	0x50, 0x97, 0x52, 0x06, 0x17, 0xca, 0xd3, 0xe7, 0x71, 0xde, 0xd9, 0x3f, 0x34, 0x98, 0x65, 0x53,
	0x06, 0x17, 0xd1, 0x06, 0x78, 0x78, 0x45, 0xe0, 0x52, 0xee, 0xba, 0x9d, 0x7d, 0x0a, 0x97, 0xf2,
	0x02, 0x96, 0xc3, 0x16, 0x31, 0x58, 0x87, 0xc3, 0xbb, 0x13, 0x59, 0x7e, 0x70, 0x62, 0x5a, 0x47,
	0x94, 0x41, 0x38, 0x31, 0x6c, 0x1a, 0xd4, 0x34, 0x5a, 0x1d, 0x0e, 0x97, 0xd1, 0x23, 0x80, 0xaf,
	0x84, 0x8d, 0x23, 0xcb, 0x70, 0x1d, 0x06, 0xc1, 0xc4, 0x8d, 0x19, 0xac, 0xe9, 0xd8, 0xfa, 0xbd,
	0x87, 0x8d, 0xca, 0xf8, 0x7b, 0xf8, 0xa6, 0x32, 0xfe, 0x01, 0xbe, 0xad, 0x8a, 0x37, 0x76, 0xe0,
	0xbb, 0xca, 0xf8, 0x2e, 0xdc, 0x2b, 0xaf, 0x9e, 0x9f, 0x1c, 0x1a, 0xb6, 0x4b, 0x89, 0x38, 0xe6,
	0x82, 0xb8, 0x70, 0xb5, 0x4c, 0x17, 0xa3, 0x47, 0xb4, 0x69, 0x40, 0x5c, 0x2e, 0x4b, 0x31, 0x6e,
	0x52, 0x9b, 0xee, 0xc3, 0xcd, 0x72, 0x57, 0x14, 0x05, 0xdc, 0x62, 0x56, 0x87, 0xc3, 0xe7, 0xe8,
	0x21, 0x58, 0x29, 0xf4, 0x8d, 0x73, 0xe8, 0x98, 0x8e, 0x30, 0x1d, 0xd3, 0x81, 0xcf, 0xca, 0x6b,
	0xcb, 0x43, 0x87, 0xc6, 0xe1, 0xa1, 0xe3, 0x7e, 0x84, 0xdf, 0x97, 0xdb, 0x29, 0x8f, 0x72, 0x63,
	0x9f, 0x51, 0xf8, 0xb2, 0x1c, 0x3b, 0x76, 0xb8, 0xd5, 0xb4, 0x05, 0x71, 0x5b, 0x4e, 0x13, 0x3e,
	0x44, 0xab, 0xe0, 0xde, 0x84, 0x18, 0x5c, 0x2b, 0x6f, 0xf6, 0x88, 0xda, 0xc6, 0x4f, 0x96, 0x30,
	0xda, 0x56, 0x0e, 0x3e, 0x41, 0xcf, 0xc0, 0x93, 0x6f, 0xc7, 0x45, 0x93, 0xda, 0x0d, 0xf8, 0x02,
	0xad, 0x80, 0xe5, 0x42, 0xbe, 0x4d, 0x62, 0x68, 0xef, 0x47, 0x08, 0x83, 0xfb, 0x57, 0x86, 0xb5,
	0xeb, 0x56, 0xb9, 0x33, 0xad, 0x7f, 0x50, 0xbb, 0x4d, 0xf5, 0x6e, 0x34, 0xb6, 0x32, 0x31, 0x48,
	0x3b, 0x69, 0x71, 0x1e, 0x94, 0x37, 0x69, 0x5a, 0x4d, 0x8b, 0x50, 0xe6, 0x0a, 0xe2, 0xee, 0xc2,
	0xfb, 0xdf, 0x8c, 0x35, 0xe0, 0xf6, 0xa5, 0xb4, 0xda, 0xb6, 0x45, 0x04, 0x71, 0xb3, 0x47, 0xb8,
	0x5e, 0xae, 0x07, 0x77, 0x0d, 0xe6, 0x72, 0xde, 0xd2, 0x13, 0x6e, 0x94, 0x37, 0xa1, 0x0d, 0x5b,
	0x27, 0x42, 0x6f, 0xe2, 0x71, 0xb9, 0xcc, 0xb6, 0xc3, 0xcc, 0x8e, 0xd0, 0x87, 0x4e, 0xcb, 0xe8,
	0x70, 0x6b, 0xbf, 0x45, 0xe1, 0x53, 0xf4, 0x23, 0x78, 0x51, 0x66, 0xad, 0x9f, 0x2c, 0x62, 0xb8,
	0xd4, 0x65, 0x86, 0xcd, 0xdb, 0x06, 0xa3, 0x36, 0x39, 0x11, 0x24, 0xad, 0xc4, 0x0f, 0xe5, 0x43,
	0x88, 0x7f, 0xa4, 0x59, 0x16, 0x5f, 0x95, 0x2b, 0xd0, 0xa2, 0x2e, 0xa7, 0x36, 0x61, 0x27, 0x6d,
	0xbd, 0x31, 0x41, 0x5a, 0x16, 0x39, 0x38, 0x81, 0xbf, 0x3d, 0x9d, 0x4d, 0xff, 0x1a, 0x79, 0xf3,
	0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x91, 0x0e, 0xab, 0x26, 0x11, 0x00, 0x00,
}
