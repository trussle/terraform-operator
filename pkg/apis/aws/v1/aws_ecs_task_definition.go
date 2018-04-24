
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsTaskDefinition struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEcsTaskDefinitionSpec	`json:"spec"`
}

type AwsEcsTaskDefinitionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcsTaskDefinition	`json:"items"`
}

type AwsEcsTaskDefinitionSpec struct {
	ContainerDefinitions	string	`json:"container_definitions"`
	TaskRoleArn	string	`json:"task_role_arn"`
	ExecutionRoleArn	string	`json:"execution_role_arn"`
	PlacementConstraints	interface{}	`json:"placement_constraints"`
	Cpu	string	`json:"cpu"`
	Family	string	`json:"family"`
	RequiresCompatibilities	interface{}	`json:"requires_compatibilities"`
	Memory	string	`json:"memory"`
	Volume	interface{}	`json:"volume"`
}
