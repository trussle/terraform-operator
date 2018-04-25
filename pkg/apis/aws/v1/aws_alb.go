
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlb describes a AwsAlb resource
type AwsAlb struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbSpec	`json:"spec"`
}


// AwsAlbSpec is the spec for a AwsAlb Resource
type AwsAlbSpec struct {
	NamePrefix	string	`json:"name_prefix"`
	EnableDeletionProtection	bool	`json:"enable_deletion_protection"`
	EnableHttp2	bool	`json:"enable_http2"`
	EnableCrossZoneLoadBalancing	bool	`json:"enable_cross_zone_load_balancing"`
	IdleTimeout	int	`json:"idle_timeout"`
	Tags	map[string]interface{}	`json:"tags"`
	LoadBalancerType	string	`json:"load_balancer_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbList is a list of AwsAlb resources
type AwsAlbList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlb	`json:"items"`
}

