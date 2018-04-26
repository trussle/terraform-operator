
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Properties	[]Properties	`json:"properties"`
	Deprecated	bool	`json:"deprecated"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotThingTypeList is a list of AwsIotThingType resources
type AwsIotThingTypeList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotThingType	`json:"items"`
}


// Properties is a Properties
type Properties struct {
	Description	string	`json:"description"`
}
