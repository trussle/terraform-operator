
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Description	string	`json:"description"`
	RamdiskId	string	`json:"ramdisk_id"`
	VirtualizationType	string	`json:"virtualization_type"`
	Tags	map[string]string	`json:"tags"`
	Architecture	string	`json:"architecture"`
	Name	string	`json:"name"`
	RootDeviceName	string	`json:"root_device_name"`
	KernelId	string	`json:"kernel_id"`
	SriovNetSupport	string	`json:"sriov_net_support"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiList is a list of AwsAmi resources
type AwsAmiList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAmi	`json:"items"`
}


// EphemeralBlockDevice is a EphemeralBlockDevice
type EphemeralBlockDevice struct {
	DeviceName	string	`json:"device_name"`
	VirtualName	string	`json:"virtual_name"`
}

// EbsBlockDevice is a EbsBlockDevice
type EbsBlockDevice struct {
	Encrypted	bool	`json:"encrypted"`
	Iops	int	`json:"iops"`
	SnapshotId	string	`json:"snapshot_id"`
	VolumeType	string	`json:"volume_type"`
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceName	string	`json:"device_name"`
}
