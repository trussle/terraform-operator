
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryPublicDnsNamespace describes a AwsServiceDiscoveryPublicDnsNamespace resource
type AwsServiceDiscoveryPublicDnsNamespace struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsServiceDiscoveryPublicDnsNamespaceSpec	`json:"spec"`
}


// AwsServiceDiscoveryPublicDnsNamespaceSpec is the spec for a AwsServiceDiscoveryPublicDnsNamespace Resource
type AwsServiceDiscoveryPublicDnsNamespaceSpec struct {
	Description	string	`json:"description"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryPublicDnsNamespaceList is a list of AwsServiceDiscoveryPublicDnsNamespace resources
type AwsServiceDiscoveryPublicDnsNamespaceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsServiceDiscoveryPublicDnsNamespace	`json:"items"`
}

