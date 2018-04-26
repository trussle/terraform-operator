
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcPeeringConnectionAccepter describes a AwsVpcPeeringConnectionAccepter resource
type AwsVpcPeeringConnectionAccepter struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcPeeringConnectionAccepterSpec	`json:"spec"`
}


// AwsVpcPeeringConnectionAccepterSpec is the spec for a AwsVpcPeeringConnectionAccepter Resource
type AwsVpcPeeringConnectionAccepterSpec struct {
	Tags	map[string]string	`json:"tags"`
	VpcPeeringConnectionId	string	`json:"vpc_peering_connection_id"`
	AutoAccept	bool	`json:"auto_accept"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcPeeringConnectionAccepterList is a list of AwsVpcPeeringConnectionAccepter resources
type AwsVpcPeeringConnectionAccepterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcPeeringConnectionAccepter	`json:"items"`
}


// Requester is a Requester
type Requester struct {
	AllowRemoteVpcDnsResolution	bool	`json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc	bool	`json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink	bool	`json:"allow_vpc_to_remote_classic_link"`
}

// Accepter is a Accepter
type Accepter struct {
	AllowVpcToRemoteClassicLink	bool	`json:"allow_vpc_to_remote_classic_link"`
	AllowRemoteVpcDnsResolution	bool	`json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc	bool	`json:"allow_classic_link_to_remote_vpc"`
}
