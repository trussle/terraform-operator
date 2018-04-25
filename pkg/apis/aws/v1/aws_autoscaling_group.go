
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	TerminationPolicies	[]interface{}	`json:"termination_policies"`
	WaitForElbCapacity	int	`json:"wait_for_elb_capacity"`
	LaunchConfiguration	string	`json:"launch_configuration"`
	ProtectFromScaleIn	bool	`json:"protect_from_scale_in"`
	InitialLifecycleHook	string	`json:"initial_lifecycle_hook"`
	PlacementGroup	string	`json:"placement_group"`
	EnabledMetrics	string	`json:"enabled_metrics"`
	Tag	string	`json:"tag"`
	HealthCheckGracePeriod	int	`json:"health_check_grace_period"`
	MetricsGranularity	string	`json:"metrics_granularity"`
	MinSize	int	`json:"min_size"`
	SuspendedProcesses	string	`json:"suspended_processes"`
	MaxSize	int	`json:"max_size"`
	MinElbCapacity	int	`json:"min_elb_capacity"`
	ForceDelete	bool	`json:"force_delete"`
	WaitForCapacityTimeout	string	`json:"wait_for_capacity_timeout"`
	Tags	[]interface{}	`json:"tags"`
	NamePrefix	string	`json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingGroupList is a list of AwsAutoscalingGroup resources
type AwsAutoscalingGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingGroup	`json:"items"`
}

