
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmPatchBaseline struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmPatchBaselineSpec	`json:"spec"`
}

type AwsSsmPatchBaselineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmPatchBaseline	`json:"items"`
}

type AwsSsmPatchBaselineSpec struct {
	OperatingSystem	string	`json:"operating_system"`
	ApprovedPatchesComplianceLevel	string	`json:"approved_patches_compliance_level"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	GlobalFilter	[]interface{}	`json:"global_filter"`
	ApprovalRule	[]interface{}	`json:"approval_rule"`
	ApprovedPatches	interface{}	`json:"approved_patches"`
	RejectedPatches	interface{}	`json:"rejected_patches"`
}
