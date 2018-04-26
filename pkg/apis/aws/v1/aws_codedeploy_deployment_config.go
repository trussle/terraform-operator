
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodedeployDeploymentConfig describes a AwsCodedeployDeploymentConfig resource
type AwsCodedeployDeploymentConfig struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCodedeployDeploymentConfigSpec	`json:"spec"`
}


// AwsCodedeployDeploymentConfigSpec is the spec for a AwsCodedeployDeploymentConfig Resource
type AwsCodedeployDeploymentConfigSpec struct {
	DeploymentConfigName	string	`json:"deployment_config_name"`
	MinimumHealthyHosts	[]EYeqkKHq	`json:"minimum_healthy_hosts"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodedeployDeploymentConfigList is a list of AwsCodedeployDeploymentConfig resources
type AwsCodedeployDeploymentConfigList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodedeployDeploymentConfig	`json:"items"`
}


// EYeqkKHq is a EYeqkKHq
type EYeqkKHq struct {
	Value	int	`json:"value"`
	Type	string	`json:"type"`
}
