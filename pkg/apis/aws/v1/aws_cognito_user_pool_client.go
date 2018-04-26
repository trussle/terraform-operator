
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	ReadAttributes	string	`json:"read_attributes"`
	AllowedOauthFlowsUserPoolClient	bool	`json:"allowed_oauth_flows_user_pool_client"`
	Name	string	`json:"name"`
	ExplicitAuthFlows	string	`json:"explicit_auth_flows"`
	CallbackUrls	[]string	`json:"callback_urls"`
	RefreshTokenValidity	int	`json:"refresh_token_validity"`
	AllowedOauthScopes	string	`json:"allowed_oauth_scopes"`
	DefaultRedirectUri	string	`json:"default_redirect_uri"`
	UserPoolId	string	`json:"user_pool_id"`
	WriteAttributes	string	`json:"write_attributes"`
	AllowedOauthFlows	string	`json:"allowed_oauth_flows"`
	LogoutUrls	[]string	`json:"logout_urls"`
	SupportedIdentityProviders	[]string	`json:"supported_identity_providers"`
	GenerateSecret	bool	`json:"generate_secret"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolClientList is a list of AwsCognitoUserPoolClient resources
type AwsCognitoUserPoolClientList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserPoolClient	`json:"items"`
}

