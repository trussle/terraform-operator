
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEipAssociation describes a AwsEipAssociation resource
type AwsEipAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEipAssociationSpec	`json:"spec"`
}


// AwsEipAssociationSpec is the spec for a AwsEipAssociation Resource
type AwsEipAssociationSpec struct {
	AllowReassociation	bool	`json:"allow_reassociation"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEipAssociationList is a list of AwsEipAssociation resources
type AwsEipAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEipAssociation	`json:"items"`
}

