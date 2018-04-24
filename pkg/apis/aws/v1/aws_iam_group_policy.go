
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroupPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamGroupPolicySpec	`json:"spec"`
}

type AwsIamGroupPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamGroupPolicy	`json:"items"`
}

type AwsIamGroupPolicySpec struct {
	Policy	string	`json:"policy"`
	NamePrefix	string	`json:"name_prefix"`
	Group	string	`json:"group"`
}
