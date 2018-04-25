
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnsTopic describes a AwsSnsTopic resource
type AwsSnsTopic struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSnsTopicSpec	`json:"spec"`
}


// AwsSnsTopicSpec is the spec for a AwsSnsTopic Resource
type AwsSnsTopicSpec struct {
	ApplicationSuccessFeedbackSampleRate	int	`json:"application_success_feedback_sample_rate"`
	HttpFailureFeedbackRoleArn	string	`json:"http_failure_feedback_role_arn"`
	LambdaSuccessFeedbackSampleRate	int	`json:"lambda_success_feedback_sample_rate"`
	SqsSuccessFeedbackSampleRate	int	`json:"sqs_success_feedback_sample_rate"`
	DeliveryPolicy	string	`json:"delivery_policy"`
	ApplicationFailureFeedbackRoleArn	string	`json:"application_failure_feedback_role_arn"`
	LambdaFailureFeedbackRoleArn	string	`json:"lambda_failure_feedback_role_arn"`
	NamePrefix	string	`json:"name_prefix"`
	DisplayName	string	`json:"display_name"`
	HttpSuccessFeedbackSampleRate	int	`json:"http_success_feedback_sample_rate"`
	LambdaSuccessFeedbackRoleArn	string	`json:"lambda_success_feedback_role_arn"`
	SqsSuccessFeedbackRoleArn	string	`json:"sqs_success_feedback_role_arn"`
	ApplicationSuccessFeedbackRoleArn	string	`json:"application_success_feedback_role_arn"`
	HttpSuccessFeedbackRoleArn	string	`json:"http_success_feedback_role_arn"`
	SqsFailureFeedbackRoleArn	string	`json:"sqs_failure_feedback_role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnsTopicList is a list of AwsSnsTopic resources
type AwsSnsTopicList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSnsTopic	`json:"items"`
}

