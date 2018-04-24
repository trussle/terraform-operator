
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRolePolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamRolePolicySpec	`json:"spec"`
}

type AwsIamRolePolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamRolePolicy	`json:"items"`
}

type AwsIamRolePolicySpec struct {
	Role	string	`json:"role"`
	Policy	string	`json:"policy"`
	NamePrefix	string	`json:"name_prefix"`
}
