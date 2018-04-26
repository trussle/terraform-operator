
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
	HealthCheckCustomConfig	[]ERwVhGCM	`json:"health_check_custom_config"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	DnsConfig	[]dfLitTqw	`json:"dns_config"`
	HealthCheckConfig	[]XTbRMGxq	`json:"health_check_config"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryServiceList is a list of AwsServiceDiscoveryService resources
type AwsServiceDiscoveryServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsServiceDiscoveryService	`json:"items"`
}


// ERwVhGCM is a ERwVhGCM
type ERwVhGCM struct {
	FailureThreshold	int	`json:"failure_threshold"`
}

// LUecgOcz is a LUecgOcz
type LUecgOcz struct {
	Ttl	int	`json:"ttl"`
	Type	string	`json:"type"`
}

// dfLitTqw is a dfLitTqw
type dfLitTqw struct {
	DnsRecords	[]LUecgOcz	`json:"dns_records"`
	RoutingPolicy	string	`json:"routing_policy"`
	NamespaceId	string	`json:"namespace_id"`
}

// XTbRMGxq is a XTbRMGxq
type XTbRMGxq struct {
	ResourcePath	string	`json:"resource_path"`
	Type	string	`json:"type"`
	FailureThreshold	int	`json:"failure_threshold"`
}
