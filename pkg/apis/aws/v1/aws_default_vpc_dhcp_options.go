
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	NetbiosNodeType	string	`json:"netbios_node_type"`
	NetbiosNameServers	[]Generic	`json:"netbios_name_servers"`
	Tags	map[string]Generic	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultVpcDhcpOptionsList is a list of AwsDefaultVpcDhcpOptions resources
type AwsDefaultVpcDhcpOptionsList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultVpcDhcpOptions	`json:"items"`
}

