
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegrationResponse struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayIntegrationResponseSpec	`json:"spec"`
}

type AwsApiGatewayIntegrationResponseList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayIntegrationResponse	`json:"items"`
}

type AwsApiGatewayIntegrationResponseSpec struct {
	StatusCode	string	`json:"status_code"`
	ResponseTemplates	map[string]interface{}	`json:"response_templates"`
	ResponseParameters	map[string]interface{}	`json:"response_parameters"`
	ContentHandling	string	`json:"content_handling"`
	ResourceId	string	`json:"resource_id"`
	HttpMethod	string	`json:"http_method"`
	ResponseParametersInJson	string	`json:"response_parameters_in_json"`
	RestApiId	string	`json:"rest_api_id"`
	SelectionPattern	string	`json:"selection_pattern"`
}
