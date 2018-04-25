
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoIdentityPoolRolesAttachment describes a AwsCognitoIdentityPoolRolesAttachment resource
type AwsCognitoIdentityPoolRolesAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCognitoIdentityPoolRolesAttachmentSpec	`json:"spec"`
}


// AwsCognitoIdentityPoolRolesAttachmentSpec is the spec for a AwsCognitoIdentityPoolRolesAttachment Resource
type AwsCognitoIdentityPoolRolesAttachmentSpec struct {
	IdentityPoolId	string	`json:"identity_pool_id"`
	RoleMapping	string	`json:"role_mapping"`
	Roles	map[string]interface{}	`json:"roles"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoIdentityPoolRolesAttachmentList is a list of AwsCognitoIdentityPoolRolesAttachment resources
type AwsCognitoIdentityPoolRolesAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoIdentityPoolRolesAttachment	`json:"items"`
}

