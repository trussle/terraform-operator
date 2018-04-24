
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafSqlInjectionMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafSqlInjectionMatchSetSpec	`json:"spec"`
}

type AwsWafSqlInjectionMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafSqlInjectionMatchSet	`json:"items"`
}

type AwsWafSqlInjectionMatchSetSpec struct {
	Name	string	`json:"name"`
	SqlInjectionMatchTuples	interface{}	`json:"sql_injection_match_tuples"`
}
