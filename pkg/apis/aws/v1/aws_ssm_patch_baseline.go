
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
	Description	string	`json:"description"`
	GlobalFilter	[]AwsSsmPatchBaselineGlobalFilter	`json:"global_filter"`
	ApprovalRule	[]AwsSsmPatchBaselineApprovalRule	`json:"approval_rule"`
	ApprovedPatches	string	`json:"approved_patches"`
	RejectedPatches	string	`json:"rejected_patches"`
	OperatingSystem	string	`json:"operating_system"`
	ApprovedPatchesComplianceLevel	string	`json:"approved_patches_compliance_level"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmPatchBaselineList is a list of AwsSsmPatchBaseline resources
type AwsSsmPatchBaselineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmPatchBaseline	`json:"items"`
}


// AwsSsmPatchBaselineGlobalFilter is a AwsSsmPatchBaselineGlobalFilter
type AwsSsmPatchBaselineGlobalFilter struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}

// AwsSsmPatchBaselinePatchFilter is a AwsSsmPatchBaselinePatchFilter
type AwsSsmPatchBaselinePatchFilter struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}

// AwsSsmPatchBaselineApprovalRule is a AwsSsmPatchBaselineApprovalRule
type AwsSsmPatchBaselineApprovalRule struct {
	ApproveAfterDays	int	`json:"approve_after_days"`
	ComplianceLevel	string	`json:"compliance_level"`
	PatchFilter	[]AwsSsmPatchBaselinePatchFilter	`json:"patch_filter"`
}
