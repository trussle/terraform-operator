
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
	AdjustmentType	string	`json:"adjustment_type"`
	Cooldown	int	`json:"cooldown"`
	ScalableDimension	string	`json:"scalable_dimension"`
	MetricAggregationType	string	`json:"metric_aggregation_type"`
	StepAdjustment	StepAdjustment	`json:"step_adjustment"`
	Name	string	`json:"name"`
	ServiceNamespace	string	`json:"service_namespace"`
	Alarms	[]string	`json:"alarms"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	TargetTrackingScalingPolicyConfiguration	[]TargetTrackingScalingPolicyConfiguration	`json:"target_tracking_scaling_policy_configuration"`
	PolicyType	string	`json:"policy_type"`
	ResourceId	string	`json:"resource_id"`
	StepScalingPolicyConfiguration	[]StepScalingPolicyConfiguration	`json:"step_scaling_policy_configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingPolicyList is a list of AwsAppautoscalingPolicy resources
type AwsAppautoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppautoscalingPolicy	`json:"items"`
}


// StepAdjustment is a StepAdjustment
type StepAdjustment struct {
	MetricIntervalLowerBound	float64	`json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound	float64	`json:"metric_interval_upper_bound"`
	ScalingAdjustment	int	`json:"scaling_adjustment"`
}

// Dimensions is a Dimensions
type Dimensions struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}

// CustomizedMetricSpecification is a CustomizedMetricSpecification
type CustomizedMetricSpecification struct {
	Namespace	string	`json:"namespace"`
	Statistic	string	`json:"statistic"`
	Unit	string	`json:"unit"`
	Dimensions	Dimensions	`json:"dimensions"`
	MetricName	string	`json:"metric_name"`
}

// PredefinedMetricSpecification is a PredefinedMetricSpecification
type PredefinedMetricSpecification struct {
	PredefinedMetricType	string	`json:"predefined_metric_type"`
	ResourceLabel	string	`json:"resource_label"`
}

// TargetTrackingScalingPolicyConfiguration is a TargetTrackingScalingPolicyConfiguration
type TargetTrackingScalingPolicyConfiguration struct {
	CustomizedMetricSpecification	[]CustomizedMetricSpecification	`json:"customized_metric_specification"`
	PredefinedMetricSpecification	[]PredefinedMetricSpecification	`json:"predefined_metric_specification"`
	DisableScaleIn	bool	`json:"disable_scale_in"`
	ScaleInCooldown	int	`json:"scale_in_cooldown"`
	ScaleOutCooldown	int	`json:"scale_out_cooldown"`
	TargetValue	float64	`json:"target_value"`
}

// StepScalingPolicyConfiguration is a StepScalingPolicyConfiguration
type StepScalingPolicyConfiguration struct {
	AdjustmentType	string	`json:"adjustment_type"`
	Cooldown	int	`json:"cooldown"`
	MetricAggregationType	string	`json:"metric_aggregation_type"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	StepAdjustment	StepAdjustment	`json:"step_adjustment"`
}
