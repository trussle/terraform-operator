
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnConnection struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpnConnectionSpec	`json:"spec"`
}

type AwsVpnConnectionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpnConnection	`json:"items"`
}

type AwsVpnConnectionSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	CustomerGatewayId	string	`json:"customer_gateway_id"`
	Type	string	`json:"type"`
	VpnGatewayId	string	`json:"vpn_gateway_id"`
}
