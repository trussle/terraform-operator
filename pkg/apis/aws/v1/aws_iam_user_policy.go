
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamUserPolicySpec	`json:"spec"`
}

type AwsIamUserPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamUserPolicy	`json:"items"`
}

type AwsIamUserPolicySpec struct {
	Policy	string	`json:"policy"`
	NamePrefix	string	`json:"name_prefix"`
	User	string	`json:"user"`
}
