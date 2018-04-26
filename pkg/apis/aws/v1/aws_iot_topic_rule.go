
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotTopicRule describes a AwsIotTopicRule resource
type AwsIotTopicRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIotTopicRuleSpec	`json:"spec"`
}


// AwsIotTopicRuleSpec is the spec for a AwsIotTopicRule Resource
type AwsIotTopicRuleSpec struct {
	Elasticsearch	Elasticsearch	`json:"elasticsearch"`
	Name	string	`json:"name"`
	Lambda	Lambda	`json:"lambda"`
	Sql	string	`json:"sql"`
	Enabled	bool	`json:"enabled"`
	CloudwatchMetric	CloudwatchMetric	`json:"cloudwatch_metric"`
	Dynamodb	Dynamodb	`json:"dynamodb"`
	S3	S3	`json:"s3"`
	Sns	Sns	`json:"sns"`
	Description	string	`json:"description"`
	CloudwatchAlarm	CloudwatchAlarm	`json:"cloudwatch_alarm"`
	Firehose	Firehose	`json:"firehose"`
	Kinesis	Kinesis	`json:"kinesis"`
	Republish	Republish	`json:"republish"`
	Sqs	Sqs	`json:"sqs"`
	SqlVersion	string	`json:"sql_version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotTopicRuleList is a list of AwsIotTopicRule resources
type AwsIotTopicRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotTopicRule	`json:"items"`
}


// Elasticsearch is a Elasticsearch
type Elasticsearch struct {
	Endpoint	string	`json:"endpoint"`
	Id	string	`json:"id"`
	Index	string	`json:"index"`
	RoleArn	string	`json:"role_arn"`
	Type	string	`json:"type"`
}

// CloudwatchAlarm is a CloudwatchAlarm
type CloudwatchAlarm struct {
	AlarmName	string	`json:"alarm_name"`
	RoleArn	string	`json:"role_arn"`
	StateReason	string	`json:"state_reason"`
	StateValue	string	`json:"state_value"`
}

// Republish is a Republish
type Republish struct {
	Topic	string	`json:"topic"`
	RoleArn	string	`json:"role_arn"`
}

// Lambda is a Lambda
type Lambda struct {
	FunctionArn	string	`json:"function_arn"`
}

// CloudwatchMetric is a CloudwatchMetric
type CloudwatchMetric struct {
	MetricValue	string	`json:"metric_value"`
	RoleArn	string	`json:"role_arn"`
	MetricName	string	`json:"metric_name"`
	MetricNamespace	string	`json:"metric_namespace"`
	MetricTimestamp	string	`json:"metric_timestamp"`
	MetricUnit	string	`json:"metric_unit"`
}

// Dynamodb is a Dynamodb
type Dynamodb struct {
	HashKeyField	string	`json:"hash_key_field"`
	PayloadField	string	`json:"payload_field"`
	RangeKeyField	string	`json:"range_key_field"`
	RoleArn	string	`json:"role_arn"`
	TableName	string	`json:"table_name"`
	HashKeyValue	string	`json:"hash_key_value"`
	HashKeyType	string	`json:"hash_key_type"`
	RangeKeyValue	string	`json:"range_key_value"`
	RangeKeyType	string	`json:"range_key_type"`
}

// S3 is a S3
type S3 struct {
	BucketName	string	`json:"bucket_name"`
	Key	string	`json:"key"`
	RoleArn	string	`json:"role_arn"`
}

// Sns is a Sns
type Sns struct {
	MessageFormat	string	`json:"message_format"`
	TargetArn	string	`json:"target_arn"`
	RoleArn	string	`json:"role_arn"`
}

// Firehose is a Firehose
type Firehose struct {
	DeliveryStreamName	string	`json:"delivery_stream_name"`
	RoleArn	string	`json:"role_arn"`
}

// Kinesis is a Kinesis
type Kinesis struct {
	StreamName	string	`json:"stream_name"`
	PartitionKey	string	`json:"partition_key"`
	RoleArn	string	`json:"role_arn"`
}

// Sqs is a Sqs
type Sqs struct {
	QueueUrl	string	`json:"queue_url"`
	RoleArn	string	`json:"role_arn"`
	UseBase64	bool	`json:"use_base64"`
}
