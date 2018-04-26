
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Runtime	string	`json:"runtime"`
	S3ObjectVersion	string	`json:"s3_object_version"`
	FunctionName	string	`json:"function_name"`
	Handler	string	`json:"handler"`
	ReservedConcurrentExecutions	int	`json:"reserved_concurrent_executions"`
	Tags	map[string]string	`json:"tags"`
	S3Bucket	string	`json:"s3_bucket"`
	S3Key	string	`json:"s3_key"`
	Role	string	`json:"role"`
	Timeout	int	`json:"timeout"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	Environment	[]Environment	`json:"environment"`
	Description	string	`json:"description"`
	DeadLetterConfig	[]DeadLetterConfig	`json:"dead_letter_config"`
	Publish	bool	`json:"publish"`
	VpcConfig	[]VpcConfig	`json:"vpc_config"`
	Filename	string	`json:"filename"`
	MemorySize	int	`json:"memory_size"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaFunctionList is a list of AwsLambdaFunction resources
type AwsLambdaFunctionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLambdaFunction	`json:"items"`
}


// Environment is a Environment
type Environment struct {
	Variables	map[string]string	`json:"variables"`
}

// DeadLetterConfig is a DeadLetterConfig
type DeadLetterConfig struct {
	TargetArn	string	`json:"target_arn"`
}

// VpcConfig is a VpcConfig
type VpcConfig struct {
	SubnetIds	string	`json:"subnet_ids"`
	SecurityGroupIds	string	`json:"security_group_ids"`
}

// TracingConfig is a TracingConfig
type TracingConfig struct {
	Mode	string	`json:"mode"`
}
