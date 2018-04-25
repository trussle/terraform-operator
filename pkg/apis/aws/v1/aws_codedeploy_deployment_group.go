
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodedeployDeploymentGroup describes a AwsCodedeployDeploymentGroup resource
type AwsCodedeployDeploymentGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCodedeployDeploymentGroupSpec	`json:"spec"`
}


// AwsCodedeployDeploymentGroupSpec is the spec for a AwsCodedeployDeploymentGroup Resource
type AwsCodedeployDeploymentGroupSpec struct {
	ServiceRoleArn	string	`json:"service_role_arn"`
	AutoRollbackConfiguration	[]interface{}	`json:"auto_rollback_configuration"`
	AutoscalingGroups	string	`json:"autoscaling_groups"`
	Ec2TagFilter	string	`json:"ec2_tag_filter"`
	OnPremisesInstanceTagFilter	string	`json:"on_premises_instance_tag_filter"`
	AppName	string	`json:"app_name"`
	AlarmConfiguration	[]interface{}	`json:"alarm_configuration"`
	DeploymentConfigName	string	`json:"deployment_config_name"`
	TriggerConfiguration	string	`json:"trigger_configuration"`
	DeploymentGroupName	string	`json:"deployment_group_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodedeployDeploymentGroupList is a list of AwsCodedeployDeploymentGroup resources
type AwsCodedeployDeploymentGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodedeployDeploymentGroup	`json:"items"`
}

