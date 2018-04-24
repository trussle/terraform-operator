
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbTargetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbTargetGroupSpec	`json:"spec"`
}

type AwsAlbTargetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbTargetGroup	`json:"items"`
}

type AwsAlbTargetGroupSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	NamePrefix	string	`json:"name_prefix"`
	VpcId	string	`json:"vpc_id"`
	TargetType	string	`json:"target_type"`
	Port	int	`json:"port"`
	Protocol	string	`json:"protocol"`
	DeregistrationDelay	int	`json:"deregistration_delay"`
}
