
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayGatewayResponse describes a AwsApiGatewayGatewayResponse resource
type AwsApiGatewayGatewayResponse struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayGatewayResponseSpec	`json:"spec"`
}


// AwsApiGatewayGatewayResponseSpec is the spec for a AwsApiGatewayGatewayResponse Resource
type AwsApiGatewayGatewayResponseSpec struct {
	ResponseTemplates	map[string]interface{}	`json:"response_templates"`
	ResponseParameters	map[string]interface{}	`json:"response_parameters"`
	RestApiId	string	`json:"rest_api_id"`
	ResponseType	string	`json:"response_type"`
	StatusCode	string	`json:"status_code"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayGatewayResponseList is a list of AwsApiGatewayGatewayResponse resources
type AwsApiGatewayGatewayResponseList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayGatewayResponse	`json:"items"`
}

