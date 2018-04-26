
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxConnection describes a AwsDxConnection resource
type AwsDxConnection struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDxConnectionSpec	`json:"spec"`
}


// AwsDxConnectionSpec is the spec for a AwsDxConnection Resource
type AwsDxConnectionSpec struct {
	Name	string	`json:"name"`
	Bandwidth	string	`json:"bandwidth"`
	Location	string	`json:"location"`
	Tags	map[string]???	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxConnectionList is a list of AwsDxConnection resources
type AwsDxConnectionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDxConnection	`json:"items"`
}

