
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
	EcsTarget	[]AwsCloudwatchEventTargetEcsTarget	`json:"ecs_target"`
	BatchTarget	[]AwsCloudwatchEventTargetBatchTarget	`json:"batch_target"`
	KinesisTarget	[]AwsCloudwatchEventTargetKinesisTarget	`json:"kinesis_target"`
	Rule	string	`json:"rule"`
	Arn	string	`json:"arn"`
	Input	string	`json:"input"`
	RunCommandTargets	[]AwsCloudwatchEventTargetRunCommandTargets	`json:"run_command_targets"`
	SqsTarget	[]AwsCloudwatchEventTargetSqsTarget	`json:"sqs_target"`
	InputPath	string	`json:"input_path"`
	RoleArn	string	`json:"role_arn"`
	InputTransformer	[]AwsCloudwatchEventTargetInputTransformer	`json:"input_transformer"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchEventTargetList is a list of AwsCloudwatchEventTarget resources
type AwsCloudwatchEventTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchEventTarget	`json:"items"`
}


// AwsCloudwatchEventTargetEcsTarget is a AwsCloudwatchEventTargetEcsTarget
type AwsCloudwatchEventTargetEcsTarget struct {
	TaskCount	int	`json:"task_count"`
	TaskDefinitionArn	string	`json:"task_definition_arn"`
}

// AwsCloudwatchEventTargetBatchTarget is a AwsCloudwatchEventTargetBatchTarget
type AwsCloudwatchEventTargetBatchTarget struct {
	JobAttempts	int	`json:"job_attempts"`
	JobDefinition	string	`json:"job_definition"`
	JobName	string	`json:"job_name"`
	ArraySize	int	`json:"array_size"`
}

// AwsCloudwatchEventTargetKinesisTarget is a AwsCloudwatchEventTargetKinesisTarget
type AwsCloudwatchEventTargetKinesisTarget struct {
	PartitionKeyPath	string	`json:"partition_key_path"`
}

// AwsCloudwatchEventTargetRunCommandTargets is a AwsCloudwatchEventTargetRunCommandTargets
type AwsCloudwatchEventTargetRunCommandTargets struct {
	Values	[]string	`json:"values"`
	Key	string	`json:"key"`
}

// AwsCloudwatchEventTargetSqsTarget is a AwsCloudwatchEventTargetSqsTarget
type AwsCloudwatchEventTargetSqsTarget struct {
	MessageGroupId	string	`json:"message_group_id"`
}

// AwsCloudwatchEventTargetInputTransformer is a AwsCloudwatchEventTargetInputTransformer
type AwsCloudwatchEventTargetInputTransformer struct {
	InputPaths	map[string]string	`json:"input_paths"`
	InputTemplate	string	`json:"input_template"`
}
