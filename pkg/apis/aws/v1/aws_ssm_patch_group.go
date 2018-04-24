
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmPatchGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmPatchGroupSpec	`json:"spec"`
}

type AwsSsmPatchGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmPatchGroup	`json:"items"`
}

type AwsSsmPatchGroupSpec struct {
	BaselineId	string	`json:"baseline_id"`
	PatchGroup	string	`json:"patch_group"`
}
