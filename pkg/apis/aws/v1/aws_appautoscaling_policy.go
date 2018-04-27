
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	MetricAggregationType	string	`json:"metric_aggregation_type"`
	StepAdjustment	AwsAppautoscalingPolicyStepAdjustment	`json:"step_adjustment"`
	TargetTrackingScalingPolicyConfiguration	[]AwsAppautoscalingPolicyTargetTrackingScalingPolicyConfiguration	`json:"target_tracking_scaling_policy_configuration"`
	Name	string	`json:"name"`
	ScalableDimension	string	`json:"scalable_dimension"`
	Alarms	[]string	`json:"alarms"`
	ResourceId	string	`json:"resource_id"`
	ServiceNamespace	string	`json:"service_namespace"`
	StepScalingPolicyConfiguration	[]AwsAppautoscalingPolicyStepScalingPolicyConfiguration	`json:"step_scaling_policy_configuration"`
	AdjustmentType	string	`json:"adjustment_type"`
	Cooldown	int	`json:"cooldown"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	PolicyType	string	`json:"policy_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingPolicyList is a list of AwsAppautoscalingPolicy resources
type AwsAppautoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppautoscalingPolicy	`json:"items"`
}


// AwsAppautoscalingPolicyPredefinedMetricSpecification is a AwsAppautoscalingPolicyPredefinedMetricSpecification
type AwsAppautoscalingPolicyPredefinedMetricSpecification struct {
	PredefinedMetricType	string	`json:"predefined_metric_type"`
	ResourceLabel	string	`json:"resource_label"`
}

// AwsAppautoscalingPolicyTargetTrackingScalingPolicyConfiguration is a AwsAppautoscalingPolicyTargetTrackingScalingPolicyConfiguration
type AwsAppautoscalingPolicyTargetTrackingScalingPolicyConfiguration struct {
	ScaleOutCooldown	int	`json:"scale_out_cooldown"`
	TargetValue	float64	`json:"target_value"`
	CustomizedMetricSpecification	[]AwsAppautoscalingPolicyCustomizedMetricSpecification	`json:"customized_metric_specification"`
	PredefinedMetricSpecification	[]AwsAppautoscalingPolicyPredefinedMetricSpecification	`json:"predefined_metric_specification"`
	DisableScaleIn	bool	`json:"disable_scale_in"`
	ScaleInCooldown	int	`json:"scale_in_cooldown"`
}

// AwsAppautoscalingPolicyStepScalingPolicyConfiguration is a AwsAppautoscalingPolicyStepScalingPolicyConfiguration
type AwsAppautoscalingPolicyStepScalingPolicyConfiguration struct {
	AdjustmentType	string	`json:"adjustment_type"`
	Cooldown	int	`json:"cooldown"`
	MetricAggregationType	string	`json:"metric_aggregation_type"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	StepAdjustment	AwsAppautoscalingPolicyStepAdjustment	`json:"step_adjustment"`
}

// AwsAppautoscalingPolicyStepAdjustment is a AwsAppautoscalingPolicyStepAdjustment
type AwsAppautoscalingPolicyStepAdjustment struct {
	MetricIntervalUpperBound	string	`json:"metric_interval_upper_bound"`
	ScalingAdjustment	int	`json:"scaling_adjustment"`
	MetricIntervalLowerBound	string	`json:"metric_interval_lower_bound"`
}

// AwsAppautoscalingPolicyDimensions is a AwsAppautoscalingPolicyDimensions
type AwsAppautoscalingPolicyDimensions struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}

// AwsAppautoscalingPolicyCustomizedMetricSpecification is a AwsAppautoscalingPolicyCustomizedMetricSpecification
type AwsAppautoscalingPolicyCustomizedMetricSpecification struct {
	Dimensions	AwsAppautoscalingPolicyDimensions	`json:"dimensions"`
	MetricName	string	`json:"metric_name"`
	Namespace	string	`json:"namespace"`
	Statistic	string	`json:"statistic"`
	Unit	string	`json:"unit"`
}
