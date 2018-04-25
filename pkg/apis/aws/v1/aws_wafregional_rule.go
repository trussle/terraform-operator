
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRule describes a AwsWafregionalRule resource
type AwsWafregionalRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalRuleSpec	`json:"spec"`
}


// AwsWafregionalRuleSpec is the spec for a AwsWafregionalRule Resource
type AwsWafregionalRuleSpec struct {
	Name	string	`json:"name"`
	MetricName	string	`json:"metric_name"`
	Predicate	string	`json:"predicate"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRuleList is a list of AwsWafregionalRule resources
type AwsWafregionalRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalRule	`json:"items"`
}

