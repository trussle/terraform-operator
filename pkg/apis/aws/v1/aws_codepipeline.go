
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	RoleArn	string	`json:"role_arn"`
	ArtifactStore	[]interface{}	`json:"artifact_store"`
	Stage	[]interface{}	`json:"stage"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodepipelineList is a list of AwsCodepipeline resources
type AwsCodepipelineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodepipeline	`json:"items"`
}

