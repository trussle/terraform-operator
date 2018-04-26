
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmPatchBaseline describes a AwsSsmPatchBaseline resource
type AwsSsmPatchBaseline struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmPatchBaselineSpec	`json:"spec"`
}


// AwsSsmPatchBaselineSpec is the spec for a AwsSsmPatchBaseline Resource
type AwsSsmPatchBaselineSpec struct {
	ApprovedPatchesComplianceLevel	string	`json:"approved_patches_compliance_level"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	GlobalFilter	[]Generic	`json:"global_filter"`
	ApprovalRule	[]Generic	`json:"approval_rule"`
	ApprovedPatches	Generic	`json:"approved_patches"`
	RejectedPatches	Generic	`json:"rejected_patches"`
	OperatingSystem	string	`json:"operating_system"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmPatchBaselineList is a list of AwsSsmPatchBaseline resources
type AwsSsmPatchBaselineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmPatchBaseline	`json:"items"`
}

