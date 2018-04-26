
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserLoginProfile describes a AwsIamUserLoginProfile resource
type AwsIamUserLoginProfile struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamUserLoginProfileSpec	`json:"spec"`
}


// AwsIamUserLoginProfileSpec is the spec for a AwsIamUserLoginProfile Resource
type AwsIamUserLoginProfileSpec struct {
	User	string	`json:"user"`
	PgpKey	string	`json:"pgp_key"`
	PasswordResetRequired	bool	`json:"password_reset_required"`
	PasswordLength	int	`json:"password_length"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserLoginProfileList is a list of AwsIamUserLoginProfile resources
type AwsIamUserLoginProfileList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamUserLoginProfile	`json:"items"`
}

