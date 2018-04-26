
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
	DeploymentConfigName	string	`json:"deployment_config_name"`
	Ec2TagFilter	Ec2TagFilter	`json:"ec2_tag_filter"`
	DeploymentGroupName	string	`json:"deployment_group_name"`
	AlarmConfiguration	[]AlarmConfiguration	`json:"alarm_configuration"`
	AutoscalingGroups	string	`json:"autoscaling_groups"`
	TriggerConfiguration	TriggerConfiguration	`json:"trigger_configuration"`
	AppName	string	`json:"app_name"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	AutoRollbackConfiguration	[]AutoRollbackConfiguration	`json:"auto_rollback_configuration"`
	OnPremisesInstanceTagFilter	OnPremisesInstanceTagFilter	`json:"on_premises_instance_tag_filter"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodedeployDeploymentGroupList is a list of AwsCodedeployDeploymentGroup resources
type AwsCodedeployDeploymentGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodedeployDeploymentGroup	`json:"items"`
}


// TriggerConfiguration is a TriggerConfiguration
type TriggerConfiguration struct {
	TriggerTargetArn	string	`json:"trigger_target_arn"`
	TriggerEvents	string	`json:"trigger_events"`
	TriggerName	string	`json:"trigger_name"`
}

// DeploymentStyle is a DeploymentStyle
type DeploymentStyle struct {
	DeploymentOption	string	`json:"deployment_option"`
	DeploymentType	string	`json:"deployment_type"`
}

// DeploymentReadyOption is a DeploymentReadyOption
type DeploymentReadyOption struct {
	ActionOnTimeout	string	`json:"action_on_timeout"`
	WaitTimeInMinutes	int	`json:"wait_time_in_minutes"`
}

// GreenFleetProvisioningOption is a GreenFleetProvisioningOption
type GreenFleetProvisioningOption struct {
	Action	string	`json:"action"`
}

// TerminateBlueInstancesOnDeploymentSuccess is a TerminateBlueInstancesOnDeploymentSuccess
type TerminateBlueInstancesOnDeploymentSuccess struct {
	Action	string	`json:"action"`
	TerminationWaitTimeInMinutes	int	`json:"termination_wait_time_in_minutes"`
}

// BlueGreenDeploymentConfig is a BlueGreenDeploymentConfig
type BlueGreenDeploymentConfig struct {
	DeploymentReadyOption	[]DeploymentReadyOption	`json:"deployment_ready_option"`
	TerminateBlueInstancesOnDeploymentSuccess	[]TerminateBlueInstancesOnDeploymentSuccess	`json:"terminate_blue_instances_on_deployment_success"`
}

// ElbInfo is a ElbInfo
type ElbInfo struct {
	Name	string	`json:"name"`
}

// LoadBalancerInfo is a LoadBalancerInfo
type LoadBalancerInfo struct {
	ElbInfo	ElbInfo	`json:"elb_info"`
	TargetGroupInfo	TargetGroupInfo	`json:"target_group_info"`
}

// OnPremisesInstanceTagFilter is a OnPremisesInstanceTagFilter
type OnPremisesInstanceTagFilter struct {
	Key	string	`json:"key"`
	Type	string	`json:"type"`
	Value	string	`json:"value"`
}

// Ec2TagFilter is a Ec2TagFilter
type Ec2TagFilter struct {
	Value	string	`json:"value"`
	Key	string	`json:"key"`
	Type	string	`json:"type"`
}

// AlarmConfiguration is a AlarmConfiguration
type AlarmConfiguration struct {
	Alarms	string	`json:"alarms"`
	Enabled	bool	`json:"enabled"`
	IgnorePollAlarmFailure	bool	`json:"ignore_poll_alarm_failure"`
}

// TargetGroupInfo is a TargetGroupInfo
type TargetGroupInfo struct {
	Name	string	`json:"name"`
}

// AutoRollbackConfiguration is a AutoRollbackConfiguration
type AutoRollbackConfiguration struct {
	Enabled	bool	`json:"enabled"`
	Events	string	`json:"events"`
}
