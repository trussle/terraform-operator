
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRuleGroup describes a AwsWafregionalRuleGroup resource
type AwsWafregionalRuleGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalRuleGroupSpec	`json:"spec"`
}


// AwsWafregionalRuleGroupSpec is the spec for a AwsWafregionalRuleGroup Resource
type AwsWafregionalRuleGroupSpec struct {
	Name	string	`json:"name"`
	MetricName	string	`json:"metric_name"`
	ActivatedRule	AwsWafregionalRuleGroupActivatedRule	`json:"activated_rule"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRuleGroupList is a list of AwsWafregionalRuleGroup resources
type AwsWafregionalRuleGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalRuleGroup	`json:"items"`
}


// AwsWafregionalRuleGroupActivatedRule is a AwsWafregionalRuleGroupActivatedRule
type AwsWafregionalRuleGroupActivatedRule struct {
	Action	[]AwsWafregionalRuleGroupAction	`json:"action"`
	Priority	int	`json:"priority"`
	RuleId	string	`json:"rule_id"`
	Type	string	`json:"type"`
}

// AwsWafregionalRuleGroupAction is a AwsWafregionalRuleGroupAction
type AwsWafregionalRuleGroupAction struct {
	Type	string	`json:"type"`
}
