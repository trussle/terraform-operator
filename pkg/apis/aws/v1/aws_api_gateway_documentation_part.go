
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDocumentationPart describes a AwsApiGatewayDocumentationPart resource
type AwsApiGatewayDocumentationPart struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayDocumentationPartSpec	`json:"spec"`
}


// AwsApiGatewayDocumentationPartSpec is the spec for a AwsApiGatewayDocumentationPart Resource
type AwsApiGatewayDocumentationPartSpec struct {
	RestApiId	string	`json:"rest_api_id"`
	Location	[]mBdtZtcO	`json:"location"`
	Properties	string	`json:"properties"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDocumentationPartList is a list of AwsApiGatewayDocumentationPart resources
type AwsApiGatewayDocumentationPartList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayDocumentationPart	`json:"items"`
}


// mBdtZtcO is a mBdtZtcO
type mBdtZtcO struct {
	Method	string	`json:"method"`
	Name	string	`json:"name"`
	Path	string	`json:"path"`
	StatusCode	string	`json:"status_code"`
	Type	string	`json:"type"`
}
