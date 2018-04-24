
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkInterfaceSgAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNetworkInterfaceSgAttachmentSpec	`json:"spec"`
}

type AwsNetworkInterfaceSgAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNetworkInterfaceSgAttachment	`json:"items"`
}

type AwsNetworkInterfaceSgAttachmentSpec struct {
	NetworkInterfaceId	string	`json:"network_interface_id"`
	SecurityGroupId	string	`json:"security_group_id"`
}
