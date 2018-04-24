
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodedeployDeploymentGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCodedeployDeploymentGroupSpec	`json:"spec"`
}

type AwsCodedeployDeploymentGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodedeployDeploymentGroup	`json:"items"`
}

type AwsCodedeployDeploymentGroupSpec struct {
	DeploymentConfigName	string	`json:"deployment_config_name"`
	OnPremisesInstanceTagFilter	interface{}	`json:"on_premises_instance_tag_filter"`
	TriggerConfiguration	interface{}	`json:"trigger_configuration"`
	DeploymentGroupName	string	`json:"deployment_group_name"`
	AlarmConfiguration	[]interface{}	`json:"alarm_configuration"`
	AutoRollbackConfiguration	[]interface{}	`json:"auto_rollback_configuration"`
	AutoscalingGroups	interface{}	`json:"autoscaling_groups"`
	Ec2TagFilter	interface{}	`json:"ec2_tag_filter"`
	AppName	string	`json:"app_name"`
	ServiceRoleArn	string	`json:"service_role_arn"`
}
