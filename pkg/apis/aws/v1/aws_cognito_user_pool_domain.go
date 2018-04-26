
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolDomain describes a AwsCognitoUserPoolDomain resource
type AwsCognitoUserPoolDomain struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCognitoUserPoolDomainSpec	`json:"spec"`
}


// AwsCognitoUserPoolDomainSpec is the spec for a AwsCognitoUserPoolDomain Resource
type AwsCognitoUserPoolDomainSpec struct {
	Domain	string	`json:"domain"`
	UserPoolId	string	`json:"user_pool_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolDomainList is a list of AwsCognitoUserPoolDomain resources
type AwsCognitoUserPoolDomainList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserPoolDomain	`json:"items"`
}

