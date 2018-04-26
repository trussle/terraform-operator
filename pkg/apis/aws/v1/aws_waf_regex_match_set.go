
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
	RegexMatchTuple	string	`json:"regex_match_tuple"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRegexMatchSetList is a list of AwsWafRegexMatchSet resources
type AwsWafRegexMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafRegexMatchSet	`json:"items"`
}

