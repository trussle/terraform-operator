
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	GlobalFilter	[]GlobalFilter	`json:"global_filter"`
	ApprovalRule	[]ApprovalRule	`json:"approval_rule"`
	ApprovedPatches	string	`json:"approved_patches"`
	RejectedPatches	string	`json:"rejected_patches"`
	OperatingSystem	string	`json:"operating_system"`
	ApprovedPatchesComplianceLevel	string	`json:"approved_patches_compliance_level"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmPatchBaselineList is a list of AwsSsmPatchBaseline resources
type AwsSsmPatchBaselineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmPatchBaseline	`json:"items"`
}


// GlobalFilter is a GlobalFilter
type GlobalFilter struct {
	Values	[]string	`json:"values"`
	Key	string	`json:"key"`
}

// PatchFilter is a PatchFilter
type PatchFilter struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}

// ApprovalRule is a ApprovalRule
type ApprovalRule struct {
	ApproveAfterDays	int	`json:"approve_after_days"`
	ComplianceLevel	string	`json:"compliance_level"`
	PatchFilter	[]PatchFilter	`json:"patch_filter"`
}
