
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDeployment describes a AwsApiGatewayDeployment resource
type AwsApiGatewayDeployment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayDeploymentSpec	`json:"spec"`
}


// AwsApiGatewayDeploymentSpec is the spec for a AwsApiGatewayDeployment Resource
type AwsApiGatewayDeploymentSpec struct {
	RestApiId	string	`json:"rest_api_id"`
	StageName	string	`json:"stage_name"`
	Description	string	`json:"description"`
	StageDescription	string	`json:"stage_description"`
	Variables	map[string]string	`json:"variables"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDeploymentList is a list of AwsApiGatewayDeployment resources
type AwsApiGatewayDeploymentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayDeployment	`json:"items"`
}

