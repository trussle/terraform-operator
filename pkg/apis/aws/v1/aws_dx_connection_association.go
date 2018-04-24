
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxConnectionAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDxConnectionAssociationSpec	`json:"spec"`
}

type AwsDxConnectionAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDxConnectionAssociation	`json:"items"`
}

type AwsDxConnectionAssociationSpec struct {
	ConnectionId	string	`json:"connection_id"`
	LagId	string	`json:"lag_id"`
}
