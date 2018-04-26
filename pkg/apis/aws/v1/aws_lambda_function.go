
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
	ReservedConcurrentExecutions	int	`json:"reserved_concurrent_executions"`
	Publish	bool	`json:"publish"`
	S3ObjectVersion	string	`json:"s3_object_version"`
	DeadLetterConfig	[]Generic	`json:"dead_letter_config"`
	FunctionName	string	`json:"function_name"`
	Handler	string	`json:"handler"`
	MemorySize	int	`json:"memory_size"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	Timeout	int	`json:"timeout"`
	Tags	map[string]Generic	`json:"tags"`
	Filename	string	`json:"filename"`
	Description	string	`json:"description"`
	Role	string	`json:"role"`
	VpcConfig	[]Generic	`json:"vpc_config"`
	S3Bucket	string	`json:"s3_bucket"`
	S3Key	string	`json:"s3_key"`
	Runtime	string	`json:"runtime"`
	Environment	[]Generic	`json:"environment"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaFunctionList is a list of AwsLambdaFunction resources
type AwsLambdaFunctionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLambdaFunction	`json:"items"`
}

