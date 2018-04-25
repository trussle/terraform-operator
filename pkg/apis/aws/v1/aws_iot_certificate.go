
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotCertificate describes a AwsIotCertificate resource
type AwsIotCertificate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIotCertificateSpec	`json:"spec"`
}


// AwsIotCertificateSpec is the spec for a AwsIotCertificate Resource
type AwsIotCertificateSpec struct {
	Csr	string	`json:"csr"`
	Active	bool	`json:"active"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotCertificateList is a list of AwsIotCertificate resources
type AwsIotCertificateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotCertificate	`json:"items"`
}

