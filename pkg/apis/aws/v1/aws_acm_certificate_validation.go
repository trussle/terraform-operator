
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAcmCertificateValidation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAcmCertificateValidationSpec	`json:"spec"`
}

type AwsAcmCertificateValidationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAcmCertificateValidation	`json:"items"`
}

type AwsAcmCertificateValidationSpec struct {
	CertificateArn	string	`json:"certificate_arn"`
	ValidationRecordFqdns	interface{}	`json:"validation_record_fqdns"`
}
