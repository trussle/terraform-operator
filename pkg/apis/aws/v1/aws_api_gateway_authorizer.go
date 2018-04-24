
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayAuthorizer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayAuthorizerSpec	`json:"spec"`
}

type AwsApiGatewayAuthorizerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayAuthorizer	`json:"items"`
}

type AwsApiGatewayAuthorizerSpec struct {
	IdentitySource	string	`json:"identity_source"`
	Type	string	`json:"type"`
	AuthorizerResultTtlInSeconds	int	`json:"authorizer_result_ttl_in_seconds"`
	IdentityValidationExpression	string	`json:"identity_validation_expression"`
	AuthorizerUri	string	`json:"authorizer_uri"`
	Name	string	`json:"name"`
	RestApiId	string	`json:"rest_api_id"`
	AuthorizerCredentials	string	`json:"authorizer_credentials"`
	ProviderArns	interface{}	`json:"provider_arns"`
}
