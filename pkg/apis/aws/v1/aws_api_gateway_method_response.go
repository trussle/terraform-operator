
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayMethodResponse describes a AwsApiGatewayMethodResponse resource
type AwsApiGatewayMethodResponse struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayMethodResponseSpec	`json:"spec"`
}


// AwsApiGatewayMethodResponseSpec is the spec for a AwsApiGatewayMethodResponse Resource
type AwsApiGatewayMethodResponseSpec struct {
	RestApiId	string	`json:"rest_api_id"`
	ResourceId	string	`json:"resource_id"`
	HttpMethod	string	`json:"http_method"`
	StatusCode	string	`json:"status_code"`
	ResponseModels	map[string]???	`json:"response_models"`
	ResponseParameters	map[string]???	`json:"response_parameters"`
	ResponseParametersInJson	string	`json:"response_parameters_in_json"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayMethodResponseList is a list of AwsApiGatewayMethodResponse resources
type AwsApiGatewayMethodResponseList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayMethodResponse	`json:"items"`
}

