
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElbAttachment describes a AwsElbAttachment resource
type AwsElbAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElbAttachmentSpec	`json:"spec"`
}


// AwsElbAttachmentSpec is the spec for a AwsElbAttachment Resource
type AwsElbAttachmentSpec struct {
	Elb	string	`json:"elb"`
	Instance	string	`json:"instance"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElbAttachmentList is a list of AwsElbAttachment resources
type AwsElbAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElbAttachment	`json:"items"`
}

