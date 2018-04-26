
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
	Notifications	Notifications	`json:"notifications"`
	ThumbnailConfigPermissions	ThumbnailConfigPermissions	`json:"thumbnail_config_permissions"`
	ContentConfigPermissions	ContentConfigPermissions	`json:"content_config_permissions"`
	InputBucket	string	`json:"input_bucket"`
	Role	string	`json:"role"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPipelineList is a list of AwsElastictranscoderPipeline resources
type AwsElastictranscoderPipelineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElastictranscoderPipeline	`json:"items"`
}


// Notifications is a Notifications
type Notifications struct {
	Completed	string	`json:"completed"`
	Error	string	`json:"error"`
	Progressing	string	`json:"progressing"`
	Warning	string	`json:"warning"`
}

// ThumbnailConfigPermissions is a ThumbnailConfigPermissions
type ThumbnailConfigPermissions struct {
	Grantee	string	`json:"grantee"`
	GranteeType	string	`json:"grantee_type"`
	Access	[]string	`json:"access"`
}

// ContentConfigPermissions is a ContentConfigPermissions
type ContentConfigPermissions struct {
	GranteeType	string	`json:"grantee_type"`
	Access	[]string	`json:"access"`
	Grantee	string	`json:"grantee"`
}

// ThumbnailConfig is a ThumbnailConfig
type ThumbnailConfig struct {
	StorageClass	string	`json:"storage_class"`
}

// ContentConfig is a ContentConfig
type ContentConfig struct {
	StorageClass	string	`json:"storage_class"`
}
