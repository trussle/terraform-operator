
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserGroup describes a AwsCognitoUserGroup resource
type AwsCognitoUserGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCognitoUserGroupSpec	`json:"spec"`
}


// AwsCognitoUserGroupSpec is the spec for a AwsCognitoUserGroup Resource
type AwsCognitoUserGroupSpec struct {
	Name	string	`json:"name"`
	Precedence	int	`json:"precedence"`
	RoleArn	string	`json:"role_arn"`
	UserPoolId	string	`json:"user_pool_id"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserGroupList is a list of AwsCognitoUserGroup resources
type AwsCognitoUserGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserGroup	`json:"items"`
}

