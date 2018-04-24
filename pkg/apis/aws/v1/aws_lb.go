
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLb struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbSpec	`json:"spec"`
}

type AwsLbList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLb	`json:"items"`
}

type AwsLbSpec struct {
	LoadBalancerType	string	`json:"load_balancer_type"`
	EnableHttp2	bool	`json:"enable_http2"`
	IdleTimeout	int	`json:"idle_timeout"`
	EnableCrossZoneLoadBalancing	bool	`json:"enable_cross_zone_load_balancing"`
	EnableDeletionProtection	bool	`json:"enable_deletion_protection"`
	NamePrefix	string	`json:"name_prefix"`
	Tags	map[string]interface{}	`json:"tags"`
}
