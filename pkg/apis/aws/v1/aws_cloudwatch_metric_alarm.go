
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchMetricAlarm describes a AwsCloudwatchMetricAlarm resource
type AwsCloudwatchMetricAlarm struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchMetricAlarmSpec	`json:"spec"`
}


// AwsCloudwatchMetricAlarmSpec is the spec for a AwsCloudwatchMetricAlarm Resource
type AwsCloudwatchMetricAlarmSpec struct {
	AlarmName	string	`json:"alarm_name"`
	EvaluationPeriods	int	`json:"evaluation_periods"`
	Threshold	float64	`json:"threshold"`
	AlarmActions	string	`json:"alarm_actions"`
	InsufficientDataActions	string	`json:"insufficient_data_actions"`
	Namespace	string	`json:"namespace"`
	AlarmDescription	string	`json:"alarm_description"`
	Dimensions	map[string]string	`json:"dimensions"`
	OkActions	string	`json:"ok_actions"`
	ExtendedStatistic	string	`json:"extended_statistic"`
	ActionsEnabled	bool	`json:"actions_enabled"`
	DatapointsToAlarm	int	`json:"datapoints_to_alarm"`
	Unit	string	`json:"unit"`
	TreatMissingData	string	`json:"treat_missing_data"`
	ComparisonOperator	string	`json:"comparison_operator"`
	MetricName	string	`json:"metric_name"`
	Period	int	`json:"period"`
	Statistic	string	`json:"statistic"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchMetricAlarmList is a list of AwsCloudwatchMetricAlarm resources
type AwsCloudwatchMetricAlarmList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchMetricAlarm	`json:"items"`
}

