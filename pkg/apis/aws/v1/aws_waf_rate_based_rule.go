
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRateBasedRule describes a AwsWafRateBasedRule resource
type AwsWafRateBasedRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafRateBasedRuleSpec	`json:"spec"`
}


// AwsWafRateBasedRuleSpec is the spec for a AwsWafRateBasedRule Resource
type AwsWafRateBasedRuleSpec struct {
	RateKey	string	`json:"rate_key"`
	RateLimit	int	`json:"rate_limit"`
	Name	string	`json:"name"`
	MetricName	string	`json:"metric_name"`
	Predicates	string	`json:"predicates"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRateBasedRuleList is a list of AwsWafRateBasedRule resources
type AwsWafRateBasedRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafRateBasedRule	`json:"items"`
}

