
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoIdentityPool struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCognitoIdentityPoolSpec	`json:"spec"`
}

type AwsCognitoIdentityPoolList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoIdentityPool	`json:"items"`
}

type AwsCognitoIdentityPoolSpec struct {
	DeveloperProviderName	string	`json:"developer_provider_name"`
	AllowUnauthenticatedIdentities	bool	`json:"allow_unauthenticated_identities"`
	OpenidConnectProviderArns	[]interface{}	`json:"openid_connect_provider_arns"`
	SamlProviderArns	[]interface{}	`json:"saml_provider_arns"`
	SupportedLoginProviders	map[string]interface{}	`json:"supported_login_providers"`
	IdentityPoolName	string	`json:"identity_pool_name"`
	CognitoIdentityProviders	interface{}	`json:"cognito_identity_providers"`
}
