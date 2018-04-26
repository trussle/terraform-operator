
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRegexPatternSet describes a AwsWafRegexPatternSet resource
type AwsWafRegexPatternSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafRegexPatternSetSpec	`json:"spec"`
}


// AwsWafRegexPatternSetSpec is the spec for a AwsWafRegexPatternSet Resource
type AwsWafRegexPatternSetSpec struct {
	RegexPatternStrings	string	`json:"regex_pattern_strings"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRegexPatternSetList is a list of AwsWafRegexPatternSet resources
type AwsWafRegexPatternSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafRegexPatternSet	`json:"items"`
}

