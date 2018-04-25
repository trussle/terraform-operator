
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	VpcId	string	`json:"vpc_id"`
	ServiceName	string	`json:"service_name"`
	VpcEndpointType	string	`json:"vpc_endpoint_type"`
	PrivateDnsEnabled	bool	`json:"private_dns_enabled"`
	AutoAccept	bool	`json:"auto_accept"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointList is a list of AwsVpcEndpoint resources
type AwsVpcEndpointList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpoint	`json:"items"`
}

