
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
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	DnsConfig	[]AwsServiceDiscoveryServiceDnsConfig	`json:"dns_config"`
	HealthCheckConfig	[]AwsServiceDiscoveryServiceHealthCheckConfig	`json:"health_check_config"`
	HealthCheckCustomConfig	[]AwsServiceDiscoveryServiceHealthCheckCustomConfig	`json:"health_check_custom_config"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryServiceList is a list of AwsServiceDiscoveryService resources
type AwsServiceDiscoveryServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsServiceDiscoveryService	`json:"items"`
}


// AwsServiceDiscoveryServiceDnsRecords is a AwsServiceDiscoveryServiceDnsRecords
type AwsServiceDiscoveryServiceDnsRecords struct {
	Ttl	int	`json:"ttl"`
	Type	string	`json:"type"`
}

// AwsServiceDiscoveryServiceDnsConfig is a AwsServiceDiscoveryServiceDnsConfig
type AwsServiceDiscoveryServiceDnsConfig struct {
	NamespaceId	string	`json:"namespace_id"`
	DnsRecords	[]AwsServiceDiscoveryServiceDnsRecords	`json:"dns_records"`
	RoutingPolicy	string	`json:"routing_policy"`
}

// AwsServiceDiscoveryServiceHealthCheckConfig is a AwsServiceDiscoveryServiceHealthCheckConfig
type AwsServiceDiscoveryServiceHealthCheckConfig struct {
	Type	string	`json:"type"`
	FailureThreshold	int	`json:"failure_threshold"`
	ResourcePath	string	`json:"resource_path"`
}

// AwsServiceDiscoveryServiceHealthCheckCustomConfig is a AwsServiceDiscoveryServiceHealthCheckCustomConfig
type AwsServiceDiscoveryServiceHealthCheckCustomConfig struct {
	FailureThreshold	int	`json:"failure_threshold"`
}
