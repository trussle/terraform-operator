
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudtrail describes a AwsCloudtrail resource
type AwsCloudtrail struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudtrailSpec	`json:"spec"`
}


// AwsCloudtrailSpec is the spec for a AwsCloudtrail Resource
type AwsCloudtrailSpec struct {
	CloudWatchLogsRoleArn	string	`json:"cloud_watch_logs_role_arn"`
	IsMultiRegionTrail	bool	`json:"is_multi_region_trail"`
	EventSelector	[]AwsCloudtrailEventSelector	`json:"event_selector"`
	Tags	map[string]string	`json:"tags"`
	EnableLogging	bool	`json:"enable_logging"`
	IncludeGlobalServiceEvents	bool	`json:"include_global_service_events"`
	SnsTopicName	string	`json:"sns_topic_name"`
	EnableLogFileValidation	bool	`json:"enable_log_file_validation"`
	KmsKeyId	string	`json:"kms_key_id"`
	Name	string	`json:"name"`
	S3BucketName	string	`json:"s3_bucket_name"`
	S3KeyPrefix	string	`json:"s3_key_prefix"`
	CloudWatchLogsGroupArn	string	`json:"cloud_watch_logs_group_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudtrailList is a list of AwsCloudtrail resources
type AwsCloudtrailList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudtrail	`json:"items"`
}


// AwsCloudtrailDataResource is a AwsCloudtrailDataResource
type AwsCloudtrailDataResource struct {
	Type	string	`json:"type"`
	Values	[]string	`json:"values"`
}

// AwsCloudtrailEventSelector is a AwsCloudtrailEventSelector
type AwsCloudtrailEventSelector struct {
	ReadWriteType	string	`json:"read_write_type"`
	IncludeManagementEvents	bool	`json:"include_management_events"`
	DataResource	[]AwsCloudtrailDataResource	`json:"data_resource"`
}
