
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElastictranscoderPipeline struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElastictranscoderPipelineSpec	`json:"spec"`
}

type AwsElastictranscoderPipelineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElastictranscoderPipeline	`json:"items"`
}

type AwsElastictranscoderPipelineSpec struct {
	AwsKmsKeyArn	string	`json:"aws_kms_key_arn"`
	Notifications	interface{}	`json:"notifications"`
	ThumbnailConfigPermissions	interface{}	`json:"thumbnail_config_permissions"`
	InputBucket	string	`json:"input_bucket"`
	Role	string	`json:"role"`
	ContentConfigPermissions	interface{}	`json:"content_config_permissions"`
}
