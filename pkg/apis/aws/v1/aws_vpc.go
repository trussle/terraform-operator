
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpc struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcSpec	`json:"spec"`
}

type AwsVpcList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpc	`json:"items"`
}

type AwsVpcSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	CidrBlock	string	`json:"cidr_block"`
	EnableDnsSupport	bool	`json:"enable_dns_support"`
	AssignGeneratedIpv6CidrBlock	bool	`json:"assign_generated_ipv6_cidr_block"`
}
