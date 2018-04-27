
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
	EstimatedInstanceWarmup	int	`json:"estimated_instance_warmup"`
	ScalingAdjustment	int	`json:"scaling_adjustment"`
	StepAdjustment	AwsAutoscalingPolicyStepAdjustment	`json:"step_adjustment"`
	TargetTrackingConfiguration	[]AwsAutoscalingPolicyTargetTrackingConfiguration	`json:"target_tracking_configuration"`
	AdjustmentType	string	`json:"adjustment_type"`
	PolicyType	string	`json:"policy_type"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	MinAdjustmentStep	int	`json:"min_adjustment_step"`
	Name	string	`json:"name"`
	AutoscalingGroupName	string	`json:"autoscaling_group_name"`
	Cooldown	int	`json:"cooldown"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingPolicyList is a list of AwsAutoscalingPolicy resources
type AwsAutoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingPolicy	`json:"items"`
}


// AwsAutoscalingPolicyStepAdjustment is a AwsAutoscalingPolicyStepAdjustment
type AwsAutoscalingPolicyStepAdjustment struct {
	MetricIntervalLowerBound	string	`json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound	string	`json:"metric_interval_upper_bound"`
	ScalingAdjustment	int	`json:"scaling_adjustment"`
}

// AwsAutoscalingPolicyPredefinedMetricSpecification is a AwsAutoscalingPolicyPredefinedMetricSpecification
type AwsAutoscalingPolicyPredefinedMetricSpecification struct {
	PredefinedMetricType	string	`json:"predefined_metric_type"`
	ResourceLabel	string	`json:"resource_label"`
}

// AwsAutoscalingPolicyMetricDimension is a AwsAutoscalingPolicyMetricDimension
type AwsAutoscalingPolicyMetricDimension struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}

// AwsAutoscalingPolicyCustomizedMetricSpecification is a AwsAutoscalingPolicyCustomizedMetricSpecification
type AwsAutoscalingPolicyCustomizedMetricSpecification struct {
	MetricDimension	[]AwsAutoscalingPolicyMetricDimension	`json:"metric_dimension"`
	MetricName	string	`json:"metric_name"`
	Namespace	string	`json:"namespace"`
	Statistic	string	`json:"statistic"`
	Unit	string	`json:"unit"`
}

// AwsAutoscalingPolicyTargetTrackingConfiguration is a AwsAutoscalingPolicyTargetTrackingConfiguration
type AwsAutoscalingPolicyTargetTrackingConfiguration struct {
	PredefinedMetricSpecification	[]AwsAutoscalingPolicyPredefinedMetricSpecification	`json:"predefined_metric_specification"`
	CustomizedMetricSpecification	[]AwsAutoscalingPolicyCustomizedMetricSpecification	`json:"customized_metric_specification"`
	TargetValue	float64	`json:"target_value"`
	DisableScaleIn	bool	`json:"disable_scale_in"`
}
