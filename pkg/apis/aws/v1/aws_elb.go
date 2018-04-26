
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	CrossZoneLoadBalancing	bool	`json:"cross_zone_load_balancing"`
	AccessLogs	[]Generic	`json:"access_logs"`
	Tags	map[string]Generic	`json:"tags"`
	NamePrefix	string	`json:"name_prefix"`
	ConnectionDrainingTimeout	int	`json:"connection_draining_timeout"`
	Listener	Generic	`json:"listener"`
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

