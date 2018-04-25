
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigConfigurationRecorder describes a AwsConfigConfigurationRecorder resource
type AwsConfigConfigurationRecorder struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsConfigConfigurationRecorderSpec	`json:"spec"`
}


// AwsConfigConfigurationRecorderSpec is the spec for a AwsConfigConfigurationRecorder Resource
type AwsConfigConfigurationRecorderSpec struct {
	Name	string	`json:"name"`
	RoleArn	string	`json:"role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigConfigurationRecorderList is a list of AwsConfigConfigurationRecorder resources
type AwsConfigConfigurationRecorderList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsConfigConfigurationRecorder	`json:"items"`
}

