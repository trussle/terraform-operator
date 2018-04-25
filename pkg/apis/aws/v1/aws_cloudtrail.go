
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	EnableLogging	bool	`json:"enable_logging"`
	IncludeGlobalServiceEvents	bool	`json:"include_global_service_events"`
	EventSelector	[]interface{}	`json:"event_selector"`
	Name	string	`json:"name"`
	S3KeyPrefix	string	`json:"s3_key_prefix"`
	CloudWatchLogsRoleArn	string	`json:"cloud_watch_logs_role_arn"`
	CloudWatchLogsGroupArn	string	`json:"cloud_watch_logs_group_arn"`
	IsMultiRegionTrail	bool	`json:"is_multi_region_trail"`
	EnableLogFileValidation	bool	`json:"enable_log_file_validation"`
	S3BucketName	string	`json:"s3_bucket_name"`
	SnsTopicName	string	`json:"sns_topic_name"`
	KmsKeyId	string	`json:"kms_key_id"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudtrailList is a list of AwsCloudtrail resources
type AwsCloudtrailList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudtrail	`json:"items"`
}

