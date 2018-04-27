
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnsTopicSubscription describes a AwsSnsTopicSubscription resource
type AwsSnsTopicSubscription struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSnsTopicSubscriptionSpec	`json:"spec"`
}


// AwsSnsTopicSubscriptionSpec is the spec for a AwsSnsTopicSubscription Resource
type AwsSnsTopicSubscriptionSpec struct {
	Endpoint	string	`json:"endpoint"`
	EndpointAutoConfirms	bool	`json:"endpoint_auto_confirms"`
	ConfirmationTimeoutInMinutes	int	`json:"confirmation_timeout_in_minutes"`
	FilterPolicy	string	`json:"filter_policy"`
	Protocol	string	`json:"protocol"`
	TopicArn	string	`json:"topic_arn"`
	DeliveryPolicy	string	`json:"delivery_policy"`
	RawMessageDelivery	bool	`json:"raw_message_delivery"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnsTopicSubscriptionList is a list of AwsSnsTopicSubscription resources
type AwsSnsTopicSubscriptionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSnsTopicSubscription	`json:"items"`
}

