
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesActiveReceiptRuleSet describes a AwsSesActiveReceiptRuleSet resource
type AwsSesActiveReceiptRuleSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesActiveReceiptRuleSetSpec	`json:"spec"`
}


// AwsSesActiveReceiptRuleSetSpec is the spec for a AwsSesActiveReceiptRuleSet Resource
type AwsSesActiveReceiptRuleSetSpec struct {
	RuleSetName	string	`json:"rule_set_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesActiveReceiptRuleSetList is a list of AwsSesActiveReceiptRuleSet resources
type AwsSesActiveReceiptRuleSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesActiveReceiptRuleSet	`json:"items"`
}

