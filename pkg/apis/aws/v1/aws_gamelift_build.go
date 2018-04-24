
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGameliftBuild struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGameliftBuildSpec	`json:"spec"`
}

type AwsGameliftBuildList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGameliftBuild	`json:"items"`
}

type AwsGameliftBuildSpec struct {
	Name	string	`json:"name"`
	OperatingSystem	string	`json:"operating_system"`
	StorageLocation	[]interface{}	`json:"storage_location"`
	Version	string	`json:"version"`
}
