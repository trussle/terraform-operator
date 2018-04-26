
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesDomainIdentityVerification describes a AwsSesDomainIdentityVerification resource
type AwsSesDomainIdentityVerification struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesDomainIdentityVerificationSpec	`json:"spec"`
}


// AwsSesDomainIdentityVerificationSpec is the spec for a AwsSesDomainIdentityVerification Resource
type AwsSesDomainIdentityVerificationSpec struct {
	Domain	string	`json:"domain"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesDomainIdentityVerificationList is a list of AwsSesDomainIdentityVerification resources
type AwsSesDomainIdentityVerificationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesDomainIdentityVerification	`json:"items"`
}

