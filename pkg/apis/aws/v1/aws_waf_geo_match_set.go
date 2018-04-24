
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafGeoMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafGeoMatchSetSpec	`json:"spec"`
}

type AwsWafGeoMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafGeoMatchSet	`json:"items"`
}

type AwsWafGeoMatchSetSpec struct {
	Name	string	`json:"name"`
	GeoMatchConstraint	interface{}	`json:"geo_match_constraint"`
}
