
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Tags	map[string]Generic	`json:"tags"`
	ViewerCertificate	Generic	`json:"viewer_certificate"`
	Enabled	bool	`json:"enabled"`
	IsIpv6Enabled	bool	`json:"is_ipv6_enabled"`
	Comment	string	`json:"comment"`
	WebAclId	string	`json:"web_acl_id"`
	Restrictions	Generic	`json:"restrictions"`
	RetainOnDelete	bool	`json:"retain_on_delete"`
	Aliases	Generic	`json:"aliases"`
	DefaultRootObject	string	`json:"default_root_object"`
	PriceClass	string	`json:"price_class"`
	OrderedCacheBehavior	[]Generic	`json:"ordered_cache_behavior"`
	DefaultCacheBehavior	Generic	`json:"default_cache_behavior"`
	HttpVersion	string	`json:"http_version"`
	Origin	Generic	`json:"origin"`
	CacheBehavior	Generic	`json:"cache_behavior"`
	CustomErrorResponse	Generic	`json:"custom_error_response"`
	LoggingConfig	Generic	`json:"logging_config"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudfrontDistributionList is a list of AwsCloudfrontDistribution resources
type AwsCloudfrontDistributionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudfrontDistribution	`json:"items"`
}

