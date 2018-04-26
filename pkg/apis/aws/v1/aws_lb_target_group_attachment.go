
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbTargetGroupAttachment describes a AwsLbTargetGroupAttachment resource
type AwsLbTargetGroupAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbTargetGroupAttachmentSpec	`json:"spec"`
}


// AwsLbTargetGroupAttachmentSpec is the spec for a AwsLbTargetGroupAttachment Resource
type AwsLbTargetGroupAttachmentSpec struct {
	TargetGroupArn	string	`json:"target_group_arn"`
	TargetId	string	`json:"target_id"`
	Port	int	`json:"port"`
	AvailabilityZone	string	`json:"availability_zone"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbTargetGroupAttachmentList is a list of AwsLbTargetGroupAttachment resources
type AwsLbTargetGroupAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbTargetGroupAttachment	`json:"items"`
}

