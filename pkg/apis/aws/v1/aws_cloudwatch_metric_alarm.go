
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchMetricAlarm struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchMetricAlarmSpec	`json:"spec"`
}

type AwsCloudwatchMetricAlarmList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchMetricAlarm	`json:"items"`
}

type AwsCloudwatchMetricAlarmSpec struct {
	Period	int	`json:"period"`
	Statistic	string	`json:"statistic"`
	AlarmActions	interface{}	`json:"alarm_actions"`
	DatapointsToAlarm	int	`json:"datapoints_to_alarm"`
	Dimensions	map[string]interface{}	`json:"dimensions"`
	TreatMissingData	string	`json:"treat_missing_data"`
	InsufficientDataActions	interface{}	`json:"insufficient_data_actions"`
	Unit	string	`json:"unit"`
	ExtendedStatistic	string	`json:"extended_statistic"`
	AlarmName	string	`json:"alarm_name"`
	EvaluationPeriods	int	`json:"evaluation_periods"`
	Namespace	string	`json:"namespace"`
	AlarmDescription	string	`json:"alarm_description"`
	OkActions	interface{}	`json:"ok_actions"`
	ComparisonOperator	string	`json:"comparison_operator"`
	MetricName	string	`json:"metric_name"`
	Threshold	float	`json:"threshold"`
	ActionsEnabled	bool	`json:"actions_enabled"`
}
