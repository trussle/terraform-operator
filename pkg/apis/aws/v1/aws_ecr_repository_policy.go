
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcrRepositoryPolicy describes a AwsEcrRepositoryPolicy resource
type AwsEcrRepositoryPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEcrRepositoryPolicySpec	`json:"spec"`
}


// AwsEcrRepositoryPolicySpec is the spec for a AwsEcrRepositoryPolicy Resource
type AwsEcrRepositoryPolicySpec struct {
	Repository	string	`json:"repository"`
	Policy	string	`json:"policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcrRepositoryPolicyList is a list of AwsEcrRepositoryPolicy resources
type AwsEcrRepositoryPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcrRepositoryPolicy	`json:"items"`
}

