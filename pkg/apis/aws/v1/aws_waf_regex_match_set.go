
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRegexMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafRegexMatchSetSpec	`json:"spec"`
}

type AwsWafRegexMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafRegexMatchSet	`json:"items"`
}

type AwsWafRegexMatchSetSpec struct {
	Name	string	`json:"name"`
	RegexMatchTuple	interface{}	`json:"regex_match_tuple"`
}
