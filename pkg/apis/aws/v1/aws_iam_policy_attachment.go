
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamPolicyAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamPolicyAttachmentSpec	`json:"spec"`
}

type AwsIamPolicyAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamPolicyAttachment	`json:"items"`
}

type AwsIamPolicyAttachmentSpec struct {
	Name	string	`json:"name"`
	Users	interface{}	`json:"users"`
	Roles	interface{}	`json:"roles"`
	Groups	interface{}	`json:"groups"`
	PolicyArn	string	`json:"policy_arn"`
}
