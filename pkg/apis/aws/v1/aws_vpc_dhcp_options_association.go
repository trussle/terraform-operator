
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcDhcpOptionsAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcDhcpOptionsAssociationSpec	`json:"spec"`
}

type AwsVpcDhcpOptionsAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcDhcpOptionsAssociation	`json:"items"`
}

type AwsVpcDhcpOptionsAssociationSpec struct {
	VpcId	string	`json:"vpc_id"`
	DhcpOptionsId	string	`json:"dhcp_options_id"`
}
