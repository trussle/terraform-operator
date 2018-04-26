
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElb describes a AwsElb resource
type AwsElb struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElbSpec	`json:"spec"`
}


// AwsElbSpec is the spec for a AwsElb Resource
type AwsElbSpec struct {
	NamePrefix	string	`json:"name_prefix"`
	Listener	string	`json:"listener"`
	ConnectionDrainingTimeout	int	`json:"connection_draining_timeout"`
	Tags	map[string]???	`json:"tags"`
	CrossZoneLoadBalancing	bool	`json:"cross_zone_load_balancing"`
	AccessLogs	[]HMLUIDjU	`json:"access_logs"`
	IdleTimeout	int	`json:"idle_timeout"`
	ConnectionDraining	bool	`json:"connection_draining"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElbList is a list of AwsElb resources
type AwsElbList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElb	`json:"items"`
}


// gBpnbPbg is a gBpnbPbg
type gBpnbPbg struct {
	HealthyThreshold	int	`json:"healthy_threshold"`
	UnhealthyThreshold	int	`json:"unhealthy_threshold"`
	Target	string	`json:"target"`
	Interval	int	`json:"interval"`
	Timeout	int	`json:"timeout"`
}

// HMLUIDjU is a HMLUIDjU
type HMLUIDjU struct {
	Interval	int	`json:"interval"`
	Bucket	string	`json:"bucket"`
	BucketPrefix	string	`json:"bucket_prefix"`
	Enabled	bool	`json:"enabled"`
}
