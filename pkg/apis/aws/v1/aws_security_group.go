
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecurityGroup describes a AwsSecurityGroup resource
type AwsSecurityGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSecurityGroupSpec	`json:"spec"`
}


// AwsSecurityGroupSpec is the spec for a AwsSecurityGroup Resource
type AwsSecurityGroupSpec struct {
	RevokeRulesOnDelete	bool	`json:"revoke_rules_on_delete"`
	NamePrefix	string	`json:"name_prefix"`
	Description	string	`json:"description"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecurityGroupList is a list of AwsSecurityGroup resources
type AwsSecurityGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSecurityGroup	`json:"items"`
}


// Ingress is a Ingress
type Ingress struct {
	Self	bool	`json:"self"`
	Description	string	`json:"description"`
	FromPort	int	`json:"from_port"`
	ToPort	int	`json:"to_port"`
	Protocol	string	`json:"protocol"`
	CidrBlocks	[]string	`json:"cidr_blocks"`
	Ipv6CidrBlocks	[]string	`json:"ipv6_cidr_blocks"`
	SecurityGroups	string	`json:"security_groups"`
}

// Egress is a Egress
type Egress struct {
	SecurityGroups	string	`json:"security_groups"`
	Self	bool	`json:"self"`
	Protocol	string	`json:"protocol"`
	Ipv6CidrBlocks	[]string	`json:"ipv6_cidr_blocks"`
	CidrBlocks	[]string	`json:"cidr_blocks"`
	PrefixListIds	[]string	`json:"prefix_list_ids"`
	Description	string	`json:"description"`
	FromPort	int	`json:"from_port"`
	ToPort	int	`json:"to_port"`
}
