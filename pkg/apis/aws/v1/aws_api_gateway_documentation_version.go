
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDocumentationVersion struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayDocumentationVersionSpec	`json:"spec"`
}

type AwsApiGatewayDocumentationVersionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayDocumentationVersion	`json:"items"`
}

type AwsApiGatewayDocumentationVersionSpec struct {
	Version	string	`json:"version"`
	RestApiId	string	`json:"rest_api_id"`
	Description	string	`json:"description"`
}
