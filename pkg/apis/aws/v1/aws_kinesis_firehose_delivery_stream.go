
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
	S3Configuration	[]S3Configuration	`json:"s3_configuration"`
	ExtendedS3Configuration	[]ExtendedS3Configuration	`json:"extended_s3_configuration"`
	RedshiftConfiguration	[]RedshiftConfiguration	`json:"redshift_configuration"`
	SplunkConfiguration	[]SplunkConfiguration	`json:"splunk_configuration"`
	Name	string	`json:"name"`
	KinesisSourceConfiguration	[]KinesisSourceConfiguration	`json:"kinesis_source_configuration"`
	Destination	string	`json:"destination"`
	ElasticsearchConfiguration	[]ElasticsearchConfiguration	`json:"elasticsearch_configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKinesisFirehoseDeliveryStreamList is a list of AwsKinesisFirehoseDeliveryStream resources
type AwsKinesisFirehoseDeliveryStreamList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKinesisFirehoseDeliveryStream	`json:"items"`
}


// Parameters is a Parameters
type Parameters struct {
	ParameterName	string	`json:"parameter_name"`
	ParameterValue	string	`json:"parameter_value"`
}

// ExtendedS3Configuration is a ExtendedS3Configuration
type ExtendedS3Configuration struct {
	BufferInterval	int	`json:"buffer_interval"`
	CompressionFormat	string	`json:"compression_format"`
	Prefix	string	`json:"prefix"`
	S3BackupMode	string	`json:"s3_backup_mode"`
	S3BackupConfiguration	[]S3BackupConfiguration	`json:"s3_backup_configuration"`
	BucketArn	string	`json:"bucket_arn"`
	BufferSize	int	`json:"buffer_size"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	RoleArn	string	`json:"role_arn"`
	ProcessingConfiguration	[]ProcessingConfiguration	`json:"processing_configuration"`
}

// SplunkConfiguration is a SplunkConfiguration
type SplunkConfiguration struct {
	ProcessingConfiguration	[]ProcessingConfiguration	`json:"processing_configuration"`
	HecAcknowledgmentTimeout	int	`json:"hec_acknowledgment_timeout"`
	HecEndpoint	string	`json:"hec_endpoint"`
	HecEndpointType	string	`json:"hec_endpoint_type"`
	HecToken	string	`json:"hec_token"`
	S3BackupMode	string	`json:"s3_backup_mode"`
	RetryDuration	int	`json:"retry_duration"`
}

// ElasticsearchConfiguration is a ElasticsearchConfiguration
type ElasticsearchConfiguration struct {
	DomainArn	string	`json:"domain_arn"`
	IndexName	string	`json:"index_name"`
	IndexRotationPeriod	string	`json:"index_rotation_period"`
	S3BackupMode	string	`json:"s3_backup_mode"`
	TypeName	string	`json:"type_name"`
	BufferingInterval	int	`json:"buffering_interval"`
	BufferingSize	int	`json:"buffering_size"`
	RetryDuration	int	`json:"retry_duration"`
	RoleArn	string	`json:"role_arn"`
	ProcessingConfiguration	[]ProcessingConfiguration	`json:"processing_configuration"`
}

// KinesisSourceConfiguration is a KinesisSourceConfiguration
type KinesisSourceConfiguration struct {
	KinesisStreamArn	string	`json:"kinesis_stream_arn"`
	RoleArn	string	`json:"role_arn"`
}

// CloudwatchLoggingOptions is a CloudwatchLoggingOptions
type CloudwatchLoggingOptions struct {
	Enabled	bool	`json:"enabled"`
	LogGroupName	string	`json:"log_group_name"`
	LogStreamName	string	`json:"log_stream_name"`
}

// S3Configuration is a S3Configuration
type S3Configuration struct {
	BufferInterval	int	`json:"buffer_interval"`
	CompressionFormat	string	`json:"compression_format"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	RoleArn	string	`json:"role_arn"`
	Prefix	string	`json:"prefix"`
	BucketArn	string	`json:"bucket_arn"`
	BufferSize	int	`json:"buffer_size"`
}

// S3BackupConfiguration is a S3BackupConfiguration
type S3BackupConfiguration struct {
	KmsKeyArn	string	`json:"kms_key_arn"`
	RoleArn	string	`json:"role_arn"`
	Prefix	string	`json:"prefix"`
	BucketArn	string	`json:"bucket_arn"`
	BufferSize	int	`json:"buffer_size"`
	BufferInterval	int	`json:"buffer_interval"`
	CompressionFormat	string	`json:"compression_format"`
}

// Processors is a Processors
type Processors struct {
	Type	string	`json:"type"`
	Parameters	[]Parameters	`json:"parameters"`
}

// ProcessingConfiguration is a ProcessingConfiguration
type ProcessingConfiguration struct {
	Enabled	bool	`json:"enabled"`
	Processors	[]Processors	`json:"processors"`
}

// RedshiftConfiguration is a RedshiftConfiguration
type RedshiftConfiguration struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	S3BackupConfiguration	[]S3BackupConfiguration	`json:"s3_backup_configuration"`
	RetryDuration	int	`json:"retry_duration"`
	CopyOptions	string	`json:"copy_options"`
	ClusterJdbcurl	string	`json:"cluster_jdbcurl"`
	ProcessingConfiguration	[]ProcessingConfiguration	`json:"processing_configuration"`
	RoleArn	string	`json:"role_arn"`
	S3BackupMode	string	`json:"s3_backup_mode"`
	DataTableColumns	string	`json:"data_table_columns"`
	DataTableName	string	`json:"data_table_name"`
}
