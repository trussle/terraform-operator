
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserSshKey struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamUserSshKeySpec	`json:"spec"`
}

type AwsIamUserSshKeyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamUserSshKey	`json:"items"`
}

type AwsIamUserSshKeySpec struct {
	Username	string	`json:"username"`
	PublicKey	string	`json:"public_key"`
	Encoding	string	`json:"encoding"`
}
