
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingGroup describes a AwsAutoscalingGroup resource
type AwsAutoscalingGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAutoscalingGroupSpec	`json:"spec"`
}


// AwsAutoscalingGroupSpec is the spec for a AwsAutoscalingGroup Resource
type AwsAutoscalingGroupSpec struct {
	PlacementGroup	string	`json:"placement_group"`
	WaitForCapacityTimeout	string	`json:"wait_for_capacity_timeout"`
	ProtectFromScaleIn	bool	`json:"protect_from_scale_in"`
	NamePrefix	string	`json:"name_prefix"`
	MinElbCapacity	int	`json:"min_elb_capacity"`
	WaitForElbCapacity	int	`json:"wait_for_elb_capacity"`
	MaxSize	int	`json:"max_size"`
	HealthCheckGracePeriod	int	`json:"health_check_grace_period"`
	Tags	[]map[string]string	`json:"tags"`
	TerminationPolicies	[]string	`json:"termination_policies"`
	InitialLifecycleHook	AwsAutoscalingGroupInitialLifecycleHook	`json:"initial_lifecycle_hook"`
	MinSize	int	`json:"min_size"`
	LaunchConfiguration	string	`json:"launch_configuration"`
	ForceDelete	bool	`json:"force_delete"`
	EnabledMetrics	string	`json:"enabled_metrics"`
	SuspendedProcesses	string	`json:"suspended_processes"`
	Tag	AwsAutoscalingGroupTag	`json:"tag"`
	MetricsGranularity	string	`json:"metrics_granularity"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingGroupList is a list of AwsAutoscalingGroup resources
type AwsAutoscalingGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingGroup	`json:"items"`
}


// AwsAutoscalingGroupInitialLifecycleHook is a AwsAutoscalingGroupInitialLifecycleHook
type AwsAutoscalingGroupInitialLifecycleHook struct {
	NotificationMetadata	string	`json:"notification_metadata"`
	NotificationTargetArn	string	`json:"notification_target_arn"`
	RoleArn	string	`json:"role_arn"`
	Name	string	`json:"name"`
	HeartbeatTimeout	int	`json:"heartbeat_timeout"`
	LifecycleTransition	string	`json:"lifecycle_transition"`
}

// AwsAutoscalingGroupTag is a AwsAutoscalingGroupTag
type AwsAutoscalingGroupTag struct {
	PropagateAtLaunch	bool	`json:"propagate_at_launch"`
	Key	string	`json:"key"`
	Value	string	`json:"value"`
}
