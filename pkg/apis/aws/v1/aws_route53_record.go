
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
	Records	string	`json:"records"`
	AllowOverwrite	bool	`json:"allow_overwrite"`
	Ttl	int	`json:"ttl"`
	WeightedRoutingPolicy	[]kUvYxToc	`json:"weighted_routing_policy"`
	MultivalueAnswerRoutingPolicy	bool	`json:"multivalue_answer_routing_policy"`
	Name	string	`json:"name"`
	GeolocationRoutingPolicy	[]DoTiUUJe	`json:"geolocation_routing_policy"`
	HealthCheckId	string	`json:"health_check_id"`
	Type	string	`json:"type"`
	SetIdentifier	string	`json:"set_identifier"`
	Failover	string	`json:"failover"`
	Alias	string	`json:"alias"`
	FailoverRoutingPolicy	[]VakGzAAL	`json:"failover_routing_policy"`
	LatencyRoutingPolicy	[]tLhTUZHJ	`json:"latency_routing_policy"`
	ZoneId	string	`json:"zone_id"`
	Weight	int	`json:"weight"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53RecordList is a list of AwsRoute53Record resources
type AwsRoute53RecordList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53Record	`json:"items"`
}


// kUvYxToc is a kUvYxToc
type kUvYxToc struct {
	Weight	int	`json:"weight"`
}

// DoTiUUJe is a DoTiUUJe
type DoTiUUJe struct {
	Continent	string	`json:"continent"`
	Country	string	`json:"country"`
	Subdivision	string	`json:"subdivision"`
}

// VakGzAAL is a VakGzAAL
type VakGzAAL struct {
	Type	string	`json:"type"`
}

// tLhTUZHJ is a tLhTUZHJ
type tLhTUZHJ struct {
	Region	string	`json:"region"`
}
