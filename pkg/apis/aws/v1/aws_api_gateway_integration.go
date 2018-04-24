
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayIntegration struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayIntegrationSpec	`json:"spec"`
}

type AwsApiGatewayIntegrationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayIntegration	`json:"items"`
}

type AwsApiGatewayIntegrationSpec struct {
	Credentials	string	`json:"credentials"`
	ContentHandling	string	`json:"content_handling"`
	RestApiId	string	`json:"rest_api_id"`
	ConnectionId	string	`json:"connection_id"`
	Uri	string	`json:"uri"`
	ResourceId	string	`json:"resource_id"`
	HttpMethod	string	`json:"http_method"`
	RequestParametersInJson	string	`json:"request_parameters_in_json"`
	CacheKeyParameters	interface{}	`json:"cache_key_parameters"`
	RequestTemplates	map[string]interface{}	`json:"request_templates"`
	RequestParameters	map[string]interface{}	`json:"request_parameters"`
	IntegrationHttpMethod	string	`json:"integration_http_method"`
	Type	string	`json:"type"`
	ConnectionType	string	`json:"connection_type"`
}
