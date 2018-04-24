
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRateBasedRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafRateBasedRuleSpec	`json:"spec"`
}

type AwsWafRateBasedRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafRateBasedRule	`json:"items"`
}

type AwsWafRateBasedRuleSpec struct {
	Predicates	interface{}	`json:"predicates"`
	RateKey	string	`json:"rate_key"`
	RateLimit	int	`json:"rate_limit"`
	Name	string	`json:"name"`
	MetricName	string	`json:"metric_name"`
}
