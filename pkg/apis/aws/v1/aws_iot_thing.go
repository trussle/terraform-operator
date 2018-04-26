
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotThing describes a AwsIotThing resource
type AwsIotThing struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIotThingSpec	`json:"spec"`
}


// AwsIotThingSpec is the spec for a AwsIotThing Resource
type AwsIotThingSpec struct {
	Name	string	`json:"name"`
	Attributes	map[string]Generic	`json:"attributes"`
	ThingTypeName	string	`json:"thing_type_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotThingList is a list of AwsIotThing resources
type AwsIotThingList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotThing	`json:"items"`
}

