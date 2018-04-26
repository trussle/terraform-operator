
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamInstanceProfile describes a AwsIamInstanceProfile resource
type AwsIamInstanceProfile struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamInstanceProfileSpec	`json:"spec"`
}


// AwsIamInstanceProfileSpec is the spec for a AwsIamInstanceProfile Resource
type AwsIamInstanceProfileSpec struct {
	NamePrefix	string	`json:"name_prefix"`
	Path	string	`json:"path"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamInstanceProfileList is a list of AwsIamInstanceProfile resources
type AwsIamInstanceProfileList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamInstanceProfile	`json:"items"`
}

