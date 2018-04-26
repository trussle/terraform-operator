
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksInstance describes a AwsOpsworksInstance resource
type AwsOpsworksInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksInstanceSpec	`json:"spec"`
}


// AwsOpsworksInstanceSpec is the spec for a AwsOpsworksInstance Resource
type AwsOpsworksInstanceSpec struct {
	Architecture	string	`json:"architecture"`
	State	string	`json:"state"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	DeleteEbs	bool	`json:"delete_ebs"`
	AutoScalingType	string	`json:"auto_scaling_type"`
	InstanceType	string	`json:"instance_type"`
	StackId	string	`json:"stack_id"`
	AgentVersion	string	`json:"agent_version"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	LayerIds	[]string	`json:"layer_ids"`
	DeleteEip	bool	`json:"delete_eip"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksInstanceList is a list of AwsOpsworksInstance resources
type AwsOpsworksInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksInstance	`json:"items"`
}


// EphemeralBlockDevice is a EphemeralBlockDevice
type EphemeralBlockDevice struct {
	DeviceName	string	`json:"device_name"`
	VirtualName	string	`json:"virtual_name"`
}

// RootBlockDevice is a RootBlockDevice
type RootBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}

// EbsBlockDevice is a EbsBlockDevice
type EbsBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceName	string	`json:"device_name"`
}
