// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: codyaccess.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CodyGatewayRateLimitSource int32

const (
	CodyGatewayRateLimitSource_CODY_GATEWAY_RATE_LIMIT_SOURCE_UNSPECIFIED CodyGatewayRateLimitSource = 0
	// Indicates that a custom override for the rate limit has been configured
	// and applied.
	CodyGatewayRateLimitSource_CODY_GATEWAY_RATE_LIMIT_SOURCE_OVERRIDE CodyGatewayRateLimitSource = 1
	// Indicates that the rate limit is inferred by the subscription's active plan.
	CodyGatewayRateLimitSource_CODY_GATEWAY_RATE_LIMIT_SOURCE_PLAN CodyGatewayRateLimitSource = 2
)

// Enum value maps for CodyGatewayRateLimitSource.
var (
	CodyGatewayRateLimitSource_name = map[int32]string{
		0: "CODY_GATEWAY_RATE_LIMIT_SOURCE_UNSPECIFIED",
		1: "CODY_GATEWAY_RATE_LIMIT_SOURCE_OVERRIDE",
		2: "CODY_GATEWAY_RATE_LIMIT_SOURCE_PLAN",
	}
	CodyGatewayRateLimitSource_value = map[string]int32{
		"CODY_GATEWAY_RATE_LIMIT_SOURCE_UNSPECIFIED": 0,
		"CODY_GATEWAY_RATE_LIMIT_SOURCE_OVERRIDE":    1,
		"CODY_GATEWAY_RATE_LIMIT_SOURCE_PLAN":        2,
	}
)

func (x CodyGatewayRateLimitSource) Enum() *CodyGatewayRateLimitSource {
	p := new(CodyGatewayRateLimitSource)
	*p = x
	return p
}

func (x CodyGatewayRateLimitSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodyGatewayRateLimitSource) Descriptor() protoreflect.EnumDescriptor {
	return file_codyaccess_proto_enumTypes[0].Descriptor()
}

func (CodyGatewayRateLimitSource) Type() protoreflect.EnumType {
	return &file_codyaccess_proto_enumTypes[0]
}

func (x CodyGatewayRateLimitSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodyGatewayRateLimitSource.Descriptor instead.
func (CodyGatewayRateLimitSource) EnumDescriptor() ([]byte, []int) {
	return file_codyaccess_proto_rawDescGZIP(), []int{0}
}

type GetCodyGatewayAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Query:
	//
	//	*GetCodyGatewayAccessRequest_SubscriptionId
	//	*GetCodyGatewayAccessRequest_AccessToken
	Query isGetCodyGatewayAccessRequest_Query `protobuf_oneof:"query"`
}

func (x *GetCodyGatewayAccessRequest) Reset() {
	*x = GetCodyGatewayAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codyaccess_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodyGatewayAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodyGatewayAccessRequest) ProtoMessage() {}

func (x *GetCodyGatewayAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_codyaccess_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodyGatewayAccessRequest.ProtoReflect.Descriptor instead.
func (*GetCodyGatewayAccessRequest) Descriptor() ([]byte, []int) {
	return file_codyaccess_proto_rawDescGZIP(), []int{0}
}

func (m *GetCodyGatewayAccessRequest) GetQuery() isGetCodyGatewayAccessRequest_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *GetCodyGatewayAccessRequest) GetSubscriptionId() string {
	if x, ok := x.GetQuery().(*GetCodyGatewayAccessRequest_SubscriptionId); ok {
		return x.SubscriptionId
	}
	return ""
}

func (x *GetCodyGatewayAccessRequest) GetAccessToken() string {
	if x, ok := x.GetQuery().(*GetCodyGatewayAccessRequest_AccessToken); ok {
		return x.AccessToken
	}
	return ""
}

type isGetCodyGatewayAccessRequest_Query interface {
	isGetCodyGatewayAccessRequest_Query()
}

