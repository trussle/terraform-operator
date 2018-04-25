
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRouteTableAssociation describes a AwsRouteTableAssociation resource
type AwsRouteTableAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRouteTableAssociationSpec	`json:"spec"`
}


// AwsRouteTableAssociationSpec is the spec for a AwsRouteTableAssociation Resource
type AwsRouteTableAssociationSpec struct {
	SubnetId	string	`json:"subnet_id"`
	RouteTableId	string	`json:"route_table_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRouteTableAssociationList is a list of AwsRouteTableAssociation resources
type AwsRouteTableAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRouteTableAssociation	`json:"items"`
}

