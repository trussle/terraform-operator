
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
	MinAdjustmentMagnitude	int	`json:"min_adjustment_magnitude"`
	AdjustmentType	string	`json:"adjustment_type"`
	AutoscalingGroupName	string	`json:"autoscaling_group_name"`
	Cooldown	int	`json:"cooldown"`
	MinAdjustmentStep	int	`json:"min_adjustment_step"`
	ScalingAdjustment	int	`json:"scaling_adjustment"`
	StepAdjustment	string	`json:"step_adjustment"`
	TargetTrackingConfiguration	[]LQrCNPAN	`json:"target_tracking_configuration"`
	Name	string	`json:"name"`
	PolicyType	string	`json:"policy_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingPolicyList is a list of AwsAutoscalingPolicy resources
type AwsAutoscalingPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingPolicy	`json:"items"`
}


// WvkkOdej is a WvkkOdej
type WvkkOdej struct {
	PredefinedMetricType	string	`json:"predefined_metric_type"`
	ResourceLabel	string	`json:"resource_label"`
}

// MUqBcVFo is a MUqBcVFo
type MUqBcVFo struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}

// VcugMVTn is a VcugMVTn
type VcugMVTn struct {
	Unit	string	`json:"unit"`
	MetricDimension	[]MUqBcVFo	`json:"metric_dimension"`
	MetricName	string	`json:"metric_name"`
	Namespace	string	`json:"namespace"`
	Statistic	string	`json:"statistic"`
}

// LQrCNPAN is a LQrCNPAN
type LQrCNPAN struct {
	PredefinedMetricSpecification	[]WvkkOdej	`json:"predefined_metric_specification"`
	CustomizedMetricSpecification	[]VcugMVTn	`json:"customized_metric_specification"`
	TargetValue	float64	`json:"target_value"`
	DisableScaleIn	bool	`json:"disable_scale_in"`
}
