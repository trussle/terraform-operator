
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
	DefaultAction	DefaultAction	`json:"default_action"`
	MetricName	string	`json:"metric_name"`
	Rules	Rules	`json:"rules"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafWebAclList is a list of AwsWafWebAcl resources
type AwsWafWebAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafWebAcl	`json:"items"`
}


// Action is a Action
type Action struct {
	Type	string	`json:"type"`
}

// Rules is a Rules
type Rules struct {
	Action	Action	`json:"action"`
	Priority	int	`json:"priority"`
	Type	string	`json:"type"`
	RuleId	string	`json:"rule_id"`
}

// DefaultAction is a DefaultAction
type DefaultAction struct {
	Type	string	`json:"type"`
}
