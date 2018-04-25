
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesDomainDkim describes a AwsSesDomainDkim resource
type AwsSesDomainDkim struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesDomainDkimSpec	`json:"spec"`
}


// AwsSesDomainDkimSpec is the spec for a AwsSesDomainDkim Resource
type AwsSesDomainDkimSpec struct {
	Domain	string	`json:"domain"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesDomainDkimList is a list of AwsSesDomainDkim resources
type AwsSesDomainDkimList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesDomainDkim	`json:"items"`
}

