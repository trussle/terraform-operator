
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmi describes a AwsAmi resource
type AwsAmi struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAmiSpec	`json:"spec"`
}


// AwsAmiSpec is the spec for a AwsAmi Resource
type AwsAmiSpec struct {
	Architecture	string	`json:"architecture"`
	RamdiskId	string	`json:"ramdisk_id"`
	SriovNetSupport	string	`json:"sriov_net_support"`
	VirtualizationType	string	`json:"virtualization_type"`
	Description	string	`json:"description"`
	RootDeviceName	string	`json:"root_device_name"`
	Tags	map[string]Generic	`json:"tags"`
	KernelId	string	`json:"kernel_id"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiList is a list of AwsAmi resources
type AwsAmiList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAmi	`json:"items"`
}

