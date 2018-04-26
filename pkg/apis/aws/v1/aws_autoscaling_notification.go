
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingNotification describes a AwsAutoscalingNotification resource
type AwsAutoscalingNotification struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAutoscalingNotificationSpec	`json:"spec"`
}


// AwsAutoscalingNotificationSpec is the spec for a AwsAutoscalingNotification Resource
type AwsAutoscalingNotificationSpec struct {
	TopicArn	string	`json:"topic_arn"`
	GroupNames	string	`json:"group_names"`
	Notifications	string	`json:"notifications"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingNotificationList is a list of AwsAutoscalingNotification resources
type AwsAutoscalingNotificationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingNotification	`json:"items"`
}

