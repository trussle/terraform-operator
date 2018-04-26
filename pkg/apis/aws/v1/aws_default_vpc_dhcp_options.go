
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultVpcDhcpOptions describes a AwsDefaultVpcDhcpOptions resource
type AwsDefaultVpcDhcpOptions struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultVpcDhcpOptionsSpec	`json:"spec"`
}


// AwsDefaultVpcDhcpOptionsSpec is the spec for a AwsDefaultVpcDhcpOptions Resource
type AwsDefaultVpcDhcpOptionsSpec struct {
	Tags	map[string]???	`json:"tags"`
	NetbiosNodeType	string	`json:"netbios_node_type"`
	NetbiosNameServers	[]string	`json:"netbios_name_servers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultVpcDhcpOptionsList is a list of AwsDefaultVpcDhcpOptions resources
type AwsDefaultVpcDhcpOptionsList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultVpcDhcpOptions	`json:"items"`
}

