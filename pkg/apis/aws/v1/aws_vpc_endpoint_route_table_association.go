
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointRouteTableAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcEndpointRouteTableAssociationSpec	`json:"spec"`
}

type AwsVpcEndpointRouteTableAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpointRouteTableAssociation	`json:"items"`
}

type AwsVpcEndpointRouteTableAssociationSpec struct {
	VpcEndpointId	string	`json:"vpc_endpoint_id"`
	RouteTableId	string	`json:"route_table_id"`
}
