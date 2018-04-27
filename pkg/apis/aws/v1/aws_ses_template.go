
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesTemplate describes a AwsSesTemplate resource
type AwsSesTemplate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesTemplateSpec	`json:"spec"`
}


// AwsSesTemplateSpec is the spec for a AwsSesTemplate Resource
type AwsSesTemplateSpec struct {
	Subject	string	`json:"subject"`
	Text	string	`json:"text"`
	Name	string	`json:"name"`
	Html	string	`json:"html"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesTemplateList is a list of AwsSesTemplate resources
type AwsSesTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesTemplate	`json:"items"`
}

