
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
	Tags	map[string]string	`json:"tags"`
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


// AwsVpcPeeringConnectionRequester is a AwsVpcPeeringConnectionRequester
type AwsVpcPeeringConnectionRequester struct {
	AllowRemoteVpcDnsResolution	bool	`json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc	bool	`json:"allow_classic_link_to_remote_vpc"`
	AllowVpcToRemoteClassicLink	bool	`json:"allow_vpc_to_remote_classic_link"`
}

// AwsVpcPeeringConnectionAccepter is a AwsVpcPeeringConnectionAccepter
type AwsVpcPeeringConnectionAccepter struct {
	AllowVpcToRemoteClassicLink	bool	`json:"allow_vpc_to_remote_classic_link"`
	AllowRemoteVpcDnsResolution	bool	`json:"allow_remote_vpc_dns_resolution"`
	AllowClassicLinkToRemoteVpc	bool	`json:"allow_classic_link_to_remote_vpc"`
}
