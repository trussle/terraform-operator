
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointSubnetAssociation describes a AwsVpcEndpointSubnetAssociation resource
type AwsVpcEndpointSubnetAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcEndpointSubnetAssociationSpec	`json:"spec"`
}


// AwsVpcEndpointSubnetAssociationSpec is the spec for a AwsVpcEndpointSubnetAssociation Resource
type AwsVpcEndpointSubnetAssociationSpec struct {
	SubnetId	string	`json:"subnet_id"`
	VpcEndpointId	string	`json:"vpc_endpoint_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointSubnetAssociationList is a list of AwsVpcEndpointSubnetAssociation resources
type AwsVpcEndpointSubnetAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpointSubnetAssociation	`json:"items"`
}

