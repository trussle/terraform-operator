
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserPolicyAttachment describes a AwsIamUserPolicyAttachment resource
type AwsIamUserPolicyAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamUserPolicyAttachmentSpec	`json:"spec"`
}


// AwsIamUserPolicyAttachmentSpec is the spec for a AwsIamUserPolicyAttachment Resource
type AwsIamUserPolicyAttachmentSpec struct {
	User	string	`json:"user"`
	PolicyArn	string	`json:"policy_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserPolicyAttachmentList is a list of AwsIamUserPolicyAttachment resources
type AwsIamUserPolicyAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamUserPolicyAttachment	`json:"items"`
}

