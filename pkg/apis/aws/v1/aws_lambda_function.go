
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
	Environment	[]AwsLambdaFunctionEnvironment	`json:"environment"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	Filename	string	`json:"filename"`
	S3ObjectVersion	string	`json:"s3_object_version"`
	Description	string	`json:"description"`
	Runtime	string	`json:"runtime"`
	Tags	map[string]string	`json:"tags"`
	ReservedConcurrentExecutions	int	`json:"reserved_concurrent_executions"`
	Role	string	`json:"role"`
	Publish	bool	`json:"publish"`
	S3Bucket	string	`json:"s3_bucket"`
	DeadLetterConfig	[]AwsLambdaFunctionDeadLetterConfig	`json:"dead_letter_config"`
	Handler	string	`json:"handler"`
	MemorySize	int	`json:"memory_size"`
	FunctionName	string	`json:"function_name"`
	S3Key	string	`json:"s3_key"`
	Timeout	int	`json:"timeout"`
	VpcConfig	[]AwsLambdaFunctionVpcConfig	`json:"vpc_config"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaFunctionList is a list of AwsLambdaFunction resources
type AwsLambdaFunctionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLambdaFunction	`json:"items"`
}


// AwsLambdaFunctionEnvironment is a AwsLambdaFunctionEnvironment
type AwsLambdaFunctionEnvironment struct {
	Variables	map[string]string	`json:"variables"`
}

// AwsLambdaFunctionDeadLetterConfig is a AwsLambdaFunctionDeadLetterConfig
type AwsLambdaFunctionDeadLetterConfig struct {
	TargetArn	string	`json:"target_arn"`
}

// AwsLambdaFunctionTracingConfig is a AwsLambdaFunctionTracingConfig
type AwsLambdaFunctionTracingConfig struct {
	Mode	string	`json:"mode"`
}

// AwsLambdaFunctionVpcConfig is a AwsLambdaFunctionVpcConfig
type AwsLambdaFunctionVpcConfig struct {
	SubnetIds	string	`json:"subnet_ids"`
	SecurityGroupIds	string	`json:"security_group_ids"`
}
