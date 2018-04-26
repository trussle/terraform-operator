
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkInterfaceSgAttachment describes a AwsNetworkInterfaceSgAttachment resource
type AwsNetworkInterfaceSgAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNetworkInterfaceSgAttachmentSpec	`json:"spec"`
}


// AwsNetworkInterfaceSgAttachmentSpec is the spec for a AwsNetworkInterfaceSgAttachment Resource
type AwsNetworkInterfaceSgAttachmentSpec struct {
	SecurityGroupId	string	`json:"security_group_id"`
	NetworkInterfaceId	string	`json:"network_interface_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkInterfaceSgAttachmentList is a list of AwsNetworkInterfaceSgAttachment resources
type AwsNetworkInterfaceSgAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNetworkInterfaceSgAttachment	`json:"items"`
}

