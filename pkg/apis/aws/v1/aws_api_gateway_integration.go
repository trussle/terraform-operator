
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Type	string	`json:"type"`
	HttpMethod	string	`json:"http_method"`
	IntegrationHttpMethod	string	`json:"integration_http_method"`
	RequestParameters	map[string]???	`json:"request_parameters"`
	RequestParametersInJson	string	`json:"request_parameters_in_json"`
	CacheKeyParameters	string	`json:"cache_key_parameters"`
	RestApiId	string	`json:"rest_api_id"`
	ResourceId	string	`json:"resource_id"`
	Uri	string	`json:"uri"`
	Credentials	string	`json:"credentials"`
	RequestTemplates	map[string]???	`json:"request_templates"`
	ContentHandling	string	`json:"content_handling"`
	ConnectionType	string	`json:"connection_type"`
	ConnectionId	string	`json:"connection_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayIntegrationList is a list of AwsApiGatewayIntegration resources
type AwsApiGatewayIntegrationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayIntegration	`json:"items"`
}

