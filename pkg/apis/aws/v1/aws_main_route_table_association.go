
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMainRouteTableAssociation describes a AwsMainRouteTableAssociation resource
type AwsMainRouteTableAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsMainRouteTableAssociationSpec	`json:"spec"`
}


// AwsMainRouteTableAssociationSpec is the spec for a AwsMainRouteTableAssociation Resource
type AwsMainRouteTableAssociationSpec struct {
	VpcId	string	`json:"vpc_id"`
	RouteTableId	string	`json:"route_table_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMainRouteTableAssociationList is a list of AwsMainRouteTableAssociation resources
type AwsMainRouteTableAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsMainRouteTableAssociation	`json:"items"`
}

