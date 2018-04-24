
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAcmCertificate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAcmCertificateSpec	`json:"spec"`
}

type AwsAcmCertificateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAcmCertificate	`json:"items"`
}

type AwsAcmCertificateSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	DomainName	string	`json:"domain_name"`
	SubjectAlternativeNames	[]interface{}	`json:"subject_alternative_names"`
	ValidationMethod	string	`json:"validation_method"`
}
