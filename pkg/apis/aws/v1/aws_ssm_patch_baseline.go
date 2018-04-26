
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
	RejectedPatches	string	`json:"rejected_patches"`
	OperatingSystem	string	`json:"operating_system"`
	ApprovedPatchesComplianceLevel	string	`json:"approved_patches_compliance_level"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	GlobalFilter	[]cjiOgjhY	`json:"global_filter"`
	ApprovalRule	[]eVwBTCML	`json:"approval_rule"`
	ApprovedPatches	string	`json:"approved_patches"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmPatchBaselineList is a list of AwsSsmPatchBaseline resources
type AwsSsmPatchBaselineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmPatchBaseline	`json:"items"`
}


// cjiOgjhY is a cjiOgjhY
type cjiOgjhY struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}

// frDGXqwp is a frDGXqwp
type frDGXqwp struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}

// eVwBTCML is a eVwBTCML
type eVwBTCML struct {
	ComplianceLevel	string	`json:"compliance_level"`
	PatchFilter	[]frDGXqwp	`json:"patch_filter"`
	ApproveAfterDays	int	`json:"approve_after_days"`
}
