
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	SearchString	string	`json:"search_string"`
	MeasureLatency	bool	`json:"measure_latency"`
	CloudwatchAlarmName	string	`json:"cloudwatch_alarm_name"`
	FailureThreshold	int	`json:"failure_threshold"`
	RequestInterval	int	`json:"request_interval"`
	IpAddress	string	`json:"ip_address"`
	Fqdn	string	`json:"fqdn"`
	Port	int	`json:"port"`
	ReferenceName	string	`json:"reference_name"`
	ChildHealthThreshold	int	`json:"child_health_threshold"`
	InsufficientDataHealthStatus	string	`json:"insufficient_data_health_status"`
	Regions	string	`json:"regions"`
	ResourcePath	string	`json:"resource_path"`
	Tags	map[string]string	`json:"tags"`
	Type	string	`json:"type"`
	InvertHealthcheck	bool	`json:"invert_healthcheck"`
	ChildHealthchecks	string	`json:"child_healthchecks"`
	CloudwatchAlarmRegion	string	`json:"cloudwatch_alarm_region"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53HealthCheckList is a list of AwsRoute53HealthCheck resources
type AwsRoute53HealthCheckList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53HealthCheck	`json:"items"`
}

