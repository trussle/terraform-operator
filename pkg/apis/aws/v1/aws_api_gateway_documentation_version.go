
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDocumentationVersion describes a AwsApiGatewayDocumentationVersion resource
type AwsApiGatewayDocumentationVersion struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayDocumentationVersionSpec	`json:"spec"`
}


// AwsApiGatewayDocumentationVersionSpec is the spec for a AwsApiGatewayDocumentationVersion Resource
type AwsApiGatewayDocumentationVersionSpec struct {
	RestApiId	string	`json:"rest_api_id"`
	Description	string	`json:"description"`
	Version	string	`json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDocumentationVersionList is a list of AwsApiGatewayDocumentationVersion resources
type AwsApiGatewayDocumentationVersionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayDocumentationVersion	`json:"items"`
}

