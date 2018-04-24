
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudfrontOriginAccessIdentity struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudfrontOriginAccessIdentitySpec	`json:"spec"`
}

type AwsCloudfrontOriginAccessIdentityList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudfrontOriginAccessIdentity	`json:"items"`
}

type AwsCloudfrontOriginAccessIdentitySpec struct {
	Comment	string	`json:"comment"`
}
