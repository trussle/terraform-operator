
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcrRepository describes a AwsEcrRepository resource
type AwsEcrRepository struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEcrRepositorySpec	`json:"spec"`
}


// AwsEcrRepositorySpec is the spec for a AwsEcrRepository Resource
type AwsEcrRepositorySpec struct {
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcrRepositoryList is a list of AwsEcrRepository resources
type AwsEcrRepositoryList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcrRepository	`json:"items"`
}

