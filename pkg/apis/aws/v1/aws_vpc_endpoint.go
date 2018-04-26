
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpoint describes a AwsVpcEndpoint resource
type AwsVpcEndpoint struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcEndpointSpec	`json:"spec"`
}


// AwsVpcEndpointSpec is the spec for a AwsVpcEndpoint Resource
type AwsVpcEndpointSpec struct {
	VpcEndpointType	string	`json:"vpc_endpoint_type"`
	ServiceName	string	`json:"service_name"`
	PrivateDnsEnabled	bool	`json:"private_dns_enabled"`
	VpcId	string	`json:"vpc_id"`
	AutoAccept	bool	`json:"auto_accept"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointList is a list of AwsVpcEndpoint resources
type AwsVpcEndpointList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpoint	`json:"items"`
}


// MJjxzuai is a MJjxzuai
type MJjxzuai struct {
}
