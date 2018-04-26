
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
	Unit	string	`json:"unit"`
	Threshold	float64	`json:"threshold"`
	ActionsEnabled	bool	`json:"actions_enabled"`
	AlarmDescription	string	`json:"alarm_description"`
	DatapointsToAlarm	int	`json:"datapoints_to_alarm"`
	Dimensions	map[string]???	`json:"dimensions"`
	InsufficientDataActions	string	`json:"insufficient_data_actions"`
	Period	int	`json:"period"`
	EvaluationPeriods	int	`json:"evaluation_periods"`
	Statistic	string	`json:"statistic"`
	AlarmActions	string	`json:"alarm_actions"`
	TreatMissingData	string	`json:"treat_missing_data"`
	ComparisonOperator	string	`json:"comparison_operator"`
	MetricName	string	`json:"metric_name"`
	Namespace	string	`json:"namespace"`
	OkActions	string	`json:"ok_actions"`
	ExtendedStatistic	string	`json:"extended_statistic"`
	AlarmName	string	`json:"alarm_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchMetricAlarmList is a list of AwsCloudwatchMetricAlarm resources
type AwsCloudwatchMetricAlarmList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchMetricAlarm	`json:"items"`
}

