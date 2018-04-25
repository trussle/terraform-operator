
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRegexPatternSet describes a AwsWafregionalRegexPatternSet resource
type AwsWafregionalRegexPatternSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalRegexPatternSetSpec	`json:"spec"`
}


// AwsWafregionalRegexPatternSetSpec is the spec for a AwsWafregionalRegexPatternSet Resource
type AwsWafregionalRegexPatternSetSpec struct {
	Name	string	`json:"name"`
	RegexPatternStrings	string	`json:"regex_pattern_strings"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRegexPatternSetList is a list of AwsWafregionalRegexPatternSet resources
type AwsWafregionalRegexPatternSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalRegexPatternSet	`json:"items"`
}

