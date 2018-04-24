
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcPeeringConnectionAccepter struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcPeeringConnectionAccepterSpec	`json:"spec"`
}

type AwsVpcPeeringConnectionAccepterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcPeeringConnectionAccepter	`json:"items"`
}

type AwsVpcPeeringConnectionAccepterSpec struct {
	VpcPeeringConnectionId	string	`json:"vpc_peering_connection_id"`
	AutoAccept	bool	`json:"auto_accept"`
	Tags	map[string]interface{}	`json:"tags"`
}
