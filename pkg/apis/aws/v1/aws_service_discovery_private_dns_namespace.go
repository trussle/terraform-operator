
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryPrivateDnsNamespace describes a AwsServiceDiscoveryPrivateDnsNamespace resource
type AwsServiceDiscoveryPrivateDnsNamespace struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsServiceDiscoveryPrivateDnsNamespaceSpec	`json:"spec"`
}


// AwsServiceDiscoveryPrivateDnsNamespaceSpec is the spec for a AwsServiceDiscoveryPrivateDnsNamespace Resource
type AwsServiceDiscoveryPrivateDnsNamespaceSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Vpc	string	`json:"vpc"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryPrivateDnsNamespaceList is a list of AwsServiceDiscoveryPrivateDnsNamespace resources
type AwsServiceDiscoveryPrivateDnsNamespaceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsServiceDiscoveryPrivateDnsNamespace	`json:"items"`
}

