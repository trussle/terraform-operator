
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamPolicy describes a AwsIamPolicy resource
type AwsIamPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamPolicySpec	`json:"spec"`
}


// AwsIamPolicySpec is the spec for a AwsIamPolicy Resource
type AwsIamPolicySpec struct {
	Policy	string	`json:"policy"`
	NamePrefix	string	`json:"name_prefix"`
	Description	string	`json:"description"`
	Path	string	`json:"path"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamPolicyList is a list of AwsIamPolicy resources
type AwsIamPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamPolicy	`json:"items"`
}

