
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalSizeConstraintSet describes a AwsWafregionalSizeConstraintSet resource
type AwsWafregionalSizeConstraintSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalSizeConstraintSetSpec	`json:"spec"`
}


// AwsWafregionalSizeConstraintSetSpec is the spec for a AwsWafregionalSizeConstraintSet Resource
type AwsWafregionalSizeConstraintSetSpec struct {
	Name	string	`json:"name"`
	SizeConstraints	SizeConstraints	`json:"size_constraints"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalSizeConstraintSetList is a list of AwsWafregionalSizeConstraintSet resources
type AwsWafregionalSizeConstraintSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalSizeConstraintSet	`json:"items"`
}


// FieldToMatch is a FieldToMatch
type FieldToMatch struct {
	Type	string	`json:"type"`
	Data	string	`json:"data"`
}

// SizeConstraints is a SizeConstraints
type SizeConstraints struct {
	FieldToMatch	FieldToMatch	`json:"field_to_match"`
	ComparisonOperator	string	`json:"comparison_operator"`
	Size	int	`json:"size"`
	TextTransformation	string	`json:"text_transformation"`
}
