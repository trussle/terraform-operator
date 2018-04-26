
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Self	bool	`json:"self"`
	Type	string	`json:"type"`
	Protocol	string	`json:"protocol"`
	PrefixListIds	[]string	`json:"prefix_list_ids"`
	SecurityGroupId	string	`json:"security_group_id"`
	FromPort	int	`json:"from_port"`
	ToPort	int	`json:"to_port"`
	CidrBlocks	[]string	`json:"cidr_blocks"`
	Ipv6CidrBlocks	[]string	`json:"ipv6_cidr_blocks"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecurityGroupRuleList is a list of AwsSecurityGroupRule resources
type AwsSecurityGroupRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSecurityGroupRule	`json:"items"`
}

