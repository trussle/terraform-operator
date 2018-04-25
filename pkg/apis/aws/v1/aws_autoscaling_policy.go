
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
	EstimatedInstanceWarmup	int	`json:"estimated_instance_warmup"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	MinAdjustmentStep	int	`json:"min_adjustment_step"`
	ScalingAdjustment	int	`json:"scaling_adjustment"`
	AdjustmentType	string	`json:"adjustment_type"`
	Name	string	`json:"name"`
	AutoscalingGroupName	string	`json:"autoscaling_group_name"`
	PolicyType	string	`json:"policy_type"`
	Cooldown	int	`json:"cooldown"`
	StepAdjustment	string	`json:"step_adjustment"`
	TargetTrackingConfiguration	[]interface{}	`json:"target_tracking_configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingPolicyList is a list of AwsAutoscalingPolicy resources
type AwsAutoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingPolicy	`json:"items"`
}

