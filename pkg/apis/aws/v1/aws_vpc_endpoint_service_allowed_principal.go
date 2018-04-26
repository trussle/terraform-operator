
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointServiceAllowedPrincipal describes a AwsVpcEndpointServiceAllowedPrincipal resource
type AwsVpcEndpointServiceAllowedPrincipal struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcEndpointServiceAllowedPrincipalSpec	`json:"spec"`
}


// AwsVpcEndpointServiceAllowedPrincipalSpec is the spec for a AwsVpcEndpointServiceAllowedPrincipal Resource
type AwsVpcEndpointServiceAllowedPrincipalSpec struct {
	VpcEndpointServiceId	string	`json:"vpc_endpoint_service_id"`
	PrincipalArn	string	`json:"principal_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointServiceAllowedPrincipalList is a list of AwsVpcEndpointServiceAllowedPrincipal resources
type AwsVpcEndpointServiceAllowedPrincipalList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpointServiceAllowedPrincipal	`json:"items"`
}

