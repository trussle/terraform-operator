
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserSshKey describes a AwsIamUserSshKey resource
type AwsIamUserSshKey struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamUserSshKeySpec	`json:"spec"`
}


// AwsIamUserSshKeySpec is the spec for a AwsIamUserSshKey Resource
type AwsIamUserSshKeySpec struct {
	Encoding	string	`json:"encoding"`
	Username	string	`json:"username"`
	PublicKey	string	`json:"public_key"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserSshKeyList is a list of AwsIamUserSshKey resources
type AwsIamUserSshKeyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamUserSshKey	`json:"items"`
}

