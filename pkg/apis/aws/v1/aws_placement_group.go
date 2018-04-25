
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsPlacementGroup describes a AwsPlacementGroup resource
type AwsPlacementGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsPlacementGroupSpec	`json:"spec"`
}


// AwsPlacementGroupSpec is the spec for a AwsPlacementGroup Resource
type AwsPlacementGroupSpec struct {
	Name	string	`json:"name"`
	Strategy	string	`json:"strategy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsPlacementGroupList is a list of AwsPlacementGroup resources
type AwsPlacementGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsPlacementGroup	`json:"items"`
}

