
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
	AlarmConfiguration	[]jzaLbtZs	`json:"alarm_configuration"`
	AutoRollbackConfiguration	[]zQMDQiYC	`json:"auto_rollback_configuration"`
	DeploymentConfigName	string	`json:"deployment_config_name"`
	OnPremisesInstanceTagFilter	string	`json:"on_premises_instance_tag_filter"`
	AppName	string	`json:"app_name"`
	TriggerConfiguration	string	`json:"trigger_configuration"`
	Ec2TagFilter	string	`json:"ec2_tag_filter"`
	DeploymentGroupName	string	`json:"deployment_group_name"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	AutoscalingGroups	string	`json:"autoscaling_groups"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodedeployDeploymentGroupList is a list of AwsCodedeployDeploymentGroup resources
type AwsCodedeployDeploymentGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodedeployDeploymentGroup	`json:"items"`
}


// jzaLbtZs is a jzaLbtZs
type jzaLbtZs struct {
	Alarms	string	`json:"alarms"`
	Enabled	bool	`json:"enabled"`
	IgnorePollAlarmFailure	bool	`json:"ignore_poll_alarm_failure"`
}

// yMGeuDtR is a yMGeuDtR
type yMGeuDtR struct {
	ElbInfo	string	`json:"elb_info"`
	TargetGroupInfo	string	`json:"target_group_info"`
}

// zQMDQiYC is a zQMDQiYC
type zQMDQiYC struct {
	Enabled	bool	`json:"enabled"`
	Events	string	`json:"events"`
}

// OhgHOvgS is a OhgHOvgS
type OhgHOvgS struct {
	DeploymentOption	string	`json:"deployment_option"`
	DeploymentType	string	`json:"deployment_type"`
}

// NufNjJhh is a NufNjJhh
type NufNjJhh struct {
	Action	string	`json:"action"`
	TerminationWaitTimeInMinutes	int	`json:"termination_wait_time_in_minutes"`
}

// jUVRuSqf is a jUVRuSqf
type jUVRuSqf struct {
	ActionOnTimeout	string	`json:"action_on_timeout"`
	WaitTimeInMinutes	int	`json:"wait_time_in_minutes"`
}

// gqVMkPYV is a gqVMkPYV
type gqVMkPYV struct {
	Action	string	`json:"action"`
}

// eycJPJHY is a eycJPJHY
type eycJPJHY struct {
	TerminateBlueInstancesOnDeploymentSuccess	[]NufNjJhh	`json:"terminate_blue_instances_on_deployment_success"`
	DeploymentReadyOption	[]jUVRuSqf	`json:"deployment_ready_option"`
}
