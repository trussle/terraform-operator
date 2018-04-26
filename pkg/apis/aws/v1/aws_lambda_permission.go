
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaPermission describes a AwsLambdaPermission resource
type AwsLambdaPermission struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLambdaPermissionSpec	`json:"spec"`
}


// AwsLambdaPermissionSpec is the spec for a AwsLambdaPermission Resource
type AwsLambdaPermissionSpec struct {
	FunctionName	string	`json:"function_name"`
	Principal	string	`json:"principal"`
	Qualifier	string	`json:"qualifier"`
	SourceAccount	string	`json:"source_account"`
	SourceArn	string	`json:"source_arn"`
	StatementId	string	`json:"statement_id"`
	Action	string	`json:"action"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaPermissionList is a list of AwsLambdaPermission resources
type AwsLambdaPermissionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLambdaPermission	`json:"items"`
}

