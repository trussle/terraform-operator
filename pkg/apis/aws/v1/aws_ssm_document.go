
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmDocument struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmDocumentSpec	`json:"spec"`
}

type AwsSsmDocumentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmDocument	`json:"items"`
}

type AwsSsmDocumentSpec struct {
	Content	string	`json:"content"`
	DocumentType	string	`json:"document_type"`
	DocumentFormat	string	`json:"document_format"`
	Name	string	`json:"name"`
	Permissions	map[string]interface{}	`json:"permissions"`
}
