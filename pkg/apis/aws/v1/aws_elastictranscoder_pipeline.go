
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
	InputBucket	string	`json:"input_bucket"`
	ThumbnailConfigPermissions	AwsElastictranscoderPipelineThumbnailConfigPermissions	`json:"thumbnail_config_permissions"`
	ContentConfigPermissions	AwsElastictranscoderPipelineContentConfigPermissions	`json:"content_config_permissions"`
	Notifications	AwsElastictranscoderPipelineNotifications	`json:"notifications"`
	Role	string	`json:"role"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPipelineList is a list of AwsElastictranscoderPipeline resources
type AwsElastictranscoderPipelineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElastictranscoderPipeline	`json:"items"`
}


// AwsElastictranscoderPipelineThumbnailConfig is a AwsElastictranscoderPipelineThumbnailConfig
type AwsElastictranscoderPipelineThumbnailConfig struct {
	StorageClass	string	`json:"storage_class"`
}

// AwsElastictranscoderPipelineThumbnailConfigPermissions is a AwsElastictranscoderPipelineThumbnailConfigPermissions
type AwsElastictranscoderPipelineThumbnailConfigPermissions struct {
	Access	[]string	`json:"access"`
	Grantee	string	`json:"grantee"`
	GranteeType	string	`json:"grantee_type"`
}

// AwsElastictranscoderPipelineContentConfig is a AwsElastictranscoderPipelineContentConfig
type AwsElastictranscoderPipelineContentConfig struct {
	StorageClass	string	`json:"storage_class"`
}

// AwsElastictranscoderPipelineContentConfigPermissions is a AwsElastictranscoderPipelineContentConfigPermissions
type AwsElastictranscoderPipelineContentConfigPermissions struct {
	Access	[]string	`json:"access"`
	Grantee	string	`json:"grantee"`
	GranteeType	string	`json:"grantee_type"`
}

// AwsElastictranscoderPipelineNotifications is a AwsElastictranscoderPipelineNotifications
type AwsElastictranscoderPipelineNotifications struct {
	Warning	string	`json:"warning"`
	Completed	string	`json:"completed"`
	Error	string	`json:"error"`
	Progressing	string	`json:"progressing"`
}
