
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultRouteTable struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultRouteTableSpec	`json:"spec"`
}

type AwsDefaultRouteTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultRouteTable	`json:"items"`
}

type AwsDefaultRouteTableSpec struct {
	DefaultRouteTableId	string	`json:"default_route_table_id"`
	PropagatingVgws	interface{}	`json:"propagating_vgws"`
	Tags	map[string]interface{}	`json:"tags"`
}
