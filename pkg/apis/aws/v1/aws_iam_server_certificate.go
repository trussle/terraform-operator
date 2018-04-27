
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamServerCertificate describes a AwsIamServerCertificate resource
type AwsIamServerCertificate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamServerCertificateSpec	`json:"spec"`
}


// AwsIamServerCertificateSpec is the spec for a AwsIamServerCertificate Resource
type AwsIamServerCertificateSpec struct {
	CertificateChain	string	`json:"certificate_chain"`
	Path	string	`json:"path"`
	PrivateKey	string	`json:"private_key"`
	NamePrefix	string	`json:"name_prefix"`
	CertificateBody	string	`json:"certificate_body"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamServerCertificateList is a list of AwsIamServerCertificate resources
type AwsIamServerCertificateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamServerCertificate	`json:"items"`
}

