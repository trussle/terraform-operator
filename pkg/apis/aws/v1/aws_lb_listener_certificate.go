
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbListenerCertificate describes a AwsLbListenerCertificate resource
type AwsLbListenerCertificate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbListenerCertificateSpec	`json:"spec"`
}


// AwsLbListenerCertificateSpec is the spec for a AwsLbListenerCertificate Resource
type AwsLbListenerCertificateSpec struct {
	ListenerArn	string	`json:"listener_arn"`
	CertificateArn	string	`json:"certificate_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbListenerCertificateList is a list of AwsLbListenerCertificate resources
type AwsLbListenerCertificateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbListenerCertificate	`json:"items"`
}

