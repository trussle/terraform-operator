
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkAcl describes a AwsNetworkAcl resource
type AwsNetworkAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNetworkAclSpec	`json:"spec"`
}


// AwsNetworkAclSpec is the spec for a AwsNetworkAcl Resource
type AwsNetworkAclSpec struct {
	Tags	map[string]string	`json:"tags"`
	VpcId	string	`json:"vpc_id"`
	SubnetId	string	`json:"subnet_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkAclList is a list of AwsNetworkAcl resources
type AwsNetworkAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNetworkAcl	`json:"items"`
}


// Egress is a Egress
type Egress struct {
	FromPort	int	`json:"from_port"`
	RuleNo	int	`json:"rule_no"`
	CidrBlock	string	`json:"cidr_block"`
	Ipv6CidrBlock	string	`json:"ipv6_cidr_block"`
	IcmpType	int	`json:"icmp_type"`
	IcmpCode	int	`json:"icmp_code"`
	ToPort	int	`json:"to_port"`
	Action	string	`json:"action"`
	Protocol	string	`json:"protocol"`
}

// Ingress is a Ingress
type Ingress struct {
	RuleNo	int	`json:"rule_no"`
	Action	string	`json:"action"`
	Protocol	string	`json:"protocol"`
	FromPort	int	`json:"from_port"`
	ToPort	int	`json:"to_port"`
	CidrBlock	string	`json:"cidr_block"`
	Ipv6CidrBlock	string	`json:"ipv6_cidr_block"`
	IcmpType	int	`json:"icmp_type"`
	IcmpCode	int	`json:"icmp_code"`
}
