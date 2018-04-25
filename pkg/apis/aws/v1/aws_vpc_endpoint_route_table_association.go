
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointRouteTableAssociation describes a AwsVpcEndpointRouteTableAssociation resource
type AwsVpcEndpointRouteTableAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcEndpointRouteTableAssociationSpec	`json:"spec"`
}


// AwsVpcEndpointRouteTableAssociationSpec is the spec for a AwsVpcEndpointRouteTableAssociation Resource
type AwsVpcEndpointRouteTableAssociationSpec struct {
	VpcEndpointId	string	`json:"vpc_endpoint_id"`
	RouteTableId	string	`json:"route_table_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointRouteTableAssociationList is a list of AwsVpcEndpointRouteTableAssociation resources
type AwsVpcEndpointRouteTableAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpointRouteTableAssociation	`json:"items"`
}

