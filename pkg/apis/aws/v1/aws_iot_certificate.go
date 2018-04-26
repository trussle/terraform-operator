
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Active	bool	`json:"active"`
	Csr	string	`json:"csr"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotCertificateList is a list of AwsIotCertificate resources
type AwsIotCertificateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotCertificate	`json:"items"`
}

