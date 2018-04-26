
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute describes a AwsRoute resource
type AwsRoute struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRouteSpec	`json:"spec"`
}


// AwsRouteSpec is the spec for a AwsRoute Resource
type AwsRouteSpec struct {
	DestinationCidrBlock	string	`json:"destination_cidr_block"`
	DestinationIpv6CidrBlock	string	`json:"destination_ipv6_cidr_block"`
	RouteTableId	string	`json:"route_table_id"`
	VpcPeeringConnectionId	string	`json:"vpc_peering_connection_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRouteList is a list of AwsRoute resources
type AwsRouteList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute	`json:"items"`
}

