
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	ContentConfigPermissions	string	`json:"content_config_permissions"`
	InputBucket	string	`json:"input_bucket"`
	Notifications	string	`json:"notifications"`
	Role	string	`json:"role"`
	ThumbnailConfigPermissions	string	`json:"thumbnail_config_permissions"`
	AwsKmsKeyArn	string	`json:"aws_kms_key_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPipelineList is a list of AwsElastictranscoderPipeline resources
type AwsElastictranscoderPipelineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElastictranscoderPipeline	`json:"items"`
}

