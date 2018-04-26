
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmPatchGroup describes a AwsSsmPatchGroup resource
type AwsSsmPatchGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmPatchGroupSpec	`json:"spec"`
}


// AwsSsmPatchGroupSpec is the spec for a AwsSsmPatchGroup Resource
type AwsSsmPatchGroupSpec struct {
	BaselineId	string	`json:"baseline_id"`
	PatchGroup	string	`json:"patch_group"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmPatchGroupList is a list of AwsSsmPatchGroup resources
type AwsSsmPatchGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmPatchGroup	`json:"items"`
}

