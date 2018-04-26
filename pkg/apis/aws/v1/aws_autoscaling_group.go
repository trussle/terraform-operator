
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
	HealthCheckGracePeriod	int	`json:"health_check_grace_period"`
	SuspendedProcesses	string	`json:"suspended_processes"`
	MetricsGranularity	string	`json:"metrics_granularity"`
	WaitForElbCapacity	int	`json:"wait_for_elb_capacity"`
	Tag	string	`json:"tag"`
	MinSize	int	`json:"min_size"`
	MinElbCapacity	int	`json:"min_elb_capacity"`
	TerminationPolicies	[]string	`json:"termination_policies"`
	Tags	[]map[string]???	`json:"tags"`
	NamePrefix	string	`json:"name_prefix"`
	ForceDelete	bool	`json:"force_delete"`
	PlacementGroup	string	`json:"placement_group"`
	EnabledMetrics	string	`json:"enabled_metrics"`
	LaunchConfiguration	string	`json:"launch_configuration"`
	MaxSize	int	`json:"max_size"`
	WaitForCapacityTimeout	string	`json:"wait_for_capacity_timeout"`
	ProtectFromScaleIn	bool	`json:"protect_from_scale_in"`
	InitialLifecycleHook	string	`json:"initial_lifecycle_hook"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingGroupList is a list of AwsAutoscalingGroup resources
type AwsAutoscalingGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingGroup	`json:"items"`
}

