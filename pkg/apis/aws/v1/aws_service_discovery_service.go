
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryService struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsServiceDiscoveryServiceSpec	`json:"spec"`
}

type AwsServiceDiscoveryServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsServiceDiscoveryService	`json:"items"`
}

type AwsServiceDiscoveryServiceSpec struct {
	Description	string	`json:"description"`
	DnsConfig	[]interface{}	`json:"dns_config"`
	HealthCheckConfig	[]interface{}	`json:"health_check_config"`
	HealthCheckCustomConfig	[]interface{}	`json:"health_check_custom_config"`
	Name	string	`json:"name"`
}
