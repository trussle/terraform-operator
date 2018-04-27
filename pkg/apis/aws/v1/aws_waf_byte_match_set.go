
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
	Name	string	`json:"name"`
	ByteMatchTuples	AwsWafByteMatchSetByteMatchTuples	`json:"byte_match_tuples"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafByteMatchSetList is a list of AwsWafByteMatchSet resources
type AwsWafByteMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafByteMatchSet	`json:"items"`
}


// AwsWafByteMatchSetFieldToMatch is a AwsWafByteMatchSetFieldToMatch
type AwsWafByteMatchSetFieldToMatch struct {
	Data	string	`json:"data"`
	Type	string	`json:"type"`
}

// AwsWafByteMatchSetByteMatchTuples is a AwsWafByteMatchSetByteMatchTuples
type AwsWafByteMatchSetByteMatchTuples struct {
	FieldToMatch	AwsWafByteMatchSetFieldToMatch	`json:"field_to_match"`
	PositionalConstraint	string	`json:"positional_constraint"`
	TargetString	string	`json:"target_string"`
	TextTransformation	string	`json:"text_transformation"`
}
