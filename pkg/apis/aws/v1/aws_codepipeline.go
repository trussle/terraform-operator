
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
	ArtifactStore	[]XVlBzgba	`json:"artifact_store"`
	Stage	[]hTHctcuA	`json:"stage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodepipelineList is a list of AwsCodepipeline resources
type AwsCodepipelineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodepipeline	`json:"items"`
}


// iCMRAjWw is a iCMRAjWw
type iCMRAjWw struct {
	Id	string	`json:"id"`
	Type	string	`json:"type"`
}

// XVlBzgba is a XVlBzgba
type XVlBzgba struct {
	EncryptionKey	[]iCMRAjWw	`json:"encryption_key"`
	Location	string	`json:"location"`
	Type	string	`json:"type"`
}

// xhxKQFDa is a xhxKQFDa
type xhxKQFDa struct {
	RoleArn	string	`json:"role_arn"`
	Category	string	`json:"category"`
	Provider	string	`json:"provider"`
	Version	string	`json:"version"`
	OutputArtifacts	[]string	`json:"output_artifacts"`
	Configuration	map[string]???	`json:"configuration"`
	Owner	string	`json:"owner"`
	InputArtifacts	[]string	`json:"input_artifacts"`
	Name	string	`json:"name"`
}

// hTHctcuA is a hTHctcuA
type hTHctcuA struct {
	Name	string	`json:"name"`
	Action	[]xhxKQFDa	`json:"action"`
}
