
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamGroupPolicyAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamGroupPolicyAttachmentSpec	`json:"spec"`
}

type AwsIamGroupPolicyAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamGroupPolicyAttachment	`json:"items"`
}

type AwsIamGroupPolicyAttachmentSpec struct {
	Group	string	`json:"group"`
	PolicyArn	string	`json:"policy_arn"`
}
