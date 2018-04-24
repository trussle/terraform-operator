
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOrganizationsAccount struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOrganizationsAccountSpec	`json:"spec"`
}

type AwsOrganizationsAccountList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOrganizationsAccount	`json:"items"`
}

type AwsOrganizationsAccountSpec struct {
	Name	string	`json:"name"`
	Email	string	`json:"email"`
	IamUserAccessToBilling	string	`json:"iam_user_access_to_billing"`
	RoleName	string	`json:"role_name"`
}
