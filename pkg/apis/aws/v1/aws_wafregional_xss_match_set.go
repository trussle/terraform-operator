
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalXssMatchSet describes a AwsWafregionalXssMatchSet resource
type AwsWafregionalXssMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalXssMatchSetSpec	`json:"spec"`
}


// AwsWafregionalXssMatchSetSpec is the spec for a AwsWafregionalXssMatchSet Resource
type AwsWafregionalXssMatchSetSpec struct {
	Name	string	`json:"name"`
	XssMatchTuple	Generic	`json:"xss_match_tuple"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalXssMatchSetList is a list of AwsWafregionalXssMatchSet resources
type AwsWafregionalXssMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalXssMatchSet	`json:"items"`
}

