
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnConnectionRoute struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpnConnectionRouteSpec	`json:"spec"`
}

type AwsVpnConnectionRouteList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpnConnectionRoute	`json:"items"`
}

type AwsVpnConnectionRouteSpec struct {
	DestinationCidrBlock	string	`json:"destination_cidr_block"`
	VpnConnectionId	string	`json:"vpn_connection_id"`
}
