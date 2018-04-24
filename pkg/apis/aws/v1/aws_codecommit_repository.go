
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodecommitRepository struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCodecommitRepositorySpec	`json:"spec"`
}

type AwsCodecommitRepositoryList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodecommitRepository	`json:"items"`
}

type AwsCodecommitRepositorySpec struct {
	DefaultBranch	string	`json:"default_branch"`
	RepositoryName	string	`json:"repository_name"`
	Description	string	`json:"description"`
}
