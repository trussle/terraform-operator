
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53HealthCheck struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRoute53HealthCheckSpec	`json:"spec"`
}

type AwsRoute53HealthCheckList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53HealthCheck	`json:"items"`
}

type AwsRoute53HealthCheckSpec struct {
	Port	int	`json:"port"`
	InvertHealthcheck	bool	`json:"invert_healthcheck"`
	SearchString	string	`json:"search_string"`
	FailureThreshold	int	`json:"failure_threshold"`
	ResourcePath	string	`json:"resource_path"`
	CloudwatchAlarmName	string	`json:"cloudwatch_alarm_name"`
	CloudwatchAlarmRegion	string	`json:"cloudwatch_alarm_region"`
	Regions	interface{}	`json:"regions"`
	RequestInterval	int	`json:"request_interval"`
	IpAddress	string	`json:"ip_address"`
	Fqdn	string	`json:"fqdn"`
	MeasureLatency	bool	`json:"measure_latency"`
	InsufficientDataHealthStatus	string	`json:"insufficient_data_health_status"`
	Tags	map[string]interface{}	`json:"tags"`
	Type	string	`json:"type"`
	ChildHealthThreshold	int	`json:"child_health_threshold"`
	ReferenceName	string	`json:"reference_name"`
	ChildHealthchecks	interface{}	`json:"child_healthchecks"`
}
