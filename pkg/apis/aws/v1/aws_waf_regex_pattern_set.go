
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRegexPatternSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafRegexPatternSetSpec	`json:"spec"`
}

type AwsWafRegexPatternSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafRegexPatternSet	`json:"items"`
}

type AwsWafRegexPatternSetSpec struct {
	Name	string	`json:"name"`
	RegexPatternStrings	interface{}	`json:"regex_pattern_strings"`
}
