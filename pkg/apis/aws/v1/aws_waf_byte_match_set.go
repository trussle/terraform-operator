
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafByteMatchSet describes a AwsWafByteMatchSet resource
type AwsWafByteMatchSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafByteMatchSetSpec	`json:"spec"`
}


// AwsWafByteMatchSetSpec is the spec for a AwsWafByteMatchSet Resource
type AwsWafByteMatchSetSpec struct {
	Name	string	`json:"name"`
	ByteMatchTuples	string	`json:"byte_match_tuples"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafByteMatchSetList is a list of AwsWafByteMatchSet resources
type AwsWafByteMatchSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafByteMatchSet	`json:"items"`
}

