
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	AlarmConfiguration	[]AwsCodedeployDeploymentGroupAlarmConfiguration	`json:"alarm_configuration"`
	AutoRollbackConfiguration	[]AwsCodedeployDeploymentGroupAutoRollbackConfiguration	`json:"auto_rollback_configuration"`
	AppName	string	`json:"app_name"`
	DeploymentGroupName	string	`json:"deployment_group_name"`
	DeploymentConfigName	string	`json:"deployment_config_name"`
	Ec2TagFilter	AwsCodedeployDeploymentGroupEc2TagFilter	`json:"ec2_tag_filter"`
	TriggerConfiguration	AwsCodedeployDeploymentGroupTriggerConfiguration	`json:"trigger_configuration"`
	AutoscalingGroups	string	`json:"autoscaling_groups"`
	OnPremisesInstanceTagFilter	AwsCodedeployDeploymentGroupOnPremisesInstanceTagFilter	`json:"on_premises_instance_tag_filter"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodedeployDeploymentGroupList is a list of AwsCodedeployDeploymentGroup resources
type AwsCodedeployDeploymentGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodedeployDeploymentGroup	`json:"items"`
}


// AwsCodedeployDeploymentGroupTargetGroupInfo is a AwsCodedeployDeploymentGroupTargetGroupInfo
type AwsCodedeployDeploymentGroupTargetGroupInfo struct {
	Name	string	`json:"name"`
}

// AwsCodedeployDeploymentGroupTriggerConfiguration is a AwsCodedeployDeploymentGroupTriggerConfiguration
type AwsCodedeployDeploymentGroupTriggerConfiguration struct {
	TriggerEvents	string	`json:"trigger_events"`
	TriggerName	string	`json:"trigger_name"`
	TriggerTargetArn	string	`json:"trigger_target_arn"`
}

// AwsCodedeployDeploymentGroupBlueGreenDeploymentConfig is a AwsCodedeployDeploymentGroupBlueGreenDeploymentConfig
type AwsCodedeployDeploymentGroupBlueGreenDeploymentConfig struct {
	DeploymentReadyOption	[]AwsCodedeployDeploymentGroupDeploymentReadyOption	`json:"deployment_ready_option"`
	TerminateBlueInstancesOnDeploymentSuccess	[]AwsCodedeployDeploymentGroupTerminateBlueInstancesOnDeploymentSuccess	`json:"terminate_blue_instances_on_deployment_success"`
}

// AwsCodedeployDeploymentGroupAlarmConfiguration is a AwsCodedeployDeploymentGroupAlarmConfiguration
type AwsCodedeployDeploymentGroupAlarmConfiguration struct {
	IgnorePollAlarmFailure	bool	`json:"ignore_poll_alarm_failure"`
	Alarms	string	`json:"alarms"`
	Enabled	bool	`json:"enabled"`
}

// AwsCodedeployDeploymentGroupElbInfo is a AwsCodedeployDeploymentGroupElbInfo
type AwsCodedeployDeploymentGroupElbInfo struct {
	Name	string	`json:"name"`
}

// AwsCodedeployDeploymentGroupLoadBalancerInfo is a AwsCodedeployDeploymentGroupLoadBalancerInfo
type AwsCodedeployDeploymentGroupLoadBalancerInfo struct {
	ElbInfo	AwsCodedeployDeploymentGroupElbInfo	`json:"elb_info"`
	TargetGroupInfo	AwsCodedeployDeploymentGroupTargetGroupInfo	`json:"target_group_info"`
}

// AwsCodedeployDeploymentGroupDeploymentStyle is a AwsCodedeployDeploymentGroupDeploymentStyle
type AwsCodedeployDeploymentGroupDeploymentStyle struct {
	DeploymentType	string	`json:"deployment_type"`
	DeploymentOption	string	`json:"deployment_option"`
}

// AwsCodedeployDeploymentGroupDeploymentReadyOption is a AwsCodedeployDeploymentGroupDeploymentReadyOption
type AwsCodedeployDeploymentGroupDeploymentReadyOption struct {
	ActionOnTimeout	string	`json:"action_on_timeout"`
	WaitTimeInMinutes	int	`json:"wait_time_in_minutes"`
}

// AwsCodedeployDeploymentGroupGreenFleetProvisioningOption is a AwsCodedeployDeploymentGroupGreenFleetProvisioningOption
type AwsCodedeployDeploymentGroupGreenFleetProvisioningOption struct {
	Action	string	`json:"action"`
}

// AwsCodedeployDeploymentGroupTerminateBlueInstancesOnDeploymentSuccess is a AwsCodedeployDeploymentGroupTerminateBlueInstancesOnDeploymentSuccess
type AwsCodedeployDeploymentGroupTerminateBlueInstancesOnDeploymentSuccess struct {
	Action	string	`json:"action"`
	TerminationWaitTimeInMinutes	int	`json:"termination_wait_time_in_minutes"`
}

// AwsCodedeployDeploymentGroupAutoRollbackConfiguration is a AwsCodedeployDeploymentGroupAutoRollbackConfiguration
type AwsCodedeployDeploymentGroupAutoRollbackConfiguration struct {
	Enabled	bool	`json:"enabled"`
	Events	string	`json:"events"`
}

// AwsCodedeployDeploymentGroupEc2TagFilter is a AwsCodedeployDeploymentGroupEc2TagFilter
type AwsCodedeployDeploymentGroupEc2TagFilter struct {
	Key	string	`json:"key"`
	Type	string	`json:"type"`
	Value	string	`json:"value"`
}

// AwsCodedeployDeploymentGroupOnPremisesInstanceTagFilter is a AwsCodedeployDeploymentGroupOnPremisesInstanceTagFilter
type AwsCodedeployDeploymentGroupOnPremisesInstanceTagFilter struct {
	Type	string	`json:"type"`
	Value	string	`json:"value"`
	Key	string	`json:"key"`
}
