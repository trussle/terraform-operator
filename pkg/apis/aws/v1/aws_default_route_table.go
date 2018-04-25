
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultRouteTable describes a AwsDefaultRouteTable resource
type AwsDefaultRouteTable struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultRouteTableSpec	`json:"spec"`
}


// AwsDefaultRouteTableSpec is the spec for a AwsDefaultRouteTable Resource
type AwsDefaultRouteTableSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	DefaultRouteTableId	string	`json:"default_route_table_id"`
	PropagatingVgws	string	`json:"propagating_vgws"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultRouteTableList is a list of AwsDefaultRouteTable resources
type AwsDefaultRouteTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultRouteTable	`json:"items"`
}

