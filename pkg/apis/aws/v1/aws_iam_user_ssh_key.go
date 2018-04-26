
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserSshKey describes a AwsIamUserSshKey resource
type AwsIamUserSshKey struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamUserSshKeySpec	`json:"spec"`
}


// AwsIamUserSshKeySpec is the spec for a AwsIamUserSshKey Resource
type AwsIamUserSshKeySpec struct {
	Username	string	`json:"username"`
	PublicKey	string	`json:"public_key"`
	Encoding	string	`json:"encoding"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserSshKeyList is a list of AwsIamUserSshKey resources
type AwsIamUserSshKeyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamUserSshKey	`json:"items"`
}

