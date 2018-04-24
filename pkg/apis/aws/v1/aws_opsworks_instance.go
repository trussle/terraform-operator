
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksInstanceSpec	`json:"spec"`
}

type AwsOpsworksInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksInstance	`json:"items"`
}

type AwsOpsworksInstanceSpec struct {
	LayerIds	[]interface{}	`json:"layer_ids"`
	State	string	`json:"state"`
	StackId	string	`json:"stack_id"`
	Architecture	string	`json:"architecture"`
	AgentVersion	string	`json:"agent_version"`
	DeleteEbs	bool	`json:"delete_ebs"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	AutoScalingType	string	`json:"auto_scaling_type"`
	DeleteEip	bool	`json:"delete_eip"`
	InstanceType	string	`json:"instance_type"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
}
