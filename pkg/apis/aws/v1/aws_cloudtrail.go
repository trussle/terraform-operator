
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudtrail struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudtrailSpec	`json:"spec"`
}

type AwsCloudtrailList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudtrail	`json:"items"`
}

type AwsCloudtrailSpec struct {
	IncludeGlobalServiceEvents	bool	`json:"include_global_service_events"`
	IsMultiRegionTrail	bool	`json:"is_multi_region_trail"`
	Name	string	`json:"name"`
	CloudWatchLogsRoleArn	string	`json:"cloud_watch_logs_role_arn"`
	SnsTopicName	string	`json:"sns_topic_name"`
	EnableLogFileValidation	bool	`json:"enable_log_file_validation"`
	Tags	map[string]interface{}	`json:"tags"`
	EnableLogging	bool	`json:"enable_logging"`
	S3KeyPrefix	string	`json:"s3_key_prefix"`
	KmsKeyId	string	`json:"kms_key_id"`
	EventSelector	[]interface{}	`json:"event_selector"`
	S3BucketName	string	`json:"s3_bucket_name"`
	CloudWatchLogsGroupArn	string	`json:"cloud_watch_logs_group_arn"`
}
