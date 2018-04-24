
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotTopicRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIotTopicRuleSpec	`json:"spec"`
}

type AwsIotTopicRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotTopicRule	`json:"items"`
}

type AwsIotTopicRuleSpec struct {
	Description	string	`json:"description"`
	Sql	string	`json:"sql"`
	CloudwatchAlarm	interface{}	`json:"cloudwatch_alarm"`
	S3	interface{}	`json:"s3"`
	Sns	interface{}	`json:"sns"`
	Sqs	interface{}	`json:"sqs"`
	SqlVersion	string	`json:"sql_version"`
	Firehose	interface{}	`json:"firehose"`
	Republish	interface{}	`json:"republish"`
	Name	string	`json:"name"`
	Enabled	bool	`json:"enabled"`
	Dynamodb	interface{}	`json:"dynamodb"`
	Kinesis	interface{}	`json:"kinesis"`
	CloudwatchMetric	interface{}	`json:"cloudwatch_metric"`
	Elasticsearch	interface{}	`json:"elasticsearch"`
	Lambda	interface{}	`json:"lambda"`
}
