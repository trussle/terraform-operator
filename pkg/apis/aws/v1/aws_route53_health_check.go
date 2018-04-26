
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
	Type	string	`json:"type"`
	Fqdn	string	`json:"fqdn"`
	Port	int	`json:"port"`
	SearchString	string	`json:"search_string"`
	Tags	map[string]Generic	`json:"tags"`
	InsufficientDataHealthStatus	string	`json:"insufficient_data_health_status"`
	ReferenceName	string	`json:"reference_name"`
	RequestInterval	int	`json:"request_interval"`
	InvertHealthcheck	bool	`json:"invert_healthcheck"`
	ResourcePath	string	`json:"resource_path"`
	MeasureLatency	bool	`json:"measure_latency"`
	ChildHealthchecks	Generic	`json:"child_healthchecks"`
	ChildHealthThreshold	int	`json:"child_health_threshold"`
	Regions	Generic	`json:"regions"`
	FailureThreshold	int	`json:"failure_threshold"`
	CloudwatchAlarmRegion	string	`json:"cloudwatch_alarm_region"`
	IpAddress	string	`json:"ip_address"`
	CloudwatchAlarmName	string	`json:"cloudwatch_alarm_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53HealthCheckList is a list of AwsRoute53HealthCheck resources
type AwsRoute53HealthCheckList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53HealthCheck	`json:"items"`
}

