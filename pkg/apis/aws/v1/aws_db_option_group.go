
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbOptionGroup describes a AwsDbOptionGroup resource
type AwsDbOptionGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbOptionGroupSpec	`json:"spec"`
}


// AwsDbOptionGroupSpec is the spec for a AwsDbOptionGroup Resource
type AwsDbOptionGroupSpec struct {
	EngineName	string	`json:"engine_name"`
	MajorEngineVersion	string	`json:"major_engine_version"`
	OptionGroupDescription	string	`json:"option_group_description"`
	Option	string	`json:"option"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbOptionGroupList is a list of AwsDbOptionGroup resources
type AwsDbOptionGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbOptionGroup	`json:"items"`
}

