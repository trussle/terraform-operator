
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafXssMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafXssMatchSetSpec	`json:"spec"`
}

type AwsWafXssMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafXssMatchSet	`json:"items"`
}

type AwsWafXssMatchSetSpec struct {
	Name	string	`json:"name"`
	XssMatchTuples	interface{}	`json:"xss_match_tuples"`
}
