
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppsyncGraphqlApi describes a AwsAppsyncGraphqlApi resource
type AwsAppsyncGraphqlApi struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAppsyncGraphqlApiSpec	`json:"spec"`
}


// AwsAppsyncGraphqlApiSpec is the spec for a AwsAppsyncGraphqlApi Resource
type AwsAppsyncGraphqlApiSpec struct {
	AuthenticationType	string	`json:"authentication_type"`
	Name	string	`json:"name"`
	UserPoolConfig	[]interface{}	`json:"user_pool_config"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppsyncGraphqlApiList is a list of AwsAppsyncGraphqlApi resources
type AwsAppsyncGraphqlApiList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppsyncGraphqlApi	`json:"items"`
}

