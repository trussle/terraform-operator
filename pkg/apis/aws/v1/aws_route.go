
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRouteSpec	`json:"spec"`
}

type AwsRouteList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute	`json:"items"`
}

type AwsRouteSpec struct {
	DestinationCidrBlock	string	`json:"destination_cidr_block"`
	DestinationIpv6CidrBlock	string	`json:"destination_ipv6_cidr_block"`
	VpcPeeringConnectionId	string	`json:"vpc_peering_connection_id"`
	RouteTableId	string	`json:"route_table_id"`
}
