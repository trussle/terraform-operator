
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
	SetIdentifier	string	`json:"set_identifier"`
	LatencyRoutingPolicy	[]Generic	`json:"latency_routing_policy"`
	HealthCheckId	string	`json:"health_check_id"`
	AllowOverwrite	bool	`json:"allow_overwrite"`
	Name	string	`json:"name"`
	FailoverRoutingPolicy	[]Generic	`json:"failover_routing_policy"`
	GeolocationRoutingPolicy	[]Generic	`json:"geolocation_routing_policy"`
	Failover	string	`json:"failover"`
	Type	string	`json:"type"`
	Ttl	int	`json:"ttl"`
	Weight	int	`json:"weight"`
	MultivalueAnswerRoutingPolicy	bool	`json:"multivalue_answer_routing_policy"`
	Records	Generic	`json:"records"`
	ZoneId	string	`json:"zone_id"`
	Alias	Generic	`json:"alias"`
	WeightedRoutingPolicy	[]Generic	`json:"weighted_routing_policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53RecordList is a list of AwsRoute53Record resources
type AwsRoute53RecordList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53Record	`json:"items"`
}

