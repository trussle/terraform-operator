
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingLifecycleHook struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAutoscalingLifecycleHookSpec	`json:"spec"`
}

type AwsAutoscalingLifecycleHookList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingLifecycleHook	`json:"items"`
}

type AwsAutoscalingLifecycleHookSpec struct {
	Name	string	`json:"name"`
	AutoscalingGroupName	string	`json:"autoscaling_group_name"`
	HeartbeatTimeout	int	`json:"heartbeat_timeout"`
	LifecycleTransition	string	`json:"lifecycle_transition"`
	NotificationMetadata	string	`json:"notification_metadata"`
	NotificationTargetArn	string	`json:"notification_target_arn"`
	RoleArn	string	`json:"role_arn"`
}
