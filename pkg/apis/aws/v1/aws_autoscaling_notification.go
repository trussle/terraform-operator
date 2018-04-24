
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingNotification struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAutoscalingNotificationSpec	`json:"spec"`
}

type AwsAutoscalingNotificationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingNotification	`json:"items"`
}

type AwsAutoscalingNotificationSpec struct {
	TopicArn	string	`json:"topic_arn"`
	GroupNames	interface{}	`json:"group_names"`
	Notifications	interface{}	`json:"notifications"`
}
