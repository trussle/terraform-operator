
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	NamePrefix	string	`json:"name_prefix"`
	LoadBalancerType	string	`json:"load_balancer_type"`
	EnableDeletionProtection	bool	`json:"enable_deletion_protection"`
	EnableCrossZoneLoadBalancing	bool	`json:"enable_cross_zone_load_balancing"`
	Tags	map[string]string	`json:"tags"`
	IdleTimeout	int	`json:"idle_timeout"`
	EnableHttp2	bool	`json:"enable_http2"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbList is a list of AwsLb resources
type AwsLbList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLb	`json:"items"`
}


// AccessLogs is a AccessLogs
type AccessLogs struct {
	Bucket	string	`json:"bucket"`
}

// SubnetMapping is a SubnetMapping
type SubnetMapping struct {
	SubnetId	string	`json:"subnet_id"`
	AllocationId	string	`json:"allocation_id"`
}
