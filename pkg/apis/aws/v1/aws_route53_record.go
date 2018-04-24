
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53Record struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRoute53RecordSpec	`json:"spec"`
}

type AwsRoute53RecordList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53Record	`json:"items"`
}

type AwsRoute53RecordSpec struct {
	Records	interface{}	`json:"records"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
	Weight	int	`json:"weight"`
	Alias	interface{}	`json:"alias"`
	Failover	string	`json:"failover"`
	MultivalueAnswerRoutingPolicy	bool	`json:"multivalue_answer_routing_policy"`
	HealthCheckId	string	`json:"health_check_id"`
	SetIdentifier	string	`json:"set_identifier"`
	WeightedRoutingPolicy	[]interface{}	`json:"weighted_routing_policy"`
	FailoverRoutingPolicy	[]interface{}	`json:"failover_routing_policy"`
	LatencyRoutingPolicy	[]interface{}	`json:"latency_routing_policy"`
	GeolocationRoutingPolicy	[]interface{}	`json:"geolocation_routing_policy"`
	AllowOverwrite	bool	`json:"allow_overwrite"`
	ZoneId	string	`json:"zone_id"`
	Ttl	int	`json:"ttl"`
}
