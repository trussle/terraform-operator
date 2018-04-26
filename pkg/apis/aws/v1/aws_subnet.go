
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSubnet describes a AwsSubnet resource
type AwsSubnet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSubnetSpec	`json:"spec"`
}


// AwsSubnetSpec is the spec for a AwsSubnet Resource
type AwsSubnetSpec struct {
	Tags	map[string]???	`json:"tags"`
	VpcId	string	`json:"vpc_id"`
	CidrBlock	string	`json:"cidr_block"`
	MapPublicIpOnLaunch	bool	`json:"map_public_ip_on_launch"`
	AssignIpv6AddressOnCreation	bool	`json:"assign_ipv6_address_on_creation"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSubnetList is a list of AwsSubnet resources
type AwsSubnetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSubnet	`json:"items"`
}

