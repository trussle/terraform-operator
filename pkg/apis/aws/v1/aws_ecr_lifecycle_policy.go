
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcrLifecyclePolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEcrLifecyclePolicySpec	`json:"spec"`
}

type AwsEcrLifecyclePolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcrLifecyclePolicy	`json:"items"`
}

type AwsEcrLifecyclePolicySpec struct {
	Repository	string	`json:"repository"`
	Policy	string	`json:"policy"`
}
