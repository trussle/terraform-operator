package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Policy describes an AWS IAM Policy
type Policy struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	Spec PolicySpec `json:"spec"`
}

// PolicySpec is the spec for Policy
type PolicySpec struct {
	// The policy document. This is a json formatted string
	Document string `json:"policyDocument"`
	Name     string `json:"name"`
}
