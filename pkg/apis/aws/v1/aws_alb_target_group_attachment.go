
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbTargetGroupAttachment describes a AwsAlbTargetGroupAttachment resource
type AwsAlbTargetGroupAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbTargetGroupAttachmentSpec	`json:"spec"`
}


// AwsAlbTargetGroupAttachmentSpec is the spec for a AwsAlbTargetGroupAttachment Resource
type AwsAlbTargetGroupAttachmentSpec struct {
	AvailabilityZone	string	`json:"availability_zone"`
	TargetGroupArn	string	`json:"target_group_arn"`
	TargetId	string	`json:"target_id"`
	Port	int	`json:"port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbTargetGroupAttachmentList is a list of AwsAlbTargetGroupAttachment resources
type AwsAlbTargetGroupAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbTargetGroupAttachment	`json:"items"`
}

