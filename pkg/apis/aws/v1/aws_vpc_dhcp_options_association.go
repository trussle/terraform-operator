
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcDhcpOptionsAssociation describes a AwsVpcDhcpOptionsAssociation resource
type AwsVpcDhcpOptionsAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcDhcpOptionsAssociationSpec	`json:"spec"`
}


// AwsVpcDhcpOptionsAssociationSpec is the spec for a AwsVpcDhcpOptionsAssociation Resource
type AwsVpcDhcpOptionsAssociationSpec struct {
	VpcId	string	`json:"vpc_id"`
	DhcpOptionsId	string	`json:"dhcp_options_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcDhcpOptionsAssociationList is a list of AwsVpcDhcpOptionsAssociation resources
type AwsVpcDhcpOptionsAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcDhcpOptionsAssociation	`json:"items"`
}

