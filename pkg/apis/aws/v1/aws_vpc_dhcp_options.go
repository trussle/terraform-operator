
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcDhcpOptions describes a AwsVpcDhcpOptions resource
type AwsVpcDhcpOptions struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcDhcpOptionsSpec	`json:"spec"`
}


// AwsVpcDhcpOptionsSpec is the spec for a AwsVpcDhcpOptions Resource
type AwsVpcDhcpOptionsSpec struct {
	NetbiosNameServers	[]string	`json:"netbios_name_servers"`
	Tags	map[string]???	`json:"tags"`
	DomainName	string	`json:"domain_name"`
	DomainNameServers	[]string	`json:"domain_name_servers"`
	NtpServers	[]string	`json:"ntp_servers"`
	NetbiosNodeType	string	`json:"netbios_node_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcDhcpOptionsList is a list of AwsVpcDhcpOptions resources
type AwsVpcDhcpOptionsList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcDhcpOptions	`json:"items"`
}

