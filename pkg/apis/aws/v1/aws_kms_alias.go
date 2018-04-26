
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKmsAlias describes a AwsKmsAlias resource
type AwsKmsAlias struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKmsAliasSpec	`json:"spec"`
}


// AwsKmsAliasSpec is the spec for a AwsKmsAlias Resource
type AwsKmsAliasSpec struct {
	TargetKeyId	string	`json:"target_key_id"`
	Name	string	`json:"name"`
	NamePrefix	string	`json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKmsAliasList is a list of AwsKmsAlias resources
type AwsKmsAliasList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKmsAlias	`json:"items"`
}

