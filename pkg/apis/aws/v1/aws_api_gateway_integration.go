
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayIntegration describes a AwsApiGatewayIntegration resource
type AwsApiGatewayIntegration struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayIntegrationSpec	`json:"spec"`
}


// AwsApiGatewayIntegrationSpec is the spec for a AwsApiGatewayIntegration Resource
type AwsApiGatewayIntegrationSpec struct {
	ResourceId	string	`json:"resource_id"`
	ConnectionId	string	`json:"connection_id"`
	HttpMethod	string	`json:"http_method"`
	Type	string	`json:"type"`
	RequestTemplates	map[string]interface{}	`json:"request_templates"`
	ContentHandling	string	`json:"content_handling"`
	Uri	string	`json:"uri"`
	IntegrationHttpMethod	string	`json:"integration_http_method"`
	RequestParameters	map[string]interface{}	`json:"request_parameters"`
	RequestParametersInJson	string	`json:"request_parameters_in_json"`
	CacheKeyParameters	string	`json:"cache_key_parameters"`
	RestApiId	string	`json:"rest_api_id"`
	ConnectionType	string	`json:"connection_type"`
	Credentials	string	`json:"credentials"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayIntegrationList is a list of AwsApiGatewayIntegration resources
type AwsApiGatewayIntegrationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayIntegration	`json:"items"`
}

