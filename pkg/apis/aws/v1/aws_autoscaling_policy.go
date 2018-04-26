
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingPolicy describes a AwsAutoscalingPolicy resource
type AwsAutoscalingPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAutoscalingPolicySpec	`json:"spec"`
}


// AwsAutoscalingPolicySpec is the spec for a AwsAutoscalingPolicy Resource
type AwsAutoscalingPolicySpec struct {
	MinAdjustmentStep	int	`json:"min_adjustment_step"`
	StepAdjustment	Generic	`json:"step_adjustment"`
	TargetTrackingConfiguration	[]Generic	`json:"target_tracking_configuration"`
	AutoscalingGroupName	string	`json:"autoscaling_group_name"`
	Cooldown	int	`json:"cooldown"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	PolicyType	string	`json:"policy_type"`
	EstimatedInstanceWarmup	int	`json:"estimated_instance_warmup"`
	ScalingAdjustment	int	`json:"scaling_adjustment"`
	Name	string	`json:"name"`
	AdjustmentType	string	`json:"adjustment_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingPolicyList is a list of AwsAutoscalingPolicy resources
type AwsAutoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingPolicy	`json:"items"`
}

