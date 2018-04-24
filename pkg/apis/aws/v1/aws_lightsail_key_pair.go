
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailKeyPair struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLightsailKeyPairSpec	`json:"spec"`
}

type AwsLightsailKeyPairList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLightsailKeyPair	`json:"items"`
}

type AwsLightsailKeyPairSpec struct {
	NamePrefix	string	`json:"name_prefix"`
	PgpKey	string	`json:"pgp_key"`
}
