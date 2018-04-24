
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityGroupRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSecurityGroupRuleSpec	`json:"spec"`
}

type AwsSecurityGroupRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSecurityGroupRule	`json:"items"`
}

type AwsSecurityGroupRuleSpec struct {
	FromPort	int	`json:"from_port"`
	ToPort	int	`json:"to_port"`
	Protocol	string	`json:"protocol"`
	Ipv6CidrBlocks	[]interface{}	`json:"ipv6_cidr_blocks"`
	SecurityGroupId	string	`json:"security_group_id"`
	Self	bool	`json:"self"`
	Description	string	`json:"description"`
	Type	string	`json:"type"`
	PrefixListIds	[]interface{}	`json:"prefix_list_ids"`
	CidrBlocks	[]interface{}	`json:"cidr_blocks"`
}
