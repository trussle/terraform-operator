
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamAccountAlias describes a AwsIamAccountAlias resource
type AwsIamAccountAlias struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamAccountAliasSpec	`json:"spec"`
}


// AwsIamAccountAliasSpec is the spec for a AwsIamAccountAlias Resource
type AwsIamAccountAliasSpec struct {
	AccountAlias	string	`json:"account_alias"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamAccountAliasList is a list of AwsIamAccountAlias resources
type AwsIamAccountAliasList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamAccountAlias	`json:"items"`
}

