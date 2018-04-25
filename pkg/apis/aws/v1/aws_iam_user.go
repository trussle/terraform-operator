
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUser describes a AwsIamUser resource
type AwsIamUser struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamUserSpec	`json:"spec"`
}


// AwsIamUserSpec is the spec for a AwsIamUser Resource
type AwsIamUserSpec struct {
	Name	string	`json:"name"`
	Path	string	`json:"path"`
	ForceDestroy	bool	`json:"force_destroy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserList is a list of AwsIamUser resources
type AwsIamUserList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamUser	`json:"items"`
}

