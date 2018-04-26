
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
	Name	string	`json:"name"`
	SqlInjectionMatchTuples	string	`json:"sql_injection_match_tuples"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafSqlInjectionMatchSetList is a list of AwsWafSqlInjectionMatchSet resources
type AwsWafSqlInjectionMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafSqlInjectionMatchSet	`json:"items"`
}

