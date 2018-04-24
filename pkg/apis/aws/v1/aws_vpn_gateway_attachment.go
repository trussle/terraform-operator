
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGatewayAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpnGatewayAttachmentSpec	`json:"spec"`
}

type AwsVpnGatewayAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpnGatewayAttachment	`json:"items"`
}

type AwsVpnGatewayAttachmentSpec struct {
	VpcId	string	`json:"vpc_id"`
	VpnGatewayId	string	`json:"vpn_gateway_id"`
}
