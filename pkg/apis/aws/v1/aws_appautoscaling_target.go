
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingTarget struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAppautoscalingTargetSpec	`json:"spec"`
}

type AwsAppautoscalingTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppautoscalingTarget	`json:"items"`
}

type AwsAppautoscalingTargetSpec struct {
	ScalableDimension	string	`json:"scalable_dimension"`
	ServiceNamespace	string	`json:"service_namespace"`
	MaxCapacity	int	`json:"max_capacity"`
	MinCapacity	int	`json:"min_capacity"`
	ResourceId	string	`json:"resource_id"`
}
