
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaFunction struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLambdaFunctionSpec	`json:"spec"`
}

type AwsLambdaFunctionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLambdaFunction	`json:"items"`
}

type AwsLambdaFunctionSpec struct {
	Handler	string	`json:"handler"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	Filename	string	`json:"filename"`
	S3Key	string	`json:"s3_key"`
	S3ObjectVersion	string	`json:"s3_object_version"`
	FunctionName	string	`json:"function_name"`
	MemorySize	int	`json:"memory_size"`
	ReservedConcurrentExecutions	int	`json:"reserved_concurrent_executions"`
	Runtime	string	`json:"runtime"`
	DeadLetterConfig	[]interface{}	`json:"dead_letter_config"`
	Timeout	int	`json:"timeout"`
	VpcConfig	[]interface{}	`json:"vpc_config"`
	Environment	[]interface{}	`json:"environment"`
	Tags	map[string]interface{}	`json:"tags"`
	S3Bucket	string	`json:"s3_bucket"`
	Description	string	`json:"description"`
	Role	string	`json:"role"`
	Publish	bool	`json:"publish"`
}
