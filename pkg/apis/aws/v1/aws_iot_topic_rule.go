
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	SqlVersion	string	`json:"sql_version"`
	CloudwatchAlarm	Generic	`json:"cloudwatch_alarm"`
	Elasticsearch	Generic	`json:"elasticsearch"`
	Lambda	Generic	`json:"lambda"`
	Sqs	Generic	`json:"sqs"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Republish	Generic	`json:"republish"`
	S3	Generic	`json:"s3"`
	Sns	Generic	`json:"sns"`
	Enabled	bool	`json:"enabled"`
	Sql	string	`json:"sql"`
	Kinesis	Generic	`json:"kinesis"`
	Firehose	Generic	`json:"firehose"`
	CloudwatchMetric	Generic	`json:"cloudwatch_metric"`
	Dynamodb	Generic	`json:"dynamodb"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotTopicRuleList is a list of AwsIotTopicRule resources
type AwsIotTopicRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotTopicRule	`json:"items"`
}

