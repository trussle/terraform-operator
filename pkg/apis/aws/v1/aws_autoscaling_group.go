
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
	WaitForElbCapacity	int	`json:"wait_for_elb_capacity"`
	MinSize	int	`json:"min_size"`
	Tag	Generic	`json:"tag"`
	MinElbCapacity	int	`json:"min_elb_capacity"`
	ForceDelete	bool	`json:"force_delete"`
	ProtectFromScaleIn	bool	`json:"protect_from_scale_in"`
	Tags	[]Generic	`json:"tags"`
	TerminationPolicies	[]Generic	`json:"termination_policies"`
	InitialLifecycleHook	Generic	`json:"initial_lifecycle_hook"`
	NamePrefix	string	`json:"name_prefix"`
	MaxSize	int	`json:"max_size"`
	HealthCheckGracePeriod	int	`json:"health_check_grace_period"`
	SuspendedProcesses	Generic	`json:"suspended_processes"`
	PlacementGroup	string	`json:"placement_group"`
	MetricsGranularity	string	`json:"metrics_granularity"`
	LaunchConfiguration	string	`json:"launch_configuration"`
	WaitForCapacityTimeout	string	`json:"wait_for_capacity_timeout"`
	EnabledMetrics	Generic	`json:"enabled_metrics"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingGroupList is a list of AwsAutoscalingGroup resources
type AwsAutoscalingGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingGroup	`json:"items"`
}

