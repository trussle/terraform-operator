
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingSchedule describes a AwsAutoscalingSchedule resource
type AwsAutoscalingSchedule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAutoscalingScheduleSpec	`json:"spec"`
}


// AwsAutoscalingScheduleSpec is the spec for a AwsAutoscalingSchedule Resource
type AwsAutoscalingScheduleSpec struct {
	ScheduledActionName	string	`json:"scheduled_action_name"`
	AutoscalingGroupName	string	`json:"autoscaling_group_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingScheduleList is a list of AwsAutoscalingSchedule resources
type AwsAutoscalingScheduleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingSchedule	`json:"items"`
}

