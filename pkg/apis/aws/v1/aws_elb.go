
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElb struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElbSpec	`json:"spec"`
}

type AwsElbList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElb	`json:"items"`
}

type AwsElbSpec struct {
	ConnectionDraining	bool	`json:"connection_draining"`
	AccessLogs	[]interface{}	`json:"access_logs"`
	Tags	map[string]interface{}	`json:"tags"`
	IdleTimeout	int	`json:"idle_timeout"`
	ConnectionDrainingTimeout	int	`json:"connection_draining_timeout"`
	CrossZoneLoadBalancing	bool	`json:"cross_zone_load_balancing"`
	Listener	interface{}	`json:"listener"`
	NamePrefix	string	`json:"name_prefix"`
}
