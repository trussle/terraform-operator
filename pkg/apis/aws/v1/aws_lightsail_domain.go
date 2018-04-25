
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailDomain describes a AwsLightsailDomain resource
type AwsLightsailDomain struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLightsailDomainSpec	`json:"spec"`
}


// AwsLightsailDomainSpec is the spec for a AwsLightsailDomain Resource
type AwsLightsailDomainSpec struct {
	DomainName	string	`json:"domain_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailDomainList is a list of AwsLightsailDomain resources
type AwsLightsailDomainList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLightsailDomain	`json:"items"`
}

