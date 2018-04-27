
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
	Name	string	`json:"name"`
	DefaultAction	[]AwsWafregionalWebAclDefaultAction	`json:"default_action"`
	MetricName	string	`json:"metric_name"`
	Rule	AwsWafregionalWebAclRule	`json:"rule"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalWebAclList is a list of AwsWafregionalWebAcl resources
type AwsWafregionalWebAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalWebAcl	`json:"items"`
}


// AwsWafregionalWebAclDefaultAction is a AwsWafregionalWebAclDefaultAction
type AwsWafregionalWebAclDefaultAction struct {
	Type	string	`json:"type"`
}

// AwsWafregionalWebAclAction is a AwsWafregionalWebAclAction
type AwsWafregionalWebAclAction struct {
	Type	string	`json:"type"`
}

// AwsWafregionalWebAclRule is a AwsWafregionalWebAclRule
type AwsWafregionalWebAclRule struct {
	Action	[]AwsWafregionalWebAclAction	`json:"action"`
	Priority	int	`json:"priority"`
	RuleId	string	`json:"rule_id"`
}
