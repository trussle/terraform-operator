
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchJobDefinition describes a AwsBatchJobDefinition resource
type AwsBatchJobDefinition struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsBatchJobDefinitionSpec	`json:"spec"`
}


// AwsBatchJobDefinitionSpec is the spec for a AwsBatchJobDefinition Resource
type AwsBatchJobDefinitionSpec struct {
	Type	string	`json:"type"`
	Name	string	`json:"name"`
	ContainerProperties	string	`json:"container_properties"`
	Parameters	map[string]Generic	`json:"parameters"`
	RetryStrategy	[]Generic	`json:"retry_strategy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchJobDefinitionList is a list of AwsBatchJobDefinition resources
type AwsBatchJobDefinitionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsBatchJobDefinition	`json:"items"`
}

