
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnConnection describes a AwsVpnConnection resource
type AwsVpnConnection struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpnConnectionSpec	`json:"spec"`
}


// AwsVpnConnectionSpec is the spec for a AwsVpnConnection Resource
type AwsVpnConnectionSpec struct {
	Tags	map[string]Generic	`json:"tags"`
	VpnGatewayId	string	`json:"vpn_gateway_id"`
	CustomerGatewayId	string	`json:"customer_gateway_id"`
	Type	string	`json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnConnectionList is a list of AwsVpnConnection resources
type AwsVpnConnectionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpnConnection	`json:"items"`
}

