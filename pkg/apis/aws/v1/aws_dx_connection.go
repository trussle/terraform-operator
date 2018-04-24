
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxConnection struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDxConnectionSpec	`json:"spec"`
}

type AwsDxConnectionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDxConnection	`json:"items"`
}

type AwsDxConnectionSpec struct {
	Bandwidth	string	`json:"bandwidth"`
	Location	string	`json:"location"`
	Tags	map[string]interface{}	`json:"tags"`
	Name	string	`json:"name"`
}
