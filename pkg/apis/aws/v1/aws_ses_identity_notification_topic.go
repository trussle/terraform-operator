
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesIdentityNotificationTopic struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesIdentityNotificationTopicSpec	`json:"spec"`
}

type AwsSesIdentityNotificationTopicList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesIdentityNotificationTopic	`json:"items"`
}

type AwsSesIdentityNotificationTopicSpec struct {
	TopicArn	string	`json:"topic_arn"`
	NotificationType	string	`json:"notification_type"`
	Identity	string	`json:"identity"`
}
