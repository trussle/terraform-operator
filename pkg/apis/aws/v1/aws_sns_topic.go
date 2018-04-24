
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsTopic struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSnsTopicSpec	`json:"spec"`
}

type AwsSnsTopicList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSnsTopic	`json:"items"`
}

type AwsSnsTopicSpec struct {
	ApplicationSuccessFeedbackRoleArn	string	`json:"application_success_feedback_role_arn"`
	LambdaFailureFeedbackRoleArn	string	`json:"lambda_failure_feedback_role_arn"`
	NamePrefix	string	`json:"name_prefix"`
	DisplayName	string	`json:"display_name"`
	ApplicationSuccessFeedbackSampleRate	int	`json:"application_success_feedback_sample_rate"`
	ApplicationFailureFeedbackRoleArn	string	`json:"application_failure_feedback_role_arn"`
	HttpSuccessFeedbackRoleArn	string	`json:"http_success_feedback_role_arn"`
	LambdaSuccessFeedbackRoleArn	string	`json:"lambda_success_feedback_role_arn"`
	SqsFailureFeedbackRoleArn	string	`json:"sqs_failure_feedback_role_arn"`
	HttpSuccessFeedbackSampleRate	int	`json:"http_success_feedback_sample_rate"`
	LambdaSuccessFeedbackSampleRate	int	`json:"lambda_success_feedback_sample_rate"`
	SqsSuccessFeedbackRoleArn	string	`json:"sqs_success_feedback_role_arn"`
	DeliveryPolicy	string	`json:"delivery_policy"`
	HttpFailureFeedbackRoleArn	string	`json:"http_failure_feedback_role_arn"`
	SqsSuccessFeedbackSampleRate	int	`json:"sqs_success_feedback_sample_rate"`
}
