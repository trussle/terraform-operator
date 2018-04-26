
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	EnableCrossZoneLoadBalancing	bool	`json:"enable_cross_zone_load_balancing"`
	EnableHttp2	bool	`json:"enable_http2"`
	EnableDeletionProtection	bool	`json:"enable_deletion_protection"`
	IdleTimeout	int	`json:"idle_timeout"`
	Tags	map[string]string	`json:"tags"`
	LoadBalancerType	string	`json:"load_balancer_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbList is a list of AwsAlb resources
type AwsAlbList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlb	`json:"items"`
}


// SubnetMapping is a SubnetMapping
type SubnetMapping struct {
	SubnetId	string	`json:"subnet_id"`
	AllocationId	string	`json:"allocation_id"`
}

// AccessLogs is a AccessLogs
type AccessLogs struct {
	Bucket	string	`json:"bucket"`
}
