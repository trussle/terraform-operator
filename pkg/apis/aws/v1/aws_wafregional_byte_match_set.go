
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalByteMatchSet describes a AwsWafregionalByteMatchSet resource
type AwsWafregionalByteMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalByteMatchSetSpec	`json:"spec"`
}


// AwsWafregionalByteMatchSetSpec is the spec for a AwsWafregionalByteMatchSet Resource
type AwsWafregionalByteMatchSetSpec struct {
	Name	string	`json:"name"`
	ByteMatchTuple	AwsWafregionalByteMatchSetByteMatchTuple	`json:"byte_match_tuple"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalByteMatchSetList is a list of AwsWafregionalByteMatchSet resources
type AwsWafregionalByteMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalByteMatchSet	`json:"items"`
}


// AwsWafregionalByteMatchSetByteMatchTuple is a AwsWafregionalByteMatchSetByteMatchTuple
type AwsWafregionalByteMatchSetByteMatchTuple struct {
	TargetString	string	`json:"target_string"`
	TextTransformation	string	`json:"text_transformation"`
	FieldToMatch	AwsWafregionalByteMatchSetFieldToMatch	`json:"field_to_match"`
	PositionalConstraint	string	`json:"positional_constraint"`
}

// AwsWafregionalByteMatchSetFieldToMatch is a AwsWafregionalByteMatchSetFieldToMatch
type AwsWafregionalByteMatchSetFieldToMatch struct {
	Data	string	`json:"data"`
	Type	string	`json:"type"`
}
