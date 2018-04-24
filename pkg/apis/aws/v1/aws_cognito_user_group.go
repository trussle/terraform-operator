
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCognitoUserGroupSpec	`json:"spec"`
}

type AwsCognitoUserGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserGroup	`json:"items"`
}

type AwsCognitoUserGroupSpec struct {
	Precedence	int	`json:"precedence"`
	RoleArn	string	`json:"role_arn"`
	UserPoolId	string	`json:"user_pool_id"`
	Description	string	`json:"description"`
	Name	string	`json:"name"`
}
