
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPoolClient struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCognitoUserPoolClientSpec	`json:"spec"`
}

type AwsCognitoUserPoolClientList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserPoolClient	`json:"items"`
}

type AwsCognitoUserPoolClientSpec struct {
	UserPoolId	string	`json:"user_pool_id"`
	ExplicitAuthFlows	interface{}	`json:"explicit_auth_flows"`
	WriteAttributes	interface{}	`json:"write_attributes"`
	RefreshTokenValidity	int	`json:"refresh_token_validity"`
	SupportedIdentityProviders	[]interface{}	`json:"supported_identity_providers"`
	AllowedOauthFlows	interface{}	`json:"allowed_oauth_flows"`
	CallbackUrls	[]interface{}	`json:"callback_urls"`
	LogoutUrls	[]interface{}	`json:"logout_urls"`
	Name	string	`json:"name"`
	GenerateSecret	bool	`json:"generate_secret"`
	ReadAttributes	interface{}	`json:"read_attributes"`
	AllowedOauthFlowsUserPoolClient	bool	`json:"allowed_oauth_flows_user_pool_client"`
	AllowedOauthScopes	interface{}	`json:"allowed_oauth_scopes"`
	DefaultRedirectUri	string	`json:"default_redirect_uri"`
}
