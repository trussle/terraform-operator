
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiCopy describes a AwsAmiCopy resource
type AwsAmiCopy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAmiCopySpec	`json:"spec"`
}


// AwsAmiCopySpec is the spec for a AwsAmiCopy Resource
type AwsAmiCopySpec struct {
	Encrypted	bool	`json:"encrypted"`
	Description	string	`json:"description"`
	SourceAmiRegion	string	`json:"source_ami_region"`
	Name	string	`json:"name"`
	Tags	map[string]string	`json:"tags"`
	SourceAmiId	string	`json:"source_ami_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiCopyList is a list of AwsAmiCopy resources
type AwsAmiCopyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAmiCopy	`json:"items"`
}


// EbsBlockDevice is a EbsBlockDevice
type EbsBlockDevice struct {
}

// EphemeralBlockDevice is a EphemeralBlockDevice
type EphemeralBlockDevice struct {
}
