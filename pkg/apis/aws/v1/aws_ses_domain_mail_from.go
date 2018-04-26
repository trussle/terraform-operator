
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesDomainMailFrom describes a AwsSesDomainMailFrom resource
type AwsSesDomainMailFrom struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesDomainMailFromSpec	`json:"spec"`
}


// AwsSesDomainMailFromSpec is the spec for a AwsSesDomainMailFrom Resource
type AwsSesDomainMailFromSpec struct {
	BehaviorOnMxFailure	string	`json:"behavior_on_mx_failure"`
	Domain	string	`json:"domain"`
	MailFromDomain	string	`json:"mail_from_domain"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesDomainMailFromList is a list of AwsSesDomainMailFrom resources
type AwsSesDomainMailFromList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesDomainMailFrom	`json:"items"`
}

