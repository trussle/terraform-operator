
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
	DocumentFormat	string	`json:"document_format"`
	DocumentType	string	`json:"document_type"`
	Name	string	`json:"name"`
	Content	string	`json:"content"`
	Permissions	map[string]zgexyALB	`json:"permissions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmDocumentList is a list of AwsSsmDocument resources
type AwsSsmDocumentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmDocument	`json:"items"`
}


// XVbWGvbq is a XVbWGvbq
type XVbWGvbq struct {
	Name	string	`json:"name"`
	DefaultValue	string	`json:"default_value"`
	Description	string	`json:"description"`
	Type	string	`json:"type"`
}

// zgexyALB is a zgexyALB
type zgexyALB struct {
	Type	string	`json:"type"`
	AccountIds	string	`json:"account_ids"`
}
