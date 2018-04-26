
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayRequestValidator describes a AwsApiGatewayRequestValidator resource
type AwsApiGatewayRequestValidator struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayRequestValidatorSpec	`json:"spec"`
}


// AwsApiGatewayRequestValidatorSpec is the spec for a AwsApiGatewayRequestValidator Resource
type AwsApiGatewayRequestValidatorSpec struct {
	RestApiId	string	`json:"rest_api_id"`
	Name	string	`json:"name"`
	ValidateRequestBody	bool	`json:"validate_request_body"`
	ValidateRequestParameters	bool	`json:"validate_request_parameters"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayRequestValidatorList is a list of AwsApiGatewayRequestValidator resources
type AwsApiGatewayRequestValidatorList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayRequestValidator	`json:"items"`
}

