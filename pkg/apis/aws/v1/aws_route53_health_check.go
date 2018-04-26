
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
	Regions	string	`json:"regions"`
	Tags	map[string]???	`json:"tags"`
	IpAddress	string	`json:"ip_address"`
	SearchString	string	`json:"search_string"`
	ChildHealthchecks	string	`json:"child_healthchecks"`
	FailureThreshold	int	`json:"failure_threshold"`
	Fqdn	string	`json:"fqdn"`
	ReferenceName	string	`json:"reference_name"`
	ResourcePath	string	`json:"resource_path"`
	MeasureLatency	bool	`json:"measure_latency"`
	CloudwatchAlarmRegion	string	`json:"cloudwatch_alarm_region"`
	InsufficientDataHealthStatus	string	`json:"insufficient_data_health_status"`
	Type	string	`json:"type"`
	Port	int	`json:"port"`
	InvertHealthcheck	bool	`json:"invert_healthcheck"`
	RequestInterval	int	`json:"request_interval"`
	ChildHealthThreshold	int	`json:"child_health_threshold"`
	CloudwatchAlarmName	string	`json:"cloudwatch_alarm_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53HealthCheckList is a list of AwsRoute53HealthCheck resources
type AwsRoute53HealthCheckList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53HealthCheck	`json:"items"`
}

