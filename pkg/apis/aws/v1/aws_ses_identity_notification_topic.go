
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesIdentityNotificationTopic describes a AwsSesIdentityNotificationTopic resource
type AwsSesIdentityNotificationTopic struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesIdentityNotificationTopicSpec	`json:"spec"`
}


// AwsSesIdentityNotificationTopicSpec is the spec for a AwsSesIdentityNotificationTopic Resource
type AwsSesIdentityNotificationTopicSpec struct {
	Identity	string	`json:"identity"`
	TopicArn	string	`json:"topic_arn"`
	NotificationType	string	`json:"notification_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesIdentityNotificationTopicList is a list of AwsSesIdentityNotificationTopic resources
type AwsSesIdentityNotificationTopicList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesIdentityNotificationTopic	`json:"items"`
}

