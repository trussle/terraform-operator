
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Version	string	`json:"version"`
	RestApiId	string	`json:"rest_api_id"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDocumentationVersionList is a list of AwsApiGatewayDocumentationVersion resources
type AwsApiGatewayDocumentationVersionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayDocumentationVersion	`json:"items"`
}

