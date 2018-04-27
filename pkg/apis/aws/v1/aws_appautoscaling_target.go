
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingTarget describes a AwsAppautoscalingTarget resource
type AwsAppautoscalingTarget struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAppautoscalingTargetSpec	`json:"spec"`
}


// AwsAppautoscalingTargetSpec is the spec for a AwsAppautoscalingTarget Resource
type AwsAppautoscalingTargetSpec struct {
	MaxCapacity	int	`json:"max_capacity"`
	MinCapacity	int	`json:"min_capacity"`
	ResourceId	string	`json:"resource_id"`
	ScalableDimension	string	`json:"scalable_dimension"`
	ServiceNamespace	string	`json:"service_namespace"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingTargetList is a list of AwsAppautoscalingTarget resources
type AwsAppautoscalingTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppautoscalingTarget	`json:"items"`
}

