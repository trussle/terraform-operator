
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafWebAcl describes a AwsWafWebAcl resource
type AwsWafWebAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafWebAclSpec	`json:"spec"`
}


// AwsWafWebAclSpec is the spec for a AwsWafWebAcl Resource
type AwsWafWebAclSpec struct {
	DefaultAction	AwsWafWebAclDefaultAction	`json:"default_action"`
	MetricName	string	`json:"metric_name"`
	Rules	AwsWafWebAclRules	`json:"rules"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafWebAclList is a list of AwsWafWebAcl resources
type AwsWafWebAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafWebAcl	`json:"items"`
}


// AwsWafWebAclDefaultAction is a AwsWafWebAclDefaultAction
type AwsWafWebAclDefaultAction struct {
	Type	string	`json:"type"`
}

// AwsWafWebAclAction is a AwsWafWebAclAction
type AwsWafWebAclAction struct {
	Type	string	`json:"type"`
}

// AwsWafWebAclRules is a AwsWafWebAclRules
type AwsWafWebAclRules struct {
	RuleId	string	`json:"rule_id"`
	Action	AwsWafWebAclAction	`json:"action"`
	Priority	int	`json:"priority"`
	Type	string	`json:"type"`
}
