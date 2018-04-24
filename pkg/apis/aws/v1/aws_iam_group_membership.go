
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroupMembership struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamGroupMembershipSpec	`json:"spec"`
}

type AwsIamGroupMembershipList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamGroupMembership	`json:"items"`
}

type AwsIamGroupMembershipSpec struct {
	Group	string	`json:"group"`
	Name	string	`json:"name"`
	Users	interface{}	`json:"users"`
}
