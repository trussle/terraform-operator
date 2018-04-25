
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamAccessKey describes a AwsIamAccessKey resource
type AwsIamAccessKey struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamAccessKeySpec	`json:"spec"`
}


// AwsIamAccessKeySpec is the spec for a AwsIamAccessKey Resource
type AwsIamAccessKeySpec struct {
	User	string	`json:"user"`
	PgpKey	string	`json:"pgp_key"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamAccessKeyList is a list of AwsIamAccessKey resources
type AwsIamAccessKeyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamAccessKey	`json:"items"`
}

