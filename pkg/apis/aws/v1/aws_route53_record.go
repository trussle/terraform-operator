
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
	SetIdentifier	string	`json:"set_identifier"`
	Records	string	`json:"records"`
	Type	string	`json:"type"`
	ZoneId	string	`json:"zone_id"`
	Alias	Alias	`json:"alias"`
	MultivalueAnswerRoutingPolicy	bool	`json:"multivalue_answer_routing_policy"`
	AllowOverwrite	bool	`json:"allow_overwrite"`
	GeolocationRoutingPolicy	[]GeolocationRoutingPolicy	`json:"geolocation_routing_policy"`
	WeightedRoutingPolicy	[]WeightedRoutingPolicy	`json:"weighted_routing_policy"`
	Weight	int	`json:"weight"`
	Failover	string	`json:"failover"`
	FailoverRoutingPolicy	[]FailoverRoutingPolicy	`json:"failover_routing_policy"`
	LatencyRoutingPolicy	[]LatencyRoutingPolicy	`json:"latency_routing_policy"`
	Name	string	`json:"name"`
	Ttl	int	`json:"ttl"`
	HealthCheckId	string	`json:"health_check_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53RecordList is a list of AwsRoute53Record resources
type AwsRoute53RecordList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53Record	`json:"items"`
}


// GeolocationRoutingPolicy is a GeolocationRoutingPolicy
type GeolocationRoutingPolicy struct {
	Continent	string	`json:"continent"`
	Country	string	`json:"country"`
	Subdivision	string	`json:"subdivision"`
}

// WeightedRoutingPolicy is a WeightedRoutingPolicy
type WeightedRoutingPolicy struct {
	Weight	int	`json:"weight"`
}

// FailoverRoutingPolicy is a FailoverRoutingPolicy
type FailoverRoutingPolicy struct {
	Type	string	`json:"type"`
}

// LatencyRoutingPolicy is a LatencyRoutingPolicy
type LatencyRoutingPolicy struct {
	Region	string	`json:"region"`
}

// Alias is a Alias
type Alias struct {
	ZoneId	string	`json:"zone_id"`
	Name	string	`json:"name"`
	EvaluateTargetHealth	bool	`json:"evaluate_target_health"`
}
