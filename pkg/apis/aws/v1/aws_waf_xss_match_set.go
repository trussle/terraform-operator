
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafXssMatchSet describes a AwsWafXssMatchSet resource
type AwsWafXssMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafXssMatchSetSpec	`json:"spec"`
}


// AwsWafXssMatchSetSpec is the spec for a AwsWafXssMatchSet Resource
type AwsWafXssMatchSetSpec struct {
	Name	string	`json:"name"`
	XssMatchTuples	XssMatchTuples	`json:"xss_match_tuples"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafXssMatchSetList is a list of AwsWafXssMatchSet resources
type AwsWafXssMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafXssMatchSet	`json:"items"`
}


// FieldToMatch is a FieldToMatch
type FieldToMatch struct {
	Data	string	`json:"data"`
	Type	string	`json:"type"`
}

// XssMatchTuples is a XssMatchTuples
type XssMatchTuples struct {
	FieldToMatch	FieldToMatch	`json:"field_to_match"`
	TextTransformation	string	`json:"text_transformation"`
}
