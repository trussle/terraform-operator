
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayRestApi struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayRestApiSpec	`json:"spec"`
}

type AwsApiGatewayRestApiList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayRestApi	`json:"items"`
}

type AwsApiGatewayRestApiSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	BinaryMediaTypes	[]interface{}	`json:"binary_media_types"`
	Body	string	`json:"body"`
	MinimumCompressionSize	int	`json:"minimum_compression_size"`
}
