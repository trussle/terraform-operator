
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSimpledbDomain describes a AwsSimpledbDomain resource
type AwsSimpledbDomain struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSimpledbDomainSpec	`json:"spec"`
}


// AwsSimpledbDomainSpec is the spec for a AwsSimpledbDomain Resource
type AwsSimpledbDomainSpec struct {
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSimpledbDomainList is a list of AwsSimpledbDomain resources
type AwsSimpledbDomainList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSimpledbDomain	`json:"items"`
}

