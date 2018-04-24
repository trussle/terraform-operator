
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbOptionGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbOptionGroupSpec	`json:"spec"`
}

type AwsDbOptionGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbOptionGroup	`json:"items"`
}

type AwsDbOptionGroupSpec struct {
	MajorEngineVersion	string	`json:"major_engine_version"`
	OptionGroupDescription	string	`json:"option_group_description"`
	Option	interface{}	`json:"option"`
	Tags	map[string]interface{}	`json:"tags"`
	EngineName	string	`json:"engine_name"`
}
