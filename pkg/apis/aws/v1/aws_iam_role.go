
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRole struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamRoleSpec	`json:"spec"`
}

type AwsIamRoleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamRole	`json:"items"`
}

type AwsIamRoleSpec struct {
	ForceDetachPolicies	bool	`json:"force_detach_policies"`
	Description	string	`json:"description"`
	AssumeRolePolicy	string	`json:"assume_role_policy"`
	MaxSessionDuration	int	`json:"max_session_duration"`
	NamePrefix	string	`json:"name_prefix"`
	Path	string	`json:"path"`
}
