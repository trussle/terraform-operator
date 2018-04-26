
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAcmCertificate describes a AwsAcmCertificate resource
type AwsAcmCertificate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAcmCertificateSpec	`json:"spec"`
}


// AwsAcmCertificateSpec is the spec for a AwsAcmCertificate Resource
type AwsAcmCertificateSpec struct {
	Tags	map[string]string	`json:"tags"`
	DomainName	string	`json:"domain_name"`
	SubjectAlternativeNames	[]string	`json:"subject_alternative_names"`
	ValidationMethod	string	`json:"validation_method"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAcmCertificateList is a list of AwsAcmCertificate resources
type AwsAcmCertificateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAcmCertificate	`json:"items"`
}


// DomainValidationOptions is a DomainValidationOptions
type DomainValidationOptions struct {
}
