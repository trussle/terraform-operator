
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecurityGroupRule describes a AwsSecurityGroupRule resource
type AwsSecurityGroupRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSecurityGroupRuleSpec	`json:"spec"`
}


// AwsSecurityGroupRuleSpec is the spec for a AwsSecurityGroupRule Resource
type AwsSecurityGroupRuleSpec struct {
	Type	string	`json:"type"`
	ToPort	int	`json:"to_port"`
	Protocol	string	`json:"protocol"`
	SecurityGroupId	string	`json:"security_group_id"`
	Self	bool	`json:"self"`
	Description	string	`json:"description"`
	FromPort	int	`json:"from_port"`
	CidrBlocks	[]interface{}	`json:"cidr_blocks"`
	Ipv6CidrBlocks	[]interface{}	`json:"ipv6_cidr_blocks"`
	PrefixListIds	[]interface{}	`json:"prefix_list_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecurityGroupRuleList is a list of AwsSecurityGroupRule resources
type AwsSecurityGroupRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSecurityGroupRule	`json:"items"`
}

