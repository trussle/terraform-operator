
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEfsFileSystem describes a AwsEfsFileSystem resource
type AwsEfsFileSystem struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEfsFileSystemSpec	`json:"spec"`
}


// AwsEfsFileSystemSpec is the spec for a AwsEfsFileSystem Resource
type AwsEfsFileSystemSpec struct {
	Tags	map[string]???	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEfsFileSystemList is a list of AwsEfsFileSystem resources
type AwsEfsFileSystemList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEfsFileSystem	`json:"items"`
}

