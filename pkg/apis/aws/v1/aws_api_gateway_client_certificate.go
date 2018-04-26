
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayClientCertificate describes a AwsApiGatewayClientCertificate resource
type AwsApiGatewayClientCertificate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayClientCertificateSpec	`json:"spec"`
}


// AwsApiGatewayClientCertificateSpec is the spec for a AwsApiGatewayClientCertificate Resource
type AwsApiGatewayClientCertificateSpec struct {
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayClientCertificateList is a list of AwsApiGatewayClientCertificate resources
type AwsApiGatewayClientCertificateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayClientCertificate	`json:"items"`
}

