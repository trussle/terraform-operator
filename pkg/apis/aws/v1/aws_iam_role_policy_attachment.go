
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamRolePolicyAttachment describes a AwsIamRolePolicyAttachment resource
type AwsIamRolePolicyAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamRolePolicyAttachmentSpec	`json:"spec"`
}


// AwsIamRolePolicyAttachmentSpec is the spec for a AwsIamRolePolicyAttachment Resource
type AwsIamRolePolicyAttachmentSpec struct {
	Role	string	`json:"role"`
	PolicyArn	string	`json:"policy_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamRolePolicyAttachmentList is a list of AwsIamRolePolicyAttachment resources
type AwsIamRolePolicyAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamRolePolicyAttachment	`json:"items"`
}

