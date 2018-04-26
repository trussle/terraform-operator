
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDomainName describes a AwsApiGatewayDomainName resource
type AwsApiGatewayDomainName struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayDomainNameSpec	`json:"spec"`
}


// AwsApiGatewayDomainNameSpec is the spec for a AwsApiGatewayDomainName Resource
type AwsApiGatewayDomainNameSpec struct {
	CertificateArn	string	`json:"certificate_arn"`
	CertificateBody	string	`json:"certificate_body"`
	CertificateChain	string	`json:"certificate_chain"`
	CertificateName	string	`json:"certificate_name"`
	CertificatePrivateKey	string	`json:"certificate_private_key"`
	DomainName	string	`json:"domain_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDomainNameList is a list of AwsApiGatewayDomainName resources
type AwsApiGatewayDomainNameList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayDomainName	`json:"items"`
}

