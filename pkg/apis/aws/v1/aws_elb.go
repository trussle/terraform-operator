
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
	Listener	string	`json:"listener"`
	CrossZoneLoadBalancing	bool	`json:"cross_zone_load_balancing"`
	IdleTimeout	int	`json:"idle_timeout"`
	AccessLogs	[]interface{}	`json:"access_logs"`
	ConnectionDrainingTimeout	int	`json:"connection_draining_timeout"`
	Tags	map[string]interface{}	`json:"tags"`
	NamePrefix	string	`json:"name_prefix"`
	ConnectionDraining	bool	`json:"connection_draining"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElbList is a list of AwsElb resources
type AwsElbList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElb	`json:"items"`
}

