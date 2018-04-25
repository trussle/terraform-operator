
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLb describes a AwsLb resource
type AwsLb struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbSpec	`json:"spec"`
}


// AwsLbSpec is the spec for a AwsLb Resource
type AwsLbSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	EnableHttp2	bool	`json:"enable_http2"`
	LoadBalancerType	string	`json:"load_balancer_type"`
	EnableCrossZoneLoadBalancing	bool	`json:"enable_cross_zone_load_balancing"`
	NamePrefix	string	`json:"name_prefix"`
	EnableDeletionProtection	bool	`json:"enable_deletion_protection"`
	IdleTimeout	int	`json:"idle_timeout"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbList is a list of AwsLb resources
type AwsLbList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLb	`json:"items"`
}

