
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Name	string	`json:"name"`
	Users	string	`json:"users"`
	Group	string	`json:"group"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamGroupMembershipList is a list of AwsIamGroupMembership resources
type AwsIamGroupMembershipList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamGroupMembership	`json:"items"`
}

