
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigConfigurationRecorderStatus describes a AwsConfigConfigurationRecorderStatus resource
type AwsConfigConfigurationRecorderStatus struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsConfigConfigurationRecorderStatusSpec	`json:"spec"`
}


// AwsConfigConfigurationRecorderStatusSpec is the spec for a AwsConfigConfigurationRecorderStatus Resource
type AwsConfigConfigurationRecorderStatusSpec struct {
	IsEnabled	bool	`json:"is_enabled"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigConfigurationRecorderStatusList is a list of AwsConfigConfigurationRecorderStatus resources
type AwsConfigConfigurationRecorderStatusList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsConfigConfigurationRecorderStatus	`json:"items"`
}

