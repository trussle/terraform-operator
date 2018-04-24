
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafByteMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafByteMatchSetSpec	`json:"spec"`
}

type AwsWafByteMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafByteMatchSet	`json:"items"`
}

type AwsWafByteMatchSetSpec struct {
	Name	string	`json:"name"`
	ByteMatchTuples	interface{}	`json:"byte_match_tuples"`
}
