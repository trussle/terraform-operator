
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMqConfiguration describes a AwsMqConfiguration resource
type AwsMqConfiguration struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsMqConfigurationSpec	`json:"spec"`
}


// AwsMqConfigurationSpec is the spec for a AwsMqConfiguration Resource
type AwsMqConfigurationSpec struct {
	EngineVersion	string	`json:"engine_version"`
	Name	string	`json:"name"`
	Data	string	`json:"data"`
	Description	string	`json:"description"`
	EngineType	string	`json:"engine_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMqConfigurationList is a list of AwsMqConfiguration resources
type AwsMqConfigurationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsMqConfiguration	`json:"items"`
}

