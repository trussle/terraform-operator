
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudfrontOriginAccessIdentity describes a AwsCloudfrontOriginAccessIdentity resource
type AwsCloudfrontOriginAccessIdentity struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudfrontOriginAccessIdentitySpec	`json:"spec"`
}


// AwsCloudfrontOriginAccessIdentitySpec is the spec for a AwsCloudfrontOriginAccessIdentity Resource
type AwsCloudfrontOriginAccessIdentitySpec struct {
	Comment	string	`json:"comment"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudfrontOriginAccessIdentityList is a list of AwsCloudfrontOriginAccessIdentity resources
type AwsCloudfrontOriginAccessIdentityList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudfrontOriginAccessIdentity	`json:"items"`
}

