
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamGroupMembership describes a AwsIamGroupMembership resource
type AwsIamGroupMembership struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamGroupMembershipSpec	`json:"spec"`
}


// AwsIamGroupMembershipSpec is the spec for a AwsIamGroupMembership Resource
type AwsIamGroupMembershipSpec struct {
	Users	Generic	`json:"users"`
	Group	string	`json:"group"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamGroupMembershipList is a list of AwsIamGroupMembership resources
type AwsIamGroupMembershipList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamGroupMembership	`json:"items"`
}

