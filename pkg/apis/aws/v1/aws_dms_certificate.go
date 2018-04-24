
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsCertificate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDmsCertificateSpec	`json:"spec"`
}

type AwsDmsCertificateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsCertificate	`json:"items"`
}

type AwsDmsCertificateSpec struct {
	CertificatePem	string	`json:"certificate_pem"`
	CertificateWallet	string	`json:"certificate_wallet"`
	CertificateId	string	`json:"certificate_id"`
}
