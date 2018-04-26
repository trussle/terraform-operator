
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	DatapointsToAlarm	int	`json:"datapoints_to_alarm"`
	AlarmName	string	`json:"alarm_name"`
	Namespace	string	`json:"namespace"`
	Statistic	string	`json:"statistic"`
	ActionsEnabled	bool	`json:"actions_enabled"`
	AlarmDescription	string	`json:"alarm_description"`
	Dimensions	map[string]Generic	`json:"dimensions"`
	InsufficientDataActions	Generic	`json:"insufficient_data_actions"`
	Unit	string	`json:"unit"`
	EvaluationPeriods	int	`json:"evaluation_periods"`
	ExtendedStatistic	string	`json:"extended_statistic"`
	MetricName	string	`json:"metric_name"`
	Period	int	`json:"period"`
	OkActions	Generic	`json:"ok_actions"`
	ComparisonOperator	string	`json:"comparison_operator"`
	AlarmActions	Generic	`json:"alarm_actions"`
	TreatMissingData	string	`json:"treat_missing_data"`
	Threshold	float	`json:"threshold"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchMetricAlarmList is a list of AwsCloudwatchMetricAlarm resources
type AwsCloudwatchMetricAlarmList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchMetricAlarm	`json:"items"`
}

