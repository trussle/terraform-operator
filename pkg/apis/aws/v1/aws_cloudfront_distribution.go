
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
	IsIpv6Enabled	bool	`json:"is_ipv6_enabled"`
	Tags	map[string]interface{}	`json:"tags"`
	DefaultCacheBehavior	string	`json:"default_cache_behavior"`
	LoggingConfig	string	`json:"logging_config"`
	DefaultRootObject	string	`json:"default_root_object"`
	HttpVersion	string	`json:"http_version"`
	PriceClass	string	`json:"price_class"`
	ViewerCertificate	string	`json:"viewer_certificate"`
	RetainOnDelete	bool	`json:"retain_on_delete"`
	Aliases	string	`json:"aliases"`
	CacheBehavior	string	`json:"cache_behavior"`
	Comment	string	`json:"comment"`
	CustomErrorResponse	string	`json:"custom_error_response"`
	Enabled	bool	`json:"enabled"`
	Restrictions	string	`json:"restrictions"`
	Origin	string	`json:"origin"`
	OrderedCacheBehavior	[]interface{}	`json:"ordered_cache_behavior"`
	WebAclId	string	`json:"web_acl_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudfrontDistributionList is a list of AwsCloudfrontDistribution resources
type AwsCloudfrontDistributionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudfrontDistribution	`json:"items"`
}

