
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPipeline describes a AwsElastictranscoderPipeline resource
type AwsElastictranscoderPipeline struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElastictranscoderPipelineSpec	`json:"spec"`
}


// AwsElastictranscoderPipelineSpec is the spec for a AwsElastictranscoderPipeline Resource
type AwsElastictranscoderPipelineSpec struct {
	AwsKmsKeyArn	string	`json:"aws_kms_key_arn"`
	ContentConfigPermissions	string	`json:"content_config_permissions"`
	InputBucket	string	`json:"input_bucket"`
	Role	string	`json:"role"`
	Notifications	string	`json:"notifications"`
	ThumbnailConfigPermissions	string	`json:"thumbnail_config_permissions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPipelineList is a list of AwsElastictranscoderPipeline resources
type AwsElastictranscoderPipelineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElastictranscoderPipeline	`json:"items"`
}

