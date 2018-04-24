
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDeployment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayDeploymentSpec	`json:"spec"`
}

type AwsApiGatewayDeploymentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayDeployment	`json:"items"`
}

type AwsApiGatewayDeploymentSpec struct {
	StageDescription	string	`json:"stage_description"`
	Variables	map[string]interface{}	`json:"variables"`
	RestApiId	string	`json:"rest_api_id"`
	StageName	string	`json:"stage_name"`
	Description	string	`json:"description"`
}
