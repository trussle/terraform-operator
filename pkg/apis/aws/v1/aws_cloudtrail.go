
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
	EnableLogging	bool	`json:"enable_logging"`
	IncludeGlobalServiceEvents	bool	`json:"include_global_service_events"`
	SnsTopicName	string	`json:"sns_topic_name"`
	EnableLogFileValidation	bool	`json:"enable_log_file_validation"`
	Tags	map[string]string	`json:"tags"`
	CloudWatchLogsRoleArn	string	`json:"cloud_watch_logs_role_arn"`
	CloudWatchLogsGroupArn	string	`json:"cloud_watch_logs_group_arn"`
	EventSelector	[]EventSelector	`json:"event_selector"`
	S3BucketName	string	`json:"s3_bucket_name"`
	IsMultiRegionTrail	bool	`json:"is_multi_region_trail"`
	KmsKeyId	string	`json:"kms_key_id"`
	Name	string	`json:"name"`
	S3KeyPrefix	string	`json:"s3_key_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudtrailList is a list of AwsCloudtrail resources
type AwsCloudtrailList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudtrail	`json:"items"`
}


// DataResource is a DataResource
type DataResource struct {
	Type	string	`json:"type"`
	Values	[]string	`json:"values"`
}

// EventSelector is a EventSelector
type EventSelector struct {
	ReadWriteType	string	`json:"read_write_type"`
	IncludeManagementEvents	bool	`json:"include_management_events"`
	DataResource	[]DataResource	`json:"data_resource"`
}
