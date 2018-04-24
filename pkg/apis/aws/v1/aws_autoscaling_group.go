
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAutoscalingGroupSpec	`json:"spec"`
}

type AwsAutoscalingGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingGroup	`json:"items"`
}

type AwsAutoscalingGroupSpec struct {
	ForceDelete	bool	`json:"force_delete"`
	TerminationPolicies	[]interface{}	`json:"termination_policies"`
	WaitForElbCapacity	int	`json:"wait_for_elb_capacity"`
	NamePrefix	string	`json:"name_prefix"`
	Tags	[]interface{}	`json:"tags"`
	LaunchConfiguration	string	`json:"launch_configuration"`
	MinSize	int	`json:"min_size"`
	SuspendedProcesses	interface{}	`json:"suspended_processes"`
	MetricsGranularity	string	`json:"metrics_granularity"`
	MinElbCapacity	int	`json:"min_elb_capacity"`
	PlacementGroup	string	`json:"placement_group"`
	WaitForCapacityTimeout	string	`json:"wait_for_capacity_timeout"`
	Tag	interface{}	`json:"tag"`
	MaxSize	int	`json:"max_size"`
	ProtectFromScaleIn	bool	`json:"protect_from_scale_in"`
	HealthCheckGracePeriod	int	`json:"health_check_grace_period"`
	EnabledMetrics	interface{}	`json:"enabled_metrics"`
	InitialLifecycleHook	interface{}	`json:"initial_lifecycle_hook"`
}
