
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGameliftBuild describes a AwsGameliftBuild resource
type AwsGameliftBuild struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGameliftBuildSpec	`json:"spec"`
}


// AwsGameliftBuildSpec is the spec for a AwsGameliftBuild Resource
type AwsGameliftBuildSpec struct {
	StorageLocation	[]interface{}	`json:"storage_location"`
	Version	string	`json:"version"`
	Name	string	`json:"name"`
	OperatingSystem	string	`json:"operating_system"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGameliftBuildList is a list of AwsGameliftBuild resources
type AwsGameliftBuildList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGameliftBuild	`json:"items"`
}

