
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOrganizationsPolicyAttachment describes a AwsOrganizationsPolicyAttachment resource
type AwsOrganizationsPolicyAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOrganizationsPolicyAttachmentSpec	`json:"spec"`
}


// AwsOrganizationsPolicyAttachmentSpec is the spec for a AwsOrganizationsPolicyAttachment Resource
type AwsOrganizationsPolicyAttachmentSpec struct {
	TargetId	string	`json:"target_id"`
	PolicyId	string	`json:"policy_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOrganizationsPolicyAttachmentList is a list of AwsOrganizationsPolicyAttachment resources
type AwsOrganizationsPolicyAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOrganizationsPolicyAttachment	`json:"items"`
}

