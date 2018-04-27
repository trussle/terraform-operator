
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRouteTable describes a AwsRouteTable resource
type AwsRouteTable struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRouteTableSpec	`json:"spec"`
}


// AwsRouteTableSpec is the spec for a AwsRouteTable Resource
type AwsRouteTableSpec struct {
	VpcId	string	`json:"vpc_id"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRouteTableList is a list of AwsRouteTable resources
type AwsRouteTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRouteTable	`json:"items"`
}


// AwsRouteTableRoute is a AwsRouteTableRoute
type AwsRouteTableRoute struct {
	NetworkInterfaceId	string	`json:"network_interface_id"`
	CidrBlock	string	`json:"cidr_block"`
	Ipv6CidrBlock	string	`json:"ipv6_cidr_block"`
	EgressOnlyGatewayId	string	`json:"egress_only_gateway_id"`
	GatewayId	string	`json:"gateway_id"`
	InstanceId	string	`json:"instance_id"`
	NatGatewayId	string	`json:"nat_gateway_id"`
	VpcPeeringConnectionId	string	`json:"vpc_peering_connection_id"`
}
