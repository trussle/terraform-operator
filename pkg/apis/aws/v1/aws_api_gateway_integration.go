
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
	Uri	string	`json:"uri"`
	CacheKeyParameters	Generic	`json:"cache_key_parameters"`
	RestApiId	string	`json:"rest_api_id"`
	ResourceId	string	`json:"resource_id"`
	HttpMethod	string	`json:"http_method"`
	ConnectionType	string	`json:"connection_type"`
	RequestTemplates	map[string]Generic	`json:"request_templates"`
	RequestParametersInJson	string	`json:"request_parameters_in_json"`
	Type	string	`json:"type"`
	ConnectionId	string	`json:"connection_id"`
	Credentials	string	`json:"credentials"`
	RequestParameters	map[string]Generic	`json:"request_parameters"`
	ContentHandling	string	`json:"content_handling"`
	IntegrationHttpMethod	string	`json:"integration_http_method"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayIntegrationList is a list of AwsApiGatewayIntegration resources
type AwsApiGatewayIntegrationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayIntegration	`json:"items"`
}

