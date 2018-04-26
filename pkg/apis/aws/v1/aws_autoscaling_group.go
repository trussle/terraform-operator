
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
	SuspendedProcesses	string	`json:"suspended_processes"`
	ProtectFromScaleIn	bool	`json:"protect_from_scale_in"`
	MinSize	int	`json:"min_size"`
	MaxSize	int	`json:"max_size"`
	Tags	[]map[string]string	`json:"tags"`
	HealthCheckGracePeriod	int	`json:"health_check_grace_period"`
	WaitForCapacityTimeout	string	`json:"wait_for_capacity_timeout"`
	Tag	Tag	`json:"tag"`
	NamePrefix	string	`json:"name_prefix"`
	PlacementGroup	string	`json:"placement_group"`
	EnabledMetrics	string	`json:"enabled_metrics"`
	InitialLifecycleHook	InitialLifecycleHook	`json:"initial_lifecycle_hook"`
	LaunchConfiguration	string	`json:"launch_configuration"`
	MinElbCapacity	int	`json:"min_elb_capacity"`
	WaitForElbCapacity	int	`json:"wait_for_elb_capacity"`
	TerminationPolicies	[]string	`json:"termination_policies"`
	MetricsGranularity	string	`json:"metrics_granularity"`
	ForceDelete	bool	`json:"force_delete"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingGroupList is a list of AwsAutoscalingGroup resources
type AwsAutoscalingGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingGroup	`json:"items"`
}


// Tag is a Tag
type Tag struct {
	Key	string	`json:"key"`
	Value	string	`json:"value"`
	PropagateAtLaunch	bool	`json:"propagate_at_launch"`
}

// InitialLifecycleHook is a InitialLifecycleHook
type InitialLifecycleHook struct {
	NotificationTargetArn	string	`json:"notification_target_arn"`
	RoleArn	string	`json:"role_arn"`
	Name	string	`json:"name"`
	HeartbeatTimeout	int	`json:"heartbeat_timeout"`
	LifecycleTransition	string	`json:"lifecycle_transition"`
	NotificationMetadata	string	`json:"notification_metadata"`
}
