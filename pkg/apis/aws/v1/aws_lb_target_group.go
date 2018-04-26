
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbTargetGroup describes a AwsLbTargetGroup resource
type AwsLbTargetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbTargetGroupSpec	`json:"spec"`
}


// AwsLbTargetGroupSpec is the spec for a AwsLbTargetGroup Resource
type AwsLbTargetGroupSpec struct {
	NamePrefix	string	`json:"name_prefix"`
	Port	int	`json:"port"`
	Tags	map[string]Generic	`json:"tags"`
	TargetType	string	`json:"target_type"`
	Protocol	string	`json:"protocol"`
	VpcId	string	`json:"vpc_id"`
	DeregistrationDelay	int	`json:"deregistration_delay"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbTargetGroupList is a list of AwsLbTargetGroup resources
type AwsLbTargetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbTargetGroup	`json:"items"`
}

