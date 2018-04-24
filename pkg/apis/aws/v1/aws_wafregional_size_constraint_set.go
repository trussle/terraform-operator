
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalSizeConstraintSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalSizeConstraintSetSpec	`json:"spec"`
}

type AwsWafregionalSizeConstraintSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalSizeConstraintSet	`json:"items"`
}

type AwsWafregionalSizeConstraintSetSpec struct {
	Name	string	`json:"name"`
	SizeConstraints	interface{}	`json:"size_constraints"`
}
