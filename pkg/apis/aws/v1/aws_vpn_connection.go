
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	VpnGatewayId	string	`json:"vpn_gateway_id"`
	Type	string	`json:"type"`
	Tags	map[string]string	`json:"tags"`
	CustomerGatewayId	string	`json:"customer_gateway_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnConnectionList is a list of AwsVpnConnection resources
type AwsVpnConnectionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpnConnection	`json:"items"`
}


// Routes is a Routes
type Routes struct {
}

// VgwTelemetry is a VgwTelemetry
type VgwTelemetry struct {
}
