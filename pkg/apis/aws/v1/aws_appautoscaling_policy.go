
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
	ScalableDimension	string	`json:"scalable_dimension"`
	StepScalingPolicyConfiguration	[]interface{}	`json:"step_scaling_policy_configuration"`
	Cooldown	int	`json:"cooldown"`
	Name	string	`json:"name"`
	PolicyType	string	`json:"policy_type"`
	MetricAggregationType	string	`json:"metric_aggregation_type"`
	StepAdjustment	string	`json:"step_adjustment"`
	ServiceNamespace	string	`json:"service_namespace"`
	AdjustmentType	string	`json:"adjustment_type"`
	Alarms	[]interface{}	`json:"alarms"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	TargetTrackingScalingPolicyConfiguration	[]interface{}	`json:"target_tracking_scaling_policy_configuration"`
	ResourceId	string	`json:"resource_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingPolicyList is a list of AwsAppautoscalingPolicy resources
type AwsAppautoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppautoscalingPolicy	`json:"items"`
}

