
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksUserProfile describes a AwsOpsworksUserProfile resource
type AwsOpsworksUserProfile struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksUserProfileSpec	`json:"spec"`
}


// AwsOpsworksUserProfileSpec is the spec for a AwsOpsworksUserProfile Resource
type AwsOpsworksUserProfileSpec struct {
	UserArn	string	`json:"user_arn"`
	AllowSelfManagement	bool	`json:"allow_self_management"`
	SshUsername	string	`json:"ssh_username"`
	SshPublicKey	string	`json:"ssh_public_key"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksUserProfileList is a list of AwsOpsworksUserProfile resources
type AwsOpsworksUserProfileList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksUserProfile	`json:"items"`
}

