
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigConfigurationRecorderStatus struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsConfigConfigurationRecorderStatusSpec	`json:"spec"`
}

type AwsConfigConfigurationRecorderStatusList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsConfigConfigurationRecorderStatus	`json:"items"`
}

type AwsConfigConfigurationRecorderStatusSpec struct {
	Name	string	`json:"name"`
	IsEnabled	bool	`json:"is_enabled"`
}
