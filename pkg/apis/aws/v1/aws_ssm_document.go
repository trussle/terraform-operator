
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmDocument describes a AwsSsmDocument resource
type AwsSsmDocument struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmDocumentSpec	`json:"spec"`
}


// AwsSsmDocumentSpec is the spec for a AwsSsmDocument Resource
type AwsSsmDocumentSpec struct {
	DocumentType	string	`json:"document_type"`
	Content	string	`json:"content"`
	DocumentFormat	string	`json:"document_format"`
	Permissions	map[string]Generic	`json:"permissions"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmDocumentList is a list of AwsSsmDocument resources
type AwsSsmDocumentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmDocument	`json:"items"`
}

