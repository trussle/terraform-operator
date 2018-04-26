
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesReceiptRuleSet describes a AwsSesReceiptRuleSet resource
type AwsSesReceiptRuleSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesReceiptRuleSetSpec	`json:"spec"`
}


// AwsSesReceiptRuleSetSpec is the spec for a AwsSesReceiptRuleSet Resource
type AwsSesReceiptRuleSetSpec struct {
	RuleSetName	string	`json:"rule_set_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesReceiptRuleSetList is a list of AwsSesReceiptRuleSet resources
type AwsSesReceiptRuleSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesReceiptRuleSet	`json:"items"`
}

