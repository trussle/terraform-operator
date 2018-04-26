
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDaxParameterGroup describes a AwsDaxParameterGroup resource
type AwsDaxParameterGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDaxParameterGroupSpec	`json:"spec"`
}


// AwsDaxParameterGroupSpec is the spec for a AwsDaxParameterGroup Resource
type AwsDaxParameterGroupSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDaxParameterGroupList is a list of AwsDaxParameterGroup resources
type AwsDaxParameterGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDaxParameterGroup	`json:"items"`
}


// Parameters is a Parameters
type Parameters struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}
