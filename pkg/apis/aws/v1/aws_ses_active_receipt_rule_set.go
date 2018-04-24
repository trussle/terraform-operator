
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesActiveReceiptRuleSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesActiveReceiptRuleSetSpec	`json:"spec"`
}

type AwsSesActiveReceiptRuleSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesActiveReceiptRuleSet	`json:"items"`
}

type AwsSesActiveReceiptRuleSetSpec struct {
	RuleSetName	string	`json:"rule_set_name"`
}
