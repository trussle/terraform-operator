
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
	Description	string	`json:"description"`
	Name	string	`json:"name"`
	SourceAmiId	string	`json:"source_ami_id"`
	Encrypted	bool	`json:"encrypted"`
	Tags	map[string]string	`json:"tags"`
	SourceAmiRegion	string	`json:"source_ami_region"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiCopyList is a list of AwsAmiCopy resources
type AwsAmiCopyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAmiCopy	`json:"items"`
}


// AwsAmiCopyEbsBlockDevice is a AwsAmiCopyEbsBlockDevice
type AwsAmiCopyEbsBlockDevice struct {
}

// AwsAmiCopyEphemeralBlockDevice is a AwsAmiCopyEphemeralBlockDevice
type AwsAmiCopyEphemeralBlockDevice struct {
}
