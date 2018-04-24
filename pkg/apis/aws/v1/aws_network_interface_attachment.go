
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkInterfaceAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNetworkInterfaceAttachmentSpec	`json:"spec"`
}

type AwsNetworkInterfaceAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNetworkInterfaceAttachment	`json:"items"`
}

type AwsNetworkInterfaceAttachmentSpec struct {
	InstanceId	string	`json:"instance_id"`
	NetworkInterfaceId	string	`json:"network_interface_id"`
	DeviceIndex	int	`json:"device_index"`
}
