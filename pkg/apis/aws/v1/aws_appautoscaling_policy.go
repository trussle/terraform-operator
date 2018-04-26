
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingPolicy describes a AwsAppautoscalingPolicy resource
type AwsAppautoscalingPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAppautoscalingPolicySpec	`json:"spec"`
}


// AwsAppautoscalingPolicySpec is the spec for a AwsAppautoscalingPolicy Resource
type AwsAppautoscalingPolicySpec struct {
	PolicyType	string	`json:"policy_type"`
	ResourceId	string	`json:"resource_id"`
	ScalableDimension	string	`json:"scalable_dimension"`
	StepScalingPolicyConfiguration	[]Generic	`json:"step_scaling_policy_configuration"`
	Alarms	[]Generic	`json:"alarms"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	MetricAggregationType	string	`json:"metric_aggregation_type"`
	TargetTrackingScalingPolicyConfiguration	[]Generic	`json:"target_tracking_scaling_policy_configuration"`
	Name	string	`json:"name"`
	ServiceNamespace	string	`json:"service_namespace"`
	AdjustmentType	string	`json:"adjustment_type"`
	Cooldown	int	`json:"cooldown"`
	StepAdjustment	Generic	`json:"step_adjustment"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingPolicyList is a list of AwsAppautoscalingPolicy resources
type AwsAppautoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppautoscalingPolicy	`json:"items"`
}

