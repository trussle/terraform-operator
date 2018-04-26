
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbListenerCertificate describes a AwsAlbListenerCertificate resource
type AwsAlbListenerCertificate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbListenerCertificateSpec	`json:"spec"`
}


// AwsAlbListenerCertificateSpec is the spec for a AwsAlbListenerCertificate Resource
type AwsAlbListenerCertificateSpec struct {
	ListenerArn	string	`json:"listener_arn"`
	CertificateArn	string	`json:"certificate_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbListenerCertificateList is a list of AwsAlbListenerCertificate resources
type AwsAlbListenerCertificateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbListenerCertificate	`json:"items"`
}

