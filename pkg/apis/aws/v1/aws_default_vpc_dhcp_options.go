
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultVpcDhcpOptions struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultVpcDhcpOptionsSpec	`json:"spec"`
}

type AwsDefaultVpcDhcpOptionsList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultVpcDhcpOptions	`json:"items"`
}

type AwsDefaultVpcDhcpOptionsSpec struct {
	NetbiosNodeType	string	`json:"netbios_node_type"`
	NetbiosNameServers	[]interface{}	`json:"netbios_name_servers"`
	Tags	map[string]interface{}	`json:"tags"`
}
