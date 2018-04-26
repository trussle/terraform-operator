
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRule describes a AwsWafRule resource
type AwsWafRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafRuleSpec	`json:"spec"`
}


// AwsWafRuleSpec is the spec for a AwsWafRule Resource
type AwsWafRuleSpec struct {
	Predicates	Generic	`json:"predicates"`
	Name	string	`json:"name"`
	MetricName	string	`json:"metric_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRuleList is a list of AwsWafRule resources
type AwsWafRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafRule	`json:"items"`
}

