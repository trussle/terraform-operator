
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoIdentityPool describes a AwsCognitoIdentityPool resource
type AwsCognitoIdentityPool struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCognitoIdentityPoolSpec	`json:"spec"`
}


// AwsCognitoIdentityPoolSpec is the spec for a AwsCognitoIdentityPool Resource
type AwsCognitoIdentityPoolSpec struct {
	IdentityPoolName	string	`json:"identity_pool_name"`
	CognitoIdentityProviders	string	`json:"cognito_identity_providers"`
	DeveloperProviderName	string	`json:"developer_provider_name"`
	AllowUnauthenticatedIdentities	bool	`json:"allow_unauthenticated_identities"`
	OpenidConnectProviderArns	[]string	`json:"openid_connect_provider_arns"`
	SamlProviderArns	[]string	`json:"saml_provider_arns"`
	SupportedLoginProviders	map[string]string	`json:"supported_login_providers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoIdentityPoolList is a list of AwsCognitoIdentityPool resources
type AwsCognitoIdentityPoolList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoIdentityPool	`json:"items"`
}

