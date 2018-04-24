
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalGeoMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalGeoMatchSetSpec	`json:"spec"`
}

type AwsWafregionalGeoMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalGeoMatchSet	`json:"items"`
}

type AwsWafregionalGeoMatchSetSpec struct {
	Name	string	`json:"name"`
	GeoMatchConstraint	interface{}	`json:"geo_match_constraint"`
}
