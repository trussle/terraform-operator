
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	TaskRoleArn	string	`json:"task_role_arn"`
	Volume	AwsEcsTaskDefinitionVolume	`json:"volume"`
	RequiresCompatibilities	string	`json:"requires_compatibilities"`
	Cpu	string	`json:"cpu"`
	Family	string	`json:"family"`
	ContainerDefinitions	string	`json:"container_definitions"`
	PlacementConstraints	AwsEcsTaskDefinitionPlacementConstraints	`json:"placement_constraints"`
	ExecutionRoleArn	string	`json:"execution_role_arn"`
	Memory	string	`json:"memory"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsTaskDefinitionList is a list of AwsEcsTaskDefinition resources
type AwsEcsTaskDefinitionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcsTaskDefinition	`json:"items"`
}


// AwsEcsTaskDefinitionVolume is a AwsEcsTaskDefinitionVolume
type AwsEcsTaskDefinitionVolume struct {
	Name	string	`json:"name"`
	HostPath	string	`json:"host_path"`
}

// AwsEcsTaskDefinitionPlacementConstraints is a AwsEcsTaskDefinitionPlacementConstraints
type AwsEcsTaskDefinitionPlacementConstraints struct {
	Expression	string	`json:"expression"`
	Type	string	`json:"type"`
}
