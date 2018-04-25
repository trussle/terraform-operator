
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
	EvaluationPeriods	int	`json:"evaluation_periods"`
	Threshold	float	`json:"threshold"`
	ActionsEnabled	bool	`json:"actions_enabled"`
	AlarmDescription	string	`json:"alarm_description"`
	Dimensions	map[string]interface{}	`json:"dimensions"`
	OkActions	string	`json:"ok_actions"`
	AlarmName	string	`json:"alarm_name"`
	ComparisonOperator	string	`json:"comparison_operator"`
	TreatMissingData	string	`json:"treat_missing_data"`
	ExtendedStatistic	string	`json:"extended_statistic"`
	Period	int	`json:"period"`
	Statistic	string	`json:"statistic"`
	DatapointsToAlarm	int	`json:"datapoints_to_alarm"`
	InsufficientDataActions	string	`json:"insufficient_data_actions"`
	MetricName	string	`json:"metric_name"`
	Namespace	string	`json:"namespace"`
	AlarmActions	string	`json:"alarm_actions"`
	Unit	string	`json:"unit"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchMetricAlarmList is a list of AwsCloudwatchMetricAlarm resources
type AwsCloudwatchMetricAlarmList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchMetricAlarm	`json:"items"`
}

