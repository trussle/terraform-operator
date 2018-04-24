
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcDhcpOptions struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcDhcpOptionsSpec	`json:"spec"`
}

type AwsVpcDhcpOptionsList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcDhcpOptions	`json:"items"`
}

type AwsVpcDhcpOptionsSpec struct {
	DomainName	string	`json:"domain_name"`
	DomainNameServers	[]interface{}	`json:"domain_name_servers"`
	NtpServers	[]interface{}	`json:"ntp_servers"`
	NetbiosNodeType	string	`json:"netbios_node_type"`
	NetbiosNameServers	[]interface{}	`json:"netbios_name_servers"`
	Tags	map[string]interface{}	`json:"tags"`
}
