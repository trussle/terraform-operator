
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayDomainName struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayDomainNameSpec	`json:"spec"`
}

type AwsApiGatewayDomainNameList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayDomainName	`json:"items"`
}

type AwsApiGatewayDomainNameSpec struct {
	CertificateBody	string	`json:"certificate_body"`
	DomainName	string	`json:"domain_name"`
	CertificateArn	string	`json:"certificate_arn"`
	CertificateChain	string	`json:"certificate_chain"`
	CertificateName	string	`json:"certificate_name"`
	CertificatePrivateKey	string	`json:"certificate_private_key"`
}
