
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingScheduledAction describes a AwsAppautoscalingScheduledAction resource
type AwsAppautoscalingScheduledAction struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAppautoscalingScheduledActionSpec	`json:"spec"`
}


// AwsAppautoscalingScheduledActionSpec is the spec for a AwsAppautoscalingScheduledAction Resource
type AwsAppautoscalingScheduledActionSpec struct {
	ServiceNamespace	string	`json:"service_namespace"`
	ScalableDimension	string	`json:"scalable_dimension"`
	StartTime	string	`json:"start_time"`
	Name	string	`json:"name"`
	ScalableTargetAction	[]interface{}	`json:"scalable_target_action"`
	Schedule	string	`json:"schedule"`
	EndTime	string	`json:"end_time"`
	ResourceId	string	`json:"resource_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingScheduledActionList is a list of AwsAppautoscalingScheduledAction resources
type AwsAppautoscalingScheduledActionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppautoscalingScheduledAction	`json:"items"`
}

