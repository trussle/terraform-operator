
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsGrant struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKmsGrantSpec	`json:"spec"`
}

type AwsKmsGrantList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKmsGrant	`json:"items"`
}

type AwsKmsGrantSpec struct {
	KeyId	string	`json:"key_id"`
	GranteePrincipal	string	`json:"grantee_principal"`
	Constraints	interface{}	`json:"constraints"`
	RetiringPrincipal	string	`json:"retiring_principal"`
	GrantCreationTokens	interface{}	`json:"grant_creation_tokens"`
	RetireOnDelete	bool	`json:"retire_on_delete"`
	Name	string	`json:"name"`
	Operations	interface{}	`json:"operations"`
}
