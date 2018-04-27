
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
	Tags	map[string]string	`json:"tags"`
	DefaultRootObject	string	`json:"default_root_object"`
	LoggingConfig	AwsCloudfrontDistributionLoggingConfig	`json:"logging_config"`
	Origin	AwsCloudfrontDistributionOrigin	`json:"origin"`
	ViewerCertificate	AwsCloudfrontDistributionViewerCertificate	`json:"viewer_certificate"`
	Aliases	string	`json:"aliases"`
	HttpVersion	string	`json:"http_version"`
	IsIpv6Enabled	bool	`json:"is_ipv6_enabled"`
	PriceClass	string	`json:"price_class"`
	WebAclId	string	`json:"web_acl_id"`
	Restrictions	AwsCloudfrontDistributionRestrictions	`json:"restrictions"`
	RetainOnDelete	bool	`json:"retain_on_delete"`
	CacheBehavior	AwsCloudfrontDistributionCacheBehavior	`json:"cache_behavior"`
	CustomErrorResponse	AwsCloudfrontDistributionCustomErrorResponse	`json:"custom_error_response"`
	Enabled	bool	`json:"enabled"`
	OrderedCacheBehavior	[]AwsCloudfrontDistributionOrderedCacheBehavior	`json:"ordered_cache_behavior"`
	DefaultCacheBehavior	AwsCloudfrontDistributionDefaultCacheBehavior	`json:"default_cache_behavior"`
	Comment	string	`json:"comment"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudfrontDistributionList is a list of AwsCloudfrontDistribution resources
type AwsCloudfrontDistributionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudfrontDistribution	`json:"items"`
}


// AwsCloudfrontDistributionGeoRestriction is a AwsCloudfrontDistributionGeoRestriction
type AwsCloudfrontDistributionGeoRestriction struct {
	Locations	[]string	`json:"locations"`
	RestrictionType	string	`json:"restriction_type"`
}

// AwsCloudfrontDistributionLoggingConfig is a AwsCloudfrontDistributionLoggingConfig
type AwsCloudfrontDistributionLoggingConfig struct {
	Bucket	string	`json:"bucket"`
	IncludeCookies	bool	`json:"include_cookies"`
	Prefix	string	`json:"prefix"`
}

// AwsCloudfrontDistributionCacheBehavior is a AwsCloudfrontDistributionCacheBehavior
type AwsCloudfrontDistributionCacheBehavior struct {
	AllowedMethods	[]string	`json:"allowed_methods"`
	ForwardedValues	AwsCloudfrontDistributionForwardedValues	`json:"forwarded_values"`
	MinTtl	int	`json:"min_ttl"`
	Compress	bool	`json:"compress"`
	LambdaFunctionAssociation	AwsCloudfrontDistributionLambdaFunctionAssociation	`json:"lambda_function_association"`
	CachedMethods	[]string	`json:"cached_methods"`
	DefaultTtl	int	`json:"default_ttl"`
	FieldLevelEncryptionId	string	`json:"field_level_encryption_id"`
	MaxTtl	int	`json:"max_ttl"`
	PathPattern	string	`json:"path_pattern"`
	SmoothStreaming	bool	`json:"smooth_streaming"`
	TrustedSigners	[]string	`json:"trusted_signers"`
	ViewerProtocolPolicy	string	`json:"viewer_protocol_policy"`
	TargetOriginId	string	`json:"target_origin_id"`
}

// AwsCloudfrontDistributionOrderedCacheBehavior is a AwsCloudfrontDistributionOrderedCacheBehavior
type AwsCloudfrontDistributionOrderedCacheBehavior struct {
	SmoothStreaming	bool	`json:"smooth_streaming"`
	ViewerProtocolPolicy	string	`json:"viewer_protocol_policy"`
	CachedMethods	string	`json:"cached_methods"`
	FieldLevelEncryptionId	string	`json:"field_level_encryption_id"`
	ForwardedValues	AwsCloudfrontDistributionForwardedValues	`json:"forwarded_values"`
	MaxTtl	int	`json:"max_ttl"`
	MinTtl	int	`json:"min_ttl"`
	PathPattern	string	`json:"path_pattern"`
	AllowedMethods	string	`json:"allowed_methods"`
	DefaultTtl	int	`json:"default_ttl"`
	TrustedSigners	[]string	`json:"trusted_signers"`
	Compress	bool	`json:"compress"`
	LambdaFunctionAssociation	AwsCloudfrontDistributionLambdaFunctionAssociation	`json:"lambda_function_association"`
	TargetOriginId	string	`json:"target_origin_id"`
}

// AwsCloudfrontDistributionDefaultCacheBehavior is a AwsCloudfrontDistributionDefaultCacheBehavior
type AwsCloudfrontDistributionDefaultCacheBehavior struct {
	MinTtl	int	`json:"min_ttl"`
	SmoothStreaming	bool	`json:"smooth_streaming"`
	TargetOriginId	string	`json:"target_origin_id"`
	TrustedSigners	[]string	`json:"trusted_signers"`
	CachedMethods	[]string	`json:"cached_methods"`
	FieldLevelEncryptionId	string	`json:"field_level_encryption_id"`
	LambdaFunctionAssociation	AwsCloudfrontDistributionLambdaFunctionAssociation	`json:"lambda_function_association"`
	MaxTtl	int	`json:"max_ttl"`
	ViewerProtocolPolicy	string	`json:"viewer_protocol_policy"`
	AllowedMethods	[]string	`json:"allowed_methods"`
	Compress	bool	`json:"compress"`
	DefaultTtl	int	`json:"default_ttl"`
	ForwardedValues	AwsCloudfrontDistributionForwardedValues	`json:"forwarded_values"`
}

// AwsCloudfrontDistributionRestrictions is a AwsCloudfrontDistributionRestrictions
type AwsCloudfrontDistributionRestrictions struct {
	GeoRestriction	AwsCloudfrontDistributionGeoRestriction	`json:"geo_restriction"`
}

// AwsCloudfrontDistributionCookies is a AwsCloudfrontDistributionCookies
type AwsCloudfrontDistributionCookies struct {
	Forward	string	`json:"forward"`
	WhitelistedNames	[]string	`json:"whitelisted_names"`
}

// AwsCloudfrontDistributionLambdaFunctionAssociation is a AwsCloudfrontDistributionLambdaFunctionAssociation
type AwsCloudfrontDistributionLambdaFunctionAssociation struct {
	EventType	string	`json:"event_type"`
	LambdaArn	string	`json:"lambda_arn"`
}

// AwsCloudfrontDistributionCustomErrorResponse is a AwsCloudfrontDistributionCustomErrorResponse
type AwsCloudfrontDistributionCustomErrorResponse struct {
	ErrorCachingMinTtl	int	`json:"error_caching_min_ttl"`
	ErrorCode	int	`json:"error_code"`
	ResponseCode	int	`json:"response_code"`
	ResponsePagePath	string	`json:"response_page_path"`
}

// AwsCloudfrontDistributionCustomHeader is a AwsCloudfrontDistributionCustomHeader
type AwsCloudfrontDistributionCustomHeader struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}

// AwsCloudfrontDistributionS3OriginConfig is a AwsCloudfrontDistributionS3OriginConfig
type AwsCloudfrontDistributionS3OriginConfig struct {
	OriginAccessIdentity	string	`json:"origin_access_identity"`
}

// AwsCloudfrontDistributionOrigin is a AwsCloudfrontDistributionOrigin
type AwsCloudfrontDistributionOrigin struct {
	CustomOriginConfig	AwsCloudfrontDistributionCustomOriginConfig	`json:"custom_origin_config"`
	DomainName	string	`json:"domain_name"`
	CustomHeader	AwsCloudfrontDistributionCustomHeader	`json:"custom_header"`
	OriginId	string	`json:"origin_id"`
	OriginPath	string	`json:"origin_path"`
	S3OriginConfig	AwsCloudfrontDistributionS3OriginConfig	`json:"s3_origin_config"`
}

// AwsCloudfrontDistributionViewerCertificate is a AwsCloudfrontDistributionViewerCertificate
type AwsCloudfrontDistributionViewerCertificate struct {
	SslSupportMethod	string	`json:"ssl_support_method"`
	AcmCertificateArn	string	`json:"acm_certificate_arn"`
	CloudfrontDefaultCertificate	bool	`json:"cloudfront_default_certificate"`
	IamCertificateId	string	`json:"iam_certificate_id"`
	MinimumProtocolVersion	string	`json:"minimum_protocol_version"`
}

// AwsCloudfrontDistributionForwardedValues is a AwsCloudfrontDistributionForwardedValues
type AwsCloudfrontDistributionForwardedValues struct {
	Headers	[]string	`json:"headers"`
	QueryString	bool	`json:"query_string"`
	QueryStringCacheKeys	[]string	`json:"query_string_cache_keys"`
	Cookies	AwsCloudfrontDistributionCookies	`json:"cookies"`
}

// AwsCloudfrontDistributionCustomOriginConfig is a AwsCloudfrontDistributionCustomOriginConfig
type AwsCloudfrontDistributionCustomOriginConfig struct {
	OriginProtocolPolicy	string	`json:"origin_protocol_policy"`
	OriginSslProtocols	[]string	`json:"origin_ssl_protocols"`
	HttpPort	int	`json:"http_port"`
	HttpsPort	int	`json:"https_port"`
	OriginKeepaliveTimeout	int	`json:"origin_keepalive_timeout"`
	OriginReadTimeout	int	`json:"origin_read_timeout"`
}
