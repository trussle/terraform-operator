
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodepipeline describes a AwsCodepipeline resource
type AwsCodepipeline struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCodepipelineSpec	`json:"spec"`
}


// AwsCodepipelineSpec is the spec for a AwsCodepipeline Resource
type AwsCodepipelineSpec struct {
	Name	string	`json:"name"`
	RoleArn	string	`json:"role_arn"`
	ArtifactStore	[]AwsCodepipelineArtifactStore	`json:"artifact_store"`
	Stage	[]AwsCodepipelineStage	`json:"stage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodepipelineList is a list of AwsCodepipeline resources
type AwsCodepipelineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodepipeline	`json:"items"`
}


// AwsCodepipelineEncryptionKey is a AwsCodepipelineEncryptionKey
type AwsCodepipelineEncryptionKey struct {
	Id	string	`json:"id"`
	Type	string	`json:"type"`
}

// AwsCodepipelineArtifactStore is a AwsCodepipelineArtifactStore
type AwsCodepipelineArtifactStore struct {
	Location	string	`json:"location"`
	Type	string	`json:"type"`
	EncryptionKey	[]AwsCodepipelineEncryptionKey	`json:"encryption_key"`
}

// AwsCodepipelineAction is a AwsCodepipelineAction
type AwsCodepipelineAction struct {
	InputArtifacts	[]string	`json:"input_artifacts"`
	OutputArtifacts	[]string	`json:"output_artifacts"`
	RoleArn	string	`json:"role_arn"`
	Category	string	`json:"category"`
	Owner	string	`json:"owner"`
	Provider	string	`json:"provider"`
	Version	string	`json:"version"`
	Configuration	map[string]string	`json:"configuration"`
	Name	string	`json:"name"`
}

// AwsCodepipelineStage is a AwsCodepipelineStage
type AwsCodepipelineStage struct {
	Name	string	`json:"name"`
	Action	[]AwsCodepipelineAction	`json:"action"`
}
