
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudfrontDistribution describes a AwsCloudfrontDistribution resource
type AwsCloudfrontDistribution struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudfrontDistributionSpec	`json:"spec"`
}


// AwsCloudfrontDistributionSpec is the spec for a AwsCloudfrontDistribution Resource
type AwsCloudfrontDistributionSpec struct {
	PriceClass	string	`json:"price_class"`
	Enabled	bool	`json:"enabled"`
	Restrictions	Restrictions	`json:"restrictions"`
	IsIpv6Enabled	bool	`json:"is_ipv6_enabled"`
	DefaultRootObject	string	`json:"default_root_object"`
	LoggingConfig	LoggingConfig	`json:"logging_config"`
	RetainOnDelete	bool	`json:"retain_on_delete"`
	WebAclId	string	`json:"web_acl_id"`
	Tags	map[string]string	`json:"tags"`
	Aliases	string	`json:"aliases"`
	DefaultCacheBehavior	DefaultCacheBehavior	`json:"default_cache_behavior"`
	Origin	Origin	`json:"origin"`
	ViewerCertificate	ViewerCertificate	`json:"viewer_certificate"`
	OrderedCacheBehavior	[]OrderedCacheBehavior	`json:"ordered_cache_behavior"`
	Comment	string	`json:"comment"`
	CustomErrorResponse	CustomErrorResponse	`json:"custom_error_response"`
	HttpVersion	string	`json:"http_version"`
	CacheBehavior	CacheBehavior	`json:"cache_behavior"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudfrontDistributionList is a list of AwsCloudfrontDistribution resources
type AwsCloudfrontDistributionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudfrontDistribution	`json:"items"`
}


// CacheBehavior is a CacheBehavior
type CacheBehavior struct {
	Compress	bool	`json:"compress"`
	PathPattern	string	`json:"path_pattern"`
	DefaultTtl	int	`json:"default_ttl"`
	MaxTtl	int	`json:"max_ttl"`
	TargetOriginId	string	`json:"target_origin_id"`
	ViewerProtocolPolicy	string	`json:"viewer_protocol_policy"`
	CachedMethods	[]string	`json:"cached_methods"`
	ForwardedValues	ForwardedValues	`json:"forwarded_values"`
	MinTtl	int	`json:"min_ttl"`
	SmoothStreaming	bool	`json:"smooth_streaming"`
	AllowedMethods	[]string	`json:"allowed_methods"`
	FieldLevelEncryptionId	string	`json:"field_level_encryption_id"`
	LambdaFunctionAssociation	LambdaFunctionAssociation	`json:"lambda_function_association"`
	TrustedSigners	[]string	`json:"trusted_signers"`
}

// Cookies is a Cookies
type Cookies struct {
	Forward	string	`json:"forward"`
	WhitelistedNames	[]string	`json:"whitelisted_names"`
}

// ForwardedValues is a ForwardedValues
type ForwardedValues struct {
	Cookies	Cookies	`json:"cookies"`
	Headers	[]string	`json:"headers"`
	QueryString	bool	`json:"query_string"`
	QueryStringCacheKeys	[]string	`json:"query_string_cache_keys"`
}

// DefaultCacheBehavior is a DefaultCacheBehavior
type DefaultCacheBehavior struct {
	Compress	bool	`json:"compress"`
	ForwardedValues	ForwardedValues	`json:"forwarded_values"`
	SmoothStreaming	bool	`json:"smooth_streaming"`
	TrustedSigners	[]string	`json:"trusted_signers"`
	LambdaFunctionAssociation	LambdaFunctionAssociation	`json:"lambda_function_association"`
	MaxTtl	int	`json:"max_ttl"`
	MinTtl	int	`json:"min_ttl"`
	TargetOriginId	string	`json:"target_origin_id"`
	AllowedMethods	[]string	`json:"allowed_methods"`
	CachedMethods	[]string	`json:"cached_methods"`
	DefaultTtl	int	`json:"default_ttl"`
	FieldLevelEncryptionId	string	`json:"field_level_encryption_id"`
	ViewerProtocolPolicy	string	`json:"viewer_protocol_policy"`
}

// Origin is a Origin
type Origin struct {
	OriginPath	string	`json:"origin_path"`
	S3OriginConfig	S3OriginConfig	`json:"s3_origin_config"`
	CustomOriginConfig	CustomOriginConfig	`json:"custom_origin_config"`
	DomainName	string	`json:"domain_name"`
	CustomHeader	CustomHeader	`json:"custom_header"`
	OriginId	string	`json:"origin_id"`
}

// OrderedCacheBehavior is a OrderedCacheBehavior
type OrderedCacheBehavior struct {
	Compress	bool	`json:"compress"`
	FieldLevelEncryptionId	string	`json:"field_level_encryption_id"`
	LambdaFunctionAssociation	LambdaFunctionAssociation	`json:"lambda_function_association"`
	MinTtl	int	`json:"min_ttl"`
	SmoothStreaming	bool	`json:"smooth_streaming"`
	TrustedSigners	[]string	`json:"trusted_signers"`
	CachedMethods	string	`json:"cached_methods"`
	DefaultTtl	int	`json:"default_ttl"`
	MaxTtl	int	`json:"max_ttl"`
	TargetOriginId	string	`json:"target_origin_id"`
	ViewerProtocolPolicy	string	`json:"viewer_protocol_policy"`
	AllowedMethods	string	`json:"allowed_methods"`
	ForwardedValues	ForwardedValues	`json:"forwarded_values"`
	PathPattern	string	`json:"path_pattern"`
}

// LambdaFunctionAssociation is a LambdaFunctionAssociation
type LambdaFunctionAssociation struct {
	LambdaArn	string	`json:"lambda_arn"`
	EventType	string	`json:"event_type"`
}

// CustomHeader is a CustomHeader
type CustomHeader struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}

// ViewerCertificate is a ViewerCertificate
type ViewerCertificate struct {
	IamCertificateId	string	`json:"iam_certificate_id"`
	MinimumProtocolVersion	string	`json:"minimum_protocol_version"`
	SslSupportMethod	string	`json:"ssl_support_method"`
	AcmCertificateArn	string	`json:"acm_certificate_arn"`
	CloudfrontDefaultCertificate	bool	`json:"cloudfront_default_certificate"`
}

// LoggingConfig is a LoggingConfig
type LoggingConfig struct {
	Bucket	string	`json:"bucket"`
	IncludeCookies	bool	`json:"include_cookies"`
	Prefix	string	`json:"prefix"`
}

// S3OriginConfig is a S3OriginConfig
type S3OriginConfig struct {
	OriginAccessIdentity	string	`json:"origin_access_identity"`
}

// GeoRestriction is a GeoRestriction
type GeoRestriction struct {
	Locations	[]string	`json:"locations"`
	RestrictionType	string	`json:"restriction_type"`
}

// Restrictions is a Restrictions
type Restrictions struct {
	GeoRestriction	GeoRestriction	`json:"geo_restriction"`
}

// CustomOriginConfig is a CustomOriginConfig
type CustomOriginConfig struct {
	OriginKeepaliveTimeout	int	`json:"origin_keepalive_timeout"`
	OriginReadTimeout	int	`json:"origin_read_timeout"`
	OriginProtocolPolicy	string	`json:"origin_protocol_policy"`
	OriginSslProtocols	[]string	`json:"origin_ssl_protocols"`
	HttpPort	int	`json:"http_port"`
	HttpsPort	int	`json:"https_port"`
}

// CustomErrorResponse is a CustomErrorResponse
type CustomErrorResponse struct {
	ErrorCachingMinTtl	int	`json:"error_caching_min_ttl"`
	ErrorCode	int	`json:"error_code"`
	ResponseCode	int	`json:"response_code"`
	ResponsePagePath	string	`json:"response_page_path"`
}
