
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailStaticIpAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLightsailStaticIpAttachmentSpec	`json:"spec"`
}

type AwsLightsailStaticIpAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLightsailStaticIpAttachment	`json:"items"`
}

type AwsLightsailStaticIpAttachmentSpec struct {
	InstanceName	string	`json:"instance_name"`
	StaticIpName	string	`json:"static_ip_name"`
}
