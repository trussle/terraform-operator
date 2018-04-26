
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKeyPair describes a AwsKeyPair resource
type AwsKeyPair struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKeyPairSpec	`json:"spec"`
}


// AwsKeyPairSpec is the spec for a AwsKeyPair Resource
type AwsKeyPairSpec struct {
	PublicKey	string	`json:"public_key"`
	KeyNamePrefix	string	`json:"key_name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKeyPairList is a list of AwsKeyPair resources
type AwsKeyPairList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKeyPair	`json:"items"`
}

