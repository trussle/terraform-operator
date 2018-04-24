
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalXssMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalXssMatchSetSpec	`json:"spec"`
}

type AwsWafregionalXssMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalXssMatchSet	`json:"items"`
}

type AwsWafregionalXssMatchSetSpec struct {
	Name	string	`json:"name"`
	XssMatchTuple	interface{}	`json:"xss_match_tuple"`
}
