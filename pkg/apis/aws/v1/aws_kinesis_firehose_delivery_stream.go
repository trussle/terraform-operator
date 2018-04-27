
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKinesisFirehoseDeliveryStream describes a AwsKinesisFirehoseDeliveryStream resource
type AwsKinesisFirehoseDeliveryStream struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKinesisFirehoseDeliveryStreamSpec	`json:"spec"`
}


// AwsKinesisFirehoseDeliveryStreamSpec is the spec for a AwsKinesisFirehoseDeliveryStream Resource
type AwsKinesisFirehoseDeliveryStreamSpec struct {
	KinesisSourceConfiguration	[]AwsKinesisFirehoseDeliveryStreamKinesisSourceConfiguration	`json:"kinesis_source_configuration"`
	Destination	string	`json:"destination"`
	ExtendedS3Configuration	[]AwsKinesisFirehoseDeliveryStreamExtendedS3Configuration	`json:"extended_s3_configuration"`
	RedshiftConfiguration	[]AwsKinesisFirehoseDeliveryStreamRedshiftConfiguration	`json:"redshift_configuration"`
	ElasticsearchConfiguration	[]AwsKinesisFirehoseDeliveryStreamElasticsearchConfiguration	`json:"elasticsearch_configuration"`
	SplunkConfiguration	[]AwsKinesisFirehoseDeliveryStreamSplunkConfiguration	`json:"splunk_configuration"`
	Name	string	`json:"name"`
	S3Configuration	[]AwsKinesisFirehoseDeliveryStreamS3Configuration	`json:"s3_configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKinesisFirehoseDeliveryStreamList is a list of AwsKinesisFirehoseDeliveryStream resources
type AwsKinesisFirehoseDeliveryStreamList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKinesisFirehoseDeliveryStream	`json:"items"`
}


// AwsKinesisFirehoseDeliveryStreamParameters is a AwsKinesisFirehoseDeliveryStreamParameters
type AwsKinesisFirehoseDeliveryStreamParameters struct {
	ParameterName	string	`json:"parameter_name"`
	ParameterValue	string	`json:"parameter_value"`
}

// AwsKinesisFirehoseDeliveryStreamProcessingConfiguration is a AwsKinesisFirehoseDeliveryStreamProcessingConfiguration
type AwsKinesisFirehoseDeliveryStreamProcessingConfiguration struct {
	Enabled	bool	`json:"enabled"`
	Processors	[]AwsKinesisFirehoseDeliveryStreamProcessors	`json:"processors"`
}

// AwsKinesisFirehoseDeliveryStreamExtendedS3Configuration is a AwsKinesisFirehoseDeliveryStreamExtendedS3Configuration
type AwsKinesisFirehoseDeliveryStreamExtendedS3Configuration struct {
	BufferSize	int	`json:"buffer_size"`
	BufferInterval	int	`json:"buffer_interval"`
	S3BackupMode	string	`json:"s3_backup_mode"`
	Prefix	string	`json:"prefix"`
	S3BackupConfiguration	[]AwsKinesisFirehoseDeliveryStreamS3BackupConfiguration	`json:"s3_backup_configuration"`
	ProcessingConfiguration	[]AwsKinesisFirehoseDeliveryStreamProcessingConfiguration	`json:"processing_configuration"`
	BucketArn	string	`json:"bucket_arn"`
	CompressionFormat	string	`json:"compression_format"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	RoleArn	string	`json:"role_arn"`
}

