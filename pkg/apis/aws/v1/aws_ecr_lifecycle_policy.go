
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcrLifecyclePolicy describes a AwsEcrLifecyclePolicy resource
type AwsEcrLifecyclePolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEcrLifecyclePolicySpec	`json:"spec"`
}


// AwsEcrLifecyclePolicySpec is the spec for a AwsEcrLifecyclePolicy Resource
type AwsEcrLifecyclePolicySpec struct {
	Policy	string	`json:"policy"`
	Repository	string	`json:"repository"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcrLifecyclePolicyList is a list of AwsEcrLifecyclePolicy resources
type AwsEcrLifecyclePolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcrLifecyclePolicy	`json:"items"`
}

