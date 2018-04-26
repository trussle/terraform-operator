
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecretsmanagerSecretVersion describes a AwsSecretsmanagerSecretVersion resource
type AwsSecretsmanagerSecretVersion struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSecretsmanagerSecretVersionSpec	`json:"spec"`
}


// AwsSecretsmanagerSecretVersionSpec is the spec for a AwsSecretsmanagerSecretVersion Resource
type AwsSecretsmanagerSecretVersionSpec struct {
	SecretId	string	`json:"secret_id"`
	SecretString	string	`json:"secret_string"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecretsmanagerSecretVersionList is a list of AwsSecretsmanagerSecretVersion resources
type AwsSecretsmanagerSecretVersionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSecretsmanagerSecretVersion	`json:"items"`
}

