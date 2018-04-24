
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsPlatformApplication struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSnsPlatformApplicationSpec	`json:"spec"`
}

type AwsSnsPlatformApplicationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSnsPlatformApplication	`json:"items"`
}

type AwsSnsPlatformApplicationSpec struct {
	Name	string	`json:"name"`
	Platform	string	`json:"platform"`
	PlatformCredential	string	`json:"platform_credential"`
	FailureFeedbackRoleArn	string	`json:"failure_feedback_role_arn"`
	PlatformPrincipal	string	`json:"platform_principal"`
	SuccessFeedbackRoleArn	string	`json:"success_feedback_role_arn"`
	EventDeliveryFailureTopicArn	string	`json:"event_delivery_failure_topic_arn"`
	EventEndpointCreatedTopicArn	string	`json:"event_endpoint_created_topic_arn"`
	EventEndpointDeletedTopicArn	string	`json:"event_endpoint_deleted_topic_arn"`
	EventEndpointUpdatedTopicArn	string	`json:"event_endpoint_updated_topic_arn"`
	SuccessFeedbackSampleRate	string	`json:"success_feedback_sample_rate"`
}
