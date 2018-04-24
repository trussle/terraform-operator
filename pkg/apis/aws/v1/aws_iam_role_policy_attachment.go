
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRolePolicyAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamRolePolicyAttachmentSpec	`json:"spec"`
}

type AwsIamRolePolicyAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamRolePolicyAttachment	`json:"items"`
}

type AwsIamRolePolicyAttachmentSpec struct {
	Role	string	`json:"role"`
	PolicyArn	string	`json:"policy_arn"`
}
