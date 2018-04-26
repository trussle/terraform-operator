
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcPeeringConnection describes a AwsVpcPeeringConnection resource
type AwsVpcPeeringConnection struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcPeeringConnectionSpec	`json:"spec"`
}


// AwsVpcPeeringConnectionSpec is the spec for a AwsVpcPeeringConnection Resource
type AwsVpcPeeringConnectionSpec struct {
	Tags	map[string]???	`json:"tags"`
	PeerVpcId	string	`json:"peer_vpc_id"`
	AutoAccept	bool	`json:"auto_accept"`
	VpcId	string	`json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcPeeringConnectionList is a list of AwsVpcPeeringConnection resources
type AwsVpcPeeringConnectionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcPeeringConnection	`json:"items"`
}

