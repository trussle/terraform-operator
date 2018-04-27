
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	MultivalueAnswerRoutingPolicy	bool	`json:"multivalue_answer_routing_policy"`
	AllowOverwrite	bool	`json:"allow_overwrite"`
	WeightedRoutingPolicy	[]AwsRoute53RecordWeightedRoutingPolicy	`json:"weighted_routing_policy"`
	LatencyRoutingPolicy	[]AwsRoute53RecordLatencyRoutingPolicy	`json:"latency_routing_policy"`
	HealthCheckId	string	`json:"health_check_id"`
	Name	string	`json:"name"`
	Failover	string	`json:"failover"`
	FailoverRoutingPolicy	[]AwsRoute53RecordFailoverRoutingPolicy	`json:"failover_routing_policy"`
	SetIdentifier	string	`json:"set_identifier"`
	Alias	AwsRoute53RecordAlias	`json:"alias"`
	GeolocationRoutingPolicy	[]AwsRoute53RecordGeolocationRoutingPolicy	`json:"geolocation_routing_policy"`
	Records	string	`json:"records"`
	Type	string	`json:"type"`
	ZoneId	string	`json:"zone_id"`
	Ttl	int	`json:"ttl"`
	Weight	int	`json:"weight"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53RecordList is a list of AwsRoute53Record resources
type AwsRoute53RecordList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53Record	`json:"items"`
}


// AwsRoute53RecordAlias is a AwsRoute53RecordAlias
type AwsRoute53RecordAlias struct {
	ZoneId	string	`json:"zone_id"`
	Name	string	`json:"name"`
	EvaluateTargetHealth	bool	`json:"evaluate_target_health"`
}

// AwsRoute53RecordGeolocationRoutingPolicy is a AwsRoute53RecordGeolocationRoutingPolicy
type AwsRoute53RecordGeolocationRoutingPolicy struct {
	Country	string	`json:"country"`
	Subdivision	string	`json:"subdivision"`
	Continent	string	`json:"continent"`
}

// AwsRoute53RecordWeightedRoutingPolicy is a AwsRoute53RecordWeightedRoutingPolicy
type AwsRoute53RecordWeightedRoutingPolicy struct {
	Weight	int	`json:"weight"`
}

// AwsRoute53RecordLatencyRoutingPolicy is a AwsRoute53RecordLatencyRoutingPolicy
type AwsRoute53RecordLatencyRoutingPolicy struct {
	Region	string	`json:"region"`
}

// AwsRoute53RecordFailoverRoutingPolicy is a AwsRoute53RecordFailoverRoutingPolicy
type AwsRoute53RecordFailoverRoutingPolicy struct {
	Type	string	`json:"type"`
}
