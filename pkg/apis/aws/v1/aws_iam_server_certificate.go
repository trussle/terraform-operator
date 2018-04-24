
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamServerCertificate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamServerCertificateSpec	`json:"spec"`
}

type AwsIamServerCertificateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamServerCertificate	`json:"items"`
}

type AwsIamServerCertificateSpec struct {
	CertificateBody	string	`json:"certificate_body"`
	CertificateChain	string	`json:"certificate_chain"`
	Path	string	`json:"path"`
	PrivateKey	string	`json:"private_key"`
	NamePrefix	string	`json:"name_prefix"`
}
