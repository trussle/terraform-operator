
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
	PolicyType	string	`json:"policy_type"`
	ScalableDimension	string	`json:"scalable_dimension"`
	StepScalingPolicyConfiguration	[]mYYtEjVg	`json:"step_scaling_policy_configuration"`
	AdjustmentType	string	`json:"adjustment_type"`
	Cooldown	int	`json:"cooldown"`
	TargetTrackingScalingPolicyConfiguration	[]wfFbbGGc	`json:"target_tracking_scaling_policy_configuration"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	Name	string	`json:"name"`
	ServiceNamespace	string	`json:"service_namespace"`
	MetricAggregationType	string	`json:"metric_aggregation_type"`
	ResourceId	string	`json:"resource_id"`
	Alarms	[]string	`json:"alarms"`
	StepAdjustment	string	`json:"step_adjustment"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingPolicyList is a list of AwsAppautoscalingPolicy resources
type AwsAppautoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppautoscalingPolicy	`json:"items"`
}


// mYYtEjVg is a mYYtEjVg
type mYYtEjVg struct {
	AdjustmentType	string	`json:"adjustment_type"`
	Cooldown	int	`json:"cooldown"`
	MetricAggregationType	string	`json:"metric_aggregation_type"`
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	StepAdjustment	string	`json:"step_adjustment"`
}

// nqbaEREu is a nqbaEREu
type nqbaEREu struct {
	Dimensions	string	`json:"dimensions"`
	MetricName	string	`json:"metric_name"`
	Namespace	string	`json:"namespace"`
	Statistic	string	`json:"statistic"`
	Unit	string	`json:"unit"`
}

// nUZjQXmZ is a nUZjQXmZ
type nUZjQXmZ struct {
	PredefinedMetricType	string	`json:"predefined_metric_type"`
	ResourceLabel	string	`json:"resource_label"`
}

// wfFbbGGc is a wfFbbGGc
type wfFbbGGc struct {
	CustomizedMetricSpecification	[]nqbaEREu	`json:"customized_metric_specification"`
	PredefinedMetricSpecification	[]nUZjQXmZ	`json:"predefined_metric_specification"`
	DisableScaleIn	bool	`json:"disable_scale_in"`
	ScaleInCooldown	int	`json:"scale_in_cooldown"`
	ScaleOutCooldown	int	`json:"scale_out_cooldown"`
	TargetValue	float64	`json:"target_value"`
}
