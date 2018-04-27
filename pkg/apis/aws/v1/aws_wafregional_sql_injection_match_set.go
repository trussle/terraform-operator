
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalSqlInjectionMatchSet describes a AwsWafregionalSqlInjectionMatchSet resource
type AwsWafregionalSqlInjectionMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalSqlInjectionMatchSetSpec	`json:"spec"`
}


// AwsWafregionalSqlInjectionMatchSetSpec is the spec for a AwsWafregionalSqlInjectionMatchSet Resource
type AwsWafregionalSqlInjectionMatchSetSpec struct {
	Name	string	`json:"name"`
	SqlInjectionMatchTuple	AwsWafregionalSqlInjectionMatchSetSqlInjectionMatchTuple	`json:"sql_injection_match_tuple"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalSqlInjectionMatchSetList is a list of AwsWafregionalSqlInjectionMatchSet resources
type AwsWafregionalSqlInjectionMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalSqlInjectionMatchSet	`json:"items"`
}


// AwsWafregionalSqlInjectionMatchSetFieldToMatch is a AwsWafregionalSqlInjectionMatchSetFieldToMatch
type AwsWafregionalSqlInjectionMatchSetFieldToMatch struct {
	Data	string	`json:"data"`
	Type	string	`json:"type"`
}

// AwsWafregionalSqlInjectionMatchSetSqlInjectionMatchTuple is a AwsWafregionalSqlInjectionMatchSetSqlInjectionMatchTuple
type AwsWafregionalSqlInjectionMatchSetSqlInjectionMatchTuple struct {
	FieldToMatch	[]AwsWafregionalSqlInjectionMatchSetFieldToMatch	`json:"field_to_match"`
	TextTransformation	string	`json:"text_transformation"`
}
