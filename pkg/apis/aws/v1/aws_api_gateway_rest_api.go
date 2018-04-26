
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayRestApi describes a AwsApiGatewayRestApi resource
type AwsApiGatewayRestApi struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayRestApiSpec	`json:"spec"`
}


// AwsApiGatewayRestApiSpec is the spec for a AwsApiGatewayRestApi Resource
type AwsApiGatewayRestApiSpec struct {
	MinimumCompressionSize	int	`json:"minimum_compression_size"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Policy	string	`json:"policy"`
	BinaryMediaTypes	[]string	`json:"binary_media_types"`
	Body	string	`json:"body"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayRestApiList is a list of AwsApiGatewayRestApi resources
type AwsApiGatewayRestApiList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayRestApi	`json:"items"`
}

