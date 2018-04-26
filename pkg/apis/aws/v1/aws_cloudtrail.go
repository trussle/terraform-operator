
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
	IsMultiRegionTrail	bool	`json:"is_multi_region_trail"`
	S3BucketName	string	`json:"s3_bucket_name"`
	CloudWatchLogsRoleArn	string	`json:"cloud_watch_logs_role_arn"`
	IncludeGlobalServiceEvents	bool	`json:"include_global_service_events"`
	SnsTopicName	string	`json:"sns_topic_name"`
	KmsKeyId	string	`json:"kms_key_id"`
	Name	string	`json:"name"`
	EnableLogging	bool	`json:"enable_logging"`
	EventSelector	[]Generic	`json:"event_selector"`
	CloudWatchLogsGroupArn	string	`json:"cloud_watch_logs_group_arn"`
	EnableLogFileValidation	bool	`json:"enable_log_file_validation"`
	Tags	map[string]Generic	`json:"tags"`
	S3KeyPrefix	string	`json:"s3_key_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudtrailList is a list of AwsCloudtrail resources
type AwsCloudtrailList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudtrail	`json:"items"`
}

