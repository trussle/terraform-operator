
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	ScalableDimension	string	`json:"scalable_dimension"`
	ScalableTargetAction	[]AwsAppautoscalingScheduledActionScalableTargetAction	`json:"scalable_target_action"`
	ServiceNamespace	string	`json:"service_namespace"`
	ResourceId	string	`json:"resource_id"`
	Schedule	string	`json:"schedule"`
	StartTime	string	`json:"start_time"`
	EndTime	string	`json:"end_time"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingScheduledActionList is a list of AwsAppautoscalingScheduledAction resources
type AwsAppautoscalingScheduledActionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppautoscalingScheduledAction	`json:"items"`
}


// AwsAppautoscalingScheduledActionScalableTargetAction is a AwsAppautoscalingScheduledActionScalableTargetAction
type AwsAppautoscalingScheduledActionScalableTargetAction struct {
	MaxCapacity	int	`json:"max_capacity"`
	MinCapacity	int	`json:"min_capacity"`
}
