
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
	Aliases	string	`json:"aliases"`
	CacheBehavior	string	`json:"cache_behavior"`
	DefaultCacheBehavior	string	`json:"default_cache_behavior"`
	Tags	map[string]???	`json:"tags"`
	Comment	string	`json:"comment"`
	ViewerCertificate	string	`json:"viewer_certificate"`
	LoggingConfig	string	`json:"logging_config"`
	RetainOnDelete	bool	`json:"retain_on_delete"`
	IsIpv6Enabled	bool	`json:"is_ipv6_enabled"`
	DefaultRootObject	string	`json:"default_root_object"`
	PriceClass	string	`json:"price_class"`
	CustomErrorResponse	string	`json:"custom_error_response"`
	WebAclId	string	`json:"web_acl_id"`
	Origin	string	`json:"origin"`
	Restrictions	string	`json:"restrictions"`
	OrderedCacheBehavior	[]VwHjKVtP	`json:"ordered_cache_behavior"`
	Enabled	bool	`json:"enabled"`
	HttpVersion	string	`json:"http_version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudfrontDistributionList is a list of AwsCloudfrontDistribution resources
type AwsCloudfrontDistributionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudfrontDistribution	`json:"items"`
}


// VwHjKVtP is a VwHjKVtP
type VwHjKVtP struct {
	ForwardedValues	string	`json:"forwarded_values"`
	LambdaFunctionAssociation	string	`json:"lambda_function_association"`
	TargetOriginId	string	`json:"target_origin_id"`
	ViewerProtocolPolicy	string	`json:"viewer_protocol_policy"`
	Compress	bool	`json:"compress"`
	FieldLevelEncryptionId	string	`json:"field_level_encryption_id"`
	MaxTtl	int	`json:"max_ttl"`
	MinTtl	int	`json:"min_ttl"`
	DefaultTtl	int	`json:"default_ttl"`
	SmoothStreaming	bool	`json:"smooth_streaming"`
	CachedMethods	string	`json:"cached_methods"`
	PathPattern	string	`json:"path_pattern"`
	TrustedSigners	[]string	`json:"trusted_signers"`
	AllowedMethods	string	`json:"allowed_methods"`
}
