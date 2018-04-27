
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultNetworkAcl describes a AwsDefaultNetworkAcl resource
type AwsDefaultNetworkAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultNetworkAclSpec	`json:"spec"`
}


// AwsDefaultNetworkAclSpec is the spec for a AwsDefaultNetworkAcl Resource
type AwsDefaultNetworkAclSpec struct {
	Ingress	AwsDefaultNetworkAclIngress	`json:"ingress"`
	Egress	AwsDefaultNetworkAclEgress	`json:"egress"`
	Tags	map[string]string	`json:"tags"`
	DefaultNetworkAclId	string	`json:"default_network_acl_id"`
	SubnetIds	string	`json:"subnet_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultNetworkAclList is a list of AwsDefaultNetworkAcl resources
type AwsDefaultNetworkAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultNetworkAcl	`json:"items"`
}


// AwsDefaultNetworkAclIngress is a AwsDefaultNetworkAclIngress
type AwsDefaultNetworkAclIngress struct {
	Action	string	`json:"action"`
	Protocol	string	`json:"protocol"`
	Ipv6CidrBlock	string	`json:"ipv6_cidr_block"`
	IcmpType	int	`json:"icmp_type"`
	RuleNo	int	`json:"rule_no"`
	ToPort	int	`json:"to_port"`
	CidrBlock	string	`json:"cidr_block"`
	IcmpCode	int	`json:"icmp_code"`
	FromPort	int	`json:"from_port"`
}

// AwsDefaultNetworkAclEgress is a AwsDefaultNetworkAclEgress
type AwsDefaultNetworkAclEgress struct {
	RuleNo	int	`json:"rule_no"`
	Action	string	`json:"action"`
	Ipv6CidrBlock	string	`json:"ipv6_cidr_block"`
	IcmpType	int	`json:"icmp_type"`
	IcmpCode	int	`json:"icmp_code"`
	ToPort	int	`json:"to_port"`
	Protocol	string	`json:"protocol"`
	CidrBlock	string	`json:"cidr_block"`
	FromPort	int	`json:"from_port"`
}
