
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsAlias struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKmsAliasSpec	`json:"spec"`
}

type AwsKmsAliasList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKmsAlias	`json:"items"`
}

type AwsKmsAliasSpec struct {
	Name	string	`json:"name"`
	NamePrefix	string	`json:"name_prefix"`
	TargetKeyId	string	`json:"target_key_id"`
}
