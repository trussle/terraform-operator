
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
	Arn	string	`json:"arn"`
	Input	string	`json:"input"`
	InputPath	string	`json:"input_path"`
	InputTransformer	[]LqLIbAat	`json:"input_transformer"`
	RoleArn	string	`json:"role_arn"`
	RunCommandTargets	[]LYHdaopo	`json:"run_command_targets"`
	EcsTarget	[]vFOkqIex	`json:"ecs_target"`
	BatchTarget	[]sFzXzrlc	`json:"batch_target"`
	KinesisTarget	[]ztxcdJJF	`json:"kinesis_target"`
	SqsTarget	[]uyZHRCov	`json:"sqs_target"`
	Rule	string	`json:"rule"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchEventTargetList is a list of AwsCloudwatchEventTarget resources
type AwsCloudwatchEventTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchEventTarget	`json:"items"`
}


// LqLIbAat is a LqLIbAat
type LqLIbAat struct {
	InputPaths	map[string]???	`json:"input_paths"`
	InputTemplate	string	`json:"input_template"`
}

// LYHdaopo is a LYHdaopo
type LYHdaopo struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}

// vFOkqIex is a vFOkqIex
type vFOkqIex struct {
	TaskCount	int	`json:"task_count"`
	TaskDefinitionArn	string	`json:"task_definition_arn"`
}

// sFzXzrlc is a sFzXzrlc
type sFzXzrlc struct {
	JobDefinition	string	`json:"job_definition"`
	JobName	string	`json:"job_name"`
	ArraySize	int	`json:"array_size"`
	JobAttempts	int	`json:"job_attempts"`
}

// ztxcdJJF is a ztxcdJJF
type ztxcdJJF struct {
	PartitionKeyPath	string	`json:"partition_key_path"`
}

// uyZHRCov is a uyZHRCov
type uyZHRCov struct {
	MessageGroupId	string	`json:"message_group_id"`
}
