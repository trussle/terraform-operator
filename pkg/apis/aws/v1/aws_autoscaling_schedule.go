
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingSchedule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAutoscalingScheduleSpec	`json:"spec"`
}

type AwsAutoscalingScheduleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingSchedule	`json:"items"`
}

type AwsAutoscalingScheduleSpec struct {
	ScheduledActionName	string	`json:"scheduled_action_name"`
	AutoscalingGroupName	string	`json:"autoscaling_group_name"`
}
