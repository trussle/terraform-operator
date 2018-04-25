
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafSizeConstraintSet describes a AwsWafSizeConstraintSet resource
type AwsWafSizeConstraintSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafSizeConstraintSetSpec	`json:"spec"`
}


// AwsWafSizeConstraintSetSpec is the spec for a AwsWafSizeConstraintSet Resource
type AwsWafSizeConstraintSetSpec struct {
	Name	string	`json:"name"`
	SizeConstraints	string	`json:"size_constraints"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafSizeConstraintSetList is a list of AwsWafSizeConstraintSet resources
type AwsWafSizeConstraintSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafSizeConstraintSet	`json:"items"`
}

