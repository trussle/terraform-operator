
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
	AppName	string	`json:"app_name"`
	DeploymentGroupName	string	`json:"deployment_group_name"`
	AutoscalingGroups	Generic	`json:"autoscaling_groups"`
	DeploymentConfigName	string	`json:"deployment_config_name"`
	Ec2TagFilter	Generic	`json:"ec2_tag_filter"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	AlarmConfiguration	[]Generic	`json:"alarm_configuration"`
	AutoRollbackConfiguration	[]Generic	`json:"auto_rollback_configuration"`
	OnPremisesInstanceTagFilter	Generic	`json:"on_premises_instance_tag_filter"`
	TriggerConfiguration	Generic	`json:"trigger_configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodedeployDeploymentGroupList is a list of AwsCodedeployDeploymentGroup resources
type AwsCodedeployDeploymentGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodedeployDeploymentGroup	`json:"items"`
}

