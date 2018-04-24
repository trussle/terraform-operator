
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksUserProfile struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksUserProfileSpec	`json:"spec"`
}

type AwsOpsworksUserProfileList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksUserProfile	`json:"items"`
}

type AwsOpsworksUserProfileSpec struct {
	UserArn	string	`json:"user_arn"`
	AllowSelfManagement	bool	`json:"allow_self_management"`
	SshUsername	string	`json:"ssh_username"`
	SshPublicKey	string	`json:"ssh_public_key"`
}
