
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAthenaNamedQuery struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAthenaNamedQuerySpec	`json:"spec"`
}

type AwsAthenaNamedQueryList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAthenaNamedQuery	`json:"items"`
}

type AwsAthenaNamedQuerySpec struct {
	Name	string	`json:"name"`
	Query	string	`json:"query"`
	Database	string	`json:"database"`
	Description	string	`json:"description"`
}