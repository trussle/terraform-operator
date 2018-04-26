
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnsPlatformApplication describes a AwsSnsPlatformApplication resource
type AwsSnsPlatformApplication struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSnsPlatformApplicationSpec	`json:"spec"`
}


// AwsSnsPlatformApplicationSpec is the spec for a AwsSnsPlatformApplication Resource
type AwsSnsPlatformApplicationSpec struct {
	Platform	string	`json:"platform"`
	PlatformCredential	string	`json:"platform_credential"`
	EventEndpointCreatedTopicArn	string	`json:"event_endpoint_created_topic_arn"`
	EventEndpointUpdatedTopicArn	string	`json:"event_endpoint_updated_topic_arn"`
	PlatformPrincipal	string	`json:"platform_principal"`
	SuccessFeedbackRoleArn	string	`json:"success_feedback_role_arn"`
	SuccessFeedbackSampleRate	string	`json:"success_feedback_sample_rate"`
	Name	string	`json:"name"`
	EventDeliveryFailureTopicArn	string	`json:"event_delivery_failure_topic_arn"`
	EventEndpointDeletedTopicArn	string	`json:"event_endpoint_deleted_topic_arn"`
	FailureFeedbackRoleArn	string	`json:"failure_feedback_role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnsPlatformApplicationList is a list of AwsSnsPlatformApplication resources
type AwsSnsPlatformApplicationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSnsPlatformApplication	`json:"items"`
}

