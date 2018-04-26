
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolClient describes a AwsCognitoUserPoolClient resource
type AwsCognitoUserPoolClient struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCognitoUserPoolClientSpec	`json:"spec"`
}


// AwsCognitoUserPoolClientSpec is the spec for a AwsCognitoUserPoolClient Resource
type AwsCognitoUserPoolClientSpec struct {
	ExplicitAuthFlows	Generic	`json:"explicit_auth_flows"`
	WriteAttributes	Generic	`json:"write_attributes"`
	AllowedOauthScopes	Generic	`json:"allowed_oauth_scopes"`
	GenerateSecret	bool	`json:"generate_secret"`
	CallbackUrls	[]Generic	`json:"callback_urls"`
	DefaultRedirectUri	string	`json:"default_redirect_uri"`
	SupportedIdentityProviders	[]Generic	`json:"supported_identity_providers"`
	AllowedOauthFlows	Generic	`json:"allowed_oauth_flows"`
	AllowedOauthFlowsUserPoolClient	bool	`json:"allowed_oauth_flows_user_pool_client"`
	LogoutUrls	[]Generic	`json:"logout_urls"`
	Name	string	`json:"name"`
	UserPoolId	string	`json:"user_pool_id"`
	ReadAttributes	Generic	`json:"read_attributes"`
	RefreshTokenValidity	int	`json:"refresh_token_validity"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolClientList is a list of AwsCognitoUserPoolClient resources
type AwsCognitoUserPoolClientList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserPoolClient	`json:"items"`
}

