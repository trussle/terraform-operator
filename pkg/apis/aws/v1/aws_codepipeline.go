
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
	ArtifactStore	[]ArtifactStore	`json:"artifact_store"`
	Stage	[]Stage	`json:"stage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodepipelineList is a list of AwsCodepipeline resources
type AwsCodepipelineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodepipeline	`json:"items"`
}


// ArtifactStore is a ArtifactStore
type ArtifactStore struct {
	Location	string	`json:"location"`
	Type	string	`json:"type"`
	EncryptionKey	[]EncryptionKey	`json:"encryption_key"`
}

// Action is a Action
type Action struct {
	Provider	string	`json:"provider"`
	InputArtifacts	[]string	`json:"input_artifacts"`
	OutputArtifacts	[]string	`json:"output_artifacts"`
	Configuration	map[string]string	`json:"configuration"`
	Category	string	`json:"category"`
	Owner	string	`json:"owner"`
	Version	string	`json:"version"`
	Name	string	`json:"name"`
	RoleArn	string	`json:"role_arn"`
}

// Stage is a Stage
type Stage struct {
	Name	string	`json:"name"`
	Action	[]Action	`json:"action"`
}

// EncryptionKey is a EncryptionKey
type EncryptionKey struct {
	Id	string	`json:"id"`
	Type	string	`json:"type"`
}
