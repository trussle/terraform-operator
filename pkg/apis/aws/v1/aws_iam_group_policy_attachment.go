
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamGroupPolicyAttachment describes a AwsIamGroupPolicyAttachment resource
type AwsIamGroupPolicyAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamGroupPolicyAttachmentSpec	`json:"spec"`
}


// AwsIamGroupPolicyAttachmentSpec is the spec for a AwsIamGroupPolicyAttachment Resource
type AwsIamGroupPolicyAttachmentSpec struct {
	Group	string	`json:"group"`
	PolicyArn	string	`json:"policy_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamGroupPolicyAttachmentList is a list of AwsIamGroupPolicyAttachment resources
type AwsIamGroupPolicyAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamGroupPolicyAttachment	`json:"items"`
}

