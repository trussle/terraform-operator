
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
	ComparisonOperator	string	`json:"comparison_operator"`
	InsufficientDataActions	string	`json:"insufficient_data_actions"`
	Unit	string	`json:"unit"`
	AlarmName	string	`json:"alarm_name"`
	Statistic	string	`json:"statistic"`
	Threshold	float64	`json:"threshold"`
	ActionsEnabled	bool	`json:"actions_enabled"`
	AlarmDescription	string	`json:"alarm_description"`
	DatapointsToAlarm	int	`json:"datapoints_to_alarm"`
	AlarmActions	string	`json:"alarm_actions"`
	EvaluationPeriods	int	`json:"evaluation_periods"`
	MetricName	string	`json:"metric_name"`
	Namespace	string	`json:"namespace"`
	Period	int	`json:"period"`
	Dimensions	map[string]string	`json:"dimensions"`
	OkActions	string	`json:"ok_actions"`
	ExtendedStatistic	string	`json:"extended_statistic"`
	TreatMissingData	string	`json:"treat_missing_data"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchMetricAlarmList is a list of AwsCloudwatchMetricAlarm resources
type AwsCloudwatchMetricAlarmList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchMetricAlarm	`json:"items"`
}

