
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcPeeringConnection struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcPeeringConnectionSpec	`json:"spec"`
}

type AwsVpcPeeringConnectionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcPeeringConnection	`json:"items"`
}

type AwsVpcPeeringConnectionSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	PeerVpcId	string	`json:"peer_vpc_id"`
	VpcId	string	`json:"vpc_id"`
	AutoAccept	bool	`json:"auto_accept"`
}
