
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayMethod describes a AwsApiGatewayMethod resource
type AwsApiGatewayMethod struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayMethodSpec	`json:"spec"`
}


// AwsApiGatewayMethodSpec is the spec for a AwsApiGatewayMethod Resource
type AwsApiGatewayMethodSpec struct {
	RestApiId	string	`json:"rest_api_id"`
	Authorization	string	`json:"authorization"`
	AuthorizerId	string	`json:"authorizer_id"`
	ApiKeyRequired	bool	`json:"api_key_required"`
	RequestModels	map[string]interface{}	`json:"request_models"`
	ResourceId	string	`json:"resource_id"`
	HttpMethod	string	`json:"http_method"`
	RequestParameters	map[string]interface{}	`json:"request_parameters"`
	RequestParametersInJson	string	`json:"request_parameters_in_json"`
	RequestValidatorId	string	`json:"request_validator_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayMethodList is a list of AwsApiGatewayMethod resources
type AwsApiGatewayMethodList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayMethod	`json:"items"`
}

