
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointService describes a AwsVpcEndpointService resource
type AwsVpcEndpointService struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcEndpointServiceSpec	`json:"spec"`
}


// AwsVpcEndpointServiceSpec is the spec for a AwsVpcEndpointService Resource
type AwsVpcEndpointServiceSpec struct {
	AcceptanceRequired	bool	`json:"acceptance_required"`
	NetworkLoadBalancerArns	string	`json:"network_load_balancer_arns"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointServiceList is a list of AwsVpcEndpointService resources
type AwsVpcEndpointServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpointService	`json:"items"`
}

