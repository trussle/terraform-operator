
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaAlias struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLambdaAliasSpec	`json:"spec"`
}

type AwsLambdaAliasList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLambdaAlias	`json:"items"`
}

type AwsLambdaAliasSpec struct {
	Description	string	`json:"description"`
	FunctionName	string	`json:"function_name"`
	FunctionVersion	string	`json:"function_version"`
	Name	string	`json:"name"`
}