type GetCodyGatewayAccessRequest_SubscriptionId struct {
	// The external, prefixed UUID-format identifier of an Enterprise subscription.
	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3,oneof"`
}

type GetCodyGatewayAccessRequest_AccessToken struct {
	// An license-based access token that is valid for an Enterprise subscription's
	// Cody Gateway access, e.g. 'slk_...'
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3,oneof"`
}

func (*GetCodyGatewayAccessRequest_SubscriptionId) isGetCodyGatewayAccessRequest_Query() {}

func (*GetCodyGatewayAccessRequest_AccessToken) isGetCodyGatewayAccessRequest_Query() {}

type GetCodyGatewayAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access *CodyGatewayAccess `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *GetCodyGatewayAccessResponse) Reset() {
	*x = GetCodyGatewayAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codyaccess_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodyGatewayAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodyGatewayAccessResponse) ProtoMessage() {}

func (x *GetCodyGatewayAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_codyaccess_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodyGatewayAccessResponse.ProtoReflect.Descriptor instead.
func (*GetCodyGatewayAccessResponse) Descriptor() ([]byte, []int) {
	return file_codyaccess_proto_rawDescGZIP(), []int{1}
}

func (x *GetCodyGatewayAccessResponse) GetAccess() *CodyGatewayAccess {
	if x != nil {
		return x.Access
	}
	return nil
}

type CodyGatewayRateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source of the rate limit configuration.
	Source CodyGatewayRateLimitSource `protobuf:"varint,1,opt,name=source,proto3,enum=enterpriseportal.codyaccess.v1.CodyGatewayRateLimitSource" json:"source,omitempty"`
	// Requests per time interval.
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// Interval for rate limiting.
	IntervalDuration *durationpb.Duration `protobuf:"bytes,3,opt,name=interval_duration,json=intervalDuration,proto3" json:"interval_duration,omitempty"`
}

func (x *CodyGatewayRateLimit) Reset() {
	*x = CodyGatewayRateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codyaccess_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodyGatewayRateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodyGatewayRateLimit) ProtoMessage() {}

func (x *CodyGatewayRateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_codyaccess_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodyGatewayRateLimit.ProtoReflect.Descriptor instead.
func (*CodyGatewayRateLimit) Descriptor() ([]byte, []int) {
	return file_codyaccess_proto_rawDescGZIP(), []int{2}
}

func (x *CodyGatewayRateLimit) GetSource() CodyGatewayRateLimitSource {
	if x != nil {
		return x.Source
	}
	return CodyGatewayRateLimitSource_CODY_GATEWAY_RATE_LIMIT_SOURCE_UNSPECIFIED
}

func (x *CodyGatewayRateLimit) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *CodyGatewayRateLimit) GetIntervalDuration() *durationpb.Duration {
	if x != nil {
		return x.IntervalDuration
	}
	return nil
}

type CodyGatewayAccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Access token for authenticating as the subscription holder with managed
	// Sourcegraph services.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CodyGatewayAccessToken) Reset() {
	*x = CodyGatewayAccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codyaccess_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodyGatewayAccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodyGatewayAccessToken) ProtoMessage() {}

func (x *CodyGatewayAccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_codyaccess_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodyGatewayAccessToken.ProtoReflect.Descriptor instead.
func (*CodyGatewayAccessToken) Descriptor() ([]byte, []int) {
	return file_codyaccess_proto_rawDescGZIP(), []int{3}
}

func (x *CodyGatewayAccessToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CodyGatewayAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The external, prefixed UUID-format identifier for the Enterprise
	// subscription corresponding to this Cody Gateway access description.
	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	// Whether or not a subscription has Cody Gateway access enabled.
	Enabled bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Rate limit for chat completions access, or null if not enabled.
	ChatCompletionsRateLimit *CodyGatewayRateLimit `protobuf:"bytes,3,opt,name=chat_completions_rate_limit,json=chatCompletionsRateLimit,proto3,oneof" json:"chat_completions_rate_limit,omitempty"`
	// Rate limit for code completions access, or null if not enabled.
	CodeCompletionsRateLimit *CodyGatewayRateLimit `protobuf:"bytes,4,opt,name=code_completions_rate_limit,json=codeCompletionsRateLimit,proto3,oneof" json:"code_completions_rate_limit,omitempty"`
	// Rate limit for embedding text chunks, or null if not enabled.
	EmbeddingsRateLimit *CodyGatewayRateLimit `protobuf:"bytes,5,opt,name=embeddings_rate_limit,json=embeddingsRateLimit,proto3,oneof" json:"embeddings_rate_limit,omitempty"`
	// Available access tokens for authenticating as the subscription holder with
	// Cody Gateway.
	AccessTokens []*CodyGatewayAccessToken `protobuf:"bytes,6,rep,name=access_tokens,json=accessTokens,proto3" json:"access_tokens,omitempty"`
}

func (x *CodyGatewayAccess) Reset() {
	*x = CodyGatewayAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codyaccess_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodyGatewayAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodyGatewayAccess) ProtoMessage() {}

func (x *CodyGatewayAccess) ProtoReflect() protoreflect.Message {
	mi := &file_codyaccess_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodyGatewayAccess.ProtoReflect.Descriptor instead.
func (*CodyGatewayAccess) Descriptor() ([]byte, []int) {
	return file_codyaccess_proto_rawDescGZIP(), []int{4}
}

func (x *CodyGatewayAccess) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *CodyGatewayAccess) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *CodyGatewayAccess) GetChatCompletionsRateLimit() *CodyGatewayRateLimit {
	if x != nil {
		return x.ChatCompletionsRateLimit
	}
	return nil
}

func (x *CodyGatewayAccess) GetCodeCompletionsRateLimit() *CodyGatewayRateLimit {
	if x != nil {
		return x.CodeCompletionsRateLimit
	}
	return nil
}

func (x *CodyGatewayAccess) GetEmbeddingsRateLimit() *CodyGatewayRateLimit {
	if x != nil {
		return x.EmbeddingsRateLimit
	}
	return nil
}

func (x *CodyGatewayAccess) GetAccessTokens() []*CodyGatewayAccessToken {
	if x != nil {
		return x.AccessTokens
	}
	return nil
}

type ListCodyGatewayAccessesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Clients use this field to specify the maximum number of results to be
	// returned by the server. The server may further constrain the maximum number
	// of results returned in a single page. If the page_size is 0, the server
	// will decide the number of results to be returned.
	//
	// See pagination concepts from https://cloud.google.com/apis/design/design_patterns#list_pagination
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The client uses this field to request a specific page of the list results.
	//
	// See pagination concepts from https://cloud.google.com/apis/design/design_patterns#list_pagination
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListCodyGatewayAccessesRequest) Reset() {
	*x = ListCodyGatewayAccessesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codyaccess_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCodyGatewayAccessesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCodyGatewayAccessesRequest) ProtoMessage() {}

func (x *ListCodyGatewayAccessesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_codyaccess_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCodyGatewayAccessesRequest.ProtoReflect.Descriptor instead.
func (*ListCodyGatewayAccessesRequest) Descriptor() ([]byte, []int) {
	return file_codyaccess_proto_rawDescGZIP(), []int{5}
}

func (x *ListCodyGatewayAccessesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCodyGatewayAccessesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListCodyGatewayAccessesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The list of Cody Gateway access that matched the given query.
	Accesses []*CodyGatewayAccess `protobuf:"bytes,2,rep,name=accesses,proto3" json:"accesses,omitempty"`
}

func (x *ListCodyGatewayAccessesResponse) Reset() {
	*x = ListCodyGatewayAccessesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codyaccess_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCodyGatewayAccessesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCodyGatewayAccessesResponse) ProtoMessage() {}

func (x *ListCodyGatewayAccessesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_codyaccess_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCodyGatewayAccessesResponse.ProtoReflect.Descriptor instead.
func (*ListCodyGatewayAccessesResponse) Descriptor() ([]byte, []int) {
	return file_codyaccess_proto_rawDescGZIP(), []int{6}
}

func (x *ListCodyGatewayAccessesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListCodyGatewayAccessesResponse) GetAccesses() []*CodyGatewayAccess {
	if x != nil {
		return x.Accesses
	}
	return nil
}

var File_codyaccess_proto protoreflect.FileDescriptor

var file_codyaccess_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x76, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x07, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x69, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f,
	0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x79,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x06, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x52,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a,
	0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x63, 0x6f, 0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x46, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x2e, 0x0a, 0x16, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0xf0, 0x04, 0x0a, 0x11, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x78, 0x0a, 0x1b, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x63, 0x6f, 0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x48, 0x00, 0x52, 0x18, 0x63, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x78, 0x0a, 0x1b, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x64, 0x79,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x48, 0x01,
	0x52, 0x18, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x6d, 0x0a,
	0x15, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x63, 0x6f, 0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x48, 0x02, 0x52, 0x13, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x5b, 0x0a, 0x0d,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x1e, 0x0a, 0x1c, 0x5f, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x1e, 0x0a, 0x1c, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x65, 0x6d,
	0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x5c, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x79, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x98, 0x01, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4d, 0x0a,
	0x08, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2a, 0xa2, 0x01, 0x0a,
	0x1a, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x2a, 0x43,
	0x4f, 0x44, 0x59, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x5f, 0x52, 0x41, 0x54, 0x45,
	0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2b, 0x0a, 0x27, 0x43,
	0x4f, 0x44, 0x59, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x5f, 0x52, 0x41, 0x54, 0x45,
	0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4f, 0x56,
	0x45, 0x52, 0x52, 0x49, 0x44, 0x45, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x4f, 0x44, 0x59,
	0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x49,
	0x4d, 0x49, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x10,
	0x02, 0x32, 0xce, 0x02, 0x0a, 0x11, 0x43, 0x6f, 0x64, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x3b, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x63, 0x6f, 0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01,
	0x12, 0x9f, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x3e, 0x2e, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x63, 0x6f, 0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x63, 0x6f, 0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x64, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90,
	0x02, 0x01, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x63, 0x6f,
	0x64, 0x79, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_codyaccess_proto_rawDescOnce sync.Once
	file_codyaccess_proto_rawDescData = file_codyaccess_proto_rawDesc
)

func file_codyaccess_proto_rawDescGZIP() []byte {
	file_codyaccess_proto_rawDescOnce.Do(func() {
		file_codyaccess_proto_rawDescData = protoimpl.X.CompressGZIP(file_codyaccess_proto_rawDescData)
	})
	return file_codyaccess_proto_rawDescData
}

var file_codyaccess_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_codyaccess_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_codyaccess_proto_goTypes = []interface{}{
	(CodyGatewayRateLimitSource)(0),         // 0: enterpriseportal.codyaccess.v1.CodyGatewayRateLimitSource
	(*GetCodyGatewayAccessRequest)(nil),     // 1: enterpriseportal.codyaccess.v1.GetCodyGatewayAccessRequest
	(*GetCodyGatewayAccessResponse)(nil),    // 2: enterpriseportal.codyaccess.v1.GetCodyGatewayAccessResponse
	(*CodyGatewayRateLimit)(nil),            // 3: enterpriseportal.codyaccess.v1.CodyGatewayRateLimit
	(*CodyGatewayAccessToken)(nil),          // 4: enterpriseportal.codyaccess.v1.CodyGatewayAccessToken
	(*CodyGatewayAccess)(nil),               // 5: enterpriseportal.codyaccess.v1.CodyGatewayAccess
	(*ListCodyGatewayAccessesRequest)(nil),  // 6: enterpriseportal.codyaccess.v1.ListCodyGatewayAccessesRequest
	(*ListCodyGatewayAccessesResponse)(nil), // 7: enterpriseportal.codyaccess.v1.ListCodyGatewayAccessesResponse
	(*durationpb.Duration)(nil),             // 8: google.protobuf.Duration
}
var file_codyaccess_proto_depIdxs = []int32{
	5,  // 0: enterpriseportal.codyaccess.v1.GetCodyGatewayAccessResponse.access:type_name -> enterpriseportal.codyaccess.v1.CodyGatewayAccess
	0,  // 1: enterpriseportal.codyaccess.v1.CodyGatewayRateLimit.source:type_name -> enterpriseportal.codyaccess.v1.CodyGatewayRateLimitSource
	8,  // 2: enterpriseportal.codyaccess.v1.CodyGatewayRateLimit.interval_duration:type_name -> google.protobuf.Duration
	3,  // 3: enterpriseportal.codyaccess.v1.CodyGatewayAccess.chat_completions_rate_limit:type_name -> enterpriseportal.codyaccess.v1.CodyGatewayRateLimit
	3,  // 4: enterpriseportal.codyaccess.v1.CodyGatewayAccess.code_completions_rate_limit:type_name -> enterpriseportal.codyaccess.v1.CodyGatewayRateLimit
	3,  // 5: enterpriseportal.codyaccess.v1.CodyGatewayAccess.embeddings_rate_limit:type_name -> enterpriseportal.codyaccess.v1.CodyGatewayRateLimit
	4,  // 6: enterpriseportal.codyaccess.v1.CodyGatewayAccess.access_tokens:type_name -> enterpriseportal.codyaccess.v1.CodyGatewayAccessToken
	5,  // 7: enterpriseportal.codyaccess.v1.ListCodyGatewayAccessesResponse.accesses:type_name -> enterpriseportal.codyaccess.v1.CodyGatewayAccess
	1,  // 8: enterpriseportal.codyaccess.v1.CodyAccessService.GetCodyGatewayAccess:input_type -> enterpriseportal.codyaccess.v1.GetCodyGatewayAccessRequest
	6,  // 9: enterpriseportal.codyaccess.v1.CodyAccessService.ListCodyGatewayAccesses:input_type -> enterpriseportal.codyaccess.v1.ListCodyGatewayAccessesRequest
	2,  // 10: enterpriseportal.codyaccess.v1.CodyAccessService.GetCodyGatewayAccess:output_type -> enterpriseportal.codyaccess.v1.GetCodyGatewayAccessResponse
	7,  // 11: enterpriseportal.codyaccess.v1.CodyAccessService.ListCodyGatewayAccesses:output_type -> enterpriseportal.codyaccess.v1.ListCodyGatewayAccessesResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_codyaccess_proto_init() }
func file_codyaccess_proto_init() {
	if File_codyaccess_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_codyaccess_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodyGatewayAccessRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_codyaccess_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodyGatewayAccessResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_codyaccess_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodyGatewayRateLimit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_codyaccess_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodyGatewayAccessToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_codyaccess_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodyGatewayAccess); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_codyaccess_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCodyGatewayAccessesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_codyaccess_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCodyGatewayAccessesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_codyaccess_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GetCodyGatewayAccessRequest_SubscriptionId)(nil),
		(*GetCodyGatewayAccessRequest_AccessToken)(nil),
	}
	file_codyaccess_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_codyaccess_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_codyaccess_proto_goTypes,
		DependencyIndexes: file_codyaccess_proto_depIdxs,
		EnumInfos:         file_codyaccess_proto_enumTypes,
		MessageInfos:      file_codyaccess_proto_msgTypes,
	}.Build()
	File_codyaccess_proto = out.File
	file_codyaccess_proto_rawDesc = nil
	file_codyaccess_proto_goTypes = nil
	file_codyaccess_proto_depIdxs = nil
}
