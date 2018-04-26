
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamPolicyAttachment describes a AwsIamPolicyAttachment resource
type AwsIamPolicyAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamPolicyAttachmentSpec	`json:"spec"`
}


// AwsIamPolicyAttachmentSpec is the spec for a AwsIamPolicyAttachment Resource
type AwsIamPolicyAttachmentSpec struct {
	Roles	string	`json:"roles"`
	Groups	string	`json:"groups"`
	PolicyArn	string	`json:"policy_arn"`
	Name	string	`json:"name"`
	Users	string	`json:"users"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamPolicyAttachmentList is a list of AwsIamPolicyAttachment resources
type AwsIamPolicyAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamPolicyAttachment	`json:"items"`
}

