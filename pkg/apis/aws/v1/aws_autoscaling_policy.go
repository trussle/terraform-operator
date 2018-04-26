
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	AdjustmentType	string	`json:"adjustment_type"`
	PolicyType	string	`json:"policy_type"`
	ScalingAdjustment	int	`json:"scaling_adjustment"`
	TargetTrackingConfiguration	[]TargetTrackingConfiguration	`json:"target_tracking_configuration"`
	MinAdjustmentStep	int	`json:"min_adjustment_step"`
	StepAdjustment	StepAdjustment	`json:"step_adjustment"`
	Name	string	`json:"name"`
	AutoscalingGroupName	string	`json:"autoscaling_group_name"`
	Cooldown	int	`json:"cooldown"`
	EstimatedInstanceWarmup	int	`json:"estimated_instance_warmup"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingPolicyList is a list of AwsAutoscalingPolicy resources
type AwsAutoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingPolicy	`json:"items"`
}


// StepAdjustment is a StepAdjustment
type StepAdjustment struct {
	MetricIntervalLowerBound	string	`json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound	string	`json:"metric_interval_upper_bound"`
	ScalingAdjustment	int	`json:"scaling_adjustment"`
}

// PredefinedMetricSpecification is a PredefinedMetricSpecification
type PredefinedMetricSpecification struct {
	ResourceLabel	string	`json:"resource_label"`
	PredefinedMetricType	string	`json:"predefined_metric_type"`
}

// MetricDimension is a MetricDimension
type MetricDimension struct {
	Value	string	`json:"value"`
	Name	string	`json:"name"`
}

// CustomizedMetricSpecification is a CustomizedMetricSpecification
type CustomizedMetricSpecification struct {
	MetricDimension	[]MetricDimension	`json:"metric_dimension"`
	MetricName	string	`json:"metric_name"`
	Namespace	string	`json:"namespace"`
	Statistic	string	`json:"statistic"`
	Unit	string	`json:"unit"`
}

// TargetTrackingConfiguration is a TargetTrackingConfiguration
type TargetTrackingConfiguration struct {
	PredefinedMetricSpecification	[]PredefinedMetricSpecification	`json:"predefined_metric_specification"`
	CustomizedMetricSpecification	[]CustomizedMetricSpecification	`json:"customized_metric_specification"`
	TargetValue	float64	`json:"target_value"`
	DisableScaleIn	bool	`json:"disable_scale_in"`
}
