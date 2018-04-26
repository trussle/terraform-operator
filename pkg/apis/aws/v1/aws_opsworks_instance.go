
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	State	string	`json:"state"`
	LayerIds	[]Generic	`json:"layer_ids"`
	DeleteEbs	bool	`json:"delete_ebs"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	AutoScalingType	string	`json:"auto_scaling_type"`
	DeleteEip	bool	`json:"delete_eip"`
	Architecture	string	`json:"architecture"`
	AgentVersion	string	`json:"agent_version"`
	InstanceType	string	`json:"instance_type"`
	StackId	string	`json:"stack_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksInstanceList is a list of AwsOpsworksInstance resources
type AwsOpsworksInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksInstance	`json:"items"`
}

