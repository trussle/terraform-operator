
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmiLaunchPermission struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAmiLaunchPermissionSpec	`json:"spec"`
}

type AwsAmiLaunchPermissionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAmiLaunchPermission	`json:"items"`
}

type AwsAmiLaunchPermissionSpec struct {
	ImageId	string	`json:"image_id"`
	AccountId	string	`json:"account_id"`
}
