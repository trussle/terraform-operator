
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesTemplate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesTemplateSpec	`json:"spec"`
}

type AwsSesTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesTemplate	`json:"items"`
}

type AwsSesTemplateSpec struct {
	Html	string	`json:"html"`
	Subject	string	`json:"subject"`
	Text	string	`json:"text"`
	Name	string	`json:"name"`
}
