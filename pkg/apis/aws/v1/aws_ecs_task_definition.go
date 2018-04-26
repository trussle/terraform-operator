
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsTaskDefinition describes a AwsEcsTaskDefinition resource
type AwsEcsTaskDefinition struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEcsTaskDefinitionSpec	`json:"spec"`
}


// AwsEcsTaskDefinitionSpec is the spec for a AwsEcsTaskDefinition Resource
type AwsEcsTaskDefinitionSpec struct {
	Cpu	string	`json:"cpu"`
	TaskRoleArn	string	`json:"task_role_arn"`
	ExecutionRoleArn	string	`json:"execution_role_arn"`
	Memory	string	`json:"memory"`
	Volume	Generic	`json:"volume"`
	ContainerDefinitions	string	`json:"container_definitions"`
	PlacementConstraints	Generic	`json:"placement_constraints"`
	RequiresCompatibilities	Generic	`json:"requires_compatibilities"`
	Family	string	`json:"family"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsTaskDefinitionList is a list of AwsEcsTaskDefinition resources
type AwsEcsTaskDefinitionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcsTaskDefinition	`json:"items"`
}

