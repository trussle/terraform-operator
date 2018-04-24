
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingScheduledAction struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAppautoscalingScheduledActionSpec	`json:"spec"`
}

type AwsAppautoscalingScheduledActionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppautoscalingScheduledAction	`json:"items"`
}

type AwsAppautoscalingScheduledActionSpec struct {
	Schedule	string	`json:"schedule"`
	EndTime	string	`json:"end_time"`
	ScalableTargetAction	[]interface{}	`json:"scalable_target_action"`
	ServiceNamespace	string	`json:"service_namespace"`
	ResourceId	string	`json:"resource_id"`
	ScalableDimension	string	`json:"scalable_dimension"`
	StartTime	string	`json:"start_time"`
	Name	string	`json:"name"`
}
