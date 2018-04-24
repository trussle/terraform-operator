
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointSubnetAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcEndpointSubnetAssociationSpec	`json:"spec"`
}

type AwsVpcEndpointSubnetAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpointSubnetAssociation	`json:"items"`
}

type AwsVpcEndpointSubnetAssociationSpec struct {
	VpcEndpointId	string	`json:"vpc_endpoint_id"`
	SubnetId	string	`json:"subnet_id"`
}
