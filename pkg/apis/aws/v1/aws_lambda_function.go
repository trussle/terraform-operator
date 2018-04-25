
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaFunction describes a AwsLambdaFunction resource
type AwsLambdaFunction struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLambdaFunctionSpec	`json:"spec"`
}


// AwsLambdaFunctionSpec is the spec for a AwsLambdaFunction Resource
type AwsLambdaFunctionSpec struct {
	FunctionName	string	`json:"function_name"`
	MemorySize	int	`json:"memory_size"`
	S3ObjectVersion	string	`json:"s3_object_version"`
	DeadLetterConfig	[]interface{}	`json:"dead_letter_config"`
	Tags	map[string]interface{}	`json:"tags"`
	Role	string	`json:"role"`
	VpcConfig	[]interface{}	`json:"vpc_config"`
	Environment	[]interface{}	`json:"environment"`
	Filename	string	`json:"filename"`
	S3Bucket	string	`json:"s3_bucket"`
	Runtime	string	`json:"runtime"`
	Timeout	int	`json:"timeout"`
	Publish	bool	`json:"publish"`
	S3Key	string	`json:"s3_key"`
	ReservedConcurrentExecutions	int	`json:"reserved_concurrent_executions"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	Description	string	`json:"description"`
	Handler	string	`json:"handler"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaFunctionList is a list of AwsLambdaFunction resources
type AwsLambdaFunctionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLambdaFunction	`json:"items"`
}

