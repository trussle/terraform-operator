
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayAuthorizer describes a AwsApiGatewayAuthorizer resource
type AwsApiGatewayAuthorizer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayAuthorizerSpec	`json:"spec"`
}


// AwsApiGatewayAuthorizerSpec is the spec for a AwsApiGatewayAuthorizer Resource
type AwsApiGatewayAuthorizerSpec struct {
	AuthorizerUri	string	`json:"authorizer_uri"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
	AuthorizerCredentials	string	`json:"authorizer_credentials"`
	AuthorizerResultTtlInSeconds	int	`json:"authorizer_result_ttl_in_seconds"`
	IdentityValidationExpression	string	`json:"identity_validation_expression"`
	IdentitySource	string	`json:"identity_source"`
	RestApiId	string	`json:"rest_api_id"`
	ProviderArns	Generic	`json:"provider_arns"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayAuthorizerList is a list of AwsApiGatewayAuthorizer resources
type AwsApiGatewayAuthorizerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayAuthorizer	`json:"items"`
}

