
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultRouteTable describes a AwsDefaultRouteTable resource
type AwsDefaultRouteTable struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultRouteTableSpec	`json:"spec"`
}


// AwsDefaultRouteTableSpec is the spec for a AwsDefaultRouteTable Resource
type AwsDefaultRouteTableSpec struct {
	Tags	map[string]string	`json:"tags"`
	DefaultRouteTableId	string	`json:"default_route_table_id"`
	PropagatingVgws	string	`json:"propagating_vgws"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultRouteTableList is a list of AwsDefaultRouteTable resources
type AwsDefaultRouteTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultRouteTable	`json:"items"`
}


// Route is a Route
type Route struct {
	NetworkInterfaceId	string	`json:"network_interface_id"`
	CidrBlock	string	`json:"cidr_block"`
	Ipv6CidrBlock	string	`json:"ipv6_cidr_block"`
	EgressOnlyGatewayId	string	`json:"egress_only_gateway_id"`
	GatewayId	string	`json:"gateway_id"`
	InstanceId	string	`json:"instance_id"`
	NatGatewayId	string	`json:"nat_gateway_id"`
	VpcPeeringConnectionId	string	`json:"vpc_peering_connection_id"`
}
