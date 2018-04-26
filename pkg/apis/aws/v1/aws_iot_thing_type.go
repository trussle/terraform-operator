
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotThingType describes a AwsIotThingType resource
type AwsIotThingType struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIotThingTypeSpec	`json:"spec"`
}


// AwsIotThingTypeSpec is the spec for a AwsIotThingType Resource
type AwsIotThingTypeSpec struct {
	Name	string	`json:"name"`
	Properties	[]Generic	`json:"properties"`
	Deprecated	bool	`json:"deprecated"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotThingTypeList is a list of AwsIotThingType resources
type AwsIotThingTypeList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotThingType	`json:"items"`
}

