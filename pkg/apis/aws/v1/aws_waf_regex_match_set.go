
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRegexMatchSet describes a AwsWafRegexMatchSet resource
type AwsWafRegexMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafRegexMatchSetSpec	`json:"spec"`
}


// AwsWafRegexMatchSetSpec is the spec for a AwsWafRegexMatchSet Resource
type AwsWafRegexMatchSetSpec struct {
	Name	string	`json:"name"`
	RegexMatchTuple	RegexMatchTuple	`json:"regex_match_tuple"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRegexMatchSetList is a list of AwsWafRegexMatchSet resources
type AwsWafRegexMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafRegexMatchSet	`json:"items"`
}


// FieldToMatch is a FieldToMatch
type FieldToMatch struct {
	Type	string	`json:"type"`
	Data	string	`json:"data"`
}

// RegexMatchTuple is a RegexMatchTuple
type RegexMatchTuple struct {
	FieldToMatch	[]FieldToMatch	`json:"field_to_match"`
	RegexPatternSetId	string	`json:"regex_pattern_set_id"`
	TextTransformation	string	`json:"text_transformation"`
}
