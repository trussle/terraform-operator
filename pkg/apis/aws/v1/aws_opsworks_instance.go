
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
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	AgentVersion	string	`json:"agent_version"`
	LayerIds	[]string	`json:"layer_ids"`
	State	string	`json:"state"`
	StackId	string	`json:"stack_id"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	InstanceType	string	`json:"instance_type"`
	AutoScalingType	string	`json:"auto_scaling_type"`
	DeleteEbs	bool	`json:"delete_ebs"`
	DeleteEip	bool	`json:"delete_eip"`
	Architecture	string	`json:"architecture"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksInstanceList is a list of AwsOpsworksInstance resources
type AwsOpsworksInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksInstance	`json:"items"`
}

