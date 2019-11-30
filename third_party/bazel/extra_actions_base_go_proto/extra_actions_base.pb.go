// Code generated by protoc-gen-go. DO NOT EDIT.
// source: third_party/bazel/src/main/protobuf/extra_actions_base.proto

package blaze

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ExtraActionSummary struct {
	Action               []*DetailedExtraActionInfo `protobuf:"bytes,1,rep,name=action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ExtraActionSummary) Reset()         { *m = ExtraActionSummary{} }
func (m *ExtraActionSummary) String() string { return proto.CompactTextString(m) }
func (*ExtraActionSummary) ProtoMessage()    {}
func (*ExtraActionSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{0}
}

func (m *ExtraActionSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraActionSummary.Unmarshal(m, b)
}
func (m *ExtraActionSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraActionSummary.Marshal(b, m, deterministic)
}
func (m *ExtraActionSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraActionSummary.Merge(m, src)
}
func (m *ExtraActionSummary) XXX_Size() int {
	return xxx_messageInfo_ExtraActionSummary.Size(m)
}
func (m *ExtraActionSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraActionSummary.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraActionSummary proto.InternalMessageInfo

func (m *ExtraActionSummary) GetAction() []*DetailedExtraActionInfo {
	if m != nil {
		return m.Action
	}
	return nil
}

type DetailedExtraActionInfo struct {
	TriggeringFile       *string          `protobuf:"bytes,1,opt,name=triggering_file,json=triggeringFile" json:"triggering_file,omitempty"`
	Action               *ExtraActionInfo `protobuf:"bytes,2,req,name=action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DetailedExtraActionInfo) Reset()         { *m = DetailedExtraActionInfo{} }
func (m *DetailedExtraActionInfo) String() string { return proto.CompactTextString(m) }
func (*DetailedExtraActionInfo) ProtoMessage()    {}
func (*DetailedExtraActionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{1}
}

func (m *DetailedExtraActionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailedExtraActionInfo.Unmarshal(m, b)
}
func (m *DetailedExtraActionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailedExtraActionInfo.Marshal(b, m, deterministic)
}
func (m *DetailedExtraActionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailedExtraActionInfo.Merge(m, src)
}
func (m *DetailedExtraActionInfo) XXX_Size() int {
	return xxx_messageInfo_DetailedExtraActionInfo.Size(m)
}
func (m *DetailedExtraActionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailedExtraActionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DetailedExtraActionInfo proto.InternalMessageInfo

func (m *DetailedExtraActionInfo) GetTriggeringFile() string {
	if m != nil && m.TriggeringFile != nil {
		return *m.TriggeringFile
	}
	return ""
}

func (m *DetailedExtraActionInfo) GetAction() *ExtraActionInfo {
	if m != nil {
		return m.Action
	}
	return nil
}

type ExtraActionInfo struct {
	Owner                        *string                                `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	AspectName                   *string                                `protobuf:"bytes,6,opt,name=aspect_name,json=aspectName" json:"aspect_name,omitempty"`                                                                                             // Deprecated: Do not use.
	AspectParameters             map[string]*ExtraActionInfo_StringList `protobuf:"bytes,7,rep,name=aspect_parameters,json=aspectParameters" json:"aspect_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // Deprecated: Do not use.
	Aspects                      []*ExtraActionInfo_AspectDescriptor    `protobuf:"bytes,8,rep,name=aspects" json:"aspects,omitempty"`
	Id                           *string                                `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Mnemonic                     *string                                `protobuf:"bytes,5,opt,name=mnemonic" json:"mnemonic,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                               `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *ExtraActionInfo) Reset()         { *m = ExtraActionInfo{} }
func (m *ExtraActionInfo) String() string { return proto.CompactTextString(m) }
func (*ExtraActionInfo) ProtoMessage()    {}
func (*ExtraActionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{2}
}

var extRange_ExtraActionInfo = []proto.ExtensionRange{
	{Start: 1000, End: 536870911},
}

func (*ExtraActionInfo) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ExtraActionInfo
}

func (m *ExtraActionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraActionInfo.Unmarshal(m, b)
}
func (m *ExtraActionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraActionInfo.Marshal(b, m, deterministic)
}
func (m *ExtraActionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraActionInfo.Merge(m, src)
}
func (m *ExtraActionInfo) XXX_Size() int {
	return xxx_messageInfo_ExtraActionInfo.Size(m)
}
func (m *ExtraActionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraActionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraActionInfo proto.InternalMessageInfo

func (m *ExtraActionInfo) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

// Deprecated: Do not use.
func (m *ExtraActionInfo) GetAspectName() string {
	if m != nil && m.AspectName != nil {
		return *m.AspectName
	}
	return ""
}

// Deprecated: Do not use.
func (m *ExtraActionInfo) GetAspectParameters() map[string]*ExtraActionInfo_StringList {
	if m != nil {
		return m.AspectParameters
	}
	return nil
}

func (m *ExtraActionInfo) GetAspects() []*ExtraActionInfo_AspectDescriptor {
	if m != nil {
		return m.Aspects
	}
	return nil
}

func (m *ExtraActionInfo) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *ExtraActionInfo) GetMnemonic() string {
	if m != nil && m.Mnemonic != nil {
		return *m.Mnemonic
	}
	return ""
}

// Deprecated: Do not use.
type ExtraActionInfo_StringList struct {
	Value                []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraActionInfo_StringList) Reset()         { *m = ExtraActionInfo_StringList{} }
func (m *ExtraActionInfo_StringList) String() string { return proto.CompactTextString(m) }
func (*ExtraActionInfo_StringList) ProtoMessage()    {}
func (*ExtraActionInfo_StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{2, 1}
}

func (m *ExtraActionInfo_StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraActionInfo_StringList.Unmarshal(m, b)
}
func (m *ExtraActionInfo_StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraActionInfo_StringList.Marshal(b, m, deterministic)
}
func (m *ExtraActionInfo_StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraActionInfo_StringList.Merge(m, src)
}
func (m *ExtraActionInfo_StringList) XXX_Size() int {
	return xxx_messageInfo_ExtraActionInfo_StringList.Size(m)
}
func (m *ExtraActionInfo_StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraActionInfo_StringList.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraActionInfo_StringList proto.InternalMessageInfo

func (m *ExtraActionInfo_StringList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type ExtraActionInfo_AspectDescriptor struct {
	AspectName           *string                                                 `protobuf:"bytes,1,opt,name=aspect_name,json=aspectName" json:"aspect_name,omitempty"`
	AspectParameters     map[string]*ExtraActionInfo_AspectDescriptor_StringList `protobuf:"bytes,2,rep,name=aspect_parameters,json=aspectParameters" json:"aspect_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *ExtraActionInfo_AspectDescriptor) Reset()         { *m = ExtraActionInfo_AspectDescriptor{} }
func (m *ExtraActionInfo_AspectDescriptor) String() string { return proto.CompactTextString(m) }
func (*ExtraActionInfo_AspectDescriptor) ProtoMessage()    {}
func (*ExtraActionInfo_AspectDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{2, 2}
}

func (m *ExtraActionInfo_AspectDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor.Unmarshal(m, b)
}
func (m *ExtraActionInfo_AspectDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor.Marshal(b, m, deterministic)
}
func (m *ExtraActionInfo_AspectDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraActionInfo_AspectDescriptor.Merge(m, src)
}
func (m *ExtraActionInfo_AspectDescriptor) XXX_Size() int {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor.Size(m)
}
func (m *ExtraActionInfo_AspectDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraActionInfo_AspectDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraActionInfo_AspectDescriptor proto.InternalMessageInfo

func (m *ExtraActionInfo_AspectDescriptor) GetAspectName() string {
	if m != nil && m.AspectName != nil {
		return *m.AspectName
	}
	return ""
}

func (m *ExtraActionInfo_AspectDescriptor) GetAspectParameters() map[string]*ExtraActionInfo_AspectDescriptor_StringList {
	if m != nil {
		return m.AspectParameters
	}
	return nil
}

type ExtraActionInfo_AspectDescriptor_StringList struct {
	Value                []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraActionInfo_AspectDescriptor_StringList) Reset() {
	*m = ExtraActionInfo_AspectDescriptor_StringList{}
}
func (m *ExtraActionInfo_AspectDescriptor_StringList) String() string {
	return proto.CompactTextString(m)
}
func (*ExtraActionInfo_AspectDescriptor_StringList) ProtoMessage() {}
func (*ExtraActionInfo_AspectDescriptor_StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{2, 2, 1}
}

func (m *ExtraActionInfo_AspectDescriptor_StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList.Unmarshal(m, b)
}
func (m *ExtraActionInfo_AspectDescriptor_StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList.Marshal(b, m, deterministic)
}
func (m *ExtraActionInfo_AspectDescriptor_StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList.Merge(m, src)
}
func (m *ExtraActionInfo_AspectDescriptor_StringList) XXX_Size() int {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList.Size(m)
}
func (m *ExtraActionInfo_AspectDescriptor_StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList proto.InternalMessageInfo

func (m *ExtraActionInfo_AspectDescriptor_StringList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type EnvironmentVariable struct {
	Name                 *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvironmentVariable) Reset()         { *m = EnvironmentVariable{} }
func (m *EnvironmentVariable) String() string { return proto.CompactTextString(m) }
func (*EnvironmentVariable) ProtoMessage()    {}
func (*EnvironmentVariable) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{3}
}

func (m *EnvironmentVariable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironmentVariable.Unmarshal(m, b)
}
func (m *EnvironmentVariable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironmentVariable.Marshal(b, m, deterministic)
}
func (m *EnvironmentVariable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironmentVariable.Merge(m, src)
}
func (m *EnvironmentVariable) XXX_Size() int {
	return xxx_messageInfo_EnvironmentVariable.Size(m)
}
func (m *EnvironmentVariable) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironmentVariable.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironmentVariable proto.InternalMessageInfo

func (m *EnvironmentVariable) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *EnvironmentVariable) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type SpawnInfo struct {
	Argument             []string               `protobuf:"bytes,1,rep,name=argument" json:"argument,omitempty"`
	Variable             []*EnvironmentVariable `protobuf:"bytes,2,rep,name=variable" json:"variable,omitempty"`
	InputFile            []string               `protobuf:"bytes,4,rep,name=input_file,json=inputFile" json:"input_file,omitempty"`
	OutputFile           []string               `protobuf:"bytes,5,rep,name=output_file,json=outputFile" json:"output_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SpawnInfo) Reset()         { *m = SpawnInfo{} }
func (m *SpawnInfo) String() string { return proto.CompactTextString(m) }
func (*SpawnInfo) ProtoMessage()    {}
func (*SpawnInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{4}
}

func (m *SpawnInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnInfo.Unmarshal(m, b)
}
func (m *SpawnInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnInfo.Marshal(b, m, deterministic)
}
func (m *SpawnInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnInfo.Merge(m, src)
}
func (m *SpawnInfo) XXX_Size() int {
	return xxx_messageInfo_SpawnInfo.Size(m)
}
func (m *SpawnInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnInfo proto.InternalMessageInfo

func (m *SpawnInfo) GetArgument() []string {
	if m != nil {
		return m.Argument
	}
	return nil
}

func (m *SpawnInfo) GetVariable() []*EnvironmentVariable {
	if m != nil {
		return m.Variable
	}
	return nil
}

func (m *SpawnInfo) GetInputFile() []string {
	if m != nil {
		return m.InputFile
	}
	return nil
}

func (m *SpawnInfo) GetOutputFile() []string {
	if m != nil {
		return m.OutputFile
	}
	return nil
}

var E_SpawnInfo_SpawnInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*SpawnInfo)(nil),
	Field:         1003,
	Name:          "blaze.SpawnInfo.spawn_info",
	Tag:           "bytes,1003,opt,name=spawn_info",
	Filename:      "third_party/bazel/src/main/protobuf/extra_actions_base.proto",
}

type CppCompileInfo struct {
	Tool                 *string                `protobuf:"bytes,1,opt,name=tool" json:"tool,omitempty"`
	CompilerOption       []string               `protobuf:"bytes,2,rep,name=compiler_option,json=compilerOption" json:"compiler_option,omitempty"`
	SourceFile           *string                `protobuf:"bytes,3,opt,name=source_file,json=sourceFile" json:"source_file,omitempty"`
	OutputFile           *string                `protobuf:"bytes,4,opt,name=output_file,json=outputFile" json:"output_file,omitempty"`
	SourcesAndHeaders    []string               `protobuf:"bytes,5,rep,name=sources_and_headers,json=sourcesAndHeaders" json:"sources_and_headers,omitempty"`
	Variable             []*EnvironmentVariable `protobuf:"bytes,6,rep,name=variable" json:"variable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CppCompileInfo) Reset()         { *m = CppCompileInfo{} }
func (m *CppCompileInfo) String() string { return proto.CompactTextString(m) }
func (*CppCompileInfo) ProtoMessage()    {}
func (*CppCompileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{5}
}

func (m *CppCompileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CppCompileInfo.Unmarshal(m, b)
}
func (m *CppCompileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CppCompileInfo.Marshal(b, m, deterministic)
}
func (m *CppCompileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CppCompileInfo.Merge(m, src)
}
func (m *CppCompileInfo) XXX_Size() int {
	return xxx_messageInfo_CppCompileInfo.Size(m)
}
func (m *CppCompileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CppCompileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CppCompileInfo proto.InternalMessageInfo

func (m *CppCompileInfo) GetTool() string {
	if m != nil && m.Tool != nil {
		return *m.Tool
	}
	return ""
}

func (m *CppCompileInfo) GetCompilerOption() []string {
	if m != nil {
		return m.CompilerOption
	}
	return nil
}

func (m *CppCompileInfo) GetSourceFile() string {
	if m != nil && m.SourceFile != nil {
		return *m.SourceFile
	}
	return ""
}

func (m *CppCompileInfo) GetOutputFile() string {
	if m != nil && m.OutputFile != nil {
		return *m.OutputFile
	}
	return ""
}

func (m *CppCompileInfo) GetSourcesAndHeaders() []string {
	if m != nil {
		return m.SourcesAndHeaders
	}
	return nil
}

func (m *CppCompileInfo) GetVariable() []*EnvironmentVariable {
	if m != nil {
		return m.Variable
	}
	return nil
}

var E_CppCompileInfo_CppCompileInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*CppCompileInfo)(nil),
	Field:         1001,
	Name:          "blaze.CppCompileInfo.cpp_compile_info",
	Tag:           "bytes,1001,opt,name=cpp_compile_info",
	Filename:      "third_party/bazel/src/main/protobuf/extra_actions_base.proto",
}

type CppLinkInfo struct {
	InputFile               []string `protobuf:"bytes,1,rep,name=input_file,json=inputFile" json:"input_file,omitempty"`
	OutputFile              *string  `protobuf:"bytes,2,opt,name=output_file,json=outputFile" json:"output_file,omitempty"`
	InterfaceOutputFile     *string  `protobuf:"bytes,3,opt,name=interface_output_file,json=interfaceOutputFile" json:"interface_output_file,omitempty"`
	LinkTargetType          *string  `protobuf:"bytes,4,opt,name=link_target_type,json=linkTargetType" json:"link_target_type,omitempty"`
	LinkStaticness          *string  `protobuf:"bytes,5,opt,name=link_staticness,json=linkStaticness" json:"link_staticness,omitempty"`
	LinkStamp               []string `protobuf:"bytes,6,rep,name=link_stamp,json=linkStamp" json:"link_stamp,omitempty"`
	BuildInfoHeaderArtifact []string `protobuf:"bytes,7,rep,name=build_info_header_artifact,json=buildInfoHeaderArtifact" json:"build_info_header_artifact,omitempty"`
	LinkOpt                 []string `protobuf:"bytes,8,rep,name=link_opt,json=linkOpt" json:"link_opt,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *CppLinkInfo) Reset()         { *m = CppLinkInfo{} }
func (m *CppLinkInfo) String() string { return proto.CompactTextString(m) }
func (*CppLinkInfo) ProtoMessage()    {}
func (*CppLinkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{6}
}

func (m *CppLinkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CppLinkInfo.Unmarshal(m, b)
}
func (m *CppLinkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CppLinkInfo.Marshal(b, m, deterministic)
}
func (m *CppLinkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CppLinkInfo.Merge(m, src)
}
func (m *CppLinkInfo) XXX_Size() int {
	return xxx_messageInfo_CppLinkInfo.Size(m)
}
func (m *CppLinkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CppLinkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CppLinkInfo proto.InternalMessageInfo

func (m *CppLinkInfo) GetInputFile() []string {
	if m != nil {
		return m.InputFile
	}
	return nil
}

func (m *CppLinkInfo) GetOutputFile() string {
	if m != nil && m.OutputFile != nil {
		return *m.OutputFile
	}
	return ""
}

func (m *CppLinkInfo) GetInterfaceOutputFile() string {
	if m != nil && m.InterfaceOutputFile != nil {
		return *m.InterfaceOutputFile
	}
	return ""
}

func (m *CppLinkInfo) GetLinkTargetType() string {
	if m != nil && m.LinkTargetType != nil {
		return *m.LinkTargetType
	}
	return ""
}

func (m *CppLinkInfo) GetLinkStaticness() string {
	if m != nil && m.LinkStaticness != nil {
		return *m.LinkStaticness
	}
	return ""
}

func (m *CppLinkInfo) GetLinkStamp() []string {
	if m != nil {
		return m.LinkStamp
	}
	return nil
}

func (m *CppLinkInfo) GetBuildInfoHeaderArtifact() []string {
	if m != nil {
		return m.BuildInfoHeaderArtifact
	}
	return nil
}

func (m *CppLinkInfo) GetLinkOpt() []string {
	if m != nil {
		return m.LinkOpt
	}
	return nil
}

var E_CppLinkInfo_CppLinkInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*CppLinkInfo)(nil),
	Field:         1002,
	Name:          "blaze.CppLinkInfo.cpp_link_info",
	Tag:           "bytes,1002,opt,name=cpp_link_info",
	Filename:      "third_party/bazel/src/main/protobuf/extra_actions_base.proto",
}

type JavaCompileInfo struct {
	Outputjar            *string  `protobuf:"bytes,1,opt,name=outputjar" json:"outputjar,omitempty"`
	Classpath            []string `protobuf:"bytes,2,rep,name=classpath" json:"classpath,omitempty"`
	Sourcepath           []string `protobuf:"bytes,3,rep,name=sourcepath" json:"sourcepath,omitempty"`
	SourceFile           []string `protobuf:"bytes,4,rep,name=source_file,json=sourceFile" json:"source_file,omitempty"`
	JavacOpt             []string `protobuf:"bytes,5,rep,name=javac_opt,json=javacOpt" json:"javac_opt,omitempty"`
	Processor            []string `protobuf:"bytes,6,rep,name=processor" json:"processor,omitempty"`
	Processorpath        []string `protobuf:"bytes,7,rep,name=processorpath" json:"processorpath,omitempty"`
	Bootclasspath        []string `protobuf:"bytes,8,rep,name=bootclasspath" json:"bootclasspath,omitempty"`
	Argument             []string `protobuf:"bytes,9,rep,name=argument" json:"argument,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JavaCompileInfo) Reset()         { *m = JavaCompileInfo{} }
func (m *JavaCompileInfo) String() string { return proto.CompactTextString(m) }
func (*JavaCompileInfo) ProtoMessage()    {}
func (*JavaCompileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{7}
}

func (m *JavaCompileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JavaCompileInfo.Unmarshal(m, b)
}
func (m *JavaCompileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JavaCompileInfo.Marshal(b, m, deterministic)
}
func (m *JavaCompileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JavaCompileInfo.Merge(m, src)
}
func (m *JavaCompileInfo) XXX_Size() int {
	return xxx_messageInfo_JavaCompileInfo.Size(m)
}
func (m *JavaCompileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_JavaCompileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_JavaCompileInfo proto.InternalMessageInfo

func (m *JavaCompileInfo) GetOutputjar() string {
	if m != nil && m.Outputjar != nil {
		return *m.Outputjar
	}
	return ""
}

func (m *JavaCompileInfo) GetClasspath() []string {
	if m != nil {
		return m.Classpath
	}
	return nil
}

func (m *JavaCompileInfo) GetSourcepath() []string {
	if m != nil {
		return m.Sourcepath
	}
	return nil
}

func (m *JavaCompileInfo) GetSourceFile() []string {
	if m != nil {
		return m.SourceFile
	}
	return nil
}

func (m *JavaCompileInfo) GetJavacOpt() []string {
	if m != nil {
		return m.JavacOpt
	}
	return nil
}

func (m *JavaCompileInfo) GetProcessor() []string {
	if m != nil {
		return m.Processor
	}
	return nil
}

func (m *JavaCompileInfo) GetProcessorpath() []string {
	if m != nil {
		return m.Processorpath
	}
	return nil
}

func (m *JavaCompileInfo) GetBootclasspath() []string {
	if m != nil {
		return m.Bootclasspath
	}
	return nil
}

func (m *JavaCompileInfo) GetArgument() []string {
	if m != nil {
		return m.Argument
	}
	return nil
}

var E_JavaCompileInfo_JavaCompileInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*JavaCompileInfo)(nil),
	Field:         1000,
	Name:          "blaze.JavaCompileInfo.java_compile_info",
	Tag:           "bytes,1000,opt,name=java_compile_info",
	Filename:      "third_party/bazel/src/main/protobuf/extra_actions_base.proto",
}

type PythonInfo struct {
	SourceFile           []string `protobuf:"bytes,1,rep,name=source_file,json=sourceFile" json:"source_file,omitempty"`
	DepFile              []string `protobuf:"bytes,2,rep,name=dep_file,json=depFile" json:"dep_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PythonInfo) Reset()         { *m = PythonInfo{} }
func (m *PythonInfo) String() string { return proto.CompactTextString(m) }
func (*PythonInfo) ProtoMessage()    {}
func (*PythonInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b6b55463e685df, []int{8}
}

func (m *PythonInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PythonInfo.Unmarshal(m, b)
}
func (m *PythonInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PythonInfo.Marshal(b, m, deterministic)
}
func (m *PythonInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PythonInfo.Merge(m, src)
}
func (m *PythonInfo) XXX_Size() int {
	return xxx_messageInfo_PythonInfo.Size(m)
}
func (m *PythonInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PythonInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PythonInfo proto.InternalMessageInfo

func (m *PythonInfo) GetSourceFile() []string {
	if m != nil {
		return m.SourceFile
	}
	return nil
}

func (m *PythonInfo) GetDepFile() []string {
	if m != nil {
		return m.DepFile
	}
	return nil
}

var E_PythonInfo_PythonInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*PythonInfo)(nil),
	Field:         1005,
	Name:          "blaze.PythonInfo.python_info",
	Tag:           "bytes,1005,opt,name=python_info",
	Filename:      "third_party/bazel/src/main/protobuf/extra_actions_base.proto",
}

func init() {
	proto.RegisterType((*ExtraActionSummary)(nil), "blaze.ExtraActionSummary")
	proto.RegisterType((*DetailedExtraActionInfo)(nil), "blaze.DetailedExtraActionInfo")
	proto.RegisterType((*ExtraActionInfo)(nil), "blaze.ExtraActionInfo")
	proto.RegisterMapType((map[string]*ExtraActionInfo_StringList)(nil), "blaze.ExtraActionInfo.AspectParametersEntry")
	proto.RegisterType((*ExtraActionInfo_StringList)(nil), "blaze.ExtraActionInfo.StringList")
	proto.RegisterType((*ExtraActionInfo_AspectDescriptor)(nil), "blaze.ExtraActionInfo.AspectDescriptor")
	proto.RegisterMapType((map[string]*ExtraActionInfo_AspectDescriptor_StringList)(nil), "blaze.ExtraActionInfo.AspectDescriptor.AspectParametersEntry")
	proto.RegisterType((*ExtraActionInfo_AspectDescriptor_StringList)(nil), "blaze.ExtraActionInfo.AspectDescriptor.StringList")
	proto.RegisterType((*EnvironmentVariable)(nil), "blaze.EnvironmentVariable")
	proto.RegisterExtension(E_SpawnInfo_SpawnInfo)
	proto.RegisterType((*SpawnInfo)(nil), "blaze.SpawnInfo")
	proto.RegisterExtension(E_CppCompileInfo_CppCompileInfo)
	proto.RegisterType((*CppCompileInfo)(nil), "blaze.CppCompileInfo")
	proto.RegisterExtension(E_CppLinkInfo_CppLinkInfo)
	proto.RegisterType((*CppLinkInfo)(nil), "blaze.CppLinkInfo")
	proto.RegisterExtension(E_JavaCompileInfo_JavaCompileInfo)
	proto.RegisterType((*JavaCompileInfo)(nil), "blaze.JavaCompileInfo")
	proto.RegisterExtension(E_PythonInfo_PythonInfo)
	proto.RegisterType((*PythonInfo)(nil), "blaze.PythonInfo")
}

func init() {
	proto.RegisterFile("third_party/bazel/src/main/protobuf/extra_actions_base.proto", fileDescriptor_95b6b55463e685df)
}

var fileDescriptor_95b6b55463e685df = []byte{
	// 1067 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x56, 0x9c, 0xa6, 0x8d, 0x4f, 0xb4, 0x69, 0x3a, 0xa5, 0x6c, 0x36, 0xc0, 0x6e, 0x09, 0x88,
	0xad, 0x00, 0x39, 0x52, 0x2e, 0x16, 0xc4, 0x8f, 0x50, 0xb7, 0x5b, 0x54, 0xa0, 0xa2, 0x95, 0xbb,
	0x42, 0x48, 0x08, 0x59, 0x13, 0x7b, 0x92, 0x4e, 0x6b, 0x7b, 0x46, 0xe3, 0x49, 0x4a, 0xf6, 0xaa,
	0x4f, 0xc0, 0x1d, 0x77, 0x3c, 0x0a, 0x2f, 0xc1, 0x5b, 0xec, 0x82, 0xb8, 0xe3, 0x01, 0x56, 0xf3,
	0x63, 0xbb, 0x4e, 0xd3, 0x6e, 0xef, 0x66, 0xbe, 0xf3, 0xf9, 0xcc, 0x39, 0xdf, 0x37, 0x73, 0x12,
	0xf8, 0x4a, 0x9e, 0x52, 0x11, 0x05, 0x1c, 0x0b, 0x39, 0x1f, 0x8c, 0xf0, 0x0b, 0x12, 0x0f, 0x32,
	0x11, 0x0e, 0x12, 0x4c, 0xd3, 0x01, 0x17, 0x4c, 0xb2, 0xd1, 0x74, 0x3c, 0x20, 0xbf, 0x49, 0x81,
	0x03, 0x1c, 0x4a, 0xca, 0xd2, 0x2c, 0x18, 0xe1, 0x8c, 0x78, 0x3a, 0x86, 0x1a, 0xa3, 0x18, 0xbf,
	0x20, 0xfd, 0x43, 0x40, 0xfb, 0x8a, 0xb2, 0xab, 0x19, 0x27, 0xd3, 0x24, 0xc1, 0x62, 0x8e, 0x9e,
	0xc0, 0xaa, 0xf9, 0xa4, 0x5b, 0xdb, 0xae, 0xef, 0xb4, 0x86, 0x0f, 0x3d, 0xcd, 0xf6, 0x9e, 0x11,
	0x89, 0x69, 0x4c, 0xa2, 0x2b, 0x9f, 0x7c, 0x97, 0x8e, 0x99, 0x6f, 0xd9, 0x7d, 0x01, 0xf7, 0x6f,
	0xa0, 0xa0, 0xc7, 0xb0, 0x2e, 0x05, 0x9d, 0x4c, 0x88, 0xa0, 0xe9, 0x24, 0x18, 0xd3, 0x98, 0x74,
	0x6b, 0xdb, 0xb5, 0x1d, 0xd7, 0x6f, 0x97, 0xf0, 0xb7, 0x34, 0x26, 0xc8, 0x2b, 0xce, 0x76, 0xb6,
	0x9d, 0x9d, 0xd6, 0xf0, 0x6d, 0x7b, 0xf6, 0x4d, 0x67, 0xfe, 0xdf, 0x80, 0xf5, 0xc5, 0xc3, 0xde,
	0x82, 0x06, 0xbb, 0x48, 0x89, 0xb0, 0x47, 0x98, 0x0d, 0xfa, 0x00, 0x5a, 0x38, 0xe3, 0x24, 0x94,
	0x41, 0x8a, 0x13, 0xd2, 0x5d, 0x55, 0xb1, 0xa7, 0x4e, 0xb7, 0xe6, 0x83, 0x81, 0x7f, 0xc4, 0x09,
	0x41, 0xbf, 0xc2, 0x86, 0x25, 0x71, 0x2c, 0x70, 0x42, 0x24, 0x11, 0x59, 0x77, 0x4d, 0xab, 0xf0,
	0xe9, 0xf2, 0x4a, 0xbc, 0x5d, 0xcd, 0x3f, 0x2e, 0xe8, 0xfb, 0xa9, 0x14, 0x73, 0x9d, 0xb8, 0x83,
	0x17, 0x42, 0x68, 0x17, 0xd6, 0x0c, 0x96, 0x75, 0x9b, 0x3a, 0xe9, 0xe3, 0x5b, 0x93, 0x3e, 0x23,
	0x59, 0x28, 0x28, 0x97, 0x4c, 0xf8, 0xf9, 0x77, 0xa8, 0x0d, 0x0e, 0x8d, 0xba, 0x8e, 0xee, 0xcc,
	0xa1, 0x11, 0xea, 0x41, 0x33, 0x49, 0x49, 0xc2, 0x52, 0x1a, 0x76, 0x1b, 0x1a, 0x2d, 0xf6, 0xbd,
	0x31, 0x6c, 0x2d, 0xad, 0x0e, 0x75, 0xa0, 0x7e, 0x4e, 0xe6, 0x56, 0x1f, 0xb5, 0x44, 0x9f, 0x41,
	0x63, 0x86, 0xe3, 0x29, 0xd1, 0x99, 0x5b, 0xc3, 0xf7, 0x6f, 0xa8, 0xeb, 0x44, 0x2a, 0xa7, 0x0e,
	0x69, 0x26, 0x7d, 0xc3, 0xff, 0xc2, 0xf9, 0xbc, 0xd6, 0xfb, 0x08, 0xa0, 0x0c, 0x28, 0xf9, 0x4d,
	0x2a, 0x75, 0x7b, 0xdc, 0x82, 0xd7, 0xad, 0xf5, 0xfe, 0x72, 0xa0, 0xb3, 0xd8, 0x19, 0x7a, 0x54,
	0xf5, 0xc5, 0xd4, 0x74, 0xd5, 0x93, 0xb3, 0x65, 0x9e, 0x38, 0x5a, 0xbe, 0xaf, 0xef, 0x28, 0xdf,
	0x72, 0x93, 0xae, 0x1b, 0xd4, 0xbb, 0xb8, 0xbb, 0x62, 0x07, 0x55, 0xc5, 0x86, 0x77, 0x2d, 0x65,
	0xb9, 0x84, 0xfd, 0x37, 0x4b, 0xf8, 0xb1, 0xdb, 0x7c, 0xb9, 0xd6, 0xb9, 0xbc, 0xbc, 0xbc, 0x74,
	0xfa, 0xdf, 0xc0, 0xe6, 0x7e, 0x3a, 0xa3, 0x82, 0xa5, 0x09, 0x49, 0xe5, 0x4f, 0x58, 0x50, 0x3c,
	0x8a, 0x09, 0x42, 0xb0, 0x62, 0x45, 0x74, 0x76, 0x5c, 0x5f, 0xaf, 0xcb, 0x5c, 0x8e, 0x06, 0xcd,
	0xa6, 0xff, 0xaa, 0x06, 0xee, 0x09, 0xc7, 0x17, 0xe6, 0xc5, 0xf4, 0xa0, 0x89, 0xc5, 0x64, 0xaa,
	0x72, 0xd9, 0x23, 0x8b, 0x3d, 0x7a, 0x02, 0xcd, 0x99, 0xcd, 0x6f, 0x55, 0xef, 0xe5, 0xad, 0x5e,
	0xaf, 0xc0, 0x2f, 0xb8, 0xe8, 0x3d, 0x00, 0x9a, 0xf2, 0xa9, 0x34, 0xaf, 0x7d, 0x45, 0x67, 0x75,
	0x35, 0xa2, 0x1f, 0xfa, 0x23, 0x68, 0xb1, 0xa9, 0x2c, 0xe2, 0x0d, 0x1d, 0x07, 0x03, 0x29, 0xc2,
	0xf0, 0x00, 0x20, 0x53, 0x05, 0x06, 0x54, 0x55, 0x78, 0xc3, 0x1c, 0xe8, 0xfe, 0xbb, 0xa6, 0xd5,
	0xef, 0xd8, 0x70, 0xd1, 0x92, 0xef, 0x66, 0xf9, 0xb2, 0xff, 0xb7, 0x03, 0xed, 0x3d, 0xce, 0xf7,
	0x58, 0xc2, 0x69, 0x4c, 0x74, 0xc3, 0x08, 0x56, 0x24, 0x63, 0xb1, 0xf5, 0x53, 0xaf, 0xd5, 0x8c,
	0x0a, 0x0d, 0x45, 0x04, 0x8c, 0xdb, 0x19, 0xa4, 0xaa, 0x6a, 0xe7, 0xf0, 0x91, 0x46, 0x55, 0xe9,
	0x19, 0x9b, 0x8a, 0x90, 0x98, 0xd2, 0xeb, 0xe6, 0xc6, 0x1a, 0x68, 0x59, 0x6f, 0x2b, 0x86, 0x50,
	0xf6, 0x86, 0x3c, 0xd8, 0x34, 0xf4, 0x2c, 0xc0, 0x69, 0x14, 0x9c, 0x12, 0x1c, 0xa9, 0x4b, 0x6d,
	0x44, 0xd8, 0xb0, 0xa1, 0xdd, 0x34, 0x3a, 0x30, 0x81, 0x8a, 0x07, 0xab, 0x77, 0xf7, 0x60, 0xf8,
	0x33, 0x74, 0x42, 0xce, 0x03, 0x5b, 0xff, 0xed, 0x4a, 0xbe, 0x32, 0x4a, 0x6e, 0xd9, 0x70, 0x55,
	0x30, 0xbf, 0x1d, 0x56, 0xf6, 0xfd, 0x3f, 0xeb, 0xd0, 0xda, 0xe3, 0xfc, 0x90, 0xa6, 0xe7, 0x5a,
	0xd0, 0xaa, 0xdb, 0xb5, 0x37, 0xb8, 0xed, 0x5c, 0x53, 0x64, 0x08, 0x5b, 0x34, 0x95, 0x44, 0x8c,
	0x71, 0x48, 0x82, 0xab, 0x54, 0xa3, 0xee, 0x66, 0x11, 0x3c, 0x2a, 0xbf, 0xd9, 0x81, 0x4e, 0x4c,
	0xd3, 0xf3, 0x40, 0x62, 0x31, 0x21, 0x32, 0x90, 0x73, 0x9e, 0x6b, 0xdd, 0x56, 0xf8, 0x73, 0x0d,
	0x3f, 0x9f, 0x73, 0xa2, 0xac, 0xd5, 0xcc, 0x4c, 0x62, 0x49, 0xc3, 0x94, 0x64, 0x99, 0x9d, 0x95,
	0x9a, 0x78, 0x52, 0xa0, 0xaa, 0x8d, 0x9c, 0x98, 0x70, 0x2d, 0xb5, 0xeb, 0xbb, 0x96, 0x93, 0x70,
	0xf4, 0x25, 0xf4, 0x46, 0x53, 0x1a, 0x47, 0x5a, 0x49, 0x6b, 0x5b, 0x80, 0x85, 0xa4, 0x63, 0x1c,
	0x4a, 0xfd, 0x3b, 0xe1, 0xfa, 0xf7, 0x35, 0x43, 0x89, 0x62, 0xdc, 0xdb, 0xb5, 0x61, 0xf4, 0x00,
	0x9a, 0x3a, 0x37, 0xe3, 0x52, 0x4f, 0x7f, 0xd7, 0x5f, 0x53, 0xfb, 0x23, 0x2e, 0x87, 0x47, 0x70,
	0x4f, 0xf9, 0xa4, 0xc3, 0xb7, 0x9a, 0xf4, 0x8f, 0x31, 0x09, 0x95, 0x26, 0xe5, 0x0e, 0xf8, 0xad,
	0xb0, 0xdc, 0xf4, 0x7f, 0xaf, 0xc3, 0xfa, 0xf7, 0x78, 0x86, 0xaf, 0xde, 0xf9, 0x77, 0xc1, 0x35,
	0xc2, 0x9e, 0xe1, 0xfc, 0xa7, 0xb1, 0x04, 0x54, 0x34, 0x8c, 0x71, 0x96, 0x71, 0x2c, 0x4f, 0xed,
	0xbd, 0x2f, 0x01, 0xf4, 0x10, 0xec, 0xfd, 0xd6, 0xe1, 0xba, 0x79, 0xac, 0x25, 0xb2, 0xf8, 0x24,
	0x56, 0xae, 0x12, 0xb4, 0x57, 0xef, 0x80, 0x7b, 0x86, 0x67, 0x38, 0xd4, 0xdd, 0x9b, 0x7b, 0xde,
	0xd4, 0xc0, 0x11, 0x97, 0xea, 0x6c, 0x2e, 0x58, 0x48, 0xb2, 0x8c, 0x89, 0x5c, 0xf4, 0x02, 0x40,
	0x1f, 0xc2, 0xbd, 0x62, 0xa3, 0x8f, 0x37, 0x3a, 0x57, 0x41, 0xc5, 0x1a, 0x31, 0x26, 0xcb, 0x1e,
	0x8c, 0xc4, 0x55, 0xb0, 0x32, 0xe8, 0xdc, 0xea, 0xa0, 0x1b, 0xfe, 0x02, 0x1b, 0xaa, 0xa2, 0xbb,
	0xbd, 0x96, 0x97, 0xc6, 0x88, 0x3c, 0xbc, 0xa0, 0xb5, 0xbf, 0x7e, 0x56, 0x05, 0xfa, 0x7f, 0xd4,
	0x00, 0x8e, 0xe7, 0xf2, 0xd4, 0xfe, 0x45, 0x59, 0xd0, 0xab, 0x76, 0x4d, 0xaf, 0x07, 0xd0, 0x8c,
	0x08, 0xcf, 0x5f, 0x8b, 0xbe, 0x2c, 0x11, 0xe1, 0x7a, 0x30, 0xfe, 0x00, 0x2d, 0xae, 0x33, 0xdd,
	0x5e, 0xe1, 0x7f, 0xa6, 0xc2, 0x0d, 0x1b, 0x2e, 0x0f, 0xf7, 0x81, 0x17, 0xeb, 0xa7, 0x03, 0xf8,
	0x24, 0x64, 0x89, 0x37, 0x61, 0x6c, 0x12, 0x13, 0x2f, 0x22, 0x33, 0x35, 0x0a, 0x33, 0x4f, 0xdf,
	0x61, 0x2f, 0xa6, 0x23, 0xcf, 0xfe, 0x79, 0xf4, 0xf4, 0x5f, 0xc9, 0xe3, 0xda, 0xeb, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0x0d, 0x53, 0x93, 0x78, 0x0a, 0x00, 0x00,
}
