
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
	State	string	`json:"state"`
	InstanceType	string	`json:"instance_type"`
	DeleteEip	bool	`json:"delete_eip"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	AutoScalingType	string	`json:"auto_scaling_type"`
	DeleteEbs	bool	`json:"delete_ebs"`
	LayerIds	[]string	`json:"layer_ids"`
	AgentVersion	string	`json:"agent_version"`
	Architecture	string	`json:"architecture"`
	StackId	string	`json:"stack_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksInstanceList is a list of AwsOpsworksInstance resources
type AwsOpsworksInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksInstance	`json:"items"`
}


// AwsOpsworksInstanceEbsBlockDevice is a AwsOpsworksInstanceEbsBlockDevice
type AwsOpsworksInstanceEbsBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceName	string	`json:"device_name"`
}

// AwsOpsworksInstanceRootBlockDevice is a AwsOpsworksInstanceRootBlockDevice
type AwsOpsworksInstanceRootBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}

// AwsOpsworksInstanceEphemeralBlockDevice is a AwsOpsworksInstanceEphemeralBlockDevice
type AwsOpsworksInstanceEphemeralBlockDevice struct {
	DeviceName	string	`json:"device_name"`
	VirtualName	string	`json:"virtual_name"`
}
