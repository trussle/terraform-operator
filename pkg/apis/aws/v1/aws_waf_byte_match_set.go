
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafByteMatchSet describes a AwsWafByteMatchSet resource
type AwsWafByteMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafByteMatchSetSpec	`json:"spec"`
}


// AwsWafByteMatchSetSpec is the spec for a AwsWafByteMatchSet Resource
type AwsWafByteMatchSetSpec struct {
	ByteMatchTuples	ByteMatchTuples	`json:"byte_match_tuples"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafByteMatchSetList is a list of AwsWafByteMatchSet resources
type AwsWafByteMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafByteMatchSet	`json:"items"`
}


// FieldToMatch is a FieldToMatch
type FieldToMatch struct {
	Data	string	`json:"data"`
	Type	string	`json:"type"`
}

// ByteMatchTuples is a ByteMatchTuples
type ByteMatchTuples struct {
	TextTransformation	string	`json:"text_transformation"`
	FieldToMatch	FieldToMatch	`json:"field_to_match"`
	PositionalConstraint	string	`json:"positional_constraint"`
	TargetString	string	`json:"target_string"`
}
