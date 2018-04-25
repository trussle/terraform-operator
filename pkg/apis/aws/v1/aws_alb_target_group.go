
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbTargetGroup describes a AwsAlbTargetGroup resource
type AwsAlbTargetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbTargetGroupSpec	`json:"spec"`
}


// AwsAlbTargetGroupSpec is the spec for a AwsAlbTargetGroup Resource
type AwsAlbTargetGroupSpec struct {
	NamePrefix	string	`json:"name_prefix"`
	Protocol	string	`json:"protocol"`
	VpcId	string	`json:"vpc_id"`
	Tags	map[string]interface{}	`json:"tags"`
	Port	int	`json:"port"`
	DeregistrationDelay	int	`json:"deregistration_delay"`
	TargetType	string	`json:"target_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbTargetGroupList is a list of AwsAlbTargetGroup resources
type AwsAlbTargetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbTargetGroup	`json:"items"`
}

