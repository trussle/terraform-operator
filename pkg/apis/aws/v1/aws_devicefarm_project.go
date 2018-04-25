
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDevicefarmProject describes a AwsDevicefarmProject resource
type AwsDevicefarmProject struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDevicefarmProjectSpec	`json:"spec"`
}


// AwsDevicefarmProjectSpec is the spec for a AwsDevicefarmProject Resource
type AwsDevicefarmProjectSpec struct {
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDevicefarmProjectList is a list of AwsDevicefarmProject resources
type AwsDevicefarmProjectList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDevicefarmProject	`json:"items"`
}

