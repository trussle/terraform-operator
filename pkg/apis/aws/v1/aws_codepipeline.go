
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodepipeline struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCodepipelineSpec	`json:"spec"`
}

type AwsCodepipelineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodepipeline	`json:"items"`
}

type AwsCodepipelineSpec struct {
	Stage	[]interface{}	`json:"stage"`
	Name	string	`json:"name"`
	RoleArn	string	`json:"role_arn"`
	ArtifactStore	[]interface{}	`json:"artifact_store"`
}
