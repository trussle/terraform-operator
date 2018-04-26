
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailKeyPair describes a AwsLightsailKeyPair resource
type AwsLightsailKeyPair struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLightsailKeyPairSpec	`json:"spec"`
}


// AwsLightsailKeyPairSpec is the spec for a AwsLightsailKeyPair Resource
type AwsLightsailKeyPairSpec struct {
	NamePrefix	string	`json:"name_prefix"`
	PgpKey	string	`json:"pgp_key"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailKeyPairList is a list of AwsLightsailKeyPair resources
type AwsLightsailKeyPairList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLightsailKeyPair	`json:"items"`
}

