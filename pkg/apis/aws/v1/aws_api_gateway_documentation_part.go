
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDocumentationPart struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayDocumentationPartSpec	`json:"spec"`
}

type AwsApiGatewayDocumentationPartList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayDocumentationPart	`json:"items"`
}

type AwsApiGatewayDocumentationPartSpec struct {
	Location	[]interface{}	`json:"location"`
	Properties	string	`json:"properties"`
	RestApiId	string	`json:"rest_api_id"`
}
