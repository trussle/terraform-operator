
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53HealthCheck describes a AwsRoute53HealthCheck resource
type AwsRoute53HealthCheck struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRoute53HealthCheckSpec	`json:"spec"`
}


// AwsRoute53HealthCheckSpec is the spec for a AwsRoute53HealthCheck Resource
type AwsRoute53HealthCheckSpec struct {
	RequestInterval	int	`json:"request_interval"`
	Fqdn	string	`json:"fqdn"`
	ResourcePath	string	`json:"resource_path"`
	SearchString	string	`json:"search_string"`
	ChildHealthThreshold	int	`json:"child_health_threshold"`
	CloudwatchAlarmRegion	string	`json:"cloudwatch_alarm_region"`
	Type	string	`json:"type"`
	InsufficientDataHealthStatus	string	`json:"insufficient_data_health_status"`
	ReferenceName	string	`json:"reference_name"`
	FailureThreshold	int	`json:"failure_threshold"`
	IpAddress	string	`json:"ip_address"`
	InvertHealthcheck	bool	`json:"invert_healthcheck"`
	MeasureLatency	bool	`json:"measure_latency"`
	Port	int	`json:"port"`
	ChildHealthchecks	string	`json:"child_healthchecks"`
	CloudwatchAlarmName	string	`json:"cloudwatch_alarm_name"`
	Regions	string	`json:"regions"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53HealthCheckList is a list of AwsRoute53HealthCheck resources
type AwsRoute53HealthCheckList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53HealthCheck	`json:"items"`
}

