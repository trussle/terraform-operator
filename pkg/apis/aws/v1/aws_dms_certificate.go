
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsCertificate describes a AwsDmsCertificate resource
type AwsDmsCertificate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDmsCertificateSpec	`json:"spec"`
}


// AwsDmsCertificateSpec is the spec for a AwsDmsCertificate Resource
type AwsDmsCertificateSpec struct {
	CertificateId	string	`json:"certificate_id"`
	CertificatePem	string	`json:"certificate_pem"`
	CertificateWallet	string	`json:"certificate_wallet"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsCertificateList is a list of AwsDmsCertificate resources
type AwsDmsCertificateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsCertificate	`json:"items"`
}

