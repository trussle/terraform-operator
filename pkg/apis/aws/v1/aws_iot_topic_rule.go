
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
	Lambda	AwsIotTopicRuleLambda	`json:"lambda"`
	Enabled	bool	`json:"enabled"`
	CloudwatchAlarm	AwsIotTopicRuleCloudwatchAlarm	`json:"cloudwatch_alarm"`
	CloudwatchMetric	AwsIotTopicRuleCloudwatchMetric	`json:"cloudwatch_metric"`
	Firehose	AwsIotTopicRuleFirehose	`json:"firehose"`
	SqlVersion	string	`json:"sql_version"`
	Kinesis	AwsIotTopicRuleKinesis	`json:"kinesis"`
	S3	AwsIotTopicRuleS3	`json:"s3"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Sql	string	`json:"sql"`
	Republish	AwsIotTopicRuleRepublish	`json:"republish"`
	Dynamodb	AwsIotTopicRuleDynamodb	`json:"dynamodb"`
	Elasticsearch	AwsIotTopicRuleElasticsearch	`json:"elasticsearch"`
	Sns	AwsIotTopicRuleSns	`json:"sns"`
	Sqs	AwsIotTopicRuleSqs	`json:"sqs"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotTopicRuleList is a list of AwsIotTopicRule resources
type AwsIotTopicRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotTopicRule	`json:"items"`
}


// AwsIotTopicRuleFirehose is a AwsIotTopicRuleFirehose
type AwsIotTopicRuleFirehose struct {
	DeliveryStreamName	string	`json:"delivery_stream_name"`
	RoleArn	string	`json:"role_arn"`
}

// AwsIotTopicRuleKinesis is a AwsIotTopicRuleKinesis
type AwsIotTopicRuleKinesis struct {
	StreamName	string	`json:"stream_name"`
	PartitionKey	string	`json:"partition_key"`
	RoleArn	string	`json:"role_arn"`
}

// AwsIotTopicRuleS3 is a AwsIotTopicRuleS3
type AwsIotTopicRuleS3 struct {
	BucketName	string	`json:"bucket_name"`
	Key	string	`json:"key"`
	RoleArn	string	`json:"role_arn"`
}

// AwsIotTopicRuleRepublish is a AwsIotTopicRuleRepublish
type AwsIotTopicRuleRepublish struct {
	RoleArn	string	`json:"role_arn"`
	Topic	string	`json:"topic"`
}

// AwsIotTopicRuleElasticsearch is a AwsIotTopicRuleElasticsearch
type AwsIotTopicRuleElasticsearch struct {
	Endpoint	string	`json:"endpoint"`
	Id	string	`json:"id"`
	Index	string	`json:"index"`
	RoleArn	string	`json:"role_arn"`
	Type	string	`json:"type"`
}

// AwsIotTopicRuleSns is a AwsIotTopicRuleSns
type AwsIotTopicRuleSns struct {
	RoleArn	string	`json:"role_arn"`
	MessageFormat	string	`json:"message_format"`
	TargetArn	string	`json:"target_arn"`
}

// AwsIotTopicRuleLambda is a AwsIotTopicRuleLambda
type AwsIotTopicRuleLambda struct {
	FunctionArn	string	`json:"function_arn"`
}

// AwsIotTopicRuleCloudwatchMetric is a AwsIotTopicRuleCloudwatchMetric
type AwsIotTopicRuleCloudwatchMetric struct {
	MetricName	string	`json:"metric_name"`
	MetricNamespace	string	`json:"metric_namespace"`
	MetricTimestamp	string	`json:"metric_timestamp"`
	MetricUnit	string	`json:"metric_unit"`
	MetricValue	string	`json:"metric_value"`
	RoleArn	string	`json:"role_arn"`
}

// AwsIotTopicRuleSqs is a AwsIotTopicRuleSqs
type AwsIotTopicRuleSqs struct {
	UseBase64	bool	`json:"use_base64"`
	QueueUrl	string	`json:"queue_url"`
	RoleArn	string	`json:"role_arn"`
}

// AwsIotTopicRuleCloudwatchAlarm is a AwsIotTopicRuleCloudwatchAlarm
type AwsIotTopicRuleCloudwatchAlarm struct {
	AlarmName	string	`json:"alarm_name"`
	RoleArn	string	`json:"role_arn"`
	StateReason	string	`json:"state_reason"`
	StateValue	string	`json:"state_value"`
}

// AwsIotTopicRuleDynamodb is a AwsIotTopicRuleDynamodb
type AwsIotTopicRuleDynamodb struct {
	RoleArn	string	`json:"role_arn"`
	TableName	string	`json:"table_name"`
	HashKeyValue	string	`json:"hash_key_value"`
	HashKeyType	string	`json:"hash_key_type"`
	RangeKeyValue	string	`json:"range_key_value"`
	RangeKeyType	string	`json:"range_key_type"`
	HashKeyField	string	`json:"hash_key_field"`
	PayloadField	string	`json:"payload_field"`
	RangeKeyField	string	`json:"range_key_field"`
}
