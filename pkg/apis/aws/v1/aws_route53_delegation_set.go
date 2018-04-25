
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53DelegationSet describes a AwsRoute53DelegationSet resource
type AwsRoute53DelegationSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRoute53DelegationSetSpec	`json:"spec"`
}


// AwsRoute53DelegationSetSpec is the spec for a AwsRoute53DelegationSet Resource
type AwsRoute53DelegationSetSpec struct {
	ReferenceName	string	`json:"reference_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53DelegationSetList is a list of AwsRoute53DelegationSet resources
type AwsRoute53DelegationSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53DelegationSet	`json:"items"`
}

