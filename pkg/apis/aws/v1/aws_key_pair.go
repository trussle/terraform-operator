
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKeyPair struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKeyPairSpec	`json:"spec"`
}

type AwsKeyPairList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKeyPair	`json:"items"`
}

type AwsKeyPairSpec struct {
	KeyNamePrefix	string	`json:"key_name_prefix"`
	PublicKey	string	`json:"public_key"`
}
