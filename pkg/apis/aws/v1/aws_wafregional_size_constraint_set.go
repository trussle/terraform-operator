
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
	SizeConstraints	AwsWafregionalSizeConstraintSetSizeConstraints	`json:"size_constraints"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalSizeConstraintSetList is a list of AwsWafregionalSizeConstraintSet resources
type AwsWafregionalSizeConstraintSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalSizeConstraintSet	`json:"items"`
}


// AwsWafregionalSizeConstraintSetFieldToMatch is a AwsWafregionalSizeConstraintSetFieldToMatch
type AwsWafregionalSizeConstraintSetFieldToMatch struct {
	Data	string	`json:"data"`
	Type	string	`json:"type"`
}

// AwsWafregionalSizeConstraintSetSizeConstraints is a AwsWafregionalSizeConstraintSetSizeConstraints
type AwsWafregionalSizeConstraintSetSizeConstraints struct {
	FieldToMatch	AwsWafregionalSizeConstraintSetFieldToMatch	`json:"field_to_match"`
	ComparisonOperator	string	`json:"comparison_operator"`
	Size	int	`json:"size"`
	TextTransformation	string	`json:"text_transformation"`
}
