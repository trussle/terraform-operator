
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbTargetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbTargetGroupSpec	`json:"spec"`
}

type AwsLbTargetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbTargetGroup	`json:"items"`
}

type AwsLbTargetGroupSpec struct {
	VpcId	string	`json:"vpc_id"`
	TargetType	string	`json:"target_type"`
	Tags	map[string]interface{}	`json:"tags"`
	NamePrefix	string	`json:"name_prefix"`
	Port	int	`json:"port"`
	Protocol	string	`json:"protocol"`
	DeregistrationDelay	int	`json:"deregistration_delay"`
}
