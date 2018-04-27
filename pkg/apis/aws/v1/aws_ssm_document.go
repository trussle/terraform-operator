
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Content	string	`json:"content"`
	DocumentFormat	string	`json:"document_format"`
	Name	string	`json:"name"`
	DocumentType	string	`json:"document_type"`
	Permissions	map[string]AwsSsmDocumentPermissions	`json:"permissions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmDocumentList is a list of AwsSsmDocument resources
type AwsSsmDocumentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmDocument	`json:"items"`
}


// AwsSsmDocumentPermissions is a AwsSsmDocumentPermissions
type AwsSsmDocumentPermissions struct {
	Type	string	`json:"type"`
	AccountIds	string	`json:"account_ids"`
}

// AwsSsmDocumentParameter is a AwsSsmDocumentParameter
type AwsSsmDocumentParameter struct {
	DefaultValue	string	`json:"default_value"`
	Description	string	`json:"description"`
	Type	string	`json:"type"`
	Name	string	`json:"name"`
}
