
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafSizeConstraintSet describes a AwsWafSizeConstraintSet resource
type AwsWafSizeConstraintSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafSizeConstraintSetSpec	`json:"spec"`
}


// AwsWafSizeConstraintSetSpec is the spec for a AwsWafSizeConstraintSet Resource
type AwsWafSizeConstraintSetSpec struct {
	Name	string	`json:"name"`
	SizeConstraints	SizeConstraints	`json:"size_constraints"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafSizeConstraintSetList is a list of AwsWafSizeConstraintSet resources
type AwsWafSizeConstraintSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafSizeConstraintSet	`json:"items"`
}


// FieldToMatch is a FieldToMatch
type FieldToMatch struct {
	Data	string	`json:"data"`
	Type	string	`json:"type"`
}

// SizeConstraints is a SizeConstraints
type SizeConstraints struct {
	Size	int	`json:"size"`
	TextTransformation	string	`json:"text_transformation"`
	FieldToMatch	FieldToMatch	`json:"field_to_match"`
	ComparisonOperator	string	`json:"comparison_operator"`
}
