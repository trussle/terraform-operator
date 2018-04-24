
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchComputeEnvironment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsBatchComputeEnvironmentSpec	`json:"spec"`
}

type AwsBatchComputeEnvironmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsBatchComputeEnvironment	`json:"items"`
}

type AwsBatchComputeEnvironmentSpec struct {
	ComputeResources	[]interface{}	`json:"compute_resources"`
	ServiceRole	string	`json:"service_role"`
	Type	string	`json:"type"`
	ComputeEnvironmentName	string	`json:"compute_environment_name"`
	State	string	`json:"state"`
}
