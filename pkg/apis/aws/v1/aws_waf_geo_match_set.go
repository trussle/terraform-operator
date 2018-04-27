
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafGeoMatchSet describes a AwsWafGeoMatchSet resource
type AwsWafGeoMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafGeoMatchSetSpec	`json:"spec"`
}


// AwsWafGeoMatchSetSpec is the spec for a AwsWafGeoMatchSet Resource
type AwsWafGeoMatchSetSpec struct {
	Name	string	`json:"name"`
	GeoMatchConstraint	AwsWafGeoMatchSetGeoMatchConstraint	`json:"geo_match_constraint"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafGeoMatchSetList is a list of AwsWafGeoMatchSet resources
type AwsWafGeoMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafGeoMatchSet	`json:"items"`
}


// AwsWafGeoMatchSetGeoMatchConstraint is a AwsWafGeoMatchSetGeoMatchConstraint
type AwsWafGeoMatchSetGeoMatchConstraint struct {
	Type	string	`json:"type"`
	Value	string	`json:"value"`
}
