
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
	RedshiftConfiguration	[]CsNVlgTe	`json:"redshift_configuration"`
	SplunkConfiguration	[]etHsbZRj	`json:"splunk_configuration"`
	KinesisSourceConfiguration	[]ZLCtTMtT	`json:"kinesis_source_configuration"`
	S3Configuration	[]CoaNatyy	`json:"s3_configuration"`
	ExtendedS3Configuration	[]iNKAReKJ	`json:"extended_s3_configuration"`
	Name	string	`json:"name"`
	Destination	string	`json:"destination"`
	ElasticsearchConfiguration	[]iFQGZsnw	`json:"elasticsearch_configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKinesisFirehoseDeliveryStreamList is a list of AwsKinesisFirehoseDeliveryStream resources
type AwsKinesisFirehoseDeliveryStreamList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKinesisFirehoseDeliveryStream	`json:"items"`
}


// JjPjzpfR is a JjPjzpfR
type JjPjzpfR struct {
	ParameterName	string	`json:"parameter_name"`
	ParameterValue	string	`json:"parameter_value"`
}

// QYhYzRyW is a QYhYzRyW
type QYhYzRyW struct {
	Parameters	[]JjPjzpfR	`json:"parameters"`
	Type	string	`json:"type"`
}

// MaPEZQle is a MaPEZQle
type MaPEZQle struct {
	Enabled	bool	`json:"enabled"`
	Processors	[]QYhYzRyW	`json:"processors"`
}

// FEgmotaF is a FEgmotaF
type FEgmotaF struct {
	BufferSize	int	`json:"buffer_size"`
	BufferInterval	int	`json:"buffer_interval"`
	CompressionFormat	string	`json:"compression_format"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	RoleArn	string	`json:"role_arn"`
	Prefix	string	`json:"prefix"`
	BucketArn	string	`json:"bucket_arn"`
}

// CsNVlgTe is a CsNVlgTe
type CsNVlgTe struct {
	Username	string	`json:"username"`
	ProcessingConfiguration	[]MaPEZQle	`json:"processing_configuration"`
	S3BackupConfiguration	[]FEgmotaF	`json:"s3_backup_configuration"`
	DataTableColumns	string	`json:"data_table_columns"`
	DataTableName	string	`json:"data_table_name"`
	CopyOptions	string	`json:"copy_options"`
	ClusterJdbcurl	string	`json:"cluster_jdbcurl"`
	Password	string	`json:"password"`
	RoleArn	string	`json:"role_arn"`
	S3BackupMode	string	`json:"s3_backup_mode"`
	RetryDuration	int	`json:"retry_duration"`
}

// EkXBAkjQ is a EkXBAkjQ
type EkXBAkjQ struct {
	ParameterName	string	`json:"parameter_name"`
	ParameterValue	string	`json:"parameter_value"`
}

// BEmfdzdc is a BEmfdzdc
type BEmfdzdc struct {
	Type	string	`json:"type"`
	Parameters	[]EkXBAkjQ	`json:"parameters"`
}

// xAwnwekr is a xAwnwekr
type xAwnwekr struct {
	Enabled	bool	`json:"enabled"`
	Processors	[]BEmfdzdc	`json:"processors"`
}

// etHsbZRj is a etHsbZRj
type etHsbZRj struct {
	RetryDuration	int	`json:"retry_duration"`
	ProcessingConfiguration	[]xAwnwekr	`json:"processing_configuration"`
	HecAcknowledgmentTimeout	int	`json:"hec_acknowledgment_timeout"`
	HecEndpoint	string	`json:"hec_endpoint"`
	HecEndpointType	string	`json:"hec_endpoint_type"`
	HecToken	string	`json:"hec_token"`
	S3BackupMode	string	`json:"s3_backup_mode"`
}

// ZLCtTMtT is a ZLCtTMtT
type ZLCtTMtT struct {
	KinesisStreamArn	string	`json:"kinesis_stream_arn"`
	RoleArn	string	`json:"role_arn"`
}

// CoaNatyy is a CoaNatyy
type CoaNatyy struct {
	BufferSize	int	`json:"buffer_size"`
	BufferInterval	int	`json:"buffer_interval"`
	CompressionFormat	string	`json:"compression_format"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	RoleArn	string	`json:"role_arn"`
	Prefix	string	`json:"prefix"`
	BucketArn	string	`json:"bucket_arn"`
}

// yiXJrscc is a yiXJrscc
type yiXJrscc struct {
	KmsKeyArn	string	`json:"kms_key_arn"`
	RoleArn	string	`json:"role_arn"`
	Prefix	string	`json:"prefix"`
	BucketArn	string	`json:"bucket_arn"`
	BufferSize	int	`json:"buffer_size"`
	BufferInterval	int	`json:"buffer_interval"`
	CompressionFormat	string	`json:"compression_format"`
}

// zFZBsbOJ is a zFZBsbOJ
type zFZBsbOJ struct {
	ParameterName	string	`json:"parameter_name"`
	ParameterValue	string	`json:"parameter_value"`
}

// RussVmao is a RussVmao
type RussVmao struct {
	Parameters	[]zFZBsbOJ	`json:"parameters"`
	Type	string	`json:"type"`
}

// tNswYNsG is a tNswYNsG
type tNswYNsG struct {
	Enabled	bool	`json:"enabled"`
	Processors	[]RussVmao	`json:"processors"`
}

// iNKAReKJ is a iNKAReKJ
type iNKAReKJ struct {
	BufferSize	int	`json:"buffer_size"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	RoleArn	string	`json:"role_arn"`
	Prefix	string	`json:"prefix"`
	S3BackupConfiguration	[]yiXJrscc	`json:"s3_backup_configuration"`
	ProcessingConfiguration	[]tNswYNsG	`json:"processing_configuration"`
	BucketArn	string	`json:"bucket_arn"`
	BufferInterval	int	`json:"buffer_interval"`
	CompressionFormat	string	`json:"compression_format"`
	S3BackupMode	string	`json:"s3_backup_mode"`
}

// dKupdOMe is a dKupdOMe
type dKupdOMe struct {
	ParameterName	string	`json:"parameter_name"`
	ParameterValue	string	`json:"parameter_value"`
}

// LOpbUOpE is a LOpbUOpE
type LOpbUOpE struct {
	Parameters	[]dKupdOMe	`json:"parameters"`
	Type	string	`json:"type"`
}

// TKSmVoiG is a TKSmVoiG
type TKSmVoiG struct {
	Enabled	bool	`json:"enabled"`
	Processors	[]LOpbUOpE	`json:"processors"`
}

// iFQGZsnw is a iFQGZsnw
type iFQGZsnw struct {
	DomainArn	string	`json:"domain_arn"`
	RetryDuration	int	`json:"retry_duration"`
	S3BackupMode	string	`json:"s3_backup_mode"`
	TypeName	string	`json:"type_name"`
	ProcessingConfiguration	[]TKSmVoiG	`json:"processing_configuration"`
	BufferingInterval	int	`json:"buffering_interval"`
	BufferingSize	int	`json:"buffering_size"`
	IndexName	string	`json:"index_name"`
	IndexRotationPeriod	string	`json:"index_rotation_period"`
	RoleArn	string	`json:"role_arn"`
}
