
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSnsTopicSubscription struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSnsTopicSubscriptionSpec	`json:"spec"`
}

type AwsSnsTopicSubscriptionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSnsTopicSubscription	`json:"items"`
}

type AwsSnsTopicSubscriptionSpec struct {
	TopicArn	string	`json:"topic_arn"`
	DeliveryPolicy	string	`json:"delivery_policy"`
	FilterPolicy	string	`json:"filter_policy"`
	Protocol	string	`json:"protocol"`
	Endpoint	string	`json:"endpoint"`
	EndpointAutoConfirms	bool	`json:"endpoint_auto_confirms"`
	ConfirmationTimeoutInMinutes	int	`json:"confirmation_timeout_in_minutes"`
	RawMessageDelivery	bool	`json:"raw_message_delivery"`
}
