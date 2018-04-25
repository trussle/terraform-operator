
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchComputeEnvironment describes a AwsBatchComputeEnvironment resource
type AwsBatchComputeEnvironment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsBatchComputeEnvironmentSpec	`json:"spec"`
}


// AwsBatchComputeEnvironmentSpec is the spec for a AwsBatchComputeEnvironment Resource
type AwsBatchComputeEnvironmentSpec struct {
	ComputeEnvironmentName	string	`json:"compute_environment_name"`
	ComputeResources	[]interface{}	`json:"compute_resources"`
	ServiceRole	string	`json:"service_role"`
	Type	string	`json:"type"`
	State	string	`json:"state"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchComputeEnvironmentList is a list of AwsBatchComputeEnvironment resources
type AwsBatchComputeEnvironmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsBatchComputeEnvironment	`json:"items"`
}

