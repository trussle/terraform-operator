
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGatewayRoutePropagation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpnGatewayRoutePropagationSpec	`json:"spec"`
}

type AwsVpnGatewayRoutePropagationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpnGatewayRoutePropagation	`json:"items"`
}

type AwsVpnGatewayRoutePropagationSpec struct {
	RouteTableId	string	`json:"route_table_id"`
	VpnGatewayId	string	`json:"vpn_gateway_id"`
}
