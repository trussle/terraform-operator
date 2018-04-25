
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamGroup describes a AwsIamGroup resource
type AwsIamGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamGroupSpec	`json:"spec"`
}


// AwsIamGroupSpec is the spec for a AwsIamGroup Resource
type AwsIamGroupSpec struct {
	Name	string	`json:"name"`
	Path	string	`json:"path"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamGroupList is a list of AwsIamGroup resources
type AwsIamGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamGroup	`json:"items"`
}

