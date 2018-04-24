
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointService struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcEndpointServiceSpec	`json:"spec"`
}

type AwsVpcEndpointServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpointService	`json:"items"`
}

type AwsVpcEndpointServiceSpec struct {
	AcceptanceRequired	bool	`json:"acceptance_required"`
	NetworkLoadBalancerArns	interface{}	`json:"network_load_balancer_arns"`
}
