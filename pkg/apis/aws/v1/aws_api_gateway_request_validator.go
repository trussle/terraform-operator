
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayRequestValidator struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayRequestValidatorSpec	`json:"spec"`
}

type AwsApiGatewayRequestValidatorList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayRequestValidator	`json:"items"`
}

type AwsApiGatewayRequestValidatorSpec struct {
	ValidateRequestParameters	bool	`json:"validate_request_parameters"`
	RestApiId	string	`json:"rest_api_id"`
	Name	string	`json:"name"`
	ValidateRequestBody	bool	`json:"validate_request_body"`
}
