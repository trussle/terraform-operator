
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotThing struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIotThingSpec	`json:"spec"`
}

type AwsIotThingList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotThing	`json:"items"`
}

type AwsIotThingSpec struct {
	Name	string	`json:"name"`
	Attributes	map[string]interface{}	`json:"attributes"`
	ThingTypeName	string	`json:"thing_type_name"`
}
