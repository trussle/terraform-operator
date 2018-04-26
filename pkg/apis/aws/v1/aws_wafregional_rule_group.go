
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	MetricName	string	`json:"metric_name"`
	ActivatedRule	Generic	`json:"activated_rule"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRuleGroupList is a list of AwsWafregionalRuleGroup resources
type AwsWafregionalRuleGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalRuleGroup	`json:"items"`
}

