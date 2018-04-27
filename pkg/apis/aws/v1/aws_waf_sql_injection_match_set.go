
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafSqlInjectionMatchSet describes a AwsWafSqlInjectionMatchSet resource
type AwsWafSqlInjectionMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafSqlInjectionMatchSetSpec	`json:"spec"`
}


// AwsWafSqlInjectionMatchSetSpec is the spec for a AwsWafSqlInjectionMatchSet Resource
type AwsWafSqlInjectionMatchSetSpec struct {
	SqlInjectionMatchTuples	AwsWafSqlInjectionMatchSetSqlInjectionMatchTuples	`json:"sql_injection_match_tuples"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafSqlInjectionMatchSetList is a list of AwsWafSqlInjectionMatchSet resources
type AwsWafSqlInjectionMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafSqlInjectionMatchSet	`json:"items"`
}


// AwsWafSqlInjectionMatchSetFieldToMatch is a AwsWafSqlInjectionMatchSetFieldToMatch
type AwsWafSqlInjectionMatchSetFieldToMatch struct {
	Data	string	`json:"data"`
	Type	string	`json:"type"`
}

// AwsWafSqlInjectionMatchSetSqlInjectionMatchTuples is a AwsWafSqlInjectionMatchSetSqlInjectionMatchTuples
type AwsWafSqlInjectionMatchSetSqlInjectionMatchTuples struct {
	FieldToMatch	AwsWafSqlInjectionMatchSetFieldToMatch	`json:"field_to_match"`
	TextTransformation	string	`json:"text_transformation"`
}
