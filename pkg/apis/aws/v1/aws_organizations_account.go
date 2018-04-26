
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOrganizationsAccount describes a AwsOrganizationsAccount resource
type AwsOrganizationsAccount struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOrganizationsAccountSpec	`json:"spec"`
}


// AwsOrganizationsAccountSpec is the spec for a AwsOrganizationsAccount Resource
type AwsOrganizationsAccountSpec struct {
	IamUserAccessToBilling	string	`json:"iam_user_access_to_billing"`
	RoleName	string	`json:"role_name"`
	Name	string	`json:"name"`
	Email	string	`json:"email"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOrganizationsAccountList is a list of AwsOrganizationsAccount resources
type AwsOrganizationsAccountList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOrganizationsAccount	`json:"items"`
}

