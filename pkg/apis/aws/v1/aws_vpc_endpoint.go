
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpoint struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcEndpointSpec	`json:"spec"`
}

type AwsVpcEndpointList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpoint	`json:"items"`
}

type AwsVpcEndpointSpec struct {
	ServiceName	string	`json:"service_name"`
	AutoAccept	bool	`json:"auto_accept"`
	VpcEndpointType	string	`json:"vpc_endpoint_type"`
	VpcId	string	`json:"vpc_id"`
	PrivateDnsEnabled	bool	`json:"private_dns_enabled"`
}
