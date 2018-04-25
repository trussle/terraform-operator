
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRuleGroup describes a AwsWafRuleGroup resource
type AwsWafRuleGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafRuleGroupSpec	`json:"spec"`
}


// AwsWafRuleGroupSpec is the spec for a AwsWafRuleGroup Resource
type AwsWafRuleGroupSpec struct {
	Name	string	`json:"name"`
	MetricName	string	`json:"metric_name"`
	ActivatedRule	string	`json:"activated_rule"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRuleGroupList is a list of AwsWafRuleGroup resources
type AwsWafRuleGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafRuleGroup	`json:"items"`
}

