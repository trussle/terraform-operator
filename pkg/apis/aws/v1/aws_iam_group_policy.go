
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamGroupPolicy describes a AwsIamGroupPolicy resource
type AwsIamGroupPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamGroupPolicySpec	`json:"spec"`
}


// AwsIamGroupPolicySpec is the spec for a AwsIamGroupPolicy Resource
type AwsIamGroupPolicySpec struct {
	Group	string	`json:"group"`
	Policy	string	`json:"policy"`
	NamePrefix	string	`json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamGroupPolicyList is a list of AwsIamGroupPolicy resources
type AwsIamGroupPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamGroupPolicy	`json:"items"`
}

