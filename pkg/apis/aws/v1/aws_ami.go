
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmi struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAmiSpec	`json:"spec"`
}

type AwsAmiList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAmi	`json:"items"`
}

type AwsAmiSpec struct {
	RamdiskId	string	`json:"ramdisk_id"`
	Tags	map[string]interface{}	`json:"tags"`
	Description	string	`json:"description"`
	Architecture	string	`json:"architecture"`
	KernelId	string	`json:"kernel_id"`
	Name	string	`json:"name"`
	RootDeviceName	string	`json:"root_device_name"`
	SriovNetSupport	string	`json:"sriov_net_support"`
	VirtualizationType	string	`json:"virtualization_type"`
}
