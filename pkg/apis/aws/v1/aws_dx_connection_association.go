
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxConnectionAssociation describes a AwsDxConnectionAssociation resource
type AwsDxConnectionAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDxConnectionAssociationSpec	`json:"spec"`
}


// AwsDxConnectionAssociationSpec is the spec for a AwsDxConnectionAssociation Resource
type AwsDxConnectionAssociationSpec struct {
	ConnectionId	string	`json:"connection_id"`
	LagId	string	`json:"lag_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxConnectionAssociationList is a list of AwsDxConnectionAssociation resources
type AwsDxConnectionAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDxConnectionAssociation	`json:"items"`
}

