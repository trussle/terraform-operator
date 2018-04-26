
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchEventTarget describes a AwsCloudwatchEventTarget resource
type AwsCloudwatchEventTarget struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchEventTargetSpec	`json:"spec"`
}


// AwsCloudwatchEventTargetSpec is the spec for a AwsCloudwatchEventTarget Resource
type AwsCloudwatchEventTargetSpec struct {
	InputTransformer	[]InputTransformer	`json:"input_transformer"`
	Rule	string	`json:"rule"`
	Input	string	`json:"input"`
	InputPath	string	`json:"input_path"`
	RunCommandTargets	[]RunCommandTargets	`json:"run_command_targets"`
	BatchTarget	[]BatchTarget	`json:"batch_target"`
	KinesisTarget	[]KinesisTarget	`json:"kinesis_target"`
	Arn	string	`json:"arn"`
	RoleArn	string	`json:"role_arn"`
	EcsTarget	[]EcsTarget	`json:"ecs_target"`
	SqsTarget	[]SqsTarget	`json:"sqs_target"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchEventTargetList is a list of AwsCloudwatchEventTarget resources
type AwsCloudwatchEventTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchEventTarget	`json:"items"`
}


// KinesisTarget is a KinesisTarget
type KinesisTarget struct {
	PartitionKeyPath	string	`json:"partition_key_path"`
}

// EcsTarget is a EcsTarget
type EcsTarget struct {
	TaskCount	int	`json:"task_count"`
	TaskDefinitionArn	string	`json:"task_definition_arn"`
}

// SqsTarget is a SqsTarget
type SqsTarget struct {
	MessageGroupId	string	`json:"message_group_id"`
}

// InputTransformer is a InputTransformer
type InputTransformer struct {
	InputPaths	map[string]string	`json:"input_paths"`
	InputTemplate	string	`json:"input_template"`
}

// RunCommandTargets is a RunCommandTargets
type RunCommandTargets struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}

// BatchTarget is a BatchTarget
type BatchTarget struct {
	JobName	string	`json:"job_name"`
	ArraySize	int	`json:"array_size"`
	JobAttempts	int	`json:"job_attempts"`
	JobDefinition	string	`json:"job_definition"`
}
