
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSfnActivity struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSfnActivitySpec	`json:"spec"`
}

type AwsSfnActivityList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSfnActivity	`json:"items"`
}

type AwsSfnActivitySpec struct {
	Name	string	`json:"name"`
}