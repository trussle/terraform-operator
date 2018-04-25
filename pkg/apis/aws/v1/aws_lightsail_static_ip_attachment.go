
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailStaticIpAttachment describes a AwsLightsailStaticIpAttachment resource
type AwsLightsailStaticIpAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLightsailStaticIpAttachmentSpec	`json:"spec"`
}


// AwsLightsailStaticIpAttachmentSpec is the spec for a AwsLightsailStaticIpAttachment Resource
type AwsLightsailStaticIpAttachmentSpec struct {
	StaticIpName	string	`json:"static_ip_name"`
	InstanceName	string	`json:"instance_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailStaticIpAttachmentList is a list of AwsLightsailStaticIpAttachment resources
type AwsLightsailStaticIpAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLightsailStaticIpAttachment	`json:"items"`
}

