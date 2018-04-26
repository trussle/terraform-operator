
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryService describes a AwsServiceDiscoveryService resource
type AwsServiceDiscoveryService struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsServiceDiscoveryServiceSpec	`json:"spec"`
}


// AwsServiceDiscoveryServiceSpec is the spec for a AwsServiceDiscoveryService Resource
type AwsServiceDiscoveryServiceSpec struct {
	Description	string	`json:"description"`
	DnsConfig	[]DnsConfig	`json:"dns_config"`
	HealthCheckConfig	[]HealthCheckConfig	`json:"health_check_config"`
	HealthCheckCustomConfig	[]HealthCheckCustomConfig	`json:"health_check_custom_config"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryServiceList is a list of AwsServiceDiscoveryService resources
type AwsServiceDiscoveryServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsServiceDiscoveryService	`json:"items"`
}


// DnsRecords is a DnsRecords
type DnsRecords struct {
	Ttl	int	`json:"ttl"`
	Type	string	`json:"type"`
}

// DnsConfig is a DnsConfig
type DnsConfig struct {
	NamespaceId	string	`json:"namespace_id"`
	DnsRecords	[]DnsRecords	`json:"dns_records"`
	RoutingPolicy	string	`json:"routing_policy"`
}

// HealthCheckConfig is a HealthCheckConfig
type HealthCheckConfig struct {
	FailureThreshold	int	`json:"failure_threshold"`
	ResourcePath	string	`json:"resource_path"`
	Type	string	`json:"type"`
}

// HealthCheckCustomConfig is a HealthCheckCustomConfig
type HealthCheckCustomConfig struct {
	FailureThreshold	int	`json:"failure_threshold"`
}
