
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpc describes a AwsVpc resource
type AwsVpc struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcSpec	`json:"spec"`
}


// AwsVpcSpec is the spec for a AwsVpc Resource
type AwsVpcSpec struct {
	CidrBlock	string	`json:"cidr_block"`
	AssignGeneratedIpv6CidrBlock	bool	`json:"assign_generated_ipv6_cidr_block"`
	EnableDnsSupport	bool	`json:"enable_dns_support"`
	Tags	map[string]???	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcList is a list of AwsVpc resources
type AwsVpcList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpc	`json:"items"`
}

