
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodedeployDeploymentConfig struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCodedeployDeploymentConfigSpec	`json:"spec"`
}

type AwsCodedeployDeploymentConfigList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodedeployDeploymentConfig	`json:"items"`
}

type AwsCodedeployDeploymentConfigSpec struct {
	DeploymentConfigName	string	`json:"deployment_config_name"`
	MinimumHealthyHosts	[]interface{}	`json:"minimum_healthy_hosts"`
}
