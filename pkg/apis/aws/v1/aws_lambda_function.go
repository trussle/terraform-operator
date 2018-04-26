
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
	S3Key	string	`json:"s3_key"`
	FunctionName	string	`json:"function_name"`
	ReservedConcurrentExecutions	int	`json:"reserved_concurrent_executions"`
	Role	string	`json:"role"`
	Filename	string	`json:"filename"`
	DeadLetterConfig	[]RVjaRzLN	`json:"dead_letter_config"`
	Publish	bool	`json:"publish"`
	VpcConfig	[]TXYeUCWK	`json:"vpc_config"`
	Tags	map[string]???	`json:"tags"`
	S3Bucket	string	`json:"s3_bucket"`
	Handler	string	`json:"handler"`
	MemorySize	int	`json:"memory_size"`
	Runtime	string	`json:"runtime"`
	Environment	[]mBTvKSJf	`json:"environment"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	S3ObjectVersion	string	`json:"s3_object_version"`
	Description	string	`json:"description"`
	Timeout	int	`json:"timeout"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaFunctionList is a list of AwsLambdaFunction resources
type AwsLambdaFunctionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLambdaFunction	`json:"items"`
}


// RVjaRzLN is a RVjaRzLN
type RVjaRzLN struct {
	TargetArn	string	`json:"target_arn"`
}

// TXYeUCWK is a TXYeUCWK
type TXYeUCWK struct {
	SubnetIds	string	`json:"subnet_ids"`
	SecurityGroupIds	string	`json:"security_group_ids"`
}

// sXbGyRAO is a sXbGyRAO
type sXbGyRAO struct {
	Mode	string	`json:"mode"`
}

// mBTvKSJf is a mBTvKSJf
type mBTvKSJf struct {
	Variables	map[string]???	`json:"variables"`
}
