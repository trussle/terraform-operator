
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkAclRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNetworkAclRuleSpec	`json:"spec"`
}

type AwsNetworkAclRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNetworkAclRule	`json:"items"`
}

type AwsNetworkAclRuleSpec struct {
	RuleNumber	int	`json:"rule_number"`
	Ipv6CidrBlock	string	`json:"ipv6_cidr_block"`
	ToPort	int	`json:"to_port"`
	CidrBlock	string	`json:"cidr_block"`
	FromPort	int	`json:"from_port"`
	IcmpType	string	`json:"icmp_type"`
	IcmpCode	string	`json:"icmp_code"`
	NetworkAclId	string	`json:"network_acl_id"`
	Egress	bool	`json:"egress"`
	Protocol	string	`json:"protocol"`
	RuleAction	string	`json:"rule_action"`
}
