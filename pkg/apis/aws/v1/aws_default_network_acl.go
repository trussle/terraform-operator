
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
	DefaultNetworkAclId	string	`json:"default_network_acl_id"`
	SubnetIds	string	`json:"subnet_ids"`
	Ingress	Ingress	`json:"ingress"`
	Egress	Egress	`json:"egress"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultNetworkAclList is a list of AwsDefaultNetworkAcl resources
type AwsDefaultNetworkAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultNetworkAcl	`json:"items"`
}


// Ingress is a Ingress
type Ingress struct {
	IcmpType	int	`json:"icmp_type"`
	ToPort	int	`json:"to_port"`
	Protocol	string	`json:"protocol"`
	Action	string	`json:"action"`
	CidrBlock	string	`json:"cidr_block"`
	Ipv6CidrBlock	string	`json:"ipv6_cidr_block"`
	IcmpCode	int	`json:"icmp_code"`
	FromPort	int	`json:"from_port"`
	RuleNo	int	`json:"rule_no"`
}

// Egress is a Egress
type Egress struct {
	Protocol	string	`json:"protocol"`
	FromPort	int	`json:"from_port"`
	ToPort	int	`json:"to_port"`
	RuleNo	int	`json:"rule_no"`
	IcmpType	int	`json:"icmp_type"`
	IcmpCode	int	`json:"icmp_code"`
	Action	string	`json:"action"`
	CidrBlock	string	`json:"cidr_block"`
	Ipv6CidrBlock	string	`json:"ipv6_cidr_block"`
}
