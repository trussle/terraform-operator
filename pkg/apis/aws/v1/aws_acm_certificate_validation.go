
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAcmCertificateValidation describes a AwsAcmCertificateValidation resource
type AwsAcmCertificateValidation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAcmCertificateValidationSpec	`json:"spec"`
}


// AwsAcmCertificateValidationSpec is the spec for a AwsAcmCertificateValidation Resource
type AwsAcmCertificateValidationSpec struct {
	CertificateArn	string	`json:"certificate_arn"`
	ValidationRecordFqdns	Generic	`json:"validation_record_fqdns"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAcmCertificateValidationList is a list of AwsAcmCertificateValidation resources
type AwsAcmCertificateValidationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAcmCertificateValidation	`json:"items"`
}

