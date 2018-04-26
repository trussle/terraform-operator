
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Parameters	map[string]string	`json:"parameters"`
	RetryStrategy	[]RetryStrategy	`json:"retry_strategy"`
	Type	string	`json:"type"`
	Name	string	`json:"name"`
	ContainerProperties	string	`json:"container_properties"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchJobDefinitionList is a list of AwsBatchJobDefinition resources
type AwsBatchJobDefinitionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsBatchJobDefinition	`json:"items"`
}


// RetryStrategy is a RetryStrategy
type RetryStrategy struct {
	Attempts	int	`json:"attempts"`
}
