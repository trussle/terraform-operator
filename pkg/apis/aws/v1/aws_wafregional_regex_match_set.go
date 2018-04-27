
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRegexMatchSet describes a AwsWafregionalRegexMatchSet resource
type AwsWafregionalRegexMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalRegexMatchSetSpec	`json:"spec"`
}


// AwsWafregionalRegexMatchSetSpec is the spec for a AwsWafregionalRegexMatchSet Resource
type AwsWafregionalRegexMatchSetSpec struct {
	Name	string	`json:"name"`
	RegexMatchTuple	AwsWafregionalRegexMatchSetRegexMatchTuple	`json:"regex_match_tuple"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRegexMatchSetList is a list of AwsWafregionalRegexMatchSet resources
type AwsWafregionalRegexMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalRegexMatchSet	`json:"items"`
}


// AwsWafregionalRegexMatchSetFieldToMatch is a AwsWafregionalRegexMatchSetFieldToMatch
type AwsWafregionalRegexMatchSetFieldToMatch struct {
	Data	string	`json:"data"`
	Type	string	`json:"type"`
}

// AwsWafregionalRegexMatchSetRegexMatchTuple is a AwsWafregionalRegexMatchSetRegexMatchTuple
type AwsWafregionalRegexMatchSetRegexMatchTuple struct {
	FieldToMatch	[]AwsWafregionalRegexMatchSetFieldToMatch	`json:"field_to_match"`
	RegexPatternSetId	string	`json:"regex_pattern_set_id"`
	TextTransformation	string	`json:"text_transformation"`
}
