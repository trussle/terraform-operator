
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnGatewayRoutePropagation describes a AwsVpnGatewayRoutePropagation resource
type AwsVpnGatewayRoutePropagation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpnGatewayRoutePropagationSpec	`json:"spec"`
}


// AwsVpnGatewayRoutePropagationSpec is the spec for a AwsVpnGatewayRoutePropagation Resource
type AwsVpnGatewayRoutePropagationSpec struct {
	VpnGatewayId	string	`json:"vpn_gateway_id"`
	RouteTableId	string	`json:"route_table_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnGatewayRoutePropagationList is a list of AwsVpnGatewayRoutePropagation resources
type AwsVpnGatewayRoutePropagationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpnGatewayRoutePropagation	`json:"items"`
}

