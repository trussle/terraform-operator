
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalRegexPatternSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalRegexPatternSetSpec	`json:"spec"`
}

type AwsWafregionalRegexPatternSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalRegexPatternSet	`json:"items"`
}

type AwsWafregionalRegexPatternSetSpec struct {
	RegexPatternStrings	interface{}	`json:"regex_pattern_strings"`
	Name	string	`json:"name"`
}
