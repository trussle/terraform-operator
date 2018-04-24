
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayMethod struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayMethodSpec	`json:"spec"`
}

type AwsApiGatewayMethodList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayMethod	`json:"items"`
}

type AwsApiGatewayMethodSpec struct {
	ResourceId	string	`json:"resource_id"`
	HttpMethod	string	`json:"http_method"`
	Authorization	string	`json:"authorization"`
	RequestValidatorId	string	`json:"request_validator_id"`
	RestApiId	string	`json:"rest_api_id"`
	AuthorizerId	string	`json:"authorizer_id"`
	ApiKeyRequired	bool	`json:"api_key_required"`
	RequestModels	map[string]interface{}	`json:"request_models"`
	RequestParameters	map[string]interface{}	`json:"request_parameters"`
	RequestParametersInJson	string	`json:"request_parameters_in_json"`
}
