
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccountPasswordPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamAccountPasswordPolicySpec	`json:"spec"`
}

type AwsIamAccountPasswordPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamAccountPasswordPolicy	`json:"items"`
}

type AwsIamAccountPasswordPolicySpec struct {
	AllowUsersToChangePassword	bool	`json:"allow_users_to_change_password"`
	MinimumPasswordLength	int	`json:"minimum_password_length"`
}
