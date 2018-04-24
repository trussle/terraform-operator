
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMainRouteTableAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsMainRouteTableAssociationSpec	`json:"spec"`
}

type AwsMainRouteTableAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsMainRouteTableAssociation	`json:"items"`
}

type AwsMainRouteTableAssociationSpec struct {
	VpcId	string	`json:"vpc_id"`
	RouteTableId	string	`json:"route_table_id"`
}
