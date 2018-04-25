
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesDomainIdentity describes a AwsSesDomainIdentity resource
type AwsSesDomainIdentity struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesDomainIdentitySpec	`json:"spec"`
}


// AwsSesDomainIdentitySpec is the spec for a AwsSesDomainIdentity Resource
type AwsSesDomainIdentitySpec struct {
	Domain	string	`json:"domain"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesDomainIdentityList is a list of AwsSesDomainIdentity resources
type AwsSesDomainIdentityList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesDomainIdentity	`json:"items"`
}

