
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53ZoneAssociation describes a AwsRoute53ZoneAssociation resource
type AwsRoute53ZoneAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRoute53ZoneAssociationSpec	`json:"spec"`
}


// AwsRoute53ZoneAssociationSpec is the spec for a AwsRoute53ZoneAssociation Resource
type AwsRoute53ZoneAssociationSpec struct {
	VpcId	string	`json:"vpc_id"`
	ZoneId	string	`json:"zone_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53ZoneAssociationList is a list of AwsRoute53ZoneAssociation resources
type AwsRoute53ZoneAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53ZoneAssociation	`json:"items"`
}

