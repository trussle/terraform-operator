
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchJobDefinition struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsBatchJobDefinitionSpec	`json:"spec"`
}

type AwsBatchJobDefinitionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsBatchJobDefinition	`json:"items"`
}

type AwsBatchJobDefinitionSpec struct {
	Name	string	`json:"name"`
	ContainerProperties	string	`json:"container_properties"`
	Parameters	map[string]interface{}	`json:"parameters"`
	RetryStrategy	[]interface{}	`json:"retry_strategy"`
	Type	string	`json:"type"`
}
