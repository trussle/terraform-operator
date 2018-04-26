
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKmsGrant describes a AwsKmsGrant resource
type AwsKmsGrant struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKmsGrantSpec	`json:"spec"`
}


// AwsKmsGrantSpec is the spec for a AwsKmsGrant Resource
type AwsKmsGrantSpec struct {
	Name	string	`json:"name"`
	KeyId	string	`json:"key_id"`
	RetireOnDelete	bool	`json:"retire_on_delete"`
	GrantCreationTokens	string	`json:"grant_creation_tokens"`
	GranteePrincipal	string	`json:"grantee_principal"`
	Operations	string	`json:"operations"`
	Constraints	string	`json:"constraints"`
	RetiringPrincipal	string	`json:"retiring_principal"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKmsGrantList is a list of AwsKmsGrant resources
type AwsKmsGrantList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKmsGrant	`json:"items"`
}

