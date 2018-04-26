
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	RoleMapping	RoleMapping	`json:"role_mapping"`
	Roles	map[string]Roles	`json:"roles"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoIdentityPoolRolesAttachmentList is a list of AwsCognitoIdentityPoolRolesAttachment resources
type AwsCognitoIdentityPoolRolesAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoIdentityPoolRolesAttachment	`json:"items"`
}


// Roles is a Roles
type Roles struct {
	Authenticated	string	`json:"authenticated"`
	Unauthenticated	string	`json:"unauthenticated"`
}

// MappingRule is a MappingRule
type MappingRule struct {
	Claim	string	`json:"claim"`
	MatchType	string	`json:"match_type"`
	RoleArn	string	`json:"role_arn"`
	Value	string	`json:"value"`
}

// RoleMapping is a RoleMapping
type RoleMapping struct {
	IdentityProvider	string	`json:"identity_provider"`
	AmbiguousRoleResolution	string	`json:"ambiguous_role_resolution"`
	MappingRule	[]MappingRule	`json:"mapping_rule"`
	Type	string	`json:"type"`
}
