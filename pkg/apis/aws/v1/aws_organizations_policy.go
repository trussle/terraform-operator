
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOrganizationsPolicy describes a AwsOrganizationsPolicy resource
type AwsOrganizationsPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOrganizationsPolicySpec	`json:"spec"`
}


// AwsOrganizationsPolicySpec is the spec for a AwsOrganizationsPolicy Resource
type AwsOrganizationsPolicySpec struct {
	Content	string	`json:"content"`
	Description	string	`json:"description"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOrganizationsPolicyList is a list of AwsOrganizationsPolicy resources
type AwsOrganizationsPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOrganizationsPolicy	`json:"items"`
}

