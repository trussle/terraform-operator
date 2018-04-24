
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAutoscalingPolicySpec	`json:"spec"`
}

type AwsAutoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingPolicy	`json:"items"`
}

type AwsAutoscalingPolicySpec struct {
	StepAdjustment	interface{}	`json:"step_adjustment"`
	TargetTrackingConfiguration	[]interface{}	`json:"target_tracking_configuration"`
	Name	string	`json:"name"`
	AdjustmentType	string	`json:"adjustment_type"`
	AutoscalingGroupName	string	`json:"autoscaling_group_name"`
	PolicyType	string	`json:"policy_type"`
	ScalingAdjustment	int	`json:"scaling_adjustment"`
	Cooldown	int	`json:"cooldown"`
	EstimatedInstanceWarmup	int	`json:"estimated_instance_warmup"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	MinAdjustmentStep	int	`json:"min_adjustment_step"`
}
