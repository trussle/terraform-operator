
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	HealthCheckCustomConfig	[]interface{}	`json:"health_check_custom_config"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	DnsConfig	[]interface{}	`json:"dns_config"`
	HealthCheckConfig	[]interface{}	`json:"health_check_config"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryServiceList is a list of AwsServiceDiscoveryService resources
type AwsServiceDiscoveryServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsServiceDiscoveryService	`json:"items"`
}

