
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamAccountPasswordPolicy describes a AwsIamAccountPasswordPolicy resource
type AwsIamAccountPasswordPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamAccountPasswordPolicySpec	`json:"spec"`
}


// AwsIamAccountPasswordPolicySpec is the spec for a AwsIamAccountPasswordPolicy Resource
type AwsIamAccountPasswordPolicySpec struct {
	AllowUsersToChangePassword	bool	`json:"allow_users_to_change_password"`
	MinimumPasswordLength	int	`json:"minimum_password_length"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamAccountPasswordPolicyList is a list of AwsIamAccountPasswordPolicy resources
type AwsIamAccountPasswordPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamAccountPasswordPolicy	`json:"items"`
}

