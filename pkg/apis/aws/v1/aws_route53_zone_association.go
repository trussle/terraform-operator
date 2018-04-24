
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53ZoneAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRoute53ZoneAssociationSpec	`json:"spec"`
}

type AwsRoute53ZoneAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53ZoneAssociation	`json:"items"`
}

type AwsRoute53ZoneAssociationSpec struct {
	ZoneId	string	`json:"zone_id"`
	VpcId	string	`json:"vpc_id"`
}
