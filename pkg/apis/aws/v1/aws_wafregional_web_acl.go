
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalWebAcl describes a AwsWafregionalWebAcl resource
type AwsWafregionalWebAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalWebAclSpec	`json:"spec"`
}


// AwsWafregionalWebAclSpec is the spec for a AwsWafregionalWebAcl Resource
type AwsWafregionalWebAclSpec struct {
	MetricName	string	`json:"metric_name"`
	Rule	Rule	`json:"rule"`
	Name	string	`json:"name"`
	DefaultAction	[]DefaultAction	`json:"default_action"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalWebAclList is a list of AwsWafregionalWebAcl resources
type AwsWafregionalWebAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalWebAcl	`json:"items"`
}


// Action is a Action
type Action struct {
	Type	string	`json:"type"`
}

// Rule is a Rule
type Rule struct {
	Action	[]Action	`json:"action"`
	Priority	int	`json:"priority"`
	RuleId	string	`json:"rule_id"`
}

// DefaultAction is a DefaultAction
type DefaultAction struct {
	Type	string	`json:"type"`
}
