
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryPrivateDnsNamespace struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsServiceDiscoveryPrivateDnsNamespaceSpec	`json:"spec"`
}

type AwsServiceDiscoveryPrivateDnsNamespaceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsServiceDiscoveryPrivateDnsNamespace	`json:"items"`
}

type AwsServiceDiscoveryPrivateDnsNamespaceSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Vpc	string	`json:"vpc"`
}
