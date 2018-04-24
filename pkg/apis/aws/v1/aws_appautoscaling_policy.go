
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAppautoscalingPolicySpec	`json:"spec"`
}

type AwsAppautoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppautoscalingPolicy	`json:"items"`
}

type AwsAppautoscalingPolicySpec struct {
	Name	string	`json:"name"`
	ServiceNamespace	string	`json:"service_namespace"`
	Cooldown	int	`json:"cooldown"`
	PolicyType	string	`json:"policy_type"`
	ResourceId	string	`json:"resource_id"`
	Alarms	[]interface{}	`json:"alarms"`
	StepAdjustment	interface{}	`json:"step_adjustment"`
	TargetTrackingScalingPolicyConfiguration	[]interface{}	`json:"target_tracking_scaling_policy_configuration"`
	MetricAggregationType	string	`json:"metric_aggregation_type"`
	ScalableDimension	string	`json:"scalable_dimension"`
	StepScalingPolicyConfiguration	[]interface{}	`json:"step_scaling_policy_configuration"`
	AdjustmentType	string	`json:"adjustment_type"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
}
