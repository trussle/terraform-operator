
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamRolePolicy describes a AwsIamRolePolicy resource
type AwsIamRolePolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamRolePolicySpec	`json:"spec"`
}


// AwsIamRolePolicySpec is the spec for a AwsIamRolePolicy Resource
type AwsIamRolePolicySpec struct {
	Policy	string	`json:"policy"`
	NamePrefix	string	`json:"name_prefix"`
	Role	string	`json:"role"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamRolePolicyList is a list of AwsIamRolePolicy resources
type AwsIamRolePolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamRolePolicy	`json:"items"`
}