// AwsKinesisFirehoseDeliveryStreamRedshiftConfiguration is a AwsKinesisFirehoseDeliveryStreamRedshiftConfiguration
type AwsKinesisFirehoseDeliveryStreamRedshiftConfiguration struct {
	ClusterJdbcurl	string	`json:"cluster_jdbcurl"`
	ProcessingConfiguration	[]AwsKinesisFirehoseDeliveryStreamProcessingConfiguration	`json:"processing_configuration"`
	S3BackupConfiguration	[]AwsKinesisFirehoseDeliveryStreamS3BackupConfiguration	`json:"s3_backup_configuration"`
	DataTableColumns	string	`json:"data_table_columns"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	RoleArn	string	`json:"role_arn"`
	S3BackupMode	string	`json:"s3_backup_mode"`
	RetryDuration	int	`json:"retry_duration"`
	CopyOptions	string	`json:"copy_options"`
	DataTableName	string	`json:"data_table_name"`
}

// AwsKinesisFirehoseDeliveryStreamElasticsearchConfiguration is a AwsKinesisFirehoseDeliveryStreamElasticsearchConfiguration
type AwsKinesisFirehoseDeliveryStreamElasticsearchConfiguration struct {
	RoleArn	string	`json:"role_arn"`
	TypeName	string	`json:"type_name"`
	BufferingInterval	int	`json:"buffering_interval"`
	DomainArn	string	`json:"domain_arn"`
	IndexName	string	`json:"index_name"`
	S3BackupMode	string	`json:"s3_backup_mode"`
	ProcessingConfiguration	[]AwsKinesisFirehoseDeliveryStreamProcessingConfiguration	`json:"processing_configuration"`
	BufferingSize	int	`json:"buffering_size"`
	IndexRotationPeriod	string	`json:"index_rotation_period"`
	RetryDuration	int	`json:"retry_duration"`
}

// AwsKinesisFirehoseDeliveryStreamKinesisSourceConfiguration is a AwsKinesisFirehoseDeliveryStreamKinesisSourceConfiguration
type AwsKinesisFirehoseDeliveryStreamKinesisSourceConfiguration struct {
	KinesisStreamArn	string	`json:"kinesis_stream_arn"`
	RoleArn	string	`json:"role_arn"`
}

// AwsKinesisFirehoseDeliveryStreamCloudwatchLoggingOptions is a AwsKinesisFirehoseDeliveryStreamCloudwatchLoggingOptions
type AwsKinesisFirehoseDeliveryStreamCloudwatchLoggingOptions struct {
	LogGroupName	string	`json:"log_group_name"`
	LogStreamName	string	`json:"log_stream_name"`
	Enabled	bool	`json:"enabled"`
}

// AwsKinesisFirehoseDeliveryStreamS3BackupConfiguration is a AwsKinesisFirehoseDeliveryStreamS3BackupConfiguration
type AwsKinesisFirehoseDeliveryStreamS3BackupConfiguration struct {
	BucketArn	string	`json:"bucket_arn"`
	BufferSize	int	`json:"buffer_size"`
	BufferInterval	int	`json:"buffer_interval"`
	CompressionFormat	string	`json:"compression_format"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	RoleArn	string	`json:"role_arn"`
	Prefix	string	`json:"prefix"`
}

// AwsKinesisFirehoseDeliveryStreamProcessors is a AwsKinesisFirehoseDeliveryStreamProcessors
type AwsKinesisFirehoseDeliveryStreamProcessors struct {
	Parameters	[]AwsKinesisFirehoseDeliveryStreamParameters	`json:"parameters"`
	Type	string	`json:"type"`
}

// AwsKinesisFirehoseDeliveryStreamSplunkConfiguration is a AwsKinesisFirehoseDeliveryStreamSplunkConfiguration
type AwsKinesisFirehoseDeliveryStreamSplunkConfiguration struct {
	ProcessingConfiguration	[]AwsKinesisFirehoseDeliveryStreamProcessingConfiguration	`json:"processing_configuration"`
	HecAcknowledgmentTimeout	int	`json:"hec_acknowledgment_timeout"`
	HecEndpoint	string	`json:"hec_endpoint"`
	HecEndpointType	string	`json:"hec_endpoint_type"`
	HecToken	string	`json:"hec_token"`
	S3BackupMode	string	`json:"s3_backup_mode"`
	RetryDuration	int	`json:"retry_duration"`
}

// AwsKinesisFirehoseDeliveryStreamS3Configuration is a AwsKinesisFirehoseDeliveryStreamS3Configuration
type AwsKinesisFirehoseDeliveryStreamS3Configuration struct {
	Prefix	string	`json:"prefix"`
	BucketArn	string	`json:"bucket_arn"`
	BufferSize	int	`json:"buffer_size"`
	BufferInterval	int	`json:"buffer_interval"`
	CompressionFormat	string	`json:"compression_format"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	RoleArn	string	`json:"role_arn"`
}
