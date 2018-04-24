
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigConfigurationRecorder struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsConfigConfigurationRecorderSpec	`json:"spec"`
}

type AwsConfigConfigurationRecorderList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsConfigConfigurationRecorder	`json:"items"`
}

type AwsConfigConfigurationRecorderSpec struct {
	RoleArn	string	`json:"role_arn"`
	Name	string	`json:"name"`
}
