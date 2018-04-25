
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53Record describes a AwsRoute53Record resource
type AwsRoute53Record struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRoute53RecordSpec	`json:"spec"`
}


// AwsRoute53RecordSpec is the spec for a AwsRoute53Record Resource
type AwsRoute53RecordSpec struct {
	AllowOverwrite	bool	`json:"allow_overwrite"`
	Alias	string	`json:"alias"`
	Failover	string	`json:"failover"`
	LatencyRoutingPolicy	[]interface{}	`json:"latency_routing_policy"`
	WeightedRoutingPolicy	[]interface{}	`json:"weighted_routing_policy"`
	HealthCheckId	string	`json:"health_check_id"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
	Weight	int	`json:"weight"`
	GeolocationRoutingPolicy	[]interface{}	`json:"geolocation_routing_policy"`
	Records	string	`json:"records"`
	SetIdentifier	string	`json:"set_identifier"`
	MultivalueAnswerRoutingPolicy	bool	`json:"multivalue_answer_routing_policy"`
	ZoneId	string	`json:"zone_id"`
	Ttl	int	`json:"ttl"`
	FailoverRoutingPolicy	[]interface{}	`json:"failover_routing_policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53RecordList is a list of AwsRoute53Record resources
type AwsRoute53RecordList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53Record	`json:"items"`
}

