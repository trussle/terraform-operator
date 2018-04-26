
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalGeoMatchSet describes a AwsWafregionalGeoMatchSet resource
type AwsWafregionalGeoMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalGeoMatchSetSpec	`json:"spec"`
}


// AwsWafregionalGeoMatchSetSpec is the spec for a AwsWafregionalGeoMatchSet Resource
type AwsWafregionalGeoMatchSetSpec struct {
	Name	string	`json:"name"`
	GeoMatchConstraint	GeoMatchConstraint	`json:"geo_match_constraint"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalGeoMatchSetList is a list of AwsWafregionalGeoMatchSet resources
type AwsWafregionalGeoMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalGeoMatchSet	`json:"items"`
}


// GeoMatchConstraint is a GeoMatchConstraint
type GeoMatchConstraint struct {
	Type	string	`json:"type"`
	Value	string	`json:"value"`
}
