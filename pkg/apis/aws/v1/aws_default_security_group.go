
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultSecurityGroup describes a AwsDefaultSecurityGroup resource
type AwsDefaultSecurityGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultSecurityGroupSpec	`json:"spec"`
}


// AwsDefaultSecurityGroupSpec is the spec for a AwsDefaultSecurityGroup Resource
type AwsDefaultSecurityGroupSpec struct {
	Egress	Egress	`json:"egress"`
	Ingress	Ingress	`json:"ingress"`
	Tags	map[string]string	`json:"tags"`
	RevokeRulesOnDelete	bool	`json:"revoke_rules_on_delete"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultSecurityGroupList is a list of AwsDefaultSecurityGroup resources
type AwsDefaultSecurityGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultSecurityGroup	`json:"items"`
}


// Egress is a Egress
type Egress struct {
	Protocol	string	`json:"protocol"`
	CidrBlocks	[]string	`json:"cidr_blocks"`
	Ipv6CidrBlocks	[]string	`json:"ipv6_cidr_blocks"`
	FromPort	int	`json:"from_port"`
	ToPort	int	`json:"to_port"`
	PrefixListIds	[]string	`json:"prefix_list_ids"`
	SecurityGroups	string	`json:"security_groups"`
	Self	bool	`json:"self"`
	Description	string	`json:"description"`
}

// Ingress is a Ingress
type Ingress struct {
	Protocol	string	`json:"protocol"`
	CidrBlocks	[]string	`json:"cidr_blocks"`
	Ipv6CidrBlocks	[]string	`json:"ipv6_cidr_blocks"`
	SecurityGroups	string	`json:"security_groups"`
	Self	bool	`json:"self"`
	Description	string	`json:"description"`
	FromPort	int	`json:"from_port"`
	ToPort	int	`json:"to_port"`
}
