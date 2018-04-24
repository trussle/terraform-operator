
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudfrontDistribution struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudfrontDistributionSpec	`json:"spec"`
}

type AwsCloudfrontDistributionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudfrontDistribution	`json:"items"`
}

type AwsCloudfrontDistributionSpec struct {
	DefaultCacheBehavior	interface{}	`json:"default_cache_behavior"`
	LoggingConfig	interface{}	`json:"logging_config"`
	WebAclId	string	`json:"web_acl_id"`
	CustomErrorResponse	interface{}	`json:"custom_error_response"`
	IsIpv6Enabled	bool	`json:"is_ipv6_enabled"`
	CacheBehavior	interface{}	`json:"cache_behavior"`
	Enabled	bool	`json:"enabled"`
	Origin	interface{}	`json:"origin"`
	PriceClass	string	`json:"price_class"`
	Restrictions	interface{}	`json:"restrictions"`
	Aliases	interface{}	`json:"aliases"`
	RetainOnDelete	bool	`json:"retain_on_delete"`
	DefaultRootObject	string	`json:"default_root_object"`
	Tags	map[string]interface{}	`json:"tags"`
	OrderedCacheBehavior	[]interface{}	`json:"ordered_cache_behavior"`
	Comment	string	`json:"comment"`
	HttpVersion	string	`json:"http_version"`
	ViewerCertificate	interface{}	`json:"viewer_certificate"`
}
