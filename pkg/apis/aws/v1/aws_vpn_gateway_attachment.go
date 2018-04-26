
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnGatewayAttachment describes a AwsVpnGatewayAttachment resource
type AwsVpnGatewayAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpnGatewayAttachmentSpec	`json:"spec"`
}


// AwsVpnGatewayAttachmentSpec is the spec for a AwsVpnGatewayAttachment Resource
type AwsVpnGatewayAttachmentSpec struct {
	VpcId	string	`json:"vpc_id"`
	VpnGatewayId	string	`json:"vpn_gateway_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnGatewayAttachmentList is a list of AwsVpnGatewayAttachment resources
type AwsVpnGatewayAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpnGatewayAttachment	`json:"items"`
}

