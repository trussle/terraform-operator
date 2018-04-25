
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserPolicy describes a AwsIamUserPolicy resource
type AwsIamUserPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamUserPolicySpec	`json:"spec"`
}


// AwsIamUserPolicySpec is the spec for a AwsIamUserPolicy Resource
type AwsIamUserPolicySpec struct {
	User	string	`json:"user"`
	Policy	string	`json:"policy"`
	NamePrefix	string	`json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserPolicyList is a list of AwsIamUserPolicy resources
type AwsIamUserPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamUserPolicy	`json:"items"`
}

