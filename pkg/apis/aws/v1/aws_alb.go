
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlb struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbSpec	`json:"spec"`
}

type AwsAlbList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlb	`json:"items"`
}

type AwsAlbSpec struct {
	EnableCrossZoneLoadBalancing	bool	`json:"enable_cross_zone_load_balancing"`
	EnableDeletionProtection	bool	`json:"enable_deletion_protection"`
	IdleTimeout	int	`json:"idle_timeout"`
	LoadBalancerType	string	`json:"load_balancer_type"`
	Tags	map[string]interface{}	`json:"tags"`
	EnableHttp2	bool	`json:"enable_http2"`
	NamePrefix	string	`json:"name_prefix"`
}
