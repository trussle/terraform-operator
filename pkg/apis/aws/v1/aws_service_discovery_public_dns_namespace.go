
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryPublicDnsNamespace struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsServiceDiscoveryPublicDnsNamespaceSpec	`json:"spec"`
}

type AwsServiceDiscoveryPublicDnsNamespaceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsServiceDiscoveryPublicDnsNamespace	`json:"items"`
}

type AwsServiceDiscoveryPublicDnsNamespaceSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
}
