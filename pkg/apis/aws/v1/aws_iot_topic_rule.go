
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
	Kinesis	string	`json:"kinesis"`
	Description	string	`json:"description"`
	Enabled	bool	`json:"enabled"`
	Dynamodb	string	`json:"dynamodb"`
	Elasticsearch	string	`json:"elasticsearch"`
	Firehose	string	`json:"firehose"`
	Name	string	`json:"name"`
	SqlVersion	string	`json:"sql_version"`
	CloudwatchAlarm	string	`json:"cloudwatch_alarm"`
	CloudwatchMetric	string	`json:"cloudwatch_metric"`
	Republish	string	`json:"republish"`
	Sql	string	`json:"sql"`
	Lambda	string	`json:"lambda"`
	S3	string	`json:"s3"`
	Sns	string	`json:"sns"`
	Sqs	string	`json:"sqs"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotTopicRuleList is a list of AwsIotTopicRule resources
type AwsIotTopicRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotTopicRule	`json:"items"`
}

