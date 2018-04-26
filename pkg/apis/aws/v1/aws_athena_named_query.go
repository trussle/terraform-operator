
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAthenaNamedQuery describes a AwsAthenaNamedQuery resource
type AwsAthenaNamedQuery struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAthenaNamedQuerySpec	`json:"spec"`
}


// AwsAthenaNamedQuerySpec is the spec for a AwsAthenaNamedQuery Resource
type AwsAthenaNamedQuerySpec struct {
	Description	string	`json:"description"`
	Name	string	`json:"name"`
	Query	string	`json:"query"`
	Database	string	`json:"database"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAthenaNamedQueryList is a list of AwsAthenaNamedQuery resources
type AwsAthenaNamedQueryList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAthenaNamedQuery	`json:"items"`
}

